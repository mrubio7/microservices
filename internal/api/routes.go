package api

import (
	"ibercs/pkg/logger"
	pb_players "ibercs/proto/players"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Router struct {
	gin           *gin.Engine
	PlayersServer pb_players.PlayerServiceClient
}

func NewRouter() *Router {
	r := &Router{
		gin: gin.Default(),
	}

	r.registerPlayersServer()

	return r
}

func (r *Router) Listen() {
	if err := http.ListenAndServe(":8080", r.gin); err != nil {
		panic(err)
	}
}

func (r *Router) registerPlayersServer() {
	creds := credentials.NewTLS(nil)
	conn, err := grpc.NewClient(os.Getenv("MICROSERVICE_PLAYERS_HOST"), grpc.WithTransportCredentials(creds))
	if err != nil {
		logger.Error("Cannot connect to players grpc server: %s", err.Error())
		return
	}

	playersGrpc := pb_players.NewPlayerServiceClient(conn)

	r.PlayersServer = playersGrpc
	logger.Trace("Players grpc server connected succesfully")
}
