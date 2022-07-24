// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/ld/types/ld.proto

package types

import (
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

// ProviderID is a request/response model for operations that involve remote provider ID.
type ProviderID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProviderID) Reset() {
	*x = ProviderID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ld_types_ld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderID) ProtoMessage() {}

func (x *ProviderID) ProtoReflect() protoreflect.Message {
	mi := &file_api_ld_types_ld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderID.ProtoReflect.Descriptor instead.
func (*ProviderID) Descriptor() ([]byte, []int) {
	return file_api_ld_types_ld_proto_rawDescGZIP(), []int{0}
}

func (x *ProviderID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Document is a JSON-LD context document with associated metadata.
type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	DocumentURL string `protobuf:"bytes,2,opt,name=documentURL,proto3" json:"documentURL,omitempty"`
	Content     []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"` // json.RawMessage Content of the context document.
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ld_types_ld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_api_ld_types_ld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_api_ld_types_ld_proto_rawDescGZIP(), []int{1}
}

func (x *Document) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Document) GetDocumentURL() string {
	if x != nil {
		return x.DocumentURL
	}
	return ""
}

func (x *Document) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

// AddContextsRequest is a request model for adding JSON-LD contexts.
type AddContextsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Documents []*Document `protobuf:"bytes,1,rep,name=documents,proto3" json:"documents,omitempty"`
}

func (x *AddContextsRequest) Reset() {
	*x = AddContextsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ld_types_ld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddContextsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddContextsRequest) ProtoMessage() {}

func (x *AddContextsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ld_types_ld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddContextsRequest.ProtoReflect.Descriptor instead.
func (*AddContextsRequest) Descriptor() ([]byte, []int) {
	return file_api_ld_types_ld_proto_rawDescGZIP(), []int{2}
}

func (x *AddContextsRequest) GetDocuments() []*Document {
	if x != nil {
		return x.Documents
	}
	return nil
}

// AddRemoteProviderRequest is a request model for adding a new remote context provider.
type AddRemoteProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *AddRemoteProviderRequest) Reset() {
	*x = AddRemoteProviderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ld_types_ld_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRemoteProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRemoteProviderRequest) ProtoMessage() {}

func (x *AddRemoteProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ld_types_ld_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRemoteProviderRequest.ProtoReflect.Descriptor instead.
func (*AddRemoteProviderRequest) Descriptor() ([]byte, []int) {
	return file_api_ld_types_ld_proto_rawDescGZIP(), []int{3}
}

func (x *AddRemoteProviderRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

// RemoteProviderRecord is a record in store with remote provider info.
type RemoteProviderRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *RemoteProviderRecord) Reset() {
	*x = RemoteProviderRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ld_types_ld_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteProviderRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteProviderRecord) ProtoMessage() {}

func (x *RemoteProviderRecord) ProtoReflect() protoreflect.Message {
	mi := &file_api_ld_types_ld_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteProviderRecord.ProtoReflect.Descriptor instead.
func (*RemoteProviderRecord) Descriptor() ([]byte, []int) {
	return file_api_ld_types_ld_proto_rawDescGZIP(), []int{4}
}

func (x *RemoteProviderRecord) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemoteProviderRecord) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

// GetAllRemoteProvidersResponse is a response model for listing all remote providers.
type GetAllRemoteProvidersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Providers []*RemoteProviderRecord `protobuf:"bytes,1,rep,name=providers,proto3" json:"providers,omitempty"`
}

func (x *GetAllRemoteProvidersResponse) Reset() {
	*x = GetAllRemoteProvidersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ld_types_ld_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRemoteProvidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRemoteProvidersResponse) ProtoMessage() {}

func (x *GetAllRemoteProvidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ld_types_ld_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRemoteProvidersResponse.ProtoReflect.Descriptor instead.
func (*GetAllRemoteProvidersResponse) Descriptor() ([]byte, []int) {
	return file_api_ld_types_ld_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllRemoteProvidersResponse) GetProviders() []*RemoteProviderRecord {
	if x != nil {
		return x.Providers
	}
	return nil
}

var File_api_ld_types_ld_proto protoreflect.FileDescriptor

var file_api_ld_types_ld_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6c,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x64, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x1c, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x58, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x52, 0x4c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x55, 0x52, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x4a, 0x0a,
	0x12, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x64, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x36, 0x0a, 0x18, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x22, 0x42, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x61, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6c, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x68, 0x61, 0x74, 0x74, 0x69, 0x2f, 0x47, 0x53,
	0x53, 0x49, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ld_types_ld_proto_rawDescOnce sync.Once
	file_api_ld_types_ld_proto_rawDescData = file_api_ld_types_ld_proto_rawDesc
)

func file_api_ld_types_ld_proto_rawDescGZIP() []byte {
	file_api_ld_types_ld_proto_rawDescOnce.Do(func() {
		file_api_ld_types_ld_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ld_types_ld_proto_rawDescData)
	})
	return file_api_ld_types_ld_proto_rawDescData
}

var file_api_ld_types_ld_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_ld_types_ld_proto_goTypes = []interface{}{
	(*ProviderID)(nil),                    // 0: api.ld.types.ProviderID
	(*Document)(nil),                      // 1: api.ld.types.Document
	(*AddContextsRequest)(nil),            // 2: api.ld.types.AddContextsRequest
	(*AddRemoteProviderRequest)(nil),      // 3: api.ld.types.AddRemoteProviderRequest
	(*RemoteProviderRecord)(nil),          // 4: api.ld.types.RemoteProviderRecord
	(*GetAllRemoteProvidersResponse)(nil), // 5: api.ld.types.GetAllRemoteProvidersResponse
}
var file_api_ld_types_ld_proto_depIdxs = []int32{
	1, // 0: api.ld.types.AddContextsRequest.documents:type_name -> api.ld.types.Document
	4, // 1: api.ld.types.GetAllRemoteProvidersResponse.providers:type_name -> api.ld.types.RemoteProviderRecord
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_ld_types_ld_proto_init() }
func file_api_ld_types_ld_proto_init() {
	if File_api_ld_types_ld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ld_types_ld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderID); i {
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
		file_api_ld_types_ld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
		file_api_ld_types_ld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddContextsRequest); i {
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
		file_api_ld_types_ld_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRemoteProviderRequest); i {
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
		file_api_ld_types_ld_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteProviderRecord); i {
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
		file_api_ld_types_ld_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRemoteProvidersResponse); i {
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
			RawDescriptor: file_api_ld_types_ld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_ld_types_ld_proto_goTypes,
		DependencyIndexes: file_api_ld_types_ld_proto_depIdxs,
		MessageInfos:      file_api_ld_types_ld_proto_msgTypes,
	}.Build()
	File_api_ld_types_ld_proto = out.File
	file_api_ld_types_ld_proto_rawDesc = nil
	file_api_ld_types_ld_proto_goTypes = nil
	file_api_ld_types_ld_proto_depIdxs = nil
}