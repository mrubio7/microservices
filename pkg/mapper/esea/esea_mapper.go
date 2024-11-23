package esea_mapper

import (
	"encoding/json"
	"ibercs/internal/model"
	"ibercs/pkg/logger"
	teams_mapper "ibercs/pkg/mapper/teams"
	pb "ibercs/proto/tournaments"
)

type EseaMapper struct{}

func (EseaMapper) Proto(entity model.EseaLeagueModel, _ ...interface{}) *pb.Esea {
	divisions := make([]*pb.EseaDivision, 0)

	for _, division := range entity.Divisions {
		divisions = append(divisions, EseaDivisionMapper{}.Proto(division))
	}

	return &pb.Esea{
		FaceitId:  entity.FaceitId,
		Name:      entity.Name,
		Season:    entity.Season,
		Divisions: divisions,
	}
}

func (EseaMapper) Model(proto *pb.Esea, _ ...interface{}) model.EseaLeagueModel {
	divisions := make([]model.EseaDivisionModel, 0)

	for _, division := range proto.Divisions {
		divisions = append(divisions, EseaDivisionMapper{}.Model(division))
	}

	return model.EseaLeagueModel{
		Name:      proto.Name,
		FaceitId:  proto.FaceitId,
		Season:    proto.Season,
		Divisions: divisions,
	}
}

type EseaDivisionMapper struct{}

func (EseaDivisionMapper) Proto(entity model.EseaDivisionModel, _ ...interface{}) *pb.EseaDivision {
	standings := make([]*pb.EseaStanding, 0)

	for _, standing := range entity.Standings {
		standings = append(standings, EseaStandingMapper{}.Proto(standing))
	}

	var jsonData string
	err := json.Unmarshal([]byte(entity.PlayoffsData), &jsonData)
	if err != nil {
		logger.Error("Error mapping playoffs data")
		return nil
	}

	return &pb.EseaDivision{
		FaceitId:           entity.FaceitId,
		EseaLeagueFaceitId: entity.EseaLeagueFaceitId,
		Name:               entity.Name,
		Standings:          standings,
		Playoffs:           entity.Playoffs,
		PlayoffsData:       jsonData,
	}
}

func (EseaDivisionMapper) Model(proto *pb.EseaDivision, _ ...interface{}) model.EseaDivisionModel {
	standings := make([]model.EseaStandingModel, 0)

	for _, standing := range proto.Standings {
		standings = append(standings, EseaStandingMapper{}.Model(standing))
	}

	return model.EseaDivisionModel{
		FaceitId:           proto.FaceitId,
		Name:               proto.Name,
		EseaLeagueFaceitId: proto.EseaLeagueFaceitId,
		Standings:          standings,
		Playoffs:           proto.Playoffs,
		PlayoffsData:       model.JSONString(proto.PlayoffsData),
	}
}

type EseaStandingMapper struct{}

func (EseaStandingMapper) Proto(entity model.EseaStandingModel, _ ...interface{}) *pb.EseaStanding {
	return &pb.EseaStanding{
		IsDisqualified: entity.IsDisqualified,
		RankStart:      int32(entity.RankStart),
		RankEnd:        int32(entity.RankEnd),
		Points:         int32(entity.Points),
		MatchesPlayed:  int32(entity.MatchesPlayed),
		MatchesWon:     int32(entity.MatchesWon),
		MatchesLost:    int32(entity.MatchesLost),
		MatchesTied:    int32(entity.MatchesTied),
		BuchholzScore:  int32(entity.BuchholzScore),
		Team:           teams_mapper.TeamMapper{}.Proto(entity.Team),
	}
}

func (EseaStandingMapper) Model(proto *pb.EseaStanding, _ ...interface{}) model.EseaStandingModel {
	return model.EseaStandingModel{
		IsDisqualified: proto.IsDisqualified,
		RankStart:      int(proto.RankStart),
		RankEnd:        int(proto.RankEnd),
		Points:         int(proto.Points),
		MatchesPlayed:  int(proto.MatchesPlayed),
		MatchesWon:     int(proto.MatchesWon),
		MatchesLost:    int(proto.MatchesLost),
		MatchesTied:    int(proto.MatchesTied),
		BuchholzScore:  int(proto.BuchholzScore),
		Team:           teams_mapper.TeamMapper{}.Model(proto.Team),
	}
}
