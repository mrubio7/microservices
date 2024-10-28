package api

import (
	"ibercs/pkg/config"
	"ibercs/pkg/microservices"
	pb_players "ibercs/proto/players"
	pb_teams "ibercs/proto/teams"
	pb_users "ibercs/proto/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	gin           *gin.Engine
	PlayersServer *pb_players.PlayerServiceClient
	TeamsServer   *pb_teams.TeamServiceClient
	UsersServer   *pb_users.UserServiceClient
}

func NewRouter(cfg config.Config) *Router {
	r := &Router{
		gin: gin.Default(),
	}

	r.registerPlayersServer(cfg.Microservices)
	r.registerTeamsServer(cfg.Microservices)
	r.registerUsersServer(cfg.Microservices)

	return r
}

func (r *Router) Listen() {
	if err := http.ListenAndServe(":8080", r.gin); err != nil {
		panic(err)
	}
}

func (r *Router) registerPlayersServer(cfg config.MicroservicesConfig) {
	r.PlayersServer = microservices.New(cfg.PlayersHost, cfg.PlayersPort, pb_players.NewPlayerServiceClient)
}

func (r *Router) registerTeamsServer(cfg config.MicroservicesConfig) {
	r.TeamsServer = microservices.New(cfg.TeamsHost, cfg.TeamsPort, pb_teams.NewTeamServiceClient)
}

func (r *Router) registerUsersServer(cfg config.MicroservicesConfig) {
	r.UsersServer = microservices.New(cfg.UserHost, cfg.UserPort, pb_users.NewUserServiceClient)
}
