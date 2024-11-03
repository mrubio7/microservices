// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: proto/matches.proto

package matches

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	teams "ibercs/proto/teams"
	tournaments "ibercs/proto/tournaments"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Match struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           int32                   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FaceitId     string                  `protobuf:"bytes,2,opt,name=FaceitId,proto3" json:"FaceitId,omitempty"`
	TeamA        *teams.Team             `protobuf:"bytes,3,opt,name=TeamA,proto3" json:"TeamA,omitempty"`
	TeamAName    string                  `protobuf:"bytes,4,opt,name=TeamAName,proto3" json:"TeamAName,omitempty"`
	IsTeamAKnown bool                    `protobuf:"varint,5,opt,name=IsTeamAKnown,proto3" json:"IsTeamAKnown,omitempty"`
	ScoreTeamA   int32                   `protobuf:"varint,6,opt,name=ScoreTeamA,proto3" json:"ScoreTeamA,omitempty"`
	TeamB        *teams.Team             `protobuf:"bytes,7,opt,name=TeamB,proto3" json:"TeamB,omitempty"`
	TeamBName    string                  `protobuf:"bytes,8,opt,name=TeamBName,proto3" json:"TeamBName,omitempty"`
	IsTeamBKnown bool                    `protobuf:"varint,9,opt,name=IsTeamBKnown,proto3" json:"IsTeamBKnown,omitempty"`
	ScoreTeamB   int32                   `protobuf:"varint,10,opt,name=ScoreTeamB,proto3" json:"ScoreTeamB,omitempty"`
	BestOf       int32                   `protobuf:"varint,11,opt,name=BestOf,proto3" json:"BestOf,omitempty"`
	Timestamp    int64                   `protobuf:"varint,12,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Streams      []string                `protobuf:"bytes,13,rep,name=Streams,proto3" json:"Streams,omitempty"`
	Tournament   *tournaments.Tournament `protobuf:"bytes,14,opt,name=Tournament,proto3" json:"Tournament,omitempty"`
	Map          []string                `protobuf:"bytes,15,rep,name=Map,proto3" json:"Map,omitempty"`
}

func (x *Match) Reset() {
	*x = Match{}
	mi := &file_proto_matches_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Match) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Match) ProtoMessage() {}

func (x *Match) ProtoReflect() protoreflect.Message {
	mi := &file_proto_matches_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Match.ProtoReflect.Descriptor instead.
func (*Match) Descriptor() ([]byte, []int) {
	return file_proto_matches_proto_rawDescGZIP(), []int{0}
}

func (x *Match) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Match) GetFaceitId() string {
	if x != nil {
		return x.FaceitId
	}
	return ""
}

func (x *Match) GetTeamA() *teams.Team {
	if x != nil {
		return x.TeamA
	}
	return nil
}

func (x *Match) GetTeamAName() string {
	if x != nil {
		return x.TeamAName
	}
	return ""
}

func (x *Match) GetIsTeamAKnown() bool {
	if x != nil {
		return x.IsTeamAKnown
	}
	return false
}

func (x *Match) GetScoreTeamA() int32 {
	if x != nil {
		return x.ScoreTeamA
	}
	return 0
}

func (x *Match) GetTeamB() *teams.Team {
	if x != nil {
		return x.TeamB
	}
	return nil
}

func (x *Match) GetTeamBName() string {
	if x != nil {
		return x.TeamBName
	}
	return ""
}

func (x *Match) GetIsTeamBKnown() bool {
	if x != nil {
		return x.IsTeamBKnown
	}
	return false
}

func (x *Match) GetScoreTeamB() int32 {
	if x != nil {
		return x.ScoreTeamB
	}
	return 0
}

func (x *Match) GetBestOf() int32 {
	if x != nil {
		return x.BestOf
	}
	return 0
}

func (x *Match) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Match) GetStreams() []string {
	if x != nil {
		return x.Streams
	}
	return nil
}

func (x *Match) GetTournament() *tournaments.Tournament {
	if x != nil {
		return x.Tournament
	}
	return nil
}

func (x *Match) GetMap() []string {
	if x != nil {
		return x.Map
	}
	return nil
}

type MatchList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matches []*Match `protobuf:"bytes,1,rep,name=Matches,proto3" json:"Matches,omitempty"`
}

func (x *MatchList) Reset() {
	*x = MatchList{}
	mi := &file_proto_matches_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchList) ProtoMessage() {}

func (x *MatchList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_matches_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchList.ProtoReflect.Descriptor instead.
func (*MatchList) Descriptor() ([]byte, []int) {
	return file_proto_matches_proto_rawDescGZIP(), []int{1}
}

func (x *MatchList) GetMatches() []*Match {
	if x != nil {
		return x.Matches
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_matches_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_matches_proto_msgTypes[2]
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
	return file_proto_matches_proto_rawDescGZIP(), []int{2}
}

type GetMatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceitId string `protobuf:"bytes,1,opt,name=FaceitId,proto3" json:"FaceitId,omitempty"`
}

func (x *GetMatchRequest) Reset() {
	*x = GetMatchRequest{}
	mi := &file_proto_matches_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMatchRequest) ProtoMessage() {}

func (x *GetMatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_matches_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMatchRequest.ProtoReflect.Descriptor instead.
func (*GetMatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_matches_proto_rawDescGZIP(), []int{3}
}

func (x *GetMatchRequest) GetFaceitId() string {
	if x != nil {
		return x.FaceitId
	}
	return ""
}

type GetUpcomingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Days int32 `protobuf:"varint,1,opt,name=Days,proto3" json:"Days,omitempty"`
}

func (x *GetUpcomingRequest) Reset() {
	*x = GetUpcomingRequest{}
	mi := &file_proto_matches_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUpcomingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpcomingRequest) ProtoMessage() {}

func (x *GetUpcomingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_matches_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpcomingRequest.ProtoReflect.Descriptor instead.
func (*GetUpcomingRequest) Descriptor() ([]byte, []int) {
	return file_proto_matches_proto_rawDescGZIP(), []int{4}
}

func (x *GetUpcomingRequest) GetDays() int32 {
	if x != nil {
		return x.Days
	}
	return 0
}

var File_proto_matches_proto protoreflect.FileDescriptor

var file_proto_matches_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x1a, 0x11,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x03, 0x0a, 0x05, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x05, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x05, 0x54, 0x65,
	0x61, 0x6d, 0x41, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x73, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x4b, 0x6e, 0x6f, 0x77,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x49, 0x73, 0x54, 0x65, 0x61, 0x6d, 0x41,
	0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x41, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x54, 0x65, 0x61, 0x6d, 0x41, 0x12, 0x21, 0x0a, 0x05, 0x54, 0x65, 0x61, 0x6d, 0x42, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x05, 0x54, 0x65, 0x61, 0x6d, 0x42, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65, 0x61, 0x6d,
	0x42, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x65, 0x61,
	0x6d, 0x42, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x73, 0x54, 0x65, 0x61, 0x6d,
	0x42, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x49, 0x73,
	0x54, 0x65, 0x61, 0x6d, 0x42, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x42, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x42, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x65,
	0x73, 0x74, 0x4f, 0x66, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x42, 0x65, 0x73, 0x74,
	0x4f, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x54, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x54, 0x6f, 0x75,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x61, 0x70, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x03, 0x4d, 0x61, 0x70, 0x22, 0x35, 0x0a, 0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x61, 0x63, 0x65,
	0x69, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x61, 0x63, 0x65,
	0x69, 0x74, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x70, 0x63, 0x6f, 0x6d,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61,
	0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x44, 0x61, 0x79, 0x73, 0x32, 0xcc,
	0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x33, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x12, 0x0e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x12, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x70, 0x63,
	0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x63, 0x6f, 0x6d, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x79, 0x46, 0x61, 0x63, 0x65, 0x69,
	0x74, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x1e, 0x5a,
	0x1c, 0x69, 0x62, 0x65, 0x72, 0x63, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x3b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_matches_proto_rawDescOnce sync.Once
	file_proto_matches_proto_rawDescData = file_proto_matches_proto_rawDesc
)

func file_proto_matches_proto_rawDescGZIP() []byte {
	file_proto_matches_proto_rawDescOnce.Do(func() {
		file_proto_matches_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_matches_proto_rawDescData)
	})
	return file_proto_matches_proto_rawDescData
}

var file_proto_matches_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_matches_proto_goTypes = []any{
	(*Match)(nil),                  // 0: matches.Match
	(*MatchList)(nil),              // 1: matches.MatchList
	(*Empty)(nil),                  // 2: matches.Empty
	(*GetMatchRequest)(nil),        // 3: matches.GetMatchRequest
	(*GetUpcomingRequest)(nil),     // 4: matches.GetUpcomingRequest
	(*teams.Team)(nil),             // 5: teams.Team
	(*tournaments.Tournament)(nil), // 6: tournaments.Tournament
}
var file_proto_matches_proto_depIdxs = []int32{
	5, // 0: matches.Match.TeamA:type_name -> teams.Team
	5, // 1: matches.Match.TeamB:type_name -> teams.Team
	6, // 2: matches.Match.Tournament:type_name -> tournaments.Tournament
	0, // 3: matches.MatchList.Matches:type_name -> matches.Match
	2, // 4: matches.MatchesService.GetAllMatches:input_type -> matches.Empty
	4, // 5: matches.MatchesService.GetUpcomingMatches:input_type -> matches.GetUpcomingRequest
	3, // 6: matches.MatchesService.GetMatchByFaceitId:input_type -> matches.GetMatchRequest
	1, // 7: matches.MatchesService.GetAllMatches:output_type -> matches.MatchList
	1, // 8: matches.MatchesService.GetUpcomingMatches:output_type -> matches.MatchList
	0, // 9: matches.MatchesService.GetMatchByFaceitId:output_type -> matches.Match
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_matches_proto_init() }
func file_proto_matches_proto_init() {
	if File_proto_matches_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_matches_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_matches_proto_goTypes,
		DependencyIndexes: file_proto_matches_proto_depIdxs,
		MessageInfos:      file_proto_matches_proto_msgTypes,
	}.Build()
	File_proto_matches_proto = out.File
	file_proto_matches_proto_rawDesc = nil
	file_proto_matches_proto_goTypes = nil
	file_proto_matches_proto_depIdxs = nil
}
