// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: proto/tournaments.proto

package tournaments

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	teams "ibercs/proto/teams"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Tournament struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	FaceitId        string   `protobuf:"bytes,2,opt,name=FaceitId,proto3" json:"FaceitId,omitempty"`
	OrganizerId     string   `protobuf:"bytes,3,opt,name=OrganizerId,proto3" json:"OrganizerId,omitempty"`
	Name            string   `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	RegisterDate    int64    `protobuf:"varint,5,opt,name=RegisterDate,proto3" json:"RegisterDate,omitempty"`
	StartDate       int64    `protobuf:"varint,6,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	JoinPolicy      string   `protobuf:"bytes,7,opt,name=JoinPolicy,proto3" json:"JoinPolicy,omitempty"`
	GeoCountries    []string `protobuf:"bytes,8,rep,name=GeoCountries,proto3" json:"GeoCountries,omitempty"`
	MinLevel        int32    `protobuf:"varint,9,opt,name=MinLevel,proto3" json:"MinLevel,omitempty"`
	MaxLevel        int32    `protobuf:"varint,10,opt,name=MaxLevel,proto3" json:"MaxLevel,omitempty"`
	Status          string   `protobuf:"bytes,11,opt,name=Status,proto3" json:"Status,omitempty"`
	BackgroundImage string   `protobuf:"bytes,12,opt,name=BackgroundImage,proto3" json:"BackgroundImage,omitempty"`
	CoverImage      string   `protobuf:"bytes,13,opt,name=CoverImage,proto3" json:"CoverImage,omitempty"`
	Avatar          string   `protobuf:"bytes,14,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	TeamsId         []string `protobuf:"bytes,15,rep,name=TeamsId,proto3" json:"TeamsId,omitempty"`
}

func (x *Tournament) Reset() {
	*x = Tournament{}
	mi := &file_proto_tournaments_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tournament) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tournament) ProtoMessage() {}

func (x *Tournament) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tournaments_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tournament.ProtoReflect.Descriptor instead.
func (*Tournament) Descriptor() ([]byte, []int) {
	return file_proto_tournaments_proto_rawDescGZIP(), []int{0}
}

func (x *Tournament) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tournament) GetFaceitId() string {
	if x != nil {
		return x.FaceitId
	}
	return ""
}

func (x *Tournament) GetOrganizerId() string {
	if x != nil {
		return x.OrganizerId
	}
	return ""
}

func (x *Tournament) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tournament) GetRegisterDate() int64 {
	if x != nil {
		return x.RegisterDate
	}
	return 0
}

func (x *Tournament) GetStartDate() int64 {
	if x != nil {
		return x.StartDate
	}
	return 0
}

func (x *Tournament) GetJoinPolicy() string {
	if x != nil {
		return x.JoinPolicy
	}
	return ""
}

func (x *Tournament) GetGeoCountries() []string {
	if x != nil {
		return x.GeoCountries
	}
	return nil
}

func (x *Tournament) GetMinLevel() int32 {
	if x != nil {
		return x.MinLevel
	}
	return 0
}

func (x *Tournament) GetMaxLevel() int32 {
	if x != nil {
		return x.MaxLevel
	}
	return 0
}

func (x *Tournament) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Tournament) GetBackgroundImage() string {
	if x != nil {
		return x.BackgroundImage
	}
	return ""
}

func (x *Tournament) GetCoverImage() string {
	if x != nil {
		return x.CoverImage
	}
	return ""
}

func (x *Tournament) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Tournament) GetTeamsId() []string {
	if x != nil {
		return x.TeamsId
	}
	return nil
}

type Organizer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	FaceitId string `protobuf:"bytes,2,opt,name=FaceitId,proto3" json:"FaceitId,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Twitter  string `protobuf:"bytes,4,opt,name=Twitter,proto3" json:"Twitter,omitempty"`
	Twitch   string `protobuf:"bytes,6,opt,name=Twitch,proto3" json:"Twitch,omitempty"`
	Avatar   string `protobuf:"bytes,5,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Type     string `protobuf:"bytes,7,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *Organizer) Reset() {
	*x = Organizer{}
	mi := &file_proto_tournaments_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Organizer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Organizer) ProtoMessage() {}

func (x *Organizer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tournaments_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Organizer.ProtoReflect.Descriptor instead.
func (*Organizer) Descriptor() ([]byte, []int) {
	return file_proto_tournaments_proto_rawDescGZIP(), []int{1}
}

func (x *Organizer) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Organizer) GetFaceitId() string {
	if x != nil {
		return x.FaceitId
	}
	return ""
}

