// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/did/jwk/types/jwk.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Format describes PresentationDefinition`s Format field.
type Format struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt   *JwtType `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	JwtVc *JwtType `protobuf:"bytes,2,opt,name=jwt_vc,json=jwtVc,proto3" json:"jwt_vc,omitempty"`
	JwtVp *JwtType `protobuf:"bytes,3,opt,name=jwt_vp,json=jwtVp,proto3" json:"jwt_vp,omitempty"`
	Ldp   *LdpType `protobuf:"bytes,4,opt,name=ldp,proto3" json:"ldp,omitempty"`
	LdpVc *LdpType `protobuf:"bytes,5,opt,name=ldp_vc,json=ldpVc,proto3" json:"ldp_vc,omitempty"`
	LdpVp *LdpType `protobuf:"bytes,6,opt,name=ldp_vp,json=ldpVp,proto3" json:"ldp_vp,omitempty"`
}

func (x *Format) Reset() {
	*x = Format{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_did_jwk_types_jwk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Format) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Format) ProtoMessage() {}

func (x *Format) ProtoReflect() protoreflect.Message {
	mi := &file_api_did_jwk_types_jwk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Format.ProtoReflect.Descriptor instead.
func (*Format) Descriptor() ([]byte, []int) {
	return file_api_did_jwk_types_jwk_proto_rawDescGZIP(), []int{0}
}

func (x *Format) GetJwt() *JwtType {
	if x != nil {
		return x.Jwt
	}
	return nil
}

func (x *Format) GetJwtVc() *JwtType {
	if x != nil {
		return x.JwtVc
	}
	return nil
}

func (x *Format) GetJwtVp() *JwtType {
	if x != nil {
		return x.JwtVp
	}
	return nil
}

func (x *Format) GetLdp() *LdpType {
	if x != nil {
		return x.Ldp
	}
	return nil
}

func (x *Format) GetLdpVc() *LdpType {
	if x != nil {
		return x.LdpVc
	}
	return nil
}

func (x *Format) GetLdpVp() *LdpType {
	if x != nil {
		return x.LdpVp
	}
	return nil
}

// CredentialApplication represents a credential_application object as defined in
// https://identity.foundation/credential-manifest/#credential-application.
// Note that the term "Credential Application" is overloaded in the spec - a "Credential Application" may be referring
// to one of two different, but related, concepts. A "Credential Application" can be the object defined below, which is
// intended to be embedded in an envelope like a Verifiable Presentation. Additionally, when that envelope contains
// the object defined below under a field named "credential_application", then that envelope itself can be called
// a "Credential Application". The larger "envelope version" of a Credential Application may also have a sibling
// presentation_submission object within the envelope, as demonstrated by the PresentCredentialApplication method.
// See https://github.com/decentralized-identity/credential-manifest/issues/73 for more information about this name
// overloading.
// swagger:model
type CredentialApplication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The value of this property MUST be the ID of a valid Credential Manifest.
	ManifestId string `protobuf:"bytes,2,opt,name=manifest_id,json=manifestId,proto3" json:"manifest_id,omitempty"`
	// Must be a subset of the format property of the CredentialManifest that this CredentialApplication is related to
	Format *Format `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"`
}

func (x *CredentialApplication) Reset() {
	*x = CredentialApplication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_did_jwk_types_jwk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CredentialApplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialApplication) ProtoMessage() {}

func (x *CredentialApplication) ProtoReflect() protoreflect.Message {
	mi := &file_api_did_jwk_types_jwk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialApplication.ProtoReflect.Descriptor instead.
func (*CredentialApplication) Descriptor() ([]byte, []int) {
	return file_api_did_jwk_types_jwk_proto_rawDescGZIP(), []int{1}
}

func (x *CredentialApplication) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CredentialApplication) GetManifestId() string {
	if x != nil {
		return x.ManifestId
	}
	return ""
}

func (x *CredentialApplication) GetFormat() *Format {
	if x != nil {
		return x.Format
	}
	return nil
}

// JwtType contains alg.
type JwtType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alg []string `protobuf:"bytes,1,rep,name=alg,proto3" json:"alg,omitempty"`
}

func (x *JwtType) Reset() {
	*x = JwtType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_did_jwk_types_jwk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JwtType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtType) ProtoMessage() {}

