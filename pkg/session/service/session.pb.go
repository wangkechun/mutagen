// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/havoc-io/mutagen/pkg/session/service/session.proto

package session

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import session "github.com/havoc-io/mutagen/pkg/session"
import url "github.com/havoc-io/mutagen/pkg/url"

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

type CreateRequest struct {
	Alpha                *url.URL               `protobuf:"bytes,1,opt,name=alpha,proto3" json:"alpha,omitempty"`
	Beta                 *url.URL               `protobuf:"bytes,2,opt,name=beta,proto3" json:"beta,omitempty"`
	Configuration        *session.Configuration `protobuf:"bytes,3,opt,name=configuration,proto3" json:"configuration,omitempty"`
	Response             string                 `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_004cdd49a5f4cc5f, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetAlpha() *url.URL {
	if m != nil {
		return m.Alpha
	}
	return nil
}

func (m *CreateRequest) GetBeta() *url.URL {
	if m != nil {
		return m.Beta
	}
	return nil
}

func (m *CreateRequest) GetConfiguration() *session.Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *CreateRequest) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type CreateResponse struct {
	Session              string   `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Prompt               string   `protobuf:"bytes,3,opt,name=prompt,proto3" json:"prompt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_004cdd49a5f4cc5f, []int{1}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (dst *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(dst, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *CreateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateResponse) GetPrompt() string {
	if m != nil {
		return m.Prompt
	}
	return ""
}

type ListRequest struct {
	PreviousStateIndex   uint64   `protobuf:"varint,1,opt,name=previousStateIndex,proto3" json:"previousStateIndex,omitempty"`
	Specifications       []string `protobuf:"bytes,2,rep,name=specifications,proto3" json:"specifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_004cdd49a5f4cc5f, []int{2}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetPreviousStateIndex() uint64 {
	if m != nil {
		return m.PreviousStateIndex
	}
	return 0
}

func (m *ListRequest) GetSpecifications() []string {
	if m != nil {
		return m.Specifications
	}
	return nil
}

