package main

import (
	microservice_users "ibercs/cmd/microservices/users/server"
	"ibercs/pkg/config"
	"ibercs/pkg/logger"
	"net"

	pb "ibercs/proto/users"

	"google.golang.org/grpc"
)

func main() {
	logger.Initialize()
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	port := cfg.Microservices.UserPort
	listener, err := net.Listen("tcp", port)
	if err != nil {
		logger.Error("Cannot create tcp connection: %s", err.Error())
		return
	} else {
		logger.Info("gRPC server started on port %s", port)
	}

	grpcServer := grpc.NewServer()
	microserviceUsers := microservice_users.New()

	pb.RegisterUserServiceServer(grpcServer, microserviceUsers)

	if err := grpcServer.Serve(listener); err != nil {
		logger.Error("Cannot initialize the server: %s", err.Error())
		return
	}
}
