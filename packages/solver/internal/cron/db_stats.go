package cron

import (
	"log"
	"solver/internal/database"
)

func LogDBStats() {
	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Printf("Error getting DB stats: %v", err)
		return
	}

	stats := sqlDB.Stats()
	log.Printf("DB Stats: %+v", stats)
}
