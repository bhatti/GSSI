// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/spi/types/store.proto

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

// SortOrder specifies the sort order of query results.
type SortOrder int32

const (
	SortOrder_ASCENDING  SortOrder = 0
	SortOrder_DESCENDING SortOrder = 1
)

// Enum value maps for SortOrder.
var (
	SortOrder_name = map[int32]string{
		0: "ASCENDING",
		1: "DESCENDING",
	}
	SortOrder_value = map[string]int32{
		"ASCENDING":  0,
		"DESCENDING": 1,
	}
)

func (x SortOrder) Enum() *SortOrder {
	p := new(SortOrder)
	*p = x
	return p
}

func (x SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_api_spi_types_store_proto_enumTypes[0].Descriptor()
}

func (SortOrder) Type() protoreflect.EnumType {
	return &file_api_spi_types_store_proto_enumTypes[0]
}

func (x SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortOrder.Descriptor instead.
func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_api_spi_types_store_proto_rawDescGZIP(), []int{0}
}

// OperationType
type OperationType int32

const (
	OperationType_CREATE OperationType = 0
	OperationType_UPDATE OperationType = 1
	OperationType_DELETE OperationType = 2
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "CREATE",
		1: "UPDATE",
		2: "DELETE",
	}
	OperationType_value = map[string]int32{
		"CREATE": 0,
		"UPDATE": 1,
		"DELETE": 2,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_spi_types_store_proto_enumTypes[1].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_api_spi_types_store_proto_enumTypes[1]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_api_spi_types_store_proto_rawDescGZIP(), []int{1}
}

// SortOptions sets the order that results from an Iterator will be returned in. Sorting is based on the tag values
// associated with the TagName chosen below. The TagName you use below can be the same as the one you're querying on,
// or it can be a different one. Depending on the storage implementation, you may need to ensure that the TagName set
// below is in the Store's StoreConfiguration before trying to use it for sorting (or it may be optional,
// but recommended). If tag value strings are decimal numbers, then the sorting must be based on their numerical value
// instead of the string representations of those numbers (i.e. numerical sorting, not lexicographic).
// TagName cannot be blank.
type SortOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order   SortOrder `protobuf:"varint,1,opt,name=order,proto3,enum=api.spi.types.SortOrder" json:"order,omitempty"`
	TagName string    `protobuf:"bytes,2,opt,name=tagName,proto3" json:"tagName,omitempty"`
}

func (x *SortOptions) Reset() {
	*x = SortOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_spi_types_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortOptions) ProtoMessage() {}

func (x *SortOptions) ProtoReflect() protoreflect.Message {
	mi := &file_api_spi_types_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortOptions.ProtoReflect.Descriptor instead.
func (*SortOptions) Descriptor() ([]byte, []int) {
	return file_api_spi_types_store_proto_rawDescGZIP(), []int{0}
}

func (x *SortOptions) GetOrder() SortOrder {
	if x != nil {
		return x.Order
	}
	return SortOrder_ASCENDING
}

func (x *SortOptions) GetTagName() string {
	if x != nil {
		return x.TagName
	}
	return ""
}

// QueryOptions represents various options for Query calls in a store.
type QueryOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PageSize sets the page size used by the Store.Query method.
	PageSize int32 `protobuf:"varint,1,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	// InitialPageNum sets the page for the iterator returned from Store.Query to start from.
	PageNum int32 `protobuf:"varint,2,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	// SortOptions defines the sort order.
	SortedOptions *SortOptions `protobuf:"bytes,3,opt,name=sortedOptions,proto3" json:"sortedOptions,omitempty"`
}

func (x *QueryOptions) Reset() {
	*x = QueryOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_spi_types_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOptions) ProtoMessage() {}

