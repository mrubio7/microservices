package players_mapper

import (
	"ibercs/internal/model"
	pb "ibercs/proto/players"
)

type PlayerMapper struct{}

func (PlayerMapper) Proto(model model.PlayerModel) *pb.Player {
	return &pb.Player{
		Id: int32(model.ID),
	}
}

func (PlayerMapper) Model(proto *pb.Player) model.PlayerModel {
	return model.PlayerModel{
		ID: int32(proto.Id),
	}
}
