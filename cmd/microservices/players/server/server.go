package microservice_players

import (
	"context"
	"ibercs/internal/model"
	"ibercs/pkg/config"
	"ibercs/pkg/consts"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/managers"
	"ibercs/pkg/mapper"
	"ibercs/pkg/microservices"
	pb "ibercs/proto/players"
	pb_users "ibercs/proto/users"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedPlayerServiceServer
	PlayerManager *managers.PlayerManager
	FaceitService *faceit.FaceitClient
	UserServer    pb_users.UserServiceClient
}

func registerUsersClient(cfg config.MicroserviceConfig) *pb_users.UserServiceClient {
	return microservices.New(cfg.Host_gRPC, cfg.Port_gRPC, pb_users.NewUserServiceClient)
}

func New(cfg, usersCfg config.MicroserviceConfig, cfgThirdParty config.ThirdPartyApiTokens) *Server {
	db := database.NewDatabase(cfg.Database)
	playerManager := managers.NewPlayerManager(db.GetDB())
	faceit := faceit.New(cfgThirdParty.FaceitApiToken)

	return &Server{
		PlayerManager: playerManager,
		FaceitService: faceit,
		UserServer:    *registerUsersClient(usersCfg),
	}
}

func (s *Server) GetPlayersByFaceitId(ctx context.Context, req *pb.GetPlayerRequest) (*pb.PlayerList, error) {
	playersRes := make([]*pb.Player, len(req.FaceitId))

	for i, id := range req.FaceitId {
		p, err := s.PlayerManager.GetByFaceitId(id)
		if err != nil {
			logger.Error("Error getting player: %s", err.Error())
			err := status.Errorf(codes.NotFound, "player not found")
			return nil, err
		}

		playerUpdated := s.FaceitService.GetPlayerAverageDetails(id, consts.LAST_MATCHES_NUMBER)
		playerUpdated.Id = p.Id
		playerUpdated.Stats.Id = p.Stats.Id
		err = s.PlayerManager.Update(playerUpdated)
		if err != nil {
			logger.Warning("Error updating player: %s", err.Error())
		}

		pbPlayer := mapper.Convert[model.PlayerModel, *pb.Player](*p)
		playersRes[i] = pbPlayer
	}

	return &pb.PlayerList{Players: playersRes}, nil
}

func (s *Server) GetPlayerByNickname(ctx context.Context, req *pb.GetPlayerByNicknameRequest) (*pb.Player, error) {
	p, err := s.PlayerManager.GetByNickname(req.Nickname)
	if err != nil {
		logger.Error("Error getting player: %s", err.Error())
		err := status.Errorf(codes.NotFound, "player not found")
		return nil, err
	}

	playerUpdated := s.FaceitService.GetPlayerAverageDetails(p.FaceitId, consts.LAST_MATCHES_NUMBER)
	playerUpdated.Id = p.Id
	playerUpdated.Stats.Id = p.Stats.Id
	err = s.PlayerManager.Update(playerUpdated)
	if err != nil {
		logger.Warning("Error updating player: %s", err.Error())
	}

	res, err := s.PlayerManager.GetByFaceitId(playerUpdated.FaceitId)
	if err != nil {
		logger.Error("Error getting player: %s", err.Error())
		err := status.Errorf(codes.NotFound, "player not found")
		return nil, err
	}

	pbPlayer := mapper.Convert[model.PlayerModel, *pb.Player](*res)

	return pbPlayer, nil
}

func (s *Server) GetAllPlayers(ctx context.Context, _ *pb.Empty) (*pb.PlayerList, error) {
	players, err := s.PlayerManager.GetAll()
	if err != nil {
		logger.Error("Error getting players: %s", err.Error())
		err := status.Errorf(codes.NotFound, "players not found")
		return nil, err
	}

	playersRes := make([]*pb.Player, len(players))
	for i, p := range players {
		pbPlayer := mapper.Convert[model.PlayerModel, *pb.Player](p)
		playersRes[i] = pbPlayer
	}

	return &pb.PlayerList{Players: playersRes}, nil
}

func (s *Server) CreatePlayerFromFaceitId(ctx context.Context, req *pb.CreatePlayerByFaceitIdRequest) (*pb.Player, error) {
	var player *model.PlayerModel
	player, _ = s.PlayerManager.GetByFaceitId(req.FaceitId)
	if player != nil {
		err := status.Errorf(codes.AlreadyExists, "Player already exist")
		return nil, err
	}

	playerData := s.FaceitService.GetPlayerAverageDetails(req.FaceitId, consts.LAST_MATCHES_NUMBER)
	if playerData == nil {
		err := status.Errorf(codes.NotFound, "Player not found")
		return nil, err
	}

	player, err := s.PlayerManager.Create(playerData)
	if err != nil {
		err := status.Errorf(codes.NotFound, "Error creating player")
		return nil, err
	}

	pbPlayer := mapper.Convert[model.PlayerModel, *pb.Player](*player)
	return pbPlayer, nil
}

