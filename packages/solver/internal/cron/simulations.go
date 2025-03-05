package cron

import (
	"fmt"
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
		if intent.Saved {
			solution, err := s.RebuildSolutionFromModels(&intent)

			fmt.Printf("Solution Status: %+v\n", solution.Status)
			if solution.Transactions != nil {
				fmt.Printf("Transactions: %+v\n", *solution.Transactions)
			}
			if solution.LivePlugs != nil {
				fmt.Printf("LivePlugs: %+v\n", *solution.LivePlugs)
			}
			if solution.Intent != nil {
				fmt.Printf("Intent: %+v\n", *solution.Intent)
			}
			if solution.Run != nil {
				fmt.Printf("Run: %+v\n", *solution.Run)
			}
			if solution.Transaction != nil {
				fmt.Printf("Transaction: %+v\n", *solution.Transaction)
			}

			if err != nil {
				solutions[index] = solver.Solution{
					Status: solver.SolutionStatus{
						Success: false,
						Error:   err.Error(),
					},
				}
				continue
			}
			solutions[index] = *solution
		} else if solution, err := s.Solve(&intent); err != nil {
			solutions[index] = solver.Solution{
				Status: solver.SolutionStatus{
					Success: false,
					Error:   err.Error(),
				},
			}
		} else {
			solutions[index] = *solution
		}
	}

	// TODO: Enable execution.
}
