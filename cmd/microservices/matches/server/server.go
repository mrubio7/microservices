package microservice_matches

import (
	"context"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/microservices"
	"ibercs/pkg/service"
	pb "ibercs/proto/matches"
	pb_teams "ibercs/proto/teams"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedMatchesServiceServer
	MatchesService *service.Matches
	TeamsClient    pb_teams.TeamServiceClient
	FaceitService  *faceit.FaceitClient
}

func registerTeamsClient(cfg config.MicroservicesConfig) *pb_teams.TeamServiceClient {
	return microservices.New(cfg.TeamsHost, cfg.TeamsPort, pb_teams.NewTeamServiceClient)
}

func New() *Server {
	cfg, err := config.Load()
	if err != nil {
		logger.Error("Unable to create grpc player server")
		return nil
	}
	db := database.New(cfg.Database)
	matchesService := service.NewMatchesService(db)
	faceit := faceit.New(cfg.FaceitApiToken)

	return &Server{
		MatchesService: matchesService,
		TeamsClient:    *registerTeamsClient(cfg.Microservices),
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

func (s *Server) GetUpcomingMatches(ctx context.Context, req *pb.GetUpcomingRequest) (*pb.MatchList, error) {
	matches := s.MatchesService.GetAllMatches()
	teams, err := s.TeamsClient.GetTeams(ctx, &pb_teams.GetTeamsRequest{Active: true})
	if err != nil {
		return nil, err
	}

	mapa := make(map[string]*pb_teams.Team, len(teams.Teams))

	for _, m := range teams.Teams {
		mapa[m.FaceitId] = m
	}

	var res []*pb.Match
	days := time.Duration(req.Days) * 24 * time.Hour
	now := time.Now()
	for _, m := range matches {
		if m.Timestamp.Before(now.Add(-days)) || m.Timestamp.After(now.Add(days)) {
			continue
		}

		res = append(res, &pb.Match{
			ID:           int32(m.ID),
			FaceitId:     m.FaceitId,
			TeamAName:    m.TeamAName,
			TeamBName:    m.TeamBName,
			ScoreTeamA:   m.ScoreTeamA,
			ScoreTeamB:   m.ScoreTeamB,
			IsTeamAKnown: m.IsTeamAKnown,
			IsTeamBKnown: m.IsTeamBKnown,
			BestOf:       m.BestOf,
			Map:          m.Map,
			TeamA:        mapa[m.TeamAFaceitId],
			TeamB:        mapa[m.TeamBFaceitId],
			Timestamp:    m.Timestamp.Unix(),
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