func (x *JwtType) ProtoReflect() protoreflect.Message {
	mi := &file_api_did_jwk_types_jwk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtType.ProtoReflect.Descriptor instead.
func (*JwtType) Descriptor() ([]byte, []int) {
	return file_api_did_jwk_types_jwk_proto_rawDescGZIP(), []int{2}
}

func (x *JwtType) GetAlg() []string {
	if x != nil {
		return x.Alg
	}
	return nil
}

// LdpType contains proof_type.
type LdpType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProofType []string `protobuf:"bytes,1,rep,name=proof_type,json=proofType,proto3" json:"proof_type,omitempty"`
}

func (x *LdpType) Reset() {
	*x = LdpType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_did_jwk_types_jwk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LdpType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LdpType) ProtoMessage() {}

func (x *LdpType) ProtoReflect() protoreflect.Message {
	mi := &file_api_did_jwk_types_jwk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LdpType.ProtoReflect.Descriptor instead.
func (*LdpType) Descriptor() ([]byte, []int) {
	return file_api_did_jwk_types_jwk_proto_rawDescGZIP(), []int{3}
}

func (x *LdpType) GetProofType() []string {
	if x != nil {
		return x.ProofType
	}
	return nil
}

// swagger:model
type JSONWebKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cryptographic key, can be a symmetric or asymmetric key.
	Key *anypb.Any `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Key identifier, parsed from `kid` header.
	Kid string `protobuf:"bytes,2,opt,name=kid,proto3" json:"kid,omitempty"` //keyID
	// Key algorithm, parsed from `alg` header.
	Alg string `protobuf:"bytes,3,opt,name=alg,proto3" json:"alg,omitempty"`
	// Key use, parsed from `use` header.
	Use string `protobuf:"bytes,4,opt,name=use,proto3" json:"use,omitempty"`
	Kty string `protobuf:"bytes,5,opt,name=kty,proto3" json:"kty,omitempty"`
	Crv string `protobuf:"bytes,6,opt,name=crv,proto3" json:"crv,omitempty"`
	X   string `protobuf:"bytes,7,opt,name=x,proto3" json:"x,omitempty"`
	Y   string `protobuf:"bytes,8,opt,name=y,proto3" json:"y,omitempty"`
	D   string `protobuf:"bytes,9,opt,name=d,proto3" json:"d,omitempty"`
	// X.509 certificate chain, parsed from `x5c` header.
	Certificate []*anypb.Any `protobuf:"bytes,10,rep,name=certificate,proto3" json:"certificate,omitempty"`
	// X.509 certificate URL, parsed from `x5u` header.
	CertificatesUrl string `protobuf:"bytes,11,opt,name=CertificatesUrl,proto3" json:"CertificatesUrl,omitempty"`
	// X.509 certificate thumbprint (SHA-1), parsed from `x5t` header.
	CertificateThumbprintSHA1 []byte `protobuf:"bytes,12,opt,name=certificateThumbprintSHA1,proto3" json:"certificateThumbprintSHA1,omitempty"`
	// X.509 certificate thumbprint (SHA-256), parsed from `x5t#S256` header.
	CertificateThumbprintSHA256 []byte `protobuf:"bytes,13,opt,name=certificateThumbprintSHA256,proto3" json:"certificateThumbprintSHA256,omitempty"`
}

func (x *JSONWebKey) Reset() {
	*x = JSONWebKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_did_jwk_types_jwk_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JSONWebKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JSONWebKey) ProtoMessage() {}

