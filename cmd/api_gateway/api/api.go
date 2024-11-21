package api

import (
	"ibercs/cmd/api_gateway/api/handlers"
	"ibercs/cmd/api_gateway/api/middlewares"
	"ibercs/internal/router"
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
	cfg    config.ConfigV2
	router *router.Router
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
	api.router = router.NewRouter(api.cfg)
	api.db = database.NewDatabase(api.cfg.MicroserviceUsers.Database).GetDB()

	cache := cache.NewCache()

	matchHandler := handlers.NewMatchesHandlers(*api.router.MatchesServer)
	playerHandler := handlers.NewPlayersHandlers(*api.router.PlayersServer)
	userHandler := handlers.NewUsersHandlers(*api.router.UsersServer)
	teamHandler := handlers.NewTeamsHandlers(*api.router.TeamsServer)
	tournamentHandler := handlers.NewTournamentsHandlers(*api.router.TournamentsServer)

	api.router.Use(middlewares.CORSMiddleware())
	cacheMiddleware := middlewares.Cache(cache, consts.CACHE_DURATION)
	authMiddleware := middlewares.Auth(api.db)

	api.router.GET("/api/v2/user", userHandler.Get, cacheMiddleware) // query param: id or faceit_id
	api.router.PUT("/api/v2/user", userHandler.Update, authMiddleware)
	api.router.GET("/api/v2/user/streams", userHandler.GetStreams, cacheMiddleware)
	api.router.GET("/api/v2/auth/callback/faceit", userHandler.AuthCallback_Faceit)
	api.router.POST("/api/v2/auth", userHandler.Login, authMiddleware)
	api.router.DELETE("/api/v2/auth", userHandler.Logout, authMiddleware)

	api.router.GET("/api/v2/player", playerHandler.Get, cacheMiddleware) // query param: ids or nickname
	api.router.GET("/api/v2/players", playerHandler.GetAll, cacheMiddleware)
	api.router.GET("/api/v2/players/looking-for-team", playerHandler.GetLookingForTeamPlayers)
	api.router.POST("/api/v2/players/looking-for-team", playerHandler.CreateLookingForTeam, authMiddleware)
	api.router.PUT("/api/v2/players/looking-for-team", playerHandler.UpdateLookingForTeam, authMiddleware)
	api.router.DELETE("/api/v2/players/looking-for-team", playerHandler.DeleteLookingForTeam, authMiddleware)
	api.router.GET("/api/v2/players/prominent", playerHandler.GetProminentPlayers, cacheMiddleware)

	api.router.POST("/api/v2/team/faceit", teamHandler.CreateFromFaceit, authMiddleware)
	api.router.GET("/api/v2/team", teamHandler.Get, cacheMiddleware) // query param: id or nickname
	api.router.GET("/api/v2/teams", teamHandler.GetAll, cacheMiddleware)
	api.router.GET("/api/v2/teams/active", teamHandler.GetActiveTeams, cacheMiddleware)

	api.router.GET("/api/v2/tournaments", tournamentHandler.GetAll, cacheMiddleware)
	api.router.POST("/api/v2/organizer", tournamentHandler.CreateOrganizer, authMiddleware)
	api.router.GET("/api/v2/esea", tournamentHandler.GetEseaLeagues, cacheMiddleware)

	api.router.GET("/api/v2/match", matchHandler.Get, cacheMiddleware)      // query param: id
	api.router.GET("/api/v2/matches", matchHandler.GetAll, cacheMiddleware) // query param: team_id
	api.router.GET("/api/v2/matches/range", matchHandler.GetRange, cacheMiddleware)

	logger.Debug("API initialized")
	api.router.Listen()
}
