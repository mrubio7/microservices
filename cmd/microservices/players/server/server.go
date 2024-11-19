package microservice_players

import (
	"context"
	"ibercs/pkg/config"
	"ibercs/pkg/consts"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
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

		player := &pb.Player{
			Id:       p.ID,
			Nickname: p.Nickname,
			FaceitId: p.FaceitId,
			SteamId:  p.SteamId,
			Avatar:   p.Avatar,
			Stats: &pb.PlayerStats{
				PlayerId:               p.Stats.ID,
				KdRatio:                p.Stats.KdRatio,
				KrRatio:                p.Stats.KrRatio,
				KillsAverage:           p.Stats.KillsAverage,
				DeathsAverage:          p.Stats.DeathsAverage,
				HeadshotPercentAverage: p.Stats.HeadshotPercentAverage,
				MVPAverage:             p.Stats.MVPAverage,
				AssistAverage:          p.Stats.AssistAverage,
				Elo:                    p.Stats.Elo,
			},
		}

		playersRes[i] = player
	}

	return &pb.PlayerList{Players: playersRes}, nil
}
