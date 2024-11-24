package faker

import (
	"ibercs/internal/model"

	"github.com/brianvoe/gofakeit/v6"
)

func GenerateEseaLeague(seed int64) model.EseaLeagueModel {
	gofakeit.Seed(seed)

	return model.EseaLeagueModel{
		FaceitId: gofakeit.UUID(),
		Name:     gofakeit.Company(),
		Season:   int32(gofakeit.Number(0, 16)),
		Divisions: []model.EseaDivisionModel{
			GenerateEseaDivision(seed),
			GenerateEseaDivision(seed),
		},
	}
}

func GenerateEseaDivision(seed int64) model.EseaDivisionModel {
	gofakeit.Seed(seed)

	return model.EseaDivisionModel{
		FaceitId:     gofakeit.UUID(),
		DivisionName: gofakeit.Company(),
		StageName:    gofakeit.Company(),
		Standings: []model.EseaStandingModel{
			GenerateEseaStanding(seed),
			GenerateEseaStanding(seed),
		},
		Playoffs:     gofakeit.Bool(),
		PlayoffsData: model.JSONString(gofakeit.UUID()),
	}
}

func GenerateEseaStanding(seed int64) model.EseaStandingModel {
	gofakeit.Seed(seed)

	return model.EseaStandingModel{
		TeamFaceitId:   gofakeit.UUID(),
		IsDisqualified: gofakeit.Bool(),
		TournamentName: gofakeit.Company(),
		RankStart:      gofakeit.Number(0, 16),
		RankEnd:        gofakeit.Number(0, 16),
		Points:         gofakeit.Number(0, 16),
		MatchesPlayed:  gofakeit.Number(0, 16),
		MatchesWon:     gofakeit.Number(0, 16),
		MatchesLost:    gofakeit.Number(0, 16),
		MatchesTied:    gofakeit.Number(0, 16),
		BuchholzScore:  gofakeit.Number(0, 16),
	}
}
