package requests

import (
	pb_players "ibercs/proto/players"

	"github.com/invopop/validation"
)

type CreateLookingForTeam struct {
	InGameRole   []string `json:"in_game_role"`
	TimeTable    string   `json:"time_table"`
	OldTeams     string   `json:"old_teams"`
	PlayingYears int32    `json:"playing_years"`
	Description  string   `json:"description"`
	FaceitId     string   `json:"faceit_id"`
	UserId       int32    `json:"-"`
}

func (req CreateLookingForTeam) validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.InGameRole, validation.Required),
	)
}

func (req CreateLookingForTeam) ToProto(identity int32) (*pb_players.CreatePlayerLookingForTeamRequest, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	return &pb_players.CreatePlayerLookingForTeamRequest{
		InGameRole:   req.InGameRole,
		TimeTable:    req.TimeTable,
		OldTeams:     req.OldTeams,
		PlayingYears: req.PlayingYears,
		Description:  req.Description,
		FaceitId:     req.FaceitId,
		UserId:       identity,
	}, nil
}
