// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pricelist/rpc/pricelist.proto

package pricelist

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetPricelistRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPricelistRequest) Reset()         { *m = GetPricelistRequest{} }
func (m *GetPricelistRequest) String() string { return proto.CompactTextString(m) }
func (*GetPricelistRequest) ProtoMessage()    {}
func (*GetPricelistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe030b0dfe126bc, []int{0}
}

func (m *GetPricelistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPricelistRequest.Unmarshal(m, b)
}
func (m *GetPricelistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPricelistRequest.Marshal(b, m, deterministic)
}
func (m *GetPricelistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPricelistRequest.Merge(m, src)
}
func (m *GetPricelistRequest) XXX_Size() int {
	return xxx_messageInfo_GetPricelistRequest.Size(m)
}
func (m *GetPricelistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPricelistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPricelistRequest proto.InternalMessageInfo

func (m *GetPricelistRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetPricelistResponse struct {
	Positions            []*Position `protobuf:"bytes,1,rep,name=Positions,proto3" json:"Positions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetPricelistResponse) Reset()         { *m = GetPricelistResponse{} }
func (m *GetPricelistResponse) String() string { return proto.CompactTextString(m) }
func (*GetPricelistResponse) ProtoMessage()    {}
func (*GetPricelistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe030b0dfe126bc, []int{1}
}

func (m *GetPricelistResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPricelistResponse.Unmarshal(m, b)
}
func (m *GetPricelistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPricelistResponse.Marshal(b, m, deterministic)
}
func (m *GetPricelistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPricelistResponse.Merge(m, src)
}
func (m *GetPricelistResponse) XXX_Size() int {
	return xxx_messageInfo_GetPricelistResponse.Size(m)
}
func (m *GetPricelistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPricelistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPricelistResponse proto.InternalMessageInfo

func (m *GetPricelistResponse) GetPositions() []*Position {
	if m != nil {
		return m.Positions
	}
	return nil
}

type Position struct {
	Rank                 int32    `protobuf:"varint,1,opt,name=Rank,proto3" json:"Rank,omitempty"`
	Symbol               string   `protobuf:"bytes,2,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	PriceUSD             float64  `protobuf:"fixed64,3,opt,name=PriceUSD,proto3" json:"PriceUSD,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe030b0dfe126bc, []int{2}
}

func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *Position) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Position) GetPriceUSD() float64 {
	if m != nil {
		return m.PriceUSD
	}
	return 0
}

func init() {
	proto.RegisterType((*GetPricelistRequest)(nil), "pricelist.GetPricelistRequest")
	proto.RegisterType((*GetPricelistResponse)(nil), "pricelist.GetPricelistResponse")
	proto.RegisterType((*Position)(nil), "pricelist.Position")
}

func init() { proto.RegisterFile("pricelist/rpc/pricelist.proto", fileDescriptor_abe030b0dfe126bc) }

var fileDescriptor_abe030b0dfe126bc = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0x8d, 0x73, 0x63, 0xbd, 0xfa, 0x20, 0x77, 0x43, 0xc2, 0x40, 0x0d, 0x79, 0x0a, 0x08,
	0x1b, 0xce, 0xaf, 0x20, 0x88, 0x6f, 0x35, 0xc5, 0x0f, 0x60, 0xcb, 0x7d, 0x08, 0xb6, 0x4d, 0x4c,
	0xa2, 0xe0, 0xb7, 0x17, 0xfa, 0x27, 0xad, 0xa0, 0x6f, 0xf9, 0xe5, 0x1c, 0x38, 0xe7, 0x5c, 0xb8,
	0x76, 0xde, 0x54, 0x54, 0x9b, 0x10, 0x0f, 0xde, 0x55, 0x87, 0x44, 0x7b, 0xe7, 0x6d, 0xb4, 0x98,
	0xa5, 0x0f, 0x79, 0x07, 0x9b, 0x27, 0x8a, 0xf9, 0xc8, 0x9a, 0x3e, 0x3e, 0x29, 0x44, 0xdc, 0xc2,
	0xb2, 0x36, 0x8d, 0x89, 0x9c, 0x09, 0xa6, 0x96, 0xba, 0x07, 0xf9, 0x0c, 0xdb, 0xdf, 0xe6, 0xe0,
	0x6c, 0x1b, 0x08, 0xef, 0x21, 0xcb, 0x6d, 0x30, 0xd1, 0xd8, 0x36, 0x70, 0x26, 0x16, 0xea, 0xfc,
	0xb8, 0xd9, 0x4f, 0xa1, 0xa3, 0xa6, 0x27, 0x97, 0xd4, 0xb0, 0x1e, 0x01, 0x11, 0xce, 0xf4, 0x5b,
	0xfb, 0x3e, 0x64, 0x75, 0x6f, 0xbc, 0x82, 0x55, 0xf1, 0xdd, 0x94, 0xb6, 0xe6, 0xa7, 0x82, 0xa9,
	0x4c, 0x0f, 0x84, 0x3b, 0x58, 0x77, 0xf9, 0xaf, 0xc5, 0x23, 0x5f, 0x08, 0xa6, 0x98, 0x4e, 0x7c,
	0x24, 0xb8, 0x4c, 0xdd, 0x0a, 0xf2, 0x5f, 0xa6, 0x22, 0x7c, 0x81, 0x8b, 0x79, 0x65, 0xbc, 0x99,
	0xf5, 0xfa, 0x63, 0xf8, 0xee, 0xf6, 0x5f, 0xbd, 0xdf, 0x2a, 0x4f, 0xca, 0x55, 0x77, 0xc4, 0x87,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0x49, 0xfb, 0xdb, 0x65, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PricelistServiceClient is the client API for PricelistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PricelistServiceClient interface {
	GetPricelist(ctx context.Context, in *GetPricelistRequest, opts ...grpc.CallOption) (*GetPricelistResponse, error)
}

type pricelistServiceClient struct {
	cc *grpc.ClientConn
}

func NewPricelistServiceClient(cc *grpc.ClientConn) PricelistServiceClient {
	return &pricelistServiceClient{cc}
}

func (c *pricelistServiceClient) GetPricelist(ctx context.Context, in *GetPricelistRequest, opts ...grpc.CallOption) (*GetPricelistResponse, error) {
	out := new(GetPricelistResponse)
	err := c.cc.Invoke(ctx, "/pricelist.PricelistService/GetPricelist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PricelistServiceServer is the server API for PricelistService service.
type PricelistServiceServer interface {
	GetPricelist(context.Context, *GetPricelistRequest) (*GetPricelistResponse, error)
}

func RegisterPricelistServiceServer(s *grpc.Server, srv PricelistServiceServer) {
	s.RegisterService(&_PricelistService_serviceDesc, srv)
}

func _PricelistService_GetPricelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPricelistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricelistServiceServer).GetPricelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pricelist.PricelistService/GetPricelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricelistServiceServer).GetPricelist(ctx, req.(*GetPricelistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PricelistService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pricelist.PricelistService",
	HandlerType: (*PricelistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPricelist",
			Handler:    _PricelistService_GetPricelist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pricelist/rpc/pricelist.proto",
}
