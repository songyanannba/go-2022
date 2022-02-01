// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: food.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FoodStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FoodStreamRequest) Reset() {
	*x = FoodStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodStreamRequest) ProtoMessage() {}

func (x *FoodStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodStreamRequest.ProtoReflect.Descriptor instead.
func (*FoodStreamRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{0}
}

func (x *FoodStreamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FoodStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *FoodStreamResponse) Reset() {
	*x = FoodStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodStreamResponse) ProtoMessage() {}

func (x *FoodStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodStreamResponse.ProtoReflect.Descriptor instead.
func (*FoodStreamResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{1}
}

func (x *FoodStreamResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_food_proto protoreflect.FileDescriptor

var file_food_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x11,
	0x46, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x12, 0x46, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xb5, 0x01,
	0x0a, 0x0b, 0x46, 0x6f, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a,
	0x07, 0x53, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x46,
	0x6f, 0x6f, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x12, 0x35, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x39, 0x0a, 0x0a, 0x46, 0x75,
	0x6c, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x12, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x46,
	0x6f, 0x6f, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_food_proto_rawDescOnce sync.Once
	file_food_proto_rawDescData = file_food_proto_rawDesc
)

func file_food_proto_rawDescGZIP() []byte {
	file_food_proto_rawDescOnce.Do(func() {
		file_food_proto_rawDescData = protoimpl.X.CompressGZIP(file_food_proto_rawDescData)
	})
	return file_food_proto_rawDescData
}

var file_food_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_food_proto_goTypes = []interface{}{
	(*FoodStreamRequest)(nil),  // 0: FoodStreamRequest
	(*FoodStreamResponse)(nil), // 1: FoodStreamResponse
}
var file_food_proto_depIdxs = []int32{
	0, // 0: FoodService.SayName:input_type -> FoodStreamRequest
	0, // 1: FoodService.PostName:input_type -> FoodStreamRequest
	0, // 2: FoodService.FullStream:input_type -> FoodStreamRequest
	1, // 3: FoodService.SayName:output_type -> FoodStreamResponse
	1, // 4: FoodService.PostName:output_type -> FoodStreamResponse
	1, // 5: FoodService.FullStream:output_type -> FoodStreamResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_food_proto_init() }
func file_food_proto_init() {
	if File_food_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_food_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodStreamRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_food_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodStreamResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_food_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_food_proto_goTypes,
		DependencyIndexes: file_food_proto_depIdxs,
		MessageInfos:      file_food_proto_msgTypes,
	}.Build()
	File_food_proto = out.File
	file_food_proto_rawDesc = nil
	file_food_proto_goTypes = nil
	file_food_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FoodServiceClient is the client API for FoodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FoodServiceClient interface {
	SayName(ctx context.Context, in *FoodStreamRequest, opts ...grpc.CallOption) (FoodService_SayNameClient, error)
	PostName(ctx context.Context, opts ...grpc.CallOption) (FoodService_PostNameClient, error)
	FullStream(ctx context.Context, opts ...grpc.CallOption) (FoodService_FullStreamClient, error)
}

type foodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFoodServiceClient(cc grpc.ClientConnInterface) FoodServiceClient {
	return &foodServiceClient{cc}
}

func (c *foodServiceClient) SayName(ctx context.Context, in *FoodStreamRequest, opts ...grpc.CallOption) (FoodService_SayNameClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FoodService_serviceDesc.Streams[0], "/FoodService/SayName", opts...)
	if err != nil {
		return nil, err
	}
	x := &foodServiceSayNameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FoodService_SayNameClient interface {
	Recv() (*FoodStreamResponse, error)
	grpc.ClientStream
}

type foodServiceSayNameClient struct {
	grpc.ClientStream
}

func (x *foodServiceSayNameClient) Recv() (*FoodStreamResponse, error) {
	m := new(FoodStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *foodServiceClient) PostName(ctx context.Context, opts ...grpc.CallOption) (FoodService_PostNameClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FoodService_serviceDesc.Streams[1], "/FoodService/PostName", opts...)
	if err != nil {
		return nil, err
	}
	x := &foodServicePostNameClient{stream}
	return x, nil
}

type FoodService_PostNameClient interface {
	Send(*FoodStreamRequest) error
	CloseAndRecv() (*FoodStreamResponse, error)
	grpc.ClientStream
}

type foodServicePostNameClient struct {
	grpc.ClientStream
}

func (x *foodServicePostNameClient) Send(m *FoodStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *foodServicePostNameClient) CloseAndRecv() (*FoodStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FoodStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *foodServiceClient) FullStream(ctx context.Context, opts ...grpc.CallOption) (FoodService_FullStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FoodService_serviceDesc.Streams[2], "/FoodService/FullStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &foodServiceFullStreamClient{stream}
	return x, nil
}

type FoodService_FullStreamClient interface {
	Send(*FoodStreamRequest) error
	Recv() (*FoodStreamResponse, error)
	grpc.ClientStream
}

type foodServiceFullStreamClient struct {
	grpc.ClientStream
}

func (x *foodServiceFullStreamClient) Send(m *FoodStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *foodServiceFullStreamClient) Recv() (*FoodStreamResponse, error) {
	m := new(FoodStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FoodServiceServer is the server API for FoodService service.
type FoodServiceServer interface {
	SayName(*FoodStreamRequest, FoodService_SayNameServer) error
	PostName(FoodService_PostNameServer) error
	FullStream(FoodService_FullStreamServer) error
}

// UnimplementedFoodServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFoodServiceServer struct {
}

func (*UnimplementedFoodServiceServer) SayName(*FoodStreamRequest, FoodService_SayNameServer) error {
	return status.Errorf(codes.Unimplemented, "method SayName not implemented")
}
func (*UnimplementedFoodServiceServer) PostName(FoodService_PostNameServer) error {
	return status.Errorf(codes.Unimplemented, "method PostName not implemented")
}
func (*UnimplementedFoodServiceServer) FullStream(FoodService_FullStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method FullStream not implemented")
}

func RegisterFoodServiceServer(s *grpc.Server, srv FoodServiceServer) {
	s.RegisterService(&_FoodService_serviceDesc, srv)
}

func _FoodService_SayName_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FoodStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FoodServiceServer).SayName(m, &foodServiceSayNameServer{stream})
}

type FoodService_SayNameServer interface {
	Send(*FoodStreamResponse) error
	grpc.ServerStream
}

type foodServiceSayNameServer struct {
	grpc.ServerStream
}

func (x *foodServiceSayNameServer) Send(m *FoodStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FoodService_PostName_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FoodServiceServer).PostName(&foodServicePostNameServer{stream})
}

type FoodService_PostNameServer interface {
	SendAndClose(*FoodStreamResponse) error
	Recv() (*FoodStreamRequest, error)
	grpc.ServerStream
}

type foodServicePostNameServer struct {
	grpc.ServerStream
}

func (x *foodServicePostNameServer) SendAndClose(m *FoodStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *foodServicePostNameServer) Recv() (*FoodStreamRequest, error) {
	m := new(FoodStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FoodService_FullStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FoodServiceServer).FullStream(&foodServiceFullStreamServer{stream})
}

type FoodService_FullStreamServer interface {
	Send(*FoodStreamResponse) error
	Recv() (*FoodStreamRequest, error)
	grpc.ServerStream
}

type foodServiceFullStreamServer struct {
	grpc.ServerStream
}

func (x *foodServiceFullStreamServer) Send(m *FoodStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *foodServiceFullStreamServer) Recv() (*FoodStreamRequest, error) {
	m := new(FoodStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FoodService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FoodService",
	HandlerType: (*FoodServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayName",
			Handler:       _FoodService_SayName_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PostName",
			Handler:       _FoodService_PostName_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FullStream",
			Handler:       _FoodService_FullStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "food.proto",
}
