package requests

import (
	pb "ibercs/proto/teams"

	"github.com/invopop/validation"
)

type CreateTeamFromFaceitRequest struct {
	FaceitId string `json:"faceit_id"`
}

func (req CreateTeamFromFaceitRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.FaceitId, validation.Required),
	)
}

func (req CreateTeamFromFaceitRequest) ToProto() (*pb.NewTeamFromFaceitRequest, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	return &pb.NewTeamFromFaceitRequest{
		FaceitId: req.FaceitId,
	}, nil
}
