package database

import (
	"fmt"
	"ibercs/internal/model"
	"ibercs/pkg/config"

	"ibercs/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	lgr "gorm.io/gorm/logger"
)

func New(cfg config.DatabaseConfig) *gorm.DB {
	logger.Debug("Initializing database...")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Madrid prefer_simple_protocol=true", cfg.Host, cfg.User, cfg.Password, cfg.DbName, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:      lgr.Default.LogMode(lgr.Silent),
		PrepareStmt: false,
	})
	if err != nil {
		return nil
	}

	logger.Debug("Database initialized")
	migrateTables(db)
	return db
}

func migrateTables(db *gorm.DB) {
	logger.Trace("Updating database tables")
	db.AutoMigrate(&model.PlayerModel{})
	db.AutoMigrate(&model.PlayerStatsModel{})
	db.AutoMigrate(&model.PlayerProminentModel{})
	db.AutoMigrate(&model.ProminentWeekModel{})
	db.AutoMigrate(&model.TeamModel{})
}
