package cron

import (
	"log"
	"solver/internal/client"
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
				log.Printf("failed to rebuild existing solution: %v", err)
			}
			livePlugs[intent.ChainId][intent.Id] = *solution.LivePlugs
		} else {
			if solution, err := s.Solve(&intent, !production, true); err == nil {
				livePlugs[intent.ChainId][intent.Id] = *solution.LivePlugs
			} else {
				log.Printf("failed to simulate: %v", err)
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

	if !production {
		log.Println("Solver is not in production so will not be executing transactions.")
		return
	}

	for chainId, chainPlugs := range livePlugs {
		client, err := client.New(chainId)
		if err != nil {
			log.Printf("failed to create client for chain %d: %v", chainId, err)
			continue
		}

		_, err = client.Plug(chainPlugs)
		if err != nil {
			log.Printf("failed to simulate plugs for chain %d: %v", chainId, err)
			continue
		}
	}
}
