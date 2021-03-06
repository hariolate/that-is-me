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
	return file_game_proto_enumTypes[0].Descriptor()
}

func (GameResultEvent_Type) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[0]
}

func (x GameResultEvent_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameResultEvent_Type.Descriptor instead.
func (GameResultEvent_Type) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{8, 0}
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
		mi := &file_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
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

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position *Position `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	ObjectId uint32    `protobuf:"varint,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

func (x *Object) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Object) GetObjectId() uint32 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

type UserAKillEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId uint32 `protobuf:"varint,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
}

func (x *UserAKillEvent) Reset() {
	*x = UserAKillEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAKillEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAKillEvent) ProtoMessage() {}

func (x *UserAKillEvent) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserAKillEvent.ProtoReflect.Descriptor instead.
func (*UserAKillEvent) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{2}
}

func (x *UserAKillEvent) GetObjectId() uint32 {
	if x != nil {
		return x.ObjectId
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
		mi := &file_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBMoveEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBMoveEvent) ProtoMessage() {}

func (x *UserBMoveEvent) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserBMoveEvent.ProtoReflect.Descriptor instead.
func (*UserBMoveEvent) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{3}
}

func (x *UserBMoveEvent) GetTo() *Position {
	if x != nil {
		return x.To
	}
	return nil
}

type UserAInitEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objects []*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *UserAInitEvent) Reset() {
	*x = UserAInitEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAInitEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAInitEvent) ProtoMessage() {}

func (x *UserAInitEvent) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserAInitEvent.ProtoReflect.Descriptor instead.
func (*UserAInitEvent) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{4}
}

func (x *UserAInitEvent) GetObjects() []*Object {
	if x != nil {
		return x.Objects
	}
	return nil
}

type UserBInitEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objects        []*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	ThatObjectIsMe uint32    `protobuf:"varint,2,opt,name=that_object_is_me,json=thatObjectIsMe,proto3" json:"that_object_is_me,omitempty"`
}

func (x *UserBInitEvent) Reset() {
	*x = UserBInitEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBInitEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBInitEvent) ProtoMessage() {}

func (x *UserBInitEvent) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserBInitEvent.ProtoReflect.Descriptor instead.
func (*UserBInitEvent) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{5}
}

func (x *UserBInitEvent) GetObjects() []*Object {
	if x != nil {
		return x.Objects
	}
	return nil
}

func (x *UserBInitEvent) GetThatObjectIsMe() uint32 {
	if x != nil {
		return x.ThatObjectIsMe
	}
	return 0
}

type ClientReadyEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientReadyEvent) Reset() {
	*x = ClientReadyEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientReadyEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientReadyEvent) ProtoMessage() {}

func (x *ClientReadyEvent) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ClientReadyEvent.ProtoReflect.Descriptor instead.
func (*ClientReadyEvent) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{6}
}

type ObjectsLocationUpdateEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objects []*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *ObjectsLocationUpdateEvent) Reset() {
	*x = ObjectsLocationUpdateEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectsLocationUpdateEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectsLocationUpdateEvent) ProtoMessage() {}

func (x *ObjectsLocationUpdateEvent) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectsLocationUpdateEvent.ProtoReflect.Descriptor instead.
func (*ObjectsLocationUpdateEvent) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{7}
}

func (x *ObjectsLocationUpdateEvent) GetObjects() []*Object {
	if x != nil {
		return x.Objects
	}
	return nil
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
		mi := &file_game_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameResultEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameResultEvent) ProtoMessage() {}

func (x *GameResultEvent) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[8]
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
	return file_game_proto_rawDescGZIP(), []int{8}
}

func (x *GameResultEvent) GetType() GameResultEvent_Type {
	if x != nil {
		return x.Type
	}
	return GameResultEvent_Success
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x08,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x01, 0x79, 0x22, 0x4c, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25,
	0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x22, 0x2d, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x4b, 0x69, 0x6c, 0x6c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x22, 0x2b, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x4d, 0x6f, 0x76, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x33,
	0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x49, 0x6e, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x21, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x22, 0x5e, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x49, 0x6e, 0x69, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x11, 0x74, 0x68, 0x61, 0x74,
	0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x73, 0x5f, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74, 0x68, 0x61, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x73, 0x4d, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x61,
	0x64, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x3f, 0x0a, 0x1a, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x5e, 0x0a, 0x0f, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x20, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x10, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_game_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_game_proto_goTypes = []interface{}{
	(GameResultEvent_Type)(0),          // 0: GameResultEvent.Type
	(*Position)(nil),                   // 1: Position
	(*Object)(nil),                     // 2: Object
	(*UserAKillEvent)(nil),             // 3: UserAKillEvent
	(*UserBMoveEvent)(nil),             // 4: UserBMoveEvent
	(*UserAInitEvent)(nil),             // 5: UserAInitEvent
	(*UserBInitEvent)(nil),             // 6: UserBInitEvent
	(*ClientReadyEvent)(nil),           // 7: ClientReadyEvent
	(*ObjectsLocationUpdateEvent)(nil), // 8: ObjectsLocationUpdateEvent
	(*GameResultEvent)(nil),            // 9: GameResultEvent
}
var file_game_proto_depIdxs = []int32{
	1, // 0: Object.position:type_name -> Position
	1, // 1: UserBMoveEvent.to:type_name -> Position
	2, // 2: UserAInitEvent.objects:type_name -> Object
	2, // 3: UserBInitEvent.objects:type_name -> Object
	2, // 4: ObjectsLocationUpdateEvent.objects:type_name -> Object
	0, // 5: GameResultEvent.type:type_name -> GameResultEvent.Type
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
			switch v := v.(*UserAKillEvent); i {
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
		file_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAInitEvent); i {
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
			switch v := v.(*UserBInitEvent); i {
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
			switch v := v.(*ClientReadyEvent); i {
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
		file_game_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectsLocationUpdateEvent); i {
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
		file_game_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
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
