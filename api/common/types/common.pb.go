// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/common/types/common.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// These are the encodings specified in the standard, not are all
// These are the encodings specified in the standard, not are all
// supported yet
type Encoding int32

const (
	Encoding_Identity          Encoding = 0
	Encoding_Base2             Encoding = 48     // '0'
	Encoding_Base8             Encoding = 55     // '7'
	Encoding_Base10            Encoding = 57     // '9'
	Encoding_Base16            Encoding = 102    //'f'
	Encoding_Base16Upper       Encoding = 70     // 'F'
	Encoding_Base32            Encoding = 98     // 'b'
	Encoding_Base32Upper       Encoding = 66     // 'B'
	Encoding_Base32pad         Encoding = 99     /// 'c'
	Encoding_Base32padUpper    Encoding = 67     // 'C'
	Encoding_Base32hex         Encoding = 118    // 'v'
	Encoding_Base32hexUpper    Encoding = 86     // 'V'
	Encoding_Base32hexPad      Encoding = 116    // 't'
	Encoding_Base32hexPadUpper Encoding = 84     // 'T'
	Encoding_Base36            Encoding = 107    // 'k'
	Encoding_Base36Upper       Encoding = 75     // 'K'
	Encoding_Base58BTC         Encoding = 122    // 'z'
	Encoding_Base58Flickr      Encoding = 90     // 'Z'
	Encoding_Base64            Encoding = 109    // 'm'
	Encoding_Base64url         Encoding = 117    // 'u'
	Encoding_Base64pad         Encoding = 77     // 'M'
	Encoding_Base64urlPad      Encoding = 85     // 'U'
	Encoding_Base256Emoji      Encoding = 128640 /// '🚀'
)

// Enum value maps for Encoding.
var (
	Encoding_name = map[int32]string{
		0:      "Identity",
		48:     "Base2",
		55:     "Base8",
		57:     "Base10",
		102:    "Base16",
		70:     "Base16Upper",
		98:     "Base32",
		66:     "Base32Upper",
		99:     "Base32pad",
		67:     "Base32padUpper",
		118:    "Base32hex",
		86:     "Base32hexUpper",
		116:    "Base32hexPad",
		84:     "Base32hexPadUpper",
		107:    "Base36",
		75:     "Base36Upper",
		122:    "Base58BTC",
		90:     "Base58Flickr",
		109:    "Base64",
		117:    "Base64url",
		77:     "Base64pad",
		85:     "Base64urlPad",
		128640: "Base256Emoji",
	}
	Encoding_value = map[string]int32{
		"Identity":          0,
		"Base2":             48,
		"Base8":             55,
		"Base10":            57,
		"Base16":            102,
		"Base16Upper":       70,
		"Base32":            98,
		"Base32Upper":       66,
		"Base32pad":         99,
		"Base32padUpper":    67,
		"Base32hex":         118,
		"Base32hexUpper":    86,
		"Base32hexPad":      116,
		"Base32hexPadUpper": 84,
		"Base36":            107,
		"Base36Upper":       75,
		"Base58BTC":         122,
		"Base58Flickr":      90,
		"Base64":            109,
		"Base64url":         117,
		"Base64pad":         77,
		"Base64urlPad":      85,
		"Base256Emoji":      128640,
	}
)

func (x Encoding) Enum() *Encoding {
	p := new(Encoding)
	*p = x
	return p
}

func (x Encoding) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Encoding) Descriptor() protoreflect.EnumDescriptor {
	return file_api_common_types_common_proto_enumTypes[0].Descriptor()
}

func (Encoding) Type() protoreflect.EnumType {
	return &file_api_common_types_common_proto_enumTypes[0]
}

func (x Encoding) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Encoding.Descriptor instead.
func (Encoding) EnumDescriptor() ([]byte, []int) {
	return file_api_common_types_common_proto_rawDescGZIP(), []int{0}
}

// TODO command types shouldn't be mentioned in common error package, [Issue#1182].
type Code int32

