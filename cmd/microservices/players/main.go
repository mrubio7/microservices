package main

import (
	microservice_players "ibercs/cmd/microservices/players/server"
	"ibercs/pkg/config"
	"ibercs/pkg/logger"
	"net"

	pb "ibercs/proto/players"

	"google.golang.org/grpc"
)

func main() {
	logger.Initialize()
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	port := cfg.Microservices.PlayersPort
	listener, err := net.Listen("tcp", port)
	if err != nil {
		logger.Error("Cannot create tcp connection: %s", err.Error())
		return
	} else {
		logger.Info("gRPC server started on port %s", port)
	}

	grpcServer := grpc.NewServer()
	microservicePlayers := microservice_players.New()

	pb.RegisterPlayerServiceServer(grpcServer, microservicePlayers)

	if err := grpcServer.Serve(listener); err != nil {
		logger.Error("Cannot initialize the server: %s", err.Error())
		return
	}
}
