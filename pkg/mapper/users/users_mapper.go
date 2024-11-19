package users_mapper

import (
	"psonder/internal/models"
	pb "psonder/proto/users"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserMapper struct{}

func (UserMapper) Proto(model models.UserModel) *pb.User {
	return &pb.User{
		ID:        int32(model.ID),
		Name:      model.Name,
		Email:     model.Email,
		BornDate:  timestamppb.New(model.BornDate.Local()),
		CreatedAt: timestamppb.New(model.CreatedAt.Local()),
		UpdatedAt: timestamppb.New(model.UpdatedAt.Local()),
		DeletedAt: timestamppb.New(model.DeletedAt.Local()),
	}
}

func (UserMapper) Model(proto *pb.User) models.UserModel {
	return models.UserModel{
		ID:        int32(proto.ID),
		Name:      proto.Name,
		Email:     proto.Email,
		BornDate:  proto.BornDate.AsTime().Local(),
		CreatedAt: proto.CreatedAt.AsTime().Local(),
		UpdatedAt: proto.UpdatedAt.AsTime().Local(),
		DeletedAt: proto.DeletedAt.AsTime().Local(),
	}
}
