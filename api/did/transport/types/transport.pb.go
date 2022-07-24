// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/did/transport/types/transport.proto

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

// Envelope holds message data and metadata for inbound and outbound messaging.
type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaTypeProfile string `protobuf:"bytes,1,opt,name=mediaTypeProfile,proto3" json:"mediaTypeProfile,omitempty"`
	Message          []byte `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FromKey          []byte `protobuf:"bytes,3,opt,name=fromKey,proto3" json:"fromKey,omitempty"`
	// ToKeys stores keys for an outbound message packing
	ToKeys []string `protobuf:"bytes,4,rep,name=toKeys,proto3" json:"toKeys,omitempty"`
	// ToKey holds the key that was used to decrypt an inbound message
	ToKey []byte `protobuf:"bytes,5,opt,name=toKey,proto3" json:"toKey,omitempty"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_did_transport_types_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_api_did_transport_types_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_api_did_transport_types_transport_proto_rawDescGZIP(), []int{0}
}

func (x *Envelope) GetMediaTypeProfile() string {
	if x != nil {
		return x.MediaTypeProfile
	}
	return ""
}

func (x *Envelope) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Envelope) GetFromKey() []byte {
	if x != nil {
		return x.FromKey
	}
	return nil
}

func (x *Envelope) GetToKeys() []string {
	if x != nil {
		return x.ToKeys
	}
	return nil
}

func (x *Envelope) GetToKey() []byte {
	if x != nil {
		return x.ToKey
	}
	return nil
}

var File_api_did_transport_types_transport_proto protoreflect.FileDescriptor

var file_api_did_transport_types_transport_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x69, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12,
	0x2a, 0x0a, 0x10, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x4b, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x4b, 0x65, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x4b, 0x65, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x68, 0x61, 0x74,
	0x74, 0x69, 0x2f, 0x47, 0x53, 0x53, 0x49, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x64, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_did_transport_types_transport_proto_rawDescOnce sync.Once
	file_api_did_transport_types_transport_proto_rawDescData = file_api_did_transport_types_transport_proto_rawDesc
)

func file_api_did_transport_types_transport_proto_rawDescGZIP() []byte {
	file_api_did_transport_types_transport_proto_rawDescOnce.Do(func() {
		file_api_did_transport_types_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_did_transport_types_transport_proto_rawDescData)
	})
	return file_api_did_transport_types_transport_proto_rawDescData
}

var file_api_did_transport_types_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_did_transport_types_transport_proto_goTypes = []interface{}{
	(*Envelope)(nil), // 0: api.did.transport.types.Envelope
}
var file_api_did_transport_types_transport_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_did_transport_types_transport_proto_init() }
func file_api_did_transport_types_transport_proto_init() {
	if File_api_did_transport_types_transport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_did_transport_types_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
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
			RawDescriptor: file_api_did_transport_types_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_did_transport_types_transport_proto_goTypes,
		DependencyIndexes: file_api_did_transport_types_transport_proto_depIdxs,
		MessageInfos:      file_api_did_transport_types_transport_proto_msgTypes,
	}.Build()
	File_api_did_transport_types_transport_proto = out.File
	file_api_did_transport_types_transport_proto_rawDesc = nil
	file_api_did_transport_types_transport_proto_goTypes = nil
	file_api_did_transport_types_transport_proto_depIdxs = nil
}