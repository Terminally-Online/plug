package cron

import (
	"context"
	"encoding/json"
	"log"
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

	if !config.UseAVS {
		return submitToMempool(solution)
	}

	return submitToAVS(solution)
}

func processIntent(s *solver.Solver, intent *models.Intent) {
	intent.PeriodEndAt, intent.NextSimulationAt = intent.GetNextSimulationAt()
	if err := database.DB.Model(&intent).Updates(map[string]any{
		"period_end_at":      intent.PeriodEndAt,
		"next_simulation_at": intent.NextSimulationAt,
	}).Error; err != nil {
		log.Printf("failed to update intent simulation interval: %v", err)
	}

	if intent.Locked {
		// TODO: Need to actually simulate this before kicking it off.
		if err := submit(s.RebuildSolutionFromModels(intent)); err != nil {
			log.Printf("failed to submit intent execution task to avs")
		}
	} else if err := submit(s.Solve(intent, true, true)); err != nil {
		log.Printf("failed to submit intent execution task to avs")
	} else {
		log.Printf("failed to simulate intent")
	}
}

func Simulations(s *solver.Solver) {
	if s.IsKilled {
		log.Println("The solver is currently killed. Skipping simulations.")
		return
	}

	var intents []models.Intent
	query := database.DB.
		Where(
			"next_simulation_at <= ? AND status = ? AND saved = ?",
			time.Now(), "active", true,
		).
		Order("next_simulation_at asc").
		Find(&intents)
	if query.Error != nil {
		log.Printf("database error: %v", query.Error.Error())
		return
	}

	const maxWorkers = 200
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
}
