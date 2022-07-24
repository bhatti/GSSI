// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/vc/services/issue.proto

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

// GetAllCredentialRequest for verifiable credentials
// Gets list of credentials or verifiable credentials
type GetAllCredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Query string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *GetAllCredentialRequest) Reset() {
	*x = GetAllCredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_vc_services_issue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCredentialRequest) ProtoMessage() {}

func (x *GetAllCredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_vc_services_issue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCredentialRequest.ProtoReflect.Descriptor instead.
func (*GetAllCredentialRequest) Descriptor() ([]byte, []int) {
	return file_api_vc_services_issue_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllCredentialRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllCredentialRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllCredentialRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

// GetCredentialRequest for verifiable credentials
// Gets a credential or verifiable credential by ID
type GetCredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCredentialRequest) Reset() {
	*x = GetCredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_vc_services_issue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCredentialRequest) ProtoMessage() {}

func (x *GetCredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_vc_services_issue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCredentialRequest.ProtoReflect.Descriptor instead.
func (*GetCredentialRequest) Descriptor() ([]byte, []int) {
	return file_api_vc_services_issue_proto_rawDescGZIP(), []int{1}
}

func (x *GetCredentialRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteCredentialRequest for verifiable credentials
// Deletes a credential or verifiable credential by ID
type DeleteCredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCredentialRequest) Reset() {
	*x = DeleteCredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_vc_services_issue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCredentialRequest) ProtoMessage() {}

func (x *DeleteCredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_vc_services_issue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCredentialRequest.ProtoReflect.Descriptor instead.
func (*DeleteCredentialRequest) Descriptor() ([]byte, []int) {
	return file_api_vc_services_issue_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteCredentialRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// IssueCredentialRequest for verifiable credentials
// POST /credentials/issue - Issues a credential and returns it in the response body.
type IssueCredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Verifies a verifiableCredential
	Credential *types.VerifiableCredential `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	// options
	Options *types.CredentialOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *IssueCredentialRequest) Reset() {
	*x = IssueCredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_vc_services_issue_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueCredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCredentialRequest) ProtoMessage() {}

func (x *IssueCredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_vc_services_issue_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCredentialRequest.ProtoReflect.Descriptor instead.
func (*IssueCredentialRequest) Descriptor() ([]byte, []int) {
	return file_api_vc_services_issue_proto_rawDescGZIP(), []int{3}
}

func (x *IssueCredentialRequest) GetCredential() *types.VerifiableCredential {
	if x != nil {
		return x.Credential
	}
	return nil
}

func (x *IssueCredentialRequest) GetOptions() *types.CredentialOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

// UpdateCredentialRequest for Updates the status of an issued credential
// POST /credentials/status
type UpdateCredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// credentialId ID
	CredentialId string `protobuf:"bytes,1,opt,name=credentialId,proto3" json:"credentialId,omitempty"`
	// credentialStatus
	CredentialStatus []*types.CredentialStatus `protobuf:"bytes,2,rep,name=credentialStatus,proto3" json:"credentialStatus,omitempty"`
}

func (x *UpdateCredentialRequest) Reset() {
	*x = UpdateCredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_vc_services_issue_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCredentialRequest) ProtoMessage() {}

func (x *UpdateCredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_vc_services_issue_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCredentialRequest.ProtoReflect.Descriptor instead.
func (*UpdateCredentialRequest) Descriptor() ([]byte, []int) {
	return file_api_vc_services_issue_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCredentialRequest) GetCredentialId() string {
	if x != nil {
		return x.CredentialId
	}
	return ""
}

func (x *UpdateCredentialRequest) GetCredentialStatus() []*types.CredentialStatus {
	if x != nil {
		return x.CredentialStatus
	}
	return nil
}

var File_api_vc_services_issue_proto protoreflect.FileDescriptor

var file_api_vc_services_issue_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x19,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x17, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x16, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x39, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x89, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x64,
	0x12, 0x4a, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x63, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x10, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xf1, 0x03, 0x0a,
	0x17, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x06, 0x67, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x5d, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x00, 0x12, 0x48,
	0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x05, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x63, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x06, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x63, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x00,
	0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x68, 0x61, 0x74, 0x74, 0x69, 0x2f, 0x47, 0x53, 0x53, 0x49, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_vc_services_issue_proto_rawDescOnce sync.Once
	file_api_vc_services_issue_proto_rawDescData = file_api_vc_services_issue_proto_rawDesc
)

func file_api_vc_services_issue_proto_rawDescGZIP() []byte {
	file_api_vc_services_issue_proto_rawDescOnce.Do(func() {
		file_api_vc_services_issue_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_vc_services_issue_proto_rawDescData)
	})
	return file_api_vc_services_issue_proto_rawDescData
}

var file_api_vc_services_issue_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_vc_services_issue_proto_goTypes = []interface{}{
	(*GetAllCredentialRequest)(nil),               // 0: api.vc.services.GetAllCredentialRequest
	(*GetCredentialRequest)(nil),                  // 1: api.vc.services.GetCredentialRequest
	(*DeleteCredentialRequest)(nil),               // 2: api.vc.services.DeleteCredentialRequest
	(*IssueCredentialRequest)(nil),                // 3: api.vc.services.IssueCredentialRequest
	(*UpdateCredentialRequest)(nil),               // 4: api.vc.services.UpdateCredentialRequest
	(*types.VerifiableCredential)(nil),            // 5: api.vc.types.VerifiableCredential
	(*types.CredentialOptions)(nil),               // 6: api.vc.types.CredentialOptions
	(*types.CredentialStatus)(nil),                // 7: api.vc.types.CredentialStatus
	(*types.RefreshableVerifiableCredential)(nil), // 8: api.vc.types.RefreshableVerifiableCredential
	(*types.Bool)(nil),                            // 9: api.vc.types.Bool
}
var file_api_vc_services_issue_proto_depIdxs = []int32{
	5, // 0: api.vc.services.IssueCredentialRequest.credential:type_name -> api.vc.types.VerifiableCredential
	6, // 1: api.vc.services.IssueCredentialRequest.options:type_name -> api.vc.types.CredentialOptions
	7, // 2: api.vc.services.UpdateCredentialRequest.credentialStatus:type_name -> api.vc.types.CredentialStatus
	0, // 3: api.vc.services.CredentialsIssueService.getAll:input_type -> api.vc.services.GetAllCredentialRequest
	1, // 4: api.vc.services.CredentialsIssueService.get:input_type -> api.vc.services.GetCredentialRequest
	2, // 5: api.vc.services.CredentialsIssueService.delete:input_type -> api.vc.services.DeleteCredentialRequest
	3, // 6: api.vc.services.CredentialsIssueService.issue:input_type -> api.vc.services.IssueCredentialRequest
	4, // 7: api.vc.services.CredentialsIssueService.update:input_type -> api.vc.services.UpdateCredentialRequest
	8, // 8: api.vc.services.CredentialsIssueService.getAll:output_type -> api.vc.types.RefreshableVerifiableCredential
	8, // 9: api.vc.services.CredentialsIssueService.get:output_type -> api.vc.types.RefreshableVerifiableCredential
	9, // 10: api.vc.services.CredentialsIssueService.delete:output_type -> api.vc.types.Bool
	8, // 11: api.vc.services.CredentialsIssueService.issue:output_type -> api.vc.types.RefreshableVerifiableCredential
	8, // 12: api.vc.services.CredentialsIssueService.update:output_type -> api.vc.types.RefreshableVerifiableCredential
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_vc_services_issue_proto_init() }
func file_api_vc_services_issue_proto_init() {
	if File_api_vc_services_issue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_vc_services_issue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCredentialRequest); i {
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
		file_api_vc_services_issue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCredentialRequest); i {
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
		file_api_vc_services_issue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCredentialRequest); i {
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
		file_api_vc_services_issue_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueCredentialRequest); i {
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
		file_api_vc_services_issue_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCredentialRequest); i {
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
			RawDescriptor: file_api_vc_services_issue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_vc_services_issue_proto_goTypes,
		DependencyIndexes: file_api_vc_services_issue_proto_depIdxs,
		MessageInfos:      file_api_vc_services_issue_proto_msgTypes,
	}.Build()
	File_api_vc_services_issue_proto = out.File
	file_api_vc_services_issue_proto_rawDesc = nil
	file_api_vc_services_issue_proto_goTypes = nil
	file_api_vc_services_issue_proto_depIdxs = nil
}