const (
	Code_DEFAULT Code = 0
	// Common error group for general command errors.
	Code_Common Code = 1000
	// DIDExchange error group for DID exchange command errors.
	Code_DIDExchange Code = 2000
	// Messaging error group for messaging service errors.
	Code_Messaging Code = 3000
	// VDR error group for VDR command errors.
	Code_VDR Code = 4000
	// ROUTE error group for Route command errors.
	Code_ROUTE Code = 5000
	// VC error group for Verifiable Credential command errors.
	Code_VC Code = 6000
	// KMS error group for key management service errors.
	Code_KMS Code = 7000
	// IssueCredential error group for issue credential command errors.
	Code_IssueCredential Code = 8000
	// PresentProof error group for present proof command errors.
	Code_PresentProof Code = 9000
	// Introduce error group for introduce command errors.
	Code_Introduce Code = 10000
	// Outofband error group for outofband command errors.
	Code_Outofband Code = 11000
	// OutofbandV2 error group for outofband command errors.
	Code_OutofbandV2 Code = 11100
	// VCWallet error group for verifiable Credential wallet command errors.
	Code_VCWallet Code = 12000
	// RFC0593 error group for RFC0593 command errors.
	Code_RFC0593 Code = 13000
	// LD error group for JSON-LD command errors.
	Code_LD Code = 14000
	// Connection error group for connection management errors.
	Code_Connection Code = 15000
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:     "DEFAULT",
		1000:  "Common",
		2000:  "DIDExchange",
		3000:  "Messaging",
		4000:  "VDR",
		5000:  "ROUTE",
		6000:  "VC",
		7000:  "KMS",
		8000:  "IssueCredential",
		9000:  "PresentProof",
		10000: "Introduce",
		11000: "Outofband",
		11100: "OutofbandV2",
		12000: "VCWallet",
		13000: "RFC0593",
		14000: "LD",
		15000: "Connection",
	}
	Code_value = map[string]int32{
		"DEFAULT":         0,
		"Common":          1000,
		"DIDExchange":     2000,
		"Messaging":       3000,
		"VDR":             4000,
		"ROUTE":           5000,
		"VC":              6000,
		"KMS":             7000,
		"IssueCredential": 8000,
		"PresentProof":    9000,
		"Introduce":       10000,
		"Outofband":       11000,
		"OutofbandV2":     11100,
		"VCWallet":        12000,
		"RFC0593":         13000,
		"LD":              14000,
		"Connection":      15000,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_api_common_types_common_proto_enumTypes[1].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_api_common_types_common_proto_enumTypes[1]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_api_common_types_common_proto_rawDescGZIP(), []int{1}
}

// Timing keeps expiration time.
type Timing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpiresTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=expires_time,json=expiresTime,proto3" json:"expires_time,omitempty"`
}

func (x *Timing) Reset() {
	*x = Timing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_types_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timing) ProtoMessage() {}

func (x *Timing) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_types_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timing.ProtoReflect.Descriptor instead.
func (*Timing) Descriptor() ([]byte, []int) {
	return file_api_common_types_common_proto_rawDescGZIP(), []int{0}
}

func (x *Timing) GetExpiresTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresTime
	}
	return nil
}

// EmptyResponse model
//
// swagger:response emptyResponse
type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// in: body
	Body *Empty `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_types_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_types_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_api_common_types_common_proto_rawDescGZIP(), []int{1}
}

func (x *EmptyResponse) GetBody() *Empty {
	if x != nil {
		return x.Body
	}
	return nil
}

// Empty model
//
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_types_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_types_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_common_types_common_proto_rawDescGZIP(), []int{2}
}

// swagger:response genericError
type DocGenericError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// in: body
	Body *GenericError `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *DocGenericError) Reset() {
	*x = DocGenericError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_types_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocGenericError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocGenericError) ProtoMessage() {}

func (x *DocGenericError) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_types_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocGenericError.ProtoReflect.Descriptor instead.
func (*DocGenericError) Descriptor() ([]byte, []int) {
	return file_api_common_types_common_proto_rawDescGZIP(), []int{3}
}

func (x *DocGenericError) GetBody() *GenericError {
	if x != nil {
		return x.Body
	}
	return nil
}

