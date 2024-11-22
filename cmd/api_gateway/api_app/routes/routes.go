package routes

import (
	"ibercs/cmd/api_gateway/api_app/handlers"
	"ibercs/internal/api/middlewares"
	"ibercs/internal/router"
	"ibercs/pkg/cache"
	"ibercs/pkg/consts"

	"github.com/gin-gonic/gin"
)

type routes struct {
	router          *router.Router
	cacheMiddleware gin.HandlerFunc
	authMiddleware  gin.HandlerFunc
}

func New(router *router.Router) *routes {
	return &routes{
		router:          router,
		cacheMiddleware: middlewares.Cache(cache.NewCache(), consts.CACHE_DURATION),
		authMiddleware:  middlewares.Auth(*router.UsersServer),
	}
}

func (r routes) RegisterUserRoutes() {
	userHandler := handlers.NewUsersHandlers(*r.router.UsersServer)

	r.router.GET("/api/v2/user", userHandler.Get, r.cacheMiddleware) // query param: id or faceit_id
	r.router.PUT("/api/v2/user", userHandler.Update, r.authMiddleware)
	r.router.GET("/api/v2/users/streams", userHandler.GetStreams, r.cacheMiddleware)

	//Auth
	r.router.GET("/api/v2/auth/callback/faceit", userHandler.AuthCallback_Faceit)
	r.router.POST("/api/v2/auth", userHandler.Login, r.authMiddleware)
	r.router.DELETE("/api/v2/auth", userHandler.Logout, r.authMiddleware)
}

func (r routes) RegisterPlayerRoutes() {
	playerHandler := handlers.NewPlayersHandlers(*r.router.PlayersServer)

	r.router.GET("/api/v2/player", playerHandler.Get, r.cacheMiddleware) // query param: ids or nickname
	r.router.GET("/api/v2/players", playerHandler.GetAll, r.cacheMiddleware)
	r.router.GET("/api/v2/players/looking-for-team", playerHandler.GetLookingForTeamPlayers)
	r.router.POST("/api/v2/players/looking-for-team", playerHandler.CreateLookingForTeam, r.authMiddleware)
	r.router.PUT("/api/v2/players/looking-for-team", playerHandler.UpdateLookingForTeam, r.authMiddleware)
	r.router.DELETE("/api/v2/players/looking-for-team", playerHandler.DeleteLookingForTeam, r.authMiddleware)
	r.router.GET("/api/v2/players/prominent", playerHandler.GetProminentPlayers, r.cacheMiddleware)
}

func (r routes) RegisterTeamRoutes() {
	teamHandler := handlers.NewTeamsHandlers(*r.router.TeamsServer)

	r.router.POST("/api/v2/team/faceit", teamHandler.CreateFromFaceit, r.authMiddleware)
	r.router.GET("/api/v2/team", teamHandler.Get, r.cacheMiddleware) // query param: id or nickname
	r.router.GET("/api/v2/teams", teamHandler.GetAll, r.cacheMiddleware)
	r.router.GET("/api/v2/teams/active", teamHandler.GetActiveTeams, r.cacheMiddleware)
	r.router.GET("/api/v2/team/player", teamHandler.FindTeamByPlayerId)
}

func (r routes) RegisterTournamentRoutes() {
	tournamentHandler := handlers.NewTournamentsHandlers(*r.router.TournamentsServer)

	r.router.GET("/api/v2/tournaments", tournamentHandler.GetAll, r.cacheMiddleware)
	r.router.POST("/api/v2/organizer", tournamentHandler.CreateOrganizer, r.authMiddleware)
	r.router.GET("/api/v2/esea", tournamentHandler.GetEseaLeagues, r.cacheMiddleware)
}

func (r routes) RegisterMatchRoutes() {
	matchHandler := handlers.NewMatchesHandlers(*r.router.MatchesServer)

	r.router.GET("/api/v2/match", matchHandler.Get, r.cacheMiddleware)      // query param: id
	r.router.GET("/api/v2/matches", matchHandler.GetAll, r.cacheMiddleware) // query param: team_id
	r.router.GET("/api/v2/matches/range", matchHandler.GetRange, r.cacheMiddleware)
	r.router.POST("/api/v2/match/stream", matchHandler.SetStreamMatch)
	r.router.GET("/api/v2/matches/team", matchHandler.GetMatchesByTeamId, r.cacheMiddleware)
}

func (r routes) RegisterStateRoutes() {
	stateHandler := handlers.NewStateHandlers()

	r.router.GET("/api/v2/state", stateHandler.GetState)
	r.router.GET("/api/v2/state/players-update", stateHandler.GetState)
}
