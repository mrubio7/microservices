// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: proto/players.proto

package players

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32        `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Nickname string       `protobuf:"bytes,2,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	FaceitId string       `protobuf:"bytes,3,opt,name=FaceitId,proto3" json:"FaceitId,omitempty"`
	SteamId  string       `protobuf:"bytes,4,opt,name=SteamId,proto3" json:"SteamId,omitempty"`
	Avatar   string       `protobuf:"bytes,5,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Stats    *PlayerStats `protobuf:"bytes,6,opt,name=Stats,proto3" json:"Stats,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	mi := &file_proto_players_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{0}
}

func (x *Player) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Player) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Player) GetFaceitId() string {
	if x != nil {
		return x.FaceitId
	}
	return ""
}

func (x *Player) GetSteamId() string {
	if x != nil {
		return x.SteamId
	}
	return ""
}

func (x *Player) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Player) GetStats() *PlayerStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type PlayerStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId               int32   `protobuf:"varint,1,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	KdRatio                float32 `protobuf:"fixed32,2,opt,name=KdRatio,proto3" json:"KdRatio,omitempty"`
	KrRatio                float32 `protobuf:"fixed32,3,opt,name=KrRatio,proto3" json:"KrRatio,omitempty"`
	KillsAverage           float32 `protobuf:"fixed32,4,opt,name=KillsAverage,proto3" json:"KillsAverage,omitempty"`
	DeathsAverage          float32 `protobuf:"fixed32,5,opt,name=DeathsAverage,proto3" json:"DeathsAverage,omitempty"`
	HeadshotPercentAverage float32 `protobuf:"fixed32,6,opt,name=HeadshotPercentAverage,proto3" json:"HeadshotPercentAverage,omitempty"`
	MVPAverage             float32 `protobuf:"fixed32,7,opt,name=MVPAverage,proto3" json:"MVPAverage,omitempty"`
	AssistAverage          float32 `protobuf:"fixed32,8,opt,name=AssistAverage,proto3" json:"AssistAverage,omitempty"`
	Elo                    int32   `protobuf:"varint,9,opt,name=Elo,proto3" json:"Elo,omitempty"`
}

func (x *PlayerStats) Reset() {
	*x = PlayerStats{}
	mi := &file_proto_players_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerStats) ProtoMessage() {}

func (x *PlayerStats) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerStats.ProtoReflect.Descriptor instead.
func (*PlayerStats) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerStats) GetPlayerId() int32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *PlayerStats) GetKdRatio() float32 {
	if x != nil {
		return x.KdRatio
	}
	return 0
}

func (x *PlayerStats) GetKrRatio() float32 {
	if x != nil {
		return x.KrRatio
	}
	return 0
}

func (x *PlayerStats) GetKillsAverage() float32 {
	if x != nil {
		return x.KillsAverage
	}
	return 0
}

func (x *PlayerStats) GetDeathsAverage() float32 {
	if x != nil {
		return x.DeathsAverage
	}
	return 0
}

func (x *PlayerStats) GetHeadshotPercentAverage() float32 {
	if x != nil {
		return x.HeadshotPercentAverage
	}
	return 0
}

func (x *PlayerStats) GetMVPAverage() float32 {
	if x != nil {
		return x.MVPAverage
	}
	return 0
}

func (x *PlayerStats) GetAssistAverage() float32 {
	if x != nil {
		return x.AssistAverage
	}
	return 0
}

func (x *PlayerStats) GetElo() int32 {
	if x != nil {
		return x.Elo
	}
	return 0
}

type PlayerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *PlayerList) Reset() {
	*x = PlayerList{}
	mi := &file_proto_players_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerList) ProtoMessage() {}

func (x *PlayerList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerList.ProtoReflect.Descriptor instead.
func (*PlayerList) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerList) GetPlayers() []*Player {
	if x != nil {
		return x.Players
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
	mi := &file_proto_players_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[3]
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
	return file_proto_players_proto_rawDescGZIP(), []int{3}
}

type GetPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceitId []string `protobuf:"bytes,1,rep,name=FaceitId,proto3" json:"FaceitId,omitempty"`
}

func (x *GetPlayerRequest) Reset() {
	*x = GetPlayerRequest{}
	mi := &file_proto_players_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerRequest) ProtoMessage() {}

func (x *GetPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerRequest) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{4}
}

func (x *GetPlayerRequest) GetFaceitId() []string {
	if x != nil {
		return x.FaceitId
	}
	return nil
}

type NewPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname string       `protobuf:"bytes,1,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	FaceitId string       `protobuf:"bytes,2,opt,name=FaceitId,proto3" json:"FaceitId,omitempty"`
	SteamId  string       `protobuf:"bytes,3,opt,name=SteamId,proto3" json:"SteamId,omitempty"`
	Stats    *PlayerStats `protobuf:"bytes,4,opt,name=Stats,proto3" json:"Stats,omitempty"`
}

