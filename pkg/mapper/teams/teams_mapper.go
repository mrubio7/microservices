package teams_mapper

import (
	"ibercs/internal/model"
	pb "ibercs/proto/teams"
)

type TeamMapper struct{}

func (TeamMapper) Proto(entity model.TeamModel, _ ...interface{}) *pb.Team {
	mapStats := make(map[string]*pb.TeamMapStats, len(entity.Stats.MapStats))

	for _, m := range entity.Stats.MapStats {
		mapStats[m.MapName] = &pb.TeamMapStats{
			MapName: m.MapName,
			Winrate: m.WinRate,
			Matches: m.Matches,
		}
	}

	return &pb.Team{
		Id:          int32(entity.Id),
		Nickname:    entity.Nickname,
		FaceitId:    entity.FaceitId,
		Name:        entity.Name,
		Avatar:      entity.Avatar,
		Active:      entity.Active,
		PlayersId:   entity.PlayersId,
		Twitter:     entity.Twitter,
		Instagram:   entity.Instagram,
		Web:         entity.Web,
		Tournaments: entity.Tournaments,
		Stats: &pb.TeamStats{
			TotalMatches:  entity.Stats.TotalMatches,
			Wins:          entity.Stats.Wins,
			Winrate:       entity.Stats.Winrate,
			RecentResults: entity.Stats.RecentResults,
			MapStats:      mapStats,
		},
	}
}

func (TeamMapper) Model(proto *pb.Team, _ ...interface{}) model.TeamModel {
	jsonMapStats := model.JSONMapStats{}
	for _, m := range proto.Stats.MapStats {
		jsonMapStats[m.MapName] = model.TeamMapStats{
			MapName: m.MapName,
			WinRate: m.Winrate,
			Matches: m.Matches,
		}
	}

	return model.TeamModel{
		Id:          int32(proto.Id),
		Nickname:    proto.Nickname,
		FaceitId:    proto.FaceitId,
		Name:        proto.Name,
		Avatar:      proto.Avatar,
		Active:      proto.Active,
		PlayersId:   proto.PlayersId,
		Twitter:     proto.Twitter,
		Instagram:   proto.Instagram,
		Web:         proto.Web,
		Tournaments: proto.Tournaments,
		Stats: model.TeamStatsModel{
			TotalMatches:  proto.Stats.TotalMatches,
			Wins:          proto.Stats.Wins,
			Winrate:       proto.Stats.Winrate,
			RecentResults: proto.Stats.RecentResults,
			MapStats:      jsonMapStats,
		},
	}
}

type TeamMapMapper struct{}

func (TeamMapMapper) Proto(model model.TeamMapStats) *pb.TeamMapStats {
	return &pb.TeamMapStats{}
}

func (TeamMapMapper) Model(proto *pb.TeamMapStats) model.TeamMapStats {
	return model.TeamMapStats{}
}
