package microservice_players_test

import (
	"context"
	"testing"

	microservice_players "ibercs/cmd/microservices/teams/server"
	"ibercs/pkg/logger"
	pb "ibercs/proto/teams"
)

// func TestNewTeam(t *testing.T) {
// 	server := microservice_players.New()
// 	req := &pb.NewTeamRequest{FaceitId: "sample_faceit_id"}

// 	resp, err := server.NewTeam(context.Background(), req)
// 	if err != nil {
// 		t.Fatalf("NewTeam failed: %v", err)
// 	}
// 	if resp.FaceitId != req.FaceitId {
// 		t.Errorf("Expected FaceitId %s, got %s", req.FaceitId, resp.FaceitId)
// 	}
// }

func TestGetTeamById(t *testing.T) {
	logger.Initialize()
	server := microservice_players.New()
	req := &pb.NewTeamRequest{FaceitId: "96b9177a-9af0-47af-bbac-0e8e1ad2f723"}

	resp, err := server.GetTeamById(context.Background(), req)
	if err != nil {
		t.Fatalf("GetTeamById failed: %v", err)
	}
	if resp.FaceitId != req.FaceitId {
		t.Errorf("Expected FaceitId %s, got %s", req.FaceitId, resp.FaceitId)
	}
}

func TestGetTeamByNickname(t *testing.T) {
	logger.Initialize()
	server := microservice_players.New()
	req := &pb.NewTeamRequest{FaceitId: "TGDpro"}

	resp, err := server.GetTeamByNickname(context.Background(), req)
	if err != nil {
		t.Fatalf("GetTeamByNickname failed: %v", err)
	}
	if resp.Nickname != req.FaceitId {
		t.Errorf("Expected Nickname %s, got %s", req.FaceitId, resp.Nickname)
	}
}

func TestGetTeams(t *testing.T) {
	logger.Initialize()
	server := microservice_players.New()
	req := &pb.GetTeamsRequest{Active: true}

	resp, err := server.GetTeams(context.Background(), req)
	if err != nil {
		t.Fatalf("GetTeams failed: %v", err)
	}
	if len(resp.Teams) == 0 {
		t.Error("Expected non-empty team list")
	}
}

func TestFindTeamByPlayerId(t *testing.T) {
	logger.Initialize()
	server := microservice_players.New()
	req := &pb.NewTeamRequest{FaceitId: "50344702-7dfd-4e81-836e-bc77448b8316"}

	resp, err := server.FindTeamByPlayerId(context.Background(), req)
	if err != nil {
		t.Fatalf("FindTeamByPlayerId failed: %v", err)
	}
	if len(resp.Teams) == 0 {
		t.Error("Expected to find at least one team")
	}
}

func TestGetTeamFromFaceit(t *testing.T) {
	logger.Initialize()
	server := microservice_players.New()
	req := &pb.NewTeamRequest{FaceitId: "01e682b8-ccac-42df-89ab-0cfc44af48c6"}

	resp, err := server.GetTeamFromFaceit(context.Background(), req)
	if err != nil {
		t.Fatalf("GetTeamFromFaceit failed: %v", err)
	}
	if resp.FaceitId != req.FaceitId {
		t.Errorf("Expected FaceitId %s, got %s", req.FaceitId, resp.FaceitId)
	}
}
