// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo/todo.proto

package todo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Todo struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Complete             bool     `protobuf:"varint,3,opt,name=complete,proto3" json:"complete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b7148f81fa63801, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Todo) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Todo) GetComplete() bool {
	if m != nil {
		return m.Complete
	}
	return false
}

type ListTodoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoRequest) Reset()         { *m = ListTodoRequest{} }
func (m *ListTodoRequest) String() string { return proto.CompactTextString(m) }
func (*ListTodoRequest) ProtoMessage()    {}
func (*ListTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b7148f81fa63801, []int{1}
}

func (m *ListTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoRequest.Unmarshal(m, b)
}
func (m *ListTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoRequest.Marshal(b, m, deterministic)
}
func (m *ListTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoRequest.Merge(m, src)
}
func (m *ListTodoRequest) XXX_Size() int {
	return xxx_messageInfo_ListTodoRequest.Size(m)
}
func (m *ListTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoRequest proto.InternalMessageInfo

type ListTodoResponse struct {
	Todos                []*Todo  `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoResponse) Reset()         { *m = ListTodoResponse{} }
func (m *ListTodoResponse) String() string { return proto.CompactTextString(m) }
func (*ListTodoResponse) ProtoMessage()    {}
func (*ListTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b7148f81fa63801, []int{2}
}

func (m *ListTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoResponse.Unmarshal(m, b)
}
func (m *ListTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoResponse.Marshal(b, m, deterministic)
}
func (m *ListTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoResponse.Merge(m, src)
}
func (m *ListTodoResponse) XXX_Size() int {
	return xxx_messageInfo_ListTodoResponse.Size(m)
}
func (m *ListTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoResponse proto.InternalMessageInfo

func (m *ListTodoResponse) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

func init() {
	proto.RegisterType((*Todo)(nil), "todo.Todo")
	proto.RegisterType((*ListTodoRequest)(nil), "todo.ListTodoRequest")
	proto.RegisterType((*ListTodoResponse)(nil), "todo.ListTodoResponse")
}

func init() { proto.RegisterFile("todo/todo.proto", fileDescriptor_7b7148f81fa63801) }

var fileDescriptor_7b7148f81fa63801 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x50, 0xc1, 0x4a, 0x03, 0x31,
	0x14, 0x24, 0xdb, 0x6d, 0x59, 0x5f, 0xa1, 0xd5, 0x07, 0xca, 0xb2, 0x78, 0x08, 0x7b, 0xda, 0x53,
	0x17, 0xaa, 0xdf, 0x20, 0x08, 0x9e, 0x56, 0x0f, 0x5e, 0xd7, 0xe6, 0x51, 0x02, 0x35, 0x2f, 0x36,
	0x4f, 0xf1, 0xec, 0x2f, 0xf8, 0x69, 0xfe, 0x82, 0x1f, 0x22, 0x49, 0x50, 0x61, 0x2f, 0x61, 0x66,
	0x32, 0xcc, 0x4c, 0x02, 0x6b, 0x61, 0xc3, 0x7d, 0x3c, 0x36, 0xfe, 0xc8, 0xc2, 0x58, 0x46, 0xdc,
	0x5c, 0xee, 0x99, 0xf7, 0x07, 0xea, 0x47, 0x6f, 0xfb, 0xd1, 0x39, 0x96, 0x51, 0x2c, 0xbb, 0x90,
	0x3d, 0xed, 0x0d, 0x94, 0x0f, 0x6c, 0x18, 0x57, 0x50, 0x58, 0x53, 0x2b, 0xad, 0xba, 0xf9, 0x50,
	0x58, 0x83, 0x08, 0xa5, 0xd0, 0xbb, 0xd4, 0x85, 0x56, 0xdd, 0xc9, 0x90, 0x30, 0x36, 0x50, 0xed,
	0xf8, 0xd9, 0x1f, 0x48, 0xa8, 0x9e, 0x69, 0xd5, 0x55, 0xc3, 0x1f, 0x6f, 0xcf, 0x60, 0x7d, 0x67,
	0x83, 0xc4, 0xac, 0x81, 0x5e, 0x5e, 0x29, 0x48, 0x7b, 0x0d, 0xa7, 0xff, 0x52, 0xf0, 0xec, 0x02,
	0xa1, 0x86, 0x79, 0x1c, 0x15, 0x6a, 0xa5, 0x67, 0xdd, 0x72, 0x0b, 0x9b, 0x34, 0x37, 0x59, 0xf2,
	0xc5, 0xf6, 0x11, 0x96, 0x91, 0xde, 0xd3, 0xf1, 0xcd, 0xee, 0x08, 0x6f, 0xa1, 0xfa, 0x0d, 0xc1,
	0xf3, 0xec, 0x9e, 0xf4, 0x34, 0x17, 0x53, 0x39, 0x77, 0xb5, 0xab, 0x8f, 0xaf, 0xef, 0xcf, 0xa2,
	0xc2, 0x45, 0xfa, 0x93, 0xf0, 0xb4, 0x48, 0x2f, 0xbe, 0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf9,
	0xee, 0x11, 0x89, 0x28, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	ListTodo(ctx context.Context, in *ListTodoRequest, opts ...grpc.CallOption) (*ListTodoResponse, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) ListTodo(ctx context.Context, in *ListTodoRequest, opts ...grpc.CallOption) (*ListTodoResponse, error) {
	out := new(ListTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/ListTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	ListTodo(context.Context, *ListTodoRequest) (*ListTodoResponse, error)
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_ListTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ListTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/ListTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ListTodo(ctx, req.(*ListTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTodo",
			Handler:    _TodoService_ListTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo/todo.proto",
}
