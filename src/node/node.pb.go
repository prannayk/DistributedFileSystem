// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node/node.proto

package ChordStructure

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Node struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{0}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetIDRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIDRequest) Reset()         { *m = GetIDRequest{} }
func (m *GetIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetIDRequest) ProtoMessage()    {}
func (*GetIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{1}
}

func (m *GetIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIDRequest.Unmarshal(m, b)
}
func (m *GetIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIDRequest.Marshal(b, m, deterministic)
}
func (m *GetIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIDRequest.Merge(m, src)
}
func (m *GetIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetIDRequest.Size(m)
}
func (m *GetIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIDRequest proto.InternalMessageInfo

type GetIDResponse struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIDResponse) Reset()         { *m = GetIDResponse{} }
func (m *GetIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetIDResponse) ProtoMessage()    {}
func (*GetIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{2}
}

func (m *GetIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIDResponse.Unmarshal(m, b)
}
func (m *GetIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIDResponse.Marshal(b, m, deterministic)
}
func (m *GetIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIDResponse.Merge(m, src)
}
func (m *GetIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetIDResponse.Size(m)
}
func (m *GetIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetIDResponse proto.InternalMessageInfo

func (m *GetIDResponse) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type LocateRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocateRequest) Reset()         { *m = LocateRequest{} }
func (m *LocateRequest) String() string { return proto.CompactTextString(m) }
func (*LocateRequest) ProtoMessage()    {}
func (*LocateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{3}
}

