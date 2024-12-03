package faker

import (
	"ibercs/internal/model"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func GenerateMatch(seed int64) model.MatchModel {
	gofakeit.Seed(seed)

	return model.MatchModel{
		FaceitId:           gofakeit.UUID(),
		TeamAFaceitId:      gofakeit.UUID(),
		TeamAName:          gofakeit.Animal(),
		IsTeamAKnown:       gofakeit.Bool(),
		ScoreTeamA:         int32(gofakeit.Number(0, 16)),
		TeamBFaceitId:      gofakeit.UUID(),
		TeamBName:          gofakeit.Vegetable(),
		IsTeamBKnown:       gofakeit.Bool(),
		ScoreTeamB:         int32(gofakeit.Number(0, 16)),
		BestOf:             int32(gofakeit.Number(1, 3)), // Best of 1 or 3
		Timestamp:          gofakeit.DateRange(time.Now().Add(-time.Hour*24*30), time.Now()),
		Streams:            model.JSONStringArray{gofakeit.URL(), gofakeit.URL()},
		TournamentFaceitId: gofakeit.UUID(),
		TournamentName:     gofakeit.Company(),
		Map:                model.JSONStringArray{"de_dust2", "de_inferno"},
		Status:             "finished",
	}
}
