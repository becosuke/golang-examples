// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: kvstore.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreatePackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pack *Pack `protobuf:"bytes,1,opt,name=pack,proto3" json:"pack,omitempty"`
}

func (x *CreatePackRequest) Reset() {
	*x = CreatePackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePackRequest) ProtoMessage() {}

func (x *CreatePackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePackRequest.ProtoReflect.Descriptor instead.
func (*CreatePackRequest) Descriptor() ([]byte, []int) {
	return file_kvstore_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePackRequest) GetPack() *Pack {
	if x != nil {
		return x.Pack
	}
	return nil
}

type CreatePackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pack *Pack `protobuf:"bytes,1,opt,name=pack,proto3" json:"pack,omitempty"`
}

func (x *CreatePackResponse) Reset() {
	*x = CreatePackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvstore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePackResponse) ProtoMessage() {}

func (x *CreatePackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePackResponse.ProtoReflect.Descriptor instead.
func (*CreatePackResponse) Descriptor() ([]byte, []int) {
	return file_kvstore_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePackResponse) GetPack() *Pack {
	if x != nil {
		return x.Pack
	}
	return nil
}

type GetPackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seal *Seal `protobuf:"bytes,1,opt,name=seal,proto3" json:"seal,omitempty"`
}

func (x *GetPackRequest) Reset() {
	*x = GetPackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvstore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackRequest) ProtoMessage() {}

func (x *GetPackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackRequest.ProtoReflect.Descriptor instead.
func (*GetPackRequest) Descriptor() ([]byte, []int) {
	return file_kvstore_proto_rawDescGZIP(), []int{2}
}

func (x *GetPackRequest) GetSeal() *Seal {
	if x != nil {
		return x.Seal
	}
	return nil
}

type GetPackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pack *Pack `protobuf:"bytes,1,opt,name=pack,proto3" json:"pack,omitempty"`
}

func (x *GetPackResponse) Reset() {
	*x = GetPackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvstore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackResponse) ProtoMessage() {}

func (x *GetPackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackResponse.ProtoReflect.Descriptor instead.
func (*GetPackResponse) Descriptor() ([]byte, []int) {
	return file_kvstore_proto_rawDescGZIP(), []int{3}
}

func (x *GetPackResponse) GetPack() *Pack {
	if x != nil {
		return x.Pack
	}
	return nil
}

type UpdatePackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pack *Pack `protobuf:"bytes,1,opt,name=pack,proto3" json:"pack,omitempty"`
}

func (x *UpdatePackRequest) Reset() {
	*x = UpdatePackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvstore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePackRequest) ProtoMessage() {}

func (x *UpdatePackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePackRequest.ProtoReflect.Descriptor instead.
func (*UpdatePackRequest) Descriptor() ([]byte, []int) {
	return file_kvstore_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePackRequest) GetPack() *Pack {
	if x != nil {
		return x.Pack
	}
	return nil
}

type UpdatePackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pack *Pack `protobuf:"bytes,1,opt,name=pack,proto3" json:"pack,omitempty"`
}

func (x *UpdatePackResponse) Reset() {
	*x = UpdatePackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvstore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePackResponse) ProtoMessage() {}

func (x *UpdatePackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePackResponse.ProtoReflect.Descriptor instead.
func (*UpdatePackResponse) Descriptor() ([]byte, []int) {
	return file_kvstore_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePackResponse) GetPack() *Pack {
	if x != nil {
		return x.Pack
	}
	return nil
}

type DeletePackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seal *Seal `protobuf:"bytes,1,opt,name=seal,proto3" json:"seal,omitempty"`
}

func (x *DeletePackRequest) Reset() {
	*x = DeletePackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvstore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePackRequest) ProtoMessage() {}

func (x *DeletePackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePackRequest.ProtoReflect.Descriptor instead.
func (*DeletePackRequest) Descriptor() ([]byte, []int) {
	return file_kvstore_proto_rawDescGZIP(), []int{6}
}

func (x *DeletePackRequest) GetSeal() *Seal {
	if x != nil {
		return x.Seal
	}
	return nil
}

type Pack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Pack) Reset() {
	*x = Pack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvstore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pack) ProtoMessage() {}

func (x *Pack) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pack.ProtoReflect.Descriptor instead.
func (*Pack) Descriptor() ([]byte, []int) {
	return file_kvstore_proto_rawDescGZIP(), []int{7}
}

func (x *Pack) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Pack) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Seal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Seal) Reset() {
	*x = Seal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvstore_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Seal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Seal) ProtoMessage() {}

func (x *Seal) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Seal.ProtoReflect.Descriptor instead.
func (*Seal) Descriptor() ([]byte, []int) {
	return file_kvstore_proto_rawDescGZIP(), []int{8}
}

func (x *Seal) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_kvstore_proto protoreflect.FileDescriptor

