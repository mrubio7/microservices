package main

import (
	microservice_teams "ibercs/cmd/microservices/teams/server"
	"ibercs/pkg/logger"
	"net"

	pb "ibercs/proto/teams"

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
	microserviceTeams := microservice_teams.New()

	pb.RegisterTeamServiceServer(grpcServer, microserviceTeams)

	if err := grpcServer.Serve(listener); err != nil {
		logger.Error("Cannot initialize the server: %s", err.Error())
		return
	}
}
