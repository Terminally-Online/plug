package cron

import (
	"log"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/solver"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"time"
)

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

	production := utils.GetEnvOrDefault("SOLVER_ENV", "development") == "production" 

	livePlugs := make(map[uint64]map[string]signature.LivePlugs, 0)
	for _, intent := range intents {
		if intent.Locked {
			solution, err := s.RebuildSolutionFromModels(&intent)
			if err != nil {
				continue
			}
			livePlugs[intent.ChainId][intent.Id] = *solution.LivePlugs
		} else {
			if solution, err := s.Solve(&intent, !production, true); err == nil {
				livePlugs[intent.ChainId][intent.Id] = *solution.LivePlugs
			} else {
				log.Printf("failed to simulation: %v", err)
			}
		}

		intent.PeriodEndAt, intent.NextSimulationAt = intent.GetNextSimulationAt()
		if err := database.DB.Model(&intent).Updates(map[string]any{
			"period_end_at":      intent.PeriodEndAt,
			"next_simulation_at": intent.NextSimulationAt,
		}).Error; err != nil {
			log.Printf("failed to update intent simulation interval: %v", err)
		}
	}

	if len(livePlugs) == 0 {
		log.Println("Simulations ran and there is nothing to execute.")
		return
	}

	if !production {
		log.Println("Solver is not in production so will not be executing transactions.")
		return
	}

	// client, err := client.New(chainId)
	// if err != nil {
	// 	log.Printf("failed to create client: %v", err)
	// 	continue
	// }

	// results, err := client.Plug(livePlugs)
	// if err != nil {
	// 	log.Printf("failed to simulate plugs: %v", err)
	// 	continue
	// }
}
