package microservice_matches

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
	pb "ibercs/proto/matches"
	pb_teams "ibercs/proto/teams"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedMatchesServiceServer
	MatchesManager *managers.MatchManager
	TeamsClient    pb_teams.TeamServiceClient
	FaceitService  *faceit.FaceitClient
}

func registerTeamsClient(cfg config.MicroserviceConfig) *pb_teams.TeamServiceClient {
	return microservices.New(cfg.Host_gRPC, cfg.Port_gRPC, pb_teams.NewTeamServiceClient)
}

func New(cfg config.MicroserviceConfig, cfgThirdParty config.ThirdPartyApiTokens) *Server {
	db := database.NewDatabase(cfg.Database)
	matchesService := managers.NewMatchManager(db.GetDB())
	faceit := faceit.New(cfgThirdParty.FaceitApiToken)

	return &Server{
		MatchesManager: matchesService,
		TeamsClient:    *registerTeamsClient(cfg),
		FaceitService:  faceit,
	}
}

func (s *Server) GetAllMatches(ctx context.Context, _ *pb.Empty) (*pb.MatchList, error) {
	matches, err := s.MatchesManager.GetAll()
	if err != nil {
		logger.Error(err.Error())
		err := status.Errorf(codes.NotFound, "matches not found")
		return nil, err
	}

	var pbMatches []*pb.Match
	for _, match := range matches {
		m := mapper.Convert[model.MatchModel, *pb.Match](match)
		pbMatches = append(pbMatches, m)
	}

	return &pb.MatchList{Matches: pbMatches}, nil
}

func (s *Server) GetMatchByFaceitId(ctx context.Context, req *pb.GetMatchRequest) (*pb.Match, error) {
	match, err := s.MatchesManager.GetMatchByFaceitId(req.FaceitId)
	if err != nil {
		logger.Error(err.Error())
		err := status.Errorf(codes.NotFound, "match not found")
		return nil, err
	}

	return mapper.Convert[model.MatchModel, *pb.Match](*match), nil
}

func (s *Server) GetMatchesByTeamId(ctx context.Context, req *pb.GetMatchRequest) (*pb.MatchList, error) {
	teamMatches, err := s.MatchesManager.GetMatchesByTeamId(req.FaceitId)
	if err != nil {
		logger.Error(err.Error())
		err := status.Errorf(codes.NotFound, "matches not found")
		return nil, err
	}

	var pbMatches []*pb.Match
	for _, match := range teamMatches {
		m := mapper.Convert[model.MatchModel, *pb.Match](match)
		pbMatches = append(pbMatches, m)
	}

	return &pb.MatchList{Matches: pbMatches}, nil
}
