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
	"ibercs/pkg/microservices"
	pb_players "ibercs/proto/players"
	pb "ibercs/proto/users"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	UserManager   *managers.UserManager
	FaceitService *faceit.FaceitClient
	PlayerServer  pb_players.PlayerServiceClient
}

func registerPlayerClient(cfg config.MicroserviceConfig) *pb_players.PlayerServiceClient {
	return microservices.New(cfg.Host_gRPC, cfg.Port_gRPC, pb_players.NewPlayerServiceClient)
}

func New(cfg config.MicroserviceConfig, cfgThirdParty config.ThirdPartyApiTokens) *Server {
	db := database.NewDatabase(cfg.Database)
	userManager := managers.NewUserManager(db.GetDB())
	faceit := faceit.New(cfgThirdParty.FaceitApiToken)

	return &Server{
		UserManager:   userManager,
		FaceitService: faceit,
		PlayerServer:  *registerPlayerClient(cfg),
	}
}

func (s *Server) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.User, error) {
	user, err := s.UserManager.GetByID(int(req.Id))
	if err != nil {
		logger.Error("User %s not found: %s", req.Id, err.Error())
		err := status.Errorf(codes.NotFound, "user not found")
		return nil, err
	}

	res := mapper.Convert[model.UserModel, *pb.User](*user)

	return res, nil
}

func (s *Server) GetUserByFaceitId(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user, err := s.UserManager.GetByFaceitId(req.Id)
	if err != nil {
		logger.Error("User %s not found: %s", req.Id, err.Error())
		err := status.Errorf(codes.NotFound, "user not found")
		return nil, err
	}

	res := mapper.Convert[model.UserModel, *pb.User](*user)

	return res, nil
}

func (s *Server) Create(ctx context.Context, req *pb.NewUserRequest) (*pb.User, error) {
	payload := mapper.Convert[*pb.NewUserRequest, model.UserModel](req)

	user, err := s.UserManager.Create(&payload)
	if err != nil {
		logger.Error("Error creating user: %s", err.Error())
		err := status.Errorf(codes.Internal, "error creating user")
		return nil, err
	}

	res := mapper.Convert[model.UserModel, *pb.User](*user)

	return res, nil
}

func (s *Server) Update(ctx context.Context, req *pb.User) (*pb.User, error) {
	payload := mapper.Convert[*pb.User, model.UserModel](req)

	user, err := s.UserManager.Update(&payload)
	if err != nil {
		logger.Error("Error updating user: %s", err.Error())
		err := status.Errorf(codes.Internal, "error updating user")
		return nil, err
	}

	res := mapper.Convert[model.UserModel, *pb.User](*user)

	return res, nil
}

// Sessions
func (s *Server) GetSessionByUserId(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.SessionResponse, error) {
	session, err := s.UserManager.GetSessionByUserId(int(req.Id))
	if err != nil {
		logger.Error("Session for user %s not found: %s", req.Id, err.Error())
		err := status.Errorf(codes.NotFound, "session not found")
		return nil, err
	}

	res := mapper.Convert[model.UserSessionModel, *pb.SessionResponse](*session)

	return res, nil
}

func (s *Server) CreateSession(ctx context.Context, req *pb.NewSessionRequest) (*pb.SessionResponse, error) {
	session, err := s.UserManager.CreateNewSession(int(req.Id))
	if err != nil {
		logger.Error("Error creating session: %s", err.Error())
		err := status.Errorf(codes.Internal, "error creating session")
		return nil, err
	}

	res := mapper.Convert[model.UserSessionModel, *pb.SessionResponse](*session)

	return res, nil
}

func (s *Server) DeleteSession(ctx context.Context, req *pb.NewSessionRequest) (*pb.Empty, error) {
	err := s.UserManager.DeleteSession(int(req.Id))
	if err != nil {
		logger.Error("Error deleting session: %s", err.Error())
		err := status.Errorf(codes.Internal, "error deleting session")
		return nil, err
	}

	return &pb.Empty{}, nil
}

// Streams
func (s *Server) GetAllStreams(ctx context.Context, req *pb.Empty) (*pb.StreamsResponse, error) {
	users, err := s.UserManager.GetAllStreams()
	if err != nil {
		logger.Error("Error getting all streams: %s", err.Error())
		err := status.Errorf(codes.Internal, "error getting all streams")
		return nil, err
	}

	var res []*pb.StreamResponse
	for _, u := range users {
		res = append(res, &pb.StreamResponse{Stream: u.Twitch, Name: u.Name})
	}

	return &pb.StreamsResponse{Streams: res}, nil
}
