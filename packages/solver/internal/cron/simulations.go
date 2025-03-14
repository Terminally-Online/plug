package cron

import (
	"log"
	"solver/internal/bindings/references"
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

	chainIds := []uint64{references.Base.ChainIds[0]}
	for _, chainId := range chainIds {
		var intents []models.Intent
		query := database.DB.
			Where(
				"next_simulation_at <= ? AND status = ? AND saved = ? AND chain_id = ?",
				time.Now(), "active", true, chainId,
			).
			Order("next_simulation_at asc").
			Find(&intents)
		if query.Error != nil {
			log.Printf("database error: %v", query.Error.Error())
			return
		}

		livePlugs := make(map[string]signature.LivePlugs, 0)
		for _, intent := range intents {
			if intent.Locked {
				solution, err := s.RebuildSolutionFromModels(&intent)
				if err != nil {
					continue
				}
				livePlugs[intent.Id] = *solution.LivePlugs
			} else {
				if solution, err := s.Solve(&intent, false, true); err == nil {
					livePlugs[intent.Id] = *solution.LivePlugs
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

		// TODO: Uncomment this to enable execution of plugs run through simulation.

		// if len(livePlugs) == 0 {
		// 	return
		// }

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
}
