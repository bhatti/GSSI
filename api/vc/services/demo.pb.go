// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/vc/services/demo.proto

package services

import (
	types "github.com/bhatti/GSSI/api/vc/types"
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

// GetAllDegreesRequest for verifiable credentials
type GetAllDegreesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Query string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *GetAllDegreesRequest) Reset() {
	*x = GetAllDegreesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_vc_services_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDegreesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDegreesRequest) ProtoMessage() {}

func (x *GetAllDegreesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_vc_services_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDegreesRequest.ProtoReflect.Descriptor instead.
func (*GetAllDegreesRequest) Descriptor() ([]byte, []int) {
	return file_api_vc_services_demo_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllDegreesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllDegreesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllDegreesRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

// GetDegreeRequest for verifiable credentials
type GetDegreeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDegreeRequest) Reset() {
	*x = GetDegreeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_vc_services_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDegreeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDegreeRequest) ProtoMessage() {}

func (x *GetDegreeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_vc_services_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDegreeRequest.ProtoReflect.Descriptor instead.
func (*GetDegreeRequest) Descriptor() ([]byte, []int) {
	return file_api_vc_services_demo_proto_rawDescGZIP(), []int{1}
}

func (x *GetDegreeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteDegreeRequest for verifiable credentials
type DeleteDegreeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDegreeRequest) Reset() {
	*x = DeleteDegreeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_vc_services_demo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDegreeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDegreeRequest) ProtoMessage() {}

func (x *DeleteDegreeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_vc_services_demo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDegreeRequest.ProtoReflect.Descriptor instead.
func (*DeleteDegreeRequest) Descriptor() ([]byte, []int) {
	return file_api_vc_services_demo_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteDegreeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// CreateDegreeRequest for verifiable credentials
type CreateDegreeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateDegreeRequest) Reset() {
	*x = CreateDegreeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_vc_services_demo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDegreeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDegreeRequest) ProtoMessage() {}

func (x *CreateDegreeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_vc_services_demo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDegreeRequest.ProtoReflect.Descriptor instead.
func (*CreateDegreeRequest) Descriptor() ([]byte, []int) {
	return file_api_vc_services_demo_proto_rawDescGZIP(), []int{3}
}

var File_api_vc_services_demo_proto protoreflect.FileDescriptor

var file_api_vc_services_demo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x19, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x63,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22,
	0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x67,
	0x72, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x32, 0xf5, 0x02, 0x0a, 0x0d, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x06, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x25, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x22, 0x00, 0x30, 0x01, 0x12, 0x59, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x21,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x67, 0x72, 0x65,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x63, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x61,
	0x62, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x68, 0x61, 0x74, 0x74, 0x69, 0x2f, 0x47,
	0x53, 0x53, 0x49, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_vc_services_demo_proto_rawDescOnce sync.Once
	file_api_vc_services_demo_proto_rawDescData = file_api_vc_services_demo_proto_rawDesc
)

func file_api_vc_services_demo_proto_rawDescGZIP() []byte {
	file_api_vc_services_demo_proto_rawDescOnce.Do(func() {
		file_api_vc_services_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_vc_services_demo_proto_rawDescData)
	})
	return file_api_vc_services_demo_proto_rawDescData
}

var file_api_vc_services_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_vc_services_demo_proto_goTypes = []interface{}{
	(*GetAllDegreesRequest)(nil),                  // 0: api.vc.services.GetAllDegreesRequest
	(*GetDegreeRequest)(nil),                      // 1: api.vc.services.GetDegreeRequest
	(*DeleteDegreeRequest)(nil),                   // 2: api.vc.services.DeleteDegreeRequest
	(*CreateDegreeRequest)(nil),                   // 3: api.vc.services.CreateDegreeRequest
	(*types.RefreshableVerifiableCredential)(nil), // 4: api.vc.types.RefreshableVerifiableCredential
	(*types.Bool)(nil),                            // 5: api.vc.types.Bool
}
var file_api_vc_services_demo_proto_depIdxs = []int32{
	0, // 0: api.vc.services.DegreeService.getAll:input_type -> api.vc.services.GetAllDegreesRequest
	1, // 1: api.vc.services.DegreeService.get:input_type -> api.vc.services.GetDegreeRequest
	2, // 2: api.vc.services.DegreeService.delete:input_type -> api.vc.services.DeleteDegreeRequest
	3, // 3: api.vc.services.DegreeService.create:input_type -> api.vc.services.CreateDegreeRequest
	4, // 4: api.vc.services.DegreeService.getAll:output_type -> api.vc.types.RefreshableVerifiableCredential
	4, // 5: api.vc.services.DegreeService.get:output_type -> api.vc.types.RefreshableVerifiableCredential
	5, // 6: api.vc.services.DegreeService.delete:output_type -> api.vc.types.Bool
	4, // 7: api.vc.services.DegreeService.create:output_type -> api.vc.types.RefreshableVerifiableCredential
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_vc_services_demo_proto_init() }
func file_api_vc_services_demo_proto_init() {
	if File_api_vc_services_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_vc_services_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDegreesRequest); i {
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
		file_api_vc_services_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDegreeRequest); i {
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
		file_api_vc_services_demo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDegreeRequest); i {
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
		file_api_vc_services_demo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDegreeRequest); i {
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
			RawDescriptor: file_api_vc_services_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_vc_services_demo_proto_goTypes,
		DependencyIndexes: file_api_vc_services_demo_proto_depIdxs,
		MessageInfos:      file_api_vc_services_demo_proto_msgTypes,
	}.Build()
	File_api_vc_services_demo_proto = out.File
	file_api_vc_services_demo_proto_rawDesc = nil
	file_api_vc_services_demo_proto_goTypes = nil
	file_api_vc_services_demo_proto_depIdxs = nil
}
