package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"solver/internal/database/models"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/joho/godotenv"
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
		Host:     utils.GetEnvOrDefault("DATABASE_HOST", "localhost"),
		User:     utils.GetEnvOrDefault("DATABASE_USER", "plug"),
		Password: utils.GetEnvOrDefault("DATABASE_PASSWORD", "plugdev"),
		DBName:   utils.GetEnvOrDefault("DATABASE_NAME", "plug_solver"),
		Port:     utils.GetEnvOrDefault("DATABASE_PORT", "6432"),
		SSLMode:  utils.GetEnvOrDefault("DATABASE_SSLMODE", "disable"),
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
			SlowThreshold:             500 * time.Millisecond,
			LogLevel:                  logger.Error,
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

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

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

	if err := Seed(db); err != nil {
		return nil, fmt.Errorf("failed to seed: %w", err)
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.ApiKey{},
		&models.Run{},
		&models.Intent{},
		&signature.LivePlugs{},
	)
}

func Seed(db *gorm.DB) error {
	appApiKey := os.Getenv("ADMIN_API_KEY")
	if appApiKey == "" {
		panic("ADMIN_API_KEY environment variable is not set")
	}

	dev := "0x0Bb5d848487B10F8CFBa21493c8f6D47e8a8B17E"
	var existingApiKey1 models.ApiKey
	if err := db.Where("socket_id = ?", dev).First(&existingApiKey1).Error; err == gorm.ErrRecordNotFound {
		apiKey := models.ApiKey{
			SocketId:  &dev,
			Key:       appApiKey,
			Role:      "admin",
			RateLimit: 1_000_000,
		}
		if err := db.Create(&apiKey).Error; err != nil {
			return utils.ErrBuild("failed to save first api key: " + err.Error())
		}
	}

	demo := "0x62180042606624f02d8a130da8a3171e9b33894d"
	var existingApiKey2 models.ApiKey
	if err := db.Where("socket_id = ?", demo).First(&existingApiKey2).Error; err == gorm.ErrRecordNotFound {
		apiKey := models.ApiKey{
			SocketId:  &demo,
			Key:       "bingbopboombam",
			RateLimit: 10,
		}
		if err := db.Create(&apiKey).Error; err != nil {
			return utils.ErrBuild("failed to save second api key: " + err.Error())
		}
	}

	testing := "0x50701f4f523766bFb5C195F93333107d1cB8cD90"
	var existingApiKey3 models.ApiKey
	if err := db.Where("socket_id = ?", testing).First(&existingApiKey3).Error; err == gorm.ErrRecordNotFound {
		apiKey := models.ApiKey{
			SocketId:  &testing,
			Key:       "testing",
			RateLimit: 1000,
			Role:      "admin",
		}
		if err := db.Create(&apiKey).Error; err != nil {
			return utils.ErrBuild("failed to save third api key: " + err.Error())
		}
	}

	return nil
}
