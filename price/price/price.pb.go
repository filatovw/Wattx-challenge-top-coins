// Code generated by protoc-gen-go. DO NOT EDIT.
// source: price/price/price.proto

package price

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

type GetPricesRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPricesRequest) Reset()         { *m = GetPricesRequest{} }
func (m *GetPricesRequest) String() string { return proto.CompactTextString(m) }
func (*GetPricesRequest) ProtoMessage()    {}
func (*GetPricesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d539350e99067cbd, []int{0}
}

func (m *GetPricesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPricesRequest.Unmarshal(m, b)
}
func (m *GetPricesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPricesRequest.Marshal(b, m, deterministic)
}
func (m *GetPricesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPricesRequest.Merge(m, src)
}
func (m *GetPricesRequest) XXX_Size() int {
	return xxx_messageInfo_GetPricesRequest.Size(m)
}
func (m *GetPricesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPricesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPricesRequest proto.InternalMessageInfo

func (m *GetPricesRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetPricesResponse struct {
	Prices               []*Price `protobuf:"bytes,1,rep,name=Prices,proto3" json:"Prices,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPricesResponse) Reset()         { *m = GetPricesResponse{} }
func (m *GetPricesResponse) String() string { return proto.CompactTextString(m) }
func (*GetPricesResponse) ProtoMessage()    {}
func (*GetPricesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d539350e99067cbd, []int{1}
}

func (m *GetPricesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPricesResponse.Unmarshal(m, b)
}
func (m *GetPricesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPricesResponse.Marshal(b, m, deterministic)
}
func (m *GetPricesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPricesResponse.Merge(m, src)
}
func (m *GetPricesResponse) XXX_Size() int {
	return xxx_messageInfo_GetPricesResponse.Size(m)
}
func (m *GetPricesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPricesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPricesResponse proto.InternalMessageInfo

func (m *GetPricesResponse) GetPrices() []*Price {
	if m != nil {
		return m.Prices
	}
	return nil
}

type Price struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	Value                float64  `protobuf:"fixed64,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Price) Reset()         { *m = Price{} }
func (m *Price) String() string { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()    {}
func (*Price) Descriptor() ([]byte, []int) {
	return fileDescriptor_d539350e99067cbd, []int{2}
}

func (m *Price) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Price.Unmarshal(m, b)
}
func (m *Price) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Price.Marshal(b, m, deterministic)
}
func (m *Price) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Price.Merge(m, src)
}
func (m *Price) XXX_Size() int {
	return xxx_messageInfo_Price.Size(m)
}
func (m *Price) XXX_DiscardUnknown() {
	xxx_messageInfo_Price.DiscardUnknown(m)
}

var xxx_messageInfo_Price proto.InternalMessageInfo

func (m *Price) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Price) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*GetPricesRequest)(nil), "price.GetPricesRequest")
	proto.RegisterType((*GetPricesResponse)(nil), "price.GetPricesResponse")
	proto.RegisterType((*Price)(nil), "price.Price")
}

func init() { proto.RegisterFile("price/price/price.proto", fileDescriptor_d539350e99067cbd) }

var fileDescriptor_d539350e99067cbd = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x4c,
	0x4e, 0xd5, 0x47, 0x22, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x58, 0xc1, 0x1c, 0x25, 0x0d,
	0x2e, 0x01, 0xf7, 0xd4, 0x92, 0x00, 0x10, 0xbb, 0x38, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44,
	0x48, 0x84, 0x8b, 0x35, 0x27, 0x33, 0x37, 0xb3, 0x44, 0x82, 0x51, 0x81, 0x51, 0x83, 0x35, 0x08,
	0xc2, 0x51, 0xb2, 0xe4, 0x12, 0x44, 0x52, 0x59, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0xa4, 0xc2,
	0xc5, 0x06, 0x11, 0x91, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0xe2, 0xd1, 0x83, 0xd8, 0x01, 0x16,
	0x0c, 0x82, 0xca, 0x29, 0x99, 0x72, 0xb1, 0x82, 0x59, 0x42, 0x62, 0x5c, 0x6c, 0xc1, 0x95, 0xb9,
	0x49, 0xf9, 0x39, 0x60, 0xa3, 0x39, 0x83, 0xa0, 0x3c, 0x90, 0x8d, 0x61, 0x89, 0x39, 0xa5, 0xa9,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x8c, 0x41, 0x10, 0x8e, 0x51, 0x00, 0x17, 0x0f, 0x58, 0x5b, 0x70,
	0x6a, 0x51, 0x19, 0x48, 0xb7, 0x03, 0x17, 0x27, 0xdc, 0x05, 0x42, 0xe2, 0x50, 0x9b, 0xd0, 0x5d,
	0x2f, 0x25, 0x81, 0x29, 0x01, 0x71, 0xac, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0xef, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x16, 0x8a, 0x7c, 0x75, 0x16, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PriceServiceClient is the client API for PriceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PriceServiceClient interface {
	GetPrices(ctx context.Context, in *GetPricesRequest, opts ...grpc.CallOption) (*GetPricesResponse, error)
}

type priceServiceClient struct {
	cc *grpc.ClientConn
}

func NewPriceServiceClient(cc *grpc.ClientConn) PriceServiceClient {
	return &priceServiceClient{cc}
}

func (c *priceServiceClient) GetPrices(ctx context.Context, in *GetPricesRequest, opts ...grpc.CallOption) (*GetPricesResponse, error) {
	out := new(GetPricesResponse)
	err := c.cc.Invoke(ctx, "/price.PriceService/GetPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PriceServiceServer is the server API for PriceService service.
type PriceServiceServer interface {
	GetPrices(context.Context, *GetPricesRequest) (*GetPricesResponse, error)
}

func RegisterPriceServiceServer(s *grpc.Server, srv PriceServiceServer) {
	s.RegisterService(&_PriceService_serviceDesc, srv)
}

func _PriceService_GetPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceServiceServer).GetPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/price.PriceService/GetPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceServiceServer).GetPrices(ctx, req.(*GetPricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PriceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "price.PriceService",
	HandlerType: (*PriceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrices",
			Handler:    _PriceService_GetPrices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "price/price/price.proto",
}
