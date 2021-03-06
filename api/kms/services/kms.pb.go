// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/kms/services/kms.proto

package types

import (
	types1 "github.com/bhatti/GSSI/api/common/types"
	types "github.com/bhatti/GSSI/api/did/jwk/types"
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

// CreateKeySetRequest is model for createKeySey request.
// swagger:parameters createKeySetReq
type CreateKeySetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Params for createKeySet
	//
	// in: body
	KeyType string `protobuf:"bytes,1,opt,name=keyType,proto3" json:"keyType,omitempty"`
}

func (x *CreateKeySetRequest) Reset() {
	*x = CreateKeySetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kms_services_kms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKeySetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKeySetRequest) ProtoMessage() {}

func (x *CreateKeySetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_kms_services_kms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKeySetRequest.ProtoReflect.Descriptor instead.
func (*CreateKeySetRequest) Descriptor() ([]byte, []int) {
	return file_api_kms_services_kms_proto_rawDescGZIP(), []int{0}
}

func (x *CreateKeySetRequest) GetKeyType() string {
	if x != nil {
		return x.KeyType
	}
	return ""
}

// CreateKeySetResponse for returning key pair.
// This is used for returning the create set response
//
// swagger:response createKeySetRes
type CreateKeySetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// in: body
	//  key id base64 encoded
	KeyID string `protobuf:"bytes,1,opt,name=keyID,proto3" json:"keyID,omitempty"`
	//  public key base64 encoded
	PublicKey string `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
}

func (x *CreateKeySetResponse) Reset() {
	*x = CreateKeySetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kms_services_kms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKeySetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKeySetResponse) ProtoMessage() {}

func (x *CreateKeySetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_kms_services_kms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKeySetResponse.ProtoReflect.Descriptor instead.
func (*CreateKeySetResponse) Descriptor() ([]byte, []int) {
	return file_api_kms_services_kms_proto_rawDescGZIP(), []int{1}
}

func (x *CreateKeySetResponse) GetKeyID() string {
	if x != nil {
		return x.KeyID
	}
	return ""
}

func (x *CreateKeySetResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

// importKeyReq model
//
// This is used for import key req
//
// swagger:parameters importKeyReq
type ImportKeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// in: body
	Params *types.JSONWebKey `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"` // kms.JSONWebKey
}

func (x *ImportKeyReq) Reset() {
	*x = ImportKeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kms_services_kms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportKeyReq) ProtoMessage() {}

func (x *ImportKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_kms_services_kms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportKeyReq.ProtoReflect.Descriptor instead.
func (*ImportKeyReq) Descriptor() ([]byte, []int) {
	return file_api_kms_services_kms_proto_rawDescGZIP(), []int{2}
}

func (x *ImportKeyReq) GetParams() *types.JSONWebKey {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_api_kms_services_kms_proto protoreflect.FileDescriptor

var file_api_kms_services_kms_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x6d, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x6b, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x70,
	0x69, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1d,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61,
	0x70, 0x69, 0x2f, 0x64, 0x69, 0x64, 0x2f, 0x6a, 0x77, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x6a, 0x77, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4a, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x45, 0x0a, 0x0c, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x12, 0x35, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x69,
	0x64, 0x2e, 0x6a, 0x77, 0x6b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4a, 0x53, 0x4f, 0x4e,
	0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x32, 0xc0,
	0x01, 0x0a, 0x0d, 0x4b, 0x4d, 0x53, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x12, 0x5f, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74,
	0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x6d,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4e, 0x0a, 0x09, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x68, 0x61, 0x74, 0x74, 0x69, 0x2f, 0x47, 0x53, 0x53, 0x49, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6b, 0x6d, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_kms_services_kms_proto_rawDescOnce sync.Once
	file_api_kms_services_kms_proto_rawDescData = file_api_kms_services_kms_proto_rawDesc
)

func file_api_kms_services_kms_proto_rawDescGZIP() []byte {
	file_api_kms_services_kms_proto_rawDescOnce.Do(func() {
		file_api_kms_services_kms_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_kms_services_kms_proto_rawDescData)
	})
	return file_api_kms_services_kms_proto_rawDescData
}

var file_api_kms_services_kms_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_kms_services_kms_proto_goTypes = []interface{}{
	(*CreateKeySetRequest)(nil),  // 0: api.kms.services.CreateKeySetRequest
	(*CreateKeySetResponse)(nil), // 1: api.kms.services.CreateKeySetResponse
	(*ImportKeyReq)(nil),         // 2: api.kms.services.ImportKeyReq
	(*types.JSONWebKey)(nil),     // 3: api.did.jwk.types.JSONWebKey
	(*types1.EmptyResponse)(nil), // 4: api.common.types.EmptyResponse
}
var file_api_kms_services_kms_proto_depIdxs = []int32{
	3, // 0: api.kms.services.ImportKeyReq.params:type_name -> api.did.jwk.types.JSONWebKey
	0, // 1: api.kms.services.KMSController.createKeySet:input_type -> api.kms.services.CreateKeySetRequest
	2, // 2: api.kms.services.KMSController.importKey:input_type -> api.kms.services.ImportKeyReq
	1, // 3: api.kms.services.KMSController.createKeySet:output_type -> api.kms.services.CreateKeySetResponse
	4, // 4: api.kms.services.KMSController.importKey:output_type -> api.common.types.EmptyResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_kms_services_kms_proto_init() }
func file_api_kms_services_kms_proto_init() {
	if File_api_kms_services_kms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_kms_services_kms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKeySetRequest); i {
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
		file_api_kms_services_kms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKeySetResponse); i {
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
		file_api_kms_services_kms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportKeyReq); i {
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
			RawDescriptor: file_api_kms_services_kms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_kms_services_kms_proto_goTypes,
		DependencyIndexes: file_api_kms_services_kms_proto_depIdxs,
		MessageInfos:      file_api_kms_services_kms_proto_msgTypes,
	}.Build()
	File_api_kms_services_kms_proto = out.File
	file_api_kms_services_kms_proto_rawDesc = nil
	file_api_kms_services_kms_proto_goTypes = nil
	file_api_kms_services_kms_proto_depIdxs = nil
}
