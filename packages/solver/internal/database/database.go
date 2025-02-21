package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"solver/internal/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

func NewConfig() *Config {
	return &Config{
		Host:     GetEnvOrDefault("DB_HOST", "localhost"),
		User:     GetEnvOrDefault("DB_USER", "plug"),
		Password: GetEnvOrDefault("DB_PASSWORD", "plugdev"),
		DBName:   GetEnvOrDefault("DB_NAME", "plug_solver"),
		Port:     GetEnvOrDefault("DB_PORT", "6432"),
		SSLMode:  GetEnvOrDefault("DB_SSLMODE", "disable"),
	}
}

func Connect(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
		config.SSLMode,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func GetEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func init() {
	var err error
	DB, err = Initialize()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize database: %v", err))
	}
}

func Initialize() (*gorm.DB, error) {
	config := NewConfig()
	db, err := Connect(config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := AutoMigrate(db); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.SimulationRequest{},
		&models.SimulationResponse{},
	)
}
