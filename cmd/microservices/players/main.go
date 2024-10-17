package main

import (
	microservice_players "ibercs/cmd/microservices/players/server"
	"ibercs/pkg/logger"
	"net"

	pb "ibercs/proto/players"

	"google.golang.org/grpc"
)

func main() {
	logger.Initialize()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Error("Cannot create tcp connection: %s", err.Error())
		return
	} else {
		logger.Info("gRPC server started on port 50051")
	}

	grpcServer := grpc.NewServer()
	microservicePlayers := microservice_players.New()

	pb.RegisterPlayerServiceServer(grpcServer, microservicePlayers)

	if err := grpcServer.Serve(listener); err != nil {
		logger.Error("Cannot initialize the server: %s", err.Error())
		return
	}
}
