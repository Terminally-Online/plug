package cron

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"solver/internal/avs/config"
	"solver/internal/avs/streams"
	"solver/internal/client"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/solver"
	"sync"
	"time"
)

func submitToAVS(solution *solver.Solution) error {
	if len(solution.LivePlugs.Plugs.Plugs) == 0 || solution.Run.Status != "success" {
		return nil
	}

	livePlugsData, err := json.Marshal(solution.LivePlugs)
	if err != nil {
		log.Printf("Error marshaling Plugs: %v", err)
		return err
	}

	data := map[string]any{
		"id":        solution.IntentId,
		"chainId":   solution.LivePlugs.ChainId,
		"from":      solution.LivePlugs.From,
		"timestamp": time.Now().Unix(),
		"plugs":     string(livePlugsData),
	}

	ctx := context.Background()
	if err := streams.Publish(ctx, solution.IntentId, data); err != nil {
		log.Printf("Error publishing to Circuit stream: %v", err)
		return err
	}

	log.Printf("Published intent %s to Circuit stream", solution.IntentId)
	return nil
}

func submitToMempool(solution *solver.Solution) error {
	client, err := client.New(solution.LivePlugs.ChainId)
	if err != nil {
		return err
	}

	_, err = client.Plug(solution.LivePlugs)
	return err
}

func submit(solution *solver.Solution, err error) error {
	if err != nil {
		return err
	}

	if !config.UseExecution {
		return nil
	}

	var submitFunc func(*solver.Solution) error
	switch config.UseAVS {
	case false:
		submitFunc = submitToMempool
	case true:
		submitFunc = submitToAVS
	}
	return submitFunc(solution)
}

func processIntent(s *solver.Solver, intent *models.Intent) {
	if intent.Locked {
		// TODO: Need to actually simulate this before kicking it off when using
		//       an AVS submission mechanism instead of local execution.
		if err := submit(s.RebuildSolutionFromModels(intent)); err != nil {
			log.Printf("failed to submit locked intent")
		}
	} else if err := submit(s.Solve(intent, true, true)); err != nil {
		log.Printf("failed to submit intent")
	}

	intent.Status, intent.PeriodEndAt, intent.NextSimulationAt = intent.GetNextSimulationAt()
	if err := database.DB.Model(&intent).Updates(map[string]any{
		"status":             intent.Status,
		"period_end_at":      intent.PeriodEndAt,
		"next_simulation_at": intent.NextSimulationAt,
	}).Error; err != nil {
		log.Printf("failed to update intent simulation interval: %v", err)
	}
}

func Simulations(s *solver.Solver) {
	if s.IsKilled {
		log.Println("The solver is currently killed. Skipping simulations.")
		return
	}

	var intentIds []string
	err := database.DB.Model(&models.Intent{}).
		Where("next_simulation_at <= ? AND status = ? AND saved = ?", time.Now(), "active", true).
		Order("next_simulation_at asc").
		Pluck("id", &intentIds).Error
	if err != nil || len(intentIds) == 0 {
		log.Printf("did not find any intents to simulate: %v", err)
		return
	}

	if err := database.DB.Model(&models.Intent{}).
		Where("id IN ?", intentIds).
		Update("status", "running").Error; err != nil {
		log.Printf("database error updating intents to running: %v", err)
		return
	}

	var intents []models.Intent
	if err := database.DB.Where("id IN ?", intentIds).Find(&intents).Error; err != nil {
		log.Printf("database error retrieving updated intents: %v", err)
		return
	}

	maxWorkers := int64(math.Min(float64(200), float64(len(intents))))
	intentChan := make(chan models.Intent)
	var wg sync.WaitGroup

	for range maxWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for intent := range intentChan {
				processIntent(s, &intent)
			}
		}()
	}

	for _, intent := range intents {
		intentChan <- intent
	}

	close(intentChan)
	wg.Wait()

	// NOTE: This is here to catch anything that errored out in the middle of processing and never
	//       got to an appropriate status update to bring it out of the "running" state. Even though
	//       we did catch it here something happened in our system above and it is marked as a critical
	//       failure that will draw attention from the team to resolve in future instances.
	err = database.DB.Model(&models.Intent{}).
		Where("id IN ? AND status = ?", intentIds, "running").
		Update("status", "failure").Error
	if err != nil {
		log.Printf("database error retrieving updated intents: %v", err)
		return
	}
}