var file_kvstore_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04,
	0x70, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x61, 0x63, 0x6b, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x70, 0x61, 0x63, 0x6b, 0x22, 0x32, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61,
	0x63, 0x6b, 0x52, 0x04, 0x70, 0x61, 0x63, 0x6b, 0x22, 0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x65,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x61, 0x6c, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x73, 0x65,
	0x61, 0x6c, 0x22, 0x2f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x04, 0x70,
	0x61, 0x63, 0x6b, 0x22, 0x3b, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x70, 0x61, 0x63, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x63, 0x6b,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x70, 0x61, 0x63, 0x6b,
	0x22, 0x32, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x04,
	0x70, 0x61, 0x63, 0x6b, 0x22, 0x3b, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x65, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61,
	0x6c, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x73, 0x65, 0x61,
	0x6c, 0x22, 0x46, 0x0a, 0x04, 0x50, 0x61, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18,
	0x80, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18,
	0x80, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x24, 0x0a, 0x04, 0x53, 0x65, 0x61,
	0x6c, 0x12, 0x1c, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0x80, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x32,
	0xea, 0x02, 0x0a, 0x0e, 0x4b, 0x56, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x50, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x63,
	0x6b, 0x3a, 0x01, 0x2a, 0x12, 0x4f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x12,
	0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x2f, 0x7b, 0x73, 0x65, 0x61, 0x6c,
	0x2e, 0x6b, 0x65, 0x79, 0x7d, 0x12, 0x5b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x63, 0x6b, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x1a, 0x13, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x63, 0x6b, 0x2f, 0x7b, 0x70, 0x61, 0x63, 0x6b, 0x2e, 0x6b, 0x65, 0x79, 0x7d, 0x3a,
	0x01, 0x2a, 0x12, 0x58, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x63,
	0x6b, 0x2f, 0x7b, 0x73, 0x65, 0x61, 0x6c, 0x2e, 0x6b, 0x65, 0x79, 0x7d, 0x42, 0x39, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x63, 0x6f, 0x73,
	0x75, 0x6b, 0x65, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x2f, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kvstore_proto_rawDescOnce sync.Once
	file_kvstore_proto_rawDescData = file_kvstore_proto_rawDesc
)

func file_kvstore_proto_rawDescGZIP() []byte {
	file_kvstore_proto_rawDescOnce.Do(func() {
		file_kvstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_kvstore_proto_rawDescData)
	})
	return file_kvstore_proto_rawDescData
}

var file_kvstore_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_kvstore_proto_goTypes = []interface{}{
	(*CreatePackRequest)(nil),  // 0: pb.CreatePackRequest
	(*CreatePackResponse)(nil), // 1: pb.CreatePackResponse
	(*GetPackRequest)(nil),     // 2: pb.GetPackRequest
	(*GetPackResponse)(nil),    // 3: pb.GetPackResponse
	(*UpdatePackRequest)(nil),  // 4: pb.UpdatePackRequest
	(*UpdatePackResponse)(nil), // 5: pb.UpdatePackResponse
	(*DeletePackRequest)(nil),  // 6: pb.DeletePackRequest
	(*Pack)(nil),               // 7: pb.Pack
	(*Seal)(nil),               // 8: pb.Seal
	(*emptypb.Empty)(nil),      // 9: google.protobuf.Empty
}
var file_kvstore_proto_depIdxs = []int32{
	7,  // 0: pb.CreatePackRequest.pack:type_name -> pb.Pack
	7,  // 1: pb.CreatePackResponse.pack:type_name -> pb.Pack
	8,  // 2: pb.GetPackRequest.seal:type_name -> pb.Seal
	7,  // 3: pb.GetPackResponse.pack:type_name -> pb.Pack
	7,  // 4: pb.UpdatePackRequest.pack:type_name -> pb.Pack
	7,  // 5: pb.UpdatePackResponse.pack:type_name -> pb.Pack
	8,  // 6: pb.DeletePackRequest.seal:type_name -> pb.Seal
	0,  // 7: pb.KVStoreService.CreatePack:input_type -> pb.CreatePackRequest
	2,  // 8: pb.KVStoreService.GetPack:input_type -> pb.GetPackRequest
	4,  // 9: pb.KVStoreService.UpdatePack:input_type -> pb.UpdatePackRequest
	6,  // 10: pb.KVStoreService.DeletePack:input_type -> pb.DeletePackRequest
	1,  // 11: pb.KVStoreService.CreatePack:output_type -> pb.CreatePackResponse
	3,  // 12: pb.KVStoreService.GetPack:output_type -> pb.GetPackResponse
	5,  // 13: pb.KVStoreService.UpdatePack:output_type -> pb.UpdatePackResponse
	9,  // 14: pb.KVStoreService.DeletePack:output_type -> google.protobuf.Empty
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_kvstore_proto_init() }
func file_kvstore_proto_init() {
	if File_kvstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kvstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePackRequest); i {
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
		file_kvstore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePackResponse); i {
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
		file_kvstore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackRequest); i {
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
		file_kvstore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackResponse); i {
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
		file_kvstore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePackRequest); i {
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
		file_kvstore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePackResponse); i {
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
		file_kvstore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePackRequest); i {
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
		file_kvstore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pack); i {
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
		file_kvstore_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Seal); i {
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
			RawDescriptor: file_kvstore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kvstore_proto_goTypes,
		DependencyIndexes: file_kvstore_proto_depIdxs,
		MessageInfos:      file_kvstore_proto_msgTypes,
	}.Build()
	File_kvstore_proto = out.File
	file_kvstore_proto_rawDesc = nil
	file_kvstore_proto_goTypes = nil
	file_kvstore_proto_depIdxs = nil
}