func (x *JSONWebKey) ProtoReflect() protoreflect.Message {
	mi := &file_api_did_jwk_types_jwk_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JSONWebKey.ProtoReflect.Descriptor instead.
func (*JSONWebKey) Descriptor() ([]byte, []int) {
	return file_api_did_jwk_types_jwk_proto_rawDescGZIP(), []int{4}
}

func (x *JSONWebKey) GetKey() *anypb.Any {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *JSONWebKey) GetKid() string {
	if x != nil {
		return x.Kid
	}
	return ""
}

func (x *JSONWebKey) GetAlg() string {
	if x != nil {
		return x.Alg
	}
	return ""
}

func (x *JSONWebKey) GetUse() string {
	if x != nil {
		return x.Use
	}
	return ""
}

func (x *JSONWebKey) GetKty() string {
	if x != nil {
		return x.Kty
	}
	return ""
}

func (x *JSONWebKey) GetCrv() string {
	if x != nil {
		return x.Crv
	}
	return ""
}

func (x *JSONWebKey) GetX() string {
	if x != nil {
		return x.X
	}
	return ""
}

func (x *JSONWebKey) GetY() string {
	if x != nil {
		return x.Y
	}
	return ""
}

func (x *JSONWebKey) GetD() string {
	if x != nil {
		return x.D
	}
	return ""
}

func (x *JSONWebKey) GetCertificate() []*anypb.Any {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *JSONWebKey) GetCertificatesUrl() string {
	if x != nil {
		return x.CertificatesUrl
	}
	return ""
}

func (x *JSONWebKey) GetCertificateThumbprintSHA1() []byte {
	if x != nil {
		return x.CertificateThumbprintSHA1
	}
	return nil
}

func (x *JSONWebKey) GetCertificateThumbprintSHA256() []byte {
	if x != nil {
		return x.CertificateThumbprintSHA256
	}
	return nil
}

// Recipient is a recipient of a JWE including the shared encryption key.
// swagger:model
type Recipient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *RecipientHeaders `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	EncryptedKey string            `protobuf:"bytes,2,opt,name=encrypted_key,json=encryptedKey,proto3" json:"encrypted_key,omitempty"`
}

func (x *Recipient) Reset() {
	*x = Recipient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_did_jwk_types_jwk_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recipient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipient) ProtoMessage() {}

func (x *Recipient) ProtoReflect() protoreflect.Message {
	mi := &file_api_did_jwk_types_jwk_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipient.ProtoReflect.Descriptor instead.
func (*Recipient) Descriptor() ([]byte, []int) {
	return file_api_did_jwk_types_jwk_proto_rawDescGZIP(), []int{5}
}

func (x *Recipient) GetHeader() *RecipientHeaders {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Recipient) GetEncryptedKey() string {
	if x != nil {
		return x.EncryptedKey
	}
	return ""
}

// RecipientHeaders are the recipient headers.
type RecipientHeaders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alg string `protobuf:"bytes,1,opt,name=alg,proto3" json:"alg,omitempty"`
	Apu string `protobuf:"bytes,2,opt,name=apu,proto3" json:"apu,omitempty"`
	Apv string `protobuf:"bytes,3,opt,name=apv,proto3" json:"apv,omitempty"`
	Iv  string `protobuf:"bytes,4,opt,name=iv,proto3" json:"iv,omitempty"`
	Tag string `protobuf:"bytes,5,opt,name=tag,proto3" json:"tag,omitempty"`
	Kid string `protobuf:"bytes,6,opt,name=kid,proto3" json:"kid,omitempty"`
	Epk []byte `protobuf:"bytes,7,opt,name=epk,proto3" json:"epk,omitempty"` // json.RawMessage
}

func (x *RecipientHeaders) Reset() {
	*x = RecipientHeaders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_did_jwk_types_jwk_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecipientHeaders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecipientHeaders) ProtoMessage() {}

func (x *RecipientHeaders) ProtoReflect() protoreflect.Message {
	mi := &file_api_did_jwk_types_jwk_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecipientHeaders.ProtoReflect.Descriptor instead.
func (*RecipientHeaders) Descriptor() ([]byte, []int) {
	return file_api_did_jwk_types_jwk_proto_rawDescGZIP(), []int{6}
}

func (x *RecipientHeaders) GetAlg() string {
	if x != nil {
		return x.Alg
	}
	return ""
}

func (x *RecipientHeaders) GetApu() string {
	if x != nil {
		return x.Apu
	}
	return ""
}

func (x *RecipientHeaders) GetApv() string {
	if x != nil {
		return x.Apv
	}
	return ""
}

func (x *RecipientHeaders) GetIv() string {
	if x != nil {
		return x.Iv
	}
	return ""
}

func (x *RecipientHeaders) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *RecipientHeaders) GetKid() string {
	if x != nil {
		return x.Kid
	}
	return ""
}

func (x *RecipientHeaders) GetEpk() []byte {
	if x != nil {
		return x.Epk
	}
	return nil
}

// JSONWebEncryption represents a RAW JWE that is used for serialization/deserialization.
// swagger:model
type JSONWebEncryption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protected    string `protobuf:"bytes,1,opt,name=protected,proto3" json:"protected,omitempty"`
	Unprotected  []byte `protobuf:"bytes,2,opt,name=unprotected,proto3" json:"unprotected,omitempty"`
	Recipients   []byte `protobuf:"bytes,3,opt,name=recipients,proto3" json:"recipients,omitempty"`
	EncryptedKey string `protobuf:"bytes,4,opt,name=encrypted_key,json=encryptedKey,proto3" json:"encrypted_key,omitempty"`
	Header       []byte `protobuf:"bytes,5,opt,name=header,proto3" json:"header,omitempty"`
	Aad          string `protobuf:"bytes,6,opt,name=aad,proto3" json:"aad,omitempty"`
	Iv           string `protobuf:"bytes,7,opt,name=iv,proto3" json:"iv,omitempty"`
	Ciphertext   string `protobuf:"bytes,8,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	Tag          string `protobuf:"bytes,9,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *JSONWebEncryption) Reset() {
	*x = JSONWebEncryption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_did_jwk_types_jwk_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JSONWebEncryption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JSONWebEncryption) ProtoMessage() {}

