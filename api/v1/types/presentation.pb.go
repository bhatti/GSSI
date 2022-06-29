// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/v1/types/presentation.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Presentation - Verifies a Presentation with or without proofs attached
type Presentation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the presentation.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The JSON-LD context of the credential. Each item in the @context array MUST be a string.
	Context []string `protobuf:"bytes,2,rep,name=context,proto3" json:"context,omitempty"`
	// The JSON-LD type of the credential. Each item in the type array MUST be a string.
	Type []string `protobuf:"bytes,3,rep,name=type,proto3" json:"type,omitempty"`
	// The holder - will be ignored if no proof is present since there is no proof of authority over the credentials.
	// The holder object MUST be The holder - will be ignored if no proof is present since there is no proof of authority
	// over the credentials (an object)
	Holder *anypb.Any `protobuf:"bytes,4,opt,name=holder,proto3" json:"holder,omitempty"`
	// The Verifiable Credentials Each item in the verifiableCredential array MUST be an object of the following form:
	VerifiableCredential []*VerifiableCredential `protobuf:"bytes,5,rep,name=verifiableCredential,proto3" json:"verifiableCredential,omitempty"`
}

func (x *Presentation) Reset() {
	*x = Presentation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_types_presentation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Presentation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Presentation) ProtoMessage() {}

func (x *Presentation) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_types_presentation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Presentation.ProtoReflect.Descriptor instead.
func (*Presentation) Descriptor() ([]byte, []int) {
	return file_api_v1_types_presentation_proto_rawDescGZIP(), []int{0}
}

func (x *Presentation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Presentation) GetContext() []string {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *Presentation) GetType() []string {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Presentation) GetHolder() *anypb.Any {
	if x != nil {
		return x.Holder
	}
	return nil
}

func (x *Presentation) GetVerifiableCredential() []*VerifiableCredential {
	if x != nil {
		return x.VerifiableCredential
	}
	return nil
}

// PresentationOptions for verifiable credentials
type PresentationOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the proof. Default is an appropriate proof type corresponding to the verification method.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The URI of the verificationMethod used for the proof. If omitted, a default verification method will be used.
	VerificationMethod string `protobuf:"bytes,2,opt,name=verificationMethod,proto3" json:"verificationMethod,omitempty"`
	// The purpose of the proof. Default 'assertionMethod'.
	ProofPurpose string `protobuf:"bytes,3,opt,name=proofPurpose,proto3" json:"proofPurpose,omitempty"`
	// A challenge provided by the requesting party of the proof. For example 6e62f66e-67de-11eb-b490-ef3eeefa55f2
	Challenge string `protobuf:"bytes,4,opt,name=challenge,proto3" json:"challenge,omitempty"`
	//     The intended domain of validity for the proof. For example website.example
	Domain string `protobuf:"bytes,5,opt,name=domain,proto3" json:"domain,omitempty"`
	// The created string
	Created *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *PresentationOptions) Reset() {
	*x = PresentationOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_types_presentation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PresentationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresentationOptions) ProtoMessage() {}

func (x *PresentationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_types_presentation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresentationOptions.ProtoReflect.Descriptor instead.
func (*PresentationOptions) Descriptor() ([]byte, []int) {
	return file_api_v1_types_presentation_proto_rawDescGZIP(), []int{1}
}

func (x *PresentationOptions) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PresentationOptions) GetVerificationMethod() string {
	if x != nil {
		return x.VerificationMethod
	}
	return ""
}

func (x *PresentationOptions) GetProofPurpose() string {
	if x != nil {
		return x.ProofPurpose
	}
	return ""
}

func (x *PresentationOptions) GetChallenge() string {
	if x != nil {
		return x.Challenge
	}
	return ""
}

func (x *PresentationOptions) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *PresentationOptions) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

