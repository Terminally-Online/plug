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

	intentIds := make([]string, len(intents))
	for i, intent := range intents {
		intentIds[i] = intent.Id
	}

	// Process all intents in a single transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		// First, find all LivePlugs IDs for these intents in one query
		var livePlugsIds []string
		if err := tx.Model(&signature.LivePlugs{}).
			Where("intent_id IN ?", intentIds).
			Pluck("id", &livePlugsIds).Error; err != nil {
			return err
		}

		// Delete runs in a single operation - delete all that reference these liveplugs
		if len(livePlugsIds) > 0 {
			if err := tx.Unscoped().
				Where("live_plugs_id IN ?", livePlugsIds).
				Delete(&models.Run{}).Error; err != nil {
				return err
			}
		}

		// Delete all runs associated with these intents in one operation
		if err := tx.Unscoped().
			Where("intent_id IN ?", intentIds).
			Delete(&models.Run{}).Error; err != nil {
			return err
		}

		// Delete all LivePlugs for these intents in one operation
		if err := tx.Unscoped().
			Where("intent_id IN ?", intentIds).
			Delete(&signature.LivePlugs{}).Error; err != nil {
			return err
		}

		// Finally, delete the intents themselves
		if err := tx.Unscoped().
			Where("id IN ?", intentIds).
			Delete(&models.Intent{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Printf("Error cleaning up intents: %v", err)
		return err
	}

	log.Printf("Cleanup complete: deleted %d unsaved intents", len(intentIds))
	return nil
}

// IntentCleanup is meant to be called directly by the cron scheduler
func IntentCleanup(olderThan time.Duration) {
	if err := CleanupUnusedIntents(database.DB, olderThan); err != nil {
		log.Printf("Error in cleanup job: %v", err)
	}
}
