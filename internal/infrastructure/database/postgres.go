package database

import (
	"fmt"
	"log"

	"github.com/arulkarim/golden-architecture/configs"
	"github.com/arulkarim/golden-architecture/internal/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewPostgresDB creates a new PostgreSQL database connection using GORM
func NewPostgresDB(cfg *configs.DatabaseConfig) (*gorm.DB, error) {
	dsn := cfg.DSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Get underlying sql.DB to configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	log.Println("Database connected successfully")
	return db, nil
}

// AutoMigrate runs auto migration for all entities
func AutoMigrate(db *gorm.DB) error {
	log.Println("Running auto migration...")
	return db.AutoMigrate(
		&entity.Todo{},
		&entity.User{},
	)
}
