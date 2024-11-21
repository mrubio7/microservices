package requests

import (
	pb "ibercs/proto/tournaments"

	"github.com/invopop/validation"
)

type CreateOrganizerRequest struct {
	FaceitId string `json:"faceit_id"`
}

func (r CreateOrganizerRequest) validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.FaceitId, validation.Required),
	)
}

func (r CreateOrganizerRequest) ToProto() (*pb.NewOrganizerRequest, error) {
	err := r.validate()
	if err != nil {
		return nil, err
	}

	return &pb.NewOrganizerRequest{
		FaceitId: r.FaceitId,
	}, nil
}
