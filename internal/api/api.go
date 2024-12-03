package api

import (
	"ibercs/internal/router"
	"ibercs/pkg/config"
	"ibercs/pkg/logger"
	"os"
)

type RouteRegister func(api *Api)

type Api struct {
	cfg    config.ConfigV2
	Router *router.Router
}

func New(grpcConnection bool) *Api {
	config, err := config.Load()
	if err != nil {
		logger.Error("config can't be loaded")
		os.Exit(0)
	}

	return &Api{
		cfg:    config,
		Router: router.NewRouter(config, grpcConnection),
	}
}

func (api *Api) RegisterRoutes(registers ...RouteRegister) {
	for _, register := range registers {
		register(api)
	}
}

func (api *Api) Start() {
	logger.Info("Starting API...")

	api.Router.Listen()
}
