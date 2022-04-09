// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/containerd/containerd/api/services/events/v1/events.proto

package events

import (
	context "context"
	fmt "fmt"
	_ "github.com/containerd/containerd/protobuf/plugin"
	github_com_containerd_typeurl "github.com/containerd/typeurl"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PublishRequest struct {
	Topic                string     `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Event                *types.Any `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PublishRequest) Reset()      { *m = PublishRequest{} }
func (*PublishRequest) ProtoMessage() {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_43fcd20dc1642376, []int{0}
}
func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(m, src)
}
func (m *PublishRequest) XXX_Size() int {
	return m.Size()
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

type ForwardRequest struct {
	Envelope             *Envelope `protobuf:"bytes,1,opt,name=envelope,proto3" json:"envelope,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ForwardRequest) Reset()      { *m = ForwardRequest{} }
func (*ForwardRequest) ProtoMessage() {}
func (*ForwardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_43fcd20dc1642376, []int{1}
}
func (m *ForwardRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ForwardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ForwardRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ForwardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwardRequest.Merge(m, src)
}
func (m *ForwardRequest) XXX_Size() int {
	return m.Size()
}
func (m *ForwardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ForwardRequest proto.InternalMessageInfo

type SubscribeRequest struct {
	Filters              []string `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()      { *m = SubscribeRequest{} }
func (*SubscribeRequest) ProtoMessage() {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_43fcd20dc1642376, []int{2}
}
func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return m.Size()
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

type Envelope struct {
	Timestamp            time.Time  `protobuf:"bytes,1,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	Namespace            string     `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Topic                string     `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Event                *types.Any `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Envelope) Reset()      { *m = Envelope{} }
func (*Envelope) ProtoMessage() {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_43fcd20dc1642376, []int{3}
}
func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(m, src)
}
func (m *Envelope) XXX_Size() int {
	return m.Size()
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PublishRequest)(nil), "containerd.services.events.v1.PublishRequest")
	proto.RegisterType((*ForwardRequest)(nil), "containerd.services.events.v1.ForwardRequest")
	proto.RegisterType((*SubscribeRequest)(nil), "containerd.services.events.v1.SubscribeRequest")
	proto.RegisterType((*Envelope)(nil), "containerd.services.events.v1.Envelope")
}

func init() {
	proto.RegisterFile("github.com/containerd/containerd/api/services/events/v1/events.proto", fileDescriptor_43fcd20dc1642376)
}

