package tournaments_mapper

import (
	"ibercs/internal/model"
	pb "ibercs/proto/tournaments"
	"time"
)

type TournamentMapper struct{}

func (TournamentMapper) Proto(entity model.TournamentModel, _ ...interface{}) *pb.Tournament {
	return &pb.Tournament{
		Id:              entity.Id,
		Name:            entity.Name,
		FaceitId:        entity.FaceitId,
		OrganizerId:     entity.OrganizerId,
		RegisterDate:    entity.RegisterDate.Unix(),
		StartDate:       entity.StartDate.Unix(),
		JoinPolicy:      entity.JoinPolicy,
		GeoCountries:    entity.GeoCountries,
		MinLevel:        int32(entity.MinLevel),
		MaxLevel:        int32(entity.MaxLevel),
		Status:          entity.Status,
		BackgroundImage: entity.BackgroundImage,
		CoverImage:      entity.CoverImage,
		Avatar:          entity.Avatar,
		TeamsId:         entity.TeamsId,
	}
}

func (TournamentMapper) Model(proto *pb.Tournament, _ ...interface{}) model.TournamentModel {
	return model.TournamentModel{
		Id:              proto.Id,
		Name:            proto.Name,
		FaceitId:        proto.FaceitId,
		OrganizerId:     proto.OrganizerId,
		RegisterDate:    time.Unix(proto.RegisterDate, 0),
		StartDate:       time.Unix(proto.StartDate, 0),
		JoinPolicy:      proto.JoinPolicy,
		GeoCountries:    proto.GeoCountries,
		MinLevel:        int(proto.MinLevel),
		MaxLevel:        int(proto.MaxLevel),
		Status:          proto.Status,
		BackgroundImage: proto.BackgroundImage,
		CoverImage:      proto.CoverImage,
		Avatar:          proto.Avatar,
		TeamsId:         proto.TeamsId,
	}
}

type OrganizerMapper struct{}

func (OrganizerMapper) Proto(entity model.OrganizerModel, _ ...interface{}) *pb.Organizer {
	return &pb.Organizer{
		Id:       entity.Id,
		Name:     entity.Name,
		FaceitId: entity.FaceitId,
		Twitter:  entity.Twitter,
		Twitch:   entity.Twitch,
		Avatar:   entity.Avatar,
		Type:     entity.Type,
	}
}

func (OrganizerMapper) Model(proto *pb.Organizer, _ ...interface{}) model.OrganizerModel {
	return model.OrganizerModel{
		Id:       proto.Id,
		Name:     proto.Name,
		FaceitId: proto.FaceitId,
		Twitter:  proto.Twitter,
		Twitch:   proto.Twitch,
		Avatar:   proto.Avatar,
		Type:     proto.Type,
	}
}
