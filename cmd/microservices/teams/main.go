package main

import (
	microservice_teams "ibercs/cmd/microservices/teams/server"
	"ibercs/pkg/config"
	"ibercs/pkg/logger"
	"net"

	pb "ibercs/proto/teams"

	"google.golang.org/grpc"
)

func main() {
	logger.Initialize()
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	port := cfg.MicroserviceTeams.Port_gRPC
	listener, err := net.Listen("tcp", port)
	if err != nil {
		logger.Error("Cannot create tcp connection: %s", err.Error())
		return
	} else {
		logger.Info("gRPC server started on port %s", port)
	}

	grpcServer := grpc.NewServer()
	microserviceTeam := microservice_teams.New(cfg.MicroserviceTeams, cfg.ThirdPartyApiTokens)

	pb.RegisterTeamServiceServer(grpcServer, microserviceTeam)

	if err := grpcServer.Serve(listener); err != nil {
		logger.Error("Cannot initialize the server: %s", err.Error())
		return
	}
}
