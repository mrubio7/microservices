package microservice_teams

import (
	"context"
	"ibercs/internal/model"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/managers"
	"ibercs/pkg/mapper"
	"ibercs/pkg/microservices"
	pb "ibercs/proto/teams"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedTeamServiceServer
	TeamsManager  *managers.TeamManager
	TeamsClient   pb.TeamServiceClient
	FaceitService *faceit.FaceitClient
}

func registerTeamsClient(cfg config.MicroserviceConfig) *pb.TeamServiceClient {
	return microservices.New(cfg.Host_gRPC, cfg.Port_gRPC, pb.NewTeamServiceClient)
}

func New(cfg config.MicroserviceConfig, cfgThirdParty config.ThirdPartyApiTokens) *Server {
	db := database.NewDatabase(cfg.Database)
	teamsService := managers.NewTeamManager(db.GetDB())
	faceit := faceit.New(cfgThirdParty.FaceitApiToken)

	return &Server{
		TeamsManager:  teamsService,
		TeamsClient:   *registerTeamsClient(cfg),
		FaceitService: faceit,
	}
}

func (s *Server) GetAllTeams(ctx context.Context, _ *pb.Empty) (*pb.TeamList, error) {
	teams, err := s.TeamsManager.GetAll()
	if err != nil {
		logger.Error("Error getting all teams: %v", err)
		err := status.Errorf(codes.Internal, "Error getting all teams: %v", err)
		return nil, err
	}

	var res []*pb.Team
	for _, team := range teams {
		pbTeam := mapper.Convert[model.TeamModel, *pb.Team](team)
		res = append(res, pbTeam)
	}

	return &pb.TeamList{Teams: res}, nil
}

func (s *Server) GetActiveTeams(ctx context.Context, _ *pb.Empty) (*pb.TeamList, error) {
	teams, err := s.TeamsManager.GetActiveTeams()
	if err != nil {
		logger.Error("Error getting active teams: %v", err)
		err := status.Errorf(codes.Internal, "Error getting active teams: %v", err)
		return nil, err
	}

	var res []*pb.Team
	for _, team := range teams {
		pbTeam := mapper.Convert[model.TeamModel, *pb.Team](team)
		res = append(res, pbTeam)
	}

	return &pb.TeamList{Teams: res}, nil
}

func (s *Server) GetById(ctx context.Context, req *pb.GetTeamByIdRequest) (*pb.Team, error) {
	team, err := s.TeamsManager.GetById(int(req.Id))
	if err != nil {
		return nil, err
	}

	res := mapper.Convert[model.TeamModel, *pb.Team](*team)

	return res, nil
}

func (s *Server) GetByNickname(ctx context.Context, req *pb.GetTeamByNicknameRequest) (*pb.Team, error) {
	team, err := s.TeamsManager.GetByNickname(req.Nickname)
	if err != nil {
		return nil, err
	}

	res := mapper.Convert[model.TeamModel, *pb.Team](*team)

	return res, nil
}

func (s *Server) CreateFromFaceit(ctx context.Context, req *pb.NewTeamFromFaceitRequest) (*pb.Team, error) {
	faceitTeam := s.FaceitService.GetTeamById(req.FaceitId)
	if faceitTeam == nil {
		logger.Error("Team with faceit id %s not found", req.FaceitId)
		err := status.Errorf(codes.NotFound, "Team with faceit id %s not found", req.FaceitId)
		return nil, err
	}

	createdTeam, err := s.TeamsManager.Create(faceitTeam)
	if err != nil {
		logger.Error("Error creating team: %v", err)
		err := status.Errorf(codes.Internal, "Error creating team: %v", err)
		return nil, err
	}

	res := mapper.Convert[model.TeamModel, *pb.Team](*createdTeam)

	return res, nil
}

func (s *Server) Update(ctx context.Context, req *pb.NewTeamFromFaceitRequest) (*pb.Team, error) {
	team, err := s.TeamsManager.GetByFaceitId(req.FaceitId)
	if err != nil {
		logger.Error("Team with faceit id %s not found", req.FaceitId)
		err := status.Errorf(codes.NotFound, "Team with faceit id %s not found", req.FaceitId)
		return nil, err
	}

	updatedTeam := s.FaceitService.GetTeamById(team.FaceitId)
	if updatedTeam == nil {
		logger.Warning("Team with faceit id %s not found", req.FaceitId)
		err := s.TeamsManager.DesactivateTeam(int(team.Id))
		if err != nil {
			logger.Error("Error desactivating team: %v", err)
			err := status.Errorf(codes.Internal, "Error desactivating team: %v", err)
			return nil, err
		}

		err = status.Errorf(codes.NotFound, "Team with faceit id %s not found, Team desactivated", req.FaceitId)
		return nil, err
	}

	err = s.TeamsManager.Update(updatedTeam)
	if err != nil {
		logger.Error("Error updating team: %v", err)
		err := status.Errorf(codes.Internal, "Error updating team: %v", err)
		return nil, err
	}

	res := mapper.Convert[model.TeamModel, *pb.Team](*updatedTeam)

	return res, nil
}

func (s *Server) FindTeamsByPlayerId(ctx context.Context, req *pb.GetTeamByPlayerIdRequest) (*pb.TeamList, error) {
	teams, err := s.TeamsManager.GetByPlayerId(req.PlayerId)
	if err != nil {
		logger.Error("Team with player id %s not found", req.PlayerId)
		err := status.Errorf(codes.NotFound, "Team with player id %s not found", req.PlayerId)
		return nil, err
	}

	var res []*pb.Team
	for _, team := range teams {
		pbTeam := mapper.Convert[model.TeamModel, *pb.Team](team)
		res = append(res, pbTeam)
	}

	return &pb.TeamList{Teams: res}, nil
}