func (x *Organizer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Organizer) GetTwitter() string {
	if x != nil {
		return x.Twitter
	}
	return ""
}

func (x *Organizer) GetTwitch() string {
	if x != nil {
		return x.Twitch
	}
	return ""
}

func (x *Organizer) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Organizer) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_tournaments_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tournaments_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_tournaments_proto_rawDescGZIP(), []int{2}
}

type GetTournamentByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceitId string `protobuf:"bytes,1,opt,name=FaceitId,proto3" json:"FaceitId,omitempty"`
}

func (x *GetTournamentByIdRequest) Reset() {
	*x = GetTournamentByIdRequest{}
	mi := &file_proto_tournaments_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTournamentByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTournamentByIdRequest) ProtoMessage() {}

func (x *GetTournamentByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tournaments_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTournamentByIdRequest.ProtoReflect.Descriptor instead.
func (*GetTournamentByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_tournaments_proto_rawDescGZIP(), []int{3}
}

func (x *GetTournamentByIdRequest) GetFaceitId() string {
	if x != nil {
		return x.FaceitId
	}
	return ""
}

type NewOrganizerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceitId string `protobuf:"bytes,1,opt,name=FaceitId,proto3" json:"FaceitId,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *NewOrganizerRequest) Reset() {
	*x = NewOrganizerRequest{}
	mi := &file_proto_tournaments_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewOrganizerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewOrganizerRequest) ProtoMessage() {}

func (x *NewOrganizerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tournaments_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewOrganizerRequest.ProtoReflect.Descriptor instead.
func (*NewOrganizerRequest) Descriptor() ([]byte, []int) {
	return file_proto_tournaments_proto_rawDescGZIP(), []int{4}
}

func (x *NewOrganizerRequest) GetFaceitId() string {
	if x != nil {
		return x.FaceitId
	}
	return ""
}

func (x *NewOrganizerRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type NewTournamentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceitId string `protobuf:"bytes,1,opt,name=FaceitId,proto3" json:"FaceitId,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *NewTournamentRequest) Reset() {
	*x = NewTournamentRequest{}
	mi := &file_proto_tournaments_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewTournamentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTournamentRequest) ProtoMessage() {}

func (x *NewTournamentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tournaments_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTournamentRequest.ProtoReflect.Descriptor instead.
func (*NewTournamentRequest) Descriptor() ([]byte, []int) {
	return file_proto_tournaments_proto_rawDescGZIP(), []int{5}
}

func (x *NewTournamentRequest) GetFaceitId() string {
	if x != nil {
		return x.FaceitId
	}
	return ""
}

func (x *NewTournamentRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type TournamentList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tournaments []*Tournament `protobuf:"bytes,1,rep,name=Tournaments,proto3" json:"Tournaments,omitempty"`
}

func (x *TournamentList) Reset() {
	*x = TournamentList{}
	mi := &file_proto_tournaments_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TournamentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TournamentList) ProtoMessage() {}

func (x *TournamentList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tournaments_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TournamentList.ProtoReflect.Descriptor instead.
func (*TournamentList) Descriptor() ([]byte, []int) {
	return file_proto_tournaments_proto_rawDescGZIP(), []int{6}
}

func (x *TournamentList) GetTournaments() []*Tournament {
	if x != nil {
		return x.Tournaments
	}
	return nil
}

// ESEA
type Esea struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceitId     string          `protobuf:"bytes,1,opt,name=FaceitId,proto3" json:"FaceitId,omitempty"`
	Name         string          `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Season       int32           `protobuf:"varint,3,opt,name=Season,proto3" json:"Season,omitempty"`
	Playoffs     bool            `protobuf:"varint,4,opt,name=Playoffs,proto3" json:"Playoffs,omitempty"`
	PlayoffsData string          `protobuf:"bytes,5,opt,name=PlayoffsData,proto3" json:"PlayoffsData,omitempty"`
	Divisions    []*EseaDivision `protobuf:"bytes,6,rep,name=Divisions,proto3" json:"Divisions,omitempty"`
}

func (x *Esea) Reset() {
	*x = Esea{}
	mi := &file_proto_tournaments_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Esea) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Esea) ProtoMessage() {}

func (x *Esea) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tournaments_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Esea.ProtoReflect.Descriptor instead.
func (*Esea) Descriptor() ([]byte, []int) {
	return file_proto_tournaments_proto_rawDescGZIP(), []int{7}
}

func (x *Esea) GetFaceitId() string {
	if x != nil {
		return x.FaceitId
	}
	return ""
}

func (x *Esea) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Esea) GetSeason() int32 {
	if x != nil {
		return x.Season
	}
	return 0
}

func (x *Esea) GetPlayoffs() bool {
	if x != nil {
		return x.Playoffs
	}
	return false
}

func (x *Esea) GetPlayoffsData() string {
	if x != nil {
		return x.PlayoffsData
	}
	return ""
}

func (x *Esea) GetDivisions() []*EseaDivision {
	if x != nil {
		return x.Divisions
	}
	return nil
}

type EseaDivision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceitId           string          `protobuf:"bytes,1,opt,name=FaceitId,proto3" json:"FaceitId,omitempty"`
	EseaLeagueFaceitId string          `protobuf:"bytes,2,opt,name=EseaLeagueFaceitId,proto3" json:"EseaLeagueFaceitId,omitempty"`
	Name               string          `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Standings          []*EseaStanding `protobuf:"bytes,4,rep,name=Standings,proto3" json:"Standings,omitempty"`
}

func (x *EseaDivision) Reset() {
	*x = EseaDivision{}
	mi := &file_proto_tournaments_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EseaDivision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EseaDivision) ProtoMessage() {}

func (x *EseaDivision) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tournaments_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EseaDivision.ProtoReflect.Descriptor instead.
func (*EseaDivision) Descriptor() ([]byte, []int) {
	return file_proto_tournaments_proto_rawDescGZIP(), []int{8}
}

func (x *EseaDivision) GetFaceitId() string {
	if x != nil {
		return x.FaceitId
	}
	return ""
}

func (x *EseaDivision) GetEseaLeagueFaceitId() string {
	if x != nil {
		return x.EseaLeagueFaceitId
	}
	return ""
}

func (x *EseaDivision) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EseaDivision) GetStandings() []*EseaStanding {
	if x != nil {
		return x.Standings
	}
	return nil
}