// ProminentPlayers
func (s *Server) GetProminentPlayers(ctx context.Context, _ *pb.Empty) (*pb.ProminentPlayerList, error) {
	var prominentWeek *model.ProminentWeekModel

	prominentWeek, err := s.PlayerManager.GetProminentPlayers()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			prominentWeek, err = s.PlayerManager.GenerateProminentPlayers()
			if err != nil {
				logger.Error("Error generating prominent players: %s", err.Error())
				err := status.Errorf(codes.Internal, "error generating prominent players")
				return nil, err
			}
		} else {
			logger.Error("Error getting players: %s", err.Error())
			err := status.Errorf(codes.NotFound, "players not found")
			return nil, err
		}
	}

	prominentPlayersRes := make([]*pb.ProminentPlayer, len(prominentWeek.Players))
	for i, p := range prominentWeek.Players {
		pbPlayer := mapper.Convert[model.PlayerProminentModel, *pb.ProminentPlayer](p)
		prominentPlayersRes[i] = pbPlayer
	}

	return &pb.ProminentPlayerList{Players: prominentPlayersRes}, nil
}

// LookingForTeam
func (s *Server) GetAllLookingForTeam(ctx context.Context, _ *pb.Empty) (*pb.PlayerLookingForTeamList, error) {
	players, err := s.PlayerManager.GetLookingForTeamPlayers()
	if err != nil {
		logger.Error("Error getting players: %s", err.Error())
		err := status.Errorf(codes.NotFound, "players not found")
		return nil, err
	}

	playersRes := make([]*pb.PlayerLookingForTeam, len(players))
	for i, p := range players {
		playerData, err := s.PlayerManager.GetByFaceitId(p.FaceitId)
		if err != nil {
			logger.Error("Error getting player data")
			return nil, status.Errorf(codes.Internal, "Error getting player data")
		}
		p.Player = *playerData

		pbPlayer := mapper.Convert[model.LookingForTeamModel, *pb.PlayerLookingForTeam](p)
		playersRes[i] = pbPlayer
	}

	return &pb.PlayerLookingForTeamList{LookingForTeam: playersRes}, nil
}

func (s *Server) CreateLookingForTeam(ctx context.Context, req *pb.CreatePlayerLookingForTeamRequest) (*pb.PlayerLookingForTeam, error) {
	user, err := s.UserServer.GetUserById(ctx, &pb_users.GetUserByIdRequest{Id: req.UserId})
	if err != nil {
		logger.Error("Error getting user: %s", err.Error())
		err := status.Errorf(codes.Internal, "error creating looking for team")
		return nil, err
	}

	lookingForTeam := &model.LookingForTeamModel{
		InGameRole:   req.InGameRole,
		TimeTable:    req.TimeTable,
		OldTeams:     req.OldTeams,
		PlayingYears: req.PlayingYears,
		FaceitId:     user.PlayerID,
		Description:  req.Description,
		Id:           req.UserId,
	}

	lft, err := s.PlayerManager.CreateLookingForTeamPlayer(lookingForTeam)
	if err != nil {
		logger.Error("Error creating looking for team: %s", err.Error())
		err := status.Errorf(codes.Internal, "error creating looking for team")
		return nil, err
	}

	res := mapper.Convert[model.LookingForTeamModel, *pb.PlayerLookingForTeam](*lft)

	return res, nil
}

func (s *Server) UpdateLookingForTeam(ctx context.Context, req *pb.CreatePlayerLookingForTeamRequest) (*pb.PlayerLookingForTeam, error) {
	lookingForTeam := mapper.Convert[*pb.CreatePlayerLookingForTeamRequest, model.LookingForTeamModel](req)

	user, err := s.UserServer.GetUserByFaceitId(ctx, &pb_users.GetUserRequest{Id: lookingForTeam.FaceitId})
	if err != nil {
		logger.Error("Error getting user: %s", err.Error())
		err := status.Errorf(codes.NotFound, "user not found")
		return nil, err
	}

	if user.PlayerID != lookingForTeam.FaceitId {
		if user.Role < consts.ROLE_ADMIN {
			logger.Error("Error creating looking for team")
			return nil, status.Errorf(codes.PermissionDenied, "Error creating looking for team")
		}
	}

	err = s.PlayerManager.UpdateLookingForTeamPlayer(&lookingForTeam)
	if err != nil {
		logger.Error("Error updating looking for team: %s", err.Error())
		err := status.Errorf(codes.Internal, "error updating looking for team")
		return nil, err
	}

	res := mapper.Convert[model.LookingForTeamModel, *pb.PlayerLookingForTeam](lookingForTeam)

	return res, nil
}

func (s *Server) DeleteLookingForTeam(ctx context.Context, req *pb.DeleteLookingForTeamRequest) (*pb.Empty, error) {
	user, err := s.UserServer.GetUserById(ctx, &pb_users.GetUserByIdRequest{Id: req.UserId})
	if err != nil {
		logger.Error("Error getting user: %s", err.Error())
		err := status.Errorf(codes.Internal, "error deleting looking for team")
		return nil, err
	}

	err = s.PlayerManager.DeleteLookingForTeamPlayer(user.PlayerID)
	if err != nil {
		logger.Error("Error deleting looking for team: %s", err.Error())
		err := status.Errorf(codes.Internal, "error deleting looking for team")
		return nil, err
	}

	return nil, nil
}
