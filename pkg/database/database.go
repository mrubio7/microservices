package database

import (
	"fmt"
	"ibercs/internal/model"
	"ibercs/pkg/config"
	"ibercs/pkg/logger"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	lgr "gorm.io/gorm/logger"
)

type Database struct {
	db     *gorm.DB
	scheme string
}

func NewDatabase(cfg config.DatabaseConfig) *Database {
	logger.Debug("Initializing database...")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Europe/Madrid sslmode=disable search_path=%s", cfg.Host, cfg.User, cfg.Password, cfg.DbName, cfg.Port, cfg.Scheme)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:      lgr.Default.LogMode(lgr.Silent),
		PrepareStmt: false,
	})
	if err != nil {
		logger.Error("Failed to connect to database: %v", err)
		return nil
	}

	logger.Debug("Database initialized")

	database := &Database{db: db, scheme: cfg.Scheme}

	database.db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", database.scheme))

	database.Automigrate()

	return database
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}

func (d *Database) Automigrate() {
	// Crear el esquema si no existe
	if err := d.db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", d.scheme)).Error; err != nil {
		log.Fatalf("Failed to create schema '%s': %v", d.scheme, err)
	}

	// Migrar los modelos
	if err := d.db.AutoMigrate(&model.MatchModel{}); err != nil {
		log.Fatalf("Failed to automigrate: %v", err)
	}
}
