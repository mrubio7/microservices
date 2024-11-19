package api

import (
	"ibercs/internal/api/handlers"
	"ibercs/internal/api/middlewares"
	"ibercs/pkg/cache"
	"ibercs/pkg/config"
	"ibercs/pkg/consts"
	"ibercs/pkg/logger"
	"os"

	"gorm.io/gorm"
)

type Api struct {
	db     *gorm.DB
	cfg    config.ConfigV2
	router *Router
}

func New() *Api {
	config, err := config.LoadV2()
	if err != nil {
		logger.Error("config can't be loaded")
		os.Exit(0)
	}

	return &Api{
		cfg: config,
	}
}

func (api *Api) Start() {
	logger.Debug("Initializing API...")
	api.router = NewRouter(api.cfg)

	cache := cache.NewCache()

	matchHandler := handlers.NewMatchesHandlers(*api.router.MatchesServer)

	api.router.gin.Use(middlewares.CORSMiddleware())
	api.router.gin.Use(middlewares.CacheMiddleware(cache, consts.CACHE_DURATION))
	api.router.gin.Use(middlewares.CacheInvalidationMiddleware(cache))

	api.router.gin.GET("/api/v2/match", matchHandler.Get)      // query param: id
	api.router.gin.GET("/api/v2/matches", matchHandler.GetAll) // query param: team_id
	api.router.gin.GET("/api/v2/matches/range", matchHandler.GetRange)

	api.router.Listen()
}
