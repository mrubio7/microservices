package microservice_players

import (
	"context"
	"ibercs/internal/model"
	"ibercs/pkg/config"
	"ibercs/pkg/consts"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/managers"
	"ibercs/pkg/mapper"
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

func (s *Server) GetPlayer(ctx context.Context, req *pb.GetPlayerRequest) (*pb.PlayerList, error) {
	playersRes := make([]*pb.Player, len(req.FaceitId))

	for i, p := range req.FaceitId {
		playerUpdated := s.FaceitService.GetPlayerAverageDetails(p, consts.LAST_MATCHES_NUMBER)
		err := s.PlayerManager.Update(playerUpdated)
		if err != nil {
			logger.Warning("Error updating player: %s", err.Error())
		}

		p, err := s.PlayerManager.GetByFaceitId(p)
		if err != nil {
			return nil, err
		}

		pbPlayer := mapper.Convert[*model.PlayerModel, pb.Player](p)
		playersRes[i] = &pbPlayer
	}

	return &pb.PlayerList{Players: playersRes}, nil
}
