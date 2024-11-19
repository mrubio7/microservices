package mapper

import (
	"psonder/internal/models"
	users_mapper "psonder/pkg/mapper/users"
	pb "psonder/proto/users"
)

func RegisterMappers() {
	// Mapper para User
	Register(Mapper[models.UserModel, *pb.User]{
		From: users_mapper.UserMapper{}.Proto,
		To:   users_mapper.UserMapper{}.Model,
	})

	Register(Mapper[*pb.User, models.UserModel]{
		From: users_mapper.UserMapper{}.Model,
		To:   users_mapper.UserMapper{}.Proto,
	})

}