// GenericError is aries rest api error response
// swagger:model
type GenericError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code   `protobuf:"varint,1,opt,name=code,proto3,enum=api.common.types.Code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GenericError) Reset() {
	*x = GenericError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_types_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericError) ProtoMessage() {}

func (x *GenericError) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_types_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericError.ProtoReflect.Descriptor instead.
func (*GenericError) Descriptor() ([]byte, []int) {
	return file_api_common_types_common_proto_rawDescGZIP(), []int{4}
}

func (x *GenericError) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_DEFAULT
}

func (x *GenericError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_common_types_common_proto protoreflect.FileDescriptor

var file_api_common_types_common_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x47, 0x0a, 0x06, 0x54, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x0c,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x0d, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x45, 0x0a, 0x0f, 0x64, 0x6f, 0x63, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x54, 0x0a, 0x0c, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a,
	0xf1, 0x02, 0x0a, 0x08, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x0c, 0x0a, 0x08,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x61,
	0x73, 0x65, 0x32, 0x10, 0x30, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x61, 0x73, 0x65, 0x38, 0x10, 0x37,
	0x12, 0x0a, 0x0a, 0x06, 0x42, 0x61, 0x73, 0x65, 0x31, 0x30, 0x10, 0x39, 0x12, 0x0a, 0x0a, 0x06,
	0x42, 0x61, 0x73, 0x65, 0x31, 0x36, 0x10, 0x66, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65,
	0x31, 0x36, 0x55, 0x70, 0x70, 0x65, 0x72, 0x10, 0x46, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x61, 0x73,
	0x65, 0x33, 0x32, 0x10, 0x62, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x33, 0x32, 0x55,
	0x70, 0x70, 0x65, 0x72, 0x10, 0x42, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x65, 0x33, 0x32,
	0x70, 0x61, 0x64, 0x10, 0x63, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x73, 0x65, 0x33, 0x32, 0x70,
	0x61, 0x64, 0x55, 0x70, 0x70, 0x65, 0x72, 0x10, 0x43, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73,
	0x65, 0x33, 0x32, 0x68, 0x65, 0x78, 0x10, 0x76, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x73, 0x65,
	0x33, 0x32, 0x68, 0x65, 0x78, 0x55, 0x70, 0x70, 0x65, 0x72, 0x10, 0x56, 0x12, 0x10, 0x0a, 0x0c,
	0x42, 0x61, 0x73, 0x65, 0x33, 0x32, 0x68, 0x65, 0x78, 0x50, 0x61, 0x64, 0x10, 0x74, 0x12, 0x15,
	0x0a, 0x11, 0x42, 0x61, 0x73, 0x65, 0x33, 0x32, 0x68, 0x65, 0x78, 0x50, 0x61, 0x64, 0x55, 0x70,
	0x70, 0x65, 0x72, 0x10, 0x54, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x61, 0x73, 0x65, 0x33, 0x36, 0x10,
	0x6b, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x33, 0x36, 0x55, 0x70, 0x70, 0x65, 0x72,
	0x10, 0x4b, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x65, 0x35, 0x38, 0x42, 0x54, 0x43, 0x10,
	0x7a, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x35, 0x38, 0x46, 0x6c, 0x69, 0x63, 0x6b,
	0x72, 0x10, 0x5a, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x10, 0x6d, 0x12,
	0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x75, 0x72, 0x6c, 0x10, 0x75, 0x12, 0x0d,
	0x0a, 0x09, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x70, 0x61, 0x64, 0x10, 0x4d, 0x12, 0x10, 0x0a,
	0x0c, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x75, 0x72, 0x6c, 0x50, 0x61, 0x64, 0x10, 0x55, 0x12,
	0x12, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x32, 0x35, 0x36, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x10,
	0x80, 0xed, 0x07, 0x2a, 0xfd, 0x01, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x10, 0xe8, 0x07, 0x12, 0x10, 0x0a, 0x0b, 0x44, 0x49, 0x44, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x10, 0xd0, 0x0f, 0x12, 0x0e, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x10, 0xb8, 0x17, 0x12, 0x08, 0x0a, 0x03, 0x56, 0x44, 0x52, 0x10,
	0xa0, 0x1f, 0x12, 0x0a, 0x0a, 0x05, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x10, 0x88, 0x27, 0x12, 0x07,
	0x0a, 0x02, 0x56, 0x43, 0x10, 0xf0, 0x2e, 0x12, 0x08, 0x0a, 0x03, 0x4b, 0x4d, 0x53, 0x10, 0xd8,
	0x36, 0x12, 0x14, 0x0a, 0x0f, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x10, 0xc0, 0x3e, 0x12, 0x11, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x10, 0xa8, 0x46, 0x12, 0x0e, 0x0a, 0x09, 0x49, 0x6e,
	0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x10, 0x90, 0x4e, 0x12, 0x0e, 0x0a, 0x09, 0x4f, 0x75,
	0x74, 0x6f, 0x66, 0x62, 0x61, 0x6e, 0x64, 0x10, 0xf8, 0x55, 0x12, 0x10, 0x0a, 0x0b, 0x4f, 0x75,
	0x74, 0x6f, 0x66, 0x62, 0x61, 0x6e, 0x64, 0x56, 0x32, 0x10, 0xdc, 0x56, 0x12, 0x0d, 0x0a, 0x08,
	0x56, 0x43, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x10, 0xe0, 0x5d, 0x12, 0x0c, 0x0a, 0x07, 0x52,
	0x46, 0x43, 0x30, 0x35, 0x39, 0x33, 0x10, 0xc8, 0x65, 0x12, 0x07, 0x0a, 0x02, 0x4c, 0x44, 0x10,
	0xb0, 0x6d, 0x12, 0x0f, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x10, 0x98, 0x75, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x68, 0x61, 0x74, 0x74, 0x69, 0x2f, 0x47, 0x53, 0x53, 0x49, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_common_types_common_proto_rawDescOnce sync.Once
	file_api_common_types_common_proto_rawDescData = file_api_common_types_common_proto_rawDesc
)

func file_api_common_types_common_proto_rawDescGZIP() []byte {
	file_api_common_types_common_proto_rawDescOnce.Do(func() {
		file_api_common_types_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_common_types_common_proto_rawDescData)
	})
	return file_api_common_types_common_proto_rawDescData
}

var file_api_common_types_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_common_types_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_common_types_common_proto_goTypes = []interface{}{
	(Encoding)(0),                 // 0: api.common.types.Encoding
	(Code)(0),                     // 1: api.common.types.Code
	(*Timing)(nil),                // 2: api.common.types.Timing
	(*EmptyResponse)(nil),         // 3: api.common.types.EmptyResponse
	(*Empty)(nil),                 // 4: api.common.types.Empty
	(*DocGenericError)(nil),       // 5: api.common.types.docGenericError
	(*GenericError)(nil),          // 6: api.common.types.GenericError
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_api_common_types_common_proto_depIdxs = []int32{
	7, // 0: api.common.types.Timing.expires_time:type_name -> google.protobuf.Timestamp
	4, // 1: api.common.types.EmptyResponse.body:type_name -> api.common.types.Empty
	6, // 2: api.common.types.docGenericError.body:type_name -> api.common.types.GenericError
	1, // 3: api.common.types.GenericError.code:type_name -> api.common.types.Code
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_common_types_common_proto_init() }
func file_api_common_types_common_proto_init() {
	if File_api_common_types_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_common_types_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timing); i {
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
		file_api_common_types_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
		file_api_common_types_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_common_types_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocGenericError); i {
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
		file_api_common_types_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericError); i {
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
			RawDescriptor: file_api_common_types_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_common_types_common_proto_goTypes,
		DependencyIndexes: file_api_common_types_common_proto_depIdxs,
		EnumInfos:         file_api_common_types_common_proto_enumTypes,
		MessageInfos:      file_api_common_types_common_proto_msgTypes,
	}.Build()
	File_api_common_types_common_proto = out.File
	file_api_common_types_common_proto_rawDesc = nil
	file_api_common_types_common_proto_goTypes = nil
	file_api_common_types_common_proto_depIdxs = nil
}
