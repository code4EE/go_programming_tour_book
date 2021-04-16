// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: proto/tag.proto

package proto

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

type GetTagListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State uint32 `protobuf:"varint,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *GetTagListRequest) Reset() {
	*x = GetTagListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagListRequest) ProtoMessage() {}

func (x *GetTagListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagListRequest.ProtoReflect.Descriptor instead.
func (*GetTagListRequest) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{0}
}

func (x *GetTagListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetTagListRequest) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State uint32 `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{1}
}

func (x *Tag) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tag) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

type GetTagListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*Tag `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Pager *Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
}

func (x *GetTagListReply) Reset() {
	*x = GetTagListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagListReply) ProtoMessage() {}

func (x *GetTagListReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagListReply.ProtoReflect.Descriptor instead.
func (*GetTagListReply) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{2}
}

func (x *GetTagListReply) GetList() []*Tag {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *GetTagListReply) GetPager() *Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

var File_proto_tag_proto protoreflect.FileDescriptor

var file_proto_tag_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x3f, 0x0a, 0x03, 0x54,
	0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x55, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x1e, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x32, 0x4e, 0x0a, 0x0a, 0x54, 0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x40, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_tag_proto_rawDescOnce sync.Once
	file_proto_tag_proto_rawDescData = file_proto_tag_proto_rawDesc
)

func file_proto_tag_proto_rawDescGZIP() []byte {
	file_proto_tag_proto_rawDescOnce.Do(func() {
		file_proto_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_tag_proto_rawDescData)
	})
	return file_proto_tag_proto_rawDescData
}

var file_proto_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_tag_proto_goTypes = []interface{}{
	(*GetTagListRequest)(nil), // 0: proto.GetTagListRequest
	(*Tag)(nil),               // 1: proto.Tag
	(*GetTagListReply)(nil),   // 2: proto.GetTagListReply
	(*Pager)(nil),             // 3: proto.Pager
}
var file_proto_tag_proto_depIdxs = []int32{
	1, // 0: proto.GetTagListReply.list:type_name -> proto.Tag
	3, // 1: proto.GetTagListReply.pager:type_name -> proto.Pager
	0, // 2: proto.TagService.GetTagList:input_type -> proto.GetTagListRequest
	2, // 3: proto.TagService.GetTagList:output_type -> proto.GetTagListReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_tag_proto_init() }
func file_proto_tag_proto_init() {
	if File_proto_tag_proto != nil {
		return
	}
	file_proto_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagListRequest); i {
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
		file_proto_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_proto_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagListReply); i {
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
			RawDescriptor: file_proto_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_tag_proto_goTypes,
		DependencyIndexes: file_proto_tag_proto_depIdxs,
		MessageInfos:      file_proto_tag_proto_msgTypes,
	}.Build()
	File_proto_tag_proto = out.File
	file_proto_tag_proto_rawDesc = nil
	file_proto_tag_proto_goTypes = nil
	file_proto_tag_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TagServiceClient interface {
	GetTagList(ctx context.Context, in *GetTagListRequest, opts ...grpc.CallOption) (*GetTagListReply, error)
}

type tagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagServiceClient(cc grpc.ClientConnInterface) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) GetTagList(ctx context.Context, in *GetTagListRequest, opts ...grpc.CallOption) (*GetTagListReply, error) {
	out := new(GetTagListReply)
	err := c.cc.Invoke(ctx, "/proto.TagService/GetTagList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
type TagServiceServer interface {
	GetTagList(context.Context, *GetTagListRequest) (*GetTagListReply, error)
}

// UnimplementedTagServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTagServiceServer struct {
}

func (*UnimplementedTagServiceServer) GetTagList(context.Context, *GetTagListRequest) (*GetTagListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagList not implemented")
}

func RegisterTagServiceServer(s *grpc.Server, srv TagServiceServer) {
	s.RegisterService(&_TagService_serviceDesc, srv)
}

func _TagService_GetTagList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).GetTagList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TagService/GetTagList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).GetTagList(ctx, req.(*GetTagListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TagService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTagList",
			Handler:    _TagService_GetTagList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tag.proto",
}
