package cron

import (
	"log"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/solver"
	"solver/internal/solver/signature"
	"time"
)

func Simulations(s *solver.Solver) {
	if s.IsKilled {
		return
	}

	var intents []models.Intent
	query := database.DB.
		Where("next_simulation_at <= ? AND status = ? AND saved = ?", time.Now(), "active", true).
		Order("next_simulation_at asc").
		Find(&intents)
	if query.Error != nil {
		log.Printf("database error: %v", query.Error.Error())
		return
	}

	livePlugs := make(map[string]*signature.LivePlugs, 0)
	for _, intent := range intents {
		if intent.Locked {
			solution, err := s.RebuildSolutionFromModels(&intent)
			if err != nil {
				continue
			}
			livePlugs[intent.Id] = solution.LivePlugs
		} else {
			if solution, err := s.Solve(&intent, false, true); err == nil {
				livePlugs[intent.Id] = solution.LivePlugs
			} else {
				log.Printf("failed to simulation: %v", err)
			}
		}
	}

	// TODO: Enable execution with culling simulation.
}
