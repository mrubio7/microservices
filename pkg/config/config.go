package config

import (
	"context"
	"ibercs/pkg/logger"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/testcontainers/testcontainers-go"
)

type (
	Config struct {
		Database       DatabaseConfig
		FaceitApiToken string
		Workers        WorkersConfig
		Microservices  MicroservicesConfig
	}

	ConfigV2 struct {
		ThirdPartyApiTokens     ThirdPartyApiTokens
		MicroservicePlayers     MicroserviceConfig
		MicroserviceTeams       MicroserviceConfig
		MicroserviceUsers       MicroserviceConfig
		MicroserviceTournaments MicroserviceConfig
		MicroserviceMatches     MicroserviceConfig
	}

	ThirdPartyApiTokens struct {
		FaceitApiToken string
	}

	DatabaseConfig struct {
		Host     string
		DbName   string
		Port     string
		User     string
		Password string
		Scheme   string
	}

	WorkersConfig struct {
		PlayersHost string
	}

	MicroserviceConfig struct {
		Database  DatabaseConfig
		Host_gRPC string
		Port_gRPC string
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
		MatchesHost     string
		MatchesPort     string
	}
)

func LoadTestDatabaseConfig(ctx context.Context, testContainer testcontainers.Container, scheme string) DatabaseConfig {
	host, err := testContainer.Host(ctx)
	if err != nil {
		log.Fatalf("Failed to get container host: %v", err)
	}
	port, err := testContainer.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("Failed to get container port: %v", err)
	}

	return DatabaseConfig{
		Host:     host,
		DbName:   "testdb",
		Port:     port.Port(),
		User:     "testuser",
		Password: "testpass",
		Scheme:   scheme,
	}
}

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
			MatchesHost:     os.Getenv("MICROSERVICE_MATCHES_HOST"),
			MatchesPort:     os.Getenv("MICROSERVICE_MATCHES_PORT"),
		},
	}

	logger.Debug("Config loaded successfully")
	return config, nil
}

func LoadV2() (ConfigV2, error) {
	logger.Debug("Loading config v2...")

	env := os.Getenv("ENV")
	if env == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error al cargar el archivo .env")
		}
	}

	config := ConfigV2{
		ThirdPartyApiTokens: ThirdPartyApiTokens{
			FaceitApiToken: os.Getenv("FACEIT_API_TOKEN"),
		},
		MicroservicePlayers: MicroserviceConfig{
			Database: DatabaseConfig{
				Host:     os.Getenv("DB_HOST"),
				DbName:   os.Getenv("DB_NAME"),
				Port:     os.Getenv("DB_PORT"),
				User:     os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASSWORD"),
			},
			Host_gRPC: os.Getenv("MICROSERVICE_PLAYERS_HOST"),
			Port_gRPC: os.Getenv("MICROSERVICE_PLAYERS_PORT"),
		},
		MicroserviceTeams: MicroserviceConfig{
			Database: DatabaseConfig{
				Host:     os.Getenv("DB_HOST"),
				DbName:   os.Getenv("DB_NAME"),
				Port:     os.Getenv("DB_PORT"),
				User:     os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASSWORD"),
			},
			Host_gRPC: os.Getenv("MICROSERVICE_TEAMS_HOST"),
			Port_gRPC: os.Getenv("MICROSERVICE_TEAMS_PORT"),
		},
		MicroserviceUsers: MicroserviceConfig{
			Database: DatabaseConfig{
				Host:     os.Getenv("DB_HOST"),
				DbName:   os.Getenv("DB_NAME"),
				Port:     os.Getenv("DB_PORT"),
				User:     os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASSWORD"),
			},
			Host_gRPC: os.Getenv("MICROSERVICE_USERS_HOST"),
			Port_gRPC: os.Getenv("MICROSERVICE_USERS_PORT"),
		},
		MicroserviceTournaments: MicroserviceConfig{
			Database: DatabaseConfig{
				Host:     os.Getenv("DB_HOST"),
				DbName:   os.Getenv("DB_NAME"),
				Port:     os.Getenv("DB_PORT"),
				User:     os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASSWORD"),
			},
			Host_gRPC: os.Getenv("MICROSERVICE_TOURNAMENTS_HOST"),
			Port_gRPC: os.Getenv("MICROSERVICE_TOURNAMENTS_PORT"),
		},
		MicroserviceMatches: MicroserviceConfig{
			Database: DatabaseConfig{
				Host:     os.Getenv("DB_HOST"),
				DbName:   os.Getenv("DB_NAME"),
				Port:     os.Getenv("DB_PORT"),
				User:     os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASSWORD"),
			},
			Host_gRPC: os.Getenv("MICROSERVICE_MATCHES_HOST"),
			Port_gRPC: os.Getenv("MICROSERVICE_MATCHES_PORT"),
		},
	}

	logger.Debug("Config loaded successfully")
	return config, nil
}
