// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: teams.proto

package proto

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

type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	FaceitId  string   `protobuf:"bytes,2,opt,name=FaceitId,proto3" json:"FaceitId,omitempty"`
	Name      string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Nickname  string   `protobuf:"bytes,4,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	Avatar    string   `protobuf:"bytes,5,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	PlayersId []string `protobuf:"bytes,6,rep,name=PlayersId,proto3" json:"PlayersId,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	mi := &file_teams_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_teams_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_teams_proto_rawDescGZIP(), []int{0}
}

func (x *Team) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Team) GetFaceitId() string {
	if x != nil {
		return x.FaceitId
	}
	return ""
}

func (x *Team) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Team) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Team) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Team) GetPlayersId() []string {
	if x != nil {
		return x.PlayersId
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
	mi := &file_teams_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_teams_proto_msgTypes[1]
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
	return file_teams_proto_rawDescGZIP(), []int{1}
}

type TeamList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teams []*Team `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
}

func (x *TeamList) Reset() {
	*x = TeamList{}
	mi := &file_teams_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TeamList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamList) ProtoMessage() {}

func (x *TeamList) ProtoReflect() protoreflect.Message {
	mi := &file_teams_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamList.ProtoReflect.Descriptor instead.
func (*TeamList) Descriptor() ([]byte, []int) {
	return file_teams_proto_rawDescGZIP(), []int{2}
}

func (x *TeamList) GetTeams() []*Team {
	if x != nil {
		return x.Teams
	}
	return nil
}

type NewTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceitId string `protobuf:"bytes,1,opt,name=FaceitId,proto3" json:"FaceitId,omitempty"`
}

func (x *NewTeamRequest) Reset() {
	*x = NewTeamRequest{}
	mi := &file_teams_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTeamRequest) ProtoMessage() {}

func (x *NewTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teams_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTeamRequest.ProtoReflect.Descriptor instead.
func (*NewTeamRequest) Descriptor() ([]byte, []int) {
	return file_teams_proto_rawDescGZIP(), []int{3}
}

func (x *NewTeamRequest) GetFaceitId() string {
	if x != nil {
		return x.FaceitId
	}
	return ""
}

var File_teams_proto protoreflect.FileDescriptor

var file_teams_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x65, 0x61, 0x6d, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x04, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x49, 0x64, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x49, 0x64, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2d, 0x0a, 0x08, 0x54, 0x65, 0x61, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x2c, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x54, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x61, 0x63,
	0x65, 0x69, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x61, 0x63,
	0x65, 0x69, 0x74, 0x49, 0x64, 0x32, 0x96, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d,
	0x73, 0x12, 0x0c, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0f, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x74, 0x65,
	0x61, 0x6d, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x12,
	0x2d, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x74, 0x65, 0x61,
	0x6d, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0b, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x42, 0x15,
	0x5a, 0x13, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teams_proto_rawDescOnce sync.Once
	file_teams_proto_rawDescData = file_teams_proto_rawDesc
)

func file_teams_proto_rawDescGZIP() []byte {
	file_teams_proto_rawDescOnce.Do(func() {
		file_teams_proto_rawDescData = protoimpl.X.CompressGZIP(file_teams_proto_rawDescData)
	})
	return file_teams_proto_rawDescData
}

var file_teams_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_teams_proto_goTypes = []any{
	(*Team)(nil),           // 0: teams.Team
	(*Empty)(nil),          // 1: teams.Empty
	(*TeamList)(nil),       // 2: teams.TeamList
	(*NewTeamRequest)(nil), // 3: teams.NewTeamRequest
}
var file_teams_proto_depIdxs = []int32{
	0, // 0: teams.TeamList.teams:type_name -> teams.Team
	1, // 1: teams.TeamService.GetTeams:input_type -> teams.Empty
	3, // 2: teams.TeamService.GetTeam:input_type -> teams.NewTeamRequest
	3, // 3: teams.TeamService.NewTeam:input_type -> teams.NewTeamRequest
	2, // 4: teams.TeamService.GetTeams:output_type -> teams.TeamList
	0, // 5: teams.TeamService.GetTeam:output_type -> teams.Team
	0, // 6: teams.TeamService.NewTeam:output_type -> teams.Team
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_teams_proto_init() }
func file_teams_proto_init() {
	if File_teams_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teams_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teams_proto_goTypes,
		DependencyIndexes: file_teams_proto_depIdxs,
		MessageInfos:      file_teams_proto_msgTypes,
	}.Build()
	File_teams_proto = out.File
	file_teams_proto_rawDesc = nil
	file_teams_proto_goTypes = nil
	file_teams_proto_depIdxs = nil
}
