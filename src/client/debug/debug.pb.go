// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client/debug/debug.proto

package debug

import (
	context "context"
	fmt "fmt"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type DumpRequest struct {
	// Recursed is true if this request is a recursive call from another request.
	// Callers should leave it unset, it's used to prevent infinite loops of
	// recursive calls.
	Recursed             bool     `protobuf:"varint,1,opt,name=recursed,proto3" json:"recursed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DumpRequest) Reset()         { *m = DumpRequest{} }
func (m *DumpRequest) String() string { return proto.CompactTextString(m) }
func (*DumpRequest) ProtoMessage()    {}
func (*DumpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d15a320d0127c22, []int{0}
}
func (m *DumpRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DumpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DumpRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DumpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpRequest.Merge(m, src)
}
func (m *DumpRequest) XXX_Size() int {
	return m.Size()
}
func (m *DumpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DumpRequest proto.InternalMessageInfo

func (m *DumpRequest) GetRecursed() bool {
	if m != nil {
		return m.Recursed
	}
	return false
}

type ProfileRequest struct {
	Profile              string          `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Duration             *types.Duration `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ProfileRequest) Reset()         { *m = ProfileRequest{} }
func (m *ProfileRequest) String() string { return proto.CompactTextString(m) }
func (*ProfileRequest) ProtoMessage()    {}
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d15a320d0127c22, []int{1}
}
func (m *ProfileRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProfileRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileRequest.Merge(m, src)
}
func (m *ProfileRequest) XXX_Size() int {
	return m.Size()
}
func (m *ProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileRequest proto.InternalMessageInfo

func (m *ProfileRequest) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *ProfileRequest) GetDuration() *types.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

type BinaryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BinaryRequest) Reset()         { *m = BinaryRequest{} }
func (m *BinaryRequest) String() string { return proto.CompactTextString(m) }
func (*BinaryRequest) ProtoMessage()    {}
func (*BinaryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d15a320d0127c22, []int{2}
}
func (m *BinaryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BinaryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BinaryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BinaryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BinaryRequest.Merge(m, src)
}
func (m *BinaryRequest) XXX_Size() int {
	return m.Size()
}
func (m *BinaryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BinaryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BinaryRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DumpRequest)(nil), "debug.DumpRequest")
	proto.RegisterType((*ProfileRequest)(nil), "debug.ProfileRequest")
	proto.RegisterType((*BinaryRequest)(nil), "debug.BinaryRequest")
}

func init() { proto.RegisterFile("client/debug/debug.proto", fileDescriptor_6d15a320d0127c22) }

var fileDescriptor_6d15a320d0127c22 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x17, 0x71, 0x7f, 0xcc, 0x50, 0x21, 0x28, 0xd4, 0x09, 0x65, 0xf4, 0x34, 0x2f, 0x89,
	0x4c, 0x3c, 0x79, 0x90, 0x8d, 0x7e, 0x00, 0xe9, 0xc1, 0x83, 0xb7, 0xb4, 0x7d, 0xd7, 0x15, 0xba,
	0x26, 0xa6, 0x09, 0xd2, 0x6f, 0xe2, 0x47, 0xda, 0xd1, 0x8f, 0x20, 0xf5, 0x8b, 0xc8, 0x9a, 0x76,
	0x74, 0xec, 0xb0, 0x4b, 0xe9, 0xf3, 0xbe, 0x4f, 0x9e, 0xf7, 0xf7, 0x26, 0xd8, 0x89, 0xb2, 0x14,
	0x72, 0xcd, 0x62, 0x08, 0x4d, 0x62, 0xbf, 0x54, 0x2a, 0xa1, 0x05, 0xe9, 0xd7, 0x62, 0xe2, 0x26,
	0x42, 0x24, 0x19, 0xb0, 0xba, 0x18, 0x9a, 0x15, 0xfb, 0x52, 0x5c, 0x4a, 0x50, 0x85, 0xb5, 0x1d,
	0xf7, 0x63, 0xa3, 0xb8, 0x4e, 0x45, 0x6e, 0xfb, 0xde, 0x03, 0x1e, 0xfb, 0x66, 0x23, 0x03, 0xf8,
	0x34, 0x50, 0x68, 0x32, 0xc1, 0x23, 0x05, 0x91, 0x51, 0x05, 0xc4, 0x0e, 0x9a, 0xa2, 0xd9, 0x28,
	0xd8, 0x6b, 0x8f, 0xe3, 0xab, 0x37, 0x25, 0x56, 0x69, 0x06, 0xad, 0xdb, 0xc1, 0x43, 0x69, 0x2b,
	0xb5, 0xf9, 0x22, 0x68, 0x25, 0x79, 0xc6, 0xa3, 0x76, 0x90, 0x73, 0x36, 0x45, 0xb3, 0xf1, 0xfc,
	0x8e, 0x5a, 0x12, 0xda, 0x92, 0x50, 0xbf, 0x31, 0x04, 0x7b, 0xab, 0x77, 0x8d, 0x2f, 0x97, 0x69,
	0xce, 0x55, 0xd9, 0x4c, 0x98, 0x6f, 0x11, 0xee, 0xfb, 0xbb, 0x45, 0xc9, 0x0b, 0x3e, 0xdf, 0x81,
	0x12, 0x42, 0xed, 0x2d, 0x74, 0xa8, 0x27, 0xf7, 0x47, 0xd9, 0xcb, 0x52, 0x43, 0xf1, 0xce, 0x33,
	0x03, 0x5e, 0xef, 0x11, 0x91, 0x05, 0x1e, 0x36, 0xe8, 0xe4, 0xb6, 0x39, 0x7f, 0xb8, 0xca, 0xe9,
	0x88, 0x57, 0x3c, 0xb0, 0x68, 0xe4, 0xa6, 0x49, 0x38, 0x20, 0x3d, 0x19, 0xb0, 0x5c, 0x6c, 0x2b,
	0x17, 0xfd, 0x54, 0x2e, 0xfa, 0xad, 0x5c, 0xf4, 0xfd, 0xe7, 0xf6, 0x3e, 0x58, 0x92, 0xea, 0xb5,
	0x09, 0x69, 0x24, 0x36, 0x4c, 0xf2, 0x68, 0x5d, 0xc6, 0xa0, 0xba, 0x7f, 0x85, 0x8a, 0x58, 0xf7,
	0xfd, 0xc3, 0x41, 0x9d, 0xfd, 0xf4, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x93, 0x2d, 0xd0, 0x8a, 0x16,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DebugClient is the client API for Debug service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DebugClient interface {
	Dump(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (Debug_DumpClient, error)
	Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (Debug_ProfileClient, error)
	Binary(ctx context.Context, in *BinaryRequest, opts ...grpc.CallOption) (Debug_BinaryClient, error)
}

type debugClient struct {
	cc *grpc.ClientConn
}

func NewDebugClient(cc *grpc.ClientConn) DebugClient {
	return &debugClient{cc}
}

func (c *debugClient) Dump(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (Debug_DumpClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Debug_serviceDesc.Streams[0], "/debug.Debug/Dump", opts...)
	if err != nil {
		return nil, err
	}
	x := &debugDumpClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_DumpClient interface {
	Recv() (*types.BytesValue, error)
	grpc.ClientStream
}

type debugDumpClient struct {
	grpc.ClientStream
}

func (x *debugDumpClient) Recv() (*types.BytesValue, error) {
	m := new(types.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugClient) Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (Debug_ProfileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Debug_serviceDesc.Streams[1], "/debug.Debug/Profile", opts...)
	if err != nil {
		return nil, err
	}
	x := &debugProfileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_ProfileClient interface {
	Recv() (*types.BytesValue, error)
	grpc.ClientStream
}

type debugProfileClient struct {
	grpc.ClientStream
}

func (x *debugProfileClient) Recv() (*types.BytesValue, error) {
	m := new(types.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugClient) Binary(ctx context.Context, in *BinaryRequest, opts ...grpc.CallOption) (Debug_BinaryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Debug_serviceDesc.Streams[2], "/debug.Debug/Binary", opts...)
	if err != nil {
		return nil, err
	}
	x := &debugBinaryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_BinaryClient interface {
	Recv() (*types.BytesValue, error)
	grpc.ClientStream
}

type debugBinaryClient struct {
	grpc.ClientStream
}

func (x *debugBinaryClient) Recv() (*types.BytesValue, error) {
	m := new(types.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DebugServer is the server API for Debug service.
type DebugServer interface {
	Dump(*DumpRequest, Debug_DumpServer) error
	Profile(*ProfileRequest, Debug_ProfileServer) error
	Binary(*BinaryRequest, Debug_BinaryServer) error
}

// UnimplementedDebugServer can be embedded to have forward compatible implementations.
type UnimplementedDebugServer struct {
}

func (*UnimplementedDebugServer) Dump(req *DumpRequest, srv Debug_DumpServer) error {
	return status.Errorf(codes.Unimplemented, "method Dump not implemented")
}
func (*UnimplementedDebugServer) Profile(req *ProfileRequest, srv Debug_ProfileServer) error {
	return status.Errorf(codes.Unimplemented, "method Profile not implemented")
}
func (*UnimplementedDebugServer) Binary(req *BinaryRequest, srv Debug_BinaryServer) error {
	return status.Errorf(codes.Unimplemented, "method Binary not implemented")
}

func RegisterDebugServer(s *grpc.Server, srv DebugServer) {
	s.RegisterService(&_Debug_serviceDesc, srv)
}

func _Debug_Dump_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DumpRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).Dump(m, &debugDumpServer{stream})
}

type Debug_DumpServer interface {
	Send(*types.BytesValue) error
	grpc.ServerStream
}

type debugDumpServer struct {
	grpc.ServerStream
}

func (x *debugDumpServer) Send(m *types.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func _Debug_Profile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProfileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).Profile(m, &debugProfileServer{stream})
}

type Debug_ProfileServer interface {
	Send(*types.BytesValue) error
	grpc.ServerStream
}

type debugProfileServer struct {
	grpc.ServerStream
}

func (x *debugProfileServer) Send(m *types.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func _Debug_Binary_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BinaryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).Binary(m, &debugBinaryServer{stream})
}

type Debug_BinaryServer interface {
	Send(*types.BytesValue) error
	grpc.ServerStream
}

type debugBinaryServer struct {
	grpc.ServerStream
}

func (x *debugBinaryServer) Send(m *types.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

var _Debug_serviceDesc = grpc.ServiceDesc{
	ServiceName: "debug.Debug",
	HandlerType: (*DebugServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Dump",
			Handler:       _Debug_Dump_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Profile",
			Handler:       _Debug_Profile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Binary",
			Handler:       _Debug_Binary_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "client/debug/debug.proto",
}

func (m *DumpRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DumpRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DumpRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Recursed {
		i--
		if m.Recursed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ProfileRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProfileRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProfileRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Duration != nil {
		{
			size, err := m.Duration.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDebug(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Profile) > 0 {
		i -= len(m.Profile)
		copy(dAtA[i:], m.Profile)
		i = encodeVarintDebug(dAtA, i, uint64(len(m.Profile)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BinaryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BinaryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BinaryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintDebug(dAtA []byte, offset int, v uint64) int {
	offset -= sovDebug(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DumpRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Recursed {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ProfileRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Profile)
	if l > 0 {
		n += 1 + l + sovDebug(uint64(l))
	}
	if m.Duration != nil {
		l = m.Duration.Size()
		n += 1 + l + sovDebug(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BinaryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDebug(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDebug(x uint64) (n int) {
	return sovDebug(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DumpRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDebug
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DumpRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DumpRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recursed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Recursed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipDebug(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProfileRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDebug
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProfileRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProfileRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Profile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDebug
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Profile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDebug
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Duration == nil {
				m.Duration = &types.Duration{}
			}
			if err := m.Duration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDebug(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BinaryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDebug
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BinaryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BinaryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDebug(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDebug(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDebug
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDebug
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthDebug
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDebug
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDebug(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthDebug
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDebug = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDebug   = fmt.Errorf("proto: integer overflow")
)