// VerifiablePresentation - Verifies a Presentation with or without proofs attached
// A verifiable presentation expresses data from one or more verifiable credentials, and
// is packaged in such a way that the authorship of the data is verifiable. If verifiable
// credentials are presented directly, they become verifiable presentations. Data formats derived
// from verifiable credentials that are cryptographically verifiable, but do not of themselves
// contain verifiable credentials, might also be verifiable presentations.
// VerifiablePresentation = Presentation metadata + VerifiableCredentials + Proofs
type VerifiablePresentation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the presentation.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The JSON-LD context of the credential. Each item in the @context array MUST be a string.
	Context []string `protobuf:"bytes,2,rep,name=context,proto3" json:"context,omitempty"`
	// The JSON-LD type of the credential. Each item in the type array MUST be a string.
	Type []string `protobuf:"bytes,3,rep,name=type,proto3" json:"type,omitempty"`
	// The holder - will be ignored if no proof is present since there is no proof of authority over the credentials.
	// The holder object MUST be The holder - will be ignored if no proof is present since there is no proof of authority
	// over the credentials (an object)
	Holder *anypb.Any `protobuf:"bytes,4,opt,name=holder,proto3" json:"holder,omitempty"`
	// The Verifiable Credentials Each item in the verifiableCredential array MUST be an object of the following form:
	VerifiableCredential []*VerifiableCredential `protobuf:"bytes,5,rep,name=verifiableCredential,proto3" json:"verifiableCredential,omitempty"`
	// A JSON-LD Linked Data proof.
	Proof *CredentialProof `protobuf:"bytes,6,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *VerifiablePresentation) Reset() {
	*x = VerifiablePresentation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_types_presentation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifiablePresentation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifiablePresentation) ProtoMessage() {}

func (x *VerifiablePresentation) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_types_presentation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifiablePresentation.ProtoReflect.Descriptor instead.
func (*VerifiablePresentation) Descriptor() ([]byte, []int) {
	return file_api_v1_types_presentation_proto_rawDescGZIP(), []int{2}
}

func (x *VerifiablePresentation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VerifiablePresentation) GetContext() []string {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *VerifiablePresentation) GetType() []string {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *VerifiablePresentation) GetHolder() *anypb.Any {
	if x != nil {
		return x.Holder
	}
	return nil
}

func (x *VerifiablePresentation) GetVerifiableCredential() []*VerifiableCredential {
	if x != nil {
		return x.VerifiableCredential
	}
	return nil
}

func (x *VerifiablePresentation) GetProof() *CredentialProof {
	if x != nil {
		return x.Proof
	}
	return nil
}

// RefreshableVerifiablePresentation for verifiable presentation
type RefreshableVerifiablePresentation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the presentation.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The JSON-LD context of the credential. Each item in the @context array MUST be a string.
	Context []string `protobuf:"bytes,2,rep,name=context,proto3" json:"context,omitempty"`
	// The JSON-LD type of the credential. Each item in the type array MUST be a string.
	Type []string `protobuf:"bytes,3,rep,name=type,proto3" json:"type,omitempty"`
	// The holder - will be ignored if no proof is present since there is no proof of authority over the credentials.
	// The holder object MUST be The holder - will be ignored if no proof is present since there is no proof of authority
	// over the credentials (an object)
	Holder *anypb.Any `protobuf:"bytes,4,opt,name=holder,proto3" json:"holder,omitempty"`
	// The Verifiable Credentials Each item in the verifiableCredential array MUST be an object of the following form:
	VerifiableCredential []*VerifiableCredential `protobuf:"bytes,5,rep,name=verifiableCredential,proto3" json:"verifiableCredential,omitempty"`
	// RefreshService
	RefreshService *RefreshService `protobuf:"bytes,9,opt,name=refreshService,proto3" json:"refreshService,omitempty"`
}

func (x *RefreshableVerifiablePresentation) Reset() {
	*x = RefreshableVerifiablePresentation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_types_presentation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshableVerifiablePresentation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshableVerifiablePresentation) ProtoMessage() {}

func (x *RefreshableVerifiablePresentation) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_types_presentation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshableVerifiablePresentation.ProtoReflect.Descriptor instead.
func (*RefreshableVerifiablePresentation) Descriptor() ([]byte, []int) {
	return file_api_v1_types_presentation_proto_rawDescGZIP(), []int{3}
}

func (x *RefreshableVerifiablePresentation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RefreshableVerifiablePresentation) GetContext() []string {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *RefreshableVerifiablePresentation) GetType() []string {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *RefreshableVerifiablePresentation) GetHolder() *anypb.Any {
	if x != nil {
		return x.Holder
	}
	return nil
}

func (x *RefreshableVerifiablePresentation) GetVerifiableCredential() []*VerifiableCredential {
	if x != nil {
		return x.VerifiableCredential
	}
	return nil
}

func (x *RefreshableVerifiablePresentation) GetRefreshService() *RefreshService {
	if x != nil {
		return x.RefreshService
	}
	return nil
}

var File_api_v1_types_presentation_proto protoreflect.FileDescriptor

var file_api_v1_types_presentation_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x2c, 0x0a, 0x06, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12,
	0x53, 0x0a, 0x14, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x14,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x22, 0xe9, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x2e, 0x0a, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x50, 0x75, 0x72,
	0x70, 0x6f, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x22, 0x8b, 0x02, 0x0a, 0x16, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x50,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x06, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x14, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x14, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x30, 0x0a, 0x05,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0xa7,
	0x02, 0x0a, 0x21, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x12, 0x53, 0x0a, 0x14, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52,
	0x14, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x41, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x0e, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x68, 0x61, 0x74, 0x74, 0x69, 0x2f, 0x47, 0x53,
	0x53, 0x49, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_types_presentation_proto_rawDescOnce sync.Once
	file_api_v1_types_presentation_proto_rawDescData = file_api_v1_types_presentation_proto_rawDesc
)

func file_api_v1_types_presentation_proto_rawDescGZIP() []byte {
	file_api_v1_types_presentation_proto_rawDescOnce.Do(func() {
		file_api_v1_types_presentation_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_types_presentation_proto_rawDescData)
	})
	return file_api_v1_types_presentation_proto_rawDescData
}

var file_api_v1_types_presentation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_v1_types_presentation_proto_goTypes = []interface{}{
	(*Presentation)(nil),                      // 0: api.types.Presentation
	(*PresentationOptions)(nil),               // 1: api.types.PresentationOptions
	(*VerifiablePresentation)(nil),            // 2: api.types.VerifiablePresentation
	(*RefreshableVerifiablePresentation)(nil), // 3: api.types.RefreshableVerifiablePresentation
	(*anypb.Any)(nil),                         // 4: google.protobuf.Any
	(*VerifiableCredential)(nil),              // 5: api.types.VerifiableCredential
	(*timestamppb.Timestamp)(nil),             // 6: google.protobuf.Timestamp
	(*CredentialProof)(nil),                   // 7: api.types.CredentialProof
	(*RefreshService)(nil),                    // 8: api.types.RefreshService
}
var file_api_v1_types_presentation_proto_depIdxs = []int32{
	4, // 0: api.types.Presentation.holder:type_name -> google.protobuf.Any
	5, // 1: api.types.Presentation.verifiableCredential:type_name -> api.types.VerifiableCredential
	6, // 2: api.types.PresentationOptions.created:type_name -> google.protobuf.Timestamp
	4, // 3: api.types.VerifiablePresentation.holder:type_name -> google.protobuf.Any
	5, // 4: api.types.VerifiablePresentation.verifiableCredential:type_name -> api.types.VerifiableCredential
	7, // 5: api.types.VerifiablePresentation.proof:type_name -> api.types.CredentialProof
	4, // 6: api.types.RefreshableVerifiablePresentation.holder:type_name -> google.protobuf.Any
	5, // 7: api.types.RefreshableVerifiablePresentation.verifiableCredential:type_name -> api.types.VerifiableCredential
	8, // 8: api.types.RefreshableVerifiablePresentation.refreshService:type_name -> api.types.RefreshService
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_api_v1_types_presentation_proto_init() }
func file_api_v1_types_presentation_proto_init() {
	if File_api_v1_types_presentation_proto != nil {
		return
	}
	file_api_v1_types_common_proto_init()
	file_api_v1_types_credential_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_types_presentation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Presentation); i {
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
		file_api_v1_types_presentation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PresentationOptions); i {
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
		file_api_v1_types_presentation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifiablePresentation); i {
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
		file_api_v1_types_presentation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshableVerifiablePresentation); i {
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
			RawDescriptor: file_api_v1_types_presentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_types_presentation_proto_goTypes,
		DependencyIndexes: file_api_v1_types_presentation_proto_depIdxs,
		MessageInfos:      file_api_v1_types_presentation_proto_msgTypes,
	}.Build()
	File_api_v1_types_presentation_proto = out.File
	file_api_v1_types_presentation_proto_rawDesc = nil
	file_api_v1_types_presentation_proto_goTypes = nil
	file_api_v1_types_presentation_proto_depIdxs = nil
}