package server

import (
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/microservices"
	"ibercs/pkg/service"
	pb_players "ibercs/proto/players"
	pb "ibercs/proto/users"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	FaceitService *faceit.FaceitClient
	UsersService  *service.Users
	PlayersClient *pb_players.PlayerServiceClient
}

func registerPlayersClient(cfg config.MicroservicesConfig) *pb_players.PlayerServiceClient {
	return microservices.New(cfg.PlayersHost, pb_players.NewPlayerServiceClient)
}

func New() *Server {
	cfg, err := config.Load()
	if err != nil {
		logger.Error("Unable to create grpc player server")
		return nil
	}
	db := database.New(cfg.Database)

	return &Server{
		FaceitService: faceit.New(cfg.FaceitApiToken),
		UsersService:  service.NewUsersService(db),
		PlayersClient: registerPlayersClient(cfg.Microservices),
	}
}