func (x *JSONWebEncryption) ProtoReflect() protoreflect.Message {
	mi := &file_api_did_jwk_types_jwk_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JSONWebEncryption.ProtoReflect.Descriptor instead.
func (*JSONWebEncryption) Descriptor() ([]byte, []int) {
	return file_api_did_jwk_types_jwk_proto_rawDescGZIP(), []int{7}
}

func (x *JSONWebEncryption) GetProtected() string {
	if x != nil {
		return x.Protected
	}
	return ""
}

func (x *JSONWebEncryption) GetUnprotected() []byte {
	if x != nil {
		return x.Unprotected
	}
	return nil
}

func (x *JSONWebEncryption) GetRecipients() []byte {
	if x != nil {
		return x.Recipients
	}
	return nil
}

func (x *JSONWebEncryption) GetEncryptedKey() string {
	if x != nil {
		return x.EncryptedKey
	}
	return ""
}

func (x *JSONWebEncryption) GetHeader() []byte {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *JSONWebEncryption) GetAad() string {
	if x != nil {
		return x.Aad
	}
	return ""
}

func (x *JSONWebEncryption) GetIv() string {
	if x != nil {
		return x.Iv
	}
	return ""
}

func (x *JSONWebEncryption) GetCiphertext() string {
	if x != nil {
		return x.Ciphertext
	}
	return ""
}

func (x *JSONWebEncryption) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

// Claims represents public claim values (as specified in RFC 7519).
// swagger:model
type Claims struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iss string   `protobuf:"bytes,1,opt,name=iss,proto3" json:"iss,omitempty"`  // issuer
	Sub string   `protobuf:"bytes,2,opt,name=sub,proto3" json:"sub,omitempty"`  // subject
	Aud []string `protobuf:"bytes,3,rep,name=aud,proto3" json:"aud,omitempty"`  // audience
	Exp int64    `protobuf:"varint,4,opt,name=exp,proto3" json:"exp,omitempty"` // expiry
	Nbf int64    `protobuf:"varint,5,opt,name=nbf,proto3" json:"nbf,omitempty"` // NotBefore
	Iat int64    `protobuf:"varint,6,opt,name=iat,proto3" json:"iat,omitempty"` // IssuedAt
	Jti string   `protobuf:"bytes,7,opt,name=jti,proto3" json:"jti,omitempty"`  // id
}

func (x *Claims) Reset() {
	*x = Claims{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_did_jwk_types_jwk_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Claims) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Claims) ProtoMessage() {}

func (x *Claims) ProtoReflect() protoreflect.Message {
	mi := &file_api_did_jwk_types_jwk_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Claims.ProtoReflect.Descriptor instead.
func (*Claims) Descriptor() ([]byte, []int) {
	return file_api_did_jwk_types_jwk_proto_rawDescGZIP(), []int{8}
}

func (x *Claims) GetIss() string {
	if x != nil {
		return x.Iss
	}
	return ""
}

func (x *Claims) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *Claims) GetAud() []string {
	if x != nil {
		return x.Aud
	}
	return nil
}

