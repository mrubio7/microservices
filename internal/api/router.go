package api

import (
	"ibercs/pkg/config"
	"ibercs/pkg/microservices"
	pb_matches "ibercs/proto/matches"
	pb_players "ibercs/proto/players"
	pb_teams "ibercs/proto/teams"
	pb_tournaments "ibercs/proto/tournaments"
	pb_users "ibercs/proto/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	gin               *gin.Engine
	PlayersServer     *pb_players.PlayerServiceClient
	TeamsServer       *pb_teams.TeamServiceClient
	UsersServer       *pb_users.UserServiceClient
	TournamentsServer *pb_tournaments.TournamentServiceClient
	MatchesServer     *pb_matches.MatchesServiceClient
}

func NewRouter(cfg config.ConfigV2) *Router {
	r := &Router{
		gin: gin.Default(),
	}

	r.registerMatchesServer(cfg.MicroserviceMatches)
	r.registerPlayersServer(cfg.MicroservicePlayers)
	r.registerTeamsServer(cfg.MicroserviceTeams)
	r.registerUsersServer(cfg.MicroserviceUsers)
	r.registerTournamentsServer(cfg.MicroserviceTournaments)

	return r
}

func (r *Router) GET(path string, handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) {
	r.gin.Handle("GET", path, func(c *gin.Context) {
		// Ejecuta todos los middlewares
		for _, middleware := range middlewares {
			middleware(c)
			if c.IsAborted() {
				return // Si alguno aborta, no se ejecuta el handler
			}
		}
		// Ejecuta el handler original
		handler(c)
	})
}

func (r *Router) POST(path string, handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) {
	r.gin.Handle("POST", path, func(c *gin.Context) {
		// Ejecuta todos los middlewares
		for _, middleware := range middlewares {
			middleware(c)
			if c.IsAborted() {
				return // Si alguno aborta, no se ejecuta el handler
			}
		}
		// Ejecuta el handler original
		handler(c)
	})
}

func (r *Router) PUT(path string, handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) {
	r.gin.Handle("PUT", path, func(c *gin.Context) {
		// Ejecuta todos los middlewares
		for _, middleware := range middlewares {
			middleware(c)
			if c.IsAborted() {
				return // Si alguno aborta, no se ejecuta el handler
			}
		}
		// Ejecuta el handler original
		handler(c)
	})
}

func (r *Router) DELETE(path string, handler gin.HandlerFunc, middlewares ...gin.HandlerFunc) {
	r.gin.Handle("DELETE", path, func(c *gin.Context) {
		// Ejecuta todos los middlewares
		for _, middleware := range middlewares {
			middleware(c)
			if c.IsAborted() {
				return // Si alguno aborta, no se ejecuta el handler
			}
		}
		// Ejecuta el handler original
		handler(c)
	})
}

func (r *Router) Listen() {
	if err := http.ListenAndServe(":8080", r.gin); err != nil {
		panic(err)
	}
}

func (r *Router) registerPlayersServer(cfg config.MicroserviceConfig) {
	r.PlayersServer = microservices.New(cfg.Host_gRPC, cfg.Port_gRPC, pb_players.NewPlayerServiceClient)
}

func (r *Router) registerTeamsServer(cfg config.MicroserviceConfig) {
	r.TeamsServer = microservices.New(cfg.Host_gRPC, cfg.Port_gRPC, pb_teams.NewTeamServiceClient)
}

func (r *Router) registerUsersServer(cfg config.MicroserviceConfig) {
	r.UsersServer = microservices.New(cfg.Host_gRPC, cfg.Port_gRPC, pb_users.NewUserServiceClient)
}

func (r *Router) registerTournamentsServer(cfg config.MicroserviceConfig) {
	r.TournamentsServer = microservices.New(cfg.Host_gRPC, cfg.Port_gRPC, pb_tournaments.NewTournamentServiceClient)
}

func (r *Router) registerMatchesServer(cfg config.MicroserviceConfig) {
	r.MatchesServer = microservices.New(cfg.Host_gRPC, cfg.Port_gRPC, pb_matches.NewMatchesServiceClient)
}