func (m *LocateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocateRequest.Unmarshal(m, b)
}
func (m *LocateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocateRequest.Marshal(b, m, deterministic)
}
func (m *LocateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocateRequest.Merge(m, src)
}
func (m *LocateRequest) XXX_Size() int {
	return xxx_messageInfo_LocateRequest.Size(m)
}
func (m *LocateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LocateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LocateRequest proto.InternalMessageInfo

func (m *LocateRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type LocateResponse struct {
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocateResponse) Reset()         { *m = LocateResponse{} }
func (m *LocateResponse) String() string { return proto.CompactTextString(m) }
func (*LocateResponse) ProtoMessage()    {}
func (*LocateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{4}
}

func (m *LocateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocateResponse.Unmarshal(m, b)
}
func (m *LocateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocateResponse.Marshal(b, m, deterministic)
}
func (m *LocateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocateResponse.Merge(m, src)
}
func (m *LocateResponse) XXX_Size() int {
	return xxx_messageInfo_LocateResponse.Size(m)
}
func (m *LocateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LocateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LocateResponse proto.InternalMessageInfo

func (m *LocateResponse) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type GetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{5}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetResponse struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{6}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type PutRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{7}
}

func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (m *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(m, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type PutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{8}
}

func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (m *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(m, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

type TransferKeysReq struct {
	FromId               []byte   `protobuf:"bytes,1,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	ToNode               *Node    `protobuf:"bytes,2,opt,name=to_node,json=toNode,proto3" json:"to_node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferKeysReq) Reset()         { *m = TransferKeysReq{} }
func (m *TransferKeysReq) String() string { return proto.CompactTextString(m) }
func (*TransferKeysReq) ProtoMessage()    {}
func (*TransferKeysReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{9}
}

func (m *TransferKeysReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferKeysReq.Unmarshal(m, b)
}
func (m *TransferKeysReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferKeysReq.Marshal(b, m, deterministic)
}
func (m *TransferKeysReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferKeysReq.Merge(m, src)
}
func (m *TransferKeysReq) XXX_Size() int {
	return xxx_messageInfo_TransferKeysReq.Size(m)
}
func (m *TransferKeysReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferKeysReq.DiscardUnknown(m)
}

var xxx_messageInfo_TransferKeysReq proto.InternalMessageInfo

func (m *TransferKeysReq) GetFromId() []byte {
	if m != nil {
		return m.FromId
	}
	return nil
}

func (m *TransferKeysReq) GetToNode() *Node {
	if m != nil {
		return m.ToNode
	}
	return nil
}

type MT struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MT) Reset()         { *m = MT{} }
func (m *MT) String() string { return proto.CompactTextString(m) }
func (*MT) ProtoMessage()    {}
func (*MT) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{10}
}

func (m *MT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MT.Unmarshal(m, b)
}
func (m *MT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MT.Marshal(b, m, deterministic)
}
func (m *MT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MT.Merge(m, src)
}
func (m *MT) XXX_Size() int {
	return xxx_messageInfo_MT.Size(m)
}
func (m *MT) XXX_DiscardUnknown() {
	xxx_messageInfo_MT.DiscardUnknown(m)
}

var xxx_messageInfo_MT proto.InternalMessageInfo

type KeyVal struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val                  []byte   `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyVal) Reset()         { *m = KeyVal{} }
func (m *KeyVal) String() string { return proto.CompactTextString(m) }
func (*KeyVal) ProtoMessage()    {}
func (*KeyVal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{11}
}

func (m *KeyVal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyVal.Unmarshal(m, b)
}
func (m *KeyVal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyVal.Marshal(b, m, deterministic)
}
func (m *KeyVal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyVal.Merge(m, src)
}
func (m *KeyVal) XXX_Size() int {
	return xxx_messageInfo_KeyVal.Size(m)
}
func (m *KeyVal) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyVal.DiscardUnknown(m)
}

var xxx_messageInfo_KeyVal proto.InternalMessageInfo

func (m *KeyVal) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyVal) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

type ID struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{12}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type Key struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{13}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Val struct {
	Val                  []byte   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Val) Reset()         { *m = Val{} }
func (m *Val) String() string { return proto.CompactTextString(m) }
func (*Val) ProtoMessage()    {}
func (*Val) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18530e439628818, []int{14}
}

func (m *Val) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Val.Unmarshal(m, b)
}
func (m *Val) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Val.Marshal(b, m, deterministic)
}
func (m *Val) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Val.Merge(m, src)
}
func (m *Val) XXX_Size() int {
	return xxx_messageInfo_Val.Size(m)
}
func (m *Val) XXX_DiscardUnknown() {
	xxx_messageInfo_Val.DiscardUnknown(m)
}

var xxx_messageInfo_Val proto.InternalMessageInfo

func (m *Val) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "ChordStructure.Node")
	proto.RegisterType((*GetIDRequest)(nil), "ChordStructure.GetIDRequest")
	proto.RegisterType((*GetIDResponse)(nil), "ChordStructure.GetIDResponse")
	proto.RegisterType((*LocateRequest)(nil), "ChordStructure.LocateRequest")
	proto.RegisterType((*LocateResponse)(nil), "ChordStructure.LocateResponse")
	proto.RegisterType((*GetRequest)(nil), "ChordStructure.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "ChordStructure.GetResponse")
	proto.RegisterType((*PutRequest)(nil), "ChordStructure.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "ChordStructure.PutResponse")
	proto.RegisterType((*TransferKeysReq)(nil), "ChordStructure.TransferKeysReq")
	proto.RegisterType((*MT)(nil), "ChordStructure.MT")
	proto.RegisterType((*KeyVal)(nil), "ChordStructure.KeyVal")
	proto.RegisterType((*ID)(nil), "ChordStructure.ID")
	proto.RegisterType((*Key)(nil), "ChordStructure.Key")
	proto.RegisterType((*Val)(nil), "ChordStructure.Val")
}

func init() { proto.RegisterFile("node/node.proto", fileDescriptor_a18530e439628818) }

var fileDescriptor_a18530e439628818 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x93, 0xcf, 0x4f, 0xc2, 0x30,
	0x14, 0xc7, 0xc3, 0x80, 0x21, 0x0f, 0x18, 0xa6, 0x21, 0x81, 0xa0, 0xa2, 0xd6, 0x0b, 0x31, 0x8a,
	0x09, 0x7a, 0xf2, 0xe0, 0x41, 0x97, 0x20, 0x41, 0x8d, 0x99, 0xc4, 0xc4, 0x13, 0x99, 0xac, 0x44,
	0x23, 0xac, 0xd8, 0x75, 0x26, 0xfc, 0x31, 0xfe, 0xaf, 0xb6, 0x5d, 0xf9, 0x39, 0xe6, 0x05, 0x5e,
	0xfb, 0xde, 0xfb, 0xf4, 0xdb, 0x7e, 0xdf, 0xa0, 0xec, 0x53, 0x8f, 0x5c, 0xc8, 0x9f, 0xd6, 0x94,
	0x51, 0x4e, 0x91, 0x75, 0xf7, 0x41, 0x99, 0xf7, 0xc2, 0x59, 0x38, 0xe4, 0x21, 0x23, 0xf8, 0x14,
	0x32, 0x4f, 0x22, 0x8b, 0x2c, 0x30, 0x3e, 0xbd, 0x5a, 0xea, 0x28, 0xd5, 0x2c, 0x3a, 0x22, 0x42,
	0x08, 0x32, 0xbe, 0x3b, 0x21, 0x35, 0x43, 0xec, 0xe4, 0x1d, 0x15, 0x63, 0x0b, 0x8a, 0x1d, 0xc2,
	0xbb, 0xb6, 0x43, 0xbe, 0x43, 0x12, 0x70, 0x7c, 0x08, 0x25, 0xbd, 0x0e, 0xa6, 0xd4, 0x0f, 0x62,
	0x10, 0x7c, 0x0c, 0xa5, 0x07, 0x3a, 0x74, 0x39, 0xd1, 0x1d, 0x68, 0x17, 0xd2, 0x5f, 0x64, 0xa6,
	0x2a, 0xf2, 0x8e, 0x0c, 0xf1, 0x35, 0x58, 0xf3, 0x12, 0x0d, 0x69, 0x8a, 0x93, 0x85, 0x22, 0x55,
	0x54, 0x68, 0x57, 0x5a, 0xeb, 0x82, 0x5b, 0x52, 0xad, 0xa3, 0x2a, 0x70, 0x03, 0x40, 0x9c, 0x9f,
	0xcc, 0x3e, 0x81, 0x82, 0xca, 0x6b, 0x70, 0x05, 0xb2, 0x3f, 0xee, 0x38, 0x24, 0x5a, 0x60, 0xb4,
	0xc0, 0x57, 0x00, 0xcf, 0x61, 0x32, 0x64, 0xd9, 0x65, 0xac, 0x76, 0x95, 0xa0, 0xa0, 0xba, 0x22,
	0x34, 0x7e, 0x83, 0x72, 0x9f, 0xb9, 0x7e, 0x30, 0x22, 0xac, 0x47, 0x66, 0x81, 0xa0, 0xa1, 0x2a,
	0xe4, 0x46, 0x8c, 0x4e, 0x06, 0x8b, 0x07, 0x31, 0xe5, 0xb2, 0xeb, 0xa1, 0x73, 0xc8, 0x71, 0x3a,
	0x50, 0x57, 0x34, 0xfe, 0xb9, 0xa2, 0xc9, 0xa9, 0xfc, 0xc7, 0x19, 0x30, 0x1e, 0xfb, 0xf8, 0x0c,
	0x4c, 0x01, 0x7e, 0x75, 0xc7, 0x5b, 0x14, 0x8a, 0x1d, 0x21, 0x4a, 0xeb, 0x93, 0x21, 0xae, 0x80,
	0xd1, 0xb5, 0x63, 0x6e, 0x54, 0x21, 0xdd, 0x8b, 0xca, 0x37, 0xde, 0x49, 0x24, 0x34, 0x59, 0x72,
	0x52, 0x0b, 0x4e, 0xfb, 0xd7, 0x80, 0x1d, 0xa5, 0xcd, 0xbe, 0xef, 0x23, 0x1b, 0xb2, 0xca, 0x6d,
	0xb4, 0xbf, 0xa9, 0x77, 0x75, 0x28, 0xea, 0x07, 0x09, 0x59, 0x6d, 0x42, 0x07, 0xcc, 0xc8, 0x6f,
	0x14, 0x2b, 0x5c, 0x1b, 0x95, 0x7a, 0x23, 0x29, 0xad, 0x41, 0xb7, 0x90, 0x13, 0x64, 0xdb, 0xe5,
	0x2e, 0xaa, 0x6f, 0x39, 0x72, 0x8e, 0xd9, 0xdb, 0x9a, 0xd3, 0x8c, 0x1b, 0x48, 0x0b, 0x17, 0xe3,
	0xfd, 0xcb, 0x81, 0x88, 0xf7, 0xaf, 0xd8, 0xfe, 0x6e, 0xaa, 0x6f, 0xea, 0xf2, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0xd0, 0xb5, 0x0a, 0xe1, 0x66, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChordDHTClient is the client API for ChordDHT service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChordDHTClient interface {
	// GetID returns the ID of the node
	GetID(ctx context.Context, in *GetIDRequest, opts ...grpc.CallOption) (*GetIDResponse, error)
	// Locate a key
	Locate(ctx context.Context, in *LocateRequest, opts ...grpc.CallOption) (*LocateResponse, error)
	// Get returns value in the ring
	GetData(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Put writes a key value pair to the Chord ring
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
}

type chordDHTClient struct {
	cc *grpc.ClientConn
}

func NewChordDHTClient(cc *grpc.ClientConn) ChordDHTClient {
	return &chordDHTClient{cc}
}

func (c *chordDHTClient) GetID(ctx context.Context, in *GetIDRequest, opts ...grpc.CallOption) (*GetIDResponse, error) {
	out := new(GetIDResponse)
	err := c.cc.Invoke(ctx, "/ChordStructure.ChordDHT/GetID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordDHTClient) Locate(ctx context.Context, in *LocateRequest, opts ...grpc.CallOption) (*LocateResponse, error) {
	out := new(LocateResponse)
	err := c.cc.Invoke(ctx, "/ChordStructure.ChordDHT/Locate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordDHTClient) GetData(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/ChordStructure.ChordDHT/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordDHTClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/ChordStructure.ChordDHT/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChordDHTServer is the server API for ChordDHT service.
type ChordDHTServer interface {
	// GetID returns the ID of the node
	GetID(context.Context, *GetIDRequest) (*GetIDResponse, error)
	// Locate a key
	Locate(context.Context, *LocateRequest) (*LocateResponse, error)
	// Get returns value in the ring
	GetData(context.Context, *GetRequest) (*GetResponse, error)
	// Put writes a key value pair to the Chord ring
	Put(context.Context, *PutRequest) (*PutResponse, error)
}

func RegisterChordDHTServer(s *grpc.Server, srv ChordDHTServer) {
	s.RegisterService(&_ChordDHT_serviceDesc, srv)
}

func _ChordDHT_GetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordDHTServer).GetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChordStructure.ChordDHT/GetID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordDHTServer).GetID(ctx, req.(*GetIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordDHT_Locate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordDHTServer).Locate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChordStructure.ChordDHT/Locate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordDHTServer).Locate(ctx, req.(*LocateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordDHT_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordDHTServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChordStructure.ChordDHT/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordDHTServer).GetData(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChordDHT_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordDHTServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChordStructure.ChordDHT/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordDHTServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChordDHT_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ChordStructure.ChordDHT",
	HandlerType: (*ChordDHTServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetID",
			Handler:    _ChordDHT_GetID_Handler,
		},
		{
			MethodName: "Locate",
			Handler:    _ChordDHT_Locate_Handler,
		},
		{
			MethodName: "GetData",
			Handler:    _ChordDHT_GetData_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _ChordDHT_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node/node.proto",
}
