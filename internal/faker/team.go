package faker

import (
	"ibercs/internal/model"

	"github.com/brianvoe/gofakeit/v6"
)

func GenerateTeam(seed int64) model.TeamModel {
	gofakeit.Seed(seed)

	return model.TeamModel{
		Id:        int32(gofakeit.Number(0, 16)),
		FaceitId:  gofakeit.UUID(),
		Name:      gofakeit.Company(),
		Nickname:  gofakeit.Color(),
		Avatar:    gofakeit.URL(),
		Active:    false,
		Twitter:   gofakeit.URL(),
		Instagram: gofakeit.URL(),
		Web:       gofakeit.URL(),
		Tournaments: model.JSONStringArray{
			gofakeit.UUID(),
			gofakeit.UUID(),
		},
		PlayersId: model.JSONStringArray{
			gofakeit.UUID(),
			gofakeit.UUID(),
			gofakeit.UUID(),
			gofakeit.UUID(),
			gofakeit.UUID(),
		},
		Stats: model.TeamStatsModel{
			ID:           int32(gofakeit.Number(0, 16)),
			TotalMatches: int32(gofakeit.Number(0, 16)),
			Wins:         int32(gofakeit.Number(0, 16)),
			Winrate:      gofakeit.Float32Range(0.5, 1.5),
			RecentResults: model.JSONInt32Slice{
				int32(gofakeit.Number(0, 1)),
				int32(gofakeit.Number(0, 1)),
				int32(gofakeit.Number(0, 1)),
				int32(gofakeit.Number(0, 1)),
				int32(gofakeit.Number(0, 1)),
			},
			MapStats: model.JSONMapStats{
				"de_dust2": model.TeamMapStats{
					MapName: "de_dust2",
					Matches: int32(gofakeit.Number(0, 16)),
					WinRate: int32(gofakeit.Number(0, 100)),
				},
				"de_inferno": model.TeamMapStats{
					MapName: "de_inferno",
					Matches: int32(gofakeit.Number(0, 16)),
					WinRate: int32(gofakeit.Number(0, 100)),
				},
				"de_anubis": model.TeamMapStats{
					MapName: "de_anubis",
					Matches: int32(gofakeit.Number(0, 16)),
					WinRate: int32(gofakeit.Number(0, 100)),
				},
			},
		},
	}
}
