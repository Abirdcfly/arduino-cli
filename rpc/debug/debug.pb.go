// Code generated by protoc-gen-go. DO NOT EDIT.
// source: debug/debug.proto

package commands

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

// The top-level message sent by the client for the `StreamingOpen` method.
// Multiple `StreamingOpenReq` messages can be sent but the first message
// must contain a `DebugReq` message to initialize the debug session.
// All subsequent messages must contain bytes to be sent to the debug session
// and must not contain a `DebugReq` message.
type StreamingOpenReq struct {
	// Content must be either a debug session config or data to be sent.
	//
	// Types that are valid to be assigned to Content:
	//	*StreamingOpenReq_DebugReq
	//	*StreamingOpenReq_Data
	Content              isStreamingOpenReq_Content `protobuf_oneof:"content"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *StreamingOpenReq) Reset()         { *m = StreamingOpenReq{} }
func (m *StreamingOpenReq) String() string { return proto.CompactTextString(m) }
func (*StreamingOpenReq) ProtoMessage()    {}
func (*StreamingOpenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae24eab94cb53d5, []int{0}
}

func (m *StreamingOpenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingOpenReq.Unmarshal(m, b)
}
func (m *StreamingOpenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingOpenReq.Marshal(b, m, deterministic)
}
func (m *StreamingOpenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingOpenReq.Merge(m, src)
}
func (m *StreamingOpenReq) XXX_Size() int {
	return xxx_messageInfo_StreamingOpenReq.Size(m)
}
func (m *StreamingOpenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingOpenReq.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingOpenReq proto.InternalMessageInfo

type isStreamingOpenReq_Content interface {
	isStreamingOpenReq_Content()
}

type StreamingOpenReq_DebugReq struct {
	DebugReq *DebugReq `protobuf:"bytes,1,opt,name=debugReq,proto3,oneof"`
}

type StreamingOpenReq_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*StreamingOpenReq_DebugReq) isStreamingOpenReq_Content() {}

func (*StreamingOpenReq_Data) isStreamingOpenReq_Content() {}

func (m *StreamingOpenReq) GetContent() isStreamingOpenReq_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *StreamingOpenReq) GetDebugReq() *DebugReq {
	if x, ok := m.GetContent().(*StreamingOpenReq_DebugReq); ok {
		return x.DebugReq
	}
	return nil
}

func (m *StreamingOpenReq) GetData() []byte {
	if x, ok := m.GetContent().(*StreamingOpenReq_Data); ok {
		return x.Data
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StreamingOpenReq) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StreamingOpenReq_DebugReq)(nil),
		(*StreamingOpenReq_Data)(nil),
	}
}

type DebugReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DebugReq) Reset()         { *m = DebugReq{} }
func (m *DebugReq) String() string { return proto.CompactTextString(m) }
func (*DebugReq) ProtoMessage()    {}
func (*DebugReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae24eab94cb53d5, []int{1}
}

func (m *DebugReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugReq.Unmarshal(m, b)
}
func (m *DebugReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugReq.Marshal(b, m, deterministic)
}
func (m *DebugReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugReq.Merge(m, src)
}
func (m *DebugReq) XXX_Size() int {
	return xxx_messageInfo_DebugReq.Size(m)
}
func (m *DebugReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugReq.DiscardUnknown(m)
}

var xxx_messageInfo_DebugReq proto.InternalMessageInfo

func (m *DebugReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

//
type StreamingOpenResp struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingOpenResp) Reset()         { *m = StreamingOpenResp{} }
func (m *StreamingOpenResp) String() string { return proto.CompactTextString(m) }
func (*StreamingOpenResp) ProtoMessage()    {}
func (*StreamingOpenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae24eab94cb53d5, []int{2}
}

func (m *StreamingOpenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingOpenResp.Unmarshal(m, b)
}
func (m *StreamingOpenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingOpenResp.Marshal(b, m, deterministic)
}
func (m *StreamingOpenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingOpenResp.Merge(m, src)
}
func (m *StreamingOpenResp) XXX_Size() int {
	return xxx_messageInfo_StreamingOpenResp.Size(m)
}
func (m *StreamingOpenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingOpenResp.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingOpenResp proto.InternalMessageInfo

func (m *StreamingOpenResp) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// duplicate from commands/common.proto
// as module imports seems not to work
type Instance struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}
func (*Instance) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae24eab94cb53d5, []int{3}
}

func (m *Instance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instance.Unmarshal(m, b)
}
func (m *Instance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instance.Marshal(b, m, deterministic)
}
func (m *Instance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instance.Merge(m, src)
}
func (m *Instance) XXX_Size() int {
	return xxx_messageInfo_Instance.Size(m)
}
func (m *Instance) XXX_DiscardUnknown() {
	xxx_messageInfo_Instance.DiscardUnknown(m)
}

var xxx_messageInfo_Instance proto.InternalMessageInfo

func (m *Instance) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*StreamingOpenReq)(nil), "cc.arduino.cli.commands.StreamingOpenReq")
	proto.RegisterType((*DebugReq)(nil), "cc.arduino.cli.commands.DebugReq")
	proto.RegisterType((*StreamingOpenResp)(nil), "cc.arduino.cli.commands.StreamingOpenResp")
	proto.RegisterType((*Instance)(nil), "cc.arduino.cli.commands.Instance")
}

func init() { proto.RegisterFile("debug/debug.proto", fileDescriptor_5ae24eab94cb53d5) }

var fileDescriptor_5ae24eab94cb53d5 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0x9b, 0xe2, 0xb4, 0x7e, 0x4e, 0x71, 0x41, 0x70, 0xec, 0x34, 0x73, 0xb1, 0x2a, 0x4b,
	0x65, 0x9e, 0x45, 0x18, 0x1e, 0xb6, 0x93, 0x10, 0x6f, 0xde, 0xd2, 0x24, 0xd4, 0x40, 0x9b, 0x64,
	0x6d, 0xfa, 0xff, 0xcb, 0x62, 0x3b, 0x70, 0x50, 0xdc, 0x25, 0x21, 0xe4, 0xfd, 0xde, 0x7b, 0xc9,
	0x07, 0x13, 0xa9, 0xf2, 0xb6, 0xc8, 0xc2, 0x4a, 0x5d, 0x6d, 0xbd, 0xc5, 0xb7, 0x42, 0x50, 0x5e,
	0xcb, 0x56, 0x1b, 0x4b, 0x45, 0xa9, 0xa9, 0xb0, 0x55, 0xc5, 0x8d, 0x6c, 0x88, 0x87, 0xeb, 0x4f,
	0x5f, 0x2b, 0x5e, 0x69, 0x53, 0x7c, 0x38, 0x65, 0x98, 0xda, 0xe2, 0x37, 0x48, 0x02, 0xcb, 0xd4,
	0x76, 0x8a, 0xe6, 0x28, 0xbd, 0x58, 0xde, 0xd1, 0x01, 0x9e, 0xbe, 0x77, 0xc2, 0x75, 0xc4, 0xf6,
	0x10, 0xbe, 0x81, 0x13, 0xc9, 0x3d, 0x9f, 0xc6, 0x73, 0x94, 0x8e, 0xd7, 0x11, 0x0b, 0xa7, 0xd5,
	0x39, 0x9c, 0x09, 0x6b, 0xbc, 0x32, 0x9e, 0x6c, 0x20, 0xe9, 0x41, 0xfc, 0x0a, 0x89, 0x36, 0x8d,
	0xe7, 0x46, 0xa8, 0x7f, 0xd3, 0x36, 0x9d, 0x90, 0xed, 0x11, 0x72, 0x0f, 0x93, 0x83, 0x07, 0x34,
	0x0e, 0xe3, 0xae, 0xc0, 0xce, 0x6f, 0xfc, 0x1b, 0x4f, 0x66, 0x90, 0xf4, 0x38, 0xbe, 0x82, 0x58,
	0xcb, 0x70, 0x3b, 0x62, 0xb1, 0x96, 0xcb, 0x16, 0x46, 0xa1, 0x0f, 0x2e, 0xe1, 0xf2, 0x8f, 0x1b,
	0x7e, 0x18, 0xec, 0x72, 0xf8, 0x6d, 0xb3, 0xc7, 0x63, 0xa5, 0x8d, 0x23, 0x51, 0x8a, 0x9e, 0xd1,
	0x6a, 0xf1, 0xf5, 0x54, 0x68, 0xff, 0xdd, 0xe6, 0x3b, 0x69, 0xd6, 0xa1, 0xfd, 0xbe, 0x10, 0xa5,
	0xce, 0x6a, 0x27, 0xb2, 0xde, 0x26, 0x3f, 0x0d, 0xb3, 0x7c, 0xf9, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xad, 0x8b, 0xed, 0xc4, 0xe0, 0x01, 0x00, 0x00,
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
	StreamingOpen(ctx context.Context, opts ...grpc.CallOption) (Debug_StreamingOpenClient, error)
}

type debugClient struct {
	cc *grpc.ClientConn
}

func NewDebugClient(cc *grpc.ClientConn) DebugClient {
	return &debugClient{cc}
}

func (c *debugClient) StreamingOpen(ctx context.Context, opts ...grpc.CallOption) (Debug_StreamingOpenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Debug_serviceDesc.Streams[0], "/cc.arduino.cli.commands.Debug/StreamingOpen", opts...)
	if err != nil {
		return nil, err
	}
	x := &debugStreamingOpenClient{stream}
	return x, nil
}

type Debug_StreamingOpenClient interface {
	Send(*StreamingOpenReq) error
	Recv() (*StreamingOpenResp, error)
	grpc.ClientStream
}

type debugStreamingOpenClient struct {
	grpc.ClientStream
}

func (x *debugStreamingOpenClient) Send(m *StreamingOpenReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *debugStreamingOpenClient) Recv() (*StreamingOpenResp, error) {
	m := new(StreamingOpenResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DebugServer is the server API for Debug service.
type DebugServer interface {
	StreamingOpen(Debug_StreamingOpenServer) error
}

func RegisterDebugServer(s *grpc.Server, srv DebugServer) {
	s.RegisterService(&_Debug_serviceDesc, srv)
}

func _Debug_StreamingOpen_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DebugServer).StreamingOpen(&debugStreamingOpenServer{stream})
}

type Debug_StreamingOpenServer interface {
	Send(*StreamingOpenResp) error
	Recv() (*StreamingOpenReq, error)
	grpc.ServerStream
}

type debugStreamingOpenServer struct {
	grpc.ServerStream
}

func (x *debugStreamingOpenServer) Send(m *StreamingOpenResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *debugStreamingOpenServer) Recv() (*StreamingOpenReq, error) {
	m := new(StreamingOpenReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Debug_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cc.arduino.cli.commands.Debug",
	HandlerType: (*DebugServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingOpen",
			Handler:       _Debug_StreamingOpen_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "debug/debug.proto",
}