type EseaStanding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDisqualified bool        `protobuf:"varint,1,opt,name=IsDisqualified,proto3" json:"IsDisqualified,omitempty"`
	RankStart      int32       `protobuf:"varint,2,opt,name=RankStart,proto3" json:"RankStart,omitempty"`
	RankEnd        int32       `protobuf:"varint,3,opt,name=RankEnd,proto3" json:"RankEnd,omitempty"`
	Points         int32       `protobuf:"varint,4,opt,name=Points,proto3" json:"Points,omitempty"`
	MatchesPlayed  int32       `protobuf:"varint,5,opt,name=MatchesPlayed,proto3" json:"MatchesPlayed,omitempty"`
	MatchesWon     int32       `protobuf:"varint,6,opt,name=MatchesWon,proto3" json:"MatchesWon,omitempty"`
	MatchesLost    int32       `protobuf:"varint,7,opt,name=MatchesLost,proto3" json:"MatchesLost,omitempty"`
	MatchesTied    int32       `protobuf:"varint,8,opt,name=MatchesTied,proto3" json:"MatchesTied,omitempty"`
	BuchholzScore  int32       `protobuf:"varint,9,opt,name=BuchholzScore,proto3" json:"BuchholzScore,omitempty"`
	Team           *teams.Team `protobuf:"bytes,10,opt,name=Team,proto3" json:"Team,omitempty"`
}

func (x *EseaStanding) Reset() {
	*x = EseaStanding{}
	mi := &file_proto_tournaments_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EseaStanding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EseaStanding) ProtoMessage() {}

func (x *EseaStanding) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tournaments_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EseaStanding.ProtoReflect.Descriptor instead.
func (*EseaStanding) Descriptor() ([]byte, []int) {
	return file_proto_tournaments_proto_rawDescGZIP(), []int{9}
}

func (x *EseaStanding) GetIsDisqualified() bool {
	if x != nil {
		return x.IsDisqualified
	}
	return false
}

func (x *EseaStanding) GetRankStart() int32 {
	if x != nil {
		return x.RankStart
	}
	return 0
}

func (x *EseaStanding) GetRankEnd() int32 {
	if x != nil {
		return x.RankEnd
	}
	return 0
}

func (x *EseaStanding) GetPoints() int32 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *EseaStanding) GetMatchesPlayed() int32 {
	if x != nil {
		return x.MatchesPlayed
	}
	return 0
}

func (x *EseaStanding) GetMatchesWon() int32 {
	if x != nil {
		return x.MatchesWon
	}
	return 0
}

