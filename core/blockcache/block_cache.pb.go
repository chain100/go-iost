// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/blockcache/block_cache.proto

package blockcache

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BcMessageType int32

const (
	BcMessageType_LinkType                    BcMessageType = 0
	BcMessageType_UpdateActiveType            BcMessageType = 1
	BcMessageType_UpdateLinkedRootWitnessType BcMessageType = 2
)

var BcMessageType_name = map[int32]string{
	0: "LinkType",
	1: "UpdateActiveType",
	2: "UpdateLinkedRootWitnessType",
}

var BcMessageType_value = map[string]int32{
	"LinkType":                    0,
	"UpdateActiveType":            1,
	"UpdateLinkedRootWitnessType": 2,
}

func (x BcMessageType) String() string {
	return proto.EnumName(BcMessageType_name, int32(x))
}

func (BcMessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96d071d2d2082821, []int{0}
}

type BcMessage struct {
	Data                 []byte        `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Type                 BcMessageType `protobuf:"varint,2,opt,name=type,proto3,enum=blockcache.BcMessageType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BcMessage) Reset()         { *m = BcMessage{} }
func (m *BcMessage) String() string { return proto.CompactTextString(m) }
func (*BcMessage) ProtoMessage()    {}
func (*BcMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d071d2d2082821, []int{0}
}

func (m *BcMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BcMessage.Unmarshal(m, b)
}
func (m *BcMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BcMessage.Marshal(b, m, deterministic)
}
func (m *BcMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BcMessage.Merge(m, src)
}
func (m *BcMessage) XXX_Size() int {
	return xxx_messageInfo_BcMessage.Size(m)
}
func (m *BcMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BcMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BcMessage proto.InternalMessageInfo

func (m *BcMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BcMessage) GetType() BcMessageType {
	if m != nil {
		return m.Type
	}
	return BcMessageType_LinkType
}

type BlockCacheRaw struct {
	BlockBytes           []byte       `protobuf:"bytes,1,opt,name=blockBytes,proto3" json:"blockBytes,omitempty"`
	WitnessList          *WitnessList `protobuf:"bytes,2,opt,name=witnessList,proto3" json:"witnessList,omitempty"`
	SerialNum            int64        `protobuf:"varint,3,opt,name=serialNum,proto3" json:"serialNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BlockCacheRaw) Reset()         { *m = BlockCacheRaw{} }
func (m *BlockCacheRaw) String() string { return proto.CompactTextString(m) }
func (*BlockCacheRaw) ProtoMessage()    {}
func (*BlockCacheRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d071d2d2082821, []int{1}
}

func (m *BlockCacheRaw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockCacheRaw.Unmarshal(m, b)
}
func (m *BlockCacheRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockCacheRaw.Marshal(b, m, deterministic)
}
func (m *BlockCacheRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockCacheRaw.Merge(m, src)
}
func (m *BlockCacheRaw) XXX_Size() int {
	return xxx_messageInfo_BlockCacheRaw.Size(m)
}
func (m *BlockCacheRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockCacheRaw.DiscardUnknown(m)
}

var xxx_messageInfo_BlockCacheRaw proto.InternalMessageInfo

func (m *BlockCacheRaw) GetBlockBytes() []byte {
	if m != nil {
		return m.BlockBytes
	}
	return nil
}

func (m *BlockCacheRaw) GetWitnessList() *WitnessList {
	if m != nil {
		return m.WitnessList
	}
	return nil
}

func (m *BlockCacheRaw) GetSerialNum() int64 {
	if m != nil {
		return m.SerialNum
	}
	return 0
}

type UpdateActiveRaw struct {
	BlockHashBytes       []byte       `protobuf:"bytes,1,opt,name=blockHashBytes,proto3" json:"blockHashBytes,omitempty"`
	WitnessList          *WitnessList `protobuf:"bytes,2,opt,name=witnessList,proto3" json:"witnessList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateActiveRaw) Reset()         { *m = UpdateActiveRaw{} }
func (m *UpdateActiveRaw) String() string { return proto.CompactTextString(m) }
func (*UpdateActiveRaw) ProtoMessage()    {}
func (*UpdateActiveRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d071d2d2082821, []int{2}
}

func (m *UpdateActiveRaw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateActiveRaw.Unmarshal(m, b)
}
func (m *UpdateActiveRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateActiveRaw.Marshal(b, m, deterministic)
}
func (m *UpdateActiveRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateActiveRaw.Merge(m, src)
}
func (m *UpdateActiveRaw) XXX_Size() int {
	return xxx_messageInfo_UpdateActiveRaw.Size(m)
}
func (m *UpdateActiveRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateActiveRaw.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateActiveRaw proto.InternalMessageInfo

func (m *UpdateActiveRaw) GetBlockHashBytes() []byte {
	if m != nil {
		return m.BlockHashBytes
	}
	return nil
}

func (m *UpdateActiveRaw) GetWitnessList() *WitnessList {
	if m != nil {
		return m.WitnessList
	}
	return nil
}

type UpdateLinkedRootWitnessRaw struct {
	LinkedRootWitness    []string `protobuf:"bytes,1,rep,name=linkedRootWitness,proto3" json:"linkedRootWitness,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateLinkedRootWitnessRaw) Reset()         { *m = UpdateLinkedRootWitnessRaw{} }
func (m *UpdateLinkedRootWitnessRaw) String() string { return proto.CompactTextString(m) }
func (*UpdateLinkedRootWitnessRaw) ProtoMessage()    {}
func (*UpdateLinkedRootWitnessRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d071d2d2082821, []int{3}
}

func (m *UpdateLinkedRootWitnessRaw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateLinkedRootWitnessRaw.Unmarshal(m, b)
}
func (m *UpdateLinkedRootWitnessRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateLinkedRootWitnessRaw.Marshal(b, m, deterministic)
}
func (m *UpdateLinkedRootWitnessRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateLinkedRootWitnessRaw.Merge(m, src)
}
func (m *UpdateLinkedRootWitnessRaw) XXX_Size() int {
	return xxx_messageInfo_UpdateLinkedRootWitnessRaw.Size(m)
}
func (m *UpdateLinkedRootWitnessRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateLinkedRootWitnessRaw.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateLinkedRootWitnessRaw proto.InternalMessageInfo

func (m *UpdateLinkedRootWitnessRaw) GetLinkedRootWitness() []string {
	if m != nil {
		return m.LinkedRootWitness
	}
	return nil
}

type WitnessList struct {
	ActiveWitnessList    []string `protobuf:"bytes,1,rep,name=activeWitnessList,proto3" json:"activeWitnessList,omitempty"`
	PendingWitnessList   []string `protobuf:"bytes,2,rep,name=pendingWitnessList,proto3" json:"pendingWitnessList,omitempty"`
	PendingWitnessNumber int64    `protobuf:"varint,3,opt,name=pendingWitnessNumber,proto3" json:"pendingWitnessNumber,omitempty"`
	WitnessInfo          []string `protobuf:"bytes,4,rep,name=witnessInfo,proto3" json:"witnessInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WitnessList) Reset()         { *m = WitnessList{} }
func (m *WitnessList) String() string { return proto.CompactTextString(m) }
func (*WitnessList) ProtoMessage()    {}
func (*WitnessList) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d071d2d2082821, []int{4}
}

func (m *WitnessList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WitnessList.Unmarshal(m, b)
}
func (m *WitnessList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WitnessList.Marshal(b, m, deterministic)
}
func (m *WitnessList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WitnessList.Merge(m, src)
}
func (m *WitnessList) XXX_Size() int {
	return xxx_messageInfo_WitnessList.Size(m)
}
func (m *WitnessList) XXX_DiscardUnknown() {
	xxx_messageInfo_WitnessList.DiscardUnknown(m)
}

var xxx_messageInfo_WitnessList proto.InternalMessageInfo

func (m *WitnessList) GetActiveWitnessList() []string {
	if m != nil {
		return m.ActiveWitnessList
	}
	return nil
}

func (m *WitnessList) GetPendingWitnessList() []string {
	if m != nil {
		return m.PendingWitnessList
	}
	return nil
}

func (m *WitnessList) GetPendingWitnessNumber() int64 {
	if m != nil {
		return m.PendingWitnessNumber
	}
	return 0
}

func (m *WitnessList) GetWitnessInfo() []string {
	if m != nil {
		return m.WitnessInfo
	}
	return nil
}

type WitnessInfo struct {
	NetID                string   `protobuf:"bytes,1,opt,name=NetID,proto3" json:"NetID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WitnessInfo) Reset()         { *m = WitnessInfo{} }
func (m *WitnessInfo) String() string { return proto.CompactTextString(m) }
func (*WitnessInfo) ProtoMessage()    {}
func (*WitnessInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d071d2d2082821, []int{5}
}

func (m *WitnessInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WitnessInfo.Unmarshal(m, b)
}
func (m *WitnessInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WitnessInfo.Marshal(b, m, deterministic)
}
func (m *WitnessInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WitnessInfo.Merge(m, src)
}
func (m *WitnessInfo) XXX_Size() int {
	return xxx_messageInfo_WitnessInfo.Size(m)
}
func (m *WitnessInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_WitnessInfo.DiscardUnknown(m)
}

var xxx_messageInfo_WitnessInfo proto.InternalMessageInfo

func (m *WitnessInfo) GetNetID() string {
	if m != nil {
		return m.NetID
	}
	return ""
}

func init() {
	proto.RegisterEnum("blockcache.BcMessageType", BcMessageType_name, BcMessageType_value)
	proto.RegisterType((*BcMessage)(nil), "blockcache.bcMessage")
	proto.RegisterType((*BlockCacheRaw)(nil), "blockcache.BlockCacheRaw")
	proto.RegisterType((*UpdateActiveRaw)(nil), "blockcache.UpdateActiveRaw")
	proto.RegisterType((*UpdateLinkedRootWitnessRaw)(nil), "blockcache.UpdateLinkedRootWitnessRaw")
	proto.RegisterType((*WitnessList)(nil), "blockcache.WitnessList")
	proto.RegisterType((*WitnessInfo)(nil), "blockcache.WitnessInfo")
}

func init() { proto.RegisterFile("core/blockcache/block_cache.proto", fileDescriptor_96d071d2d2082821) }

var fileDescriptor_96d071d2d2082821 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0x4f, 0xe3, 0x30,
	0x10, 0x5d, 0xb7, 0xdd, 0xd5, 0x66, 0xfa, 0xb1, 0x5d, 0xab, 0xd2, 0x66, 0x01, 0x41, 0x08, 0x12,
	0x8a, 0x10, 0x04, 0xa9, 0x9c, 0x38, 0x52, 0x38, 0x50, 0x54, 0x72, 0xb0, 0x8a, 0x7a, 0x44, 0x4e,
	0x62, 0xda, 0xa8, 0x6d, 0x12, 0xc5, 0x2e, 0x55, 0xff, 0x01, 0x7f, 0x8b, 0x7f, 0x86, 0xec, 0x86,
	0xc6, 0xfd, 0xe0, 0xc4, 0x6d, 0xf2, 0xde, 0x9b, 0x79, 0x33, 0x2f, 0x86, 0xe3, 0x20, 0xc9, 0xd8,
	0xa5, 0x3f, 0x49, 0x82, 0x71, 0x40, 0x83, 0x51, 0x5e, 0x3e, 0xab, 0xda, 0x4d, 0xb3, 0x44, 0x24,
	0x18, 0x0a, 0xd6, 0xf6, 0xc0, 0xf0, 0x83, 0x47, 0xc6, 0x39, 0x1d, 0x32, 0x8c, 0xa1, 0x12, 0x52,
	0x41, 0x4d, 0x64, 0x21, 0xa7, 0x46, 0x54, 0x8d, 0x2f, 0xa0, 0x22, 0x16, 0x29, 0x33, 0x4b, 0x16,
	0x72, 0x1a, 0xed, 0xff, 0x6e, 0xd1, 0xeb, 0xae, 0x1a, 0xfb, 0x8b, 0x94, 0x11, 0x25, 0xb3, 0xdf,
	0x10, 0xd4, 0x3b, 0x52, 0x72, 0x2b, 0x25, 0x84, 0xce, 0xf1, 0x21, 0x2c, 0xfd, 0x3a, 0x0b, 0xc1,
	0x78, 0x3e, 0x5a, 0x43, 0xf0, 0x35, 0x54, 0xe7, 0x91, 0x88, 0x19, 0xe7, 0xbd, 0x88, 0x0b, 0xe5,
	0x53, 0x6d, 0xff, 0xd3, 0x7d, 0x06, 0x05, 0x4d, 0x74, 0x2d, 0x3e, 0x00, 0x83, 0xb3, 0x2c, 0xa2,
	0x13, 0x6f, 0x36, 0x35, 0xcb, 0x16, 0x72, 0xca, 0xa4, 0x00, 0x6c, 0x01, 0x7f, 0x9e, 0xd2, 0x90,
	0x0a, 0x76, 0x13, 0x88, 0xe8, 0x55, 0xed, 0x72, 0x0a, 0x0d, 0x35, 0xf7, 0x9e, 0xf2, 0x91, 0xbe,
	0xcf, 0x06, 0xfa, 0x8d, 0x9d, 0xec, 0x07, 0xd8, 0x5b, 0xba, 0xf6, 0xa2, 0x78, 0xcc, 0x42, 0x92,
	0x24, 0x22, 0xd7, 0xca, 0x05, 0xce, 0xe1, 0xef, 0x64, 0x13, 0x37, 0x91, 0x55, 0x76, 0x0c, 0xb2,
	0x4d, 0xd8, 0xef, 0x08, 0xaa, 0x9a, 0x91, 0xec, 0xa6, 0xea, 0x16, 0x0d, 0xfc, 0xec, 0xde, 0x22,
	0xb0, 0x0b, 0x38, 0x65, 0x71, 0x18, 0xc5, 0xc3, 0xc1, 0xda, 0x2d, 0x52, 0xbe, 0x83, 0xc1, 0x6d,
	0x68, 0xad, 0xa3, 0xde, 0x6c, 0xea, 0xb3, 0x2c, 0x0f, 0x76, 0x27, 0x87, 0xad, 0x55, 0x50, 0xdd,
	0xf8, 0x25, 0x31, 0x2b, 0x6a, 0xb8, 0x0e, 0xd9, 0x27, 0xab, 0x13, 0xe4, 0x27, 0x6e, 0xc1, 0x4f,
	0x8f, 0x89, 0xee, 0x9d, 0x0a, 0xde, 0x20, 0xcb, 0x8f, 0xb3, 0x3e, 0xd4, 0xd7, 0x1e, 0x13, 0xae,
	0xc1, 0x6f, 0x99, 0x9f, 0xac, 0x9b, 0x3f, 0x70, 0x0b, 0x9a, 0xfa, 0x9f, 0x54, 0x28, 0xc2, 0x47,
	0xb0, 0xff, 0x45, 0xd2, 0x4a, 0x50, 0xf2, 0x7f, 0xa9, 0xe7, 0x7e, 0xf5, 0x11, 0x00, 0x00, 0xff,
	0xff, 0x00, 0xae, 0xa2, 0x06, 0x13, 0x03, 0x00, 0x00,
}
