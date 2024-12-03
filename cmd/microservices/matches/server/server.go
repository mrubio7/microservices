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
	"gorm.io/gorm"
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

func New(cfg config.MicroserviceConfig, cfgTeams config.MicroserviceConfig, cfgThirdParty config.ThirdPartyApiTokens) *Server {
	db := database.NewDatabase(cfg.Database)
	matchesService := managers.NewMatchManager(db.GetDB())
	faceit := faceit.New(cfgThirdParty.FaceitApiToken)

	return &Server{
		MatchesManager: matchesService,
		TeamsClient:    *registerTeamsClient(cfgTeams),
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

	var teamA, teamB *pb_teams.Team
	if match.IsTeamAKnown {
		teamA, err = s.TeamsClient.GetByFaceitId(ctx, &pb_teams.GetTeamByFaceitIdRequest{FaceitId: match.TeamAFaceitId})
		if err != nil {
			logger.Error(err.Error())
			err := status.Errorf(codes.NotFound, "teamA not found")
			return nil, err
		}
	} else {
		teamA, err = s.TeamsClient.GetTeamFromFaceit(ctx, &pb_teams.GetTeamFromFaceitRequest{FaceitId: match.TeamAFaceitId})
		if err != nil {
			logger.Error(err.Error())
			err := status.Errorf(codes.NotFound, "teamA not found")
			return nil, err
		}
	}

	if match.IsTeamBKnown {
		teamB, err = s.TeamsClient.GetByFaceitId(ctx, &pb_teams.GetTeamByFaceitIdRequest{FaceitId: match.TeamBFaceitId})
		if err != nil {
			logger.Error(err.Error())
			err := status.Errorf(codes.NotFound, "teamB not found")
			return nil, err
		}
	} else {
		teamB, err = s.TeamsClient.GetTeamFromFaceit(ctx, &pb_teams.GetTeamFromFaceitRequest{FaceitId: match.TeamBFaceitId})
		if err != nil {
			logger.Error(err.Error())
			err := status.Errorf(codes.NotFound, "teamB not found")
			return nil, err
		}
	}

	res := mapper.Convert[model.MatchModel, *pb.Match](*match)

	res.TeamA = teamA
	res.TeamB = teamB
	return res, nil
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

func (s *Server) GetNearbyMatches(ctx context.Context, req *pb.GetNearbyMatchesRequest) (*pb.MatchList, error) {
	matches, err := s.MatchesManager.GetNearbyMatches(int(req.Days))
	if err != nil {
		logger.Error(err.Error())
		err := status.Errorf(codes.NotFound, "matches not found")
		return nil, err
	}

	var pbMatches []*pb.Match
	for _, match := range matches {
		if match.IsTeamAKnown {
			teamA, err := s.TeamsClient.GetByFaceitId(ctx, &pb_teams.GetTeamByFaceitIdRequest{FaceitId: match.TeamAFaceitId})
			if err != nil {
				err := status.Errorf(codes.NotFound, "teamA %s not found", match.TeamAName)
				return nil, err
			}
			match.TeamA = mapper.Convert[*pb_teams.Team, model.TeamModel](teamA)
		}
		if match.IsTeamBKnown {
			teamB, err := s.TeamsClient.GetByFaceitId(ctx, &pb_teams.GetTeamByFaceitIdRequest{FaceitId: match.TeamBFaceitId})
			if err != nil {
				logger.Error(err.Error())
				err := status.Errorf(codes.NotFound, "teamB %s not found", match.TeamBName)
				return nil, err
			}
			match.TeamB = mapper.Convert[*pb_teams.Team, model.TeamModel](teamB)
		}

		m := mapper.Convert[model.MatchModel, *pb.Match](match)
		pbMatches = append(pbMatches, m)
	}

	return &pb.MatchList{Matches: pbMatches}, nil
}

func (s *Server) SetStreamToMatch(ctx context.Context, req *pb.SetStreamRequest) (*pb.Bool, error) {
	err := s.MatchesManager.SetStreamUrl(req.FaceitId, req.StreamChannel)
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			logger.Warning("Not updated, stream already exist")
			err := status.Errorf(codes.AlreadyExists, "stream already exists")
			return &pb.Bool{Res: true}, err
		}
		logger.Error("unable to set stream into match: %s", err.Error())
		err := status.Errorf(codes.Internal, "stream not setted")
		return &pb.Bool{Res: false}, err
	}

	return &pb.Bool{Res: true}, nil
}
