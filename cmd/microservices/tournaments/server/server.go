package microservice_tournaments

import (
	"context"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/service"
	pb "ibercs/proto/tournaments"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedTournamentServiceServer
	TournamentsService *service.Tournaments
	FaceitService      *faceit.FaceitClient
}

func New() *Server {
	cfg, err := config.Load()
	if err != nil {
		logger.Error("Unable to create grpc tournament server")
		return nil
	}
	db := database.New(cfg.Database)
	tournamentService := service.NewTournamentsService(db)
	faceit := faceit.New(cfg.FaceitApiToken)

	return &Server{
		TournamentsService: tournamentService,
		FaceitService:      faceit,
	}
}

func (s *Server) NewOrganizer(ctx context.Context, organizerReq *pb.NewOrganizerRequest) (*pb.Organizer, error) {
	if o := s.TournamentsService.GetOrganizer(organizerReq.FaceitId); o != nil {
		err := status.Errorf(codes.AlreadyExists, "organizer with FaceitID %s already exist", organizerReq.FaceitId)
		logger.Warning(err.Error())
		return nil, err
	}

	organizer := s.FaceitService.GetOrganizerById(organizerReq.FaceitId)
	if organizer == nil {
		err := status.Errorf(codes.NotFound, "organizer with FaceitID %s doesn't exist", organizerReq.FaceitId)
		logger.Warning(err.Error())
		return nil, err
	}

	organizer.Type = organizerReq.Type

	res := s.TournamentsService.NewOrganizer(organizer)
	if res == nil {
		err := status.Errorf(codes.Internal, "unable to create organizer with FaceitID %s", organizerReq.FaceitId)
		logger.Error(err.Error())
		return nil, err
	}

	return &pb.Organizer{
		Id:       res.ID,
		FaceitId: res.FaceitId,
		Name:     res.Name,
	}, nil
}

func (s *Server) NewTournament(ctx context.Context, tournamentReq *pb.NewTournamentRequest) (*pb.Tournament, error) {
	if o := s.TournamentsService.GetTournament(tournamentReq.FaceitId); o != nil {
		err := status.Errorf(codes.AlreadyExists, "tournament with FaceitID %s already exist", tournamentReq.FaceitId)
		logger.Warning(err.Error())
		return nil, err
	}

	t := s.FaceitService.GetChampionshipById(tournamentReq.FaceitId)
	if t == nil {
		err := status.Errorf(codes.NotFound, "tournament with FaceitID %s doesn't exist", tournamentReq.FaceitId)
		logger.Warning(err.Error())
		return nil, err
	}

	res := s.TournamentsService.NewTournament(t)
	if res == nil {
		err := status.Errorf(codes.AlreadyExists, "unable to create tournament with FaceitID %s", tournamentReq.FaceitId)
		logger.Warning(err.Error())
		return nil, err
	}

	tournament := &pb.Tournament{
		Id:              int32(t.ID),
		FaceitId:        t.FaceitId,
		OrganizerId:     t.OrganizerId,
		Name:            t.Name,
		RegisterDate:    t.RegisterDate.Unix(),
		StartDate:       t.RegisterDate.Unix(),
		CurrentTeams:    int32(t.CurrentTeams),
		Slots:           int32(t.Slots),
		JoinPolicy:      t.JoinPolicy,
		GeoCountries:    t.GeoCountries,
		MinLevel:        int32(t.MinLevel),
		MaxLavel:        int32(t.MaxLevel),
		Status:          t.Status,
		BackgroundImage: t.BackgroundImage,
		CoverImage:      t.CoverImage,
		Avatar:          t.Avatar,
	}

	return tournament, nil
}

func (s *Server) GetAllTorunaments(ctx context.Context, _ *pb.Empty) (*pb.TournamentList, error) {
	tournaments := s.TournamentsService.GetAllTournaments()

	res := make([]*pb.Tournament, len(tournaments))
	for i, t := range tournaments {
		res[i] = &pb.Tournament{
			Id:              int32(t.ID),
			FaceitId:        t.FaceitId,
			OrganizerId:     t.OrganizerId,
			Name:            t.Name,
			RegisterDate:    t.RegisterDate.Unix(),
			StartDate:       t.RegisterDate.Unix(),
			CurrentTeams:    int32(t.CurrentTeams),
			Slots:           int32(t.Slots),
			JoinPolicy:      t.JoinPolicy,
			GeoCountries:    t.GeoCountries,
			MinLevel:        int32(t.MinLevel),
			MaxLavel:        int32(t.MaxLevel),
			Status:          t.Status,
			BackgroundImage: t.BackgroundImage,
			CoverImage:      t.CoverImage,
			Avatar:          t.Avatar,
		}
	}

	return &pb.TournamentList{Tournaments: res}, nil
}

func (s *Server) GetTournamentById(ctx context.Context, req *pb.GetTournamentByIdRequest) (*pb.Tournament, error) {
	t := s.TournamentsService.GetTournament(req.FaceitId)
	if t == nil {
		err := status.Errorf(codes.NotFound, "tournament with FaceitID %s not found", req.FaceitId)
		return nil, err
	}

	err := s.TournamentsService.UpdateTournament(t)
	if err != nil {
		return nil, err
	}

	res := &pb.Tournament{
		Id:              int32(t.ID),
		FaceitId:        t.FaceitId,
		OrganizerId:     t.OrganizerId,
		Name:            t.Name,
		RegisterDate:    t.RegisterDate.Unix(),
		StartDate:       t.RegisterDate.Unix(),
		CurrentTeams:    int32(t.CurrentTeams),
		Slots:           int32(t.Slots),
		JoinPolicy:      t.JoinPolicy,
		GeoCountries:    t.GeoCountries,
		MinLevel:        int32(t.MinLevel),
		MaxLavel:        int32(t.MaxLevel),
		Status:          t.Status,
		BackgroundImage: t.BackgroundImage,
		CoverImage:      t.CoverImage,
		Avatar:          t.Avatar,
	}

	return res, nil
}
