package api

import (
	"ibercs/pkg/config"
	"ibercs/pkg/microservices"
	pb_players "ibercs/proto/players"
	pb_teams "ibercs/proto/teams"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	gin           *gin.Engine
	PlayersServer *pb_players.PlayerServiceClient
	TeamsServer   *pb_teams.TeamServiceClient
}

func NewRouter(cfg config.Config) *Router {
	r := &Router{
		gin: gin.Default(),
	}

	r.registerPlayersServer(cfg.Microservices)
	r.registerTeamsServer(cfg.Microservices)

	return r
}

func (r *Router) Listen() {
	if err := http.ListenAndServe(":8080", r.gin); err != nil {
		panic(err)
	}
}

func (r *Router) registerPlayersServer(cfg config.MicroservicesConfig) {
	r.PlayersServer = microservices.New(cfg.PlayersHost, pb_players.NewPlayerServiceClient)
}

func (r *Router) registerTeamsServer(cfg config.MicroservicesConfig) {
	r.TeamsServer = microservices.New(cfg.TeamsHost, pb_teams.NewTeamServiceClient)
}
