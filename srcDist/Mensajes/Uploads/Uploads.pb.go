// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: Uploads.proto

package Uploads

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Part    int32  `protobuf:"varint,3,opt,name=Part,proto3" json:"Part,omitempty"`
	Puerto  string `protobuf:"bytes,4,opt,name=Puerto,proto3" json:"Puerto,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Uploads_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_Uploads_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_Uploads_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Chunk) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Chunk) GetPart() int32 {
	if x != nil {
		return x.Part
	}
	return 0
}

func (x *Chunk) GetPuerto() string {
	if x != nil {
		return x.Puerto
	}
	return ""
}

type UploadStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *UploadStatus) Reset() {
	*x = UploadStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Uploads_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadStatus) ProtoMessage() {}

func (x *UploadStatus) ProtoReflect() protoreflect.Message {
	mi := &file_Uploads_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadStatus.ProtoReflect.Descriptor instead.
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return file_Uploads_proto_rawDescGZIP(), []int{1}
}

func (x *UploadStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_Uploads_proto protoreflect.FileDescriptor

var file_Uploads_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x22, 0x61, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50,
	0x61, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x75, 0x65, 0x72, 0x74, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x75, 0x65, 0x72, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0c, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x45, 0x0a, 0x0e, 0x47, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x0e, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x1a, 0x15, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x28, 0x01, 0x32, 0x48, 0x0a, 0x0f,
	0x52, 0x65, 0x70, 0x61, 0x72, 0x74, 0x69, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x35, 0x0a, 0x08, 0x52, 0x65, 0x70, 0x61, 0x72, 0x74, 0x69, 0x72, 0x12, 0x0e, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x15, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x00, 0x28, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Uploads_proto_rawDescOnce sync.Once
	file_Uploads_proto_rawDescData = file_Uploads_proto_rawDesc
)

func file_Uploads_proto_rawDescGZIP() []byte {
	file_Uploads_proto_rawDescOnce.Do(func() {
		file_Uploads_proto_rawDescData = protoimpl.X.CompressGZIP(file_Uploads_proto_rawDescData)
	})
	return file_Uploads_proto_rawDescData
}

var file_Uploads_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Uploads_proto_goTypes = []interface{}{
	(*Chunk)(nil),        // 0: Uploads.Chunk
	(*UploadStatus)(nil), // 1: Uploads.UploadStatus
}
var file_Uploads_proto_depIdxs = []int32{
	0, // 0: Uploads.GuploadService.Upload:input_type -> Uploads.Chunk
	0, // 1: Uploads.RepartirService.Repartir:input_type -> Uploads.Chunk
	1, // 2: Uploads.GuploadService.Upload:output_type -> Uploads.UploadStatus
	1, // 3: Uploads.RepartirService.Repartir:output_type -> Uploads.UploadStatus
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Uploads_proto_init() }
func file_Uploads_proto_init() {
	if File_Uploads_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Uploads_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_Uploads_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadStatus); i {
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
			RawDescriptor: file_Uploads_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_Uploads_proto_goTypes,
		DependencyIndexes: file_Uploads_proto_depIdxs,
		MessageInfos:      file_Uploads_proto_msgTypes,
	}.Build()
	File_Uploads_proto = out.File
	file_Uploads_proto_rawDesc = nil
	file_Uploads_proto_goTypes = nil
	file_Uploads_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GuploadServiceClient is the client API for GuploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GuploadServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (GuploadService_UploadClient, error)
}

type guploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGuploadServiceClient(cc grpc.ClientConnInterface) GuploadServiceClient {
	return &guploadServiceClient{cc}
}

func (c *guploadServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (GuploadService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GuploadService_serviceDesc.Streams[0], "/Uploads.GuploadService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &guploadServiceUploadClient{stream}
	return x, nil
}

type GuploadService_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type guploadServiceUploadClient struct {
	grpc.ClientStream
}

func (x *guploadServiceUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *guploadServiceUploadClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GuploadServiceServer is the server API for GuploadService service.
type GuploadServiceServer interface {
	Upload(GuploadService_UploadServer) error
}

// UnimplementedGuploadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGuploadServiceServer struct {
}

func (*UnimplementedGuploadServiceServer) Upload(GuploadService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterGuploadServiceServer(s *grpc.Server, srv GuploadServiceServer) {
	s.RegisterService(&_GuploadService_serviceDesc, srv)
}

func _GuploadService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GuploadServiceServer).Upload(&guploadServiceUploadServer{stream})
}

type GuploadService_UploadServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type guploadServiceUploadServer struct {
	grpc.ServerStream
}

func (x *guploadServiceUploadServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *guploadServiceUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GuploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Uploads.GuploadService",
	HandlerType: (*GuploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _GuploadService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "Uploads.proto",
}

// RepartirServiceClient is the client API for RepartirService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RepartirServiceClient interface {
	Repartir(ctx context.Context, opts ...grpc.CallOption) (RepartirService_RepartirClient, error)
}

type repartirServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRepartirServiceClient(cc grpc.ClientConnInterface) RepartirServiceClient {
	return &repartirServiceClient{cc}
}

func (c *repartirServiceClient) Repartir(ctx context.Context, opts ...grpc.CallOption) (RepartirService_RepartirClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RepartirService_serviceDesc.Streams[0], "/Uploads.RepartirService/Repartir", opts...)
	if err != nil {
		return nil, err
	}
	x := &repartirServiceRepartirClient{stream}
	return x, nil
}

type RepartirService_RepartirClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type repartirServiceRepartirClient struct {
	grpc.ClientStream
}

func (x *repartirServiceRepartirClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *repartirServiceRepartirClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RepartirServiceServer is the server API for RepartirService service.
type RepartirServiceServer interface {
	Repartir(RepartirService_RepartirServer) error
}

// UnimplementedRepartirServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRepartirServiceServer struct {
}

func (*UnimplementedRepartirServiceServer) Repartir(RepartirService_RepartirServer) error {
	return status.Errorf(codes.Unimplemented, "method Repartir not implemented")
}

func RegisterRepartirServiceServer(s *grpc.Server, srv RepartirServiceServer) {
	s.RegisterService(&_RepartirService_serviceDesc, srv)
}

func _RepartirService_Repartir_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RepartirServiceServer).Repartir(&repartirServiceRepartirServer{stream})
}

type RepartirService_RepartirServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type repartirServiceRepartirServer struct {
	grpc.ServerStream
}

func (x *repartirServiceRepartirServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *repartirServiceRepartirServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RepartirService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Uploads.RepartirService",
	HandlerType: (*RepartirServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Repartir",
			Handler:       _RepartirService_Repartir_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "Uploads.proto",
}
