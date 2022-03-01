// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: contact_service.proto

package contact_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_contact_service_proto protoreflect.FileDescriptor

var file_contact_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xc2, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x1a, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x49, 0x64, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x12, 0x1e, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64,
	0x1a, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_contact_service_proto_goTypes = []interface{}{
	(*Contact)(nil),               // 0: genproto.Contact
	(*GetAllContactRequest)(nil),  // 1: genproto.GetAllContactRequest
	(*ContactId)(nil),             // 2: genproto.ContactId
	(*GetAllContactResponse)(nil), // 3: genproto.GetAllContactResponse
}
var file_contact_service_proto_depIdxs = []int32{
	0, // 0: genproto.ContactService.Create:input_type -> genproto.Contact
	1, // 1: genproto.ContactService.GetAll:input_type -> genproto.GetAllContactRequest
	2, // 2: genproto.ContactService.Get:input_type -> genproto.ContactId
	2, // 3: genproto.ContactService.Create:output_type -> genproto.ContactId
	3, // 4: genproto.ContactService.GetAll:output_type -> genproto.GetAllContactResponse
	0, // 5: genproto.ContactService.Get:output_type -> genproto.Contact
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_contact_service_proto_init() }
func file_contact_service_proto_init() {
	if File_contact_service_proto != nil {
		return
	}
	file_contact_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contact_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contact_service_proto_goTypes,
		DependencyIndexes: file_contact_service_proto_depIdxs,
	}.Build()
	File_contact_service_proto = out.File
	file_contact_service_proto_rawDesc = nil
	file_contact_service_proto_goTypes = nil
	file_contact_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ContactServiceClient is the client API for ContactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContactServiceClient interface {
	Create(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*ContactId, error)
	GetAll(ctx context.Context, in *GetAllContactRequest, opts ...grpc.CallOption) (*GetAllContactResponse, error)
	Get(ctx context.Context, in *ContactId, opts ...grpc.CallOption) (*Contact, error)
}

type contactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContactServiceClient(cc grpc.ClientConnInterface) ContactServiceClient {
	return &contactServiceClient{cc}
}

func (c *contactServiceClient) Create(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*ContactId, error) {
	out := new(ContactId)
	err := c.cc.Invoke(ctx, "/genproto.ContactService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) GetAll(ctx context.Context, in *GetAllContactRequest, opts ...grpc.CallOption) (*GetAllContactResponse, error) {
	out := new(GetAllContactResponse)
	err := c.cc.Invoke(ctx, "/genproto.ContactService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) Get(ctx context.Context, in *ContactId, opts ...grpc.CallOption) (*Contact, error) {
	out := new(Contact)
	err := c.cc.Invoke(ctx, "/genproto.ContactService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactServiceServer is the server API for ContactService service.
type ContactServiceServer interface {
	Create(context.Context, *Contact) (*ContactId, error)
	GetAll(context.Context, *GetAllContactRequest) (*GetAllContactResponse, error)
	Get(context.Context, *ContactId) (*Contact, error)
}

// UnimplementedContactServiceServer can be embedded to have forward compatible implementations.
type UnimplementedContactServiceServer struct {
}

func (*UnimplementedContactServiceServer) Create(context.Context, *Contact) (*ContactId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedContactServiceServer) GetAll(context.Context, *GetAllContactRequest) (*GetAllContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedContactServiceServer) Get(context.Context, *ContactId) (*Contact, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterContactServiceServer(s *grpc.Server, srv ContactServiceServer) {
	s.RegisterService(&_ContactService_serviceDesc, srv)
}

func _ContactService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Contact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ContactService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Create(ctx, req.(*Contact))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ContactService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).GetAll(ctx, req.(*GetAllContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContactId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ContactService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Get(ctx, req.(*ContactId))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContactService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.ContactService",
	HandlerType: (*ContactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ContactService_Create_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ContactService_GetAll_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ContactService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contact_service.proto",
}
