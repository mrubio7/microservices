package config

import (
	"fmt"
	"ibercs/pkg/logger"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Database       DatabaseConfig
		FaceitApiToken string
	}

	DatabaseConfig struct {
		Host     string
		DbName   string
		Port     string
		User     string
		Password string
	}
)

func Load() (Config, error) {
	logger.Debug("Loading config...")

	// Obtener el directorio actual del ejecutable
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error getting executable path: %v\n", err)
		return Config{}, err
	}

	configPath := filepath.Join(filepath.Dir(exePath), "config.json")
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshaling config: %v\n", err)
		return Config{}, err
	}

	logger.Debug("Config loaded successfully")
	return config, nil
}
