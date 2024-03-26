// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: pb/service.proto

package pb

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

type Handshake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	GameVariant uint32   `protobuf:"varint,2,opt,name=gameVariant,proto3" json:"gameVariant,omitempty"`
	GameStatus  uint32   `protobuf:"varint,3,opt,name=gameStatus,proto3" json:"gameStatus,omitempty"`
	Addr        string   `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
	PeerAddrs   []string `protobuf:"bytes,5,rep,name=peerAddrs,proto3" json:"peerAddrs,omitempty"`
}

func (x *Handshake) Reset() {
	*x = Handshake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handshake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handshake) ProtoMessage() {}

func (x *Handshake) ProtoReflect() protoreflect.Message {
	mi := &file_pb_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handshake.ProtoReflect.Descriptor instead.
func (*Handshake) Descriptor() ([]byte, []int) {
	return file_pb_service_proto_rawDescGZIP(), []int{0}
}

func (x *Handshake) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Handshake) GetGameVariant() uint32 {
	if x != nil {
		return x.GameVariant
	}
	return 0
}

func (x *Handshake) GetGameStatus() uint32 {
	if x != nil {
		return x.GameStatus
	}
	return 0
}

func (x *Handshake) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Handshake) GetPeerAddrs() []string {
	if x != nil {
		return x.PeerAddrs
	}
	return nil
}

type TakeSeatMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *TakeSeatMsg) Reset() {
	*x = TakeSeatMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeSeatMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeSeatMsg) ProtoMessage() {}

func (x *TakeSeatMsg) ProtoReflect() protoreflect.Message {
	mi := &file_pb_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeSeatMsg.ProtoReflect.Descriptor instead.
func (*TakeSeatMsg) Descriptor() ([]byte, []int) {
	return file_pb_service_proto_rawDescGZIP(), []int{1}
}

func (x *TakeSeatMsg) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_pb_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_pb_service_proto_rawDescGZIP(), []int{2}
}

type ShuffleAndEncryptMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Deck [][]byte `protobuf:"bytes,2,rep,name=deck,proto3" json:"deck,omitempty"`
}

func (x *ShuffleAndEncryptMsg) Reset() {
	*x = ShuffleAndEncryptMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShuffleAndEncryptMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShuffleAndEncryptMsg) ProtoMessage() {}

func (x *ShuffleAndEncryptMsg) ProtoReflect() protoreflect.Message {
	mi := &file_pb_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShuffleAndEncryptMsg.ProtoReflect.Descriptor instead.
func (*ShuffleAndEncryptMsg) Descriptor() ([]byte, []int) {
	return file_pb_service_proto_rawDescGZIP(), []int{3}
}

func (x *ShuffleAndEncryptMsg) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *ShuffleAndEncryptMsg) GetDeck() [][]byte {
	if x != nil {
		return x.Deck
	}
	return nil
}

type SetGameStatusMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameStatus int32 `protobuf:"varint,1,opt,name=gameStatus,proto3" json:"gameStatus,omitempty"`
}

func (x *SetGameStatusMsg) Reset() {
	*x = SetGameStatusMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetGameStatusMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGameStatusMsg) ProtoMessage() {}

func (x *SetGameStatusMsg) ProtoReflect() protoreflect.Message {
	mi := &file_pb_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGameStatusMsg.ProtoReflect.Descriptor instead.
func (*SetGameStatusMsg) Descriptor() ([]byte, []int) {
	return file_pb_service_proto_rawDescGZIP(), []int{4}
}

func (x *SetGameStatusMsg) GetGameStatus() int32 {
	if x != nil {
		return x.GameStatus
	}
	return 0
}

type TakeActionMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr         string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	GameStatus   int32  `protobuf:"varint,2,opt,name=gameStatus,proto3" json:"gameStatus,omitempty"`
	PlayerAction int32  `protobuf:"varint,3,opt,name=playerAction,proto3" json:"playerAction,omitempty"`
	Value        int32  `protobuf:"varint,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TakeActionMsg) Reset() {
	*x = TakeActionMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeActionMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeActionMsg) ProtoMessage() {}

func (x *TakeActionMsg) ProtoReflect() protoreflect.Message {
	mi := &file_pb_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeActionMsg.ProtoReflect.Descriptor instead.
func (*TakeActionMsg) Descriptor() ([]byte, []int) {
	return file_pb_service_proto_rawDescGZIP(), []int{5}
}

func (x *TakeActionMsg) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *TakeActionMsg) GetGameStatus() int32 {
	if x != nil {
		return x.GameStatus
	}
	return 0
}