func (x *Claims) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *Claims) GetNbf() int64 {
	if x != nil {
		return x.Nbf
	}
	return 0
}

func (x *Claims) GetIat() int64 {
	if x != nil {
		return x.Iat
	}
	return 0
}

func (x *Claims) GetJti() string {
	if x != nil {
		return x.Jti
	}
	return ""
}

var File_api_did_jwk_types_jwk_proto protoreflect.FileDescriptor

var file_api_did_jwk_types_jwk_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x64, 0x2f, 0x6a, 0x77, 0x6b, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x6a, 0x77, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x69, 0x64, 0x2e, 0x6a, 0x77, 0x6b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x02, 0x0a, 0x06,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x69, 0x64, 0x2e, 0x6a, 0x77,
	0x6b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4a, 0x77, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x03, 0x6a, 0x77, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x6a, 0x77, 0x74, 0x5f, 0x76, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x69, 0x64, 0x2e, 0x6a,
	0x77, 0x6b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4a, 0x77, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x05, 0x6a, 0x77, 0x74, 0x56, 0x63, 0x12, 0x31, 0x0a, 0x06, 0x6a, 0x77, 0x74, 0x5f, 0x76,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x69,
	0x64, 0x2e, 0x6a, 0x77, 0x6b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4a, 0x77, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x05, 0x6a, 0x77, 0x74, 0x56, 0x70, 0x12, 0x2c, 0x0a, 0x03, 0x6c, 0x64,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x69,
	0x64, 0x2e, 0x6a, 0x77, 0x6b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x64, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x03, 0x6c, 0x64, 0x70, 0x12, 0x31, 0x0a, 0x06, 0x6c, 0x64, 0x70, 0x5f,
	0x76, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x69, 0x64, 0x2e, 0x6a, 0x77, 0x6b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x64, 0x70,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x6c, 0x64, 0x70, 0x56, 0x63, 0x12, 0x31, 0x0a, 0x06, 0x6c,
	0x64, 0x70, 0x5f, 0x76, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x64, 0x69, 0x64, 0x2e, 0x6a, 0x77, 0x6b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x4c, 0x64, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x6c, 0x64, 0x70, 0x56, 0x70, 0x22, 0x7b,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61,
	0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x69, 0x64, 0x2e, 0x6a, 0x77, 0x6b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x1b, 0x0a, 0x07, 0x4a,
	0x77, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x67, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x03, 0x61, 0x6c, 0x67, 0x22, 0x28, 0x0a, 0x07, 0x4c, 0x64, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x9a, 0x03, 0x0a, 0x0a, 0x4a, 0x53, 0x4f, 0x4e, 0x57, 0x65, 0x62, 0x4b, 0x65,
	0x79, 0x12, 0x26, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x6c, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x6c, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x74,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x72, 0x76, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x63, 0x72, 0x76, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x79, 0x12,
	0x0c, 0x0a, 0x01, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x64, 0x12, 0x36, 0x0a,
	0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x73, 0x55, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x55, 0x72, 0x6c, 0x12,
	0x3c, 0x0a, 0x19, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x68,
	0x75, 0x6d, 0x62, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x48, 0x41, 0x31, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x19, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54,
	0x68, 0x75, 0x6d, 0x62, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x48, 0x41, 0x31, 0x12, 0x40, 0x0a,
	0x1b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x68, 0x75, 0x6d,
	0x62, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x1b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54,
	0x68, 0x75, 0x6d, 0x62, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x22,
	0x6d, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x69, 0x64, 0x2e, 0x6a, 0x77, 0x6b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x22, 0x8e,
	0x01, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x61, 0x6c, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x75, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x75, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x76, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x76, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x76, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x70, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x65, 0x70, 0x6b, 0x22,
	0x84, 0x02, 0x0a, 0x11, 0x4a, 0x53, 0x4f, 0x4e, 0x57, 0x65, 0x62, 0x45, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x75, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x61, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x76, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x76, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x86, 0x01, 0x0a, 0x06, 0x43, 0x6c, 0x61, 0x69, 0x6d,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x69, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x75, 0x64, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x03, 0x61, 0x75, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x62, 0x66,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x62, 0x66, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x69, 0x61, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6a, 0x74, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x74, 0x69, 0x42,
	0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x68,
	0x61, 0x74, 0x74, 0x69, 0x2f, 0x47, 0x53, 0x53, 0x49, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69,
	0x64, 0x2f, 0x6a, 0x77, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_did_jwk_types_jwk_proto_rawDescOnce sync.Once
	file_api_did_jwk_types_jwk_proto_rawDescData = file_api_did_jwk_types_jwk_proto_rawDesc
)

