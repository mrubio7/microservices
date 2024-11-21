package microservice_players

import (
	"context"
	"ibercs/internal/model"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/managers"
	"ibercs/pkg/mapper"
	pb "ibercs/proto/tournaments"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedTournamentServiceServer
	TournamentManager *managers.TournamentManager
	EseaManager       *managers.EseaManager
	FaceitService     *faceit.FaceitClient
}

func New(cfg config.MicroserviceConfig, cfgThirdParty config.ThirdPartyApiTokens) *Server {
	db := database.NewDatabase(cfg.Database)
	tournamentManager := managers.NewTournamentManager(db.GetDB())
	eseaManager := managers.NewEseaManager(db.GetDB())
	faceit := faceit.New(cfgThirdParty.FaceitApiToken)

	return &Server{
		TournamentManager: tournamentManager,
		EseaManager:       eseaManager,
		FaceitService:     faceit,
	}
}

func (s *Server) CreateOrganizer(ctx context.Context, req *pb.NewOrganizerRequest) (*pb.Organizer, error) {
	if o, err := s.TournamentManager.GetOrganizerByFaceitId(req.FaceitId); o != nil {
		if err != nil {
			logger.Error("Error getting organizer by faceit id: %v", err)
			err := status.Errorf(codes.Internal, "Error getting organizer by faceit id: %v", err)
			return nil, err
		}
		logger.Warning("Organizer %s already exist", o.Name)
		err := status.Errorf(codes.AlreadyExists, "organizer %s already exist", o.Name)
		return nil, err
	}

	organizer := s.FaceitService.GetOrganizerById(req.FaceitId)
	if organizer == nil {
		err := status.Errorf(codes.NotFound, "organizer with FaceitID %s doesn't exist", req.FaceitId)
		logger.Error(err.Error())
		return nil, err
	}

	o, err := s.TournamentManager.CreateOrganizer(organizer)
	if err != nil {
		logger.Error("Error creating organizer: %v", err)
		err := status.Errorf(codes.Internal, "Error creating organizer: %v", err)
		return nil, err
	}

	res := mapper.Convert[model.OrganizerModel, *pb.Organizer](*o)

	return res, nil
}

func (s *Server) CreateTournament(ctx context.Context, req *pb.NewTournamentRequest) (*pb.Tournament, error) {
	if t, err := s.TournamentManager.GetTournamentByFaceitId(req.FaceitId); t != nil {
		if err != nil {
			logger.Error("Error getting tournament by faceit id: %v", err)
			err := status.Errorf(codes.Internal, "Error getting tournament by faceit id: %v", err)
			return nil, err
		}
		logger.Warning("Tournament %s already exist", t.Name)
		err := status.Errorf(codes.AlreadyExists, "tournament %s already exist", t.Name)
		return nil, err
	}

	tournament := s.FaceitService.GetChampionshipById(req.FaceitId)
	if tournament == nil {
		err := status.Errorf(codes.NotFound, "tournament with FaceitID %s doesn't exist", req.FaceitId)
		logger.Error(err.Error())
		return nil, err
	}

	t, err := s.TournamentManager.CreateTournament(tournament)
	if err != nil {
		logger.Error("Error creating tournament: %v", err)
		err := status.Errorf(codes.Internal, "Error creating tournament: %v", err)
		return nil, err
	}

	res := mapper.Convert[model.TournamentModel, *pb.Tournament](*t)

	return res, nil
}

func (s *Server) GetTournamentByFaceitId(ctx context.Context, req *pb.GetTournamentByIdRequest) (*pb.Tournament, error) {
	t, err := s.TournamentManager.GetTournamentByFaceitId(req.FaceitId)
	if err != nil {
		logger.Error("Error getting tournament by faceit id: %v", err)
		err := status.Errorf(codes.Internal, "Error getting tournament by faceit id: %v", err)
		return nil, err
	}

	if t == nil {
		err := status.Errorf(codes.NotFound, "tournament with FaceitID %s doesn't exist", req.FaceitId)
		logger.Error(err.Error())
		return nil, err
	}

	res := mapper.Convert[model.TournamentModel, *pb.Tournament](*t)

	return res, nil
}

func (s *Server) GetAllTournaments(ctx context.Context, _ *pb.Empty) (*pb.TournamentList, error) {
	tournaments, err := s.TournamentManager.GetAllTournaments()
	if err != nil {
		logger.Error("Error getting all tournaments: %v", err)
		err := status.Errorf(codes.Internal, "Error getting all tournaments: %v", err)
		return nil, err
	}

	tournamentsList := make([]*pb.Tournament, 0)
	for _, t := range tournaments {
		tournamentsList = append(tournamentsList, mapper.Convert[model.TournamentModel, *pb.Tournament](t))
	}

	return &pb.TournamentList{Tournaments: tournamentsList}, nil
}

// Esea
func (s *Server) GetLiveEseaDetails(ctx context.Context, _ *pb.Empty) (*pb.Esea, error) {
	eseaDetails, err := s.EseaManager.GetEseaLeagueLive()
	if err != nil {
		logger.Error("Error getting live esea details: %v", err)
		err := status.Errorf(codes.Internal, "Error getting live esea details: %v", err)
		return nil, err
	}

	res := mapper.Convert[model.EseaLeagueModel, *pb.Esea](*eseaDetails)

	return res, nil
}

func (s *Server) GetEseaLeagueBySeasonNumber(ctx context.Context, req *pb.GetEseaLeagueBySeasonNumberRequest) (*pb.Esea, error) {
	eseaDetails, err := s.EseaManager.GetEseaLeagueBySeasonNumber(req.Season)
	if err != nil {
		logger.Error("Error getting esea league by season number: %v", err)
		err := status.Errorf(codes.Internal, "Error getting esea league by season number: %v", err)
		return nil, err
	}

	res := mapper.Convert[model.EseaLeagueModel, *pb.Esea](*eseaDetails)

	return res, nil
}
