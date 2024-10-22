package microservices

import (
	"ibercs/pkg/logger"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func New[T any](cfgHost string, newClientFunc func(conn grpc.ClientConnInterface) T) *T {
	var creds credentials.TransportCredentials

	if env := os.Getenv("ENV"); env == "" {
		creds = insecure.NewCredentials()
	} else {
		creds = credentials.NewTLS(nil)
	}

	conn, err := grpc.NewClient(cfgHost, grpc.WithTransportCredentials(creds))
	if err != nil {
		logger.Error("Cannot connect to gRPC server at %s: %s", cfgHost, err.Error())
		return nil
	}

	grpcClient := newClientFunc(conn)
	logger.Trace("gRPC server at %s connected successfully", cfgHost)
	return &grpcClient
}
