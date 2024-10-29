// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: proto/users.proto

package users

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	players "ibercs/proto/players"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int32           `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PlayerID    string          `protobuf:"bytes,2,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	Name        string          `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string          `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	Twitter     string          `protobuf:"bytes,5,opt,name=Twitter,proto3" json:"Twitter,omitempty"`
	Twitch      string          `protobuf:"bytes,6,opt,name=Twitch,proto3" json:"Twitch,omitempty"`
	Role        int32           `protobuf:"varint,7,opt,name=Role,proto3" json:"Role,omitempty"`
	Player      *players.Player `protobuf:"bytes,8,opt,name=Player,proto3" json:"Player,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_proto_users_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_users_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_users_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetPlayerID() string {
	if x != nil {
		return x.PlayerID
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *User) GetTwitter() string {
	if x != nil {
		return x.Twitter
	}
	return ""
}

func (x *User) GetTwitch() string {
	if x != nil {
		return x.Twitch
	}
	return ""
}

func (x *User) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *User) GetPlayer() *players.Player {
	if x != nil {
		return x.Player
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
	mi := &file_proto_users_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_users_proto_msgTypes[1]
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
	return file_proto_users_proto_rawDescGZIP(), []int{1}
}

type StreamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Streams []*StreamResponse `protobuf:"bytes,1,rep,name=streams,proto3" json:"streams,omitempty"`
}

func (x *StreamsResponse) Reset() {
	*x = StreamsResponse{}
	mi := &file_proto_users_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamsResponse) ProtoMessage() {}

func (x *StreamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_users_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamsResponse.ProtoReflect.Descriptor instead.
func (*StreamsResponse) Descriptor() ([]byte, []int) {
	return file_proto_users_proto_rawDescGZIP(), []int{2}
}

func (x *StreamsResponse) GetStreams() []*StreamResponse {
	if x != nil {
		return x.Streams
	}
	return nil
}

type StreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stream string `protobuf:"bytes,1,opt,name=Stream,proto3" json:"Stream,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	mi := &file_proto_users_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_users_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_proto_users_proto_rawDescGZIP(), []int{3}
}

func (x *StreamResponse) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *StreamResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type NewSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NewSessionRequest) Reset() {
	*x = NewSessionRequest{}
	mi := &file_proto_users_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewSessionRequest) ProtoMessage() {}

func (x *NewSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_users_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewSessionRequest.ProtoReflect.Descriptor instead.
func (*NewSessionRequest) Descriptor() ([]byte, []int) {
	return file_proto_users_proto_rawDescGZIP(), []int{4}
}

func (x *NewSessionRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NewSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *NewSessionResponse) Reset() {
	*x = NewSessionResponse{}
	mi := &file_proto_users_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewSessionResponse) ProtoMessage() {}

func (x *NewSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_users_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewSessionResponse.ProtoReflect.Descriptor instead.
func (*NewSessionResponse) Descriptor() ([]byte, []int) {
	return file_proto_users_proto_rawDescGZIP(), []int{5}
}

func (x *NewSessionResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type NewUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceitId string `protobuf:"bytes,1,opt,name=FaceitId,proto3" json:"FaceitId,omitempty"`
}

func (x *NewUserRequest) Reset() {
	*x = NewUserRequest{}
	mi := &file_proto_users_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewUserRequest) ProtoMessage() {}

func (x *NewUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_users_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewUserRequest.ProtoReflect.Descriptor instead.
func (*NewUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_users_proto_rawDescGZIP(), []int{6}
}

func (x *NewUserRequest) GetFaceitId() string {
	if x != nil {
		return x.FaceitId
	}
	return ""
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_proto_users_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_users_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_users_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_users_proto protoreflect.FileDescriptor

var file_proto_users_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd7, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x77,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x77, 0x69,
	0x74, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x77, 0x69, 0x74, 0x63, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x27, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x42, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x3c, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x12, 0x4e, 0x65, 0x77,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x0e, 0x4e,
	0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x32, 0xcb, 0x03, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x46, 0x61, 0x63, 0x65, 0x69, 0x74, 0x49, 0x64, 0x12,
	0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x26, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x07, 0x4e, 0x65,
	0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4e, 0x65,
	0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0a, 0x4e, 0x65, 0x77,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x4e, 0x65, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x4e, 0x65, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x12, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_users_proto_rawDescOnce sync.Once
	file_proto_users_proto_rawDescData = file_proto_users_proto_rawDesc
)

func file_proto_users_proto_rawDescGZIP() []byte {
	file_proto_users_proto_rawDescOnce.Do(func() {
		file_proto_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_users_proto_rawDescData)
	})
	return file_proto_users_proto_rawDescData
}

var file_proto_users_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_users_proto_goTypes = []any{
	(*User)(nil),               // 0: users.User
	(*Empty)(nil),              // 1: users.Empty
	(*StreamsResponse)(nil),    // 2: users.StreamsResponse
	(*StreamResponse)(nil),     // 3: users.StreamResponse
	(*NewSessionRequest)(nil),  // 4: users.NewSessionRequest
	(*NewSessionResponse)(nil), // 5: users.NewSessionResponse
	(*NewUserRequest)(nil),     // 6: users.NewUserRequest
	(*GetUserRequest)(nil),     // 7: users.GetUserRequest
	(*players.Player)(nil),     // 8: players.Player
}
var file_proto_users_proto_depIdxs = []int32{
	8,  // 0: users.User.Player:type_name -> players.Player
	3,  // 1: users.StreamsResponse.streams:type_name -> users.StreamResponse
	7,  // 2: users.UserService.GetUser:input_type -> users.GetUserRequest
	7,  // 3: users.UserService.GetUserByFaceitId:input_type -> users.GetUserRequest
	7,  // 4: users.UserService.GetUserByPlayerNickname:input_type -> users.GetUserRequest
	0,  // 5: users.UserService.UpdateUser:input_type -> users.User
	6,  // 6: users.UserService.NewUser:input_type -> users.NewUserRequest
	4,  // 7: users.UserService.NewSession:input_type -> users.NewSessionRequest
	4,  // 8: users.UserService.DeleteSession:input_type -> users.NewSessionRequest
	1,  // 9: users.UserService.GetAllStreams:input_type -> users.Empty
	0,  // 10: users.UserService.GetUser:output_type -> users.User
	0,  // 11: users.UserService.GetUserByFaceitId:output_type -> users.User
	0,  // 12: users.UserService.GetUserByPlayerNickname:output_type -> users.User
	0,  // 13: users.UserService.UpdateUser:output_type -> users.User
	0,  // 14: users.UserService.NewUser:output_type -> users.User
	5,  // 15: users.UserService.NewSession:output_type -> users.NewSessionResponse
	5,  // 16: users.UserService.DeleteSession:output_type -> users.NewSessionResponse
	2,  // 17: users.UserService.GetAllStreams:output_type -> users.StreamsResponse
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_users_proto_init() }
func file_proto_users_proto_init() {
	if File_proto_users_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_users_proto_goTypes,
		DependencyIndexes: file_proto_users_proto_depIdxs,
		MessageInfos:      file_proto_users_proto_msgTypes,
	}.Build()
	File_proto_users_proto = out.File
	file_proto_users_proto_rawDesc = nil
	file_proto_users_proto_goTypes = nil
	file_proto_users_proto_depIdxs = nil
}