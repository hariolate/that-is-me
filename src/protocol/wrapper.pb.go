// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: wrapper.proto

package protocol

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Wrapper_MessageType int32

const (
	Wrapper_MatchMakingRequest         Wrapper_MessageType = 0
	Wrapper_MatchMakingResponse        Wrapper_MessageType = 1
	Wrapper_MatchStateChangeEvent      Wrapper_MessageType = 2
	Wrapper_UserAKillEvent             Wrapper_MessageType = 3
	Wrapper_UserBMoveEvent             Wrapper_MessageType = 4
	Wrapper_UserAInitEvent             Wrapper_MessageType = 5
	Wrapper_UserBInitEvent             Wrapper_MessageType = 6
	Wrapper_ClientReadyEvent           Wrapper_MessageType = 7
	Wrapper_ObjectsLocationUpdateEvent Wrapper_MessageType = 8
	Wrapper_GameResultEvent            Wrapper_MessageType = 9
	Wrapper_RawChatMessageEvent        Wrapper_MessageType = 10
	Wrapper_NewMessageEvent            Wrapper_MessageType = 11
)

// Enum value maps for Wrapper_MessageType.
var (
	Wrapper_MessageType_name = map[int32]string{
		0:  "MatchMakingRequest",
		1:  "MatchMakingResponse",
		2:  "MatchStateChangeEvent",
		3:  "UserAKillEvent",
		4:  "UserBMoveEvent",
		5:  "UserAInitEvent",
		6:  "UserBInitEvent",
		7:  "ClientReadyEvent",
		8:  "ObjectsLocationUpdateEvent",
		9:  "GameResultEvent",
		10: "RawChatMessageEvent",
		11: "NewMessageEvent",
	}
	Wrapper_MessageType_value = map[string]int32{
		"MatchMakingRequest":         0,
		"MatchMakingResponse":        1,
		"MatchStateChangeEvent":      2,
		"UserAKillEvent":             3,
		"UserBMoveEvent":             4,
		"UserAInitEvent":             5,
		"UserBInitEvent":             6,
		"ClientReadyEvent":           7,
		"ObjectsLocationUpdateEvent": 8,
		"GameResultEvent":            9,
		"RawChatMessageEvent":        10,
		"NewMessageEvent":            11,
	}
)

func (x Wrapper_MessageType) Enum() *Wrapper_MessageType {
	p := new(Wrapper_MessageType)
	*p = x
	return p
}

func (x Wrapper_MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Wrapper_MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_wrapper_proto_enumTypes[0].Descriptor()
}

func (Wrapper_MessageType) Type() protoreflect.EnumType {
	return &file_wrapper_proto_enumTypes[0]
}

func (x Wrapper_MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Wrapper_MessageType.Descriptor instead.
func (Wrapper_MessageType) EnumDescriptor() ([]byte, []int) {
	return file_wrapper_proto_rawDescGZIP(), []int{0, 0}
}

type Wrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    Wrapper_MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=Wrapper_MessageType" json:"type,omitempty"`
	Message *any.Any            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Wrapper) Reset() {
	*x = Wrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrapper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wrapper) ProtoMessage() {}

func (x *Wrapper) ProtoReflect() protoreflect.Message {
	mi := &file_wrapper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wrapper.ProtoReflect.Descriptor instead.
func (*Wrapper) Descriptor() ([]byte, []int) {
	return file_wrapper_proto_rawDescGZIP(), []int{0}
}

func (x *Wrapper) GetType() Wrapper_MessageType {
	if x != nil {
		return x.Type
	}
	return Wrapper_MatchMakingRequest
}

func (x *Wrapper) GetMessage() *any.Any {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_wrapper_proto protoreflect.FileDescriptor

var file_wrapper_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x03, 0x0a, 0x07, 0x57,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0xa2, 0x02, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x12, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x4d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10,
	0x01, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x4b, 0x69, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x03,
	0x12, 0x12, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x4d, 0x6f, 0x76, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x49, 0x6e, 0x69,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x49, 0x6e, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x10, 0x07, 0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x09, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x61, 0x77, 0x43, 0x68,
	0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x0a,
	0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x10, 0x0b, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wrapper_proto_rawDescOnce sync.Once
	file_wrapper_proto_rawDescData = file_wrapper_proto_rawDesc
)

func file_wrapper_proto_rawDescGZIP() []byte {
	file_wrapper_proto_rawDescOnce.Do(func() {
		file_wrapper_proto_rawDescData = protoimpl.X.CompressGZIP(file_wrapper_proto_rawDescData)
	})
	return file_wrapper_proto_rawDescData
}

var file_wrapper_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wrapper_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wrapper_proto_goTypes = []interface{}{
	(Wrapper_MessageType)(0), // 0: Wrapper.MessageType
	(*Wrapper)(nil),          // 1: Wrapper
	(*any.Any)(nil),          // 2: google.protobuf.Any
}
var file_wrapper_proto_depIdxs = []int32{
	0, // 0: Wrapper.type:type_name -> Wrapper.MessageType
	2, // 1: Wrapper.message:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wrapper_proto_init() }
func file_wrapper_proto_init() {
	if File_wrapper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wrapper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wrapper); i {
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
			RawDescriptor: file_wrapper_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wrapper_proto_goTypes,
		DependencyIndexes: file_wrapper_proto_depIdxs,
		EnumInfos:         file_wrapper_proto_enumTypes,
		MessageInfos:      file_wrapper_proto_msgTypes,
	}.Build()
	File_wrapper_proto = out.File
	file_wrapper_proto_rawDesc = nil
	file_wrapper_proto_goTypes = nil
	file_wrapper_proto_depIdxs = nil
}
