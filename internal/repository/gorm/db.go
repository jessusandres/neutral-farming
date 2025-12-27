package gorm

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"looker.com/neutral-farming/internal/config"
)

func NewDB() *gorm.DB {
	log.Println("🔨 Building pool...")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.EnvConfig.DBHost,
		config.EnvConfig.DBPort,
		config.EnvConfig.DBUsername,
		config.EnvConfig.DBPassword,
		config.EnvConfig.DBName,
		config.EnvConfig.DBSSLMode,
	)

	var db *gorm.DB
	var err error

	log.Println("⚙️  Using default postgres driver")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic("Failed to get db object")
	}

	sqlDB.SetMaxIdleConns(config.EnvConfig.DBMinConnections)

	sqlDB.SetMaxOpenConns(config.EnvConfig.DBMaxConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Minute * 3)

	if !config.EnvConfig.DBLogger {
		db.Logger = logger.Default.LogMode(logger.Silent)
	} else {
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	return db
}
