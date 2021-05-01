// Proto file for movie info service. Note this is gRPC proto syntax (not Go)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: gameapi/gameapi.proto

package gameapi

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

// The request message containing monster action
type MonsterAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Action string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	GameID string `protobuf:"bytes,3,opt,name=gameID,proto3" json:"gameID,omitempty"`
}

func (x *MonsterAction) Reset() {
	*x = MonsterAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gameapi_gameapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonsterAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonsterAction) ProtoMessage() {}

func (x *MonsterAction) ProtoReflect() protoreflect.Message {
	mi := &file_gameapi_gameapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonsterAction.ProtoReflect.Descriptor instead.
func (*MonsterAction) Descriptor() ([]byte, []int) {
	return file_gameapi_gameapi_proto_rawDescGZIP(), []int{0}
}

func (x *MonsterAction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MonsterAction) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *MonsterAction) GetGameID() string {
	if x != nil {
		return x.GameID
	}
	return ""
}

type HealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	GameID string `protobuf:"bytes,2,opt,name=gameID,proto3" json:"gameID,omitempty"`
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gameapi_gameapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gameapi_gameapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_gameapi_gameapi_proto_rawDescGZIP(), []int{1}
}

func (x *HealthRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HealthRequest) GetGameID() string {
	if x != nil {
		return x.GameID
	}
	return ""
}

type HealthPoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Health     int32  `protobuf:"varint,1,opt,name=health,proto3" json:"health,omitempty"`
	WhoseTurn  string `protobuf:"bytes,2,opt,name=whoseTurn,proto3" json:"whoseTurn,omitempty"`
	LastAttack string `protobuf:"bytes,3,opt,name=lastAttack,proto3" json:"lastAttack,omitempty"`
	Damage     int32  `protobuf:"varint,4,opt,name=damage,proto3" json:"damage,omitempty"`
}

func (x *HealthPoints) Reset() {
	*x = HealthPoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gameapi_gameapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthPoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthPoints) ProtoMessage() {}

func (x *HealthPoints) ProtoReflect() protoreflect.Message {
	mi := &file_gameapi_gameapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthPoints.ProtoReflect.Descriptor instead.
func (*HealthPoints) Descriptor() ([]byte, []int) {
	return file_gameapi_gameapi_proto_rawDescGZIP(), []int{2}
}

func (x *HealthPoints) GetHealth() int32 {
	if x != nil {
		return x.Health
	}
	return 0
}

func (x *HealthPoints) GetWhoseTurn() string {
	if x != nil {
		return x.WhoseTurn
	}
	return ""
}

func (x *HealthPoints) GetLastAttack() string {
	if x != nil {
		return x.LastAttack
	}
	return ""
}

func (x *HealthPoints) GetDamage() int32 {
	if x != nil {
		return x.Damage
	}
	return 0
}

type GameStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpponentName    string `protobuf:"bytes,1,opt,name=opponentName,proto3" json:"opponentName,omitempty"`
	OpponentMonster string `protobuf:"bytes,2,opt,name=opponentMonster,proto3" json:"opponentMonster,omitempty"`
	OpponentHealth  int32  `protobuf:"varint,3,opt,name=opponentHealth,proto3" json:"opponentHealth,omitempty"`
	WhoseTurn       string `protobuf:"bytes,4,opt,name=whoseTurn,proto3" json:"whoseTurn,omitempty"`
	MyHealth        int32  `protobuf:"varint,5,opt,name=myHealth,proto3" json:"myHealth,omitempty"`
	MyMonster       string `protobuf:"bytes,6,opt,name=myMonster,proto3" json:"myMonster,omitempty"`
	Code            string `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"`
	GameID          string `protobuf:"bytes,8,opt,name=gameID,proto3" json:"gameID,omitempty"`
}

func (x *GameStatus) Reset() {
	*x = GameStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gameapi_gameapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStatus) ProtoMessage() {}

func (x *GameStatus) ProtoReflect() protoreflect.Message {
	mi := &file_gameapi_gameapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStatus.ProtoReflect.Descriptor instead.
func (*GameStatus) Descriptor() ([]byte, []int) {
	return file_gameapi_gameapi_proto_rawDescGZIP(), []int{3}
}

func (x *GameStatus) GetOpponentName() string {
	if x != nil {
		return x.OpponentName
	}
	return ""
}

func (x *GameStatus) GetOpponentMonster() string {
	if x != nil {
		return x.OpponentMonster
	}
	return ""
}

func (x *GameStatus) GetOpponentHealth() int32 {
	if x != nil {
		return x.OpponentHealth
	}
	return 0
}

func (x *GameStatus) GetWhoseTurn() string {
	if x != nil {
		return x.WhoseTurn
	}
	return ""
}

func (x *GameStatus) GetMyHealth() int32 {
	if x != nil {
		return x.MyHealth
	}
	return 0
}

func (x *GameStatus) GetMyMonster() string {
	if x != nil {
		return x.MyMonster
	}
	return ""
}

func (x *GameStatus) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GameStatus) GetGameID() string {
	if x != nil {
		return x.GameID
	}
	return ""
}

type RequestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RequestInfo) Reset() {
	*x = RequestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gameapi_gameapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestInfo) ProtoMessage() {}

func (x *RequestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_gameapi_gameapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestInfo.ProtoReflect.Descriptor instead.
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return file_gameapi_gameapi_proto_rawDescGZIP(), []int{4}
}

func (x *RequestInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_gameapi_gameapi_proto protoreflect.FileDescriptor

var file_gameapi_gameapi_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x61, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x61, 0x70, 0x69,
	0x22, 0x53, 0x0a, 0x0d, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67,
	0x61, 0x6d, 0x65, 0x49, 0x44, 0x22, 0x3b, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65,
	0x49, 0x44, 0x22, 0x7c, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x68,
	0x6f, 0x73, 0x65, 0x54, 0x75, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77,
	0x68, 0x6f, 0x73, 0x65, 0x54, 0x75, 0x72, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61,
	0x73, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65,
	0x22, 0x86, 0x02, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x12, 0x26, 0x0a,
	0x0e, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x68, 0x6f, 0x73, 0x65, 0x54, 0x75,
	0x72, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x68, 0x6f, 0x73, 0x65, 0x54,
	0x75, 0x72, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12,
	0x1c, 0x0a, 0x09, 0x6d, 0x79, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x79, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x22, 0x21, 0x0a, 0x0b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xcc, 0x01, 0x0a,
	0x08, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x40, 0x0a, 0x0d, 0x4d, 0x6f, 0x6e,
	0x73, 0x74, 0x65, 0x72, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x13, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gameapi_gameapi_proto_rawDescOnce sync.Once
	file_gameapi_gameapi_proto_rawDescData = file_gameapi_gameapi_proto_rawDesc
)

func file_gameapi_gameapi_proto_rawDescGZIP() []byte {
	file_gameapi_gameapi_proto_rawDescOnce.Do(func() {
		file_gameapi_gameapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_gameapi_gameapi_proto_rawDescData)
	})
	return file_gameapi_gameapi_proto_rawDescData
}

var file_gameapi_gameapi_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gameapi_gameapi_proto_goTypes = []interface{}{
	(*MonsterAction)(nil), // 0: gameapi.MonsterAction
	(*HealthRequest)(nil), // 1: gameapi.HealthRequest
	(*HealthPoints)(nil),  // 2: gameapi.HealthPoints
	(*GameStatus)(nil),    // 3: gameapi.GameStatus
	(*RequestInfo)(nil),   // 4: gameapi.RequestInfo
}
var file_gameapi_gameapi_proto_depIdxs = []int32{
	0, // 0: gameapi.GameInfo.MonsterAttack:input_type -> gameapi.MonsterAction
	1, // 1: gameapi.GameInfo.GetHealthPoints:input_type -> gameapi.HealthRequest
	4, // 2: gameapi.GameInfo.GetGameInfo:input_type -> gameapi.RequestInfo
	2, // 3: gameapi.GameInfo.MonsterAttack:output_type -> gameapi.HealthPoints
	2, // 4: gameapi.GameInfo.GetHealthPoints:output_type -> gameapi.HealthPoints
	3, // 5: gameapi.GameInfo.GetGameInfo:output_type -> gameapi.GameStatus
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gameapi_gameapi_proto_init() }
func file_gameapi_gameapi_proto_init() {
	if File_gameapi_gameapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gameapi_gameapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonsterAction); i {
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
		file_gameapi_gameapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthRequest); i {
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
		file_gameapi_gameapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthPoints); i {
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
		file_gameapi_gameapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameStatus); i {
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
		file_gameapi_gameapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestInfo); i {
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
			RawDescriptor: file_gameapi_gameapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gameapi_gameapi_proto_goTypes,
		DependencyIndexes: file_gameapi_gameapi_proto_depIdxs,
		MessageInfos:      file_gameapi_gameapi_proto_msgTypes,
	}.Build()
	File_gameapi_gameapi_proto = out.File
	file_gameapi_gameapi_proto_rawDesc = nil
	file_gameapi_gameapi_proto_goTypes = nil
	file_gameapi_gameapi_proto_depIdxs = nil
}