func (x *NewPlayerRequest) Reset() {
	*x = NewPlayerRequest{}
	mi := &file_proto_players_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPlayerRequest) ProtoMessage() {}

func (x *NewPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPlayerRequest.ProtoReflect.Descriptor instead.
func (*NewPlayerRequest) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{5}
}

func (x *NewPlayerRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *NewPlayerRequest) GetFaceitId() string {
	if x != nil {
		return x.FaceitId
	}
	return ""
}

func (x *NewPlayerRequest) GetSteamId() string {
	if x != nil {
		return x.SteamId
	}
	return ""
}

func (x *NewPlayerRequest) GetStats() *PlayerStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type ProminentPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Score    float32 `protobuf:"fixed32,2,opt,name=Score,proto3" json:"Score,omitempty"`
	Nickname string  `protobuf:"bytes,3,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	FaceitId string  `protobuf:"bytes,4,opt,name=Faceit_id,json=FaceitId,proto3" json:"Faceit_id,omitempty"`
	SteamId  string  `protobuf:"bytes,5,opt,name=Steam_id,json=SteamId,proto3" json:"Steam_id,omitempty"`
	Avatar   string  `protobuf:"bytes,6,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
}

func (x *ProminentPlayer) Reset() {
	*x = ProminentPlayer{}
	mi := &file_proto_players_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProminentPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProminentPlayer) ProtoMessage() {}

func (x *ProminentPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProminentPlayer.ProtoReflect.Descriptor instead.
func (*ProminentPlayer) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{6}
}

func (x *ProminentPlayer) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProminentPlayer) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *ProminentPlayer) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *ProminentPlayer) GetFaceitId() string {
	if x != nil {
		return x.FaceitId
	}
	return ""
}

func (x *ProminentPlayer) GetSteamId() string {
	if x != nil {
		return x.SteamId
	}
	return ""
}

