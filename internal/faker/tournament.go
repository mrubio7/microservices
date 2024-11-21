package faker

import (
	"ibercs/internal/model"

	"github.com/brianvoe/gofakeit/v6"
)

func GenerateTournament(seed int64) model.TournamentModel {
	gofakeit.Seed(seed)

	return model.TournamentModel{
		Id:              int32(gofakeit.Number(0, 16)),
		FaceitId:        gofakeit.UUID(),
		Name:            gofakeit.Company(),
		OrganizerId:     gofakeit.UUID(),
		BackgroundImage: gofakeit.URL(),
		CoverImage:      gofakeit.URL(),
		Avatar:          gofakeit.URL(),
		RegisterDate:    gofakeit.Date(),
		StartDate:       gofakeit.Date(),
		Status:          "live",
		JoinPolicy:      "open",
		GeoCountries: model.JSONStringArray{
			gofakeit.Country(),
			gofakeit.Country(),
		},
		MinLevel: 1,
		MaxLevel: 10,
		Type:     "ESEA",
		TeamsId: model.JSONStringArray{
			gofakeit.UUID(),
			gofakeit.UUID(),
		},
	}
}

func GenerateOrganizer(seed int64) model.OrganizerModel {
	gofakeit.Seed(seed)

	return model.OrganizerModel{
		Id:       int32(gofakeit.Number(0, 16)),
		FaceitId: gofakeit.UUID(),
		Name:     gofakeit.Company(),
		Website:  gofakeit.URL(),
		Twitter:  gofakeit.URL(),
		Twitch:   gofakeit.URL(),
		Avatar:   gofakeit.URL(),
		Type:     "organizer",
	}
}
