package requests

import (
	pb_players "ibercs/proto/players"

	"github.com/invopop/validation"
)

type DeleteLookingForTeam struct {
	PlayerId string `json:"player_id"`
}

func (req DeleteLookingForTeam) validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.PlayerId, validation.Required),
	)
}

func (req DeleteLookingForTeam) ToProto(identity int32) (*pb_players.DeleteLookingForTeamRequest, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	return &pb_players.DeleteLookingForTeamRequest{
		PlayerId: req.PlayerId,
		UserId:   int32(identity),
	}, nil
}
