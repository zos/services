// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.7
// source: vts.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EndorsementType int32

const (
	EndorsementType_UNSET            EndorsementType = 0
	EndorsementType_REFERENCE_VALUE  EndorsementType = 1
	EndorsementType_VERIFICATION_KEY EndorsementType = 2
)

// Enum value maps for EndorsementType.
var (
	EndorsementType_name = map[int32]string{
		0: "UNSET",
		1: "REFERENCE_VALUE",
		2: "VERIFICATION_KEY",
	}
	EndorsementType_value = map[string]int32{
		"UNSET":            0,
		"REFERENCE_VALUE":  1,
		"VERIFICATION_KEY": 2,
	}
)

func (x EndorsementType) Enum() *EndorsementType {
	p := new(EndorsementType)
	*p = x
	return p
}

func (x EndorsementType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EndorsementType) Descriptor() protoreflect.EnumDescriptor {
	return file_vts_proto_enumTypes[0].Descriptor()
}

func (EndorsementType) Type() protoreflect.EnumType {
	return &file_vts_proto_enumTypes[0]
}

func (x EndorsementType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EndorsementType.Descriptor instead.
func (EndorsementType) EnumDescriptor() ([]byte, []int) {
	return file_vts_proto_rawDescGZIP(), []int{0}
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result      bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	ErrorDetail string `protobuf:"bytes,2,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_vts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_vts_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *Status) GetErrorDetail() string {
	if x != nil {
		return x.ErrorDetail
	}
	return ""
}

type Evidence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *structpb.Struct `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Evidence) Reset() {
	*x = Evidence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Evidence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Evidence) ProtoMessage() {}

func (x *Evidence) ProtoReflect() protoreflect.Message {
	mi := &file_vts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Evidence.ProtoReflect.Descriptor instead.
func (*Evidence) Descriptor() ([]byte, []int) {
	return file_vts_proto_rawDescGZIP(), []int{1}
}

func (x *Evidence) GetValue() *structpb.Struct {
	if x != nil {
		return x.Value
	}
	return nil
}

type AddSwComponentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SwComponents []*Endorsement `protobuf:"bytes,1,rep,name=sw_components,json=swComponents,proto3" json:"sw_components,omitempty"`
}

func (x *AddSwComponentsRequest) Reset() {
	*x = AddSwComponentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSwComponentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSwComponentsRequest) ProtoMessage() {}

func (x *AddSwComponentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSwComponentsRequest.ProtoReflect.Descriptor instead.
func (*AddSwComponentsRequest) Descriptor() ([]byte, []int) {
	return file_vts_proto_rawDescGZIP(), []int{2}
}

func (x *AddSwComponentsRequest) GetSwComponents() []*Endorsement {
	if x != nil {
		return x.SwComponents
	}
	return nil
}

type AddSwComponentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AddSwComponentsResponse) Reset() {
	*x = AddSwComponentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSwComponentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSwComponentsResponse) ProtoMessage() {}

func (x *AddSwComponentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSwComponentsResponse.ProtoReflect.Descriptor instead.
func (*AddSwComponentsResponse) Descriptor() ([]byte, []int) {
	return file_vts_proto_rawDescGZIP(), []int{3}
}

func (x *AddSwComponentsResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type Endorsement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scheme     AttestationFormat `protobuf:"varint,1,opt,name=scheme,proto3,enum=proto.AttestationFormat" json:"scheme,omitempty"`
	Type       EndorsementType   `protobuf:"varint,2,opt,name=type,proto3,enum=proto.EndorsementType" json:"type,omitempty"`
	Attributes *structpb.Struct  `protobuf:"bytes,3,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *Endorsement) Reset() {
	*x = Endorsement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endorsement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endorsement) ProtoMessage() {}

func (x *Endorsement) ProtoReflect() protoreflect.Message {
	mi := &file_vts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endorsement.ProtoReflect.Descriptor instead.
func (*Endorsement) Descriptor() ([]byte, []int) {
	return file_vts_proto_rawDescGZIP(), []int{4}
}

func (x *Endorsement) GetScheme() AttestationFormat {
	if x != nil {
		return x.Scheme
	}
	return AttestationFormat_UNKNOWN_FORMAT
}

func (x *Endorsement) GetType() EndorsementType {
	if x != nil {
		return x.Type
	}
	return EndorsementType_UNSET
}

func (x *Endorsement) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type AddTrustAnchorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrustAnchor *Endorsement `protobuf:"bytes,1,opt,name=trust_anchor,json=trustAnchor,proto3" json:"trust_anchor,omitempty"`
}

func (x *AddTrustAnchorRequest) Reset() {
	*x = AddTrustAnchorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTrustAnchorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTrustAnchorRequest) ProtoMessage() {}

func (x *AddTrustAnchorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTrustAnchorRequest.ProtoReflect.Descriptor instead.
func (*AddTrustAnchorRequest) Descriptor() ([]byte, []int) {
	return file_vts_proto_rawDescGZIP(), []int{5}
}

func (x *AddTrustAnchorRequest) GetTrustAnchor() *Endorsement {
	if x != nil {
		return x.TrustAnchor
	}
	return nil
}

type AddTrustAnchorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AddTrustAnchorResponse) Reset() {
	*x = AddTrustAnchorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTrustAnchorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTrustAnchorResponse) ProtoMessage() {}

func (x *AddTrustAnchorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTrustAnchorResponse.ProtoReflect.Descriptor instead.
func (*AddTrustAnchorResponse) Descriptor() ([]byte, []int) {
	return file_vts_proto_rawDescGZIP(), []int{6}
}

func (x *AddTrustAnchorResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type MediaTypeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaTypes []string `protobuf:"bytes,1,rep,name=media_types,json=mediaTypes,proto3" json:"media_types,omitempty"`
}

func (x *MediaTypeList) Reset() {
	*x = MediaTypeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaTypeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaTypeList) ProtoMessage() {}

func (x *MediaTypeList) ProtoReflect() protoreflect.Message {
	mi := &file_vts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaTypeList.ProtoReflect.Descriptor instead.
func (*MediaTypeList) Descriptor() ([]byte, []int) {
	return file_vts_proto_rawDescGZIP(), []int{7}
}

func (x *MediaTypeList) GetMediaTypes() []string {
	if x != nil {
		return x.MediaTypes
	}
	return nil
}

var File_vts_proto protoreflect.FileDescriptor

var file_vts_proto_rawDesc = []byte{
	0x0a, 0x09, 0x76, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x61, 0x70, 0x70, 0x72, 0x61, 0x69, 0x73,
	0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22,
	0x39, 0x0a, 0x08, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x51, 0x0a, 0x16, 0x41, 0x64,
	0x64, 0x53, 0x77, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0d, 0x73, 0x77, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0c, 0x73, 0x77, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x40, 0x0a,
	0x17, 0x41, 0x64, 0x64, 0x53, 0x77, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xa4, 0x01, 0x0a, 0x0b, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x30, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x4e, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x54, 0x72, 0x75,
	0x73, 0x74, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x35, 0x0a, 0x0c, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e,
	0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x74, 0x72, 0x75, 0x73, 0x74,
	0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x22, 0x3f, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x54, 0x72, 0x75,
	0x73, 0x74, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x30, 0x0a, 0x0d, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2a, 0x47, 0x0a, 0x0f, 0x45, 0x6e, 0x64,
	0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05,
	0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x46, 0x45, 0x52,
	0x45, 0x4e, 0x43, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10,
	0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4b, 0x45, 0x59,
	0x10, 0x02, 0x32, 0xfe, 0x02, 0x0a, 0x03, 0x56, 0x54, 0x53, 0x12, 0x3e, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x70,
	0x70, 0x72, 0x61, 0x69, 0x73, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x52,
	0x0a, 0x22, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x50, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x53, 0x77, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64,
	0x64, 0x53, 0x77, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64,
	0x53, 0x77, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x72, 0x75, 0x73, 0x74,
	0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x64, 0x64, 0x54, 0x72, 0x75, 0x73, 0x74, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64,
	0x54, 0x72, 0x75, 0x73, 0x74, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x76, 0x65, 0x72, 0x61, 0x69, 0x73, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_vts_proto_rawDescOnce sync.Once
	file_vts_proto_rawDescData = file_vts_proto_rawDesc
)

func file_vts_proto_rawDescGZIP() []byte {
	file_vts_proto_rawDescOnce.Do(func() {
		file_vts_proto_rawDescData = protoimpl.X.CompressGZIP(file_vts_proto_rawDescData)
	})
	return file_vts_proto_rawDescData
}

var file_vts_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vts_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_vts_proto_goTypes = []interface{}{
	(EndorsementType)(0),            // 0: proto.EndorsementType
	(*Status)(nil),                  // 1: proto.Status
	(*Evidence)(nil),                // 2: proto.Evidence
	(*AddSwComponentsRequest)(nil),  // 3: proto.AddSwComponentsRequest
	(*AddSwComponentsResponse)(nil), // 4: proto.AddSwComponentsResponse
	(*Endorsement)(nil),             // 5: proto.Endorsement
	(*AddTrustAnchorRequest)(nil),   // 6: proto.AddTrustAnchorRequest
	(*AddTrustAnchorResponse)(nil),  // 7: proto.AddTrustAnchorResponse
	(*MediaTypeList)(nil),           // 8: proto.MediaTypeList
	(*structpb.Struct)(nil),         // 9: google.protobuf.Struct
	(AttestationFormat)(0),          // 10: proto.AttestationFormat
	(*emptypb.Empty)(nil),           // 11: google.protobuf.Empty
	(*AttestationToken)(nil),        // 12: proto.AttestationToken
	(*ServiceState)(nil),            // 13: proto.ServiceState
	(*AppraisalContext)(nil),        // 14: proto.AppraisalContext
}
var file_vts_proto_depIdxs = []int32{
	9,  // 0: proto.Evidence.value:type_name -> google.protobuf.Struct
	5,  // 1: proto.AddSwComponentsRequest.sw_components:type_name -> proto.Endorsement
	1,  // 2: proto.AddSwComponentsResponse.status:type_name -> proto.Status
	10, // 3: proto.Endorsement.scheme:type_name -> proto.AttestationFormat
	0,  // 4: proto.Endorsement.type:type_name -> proto.EndorsementType
	9,  // 5: proto.Endorsement.attributes:type_name -> google.protobuf.Struct
	5,  // 6: proto.AddTrustAnchorRequest.trust_anchor:type_name -> proto.Endorsement
	1,  // 7: proto.AddTrustAnchorResponse.status:type_name -> proto.Status
	11, // 8: proto.VTS.GetServiceState:input_type -> google.protobuf.Empty
	12, // 9: proto.VTS.GetAttestation:input_type -> proto.AttestationToken
	11, // 10: proto.VTS.GetSupportedVerificationMediaTypes:input_type -> google.protobuf.Empty
	3,  // 11: proto.VTS.AddSwComponents:input_type -> proto.AddSwComponentsRequest
	6,  // 12: proto.VTS.AddTrustAnchor:input_type -> proto.AddTrustAnchorRequest
	13, // 13: proto.VTS.GetServiceState:output_type -> proto.ServiceState
	14, // 14: proto.VTS.GetAttestation:output_type -> proto.AppraisalContext
	8,  // 15: proto.VTS.GetSupportedVerificationMediaTypes:output_type -> proto.MediaTypeList
	4,  // 16: proto.VTS.AddSwComponents:output_type -> proto.AddSwComponentsResponse
	7,  // 17: proto.VTS.AddTrustAnchor:output_type -> proto.AddTrustAnchorResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_vts_proto_init() }
func file_vts_proto_init() {
	if File_vts_proto != nil {
		return
	}
	file_attestation_format_proto_init()
	file_appraisal_context_proto_init()
	file_token_proto_init()
	file_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_vts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Evidence); i {
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
		file_vts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSwComponentsRequest); i {
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
		file_vts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSwComponentsResponse); i {
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
		file_vts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endorsement); i {
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
		file_vts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTrustAnchorRequest); i {
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
		file_vts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTrustAnchorResponse); i {
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
		file_vts_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaTypeList); i {
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
			RawDescriptor: file_vts_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vts_proto_goTypes,
		DependencyIndexes: file_vts_proto_depIdxs,
		EnumInfos:         file_vts_proto_enumTypes,
		MessageInfos:      file_vts_proto_msgTypes,
	}.Build()
	File_vts_proto = out.File
	file_vts_proto_rawDesc = nil
	file_vts_proto_goTypes = nil
	file_vts_proto_depIdxs = nil
}