func (x *QueryOptions) ProtoReflect() protoreflect.Message {
	mi := &file_api_spi_types_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOptions.ProtoReflect.Descriptor instead.
func (*QueryOptions) Descriptor() ([]byte, []int) {
	return file_api_spi_types_store_proto_rawDescGZIP(), []int{1}
}

func (x *QueryOptions) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryOptions) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *QueryOptions) GetSortedOptions() *SortOptions {
	if x != nil {
		return x.SortedOptions
	}
	return nil
}

// Tag represents a Name + Value pair that can be associated with a key + value pair for querying later.
// swagger:model
type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name can be used to tag a given key + value pair as belonging to some sort of common
	// group. Example: Identifying a key+value pair as being a Verifiable Credential by storing it
	// with a tag Name called "VC". When used with the optional Value (see below), tag Name + Value can be used to
	// specify metadata for a given key + value pair. Example: Identifying a Verifiable Credential (stored as a
	// key+value pair) as belonging to a user account by using a tag Name called "UserAccount" and a tag Value called
	// "bob@example.com". Tag Names are intended to be static values that the store is configured with in order to build
	// indexes for queries (see TagNames in StoreConfiguration).
	// Tag Names cannot contain any ':' characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Value can optionally be used to indicate some metadata associated with a tag name for a given key + value pair.
	// See Name above for an example of how this can be used.
	// Tag Values are dynamic and are not specified in a StoreConfiguration.
	// Tag Values cannot contain any ':' characters.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_spi_types_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_api_spi_types_store_proto_msgTypes[2]
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
	return file_api_spi_types_store_proto_rawDescGZIP(), []int{2}
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tag) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Bool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *Bool) Reset() {
	*x = Bool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_spi_types_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bool) ProtoMessage() {}

func (x *Bool) ProtoReflect() protoreflect.Message {
	mi := &file_api_spi_types_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bool.ProtoReflect.Descriptor instead.
func (*Bool) Descriptor() ([]byte, []int) {
	return file_api_spi_types_store_proto_rawDescGZIP(), []int{3}
}

func (x *Bool) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

// StoreConfiguration represents the configuration of a store.
// Currently, it's only used for creating indexes in underlying storage databases.
type StoreConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []*Tag `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *StoreConfiguration) Reset() {
	*x = StoreConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_spi_types_store_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreConfiguration) ProtoMessage() {}

func (x *StoreConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_api_spi_types_store_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreConfiguration.ProtoReflect.Descriptor instead.
func (*StoreConfiguration) Descriptor() ([]byte, []int) {
	return file_api_spi_types_store_proto_rawDescGZIP(), []int{4}
}

func (x *StoreConfiguration) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

// Operation represents an operation to be performed in the Batch method.
type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key           string        `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Tags          []*Tag        `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	OperationType OperationType `protobuf:"varint,3,opt,name=operationType,proto3,enum=api.spi.types.OperationType" json:"operationType,omitempty"`
	Value         []byte        `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_spi_types_store_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_api_spi_types_store_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_api_spi_types_store_proto_rawDescGZIP(), []int{5}
}

func (x *Operation) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Operation) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Operation) GetOperationType() OperationType {
	if x != nil {
		return x.OperationType
	}
	return OperationType_CREATE
}

