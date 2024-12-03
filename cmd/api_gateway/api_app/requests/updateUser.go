package requests

import (
	pb_users "ibercs/proto/users"

	"github.com/invopop/validation"
)

type UpdateUser struct {
	Twitter string `json:"twitter"`
	Twitch  string `json:"twitch"`
	Desc    string `json:"desc"`
}

func (req UpdateUser) validate() error {
	return validation.ValidateStruct(&req)
}

func (req UpdateUser) ToProto(userToUpdate *pb_users.User) (*pb_users.User, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	return &pb_users.User{
		ID:          userToUpdate.ID,
		PlayerID:    userToUpdate.PlayerID,
		Name:        userToUpdate.Name,
		Role:        userToUpdate.Role,
		Player:      userToUpdate.Player,
		Description: req.Desc,
		Twitter:     req.Twitter,
		Twitch:      req.Twitch,
	}, nil
}
