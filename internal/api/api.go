package api

import (
	"ibercs/internal/api/handlers"
	"ibercs/internal/api/middlewares"
	"ibercs/pkg/cache"
	"ibercs/pkg/config"
	"ibercs/pkg/consts"
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

	api.router = NewRouter(api.cfg)
	cache := cache.NewCache()

	players_handlers := handlers.NewPlayersHandlers(*api.router.PlayersServer)
	teams_handlers := handlers.NewTeamsHandlers(*api.router.TeamsServer)
	users_handlers := handlers.NewUsersHandlers(*api.router.UsersServer)
	tournaments_handlers := handlers.NewTournamentsHandlers(*api.router.TournamentsServer)
	workers_handlers := handlers.NewWorkersHandlers(api.cfg.Workers)
	matches_handlers := handlers.NewMatchesHandlers(*api.router.MatchesServer)
	state_handlers := handlers.NewStateHandlers(api.db)

	api.router.gin.Use(middlewares.CORSMiddleware())
	api.router.gin.Use(middlewares.CacheMiddleware(cache, consts.CACHE_DURATION))
	api.router.gin.Use(middlewares.CacheInvalidationMiddleware(cache))

	api.router.gin.GET("/api/v1/tournaments/get-all", tournaments_handlers.GetAllTournaments)

	api.router.gin.GET("/api/v1/state/last-players-update", state_handlers.GetLastPlayerUpdate)

	api.router.gin.GET("/api/v1/users/get", users_handlers.GetUser)
	api.router.gin.POST("/api/v1/auth/callback", users_handlers.FaceitAuthCallback)
	api.router.gin.GET("/api/v1/users/get-streams", users_handlers.GetStreams)

	api.router.gin.GET("/api/v1/players/get", players_handlers.GetPlayers)
	api.router.gin.GET("/api/v1/players/get-all", players_handlers.GetAllPlayers)
	api.router.gin.GET("/api/v1/players/get-prominent-players", players_handlers.GetProminentPlayers)
	api.router.gin.GET("/api/v1/players/get-looking-for-team", players_handlers.GetAllLookingForTeam)

	api.router.gin.POST("/api/v1/teams/new", teams_handlers.New)
	api.router.gin.GET("/api/v1/teams/get", teams_handlers.Get)
	api.router.gin.GET("/api/v1/teams/get-all", teams_handlers.GetAll)
	api.router.gin.GET("/api/v1/teams/ranks", teams_handlers.GetRanks)
	api.router.gin.GET("/api/v1/teams/find-player", teams_handlers.FindTeamByPlayerId)

	api.router.gin.GET("/api/v1/matches/get-all", matches_handlers.GetAll)
	api.router.gin.GET("/api/v1/matches/get", matches_handlers.GetById)
	api.router.gin.GET("/api/v1/matches/get-range", matches_handlers.GetRange)
	api.router.gin.GET("/api/v1/matches/get-team", matches_handlers.GetTeamMatches)

	api.router.gin.GET("/api/v1/esea/details", tournaments_handlers.GetEseaDetails)

	api.router.gin.Use(middlewares.Auth(api.db))
	api.router.gin.POST("/api/v1/players/looking-for-team", players_handlers.NewLookingForTeam)
	api.router.gin.DELETE("/api/v1/players/looking-for-team", players_handlers.DeleteLookingForTeam)
	api.router.gin.POST("/api/v1/matches/set-stream", matches_handlers.SetStreamMatch)
	api.router.gin.POST("/api/v1/organizers/new", tournaments_handlers.NewOrganizer)
	api.router.gin.POST("/api/v1/tournaments/new", tournaments_handlers.NewTournament)
	api.router.gin.GET("/api/v1/workers/players/update", workers_handlers.Update)
	api.router.gin.GET("/api/v1/workers/tournaments/find", workers_handlers.TournamentsFind)
	api.router.gin.POST("/api/v1/user", users_handlers.Auth)
	api.router.gin.POST("/api/v1/users/update", users_handlers.UpdateProfile)
	api.router.gin.POST("/api/v1/auth/logout", users_handlers.Logout)
	api.router.gin.GET("/api/v1/state", state_handlers.GetState)

	api.router.Listen()
}
