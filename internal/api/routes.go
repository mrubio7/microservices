package api

import (
	"ibercs/pkg/logger"
	pb_players "ibercs/proto/players"
	pb_teams "ibercs/proto/teams"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Router struct {
	gin           *gin.Engine
	PlayersServer pb_players.PlayerServiceClient
	TeamsServer   pb_teams.TeamServiceClient
}

func NewRouter() *Router {
	r := &Router{
		gin: gin.Default(),
	}

	r.registerPlayersServer()
	r.registerTeamsServer()

	return r
}

func (r *Router) Listen() {
	if err := http.ListenAndServe(":8080", r.gin); err != nil {
		panic(err)
	}
}

func (r *Router) registerPlayersServer() {
	var creds credentials.TransportCredentials

	if env := os.Getenv("ENV"); env == "" {
		creds = insecure.NewCredentials()
	} else {
		creds = credentials.NewTLS(nil)
	}

	conn, err := grpc.NewClient(os.Getenv("MICROSERVICE_PLAYERS_HOST"), grpc.WithTransportCredentials(creds))
	if err != nil {
		logger.Error("Cannot connect to players grpc server: %s", err.Error())
		return
	}

	playersGrpc := pb_players.NewPlayerServiceClient(conn)

	r.PlayersServer = playersGrpc
	logger.Trace("Players grpc server connected succesfully")
}

func (r *Router) registerTeamsServer() {
	var creds credentials.TransportCredentials

	if env := os.Getenv("ENV"); env == "" {
		creds = insecure.NewCredentials()
	} else {
		creds = credentials.NewTLS(nil)
	}

	var host string
	if hostEnv := os.Getenv("MICROSERVICE_TEAMS_HOST"); hostEnv == "" {
		host = "localhost:50051"
	} else {
		host = hostEnv
	}

	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(creds))
	if err != nil {
		logger.Error("Cannot connect to teams grpc server: %s", err.Error())
		return
	}

	teamsGrpc := pb_teams.NewTeamServiceClient(conn)

	r.TeamsServer = teamsGrpc
	logger.Trace("Teams grpc server connected succesfully")
}