func (x *TakeActionMsg) GetPlayerAction() int32 {
	if x != nil {
		return x.PlayerAction
	}
	return 0
}

func (x *TakeActionMsg) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_pb_service_proto protoreflect.FileDescriptor

var file_pb_service_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x61,
	0x6d, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x67, 0x61, 0x6d, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x67, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x65, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x65, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x73, 0x22, 0x21,
	0x0a, 0x0b, 0x54, 0x61, 0x6b, 0x65, 0x53, 0x65, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x22, 0x05, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x22, 0x3e, 0x0a, 0x14, 0x53, 0x68, 0x75, 0x66,
	0x66, 0x6c, 0x65, 0x41, 0x6e, 0x64, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x4d, 0x73, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x65, 0x63, 0x6b, 0x22, 0x32, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x0a,
	0x67, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x7d, 0x0a, 0x0d,
	0x54, 0x61, 0x6b, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xce, 0x01, 0x0a, 0x06,
	0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x12, 0x24, 0x0a, 0x0a, 0x53, 0x68, 0x61, 0x6b, 0x65, 0x48,
	0x61, 0x6e, 0x64, 0x73, 0x12, 0x0a, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65,
	0x1a, 0x0a, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x1e, 0x0a, 0x08,
	0x54, 0x61, 0x6b, 0x65, 0x53, 0x65, 0x61, 0x74, 0x12, 0x0c, 0x2e, 0x54, 0x61, 0x6b, 0x65, 0x53,
	0x65, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x30, 0x0a, 0x11,
	0x53, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x41, 0x6e, 0x64, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x12, 0x15, 0x2e, 0x53, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x41, 0x6e, 0x64, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x28,
	0x0a, 0x0d, 0x53, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x11, 0x2e, 0x53, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x73, 0x67, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x22, 0x0a, 0x0a, 0x54, 0x61, 0x6b, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x2e, 0x54, 0x61, 0x6b, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x67, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x42, 0x05, 0x5a, 0x03,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_service_proto_rawDescOnce sync.Once
	file_pb_service_proto_rawDescData = file_pb_service_proto_rawDesc
)

func file_pb_service_proto_rawDescGZIP() []byte {
	file_pb_service_proto_rawDescOnce.Do(func() {
		file_pb_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_service_proto_rawDescData)
	})
	return file_pb_service_proto_rawDescData
}

var file_pb_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_service_proto_goTypes = []interface{}{
	(*Handshake)(nil),            // 0: Handshake
	(*TakeSeatMsg)(nil),          // 1: TakeSeatMsg
	(*Ack)(nil),                  // 2: Ack
	(*ShuffleAndEncryptMsg)(nil), // 3: ShuffleAndEncryptMsg
	(*SetGameStatusMsg)(nil),     // 4: SetGameStatusMsg
	(*TakeActionMsg)(nil),        // 5: TakeActionMsg
}
var file_pb_service_proto_depIdxs = []int32{
	0, // 0: Gossip.ShakeHands:input_type -> Handshake
	1, // 1: Gossip.TakeSeat:input_type -> TakeSeatMsg
	3, // 2: Gossip.ShuffleAndEncrypt:input_type -> ShuffleAndEncryptMsg
	4, // 3: Gossip.SetGameStatus:input_type -> SetGameStatusMsg
	5, // 4: Gossip.TakeAction:input_type -> TakeActionMsg
	0, // 5: Gossip.ShakeHands:output_type -> Handshake
	2, // 6: Gossip.TakeSeat:output_type -> Ack
	2, // 7: Gossip.ShuffleAndEncrypt:output_type -> Ack
	2, // 8: Gossip.SetGameStatus:output_type -> Ack
	2, // 9: Gossip.TakeAction:output_type -> Ack
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_service_proto_init() }
func file_pb_service_proto_init() {
	if File_pb_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Handshake); i {
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
		file_pb_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeSeatMsg); i {
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
		file_pb_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
		file_pb_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShuffleAndEncryptMsg); i {
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
		file_pb_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetGameStatusMsg); i {
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
		file_pb_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeActionMsg); i {
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
			RawDescriptor: file_pb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_service_proto_goTypes,
		DependencyIndexes: file_pb_service_proto_depIdxs,
		MessageInfos:      file_pb_service_proto_msgTypes,
	}.Build()
	File_pb_service_proto = out.File
	file_pb_service_proto_rawDesc = nil
	file_pb_service_proto_goTypes = nil
	file_pb_service_proto_depIdxs = nil
}