var fileDescriptor_43fcd20dc1642376 = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0xeb, 0xf9, 0x6d, 0x3c, 0xd2, 0x08, 0x45, 0x15, 0x2a, 0x01, 0xd2, 0xaa, 0x1b, 0x2a,
	0x04, 0x0e, 0x53, 0x76, 0x20, 0x21, 0x28, 0x94, 0xf5, 0x28, 0x80, 0x84, 0xd8, 0x25, 0xe9, 0x6d,
	0x6a, 0x29, 0xb1, 0x4d, 0xec, 0x04, 0xcd, 0x6e, 0x1e, 0x81, 0x0d, 0x6f, 0xc2, 0x86, 0x37, 0xe8,
	0x92, 0x25, 0x2b, 0x60, 0xfa, 0x24, 0xa8, 0x89, 0xdd, 0x30, 0x2d, 0x10, 0x34, 0xbb, 0x6b, 0xdf,
	0xe3, 0xcf, 0xb9, 0xe7, 0x38, 0xf8, 0x45, 0x4c, 0xd5, 0x3c, 0x0f, 0x49, 0xc4, 0x53, 0x2f, 0xe2,
	0x4c, 0x05, 0x94, 0x41, 0x36, 0xfd, 0xbd, 0x0c, 0x04, 0xf5, 0x24, 0x64, 0x05, 0x8d, 0x40, 0x7a,
	0x50, 0x00, 0x53, 0xd2, 0x2b, 0x4e, 0x74, 0x45, 0x44, 0xc6, 0x15, 0xb7, 0x6f, 0xd7, 0x7a, 0x62,
	0xb4, 0x44, 0x2b, 0x8a, 0x13, 0xe7, 0x69, 0xe3, 0x25, 0x25, 0x26, 0xcc, 0x67, 0x9e, 0x48, 0xf2,
	0x98, 0x32, 0x6f, 0x46, 0x21, 0x99, 0x8a, 0x40, 0xcd, 0xab, 0x0b, 0x9c, 0x4e, 0xcc, 0x63, 0x5e,
	0x96, 0xde, 0xaa, 0xd2, 0xbb, 0x37, 0x62, 0xce, 0xe3, 0x04, 0xea, 0xd3, 0x01, 0x3b, 0xd3, 0xad,
	0x9b, 0x9b, 0x2d, 0x48, 0x85, 0x32, 0xcd, 0xde, 0x66, 0x53, 0xd1, 0x14, 0xa4, 0x0a, 0x52, 0x51,
	0x09, 0x06, 0x3e, 0x3e, 0x3e, 0xcd, 0xc3, 0x84, 0xca, 0xb9, 0x0f, 0xef, 0x73, 0x90, 0xca, 0xee,
	0xe0, 0x7d, 0xc5, 0x05, 0x8d, 0xba, 0xa8, 0x8f, 0x86, 0x96, 0x5f, 0x2d, 0xec, 0xbb, 0x78, 0xbf,
	0x9c, 0xb2, 0xbb, 0xd3, 0x47, 0xc3, 0xa3, 0x51, 0x87, 0x54, 0x60, 0x62, 0xc0, 0xe4, 0x19, 0x3b,
	0xf3, 0x2b, 0xc9, 0xe0, 0x0d, 0x3e, 0x7e, 0xc9, 0xb3, 0x0f, 0x41, 0x36, 0x35, 0xcc, 0xe7, 0xb8,
	0x0d, 0xac, 0x80, 0x84, 0x0b, 0x28, 0xb1, 0x47, 0xa3, 0x3b, 0xe4, 0x9f, 0x46, 0x92, 0x89, 0x96,
	0xfb, 0xeb, 0x83, 0x83, 0x7b, 0xf8, 0xda, 0xab, 0x3c, 0x94, 0x51, 0x46, 0x43, 0x30, 0xe0, 0x2e,
	0x3e, 0x9c, 0xd1, 0x44, 0x41, 0x26, 0xbb, 0xa8, 0xbf, 0x3b, 0xb4, 0x7c, 0xb3, 0x1c, 0x7c, 0x46,
	0xb8, 0x6d, 0x20, 0xf6, 0x18, 0x5b, 0xeb, 0xc1, 0xf5, 0x07, 0x38, 0x5b, 0x13, 0xbc, 0x36, 0x8a,
	0x71, 0x7b, 0xf1, 0xbd, 0xd7, 0xfa, 0xf8, 0xa3, 0x87, 0xfc, 0xfa, 0x98, 0x7d, 0x0b, 0x5b, 0x2c,
	0x48, 0x41, 0x8a, 0x20, 0x82, 0xd2, 0x05, 0xcb, 0xaf, 0x37, 0x6a, 0xd7, 0x76, 0xff, 0xe8, 0xda,
	0x5e, 0xa3, 0x6b, 0x8f, 0xf6, 0xce, 0xbf, 0xf4, 0xd0, 0xe8, 0xd3, 0x0e, 0x3e, 0x98, 0x94, 0x2e,
	0xd8, 0xa7, 0xf8, 0x50, 0x47, 0x63, 0xdf, 0x6f, 0x70, 0xeb, 0x72, 0x84, 0xce, 0xf5, 0xad, 0x7b,
	0x26, 0xab, 0x37, 0xb1, 0x22, 0xea, 0x60, 0x1a, 0x89, 0x97, 0x03, 0xfc, 0x2b, 0x31, 0xc6, 0xd6,
	0x3a, 0x13, 0xdb, 0x6b, 0x60, 0x6e, 0xa6, 0xe7, 0xfc, 0xef, 0x23, 0x78, 0x80, 0xc6, 0x6f, 0x17,
	0x17, 0x6e, 0xeb, 0xdb, 0x85, 0xdb, 0x3a, 0x5f, 0xba, 0x68, 0xb1, 0x74, 0xd1, 0xd7, 0xa5, 0x8b,
	0x7e, 0x2e, 0x5d, 0xf4, 0xee, 0xc9, 0x15, 0xff, 0xeb, 0xc7, 0x55, 0x15, 0x1e, 0x94, 0x23, 0x3d,
	0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x38, 0x37, 0x72, 0x20, 0x04, 0x00, 0x00,
}

// Field returns the value for the given fieldpath as a string, if defined.
// If the value is not defined, the second value will be false.
func (m *Envelope) Field(fieldpath []string) (string, bool) {
	if len(fieldpath) == 0 {
		return "", false
	}

	switch fieldpath[0] {
	// unhandled: timestamp
	case "namespace":
		return string(m.Namespace), len(m.Namespace) > 0
	case "topic":
		return string(m.Topic), len(m.Topic) > 0
	case "event":
		decoded, err := github_com_containerd_typeurl.UnmarshalAny(m.Event)
		if err != nil {
			return "", false
		}

		adaptor, ok := decoded.(interface{ Field([]string) (string, bool) })
		if !ok {
			return "", false
		}
		return adaptor.Field(fieldpath[1:])
	}
	return "", false
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventsClient is the client API for Events service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventsClient interface {
	// Publish an event to a topic.
	//
	// The event will be packed into a timestamp envelope with the namespace
	// introspected from the context. The envelope will then be dispatched.
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*types.Empty, error)
	// Forward sends an event that has already been packaged into an envelope
	// with a timestamp and namespace.
	//
	// This is useful if earlier timestamping is required or when forwarding on
	// behalf of another component, namespace or publisher.
	Forward(ctx context.Context, in *ForwardRequest, opts ...grpc.CallOption) (*types.Empty, error)
	// Subscribe to a stream of events, possibly returning only that match any
	// of the provided filters.
	//
	// Unlike many other methods in containerd, subscribers will get messages
	// from all namespaces unless otherwise specified. If this is not desired,
	// a filter can be provided in the format 'namespace==<namespace>' to
	// restrict the received events.
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Events_SubscribeClient, error)
}

