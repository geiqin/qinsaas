// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: marketChannelService.proto

package services

import (
	common "github.com/geiqin/microkit/protobuf/common"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MarketChannelWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id       int32 `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	// @inject_tag: gorm:"-"
	Ids      []int32 `protobuf:"varint,4,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
	Status   int32   `protobuf:"varint,5,opt,name=status,proto3" json:"status"`
	Keywords string  `protobuf:"bytes,6,opt,name=keywords,proto3" json:"keywords"`
	Default  bool    `protobuf:"varint,7,opt,name=default,proto3" json:"default"`
}

func (x *MarketChannelWhere) Reset() {
	*x = MarketChannelWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketChannelService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketChannelWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketChannelWhere) ProtoMessage() {}

func (x *MarketChannelWhere) ProtoReflect() protoreflect.Message {
	mi := &file_marketChannelService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketChannelWhere.ProtoReflect.Descriptor instead.
func (*MarketChannelWhere) Descriptor() ([]byte, []int) {
	return file_marketChannelService_proto_rawDescGZIP(), []int{0}
}

func (x *MarketChannelWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *MarketChannelWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *MarketChannelWhere) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MarketChannelWhere) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *MarketChannelWhere) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MarketChannelWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *MarketChannelWhere) GetDefault() bool {
	if x != nil {
		return x.Default
	}
	return false
}

type MarketChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// @inject_tag: validate:"required" label:"名称"
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" validate:"required" label:"名称"`
	Sorting   int32  `protobuf:"varint,3,opt,name=sorting,proto3" json:"sorting"`
	Memo      string `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo"`
	Default   bool   `protobuf:"varint,5,opt,name=default,proto3" json:"default"`
	Status    int32  `protobuf:"varint,6,opt,name=status,proto3" json:"status"`
	CreatedAt string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *MarketChannel) Reset() {
	*x = MarketChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketChannelService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketChannel) ProtoMessage() {}

func (x *MarketChannel) ProtoReflect() protoreflect.Message {
	mi := &file_marketChannelService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketChannel.ProtoReflect.Descriptor instead.
func (*MarketChannel) Descriptor() ([]byte, []int) {
	return file_marketChannelService_proto_rawDescGZIP(), []int{1}
}

func (x *MarketChannel) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MarketChannel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MarketChannel) GetSorting() int32 {
	if x != nil {
		return x.Sorting
	}
	return 0
}

func (x *MarketChannel) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *MarketChannel) GetDefault() bool {
	if x != nil {
		return x.Default
	}
	return false
}

func (x *MarketChannel) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MarketChannel) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MarketChannel) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type MarketChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *common.Error    `protobuf:"bytes,1,opt,name=error,proto3" json:"error"`
	Info   *common.Info     `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	Pager  *common.Pager    `protobuf:"bytes,3,opt,name=pager,proto3" json:"pager"`
	Entity *MarketChannel   `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity"`
	Items  []*MarketChannel `protobuf:"bytes,5,rep,name=items,proto3" json:"items"`
}

func (x *MarketChannelResponse) Reset() {
	*x = MarketChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketChannelService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketChannelResponse) ProtoMessage() {}

func (x *MarketChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_marketChannelService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketChannelResponse.ProtoReflect.Descriptor instead.
func (*MarketChannelResponse) Descriptor() ([]byte, []int) {
	return file_marketChannelService_proto_rawDescGZIP(), []int{2}
}

func (x *MarketChannelResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *MarketChannelResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *MarketChannelResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *MarketChannelResponse) GetEntity() *MarketChannel {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *MarketChannelResponse) GetItems() []*MarketChannel {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_marketChannelService_proto protoreflect.FileDescriptor

var file_marketChannelService_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x12, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x57, 0x68, 0x65, 0x72, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x22, 0xd1, 0x01, 0x0a, 0x0d, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xe3, 0x01, 0x0a, 0x15, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2f, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2d,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x86, 0x04,
	0x0a, 0x14, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x1f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46,
	0x0a, 0x0a, 0x53, 0x65, 0x74, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x1f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_marketChannelService_proto_rawDescOnce sync.Once
	file_marketChannelService_proto_rawDescData = file_marketChannelService_proto_rawDesc
)

func file_marketChannelService_proto_rawDescGZIP() []byte {
	file_marketChannelService_proto_rawDescOnce.Do(func() {
		file_marketChannelService_proto_rawDescData = protoimpl.X.CompressGZIP(file_marketChannelService_proto_rawDescData)
	})
	return file_marketChannelService_proto_rawDescData
}

var file_marketChannelService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_marketChannelService_proto_goTypes = []interface{}{
	(*MarketChannelWhere)(nil),    // 0: services.MarketChannelWhere
	(*MarketChannel)(nil),         // 1: services.MarketChannel
	(*MarketChannelResponse)(nil), // 2: services.MarketChannelResponse
	(*common.Error)(nil),          // 3: common.Error
	(*common.Info)(nil),           // 4: common.Info
	(*common.Pager)(nil),          // 5: common.Pager
}
var file_marketChannelService_proto_depIdxs = []int32{
	3,  // 0: services.MarketChannelResponse.error:type_name -> common.Error
	4,  // 1: services.MarketChannelResponse.info:type_name -> common.Info
	5,  // 2: services.MarketChannelResponse.pager:type_name -> common.Pager
	1,  // 3: services.MarketChannelResponse.entity:type_name -> services.MarketChannel
	1,  // 4: services.MarketChannelResponse.items:type_name -> services.MarketChannel
	1,  // 5: services.MarketChannelService.Create:input_type -> services.MarketChannel
	1,  // 6: services.MarketChannelService.Update:input_type -> services.MarketChannel
	0,  // 7: services.MarketChannelService.Search:input_type -> services.MarketChannelWhere
	0,  // 8: services.MarketChannelService.List:input_type -> services.MarketChannelWhere
	1,  // 9: services.MarketChannelService.SetSorting:input_type -> services.MarketChannel
	1,  // 10: services.MarketChannelService.SetStatus:input_type -> services.MarketChannel
	0,  // 11: services.MarketChannelService.Delete:input_type -> services.MarketChannelWhere
	2,  // 12: services.MarketChannelService.Create:output_type -> services.MarketChannelResponse
	2,  // 13: services.MarketChannelService.Update:output_type -> services.MarketChannelResponse
	2,  // 14: services.MarketChannelService.Search:output_type -> services.MarketChannelResponse
	2,  // 15: services.MarketChannelService.List:output_type -> services.MarketChannelResponse
	2,  // 16: services.MarketChannelService.SetSorting:output_type -> services.MarketChannelResponse
	2,  // 17: services.MarketChannelService.SetStatus:output_type -> services.MarketChannelResponse
	2,  // 18: services.MarketChannelService.Delete:output_type -> services.MarketChannelResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_marketChannelService_proto_init() }
func file_marketChannelService_proto_init() {
	if File_marketChannelService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_marketChannelService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketChannelWhere); i {
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
		file_marketChannelService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketChannel); i {
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
		file_marketChannelService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketChannelResponse); i {
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
			RawDescriptor: file_marketChannelService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_marketChannelService_proto_goTypes,
		DependencyIndexes: file_marketChannelService_proto_depIdxs,
		MessageInfos:      file_marketChannelService_proto_msgTypes,
	}.Build()
	File_marketChannelService_proto = out.File
	file_marketChannelService_proto_rawDesc = nil
	file_marketChannelService_proto_goTypes = nil
	file_marketChannelService_proto_depIdxs = nil
}
