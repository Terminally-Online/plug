package cron

import (
	"log"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/solver/signature"
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

	if len(intents) == 0 {
		return nil
	}

	// For each intent, delete its related records and then the intent itself
	deletedCount := 0
	for _, intent := range intents {
		err := db.Transaction(func(tx *gorm.DB) error {
			// First find all LivePlugs associated with this intent
			var livePlugs []signature.LivePlugs
			if err := tx.Where("intent_id = ?", intent.Id).Find(&livePlugs).Error; err != nil {
				return err
			}

			// For each LivePlugs, first delete any Runs that reference it
			for _, livePlug := range livePlugs {
				if err := tx.Where("live_plugs_id = ?", livePlug.Id).Delete(&models.Run{}).Error; err != nil {
					return err
				}
			}

			// Then delete any remaining Runs associated with the intent
			if err := tx.Where("intent_id = ? AND live_plugs_id IS NULL", intent.Id).Delete(&models.Run{}).Error; err != nil {
				return err
			}

			// Now it's safe to delete the LivePlugs
			if err := tx.Where("intent_id = ?", intent.Id).Delete(&signature.LivePlugs{}).Error; err != nil {
				return err
			}

			// Finally, delete the Intent
			if err := tx.Delete(&intent).Error; err != nil {
				return err
			}
			log.Printf("Successfully deleted intent %s", intent.Id)

			return nil
		})

		if err != nil {
			log.Printf("Error cleaning up intent %s: %v", intent.Id, err)
			continue
		}
		deletedCount++
	}

	return nil
}

// IntentCleanup is meant to be called directly by the cron scheduler
func IntentCleanup(olderThan time.Duration) {
	if err := CleanupUnusedIntents(database.DB, olderThan); err != nil {
		log.Printf("Error in cleanup job: %v", err)
	}
}