func (x *Operation) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// Record model containing name, ID and other fields of interest.
// swagger:model
type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Context   []string `protobuf:"bytes,3,rep,name=context,proto3" json:"context,omitempty"`
	Type      []string `protobuf:"bytes,4,rep,name=type,proto3" json:"type,omitempty"`
	SubjectId string   `protobuf:"bytes,5,opt,name=subjectId,proto3" json:"subjectId,omitempty"`
	// MyDID and TheirDID contains information about participants who were involved in the process
	// of issuing a credential or presentation.
	MyDid    string `protobuf:"bytes,6,opt,name=my_did,json=myDid,proto3" json:"my_did,omitempty"`
	TheirDid string `protobuf:"bytes,7,opt,name=their_did,json=theirDid,proto3" json:"their_did,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_spi_types_store_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_api_spi_types_store_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_api_spi_types_store_proto_rawDescGZIP(), []int{6}
}

func (x *Record) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Record) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Record) GetContext() []string {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *Record) GetType() []string {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Record) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *Record) GetMyDid() string {
	if x != nil {
		return x.MyDid
	}
	return ""
}

func (x *Record) GetTheirDid() string {
	if x != nil {
		return x.TheirDid
	}
	return ""
}

var File_api_spi_types_store_proto protoreflect.FileDescriptor

var file_api_spi_types_store_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x57, 0x0a, 0x0b, 0x53, 0x6f,
	0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x40, 0x0a, 0x0d, 0x73, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x73,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2f, 0x0a, 0x03,
	0x54, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x16, 0x0a,
	0x04, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x3c, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x42, 0x0a, 0x0d, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x6d, 0x79, 0x5f, 0x64, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6d, 0x79, 0x44, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x65, 0x69, 0x72,
	0x5f, 0x64, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x65, 0x69,
	0x72, 0x44, 0x69, 0x64, 0x2a, 0x2a, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x2a, 0x33, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c,
	0x45, 0x54, 0x45, 0x10, 0x02, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x68, 0x61, 0x74, 0x74, 0x69, 0x2f, 0x47, 0x53, 0x53, 0x49, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_spi_types_store_proto_rawDescOnce sync.Once
	file_api_spi_types_store_proto_rawDescData = file_api_spi_types_store_proto_rawDesc
)

func file_api_spi_types_store_proto_rawDescGZIP() []byte {
	file_api_spi_types_store_proto_rawDescOnce.Do(func() {
		file_api_spi_types_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_spi_types_store_proto_rawDescData)
	})
	return file_api_spi_types_store_proto_rawDescData
}

var file_api_spi_types_store_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_spi_types_store_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_spi_types_store_proto_goTypes = []interface{}{
	(SortOrder)(0),             // 0: api.spi.types.SortOrder
	(OperationType)(0),         // 1: api.spi.types.OperationType
	(*SortOptions)(nil),        // 2: api.spi.types.SortOptions
	(*QueryOptions)(nil),       // 3: api.spi.types.QueryOptions
	(*Tag)(nil),                // 4: api.spi.types.Tag
	(*Bool)(nil),               // 5: api.spi.types.Bool
	(*StoreConfiguration)(nil), // 6: api.spi.types.StoreConfiguration
	(*Operation)(nil),          // 7: api.spi.types.Operation
	(*Record)(nil),             // 8: api.spi.types.Record
}
var file_api_spi_types_store_proto_depIdxs = []int32{
	0, // 0: api.spi.types.SortOptions.order:type_name -> api.spi.types.SortOrder
	2, // 1: api.spi.types.QueryOptions.sortedOptions:type_name -> api.spi.types.SortOptions
	4, // 2: api.spi.types.StoreConfiguration.tags:type_name -> api.spi.types.Tag
	4, // 3: api.spi.types.Operation.tags:type_name -> api.spi.types.Tag
	1, // 4: api.spi.types.Operation.operationType:type_name -> api.spi.types.OperationType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_spi_types_store_proto_init() }
func file_api_spi_types_store_proto_init() {
	if File_api_spi_types_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_spi_types_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortOptions); i {
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
		file_api_spi_types_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOptions); i {
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
		file_api_spi_types_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_spi_types_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bool); i {
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
		file_api_spi_types_store_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreConfiguration); i {
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
		file_api_spi_types_store_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
		file_api_spi_types_store_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
			RawDescriptor: file_api_spi_types_store_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_spi_types_store_proto_goTypes,
		DependencyIndexes: file_api_spi_types_store_proto_depIdxs,
		EnumInfos:         file_api_spi_types_store_proto_enumTypes,
		MessageInfos:      file_api_spi_types_store_proto_msgTypes,
	}.Build()
	File_api_spi_types_store_proto = out.File
	file_api_spi_types_store_proto_rawDesc = nil
	file_api_spi_types_store_proto_goTypes = nil
	file_api_spi_types_store_proto_depIdxs = nil
}
