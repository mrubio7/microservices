package main

import (
	microservice_tournaments "ibercs/cmd/microservices/tournaments/server"
	"ibercs/pkg/config"
	"ibercs/pkg/logger"
	"net"

	pb "ibercs/proto/tournaments"

	"google.golang.org/grpc"
)

func main() {
	logger.Initialize()
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	port := cfg.MicroserviceTournaments.Port_gRPC
	listener, err := net.Listen("tcp", port)
	if err != nil {
		logger.Error("Cannot create tcp connection: %s", err.Error())
		return
	} else {
		logger.Info("gRPC server started on port %s", port)
	}

	grpcServer := grpc.NewServer()
	microserviceTournaments := microservice_tournaments.New(cfg.MicroserviceTournaments, cfg.MicroserviceTeams, cfg.ThirdPartyApiTokens)

	pb.RegisterTournamentServiceServer(grpcServer, microserviceTournaments)

	if err := grpcServer.Serve(listener); err != nil {
		logger.Error("Cannot initialize the server: %s", err.Error())
		return
	}
}
