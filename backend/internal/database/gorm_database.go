package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/shafikshaon/url_shortener/config"
	"github.com/shafikshaon/url_shortener/internal/logger"
	"github.com/shafikshaon/url_shortener/internal/models"
)

// GormDatabase wraps the GORM database connection
type GormDatabase struct {
	DB *gorm.DB
}

// NewGormDatabase creates a new GORM database connection
func NewGormDatabase(cfg *config.Config) (*GormDatabase, error) {
	dsn := cfg.GetDSN()

	// Configure GORM logger
	gormConfig := &gorm.Config{
		Logger: gormlogger.New(
			log.New(log.Writer(), "\r\n", log.LstdFlags),
			gormlogger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  gormlogger.Warn,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Get underlying SQL database for connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// Test connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	ctx := context.Background()
	logger.Infof(ctx, "GORM database connection established successfully")

	return &GormDatabase{DB: db}, nil
}

// AutoMigrate runs the automatic migrations for all models
func (d *GormDatabase) AutoMigrate() error {
	ctx := context.Background()
	logger.Infof(ctx, "Starting database auto-migration...")

	err := d.DB.AutoMigrate(
		&models.User{},
		&models.Link{},
		&models.Click{},
	)

	if err != nil {
		logger.Errorf(ctx, "Failed to run auto-migration: %v", err)
		return fmt.Errorf("failed to run auto-migration: %w", err)
	}

	logger.Infof(ctx, "Database auto-migration completed successfully")
	return nil
}

// Close closes the database connection
func (d *GormDatabase) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// Ping checks the database connection
func (d *GormDatabase) Ping() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
