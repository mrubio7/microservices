package microservice_players

import (
	"context"
	"errors"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/logger"
	"ibercs/pkg/service"
	pb "ibercs/proto/players"
	"net/http"
	"time"
)

type Server struct {
	pb.UnimplementedPlayerServiceServer
	PlayersService *service.Players
}

func New() *Server {
	cfg, err := config.Load()
	if err != nil {
		logger.Error("Unable to create grpc player server")
		return nil
	}
	db := database.New(cfg.Database)
	playerService := service.NewPlayersService(db)

	go func() {
		http.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(time.Now().String()))
		})
		if err := http.ListenAndServe(":8080", nil); err != nil {
			logger.Error(err.Error())
		}
	}()

	return &Server{
		PlayersService: playerService,
	}
}

func (s *Server) GetPlayers(context.Context, *pb.Empty) (*pb.PlayerList, error) {
	// Obtener los jugadores desde la base de datos
	playerModels := s.PlayersService.GetPlayers()
	if playerModels == nil {
		return nil, errors.New("unable to get any player")
	}

	// Convertir los modelos de la base de datos a los modelos de protobuf
	var players []*pb.Player
	for _, p := range playerModels {
		player := &pb.Player{
			Id:       p.ID,
			Nickname: p.Nickname,
			FaceitId: p.FaceitId,
			SteamId:  p.SteamId,
			Stats: &pb.PlayerStats{
				PlayerId:               p.Stats.ID,
				KdRatio:                p.Stats.KdRatio,
				KrRatio:                p.Stats.KrRatio,
				KillsAverage:           p.Stats.KillsAverage,
				DeathsAverage:          p.Stats.DeathsAverage,
				HeadshotPercentAverage: p.Stats.HeadshotPercentAverage,
				MVPAverage:             p.Stats.MVPAverage,
				AssistAverage:          p.Stats.AssistAverage,
				Elo:                    p.Stats.Elo,
			},
		}
		players = append(players, player)
	}

	// Retornar el resultado en el formato gRPC
	return &pb.PlayerList{Players: players}, nil
}

func (s *Server) NewPlayer(context.Context, *pb.NewPlayerRequest) (*pb.PlayerResponse, error) {
	return nil, nil
}
