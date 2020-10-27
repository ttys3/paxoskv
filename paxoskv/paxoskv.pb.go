// Code generated by protoc-gen-go. DO NOT EDIT.
// source: paxoskv.proto

package paxoskv

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// BallotNum is the ballot number in paxos. It consists of a monotonically
// incremental number and a universally unique ProposerId.
type BallotNum struct {
	N                    int64    `protobuf:"varint,1,opt,name=N,proto3" json:"N,omitempty"`
	ProposerId           int64    `protobuf:"varint,2,opt,name=ProposerId,proto3" json:"ProposerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BallotNum) Reset()         { *m = BallotNum{} }
func (m *BallotNum) String() string { return proto.CompactTextString(m) }
func (*BallotNum) ProtoMessage()    {}
func (*BallotNum) Descriptor() ([]byte, []int) {
	return fileDescriptor_paxoskv_b9f0d51d1f406d18, []int{0}
}
func (m *BallotNum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BallotNum.Unmarshal(m, b)
}
func (m *BallotNum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BallotNum.Marshal(b, m, deterministic)
}
func (dst *BallotNum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BallotNum.Merge(dst, src)
}
func (m *BallotNum) XXX_Size() int {
	return xxx_messageInfo_BallotNum.Size(m)
}
func (m *BallotNum) XXX_DiscardUnknown() {
	xxx_messageInfo_BallotNum.DiscardUnknown(m)
}

var xxx_messageInfo_BallotNum proto.InternalMessageInfo

func (m *BallotNum) GetN() int64 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *BallotNum) GetProposerId() int64 {
	if m != nil {
		return m.ProposerId
	}
	return 0
}

// Value is the value part of a key-value record.
// In this demo it is just a int64
type Value struct {
	Vi64                 int64    `protobuf:"varint,1,opt,name=Vi64,proto3" json:"Vi64,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_paxoskv_b9f0d51d1f406d18, []int{1}
}
func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (dst *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(dst, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetVi64() int64 {
	if m != nil {
		return m.Vi64
	}
	return 0
}

// PaxosInstanceId specifies what paxos instance it runs on.
// A paxos instance is used to determine a specific version of a record.
// E.g.: for a key-value record foo₀=0, to set foo=2, a paxos instance is
// created to choose the value for key "foo", ver "1", i.e., foo₁
type PaxosInstanceId struct {
	// the key of a record to operate on.
	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	// the version of the record to modify.
	Ver                  int64    `protobuf:"varint,2,opt,name=Ver,proto3" json:"Ver,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaxosInstanceId) Reset()         { *m = PaxosInstanceId{} }
func (m *PaxosInstanceId) String() string { return proto.CompactTextString(m) }
func (*PaxosInstanceId) ProtoMessage()    {}
func (*PaxosInstanceId) Descriptor() ([]byte, []int) {
	return fileDescriptor_paxoskv_b9f0d51d1f406d18, []int{2}
}
func (m *PaxosInstanceId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaxosInstanceId.Unmarshal(m, b)
}
func (m *PaxosInstanceId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaxosInstanceId.Marshal(b, m, deterministic)
}
func (dst *PaxosInstanceId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaxosInstanceId.Merge(dst, src)
}
func (m *PaxosInstanceId) XXX_Size() int {
	return xxx_messageInfo_PaxosInstanceId.Size(m)
}
func (m *PaxosInstanceId) XXX_DiscardUnknown() {
	xxx_messageInfo_PaxosInstanceId.DiscardUnknown(m)
}

var xxx_messageInfo_PaxosInstanceId proto.InternalMessageInfo

func (m *PaxosInstanceId) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PaxosInstanceId) GetVer() int64 {
	if m != nil {
		return m.Ver
	}
	return 0
}

// Acceptor is the state of an Acceptor and also serves as the reply of the
// Prepare/Accept.
type Acceptor struct {
	// the last ballot number the instance knows of.
	LastBal *BallotNum `protobuf:"bytes,1,opt,name=LastBal,proto3" json:"LastBal,omitempty"`
	// the voted value by this Acceptor
	Val *Value `protobuf:"bytes,2,opt,name=Val,proto3" json:"Val,omitempty"`
	// at which ballot number the Acceptor voted it.
	VBal                 *BallotNum `protobuf:"bytes,3,opt,name=VBal,proto3" json:"VBal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Acceptor) Reset()         { *m = Acceptor{} }
func (m *Acceptor) String() string { return proto.CompactTextString(m) }
func (*Acceptor) ProtoMessage()    {}
func (*Acceptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_paxoskv_b9f0d51d1f406d18, []int{3}
}
func (m *Acceptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Acceptor.Unmarshal(m, b)
}
func (m *Acceptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Acceptor.Marshal(b, m, deterministic)
}
func (dst *Acceptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Acceptor.Merge(dst, src)
}
func (m *Acceptor) XXX_Size() int {
	return xxx_messageInfo_Acceptor.Size(m)
}
func (m *Acceptor) XXX_DiscardUnknown() {
	xxx_messageInfo_Acceptor.DiscardUnknown(m)
}

var xxx_messageInfo_Acceptor proto.InternalMessageInfo

func (m *Acceptor) GetLastBal() *BallotNum {
	if m != nil {
		return m.LastBal
	}
	return nil
}

func (m *Acceptor) GetVal() *Value {
	if m != nil {
		return m.Val
	}
	return nil
}

func (m *Acceptor) GetVBal() *BallotNum {
	if m != nil {
		return m.VBal
	}
	return nil
}

// Proposer is the state of a Proposer and also serves as the request of
// Prepare/Accept.
type Proposer struct {
	// what paxos instance it runs on
	Id *PaxosInstanceId `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// Bal is the ballot number of a Proposer
	Bal *BallotNum `protobuf:"bytes,2,opt,name=Bal,proto3" json:"Bal,omitempty"`
	// Val is the value a Proposer has chosen.
	Val                  *Value   `protobuf:"bytes,3,opt,name=Val,proto3" json:"Val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Proposer) Reset()         { *m = Proposer{} }
func (m *Proposer) String() string { return proto.CompactTextString(m) }
func (*Proposer) ProtoMessage()    {}
func (*Proposer) Descriptor() ([]byte, []int) {
	return fileDescriptor_paxoskv_b9f0d51d1f406d18, []int{4}
}
func (m *Proposer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proposer.Unmarshal(m, b)
}
func (m *Proposer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proposer.Marshal(b, m, deterministic)
}
func (dst *Proposer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposer.Merge(dst, src)
}
func (m *Proposer) XXX_Size() int {
	return xxx_messageInfo_Proposer.Size(m)
}
func (m *Proposer) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposer.DiscardUnknown(m)
}

var xxx_messageInfo_Proposer proto.InternalMessageInfo

func (m *Proposer) GetId() *PaxosInstanceId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Proposer) GetBal() *BallotNum {
	if m != nil {
		return m.Bal
	}
	return nil
}

func (m *Proposer) GetVal() *Value {
	if m != nil {
		return m.Val
	}
	return nil
}

func init() {
	proto.RegisterType((*BallotNum)(nil), "paxoskv.BallotNum")
	proto.RegisterType((*Value)(nil), "paxoskv.Value")
	proto.RegisterType((*PaxosInstanceId)(nil), "paxoskv.PaxosInstanceId")
	proto.RegisterType((*Acceptor)(nil), "paxoskv.Acceptor")
	proto.RegisterType((*Proposer)(nil), "paxoskv.Proposer")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PaxosKVClient is the client API for PaxosKV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaxosKVClient interface {
	Prepare(ctx context.Context, in *Proposer, opts ...grpc.CallOption) (*Acceptor, error)
	Accept(ctx context.Context, in *Proposer, opts ...grpc.CallOption) (*Acceptor, error)
}

type paxosKVClient struct {
	cc *grpc.ClientConn
}

func NewPaxosKVClient(cc *grpc.ClientConn) PaxosKVClient {
	return &paxosKVClient{cc}
}

func (c *paxosKVClient) Prepare(ctx context.Context, in *Proposer, opts ...grpc.CallOption) (*Acceptor, error) {
	out := new(Acceptor)
	err := c.cc.Invoke(ctx, "/paxoskv.PaxosKV/Prepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paxosKVClient) Accept(ctx context.Context, in *Proposer, opts ...grpc.CallOption) (*Acceptor, error) {
	out := new(Acceptor)
	err := c.cc.Invoke(ctx, "/paxoskv.PaxosKV/Accept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaxosKVServer is the server API for PaxosKV service.
type PaxosKVServer interface {
	Prepare(context.Context, *Proposer) (*Acceptor, error)
	Accept(context.Context, *Proposer) (*Acceptor, error)
}

func RegisterPaxosKVServer(s *grpc.Server, srv PaxosKVServer) {
	s.RegisterService(&_PaxosKV_serviceDesc, srv)
}

func _PaxosKV_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Proposer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosKVServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paxoskv.PaxosKV/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosKVServer).Prepare(ctx, req.(*Proposer))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaxosKV_Accept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Proposer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosKVServer).Accept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paxoskv.PaxosKV/Accept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosKVServer).Accept(ctx, req.(*Proposer))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaxosKV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "paxoskv.PaxosKV",
	HandlerType: (*PaxosKVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prepare",
			Handler:    _PaxosKV_Prepare_Handler,
		},
		{
			MethodName: "Accept",
			Handler:    _PaxosKV_Accept_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "paxoskv.proto",
}

func init() { proto.RegisterFile("paxoskv.proto", fileDescriptor_paxoskv_b9f0d51d1f406d18) }

var fileDescriptor_paxoskv_b9f0d51d1f406d18 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4b, 0xc3, 0x30,
	0x14, 0xb7, 0xcd, 0x5c, 0xb7, 0x37, 0x3f, 0xdf, 0xa9, 0x28, 0xc8, 0x08, 0x22, 0x3b, 0xc8, 0xd0,
	0xfa, 0x01, 0x1e, 0xed, 0xad, 0x4c, 0x4a, 0xe9, 0x21, 0xf7, 0xd8, 0xe6, 0x20, 0xd6, 0xa6, 0xa4,
	0x99, 0x28, 0x78, 0xf2, 0x2f, 0x97, 0xc4, 0xb4, 0x1b, 0xe2, 0x60, 0xb7, 0xf7, 0x7e, 0xef, 0xfd,
	0x3e, 0x5e, 0x08, 0xec, 0x37, 0xfc, 0x43, 0xb6, 0xaf, 0xef, 0xf3, 0x46, 0x49, 0x2d, 0x31, 0x70,
	0x2d, 0x7d, 0x80, 0x71, 0xcc, 0xab, 0x4a, 0xea, 0x74, 0xf9, 0x86, 0x7b, 0xe0, 0xa5, 0xa1, 0x37,
	0xf5, 0x66, 0x24, 0xf7, 0x52, 0x3c, 0x03, 0xc8, 0x94, 0x6c, 0x64, 0x2b, 0x54, 0x52, 0x86, 0xbe,
	0x85, 0xd7, 0x10, 0x7a, 0x0a, 0xbb, 0x8c, 0x57, 0x4b, 0x81, 0x08, 0x03, 0xf6, 0x72, 0x7f, 0xeb,
	0x98, 0xb6, 0xa6, 0x77, 0x70, 0x98, 0x19, 0x8b, 0xa4, 0x6e, 0x35, 0xaf, 0x0b, 0x91, 0x94, 0x78,
	0x04, 0x64, 0x21, 0x3e, 0xed, 0xd6, 0x38, 0x37, 0xa5, 0x41, 0x98, 0x50, 0x4e, 0xda, 0x94, 0xf4,
	0xdb, 0x83, 0xd1, 0x63, 0x51, 0x88, 0x46, 0x4b, 0x85, 0x97, 0x10, 0x3c, 0xf1, 0x56, 0xc7, 0xbc,
	0xb2, 0xa4, 0x49, 0x84, 0xf3, 0xee, 0x8a, 0x3e, 0x73, 0xde, 0xad, 0xe0, 0x14, 0x08, 0xe3, 0x95,
	0x15, 0x9b, 0x44, 0x07, 0xfd, 0xa6, 0x8d, 0x98, 0x9b, 0x11, 0x5e, 0xc0, 0x80, 0x19, 0x31, 0xb2,
	0x51, 0xcc, 0xce, 0xe9, 0x17, 0x8c, 0xba, 0x33, 0x71, 0x06, 0x7e, 0x52, 0x3a, 0xfb, 0xb0, 0x67,
	0xfc, 0x39, 0x2d, 0xf7, 0x93, 0x12, 0xcf, 0x81, 0xc4, 0xbd, 0xff, 0x7f, 0xe2, 0x64, 0x2d, 0x25,
	0xd9, 0x98, 0x32, 0xaa, 0x21, 0xb0, 0xf2, 0x0b, 0x86, 0xd7, 0x10, 0x64, 0x4a, 0x34, 0x5c, 0x09,
	0x3c, 0x5e, 0x79, 0xbb, 0x68, 0x27, 0x2b, 0xa8, 0x7b, 0x31, 0xba, 0x83, 0x57, 0x30, 0xfc, 0xed,
	0xb6, 0x65, 0x3c, 0x0f, 0xed, 0x8f, 0xb8, 0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0x41, 0x91, 0xc4,
	0xa0, 0x22, 0x02, 0x00, 0x00,
}
