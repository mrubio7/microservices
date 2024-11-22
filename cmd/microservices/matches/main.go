package main

import (
	microservice_matches "ibercs/cmd/microservices/matches/server"
	"ibercs/pkg/config"
	"ibercs/pkg/logger"
	"net"

	pb "ibercs/proto/matches"

	"google.golang.org/grpc"
)

func main() {
	logger.Initialize()
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	port := cfg.MicroserviceMatches.Port_gRPC
	listener, err := net.Listen("tcp", port)
	if err != nil {
		logger.Error("Cannot create tcp connection: %s", err.Error())
		return
	} else {
		logger.Info("gRPC server started on port %s", port)
	}

	grpcServer := grpc.NewServer()
	microserviceMatches := microservice_matches.New(cfg.MicroserviceMatches, cfg.ThirdPartyApiTokens)

	pb.RegisterMatchesServiceServer(grpcServer, microserviceMatches)

	if err := grpcServer.Serve(listener); err != nil {
		logger.Error("Cannot initialize the server: %s", err.Error())
		return
	}
}