type eventsClient struct {
	cc *grpc.ClientConn
}

func NewEventsClient(cc *grpc.ClientConn) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/containerd.services.events.v1.Events/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) Forward(ctx context.Context, in *ForwardRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/containerd.services.events.v1.Events/Forward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Events_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Events_serviceDesc.Streams[0], "/containerd.services.events.v1.Events/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventsSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Events_SubscribeClient interface {
	Recv() (*Envelope, error)
	grpc.ClientStream
}

type eventsSubscribeClient struct {
	grpc.ClientStream
}

func (x *eventsSubscribeClient) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventsServer is the server API for Events service.
type EventsServer interface {
	// Publish an event to a topic.
	//
	// The event will be packed into a timestamp envelope with the namespace
	// introspected from the context. The envelope will then be dispatched.
	Publish(context.Context, *PublishRequest) (*types.Empty, error)
	// Forward sends an event that has already been packaged into an envelope
	// with a timestamp and namespace.
	//
	// This is useful if earlier timestamping is required or when forwarding on
	// behalf of another component, namespace or publisher.
	Forward(context.Context, *ForwardRequest) (*types.Empty, error)
	// Subscribe to a stream of events, possibly returning only that match any
	// of the provided filters.
	//
	// Unlike many other methods in containerd, subscribers will get messages
	// from all namespaces unless otherwise specified. If this is not desired,
	// a filter can be provided in the format 'namespace==<namespace>' to
	// restrict the received events.
	Subscribe(*SubscribeRequest, Events_SubscribeServer) error
}

// UnimplementedEventsServer can be embedded to have forward compatible implementations.
type UnimplementedEventsServer struct {
}

func (*UnimplementedEventsServer) Publish(ctx context.Context, req *PublishRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedEventsServer) Forward(ctx context.Context, req *ForwardRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Forward not implemented")
}
func (*UnimplementedEventsServer) Subscribe(req *SubscribeRequest, srv Events_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.events.v1.Events/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_Forward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForwardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).Forward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.events.v1.Events/Forward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).Forward(ctx, req.(*ForwardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventsServer).Subscribe(m, &eventsSubscribeServer{stream})
}

type Events_SubscribeServer interface {
	Send(*Envelope) error
	grpc.ServerStream
}

type eventsSubscribeServer struct {
	grpc.ServerStream
}

func (x *eventsSubscribeServer) Send(m *Envelope) error {
	return x.ServerStream.SendMsg(m)
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.services.events.v1.Events",
	HandlerType: (*EventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Events_Publish_Handler,
		},
		{
			MethodName: "Forward",
			Handler:    _Events_Forward_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Events_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/containerd/containerd/api/services/events/v1/events.proto",
}

func (m *PublishRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublishRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublishRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Event != nil {
		{
			size, err := m.Event.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Topic) > 0 {
		i -= len(m.Topic)
		copy(dAtA[i:], m.Topic)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Topic)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ForwardRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ForwardRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ForwardRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Envelope != nil {
		{
			size, err := m.Envelope.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SubscribeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubscribeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubscribeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Filters) > 0 {
		for iNdEx := len(m.Filters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Filters[iNdEx])
			copy(dAtA[i:], m.Filters[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.Filters[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Envelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Envelope) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Envelope) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Event != nil {
		{
			size, err := m.Event.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Topic) > 0 {
		i -= len(m.Topic)
		copy(dAtA[i:], m.Topic)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Topic)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintEvents(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PublishRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Topic)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Event != nil {
		l = m.Event.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ForwardRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Envelope != nil {
		l = m.Envelope.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SubscribeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Filters) > 0 {
		for _, s := range m.Filters {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Envelope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovEvents(uint64(l))
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Topic)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Event != nil {
		l = m.Event.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PublishRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PublishRequest{`,
		`Topic:` + fmt.Sprintf("%v", this.Topic) + `,`,
		`Event:` + strings.Replace(fmt.Sprintf("%v", this.Event), "Any", "types.Any", 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ForwardRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ForwardRequest{`,
		`Envelope:` + strings.Replace(this.Envelope.String(), "Envelope", "Envelope", 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SubscribeRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SubscribeRequest{`,
		`Filters:` + fmt.Sprintf("%v", this.Filters) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Envelope) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Envelope{`,
		`Timestamp:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Timestamp), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Topic:` + fmt.Sprintf("%v", this.Topic) + `,`,
		`Event:` + strings.Replace(fmt.Sprintf("%v", this.Event), "Any", "types.Any", 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringEvents(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PublishRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: PublishRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublishRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Event == nil {
				m.Event = &types.Any{}
			}
			if err := m.Event.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *ForwardRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: ForwardRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ForwardRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Envelope", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Envelope == nil {
				m.Envelope = &Envelope{}
			}
			if err := m.Envelope.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *SubscribeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: SubscribeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubscribeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filters = append(m.Filters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *Envelope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: Envelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Envelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Event == nil {
				m.Event = &types.Any{}
			}
			if err := m.Event.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
