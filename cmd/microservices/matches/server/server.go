package microservice_matches

import (
	"context"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/service"
	pb "ibercs/proto/matches"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedMatchesServiceServer
	MatchesService *service.Matches
	TeamService    *service.Teams
	FaceitService  *faceit.FaceitClient
}

func New() *Server {
	cfg, err := config.Load()
	if err != nil {
		logger.Error("Unable to create grpc player server")
		return nil
	}
	db := database.New(cfg.Database)
	matchesService := service.NewMatchesService(db)
	teamsService := service.NewTeamsService(db)
	faceit := faceit.New(cfg.FaceitApiToken)

	return &Server{
		MatchesService: matchesService,
		TeamService:    teamsService,
		FaceitService:  faceit,
	}
}

func (s *Server) GetAllMatches(ctx context.Context, _ *pb.Empty) (*pb.MatchList, error) {
	matches := s.MatchesService.GetAllMatches()

	var res []*pb.Match
	for _, m := range matches {
		res = append(res, &pb.Match{
			ID:         int32(m.ID),
			FaceitId:   m.FaceitId,
			TeamAName:  m.TeamAName,
			TeamBName:  m.TeamBName,
			ScoreTeamA: m.ScoreTeamA,
			ScoreTeamB: m.ScoreTeamB,
			BestOf:     m.BestOf,
			Map:        m.Map,
			Timestamp:  m.Timestamp.Unix(),
		})
	}

	return &pb.MatchList{Matches: res}, nil
}

func (s *Server) GetMatchByFaceitId(ctx context.Context, req *pb.GetMatchRequest) (*pb.Match, error) {
	match := s.MatchesService.GetMatchByFaceitId(req.FaceitId)
	if match == nil {
		return nil, status.Errorf(codes.NotFound, "Error: match not found")
	}

	res := &pb.Match{
		ID:         int32(match.ID),
		FaceitId:   match.FaceitId,
		TeamAName:  match.TeamAName,
		TeamBName:  match.TeamBName,
		ScoreTeamA: match.ScoreTeamA,
		ScoreTeamB: match.ScoreTeamB,
		BestOf:     match.BestOf,
		Map:        match.Map,
		Timestamp:  match.Timestamp.Unix(),
	}

	return res, nil
}
