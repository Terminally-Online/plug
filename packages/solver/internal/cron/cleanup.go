package cron

import (
	"log"
	"solver/internal/database"
	"solver/internal/database/models"
	"time"

	"gorm.io/gorm"
)

// CleanupUnusedIntents deletes intents that are older than the specified duration and have not been saved
func CleanupUnusedIntents(db *gorm.DB, olderThan time.Duration) error {
	cutoffTime := time.Now().Add(-olderThan)

	// Find all unsaved intents older than the cutoff time
	var intents []models.Intent
	if err := db.Where("saved = ? AND created_at < ?", false, cutoffTime).Find(&intents).Error; err != nil {
		return err
	}

	// For each intent, delete its related records and then the intent itself
	for _, intent := range intents {
		err := db.Transaction(func(tx *gorm.DB) error {
			// Delete related LivePlugs and their Plugs
			var livePlugs []models.LivePlug
			if err := tx.Where("intent_id = ?", intent.Id).Find(&livePlugs).Error; err != nil {
				return err
			}

			for _, livePlug := range livePlugs {
				// Delete related Plugs
				if err := tx.Where("bundle_id = ?", livePlug.Id).Delete(&models.Plug{}).Error; err != nil {
					return err
				}
			}

			// Delete LivePlugs
			if err := tx.Where("intent_id = ?", intent.Id).Delete(&models.LivePlug{}).Error; err != nil {
				return err
			}

			// Delete related Runs
			if err := tx.Where("intent_id = ?", intent.Id).Delete(&models.Run{}).Error; err != nil {
				return err
			}

			// Finally, delete the Intent
			if err := tx.Delete(&intent).Error; err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			log.Printf("Error cleaning up intent %s: %v", intent.Id, err)
			continue
		}
	}

	return nil
}

func StartIntentCleanupJob(cleanupInterval, olderThan time.Duration) {
	go func() {
		ticker := time.NewTicker(cleanupInterval)
		defer ticker.Stop()

		for range ticker.C {
			if err := CleanupUnusedIntents(database.DB, olderThan); err != nil {
				log.Printf("Error in cleanup job: %v", err)
			}
		}
	}()
}
