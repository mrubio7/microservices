package api

import (
	"ibercs/internal/api/handlers"
	"ibercs/internal/api/middlewares"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/logger"
	"os"

	"gorm.io/gorm"
)

type Api struct {
	db     *gorm.DB
	cfg    config.Config
	router *Router
}

func New() *Api {
	config, err := config.Load()
	if err != nil {
		logger.Error("config can't be loaded")
		os.Exit(0)
	}

	return &Api{
		db:  database.New(config.Database),
		cfg: config,
	}
}

func (api *Api) Start() {
	logger.Debug("Initializing API...")

	api.router = NewRouter()

	players_handlers := handlers.NewPlayersHandlers(api.router.PlayersServer)

	api.router.gin.Use(middlewares.CORSMiddleware())
	api.router.gin.GET("/api/v1/players/get-all", players_handlers.GetAllPlayers)

	api.router.Listen()
}
