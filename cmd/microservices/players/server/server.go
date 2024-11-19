package microservice_players

import (
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/managers"
	pb "ibercs/proto/players"
)

type Server struct {
	pb.UnimplementedPlayerServiceServer
	PlayerManager *managers.PlayerManager
	FaceitService *faceit.FaceitClient
}

func New(cfg config.MicroserviceConfig, cfgThirdParty config.ThirdPartyApiTokens) *Server {
	db := database.NewDatabase(cfg.Database)
	playerManager := managers.NewPlayerManager(db.GetDB())
	faceit := faceit.New(cfgThirdParty.FaceitApiToken)

	return &Server{
		PlayerManager: playerManager,
		FaceitService: faceit,
	}
}
