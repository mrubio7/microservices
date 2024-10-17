package microservice_players

import (
	"context"
	"fmt"
	pb "ibercs/proto/players"
)

type Server struct {
	pb.UnimplementedPlayerServiceServer
}

func (s *Server) GetPlayers(context.Context, *pb.Empty) (*pb.PlayerList, error) {
	fmt.Println("TOOOOOOOOOOMMMMMMAAAAAAA!!!")
	return nil, nil
}

func (s *Server) NewPlayer(context.Context, *pb.NewPlayerRequest) (*pb.PlayerResponse, error) {
	return nil, nil
}