func (x *EseaStanding) GetMatchesLost() int32 {
	if x != nil {
		return x.MatchesLost
	}
	return 0
}

func (x *EseaStanding) GetMatchesTied() int32 {
	if x != nil {
		return x.MatchesTied
	}
	return 0
}

func (x *EseaStanding) GetBuchholzScore() int32 {
	if x != nil {
		return x.BuchholzScore
	}
	return 0
}

func (x *EseaStanding) GetTeam() *teams.Team {
	if x != nil {
		return x.Team
	}
	return nil
}

type GetEseaLeagueBySeasonNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Season int32 `protobuf:"varint,1,opt,name=Season,proto3" json:"Season,omitempty"`
}

func (x *GetEseaLeagueBySeasonNumberRequest) Reset() {
	*x = GetEseaLeagueBySeasonNumberRequest{}
	mi := &file_proto_tournaments_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEseaLeagueBySeasonNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEseaLeagueBySeasonNumberRequest) ProtoMessage() {}

func (x *GetEseaLeagueBySeasonNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tournaments_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEseaLeagueBySeasonNumberRequest.ProtoReflect.Descriptor instead.
func (*GetEseaLeagueBySeasonNumberRequest) Descriptor() ([]byte, []int) {
	return file_proto_tournaments_proto_rawDescGZIP(), []int{10}
}

func (x *GetEseaLeagueBySeasonNumberRequest) GetSeason() int32 {
	if x != nil {
		return x.Season
	}
	return 0
}

var File_proto_tournaments_proto protoreflect.FileDescriptor

var file_proto_tournaments_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65,
	0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x03, 0x0a, 0x0a, 0x54, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x61, 0x63, 0x65,
	0x69, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x61, 0x63, 0x65,
	0x69, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x4a, 0x6f, 0x69, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x4a, 0x6f, 0x69, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x22, 0x0a, 0x0c,
	0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0c, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x69, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x4d, 0x69, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x4d, 0x61, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x4d, 0x61, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x28, 0x0a, 0x0f, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x42, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f,
	0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x43, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x49, 0x64, 0x18, 0x0f, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x49, 0x64, 0x22, 0xa9, 0x01, 0x0a,
	0x09, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x61,
	0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x61,
	0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x77,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x77, 0x69,
	0x74, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x77, 0x69, 0x74, 0x63, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x36, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x13, 0x4e, 0x65, 0x77,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x46, 0x0a, 0x14, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x61, 0x63, 0x65,
	0x69, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x61, 0x63, 0x65,
	0x69, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4b, 0x0a, 0x0e, 0x54, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x54, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x54, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xc7, 0x01, 0x0a, 0x04, 0x45, 0x73, 0x65, 0x61, 0x12, 0x1a,
	0x0a, 0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x6f, 0x66,
	0x66, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x6f, 0x66,
	0x66, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x6f, 0x66, 0x66, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x6f, 0x66,
	0x66, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x09, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x73, 0x65, 0x61, 0x44, 0x69, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0xa7, 0x01, 0x0a, 0x0c, 0x45, 0x73, 0x65, 0x61, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12,
	0x45, 0x73, 0x65, 0x61, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x45, 0x73, 0x65, 0x61, 0x4c, 0x65,
	0x61, 0x67, 0x75, 0x65, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x37, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x45, 0x73, 0x65, 0x61, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x09,
	0x53, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xd7, 0x02, 0x0a, 0x0c, 0x45, 0x73,
	0x65, 0x61, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0e, 0x49, 0x73,
	0x44, 0x69, 0x73, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x49, 0x73, 0x44, 0x69, 0x73, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x52, 0x61, 0x6e, 0x6b, 0x45, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x52, 0x61, 0x6e, 0x6b, 0x45, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x57, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x57, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x4c, 0x6f, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x4c, 0x6f, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x54, 0x69, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x54, 0x69, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d,
	0x42, 0x75, 0x63, 0x68, 0x68, 0x6f, 0x6c, 0x7a, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x42, 0x75, 0x63, 0x68, 0x68, 0x6f, 0x6c, 0x7a, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x54, 0x65, 0x61, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x54,
	0x65, 0x61, 0x6d, 0x22, 0x3c, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x45, 0x73, 0x65, 0x61, 0x4c, 0x65,
	0x61, 0x67, 0x75, 0x65, 0x42, 0x79, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x32, 0xe6, 0x03, 0x0a, 0x11, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x65,
	0x72, 0x12, 0x4b, 0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x21, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x44,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x59, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x12,
	0x25, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x3b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x45, 0x73, 0x65, 0x61, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x74, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x73, 0x65, 0x61, 0x12, 0x5c, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x45, 0x73, 0x65, 0x61, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x79,
	0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x73, 0x65, 0x61, 0x4c, 0x65, 0x61, 0x67,
	0x75, 0x65, 0x42, 0x79, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x73, 0x65, 0x61, 0x42, 0x26, 0x5a, 0x24, 0x69, 0x62,
	0x65, 0x72, 0x63, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3b, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_tournaments_proto_rawDescOnce sync.Once
	file_proto_tournaments_proto_rawDescData = file_proto_tournaments_proto_rawDesc
)

