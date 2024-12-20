package users_mapper

import (
	"ibercs/internal/model"
	players_mapper "ibercs/pkg/mapper/players"
	pb "ibercs/proto/users"
)

type UserMapper struct{}

func (UserMapper) Proto(entity model.UserModel, _ ...interface{}) *pb.User {
	return &pb.User{
		ID:          int32(entity.ID),
		PlayerID:    entity.FaceitId,
		Name:        entity.Name,
		Description: entity.Description,
		Twitter:     entity.Twitter,
		Twitch:      entity.Twitch,
		Role:        int32(entity.Role),
		Player:      players_mapper.PlayerMapper{}.Proto(entity.Player),
	}
}

func (UserMapper) Model(proto *pb.User, _ ...interface{}) model.UserModel {
	return model.UserModel{
		ID:          int(proto.ID),
		FaceitId:    proto.PlayerID,
		Name:        proto.Name,
		Description: proto.Description,
		Twitter:     proto.Twitter,
		Twitch:      proto.Twitch,
		Role:        int(proto.Role),
		Player:      players_mapper.PlayerMapper{}.Model(proto.Player),
	}
}

type SessionMapper struct{}

func (SessionMapper) Proto(entity model.UserSessionModel, _ ...interface{}) *pb.SessionResponse {
	return &pb.SessionResponse{
		Token:  entity.SessionID,
		UserId: int32(entity.UserID),
	}
}

func (SessionMapper) Model(proto *pb.SessionResponse, _ ...interface{}) model.UserSessionModel {
	return model.UserSessionModel{
		UserID:    int(proto.UserId),
		SessionID: proto.Token,
	}
}
