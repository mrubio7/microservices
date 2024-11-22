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
	ConfigV2 struct {
		ThirdPartyApiTokens     ThirdPartyApiTokens
		MicroservicePlayers     MicroserviceConfig
		MicroserviceTeams       MicroserviceConfig
		MicroserviceUsers       MicroserviceConfig
		MicroserviceTournaments MicroserviceConfig
		MicroserviceMatches     MicroserviceConfig
		StateDb                 DatabaseConfig
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

	MicroserviceConfig struct {
		Database  DatabaseConfig
		Host_gRPC string
		Port_gRPC string
	}

	WorkerConfig struct {
		ThirdPartyApiTokens ThirdPartyApiTokens
		TournamentsDb       DatabaseConfig
		MatchesDb           DatabaseConfig
		UsersDb             DatabaseConfig
		PlayersDb           DatabaseConfig
		TeamsDb             DatabaseConfig
		StateDb             DatabaseConfig
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

func LoadWorker() (WorkerConfig, error) {
	logger.Debug("Loading config v2...")

	env := os.Getenv("ENV")
	if env == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error al cargar el archivo .env")
		}
	}

	config := WorkerConfig{
		ThirdPartyApiTokens: ThirdPartyApiTokens{
			FaceitApiToken: os.Getenv("FACEIT_API_TOKEN"),
		},
		TournamentsDb: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			DbName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Scheme:   os.Getenv("MICROSERVICE_TOURNAMENTS_DB_SCHEME"),
		},
		MatchesDb: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			DbName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Scheme:   os.Getenv("MICROSERVICE_MATCHES_DB_SCHEME"),
		},
		UsersDb: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			DbName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Scheme:   os.Getenv("MICROSERVICE_USERS_DB_SCHEME"),
		},
		PlayersDb: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			DbName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Scheme:   os.Getenv("MICROSERVICE_PLAYERS_DB_SCHEME"),
		},
		TeamsDb: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			DbName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Scheme:   os.Getenv("MICROSERVICE_TEAMS_DB_SCHEME"),
		},
		StateDb: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			DbName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Scheme:   os.Getenv("STATE_DB_SCHEME"),
		},
	}

	logger.Debug("Config loaded successfully")
	return config, nil
}

func Load() (ConfigV2, error) {
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
		StateDb: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			DbName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Scheme:   os.Getenv("STATE_DB_SCHEME"),
		},
		MicroservicePlayers: MicroserviceConfig{
			Database: DatabaseConfig{
				Host:     os.Getenv("DB_HOST"),
				DbName:   os.Getenv("DB_NAME"),
				Port:     os.Getenv("DB_PORT"),
				User:     os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASSWORD"),
				Scheme:   os.Getenv("MICROSERVICE_PLAYERS_DB_SCHEME"),
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
				Scheme:   os.Getenv("MICROSERVICE_TEAMS_DB_SCHEME"),
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
				Scheme:   os.Getenv("MICROSERVICE_USERS_DB_SCHEME"),
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
				Scheme:   os.Getenv("MICROSERVICE_TOURNAMENTS_DB_SCHEME"),
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
				Scheme:   os.Getenv("MICROSERVICE_MATCHES_DB_SCHEME"),
			},
			Host_gRPC: os.Getenv("MICROSERVICE_MATCHES_HOST"),
			Port_gRPC: os.Getenv("MICROSERVICE_MATCHES_PORT"),
		},
	}

	logger.Debug("Config loaded successfully")
	return config, nil
}
