package cron

import (
	"context"
	"encoding/json"
	"log"
	"solver/internal/avs/streams"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/solver"
	"time"
)

func SubmitAVSTask(solution *solver.Solution, err error) error {
	if err != nil {
		return err
	}

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

	for _, intent := range intents {
		intent.PeriodEndAt, intent.NextSimulationAt = intent.GetNextSimulationAt()
		if err := database.DB.Model(&intent).Updates(map[string]any{
			"period_end_at":      intent.PeriodEndAt,
			"next_simulation_at": intent.NextSimulationAt,
		}).Error; err != nil {
			log.Printf("failed to update intent simulation interval: %v", err)
		}

		if intent.Locked {
			// TODO: Need to actually simulate this before kicking it off.
			if err := SubmitAVSTask(s.RebuildSolutionFromModels(&intent)); err != nil {
				log.Printf("failed to submit intent execution task to avs")
			}
		} else if err := SubmitAVSTask(s.Solve(&intent, true, true)); err != nil {
			log.Printf("failed to submit intent execution task to avs")
		} else {
			log.Printf("failed to simulate intent")
		}
	}
}
