package cron

import (
	"log"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/solver"
	"time"
)

func Simulations(s solver.Solver) {
	if s.IsKilled {
		return
	}

	var intents []models.Intent
	query := database.DB.
		Where("next_simulation_at <= ? AND status = ?", time.Now(), "active").
		Order("next_simulation_at asc").
		Find(&intents)
	if query.Error != nil {
		log.Printf("database error: %v", query.Error.Error())
		return
	}

	solutions := make([]solver.Solution, len(intents))
	for index, intent := range intents {
		if solution, err := s.Solve(&intent); err != nil {
			log.Printf("failed to simulation: %v", err)
		} else {
			solutions[index] = *solution
		}
	}

	// TODO: Enable execution.
}