type ListResponse struct {
	StateIndex           uint64           `protobuf:"varint,1,opt,name=stateIndex,proto3" json:"stateIndex,omitempty"`
	SessionStates        []*session.State `protobuf:"bytes,2,rep,name=sessionStates,proto3" json:"sessionStates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_004cdd49a5f4cc5f, []int{3}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetStateIndex() uint64 {
	if m != nil {
		return m.StateIndex
	}
	return 0
}

func (m *ListResponse) GetSessionStates() []*session.State {
	if m != nil {
		return m.SessionStates
	}
	return nil
}

type PauseRequest struct {
	Specifications       []string `protobuf:"bytes,1,rep,name=specifications,proto3" json:"specifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PauseRequest) Reset()         { *m = PauseRequest{} }
func (m *PauseRequest) String() string { return proto.CompactTextString(m) }
func (*PauseRequest) ProtoMessage()    {}
func (*PauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_004cdd49a5f4cc5f, []int{4}
}
func (m *PauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PauseRequest.Unmarshal(m, b)
}
func (m *PauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PauseRequest.Marshal(b, m, deterministic)
}
func (dst *PauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PauseRequest.Merge(dst, src)
}
func (m *PauseRequest) XXX_Size() int {
	return xxx_messageInfo_PauseRequest.Size(m)
}
func (m *PauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PauseRequest proto.InternalMessageInfo

func (m *PauseRequest) GetSpecifications() []string {
	if m != nil {
		return m.Specifications
	}
	return nil
}

type PauseResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PauseResponse) Reset()         { *m = PauseResponse{} }
func (m *PauseResponse) String() string { return proto.CompactTextString(m) }
func (*PauseResponse) ProtoMessage()    {}
func (*PauseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_004cdd49a5f4cc5f, []int{5}
}
func (m *PauseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PauseResponse.Unmarshal(m, b)
}
func (m *PauseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PauseResponse.Marshal(b, m, deterministic)
}
func (dst *PauseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PauseResponse.Merge(dst, src)
}
func (m *PauseResponse) XXX_Size() int {
	return xxx_messageInfo_PauseResponse.Size(m)
}
func (m *PauseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PauseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PauseResponse proto.InternalMessageInfo

func (m *PauseResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ResumeRequest struct {
	Specifications       []string `protobuf:"bytes,1,rep,name=specifications,proto3" json:"specifications,omitempty"`
	Response             string   `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResumeRequest) Reset()         { *m = ResumeRequest{} }
func (m *ResumeRequest) String() string { return proto.CompactTextString(m) }
func (*ResumeRequest) ProtoMessage()    {}
func (*ResumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_004cdd49a5f4cc5f, []int{6}
}
func (m *ResumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeRequest.Unmarshal(m, b)
}
func (m *ResumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeRequest.Marshal(b, m, deterministic)
}
func (dst *ResumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeRequest.Merge(dst, src)
}
func (m *ResumeRequest) XXX_Size() int {
	return xxx_messageInfo_ResumeRequest.Size(m)
}
func (m *ResumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeRequest proto.InternalMessageInfo

func (m *ResumeRequest) GetSpecifications() []string {
	if m != nil {
		return m.Specifications
	}
	return nil
}

func (m *ResumeRequest) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type ResumeResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Prompt               string   `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResumeResponse) Reset()         { *m = ResumeResponse{} }
func (m *ResumeResponse) String() string { return proto.CompactTextString(m) }
func (*ResumeResponse) ProtoMessage()    {}
func (*ResumeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_004cdd49a5f4cc5f, []int{7}
}
func (m *ResumeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeResponse.Unmarshal(m, b)
}
func (m *ResumeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeResponse.Marshal(b, m, deterministic)
}
func (dst *ResumeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeResponse.Merge(dst, src)
}
func (m *ResumeResponse) XXX_Size() int {
	return xxx_messageInfo_ResumeResponse.Size(m)
}
func (m *ResumeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeResponse proto.InternalMessageInfo

func (m *ResumeResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ResumeResponse) GetPrompt() string {
	if m != nil {
		return m.Prompt
	}
	return ""
}

type TerminateRequest struct {
	Specifications       []string `protobuf:"bytes,1,rep,name=specifications,proto3" json:"specifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminateRequest) Reset()         { *m = TerminateRequest{} }
func (m *TerminateRequest) String() string { return proto.CompactTextString(m) }
func (*TerminateRequest) ProtoMessage()    {}
func (*TerminateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_004cdd49a5f4cc5f, []int{8}
}
func (m *TerminateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminateRequest.Unmarshal(m, b)
}
func (m *TerminateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminateRequest.Marshal(b, m, deterministic)
}
func (dst *TerminateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminateRequest.Merge(dst, src)
}
func (m *TerminateRequest) XXX_Size() int {
	return xxx_messageInfo_TerminateRequest.Size(m)
}
func (m *TerminateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TerminateRequest proto.InternalMessageInfo

func (m *TerminateRequest) GetSpecifications() []string {
	if m != nil {
		return m.Specifications
	}
	return nil
}

type TerminateResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminateResponse) Reset()         { *m = TerminateResponse{} }
func (m *TerminateResponse) String() string { return proto.CompactTextString(m) }
func (*TerminateResponse) ProtoMessage()    {}
func (*TerminateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_004cdd49a5f4cc5f, []int{9}
}
func (m *TerminateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminateResponse.Unmarshal(m, b)
}
func (m *TerminateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminateResponse.Marshal(b, m, deterministic)
}
func (dst *TerminateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminateResponse.Merge(dst, src)
}
func (m *TerminateResponse) XXX_Size() int {
	return xxx_messageInfo_TerminateResponse.Size(m)
}
func (m *TerminateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TerminateResponse proto.InternalMessageInfo

func (m *TerminateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "session.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "session.CreateResponse")
	proto.RegisterType((*ListRequest)(nil), "session.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "session.ListResponse")
	proto.RegisterType((*PauseRequest)(nil), "session.PauseRequest")
	proto.RegisterType((*PauseResponse)(nil), "session.PauseResponse")
	proto.RegisterType((*ResumeRequest)(nil), "session.ResumeRequest")
	proto.RegisterType((*ResumeResponse)(nil), "session.ResumeResponse")
	proto.RegisterType((*TerminateRequest)(nil), "session.TerminateRequest")
	proto.RegisterType((*TerminateResponse)(nil), "session.TerminateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionsClient is the client API for Sessions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionsClient interface {
	Create(ctx context.Context, opts ...grpc.CallOption) (Sessions_CreateClient, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Pause(ctx context.Context, opts ...grpc.CallOption) (Sessions_PauseClient, error)
	Resume(ctx context.Context, opts ...grpc.CallOption) (Sessions_ResumeClient, error)
	Terminate(ctx context.Context, opts ...grpc.CallOption) (Sessions_TerminateClient, error)
}

type sessionsClient struct {
	cc *grpc.ClientConn
}

func NewSessionsClient(cc *grpc.ClientConn) SessionsClient {
	return &sessionsClient{cc}
}

func (c *sessionsClient) Create(ctx context.Context, opts ...grpc.CallOption) (Sessions_CreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sessions_serviceDesc.Streams[0], "/session.Sessions/Create", opts...)
	if err != nil {
		return nil, err
	}
	x := &sessionsCreateClient{stream}
	return x, nil
}

type Sessions_CreateClient interface {
	Send(*CreateRequest) error
	Recv() (*CreateResponse, error)
	grpc.ClientStream
}

type sessionsCreateClient struct {
	grpc.ClientStream
}

func (x *sessionsCreateClient) Send(m *CreateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sessionsCreateClient) Recv() (*CreateResponse, error) {
	m := new(CreateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sessionsClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/session.Sessions/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionsClient) Pause(ctx context.Context, opts ...grpc.CallOption) (Sessions_PauseClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sessions_serviceDesc.Streams[1], "/session.Sessions/Pause", opts...)
	if err != nil {
		return nil, err
	}
	x := &sessionsPauseClient{stream}
	return x, nil
}

type Sessions_PauseClient interface {
	Send(*PauseRequest) error
	Recv() (*PauseResponse, error)
	grpc.ClientStream
}

type sessionsPauseClient struct {
	grpc.ClientStream
}

func (x *sessionsPauseClient) Send(m *PauseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sessionsPauseClient) Recv() (*PauseResponse, error) {
	m := new(PauseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sessionsClient) Resume(ctx context.Context, opts ...grpc.CallOption) (Sessions_ResumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sessions_serviceDesc.Streams[2], "/session.Sessions/Resume", opts...)
	if err != nil {
		return nil, err
	}
	x := &sessionsResumeClient{stream}
	return x, nil
}

type Sessions_ResumeClient interface {
	Send(*ResumeRequest) error
	Recv() (*ResumeResponse, error)
	grpc.ClientStream
}

type sessionsResumeClient struct {
	grpc.ClientStream
}

func (x *sessionsResumeClient) Send(m *ResumeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sessionsResumeClient) Recv() (*ResumeResponse, error) {
	m := new(ResumeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sessionsClient) Terminate(ctx context.Context, opts ...grpc.CallOption) (Sessions_TerminateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sessions_serviceDesc.Streams[3], "/session.Sessions/Terminate", opts...)
	if err != nil {
		return nil, err
	}
	x := &sessionsTerminateClient{stream}
	return x, nil
}

type Sessions_TerminateClient interface {
	Send(*TerminateRequest) error
	Recv() (*TerminateResponse, error)
	grpc.ClientStream
}

type sessionsTerminateClient struct {
	grpc.ClientStream
}

func (x *sessionsTerminateClient) Send(m *TerminateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sessionsTerminateClient) Recv() (*TerminateResponse, error) {
	m := new(TerminateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SessionsServer is the server API for Sessions service.
type SessionsServer interface {
	Create(Sessions_CreateServer) error
	List(context.Context, *ListRequest) (*ListResponse, error)
	Pause(Sessions_PauseServer) error
	Resume(Sessions_ResumeServer) error
	Terminate(Sessions_TerminateServer) error
}

func RegisterSessionsServer(s *grpc.Server, srv SessionsServer) {
	s.RegisterService(&_Sessions_serviceDesc, srv)
}

func _Sessions_Create_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SessionsServer).Create(&sessionsCreateServer{stream})
}

type Sessions_CreateServer interface {
	Send(*CreateResponse) error
	Recv() (*CreateRequest, error)
	grpc.ServerStream
}

type sessionsCreateServer struct {
	grpc.ServerStream
}

func (x *sessionsCreateServer) Send(m *CreateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sessionsCreateServer) Recv() (*CreateRequest, error) {
	m := new(CreateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Sessions_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Sessions/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionsServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sessions_Pause_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SessionsServer).Pause(&sessionsPauseServer{stream})
}

type Sessions_PauseServer interface {
	Send(*PauseResponse) error
	Recv() (*PauseRequest, error)
	grpc.ServerStream
}

type sessionsPauseServer struct {
	grpc.ServerStream
}

func (x *sessionsPauseServer) Send(m *PauseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sessionsPauseServer) Recv() (*PauseRequest, error) {
	m := new(PauseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Sessions_Resume_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SessionsServer).Resume(&sessionsResumeServer{stream})
}

type Sessions_ResumeServer interface {
	Send(*ResumeResponse) error
	Recv() (*ResumeRequest, error)
	grpc.ServerStream
}

type sessionsResumeServer struct {
	grpc.ServerStream
}

func (x *sessionsResumeServer) Send(m *ResumeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sessionsResumeServer) Recv() (*ResumeRequest, error) {
	m := new(ResumeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Sessions_Terminate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SessionsServer).Terminate(&sessionsTerminateServer{stream})
}

type Sessions_TerminateServer interface {
	Send(*TerminateResponse) error
	Recv() (*TerminateRequest, error)
	grpc.ServerStream
}

type sessionsTerminateServer struct {
	grpc.ServerStream
}

func (x *sessionsTerminateServer) Send(m *TerminateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sessionsTerminateServer) Recv() (*TerminateRequest, error) {
	m := new(TerminateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Sessions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "session.Sessions",
	HandlerType: (*SessionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Sessions_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Create",
			Handler:       _Sessions_Create_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Pause",
			Handler:       _Sessions_Pause_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Resume",
			Handler:       _Sessions_Resume_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Terminate",
			Handler:       _Sessions_Terminate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/havoc-io/mutagen/pkg/session/service/session.proto",
}

func init() {
	proto.RegisterFile("github.com/havoc-io/mutagen/pkg/session/service/session.proto", fileDescriptor_session_004cdd49a5f4cc5f)
}

var fileDescriptor_session_004cdd49a5f4cc5f = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x69, 0x1a, 0xe2, 0x49, 0x13, 0xc1, 0x0a, 0x82, 0xb1, 0x50, 0x15, 0xf9, 0x80, 0xc2,
	0x21, 0x09, 0x4a, 0x81, 0x03, 0x14, 0x55, 0xa2, 0x17, 0x90, 0x7a, 0x40, 0x1b, 0xb8, 0x71, 0xd9,
	0xb8, 0x53, 0x67, 0x45, 0xed, 0x35, 0xde, 0x75, 0xc4, 0x3f, 0xe2, 0xc7, 0xf0, 0xa7, 0x90, 0xd7,
	0x5e, 0x77, 0xed, 0x46, 0x4a, 0x38, 0xe4, 0x30, 0x6f, 0xe7, 0xe3, 0xcd, 0x7b, 0x13, 0xc3, 0xc7,
	0x88, 0xab, 0x4d, 0xbe, 0x9e, 0x87, 0x22, 0x5e, 0x6c, 0xd8, 0x56, 0x84, 0x33, 0x2e, 0x16, 0x71,
	0xae, 0x58, 0x84, 0xc9, 0x22, 0xfd, 0x19, 0x2d, 0x24, 0x4a, 0xc9, 0x45, 0xb2, 0x90, 0x98, 0x6d,
	0x79, 0x88, 0x26, 0x9e, 0xa7, 0x99, 0x50, 0x82, 0x3c, 0xac, 0x42, 0xff, 0xc3, 0xa1, 0x7d, 0x42,
	0x91, 0xdc, 0xf0, 0x28, 0xcf, 0x98, 0xaa, 0xbb, 0xf8, 0x67, 0x07, 0x93, 0x50, 0x4c, 0x61, 0x55,
	0x34, 0xdb, 0x57, 0x94, 0x67, 0xb7, 0xc5, 0xaf, 0x4c, 0x0f, 0xfe, 0x38, 0x30, 0xbc, 0xcc, 0x90,
	0x29, 0xa4, 0xf8, 0x2b, 0x47, 0xa9, 0xc8, 0x29, 0x1c, 0xb3, 0xdb, 0x74, 0xc3, 0x3c, 0x67, 0xe2,
	0x4c, 0x07, 0xcb, 0xfe, 0xbc, 0x48, 0xfe, 0x4e, 0xaf, 0x68, 0x09, 0x93, 0x17, 0xd0, 0x5d, 0xa3,
	0x62, 0x5e, 0xa7, 0xf5, 0xac, 0x51, 0x72, 0x0e, 0xc3, 0xc6, 0x2a, 0xde, 0x91, 0x4e, 0x1b, 0xcf,
	0x8d, 0x40, 0x97, 0xf6, 0x2b, 0x6d, 0x26, 0x13, 0x1f, 0xfa, 0x19, 0xca, 0x54, 0x24, 0x12, 0xbd,
	0xee, 0xc4, 0x99, 0xba, 0xb4, 0x8e, 0x83, 0x1f, 0x30, 0x32, 0x44, 0x4b, 0x84, 0x78, 0x60, 0x74,
	0xd6, 0x5c, 0x5d, 0x6a, 0xc2, 0xe2, 0x25, 0x46, 0x29, 0x59, 0x84, 0x9a, 0xa6, 0x4b, 0x4d, 0x48,
	0xc6, 0xd0, 0x4b, 0x33, 0x11, 0xa7, 0x4a, 0x13, 0x73, 0x69, 0x15, 0x05, 0x08, 0x83, 0x2b, 0x2e,
	0x95, 0x11, 0x61, 0x0e, 0x24, 0xcd, 0x70, 0xcb, 0x45, 0x2e, 0x57, 0x85, 0xb8, 0x5f, 0x92, 0x6b,
	0xfc, 0xad, 0xa7, 0x74, 0xe9, 0x8e, 0x17, 0xf2, 0x12, 0x46, 0x32, 0xc5, 0x90, 0xdf, 0xf0, 0x50,
	0x6f, 0x22, 0xbd, 0xce, 0xe4, 0x68, 0xea, 0xd2, 0x16, 0x1a, 0x5c, 0xc3, 0x49, 0x39, 0xa6, 0x5a,
	0xe1, 0x14, 0x40, 0xb6, 0xfb, 0x5b, 0x08, 0x79, 0x03, 0xc3, 0x6a, 0x27, 0x3d, 0xac, 0x6c, 0x3b,
	0x58, 0x8e, 0x6a, 0x39, 0x35, 0x4c, 0x9b, 0x49, 0xc1, 0x3b, 0x38, 0xf9, 0xca, 0x72, 0x59, 0x5b,
	0x7a, 0x9f, 0x9d, 0xb3, 0x93, 0xdd, 0x2b, 0x18, 0x56, 0x75, 0x77, 0x0a, 0x1b, 0x1d, 0x9d, 0x86,
	0x8e, 0xc1, 0x0a, 0x86, 0x14, 0x65, 0x1e, 0xff, 0xef, 0x8c, 0x86, 0xc5, 0x9d, 0x96, 0xc5, 0x9f,
	0x60, 0x64, 0x9a, 0xee, 0x23, 0x60, 0x19, 0xd9, 0x69, 0x18, 0xf9, 0x1e, 0x1e, 0x7d, 0xc3, 0x2c,
	0xe6, 0x89, 0x75, 0xd2, 0x87, 0xee, 0x3f, 0x83, 0xc7, 0x56, 0xed, 0x3e, 0x0a, 0xcb, 0xbf, 0x1d,
	0xe8, 0xaf, 0x4a, 0xe1, 0x25, 0xb9, 0x80, 0x5e, 0x79, 0x9e, 0xc4, 0xba, 0x75, 0xfb, 0x8f, 0xe5,
	0x3f, 0xbb, 0x87, 0x57, 0x6b, 0x3f, 0x98, 0x3a, 0xaf, 0x1d, 0xf2, 0x16, 0xba, 0xc5, 0x69, 0x90,
	0x27, 0x75, 0x9a, 0x75, 0x90, 0xfe, 0xd3, 0x16, 0x6a, 0x4a, 0xc9, 0x39, 0x1c, 0x6b, 0xcf, 0xc8,
	0x5d, 0x86, 0xed, 0xbd, 0x3f, 0x6e, 0xc3, 0x8d, 0xa1, 0x17, 0xd0, 0x2b, 0x15, 0xb7, 0x58, 0x37,
	0x7c, 0xb5, 0x58, 0x37, 0xad, 0xa9, 0x1a, 0x7c, 0x06, 0xb7, 0x96, 0x8c, 0x3c, 0xaf, 0x73, 0xdb,
	0x16, 0xf8, 0xfe, 0xae, 0x27, 0xbb, 0xd3, 0xba, 0xa7, 0x3f, 0x48, 0x67, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xa6, 0x62, 0xfb, 0x15, 0x7b, 0x05, 0x00, 0x00,
}