func file_proto_tournaments_proto_rawDescGZIP() []byte {
	file_proto_tournaments_proto_rawDescOnce.Do(func() {
		file_proto_tournaments_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_tournaments_proto_rawDescData)
	})
	return file_proto_tournaments_proto_rawDescData
}

var file_proto_tournaments_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_tournaments_proto_goTypes = []any{
	(*Tournament)(nil),                         // 0: tournaments.Tournament
	(*Organizer)(nil),                          // 1: tournaments.Organizer
	(*Empty)(nil),                              // 2: tournaments.Empty
	(*GetTournamentByIdRequest)(nil),           // 3: tournaments.GetTournamentByIdRequest
	(*NewOrganizerRequest)(nil),                // 4: tournaments.NewOrganizerRequest
	(*NewTournamentRequest)(nil),               // 5: tournaments.NewTournamentRequest
	(*TournamentList)(nil),                     // 6: tournaments.TournamentList
	(*Esea)(nil),                               // 7: tournaments.Esea
	(*EseaDivision)(nil),                       // 8: tournaments.EseaDivision
	(*EseaStanding)(nil),                       // 9: tournaments.EseaStanding
	(*GetEseaLeagueBySeasonNumberRequest)(nil), // 10: tournaments.GetEseaLeagueBySeasonNumberRequest
	(*teams.Team)(nil),                         // 11: teams.Team
}
var file_proto_tournaments_proto_depIdxs = []int32{
	0,  // 0: tournaments.TournamentList.Tournaments:type_name -> tournaments.Tournament
	8,  // 1: tournaments.Esea.Divisions:type_name -> tournaments.EseaDivision
	9,  // 2: tournaments.EseaDivision.Standings:type_name -> tournaments.EseaStanding
	11, // 3: tournaments.EseaStanding.Team:type_name -> teams.Team
	4,  // 4: tournaments.TournamentService.NewOrganizer:input_type -> tournaments.NewOrganizerRequest
	5,  // 5: tournaments.TournamentService.NewTournament:input_type -> tournaments.NewTournamentRequest
	2,  // 6: tournaments.TournamentService.GetAllTournaments:input_type -> tournaments.Empty
	3,  // 7: tournaments.TournamentService.GetTournamentByFaceitId:input_type -> tournaments.GetTournamentByIdRequest
	2,  // 8: tournaments.TournamentService.GetLiveEseaDetails:input_type -> tournaments.Empty
	10, // 9: tournaments.TournamentService.GetEseaDetailsBySeason:input_type -> tournaments.GetEseaLeagueBySeasonNumberRequest
	1,  // 10: tournaments.TournamentService.NewOrganizer:output_type -> tournaments.Organizer
	0,  // 11: tournaments.TournamentService.NewTournament:output_type -> tournaments.Tournament
	6,  // 12: tournaments.TournamentService.GetAllTournaments:output_type -> tournaments.TournamentList
	0,  // 13: tournaments.TournamentService.GetTournamentByFaceitId:output_type -> tournaments.Tournament
	7,  // 14: tournaments.TournamentService.GetLiveEseaDetails:output_type -> tournaments.Esea
	7,  // 15: tournaments.TournamentService.GetEseaDetailsBySeason:output_type -> tournaments.Esea
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_tournaments_proto_init() }
func file_proto_tournaments_proto_init() {
	if File_proto_tournaments_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_tournaments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_tournaments_proto_goTypes,
		DependencyIndexes: file_proto_tournaments_proto_depIdxs,
		MessageInfos:      file_proto_tournaments_proto_msgTypes,
	}.Build()
	File_proto_tournaments_proto = out.File
	file_proto_tournaments_proto_rawDesc = nil
	file_proto_tournaments_proto_goTypes = nil
	file_proto_tournaments_proto_depIdxs = nil
}