func (x *ProminentPlayer) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type ProminentPlayerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*ProminentPlayer `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	Week    int32              `protobuf:"varint,2,opt,name=Week,proto3" json:"Week,omitempty"`
	Year    int32              `protobuf:"varint,3,opt,name=Year,proto3" json:"Year,omitempty"`
}

func (x *ProminentPlayerList) Reset() {
	*x = ProminentPlayerList{}
	mi := &file_proto_players_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProminentPlayerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProminentPlayerList) ProtoMessage() {}

func (x *ProminentPlayerList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProminentPlayerList.ProtoReflect.Descriptor instead.
func (*ProminentPlayerList) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{7}
}

func (x *ProminentPlayerList) GetPlayers() []*ProminentPlayer {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *ProminentPlayerList) GetWeek() int32 {
	if x != nil {
		return x.Week
	}
	return 0
}

func (x *ProminentPlayerList) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

var File_proto_players_proto protoreflect.FileDescriptor

var file_proto_players_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0xae,
	0x01, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22,
	0xb7, 0x02, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4b,
	0x64, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x4b, 0x64,
	0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x4b, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12,
	0x22, 0x0a, 0x0c, 0x4b, 0x69, 0x6c, 0x6c, 0x73, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x4b, 0x69, 0x6c, 0x6c, 0x73, 0x41, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x44, 0x65, 0x61, 0x74, 0x68, 0x73, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x44, 0x65, 0x61, 0x74,
	0x68, 0x73, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x48, 0x65, 0x61,
	0x64, 0x73, 0x68, 0x6f, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x41, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x16, 0x48, 0x65, 0x61, 0x64, 0x73,
	0x68, 0x6f, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x56, 0x50, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x4d, 0x56, 0x50, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x41, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x6c, 0x6f, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x45, 0x6c, 0x6f, 0x22, 0x37, 0x0a, 0x0a, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2e, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x10,
	0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x65, 0x61,
	0x6d, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x12, 0x2a, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0xa3,
	0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x22, 0x71, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x6e, 0x65, 0x6e,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x6e, 0x65, 0x6e, 0x74,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x57, 0x65, 0x65, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x57,
	0x65, 0x65, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x32, 0xfd, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x6d, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x12, 0x0e, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1c, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x69,
	0x6e, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x37,
	0x0a, 0x09, 0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x17, 0x5a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x3b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_players_proto_rawDescOnce sync.Once
	file_proto_players_proto_rawDescData = file_proto_players_proto_rawDesc
)

func file_proto_players_proto_rawDescGZIP() []byte {
	file_proto_players_proto_rawDescOnce.Do(func() {
		file_proto_players_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_players_proto_rawDescData)
	})
	return file_proto_players_proto_rawDescData
}

var file_proto_players_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_players_proto_goTypes = []any{
	(*Player)(nil),              // 0: players.Player
	(*PlayerStats)(nil),         // 1: players.PlayerStats
	(*PlayerList)(nil),          // 2: players.PlayerList
	(*Empty)(nil),               // 3: players.Empty
	(*GetPlayerRequest)(nil),    // 4: players.GetPlayerRequest
	(*NewPlayerRequest)(nil),    // 5: players.NewPlayerRequest
	(*ProminentPlayer)(nil),     // 6: players.ProminentPlayer
	(*ProminentPlayerList)(nil), // 7: players.ProminentPlayerList
}
var file_proto_players_proto_depIdxs = []int32{
	1, // 0: players.Player.Stats:type_name -> players.PlayerStats
	0, // 1: players.PlayerList.players:type_name -> players.Player
	1, // 2: players.NewPlayerRequest.Stats:type_name -> players.PlayerStats
	6, // 3: players.ProminentPlayerList.players:type_name -> players.ProminentPlayer
	4, // 4: players.PlayerService.GetPlayer:input_type -> players.GetPlayerRequest
	3, // 5: players.PlayerService.GetPlayers:input_type -> players.Empty
	3, // 6: players.PlayerService.GetProminentPlayers:input_type -> players.Empty
	5, // 7: players.PlayerService.NewPlayer:input_type -> players.NewPlayerRequest
	2, // 8: players.PlayerService.GetPlayer:output_type -> players.PlayerList
	2, // 9: players.PlayerService.GetPlayers:output_type -> players.PlayerList
	7, // 10: players.PlayerService.GetProminentPlayers:output_type -> players.ProminentPlayerList
	0, // 11: players.PlayerService.NewPlayer:output_type -> players.Player
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_players_proto_init() }
func file_proto_players_proto_init() {
	if File_proto_players_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_players_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_players_proto_goTypes,
		DependencyIndexes: file_proto_players_proto_depIdxs,
		MessageInfos:      file_proto_players_proto_msgTypes,
	}.Build()
	File_proto_players_proto = out.File
	file_proto_players_proto_rawDesc = nil
	file_proto_players_proto_goTypes = nil
	file_proto_players_proto_depIdxs = nil
}