func file_api_did_jwk_types_jwk_proto_rawDescGZIP() []byte {
	file_api_did_jwk_types_jwk_proto_rawDescOnce.Do(func() {
		file_api_did_jwk_types_jwk_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_did_jwk_types_jwk_proto_rawDescData)
	})
	return file_api_did_jwk_types_jwk_proto_rawDescData
}

var file_api_did_jwk_types_jwk_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_did_jwk_types_jwk_proto_goTypes = []interface{}{
	(*Format)(nil),                // 0: api.did.jwk.types.Format
	(*CredentialApplication)(nil), // 1: api.did.jwk.types.CredentialApplication
	(*JwtType)(nil),               // 2: api.did.jwk.types.JwtType
	(*LdpType)(nil),               // 3: api.did.jwk.types.LdpType
	(*JSONWebKey)(nil),            // 4: api.did.jwk.types.JSONWebKey
	(*Recipient)(nil),             // 5: api.did.jwk.types.Recipient
	(*RecipientHeaders)(nil),      // 6: api.did.jwk.types.RecipientHeaders
	(*JSONWebEncryption)(nil),     // 7: api.did.jwk.types.JSONWebEncryption
	(*Claims)(nil),                // 8: api.did.jwk.types.Claims
	(*anypb.Any)(nil),             // 9: google.protobuf.Any
}
var file_api_did_jwk_types_jwk_proto_depIdxs = []int32{
	2,  // 0: api.did.jwk.types.Format.jwt:type_name -> api.did.jwk.types.JwtType
	2,  // 1: api.did.jwk.types.Format.jwt_vc:type_name -> api.did.jwk.types.JwtType
	2,  // 2: api.did.jwk.types.Format.jwt_vp:type_name -> api.did.jwk.types.JwtType
	3,  // 3: api.did.jwk.types.Format.ldp:type_name -> api.did.jwk.types.LdpType
	3,  // 4: api.did.jwk.types.Format.ldp_vc:type_name -> api.did.jwk.types.LdpType
	3,  // 5: api.did.jwk.types.Format.ldp_vp:type_name -> api.did.jwk.types.LdpType
	0,  // 6: api.did.jwk.types.CredentialApplication.format:type_name -> api.did.jwk.types.Format
	9,  // 7: api.did.jwk.types.JSONWebKey.key:type_name -> google.protobuf.Any
	9,  // 8: api.did.jwk.types.JSONWebKey.certificate:type_name -> google.protobuf.Any
	6,  // 9: api.did.jwk.types.Recipient.header:type_name -> api.did.jwk.types.RecipientHeaders
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_api_did_jwk_types_jwk_proto_init() }
func file_api_did_jwk_types_jwk_proto_init() {
	if File_api_did_jwk_types_jwk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_did_jwk_types_jwk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Format); i {
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
		file_api_did_jwk_types_jwk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CredentialApplication); i {
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
		file_api_did_jwk_types_jwk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JwtType); i {
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
		file_api_did_jwk_types_jwk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LdpType); i {
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
		file_api_did_jwk_types_jwk_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JSONWebKey); i {
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
		file_api_did_jwk_types_jwk_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recipient); i {
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
		file_api_did_jwk_types_jwk_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecipientHeaders); i {
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
		file_api_did_jwk_types_jwk_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JSONWebEncryption); i {
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
		file_api_did_jwk_types_jwk_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Claims); i {
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
			RawDescriptor: file_api_did_jwk_types_jwk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_did_jwk_types_jwk_proto_goTypes,
		DependencyIndexes: file_api_did_jwk_types_jwk_proto_depIdxs,
		MessageInfos:      file_api_did_jwk_types_jwk_proto_msgTypes,
	}.Build()
	File_api_did_jwk_types_jwk_proto = out.File
	file_api_did_jwk_types_jwk_proto_rawDesc = nil
	file_api_did_jwk_types_jwk_proto_goTypes = nil
	file_api_did_jwk_types_jwk_proto_depIdxs = nil
}
