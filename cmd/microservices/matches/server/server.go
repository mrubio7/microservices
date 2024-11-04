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
			ID:                 int32(m.ID),
			FaceitId:           m.FaceitId,
			TeamAName:          m.TeamAName,
			TeamBName:          m.TeamBName,
			ScoreTeamA:         m.ScoreTeamA,
			ScoreTeamB:         m.ScoreTeamB,
			TournamentName:     m.TournamentName,
			TournamentFaceitId: m.TournamentFaceitId,
			BestOf:             m.BestOf,
			Map:                m.Map,
			Timestamp:          m.Timestamp.Unix(),
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
			ID:                 int32(m.ID),
			FaceitId:           m.FaceitId,
			TournamentName:     m.TournamentName,
			TournamentFaceitId: m.TournamentFaceitId,
			TeamAName:          m.TeamAName,
			TeamBName:          m.TeamBName,
			ScoreTeamA:         m.ScoreTeamA,
			ScoreTeamB:         m.ScoreTeamB,
			IsTeamAKnown:       m.IsTeamAKnown,
			IsTeamBKnown:       m.IsTeamBKnown,
			BestOf:             m.BestOf,
			Map:                m.Map,
			TeamA:              mapa[m.TeamAFaceitId],
			TeamB:              mapa[m.TeamBFaceitId],
			Timestamp:          m.Timestamp.Unix(),
		})
	}

	return &pb.MatchList{Matches: res}, nil
}

func (s *Server) GetMatchByFaceitId(ctx context.Context, req *pb.GetMatchRequest) (*pb.Match, error) {
	match := s.MatchesService.GetMatchByFaceitId(req.FaceitId)
	if match == nil {
		return nil, status.Errorf(codes.NotFound, "Error: match not found")
	}

	var teamA *pb_teams.Team
	var teamB *pb_teams.Team
	var err error

	if match.IsTeamAKnown {
		teamA, err = s.TeamsClient.GetTeamById(ctx, &pb_teams.NewTeamRequest{FaceitId: match.TeamAFaceitId})
		if err != nil {
			logger.Error("error getting team")
			return nil, err
		}
		teamB, err = s.TeamsClient.GetTeamFromFaceit(ctx, &pb_teams.NewTeamRequest{FaceitId: match.TeamBFaceitId})
		if err != nil {
			logger.Error("error getting team from faceit")
			return nil, err
		}
	}

	if match.IsTeamBKnown {
		teamB, err = s.TeamsClient.GetTeamById(ctx, &pb_teams.NewTeamRequest{FaceitId: match.TeamBFaceitId})
		if err != nil {
			logger.Error("error getting team")
			return nil, err
		}
		teamA, err = s.TeamsClient.GetTeamFromFaceit(ctx, &pb_teams.NewTeamRequest{FaceitId: match.TeamAFaceitId})
		if err != nil {
			logger.Error("error getting team from faceit")
			return nil, err
		}
	}

	res := &pb.Match{
		ID:                 int32(match.ID),
		FaceitId:           match.FaceitId,
		TeamAName:          match.TeamAName,
		TeamA:              teamA,
		TeamB:              teamB,
		TeamBName:          match.TeamBName,
		IsTeamAKnown:       match.IsTeamAKnown,
		IsTeamBKnown:       match.IsTeamBKnown,
		ScoreTeamA:         match.ScoreTeamA,
		ScoreTeamB:         match.ScoreTeamB,
		BestOf:             match.BestOf,
		TournamentName:     match.TournamentName,
		TournamentFaceitId: match.TournamentFaceitId,
		Map:                match.Map,
		Timestamp:          match.Timestamp.Unix(),
	}

	return res, nil
}
