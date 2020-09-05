// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: game.proto

package protocol

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GameStateChangeEvent_Type int32

const (
	GameStateChangeEvent_Ready GameStateChangeEvent_Type = 0
	GameStateChangeEvent_End   GameStateChangeEvent_Type = 1
)

// Enum value maps for GameStateChangeEvent_Type.
var (
	GameStateChangeEvent_Type_name = map[int32]string{
		0: "Ready",
		1: "End",
	}
	GameStateChangeEvent_Type_value = map[string]int32{
		"Ready": 0,
		"End":   1,
	}
)

func (x GameStateChangeEvent_Type) Enum() *GameStateChangeEvent_Type {
	p := new(GameStateChangeEvent_Type)
	*p = x
	return p
}

func (x GameStateChangeEvent_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameStateChangeEvent_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[0].Descriptor()
}

func (GameStateChangeEvent_Type) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[0]
}

func (x GameStateChangeEvent_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameStateChangeEvent_Type.Descriptor instead.
func (GameStateChangeEvent_Type) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{3, 0}
}

type ClientMatchStatusRequest_Type int32

const (
	ClientMatchStatusRequest_Ready ClientMatchStatusRequest_Type = 0
	ClientMatchStatusRequest_End   ClientMatchStatusRequest_Type = 1
)

// Enum value maps for ClientMatchStatusRequest_Type.
var (
	ClientMatchStatusRequest_Type_name = map[int32]string{
		0: "Ready",
		1: "End",
	}
	ClientMatchStatusRequest_Type_value = map[string]int32{
		"Ready": 0,
		"End":   1,
	}
)

func (x ClientMatchStatusRequest_Type) Enum() *ClientMatchStatusRequest_Type {
	p := new(ClientMatchStatusRequest_Type)
	*p = x
	return p
}

func (x ClientMatchStatusRequest_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientMatchStatusRequest_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[1].Descriptor()
}

func (ClientMatchStatusRequest_Type) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[1]
}

func (x ClientMatchStatusRequest_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientMatchStatusRequest_Type.Descriptor instead.
func (ClientMatchStatusRequest_Type) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{4, 0}
}

type GameResultEvent_Type int32

const (
	GameResultEvent_Success GameResultEvent_Type = 0
	GameResultEvent_Failure GameResultEvent_Type = 1
)

// Enum value maps for GameResultEvent_Type.
var (
	GameResultEvent_Type_name = map[int32]string{
		0: "Success",
		1: "Failure",
	}
	GameResultEvent_Type_value = map[string]int32{
		"Success": 0,
		"Failure": 1,
	}
)

func (x GameResultEvent_Type) Enum() *GameResultEvent_Type {
	p := new(GameResultEvent_Type)
	*p = x
	return p
}

func (x GameResultEvent_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameResultEvent_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[2].Descriptor()
}

func (GameResultEvent_Type) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[2]
}

func (x GameResultEvent_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameResultEvent_Type.Descriptor instead.
func (GameResultEvent_Type) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{5, 0}
}

type ChatSessionEvent_Type int32

const (
	ChatSessionEvent_Begin ChatSessionEvent_Type = 0
	ChatSessionEvent_End   ChatSessionEvent_Type = 1
)

// Enum value maps for ChatSessionEvent_Type.
var (
	ChatSessionEvent_Type_name = map[int32]string{
		0: "Begin",
		1: "End",
	}
	ChatSessionEvent_Type_value = map[string]int32{
		"Begin": 0,
		"End":   1,
	}
)

func (x ChatSessionEvent_Type) Enum() *ChatSessionEvent_Type {
	p := new(ChatSessionEvent_Type)
	*p = x
	return p
}

func (x ChatSessionEvent_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatSessionEvent_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[3].Descriptor()
}

func (ChatSessionEvent_Type) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[3]
}

func (x ChatSessionEvent_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatSessionEvent_Type.Descriptor instead.
func (ChatSessionEvent_Type) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{6, 0}
}

type UserATapEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object uint32 `protobuf:"varint,1,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *UserATapEvent) Reset() {
	*x = UserATapEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserATapEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserATapEvent) ProtoMessage() {}

func (x *UserATapEvent) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserATapEvent.ProtoReflect.Descriptor instead.
func (*UserATapEvent) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

func (x *UserATapEvent) GetObject() uint32 {
	if x != nil {
		return x.Object
	}
	return 0
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X uint32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y uint32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

func (x *Position) GetX() uint32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Position) GetY() uint32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type UserBMoveEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To *Position `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *UserBMoveEvent) Reset() {
	*x = UserBMoveEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBMoveEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBMoveEvent) ProtoMessage() {}

func (x *UserBMoveEvent) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBMoveEvent.ProtoReflect.Descriptor instead.
func (*UserBMoveEvent) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{2}
}

func (x *UserBMoveEvent) GetTo() *Position {
	if x != nil {
		return x.To
	}
	return nil
}

type GameStateChangeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type GameStateChangeEvent_Type `protobuf:"varint,1,opt,name=type,proto3,enum=GameStateChangeEvent_Type" json:"type,omitempty"`
}

func (x *GameStateChangeEvent) Reset() {
	*x = GameStateChangeEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStateChangeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStateChangeEvent) ProtoMessage() {}

