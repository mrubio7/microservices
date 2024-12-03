package faker

import (
	"ibercs/internal/model"

	"github.com/brianvoe/gofakeit/v6"
)

func GenerateUser(seed int64) model.UserModel {
	return model.UserModel{
		ID:             gofakeit.Number(0, 16),
		FaceitId:       gofakeit.UUID(),
		Name:           gofakeit.Name(),
		Description:    gofakeit.Sentence(10),
		Twitter:        gofakeit.URL(),
		Twitch:         gofakeit.URL(),
		Role:           0,
		IsProfessional: gofakeit.Bool(),
		Player:         GeneratePlayer(seed),
	}
}
