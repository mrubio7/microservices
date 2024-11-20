package faker

import (
	"ibercs/internal/model"

	"github.com/brianvoe/gofakeit/v6"
)

func GeneratePlayer(seed int64) model.PlayerModel {
	gofakeit.Seed(seed)

	return model.PlayerModel{
		ID:       int32(gofakeit.Number(0, 16)),
		FaceitId: gofakeit.UUID(),
		Nickname: gofakeit.Animal(),
		SteamId:  gofakeit.UUID(),
		Avatar:   gofakeit.URL(),
		Stats: model.PlayerStatsModel{
			ID:                     int32(gofakeit.Number(0, 16)),
			KrRatio:                gofakeit.Float32Range(0.5, 1.5),
			KdRatio:                gofakeit.Float32Range(0.5, 1.5),
			KillsAverage:           gofakeit.Float32Range(0.5, 1.5),
			DeathsAverage:          gofakeit.Float32Range(0.5, 1.5),
			HeadshotPercentAverage: gofakeit.Float32Range(0.5, 1.5),
			MVPAverage:             gofakeit.Float32Range(0.5, 1.5),
			AssistAverage:          gofakeit.Float32Range(0.5, 1.5),
			TripleKillsAverage:     gofakeit.Float32Range(0.5, 1.5),
			QuadroKillsAverage:     gofakeit.Float32Range(0.5, 1.5),
			PentaKillsAverage:      gofakeit.Float32Range(0.5, 1.5),
			Elo:                    int32(gofakeit.Number(500, 3000)),
		},
	}
}

func GenerateProminentWeek(seed int64) model.ProminentWeekModel {
	gofakeit.Seed(seed)

	return model.ProminentWeekModel{
		ID:   int32(gofakeit.Number(0, 16)),
		Year: int16(gofakeit.Number(23, 26)),
		Week: int16(gofakeit.Number(1, 40)),
		Players: []model.PlayerProminentModel{
			GenerateProminentPlayer(seed),
			GenerateProminentPlayer(seed),
			GenerateProminentPlayer(seed),
			GenerateProminentPlayer(seed),
			GenerateProminentPlayer(seed),
		},
	}
}

func GenerateProminentPlayer(seed int64) model.PlayerProminentModel {
	gofakeit.Seed(seed)

	return model.PlayerProminentModel{
		ID:       int32(gofakeit.Number(0, 16)),
		FaceitId: gofakeit.UUID(),
		Nickname: gofakeit.Animal(),
		SteamId:  gofakeit.UUID(),
		Avatar:   gofakeit.URL(),
		Score:    gofakeit.Float32Range(5, 50),
	}
}

func GenerateLookingForTeam(seed int64) model.LookingForTeamModel {
	gofakeit.Seed(seed)

	return model.LookingForTeamModel{
		Id:           int32(gofakeit.Number(0, 16)),
		FaceitId:     gofakeit.UUID(),
		InGameRole:   model.JSONStringArray{gofakeit.Adjective(), gofakeit.Adjective()},
		TimeTable:    gofakeit.Date().String(),
		OldTeams:     gofakeit.Company(),
		PlayingYears: int32(gofakeit.Number(2, 20)),
		BornDate:     gofakeit.Date(),
		Description:  gofakeit.Sentence(10),
		CreatedAt:    gofakeit.Date().Unix(),
		UpdatedAt:    gofakeit.Date().Unix(),
	}
}
