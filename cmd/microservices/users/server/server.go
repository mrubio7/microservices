package server

import (
	"context"
	"errors"
	"fmt"
	"ibercs/internal/model"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/microservices"
	"ibercs/pkg/service"
	pb_players "ibercs/proto/players"
	pb "ibercs/proto/users"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	FaceitService *faceit.FaceitClient
	UsersService  *service.Users
	PlayersClient pb_players.PlayerServiceClient
}

func registerPlayersClient(cfg config.MicroservicesConfig) *pb_players.PlayerServiceClient {
	return microservices.New(cfg.PlayersHost, cfg.PlayersPort, pb_players.NewPlayerServiceClient)
}

func New() *Server {
	cfg, err := config.Load()
	if err != nil {
		logger.Error("Unable to create grpc player server")
		return nil
	}
	db := database.New(cfg.Database)

	return &Server{
		FaceitService: faceit.New(cfg.FaceitApiToken),
		UsersService:  service.NewUsersService(db),
		PlayersClient: *registerPlayersClient(cfg.Microservices),
	}
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user := s.UsersService.GetUserById(req.Id)
	if user == nil {
		return nil, errors.New("unable to get user")
	}

	player, err := s.PlayersClient.GetPlayer(ctx, &pb_players.GetPlayerRequest{FaceitId: []string{user.FaceitID}})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	result := &pb.User{
		ID:          int32(user.ID),
		PlayerID:    user.FaceitID,
		Name:        user.Name,
		Description: user.Description,
		Twitter:     user.Twitter,
		Twitch:      user.Twitch,
		Role:        int32(user.Role),
		Player:      player.Players[0],
	}

	return result, nil
}

func (s *Server) GetUserByFaceitId(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user := s.UsersService.GetUserByFaceitId(req.Id)
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user with FaceitID %s not found", req.Id)
	}

	player, err := s.PlayersClient.GetPlayer(ctx, &pb_players.GetPlayerRequest{FaceitId: []string{user.FaceitID}})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	result := &pb.User{
		ID:          int32(user.ID),
		PlayerID:    user.FaceitID,
		Name:        user.Name,
		Description: user.Description,
		Twitter:     user.Twitter,
		Twitch:      user.Twitch,
		Role:        int32(user.Role),
		Player:      player.Players[0],
	}

	return result, nil
}

func (s *Server) GetUserByPlayerNickname(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user := s.UsersService.GetUserByPlayerNickname(req.Id)
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user with FaceitID %s not found", req.Id)
	}

	result := &pb.User{
		ID:          int32(user.ID),
		PlayerID:    user.FaceitID,
		Name:        user.Name,
		Description: user.Description,
		Twitter:     user.Twitter,
		Twitch:      user.Twitch,
		Role:        int32(user.Role),
		Player: &pb_players.Player{
			Id:       user.Player.ID,
			Nickname: user.Player.Nickname,
			FaceitId: user.Player.FaceitId,
			SteamId:  user.Player.SteamId,
			Avatar:   user.Player.Avatar,
			Stats: &pb_players.PlayerStats{
				PlayerId:               user.Player.Stats.ID,
				KdRatio:                user.Player.Stats.KdRatio,
				KrRatio:                user.Player.Stats.KrRatio,
				KillsAverage:           user.Player.Stats.KillsAverage,
				DeathsAverage:          user.Player.Stats.DeathsAverage,
				HeadshotPercentAverage: user.Player.Stats.HeadshotPercentAverage,
				MVPAverage:             user.Player.Stats.MVPAverage,
				AssistAverage:          user.Player.Stats.AssistAverage,
				Elo:                    user.Player.Stats.Elo,
			},
		},
	}

	return result, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user := model.UserModel{
		ID:          int(req.ID),
		FaceitID:    req.PlayerID,
		Name:        req.Name,
		Description: req.Description,
		Twitter:     req.Twitter,
		Twitch:      req.Twitch,
		Role:        int(req.Role),
	}

	res := s.UsersService.UpdateUser(user)
	if res == nil {
		err := fmt.Errorf("error updating user %d", user.ID)
		logger.Error(err.Error())
		return nil, err
	}

	userRes := &pb.User{
		ID:          int32(res.ID),
		PlayerID:    res.FaceitID,
		Name:        res.Name,
		Description: res.Description,
		Twitter:     res.Twitter,
		Twitch:      res.Twitch,
		Role:        int32(res.Role),
	}

	return userRes, nil
}

func (s *Server) NewUser(ctx context.Context, req *pb.NewUserRequest) (*pb.User, error) {
	player, err := s.PlayersClient.GetPlayer(ctx, &pb_players.GetPlayerRequest{FaceitId: []string{req.FaceitId}})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	user := &model.UserModel{
		FaceitID: player.Players[0].FaceitId,
		Name:     player.Players[0].Nickname,
	}

	user = s.UsersService.NewUser(*user)
	if user == nil {
		err := fmt.Errorf("unable to create user %s", player.Players[0].FaceitId)
		logger.Error(err.Error())
		return nil, err
	}

	res := &pb.User{
		ID:       int32(user.ID),
		PlayerID: user.FaceitID,
		Name:     user.Name,
		Player:   player.Players[0],
	}

	return res, err
}

func (s *Server) NewSession(ctx context.Context, req *pb.NewSessionRequest) (*pb.NewSessionResponse, error) {
	token := s.UsersService.NewSession(int(req.Id))

	if token == "" {
		return &pb.NewSessionResponse{Response: ""}, nil
	}
	return &pb.NewSessionResponse{Response: token}, nil
}

func (s *Server) DeleteSession(ctx context.Context, req *pb.NewSessionRequest) (*pb.NewSessionResponse, error) {
	token := s.UsersService.DeleteSession(int(req.Id))

	if token == "" {
		return &pb.NewSessionResponse{Response: ""}, nil
	}
	return &pb.NewSessionResponse{Response: token}, nil
}

func (s *Server) GetAllStreams(ctx context.Context, _ *pb.Empty) (*pb.StreamsResponse, error) {
	users := s.UsersService.GetAllStreams()
	if users == nil {
		logger.Error("Error taking streams")
		return nil, status.Errorf(codes.NotFound, "Streams not found")
	}

	var streams []*pb.StreamResponse
	for _, u := range users {
		streams = append(streams, &pb.StreamResponse{Stream: u.Twitch, Name: u.Name})
	}

	return &pb.StreamsResponse{Streams: streams}, nil
}