func (x *GameStateChangeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStateChangeEvent.ProtoReflect.Descriptor instead.
func (*GameStateChangeEvent) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{3}
}

func (x *GameStateChangeEvent) GetType() GameStateChangeEvent_Type {
	if x != nil {
		return x.Type
	}
	return GameStateChangeEvent_Ready
}

type ClientMatchStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ClientMatchStatusRequest_Type `protobuf:"varint,1,opt,name=type,proto3,enum=ClientMatchStatusRequest_Type" json:"type,omitempty"`
}

func (x *ClientMatchStatusRequest) Reset() {
	*x = ClientMatchStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMatchStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMatchStatusRequest) ProtoMessage() {}

func (x *ClientMatchStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMatchStatusRequest.ProtoReflect.Descriptor instead.
func (*ClientMatchStatusRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{4}
}

func (x *ClientMatchStatusRequest) GetType() ClientMatchStatusRequest_Type {
	if x != nil {
		return x.Type
	}
	return ClientMatchStatusRequest_Ready
}

type GameResultEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type GameResultEvent_Type `protobuf:"varint,1,opt,name=type,proto3,enum=GameResultEvent_Type" json:"type,omitempty"`
}

func (x *GameResultEvent) Reset() {
	*x = GameResultEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameResultEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameResultEvent) ProtoMessage() {}

func (x *GameResultEvent) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameResultEvent.ProtoReflect.Descriptor instead.
func (*GameResultEvent) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{5}
}

func (x *GameResultEvent) GetType() GameResultEvent_Type {
	if x != nil {
		return x.Type
	}
	return GameResultEvent_Success
}

type ChatSessionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ChatSessionEvent_Type `protobuf:"varint,1,opt,name=type,proto3,enum=ChatSessionEvent_Type" json:"type,omitempty"`
}

func (x *ChatSessionEvent) Reset() {
	*x = ChatSessionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatSessionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatSessionEvent) ProtoMessage() {}

func (x *ChatSessionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatSessionEvent.ProtoReflect.Descriptor instead.
func (*ChatSessionEvent) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{6}
}

func (x *ChatSessionEvent) GetType() ChatSessionEvent_Type {
	if x != nil {
		return x.Type
	}
	return ChatSessionEvent_Begin
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x0d,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x54, 0x61, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x26, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x78, 0x12,
	0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x79, 0x22, 0x2b, 0x0a,
	0x0e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x4d, 0x6f, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x19, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x14, 0x47, 0x61,
	0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x1a, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x65,
	0x61, 0x64, 0x79, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x6e, 0x64, 0x10, 0x01, 0x22, 0x6a,
	0x0a, 0x18, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x1a,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x65, 0x61, 0x64, 0x79, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x6e, 0x64, 0x10, 0x01, 0x22, 0x5e, 0x0a, 0x0f, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x20, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x10, 0x01, 0x22, 0x5a, 0x0a, 0x10, 0x43, 0x68,
	0x61, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2a,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x1a, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x45, 0x6e, 0x64, 0x10, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_proto_rawDescOnce sync.Once
	file_game_proto_rawDescData = file_game_proto_rawDesc
)

func file_game_proto_rawDescGZIP() []byte {
	file_game_proto_rawDescOnce.Do(func() {
		file_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_proto_rawDescData)
	})
	return file_game_proto_rawDescData
}

var file_game_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_game_proto_goTypes = []interface{}{
	(GameStateChangeEvent_Type)(0),     // 0: GameStateChangeEvent.Type
	(ClientMatchStatusRequest_Type)(0), // 1: ClientMatchStatusRequest.Type
	(GameResultEvent_Type)(0),          // 2: GameResultEvent.Type
	(ChatSessionEvent_Type)(0),         // 3: ChatSessionEvent.Type
	(*UserATapEvent)(nil),              // 4: UserATapEvent
	(*Position)(nil),                   // 5: Position
	(*UserBMoveEvent)(nil),             // 6: UserBMoveEvent
	(*GameStateChangeEvent)(nil),       // 7: GameStateChangeEvent
	(*ClientMatchStatusRequest)(nil),   // 8: ClientMatchStatusRequest
	(*GameResultEvent)(nil),            // 9: GameResultEvent
	(*ChatSessionEvent)(nil),           // 10: ChatSessionEvent
}
var file_game_proto_depIdxs = []int32{
	5, // 0: UserBMoveEvent.to:type_name -> Position
	0, // 1: GameStateChangeEvent.type:type_name -> GameStateChangeEvent.Type
	1, // 2: ClientMatchStatusRequest.type:type_name -> ClientMatchStatusRequest.Type
	2, // 3: GameResultEvent.type:type_name -> GameResultEvent.Type
	3, // 4: ChatSessionEvent.type:type_name -> ChatSessionEvent.Type
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserATapEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBMoveEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameStateChangeEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMatchStatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameResultEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatSessionEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_proto_goTypes,
		DependencyIndexes: file_game_proto_depIdxs,
		EnumInfos:         file_game_proto_enumTypes,
		MessageInfos:      file_game_proto_msgTypes,
	}.Build()
	File_game_proto = out.File
	file_game_proto_rawDesc = nil
	file_game_proto_goTypes = nil
	file_game_proto_depIdxs = nil
}
