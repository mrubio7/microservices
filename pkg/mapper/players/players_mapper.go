package players_mapper

import (
	"ibercs/internal/model"
	pb "ibercs/proto/players"
	"time"
)

type PlayerMapper struct{}

func (PlayerMapper) Proto(model model.PlayerModel) *pb.Player {
	return &pb.Player{
		Id:       int32(model.ID),
		Nickname: model.Nickname,
		FaceitId: model.FaceitId,
		SteamId:  model.SteamId,
		Avatar:   model.Avatar,
		Stats: &pb.PlayerStats{
			KdRatio:                model.Stats.KdRatio,
			KrRatio:                model.Stats.KrRatio,
			KillsAverage:           model.Stats.KillsAverage,
			DeathsAverage:          model.Stats.DeathsAverage,
			AssistAverage:          model.Stats.AssistAverage,
			HeadshotPercentAverage: model.Stats.HeadshotPercentAverage,
			MVPAverage:             model.Stats.MVPAverage,
			Elo:                    model.Stats.Elo,
		},
	}
}

func (PlayerMapper) Model(proto *pb.Player) model.PlayerModel {
	return model.PlayerModel{
		ID:       int32(proto.Id),
		Nickname: proto.Nickname,
		FaceitId: proto.FaceitId,
		SteamId:  proto.SteamId,
		Avatar:   proto.Avatar,
		Stats: model.PlayerStatsModel{
			KdRatio:                proto.Stats.KdRatio,
			KrRatio:                proto.Stats.KrRatio,
			KillsAverage:           proto.Stats.KillsAverage,
			DeathsAverage:          proto.Stats.DeathsAverage,
			AssistAverage:          proto.Stats.AssistAverage,
			HeadshotPercentAverage: proto.Stats.HeadshotPercentAverage,
			MVPAverage:             proto.Stats.MVPAverage,
			Elo:                    proto.Stats.Elo,
		},
	}
}

type PlayerProminentMapper struct{}

func (PlayerProminentMapper) Proto(model model.PlayerProminentModel) *pb.ProminentPlayer {
	return &pb.ProminentPlayer{
		Id:       int32(model.ID),
		Score:    model.Score,
		Avatar:   model.Avatar,
		Nickname: model.Nickname,
		FaceitId: model.FaceitId,
		SteamId:  model.SteamId,
	}
}

func (PlayerProminentMapper) Model(proto *pb.ProminentPlayer) model.PlayerProminentModel {
	return model.PlayerProminentModel{
		ID:       int32(proto.Id),
		Avatar:   proto.Avatar,
		Nickname: proto.Nickname,
		FaceitId: proto.FaceitId,
		SteamId:  proto.SteamId,
		Score:    proto.Score,
	}
}

type PlayerLookingForTeamMapper struct{}

func (PlayerLookingForTeamMapper) Proto(model model.LookingForTeamModel) *pb.PlayerLookingForTeam {
	return &pb.PlayerLookingForTeam{
		Id:           int32(model.Id),
		InGameRole:   model.InGameRole,
		TimeTable:    model.TimeTable,
		OldTeams:     model.OldTeams,
		PlayingYears: model.PlayingYears,
		BornDate:     model.BornDate.Unix(),
		Description:  model.Description,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}
}

func (PlayerLookingForTeamMapper) Model(proto *pb.PlayerLookingForTeam) model.LookingForTeamModel {
	return model.LookingForTeamModel{
		Id:           int32(proto.Id),
		InGameRole:   proto.InGameRole,
		TimeTable:    proto.TimeTable,
		OldTeams:     proto.OldTeams,
		PlayingYears: proto.PlayingYears,
		BornDate:     time.Unix(proto.BornDate, 0),
		Description:  proto.Description,
		CreatedAt:    proto.CreatedAt,
		UpdatedAt:    proto.UpdatedAt,
	}
}

type CreatePlayerLookingForTeamMapper struct{}

func (CreatePlayerLookingForTeamMapper) Proto(model model.LookingForTeamModel, userId int32) *pb.CreatePlayerLookingForTeamRequest {
	return &pb.CreatePlayerLookingForTeamRequest{
		InGameRole:   model.InGameRole,
		TimeTable:    model.TimeTable,
		OldTeams:     model.OldTeams,
		PlayingYears: model.PlayingYears,
		Description:  model.Description,
		FaceitId:     model.FaceitId,
		UserId:       userId,
	}
}

func (CreatePlayerLookingForTeamMapper) Model(proto *pb.CreatePlayerLookingForTeamRequest) model.LookingForTeamModel {
	return model.LookingForTeamModel{
		InGameRole:   proto.InGameRole,
		FaceitId:     proto.FaceitId,
		TimeTable:    proto.TimeTable,
		OldTeams:     proto.OldTeams,
		PlayingYears: proto.PlayingYears,
		Description:  proto.Description,
	}
}
