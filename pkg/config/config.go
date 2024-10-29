package config

import (
	"ibercs/pkg/logger"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Database       DatabaseConfig
		FaceitApiToken string
		Workers        WorkersConfig
		Microservices  MicroservicesConfig
	}

	DatabaseConfig struct {
		Host     string
		DbName   string
		Port     string
		User     string
		Password string
	}

	WorkersConfig struct {
		PlayersHost string
	}

	MicroservicesConfig struct {
		PlayersHost     string
		PlayersPort     string
		TeamsHost       string
		TeamsPort       string
		UserHost        string
		UserPort        string
		TournamentsHost string
		TournamentsPort string
	}
)

func Load() (Config, error) {
	logger.Debug("Loading config...")

	env := os.Getenv("ENV")
	if env == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error al cargar el archivo .env")
		}
	}

	config := Config{
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			DbName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		FaceitApiToken: os.Getenv("FACEIT_API_TOKEN"),
		Workers: WorkersConfig{
			PlayersHost: os.Getenv("WORKER_PLAYERS_HOST"),
		},
		Microservices: MicroservicesConfig{
			PlayersHost:     os.Getenv("MICROSERVICE_PLAYERS_HOST"),
			PlayersPort:     os.Getenv("MICROSERVICE_PLAYERS_PORT"),
			TeamsHost:       os.Getenv("MICROSERVICE_TEAMS_HOST"),
			TeamsPort:       os.Getenv("MICROSERVICE_TEAMS_PORT"),
			UserHost:        os.Getenv("MICROSERVICE_USERS_HOST"),
			UserPort:        os.Getenv("MICROSERVICE_USERS_PORT"),
			TournamentsHost: os.Getenv("MICROSERVICE_TOURNAMENTS_HOST"),
			TournamentsPort: os.Getenv("MICROSERVICE_TOURNAMENTS_PORT"),
		},
	}

	logger.Debug("Config loaded successfully")
	return config, nil
}
