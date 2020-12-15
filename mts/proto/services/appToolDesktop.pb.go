// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: appToolDesktop.proto

package services

import (
	common "geiqin.saas.mts/app/proto/common"
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

type AppToolDesktopWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged         int32   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize      int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Keywords      string  `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Id            int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Ids           []int64 `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	ApplicationId int32   `protobuf:"varint,6,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	AppToolId     int32   `protobuf:"varint,7,opt,name=app_tool_id,json=appToolId,proto3" json:"app_tool_id,omitempty"`
}

func (x *AppToolDesktopWhere) Reset() {
	*x = AppToolDesktopWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appToolDesktop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppToolDesktopWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppToolDesktopWhere) ProtoMessage() {}

func (x *AppToolDesktopWhere) ProtoReflect() protoreflect.Message {
	mi := &file_appToolDesktop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppToolDesktopWhere.ProtoReflect.Descriptor instead.
func (*AppToolDesktopWhere) Descriptor() ([]byte, []int) {
	return file_appToolDesktop_proto_rawDescGZIP(), []int{0}
}

func (x *AppToolDesktopWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *AppToolDesktopWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AppToolDesktopWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *AppToolDesktopWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppToolDesktopWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *AppToolDesktopWhere) GetApplicationId() int32 {
	if x != nil {
		return x.ApplicationId
	}
	return 0
}

func (x *AppToolDesktopWhere) GetAppToolId() int32 {
	if x != nil {
		return x.AppToolId
	}
	return 0
}

type AppToolDesktop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ApplicationId int32    `protobuf:"varint,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	AppToolId     int32    `protobuf:"varint,3,opt,name=app_tool_id,json=appToolId,proto3" json:"app_tool_id,omitempty"`
	Sorting       int32    `protobuf:"varint,4,opt,name=sorting,proto3" json:"sorting,omitempty"`
	CreatedAt     string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Tool          *AppTool `protobuf:"bytes,7,opt,name=tool,proto3" json:"tool,omitempty"`
}

func (x *AppToolDesktop) Reset() {
	*x = AppToolDesktop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appToolDesktop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppToolDesktop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppToolDesktop) ProtoMessage() {}

func (x *AppToolDesktop) ProtoReflect() protoreflect.Message {
	mi := &file_appToolDesktop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppToolDesktop.ProtoReflect.Descriptor instead.
func (*AppToolDesktop) Descriptor() ([]byte, []int) {
	return file_appToolDesktop_proto_rawDescGZIP(), []int{1}
}

func (x *AppToolDesktop) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppToolDesktop) GetApplicationId() int32 {
	if x != nil {
		return x.ApplicationId
	}
	return 0
}

func (x *AppToolDesktop) GetAppToolId() int32 {
	if x != nil {
		return x.AppToolId
	}
	return 0
}

func (x *AppToolDesktop) GetSorting() int32 {
	if x != nil {
		return x.Sorting
	}
	return 0
}

func (x *AppToolDesktop) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AppToolDesktop) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *AppToolDesktop) GetTool() *AppTool {
	if x != nil {
		return x.Tool
	}
	return nil
}

type AppToolDesktopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *AppToolDesktop   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*AppToolDesktop `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error     `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info      `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *AppToolDesktopResponse) Reset() {
	*x = AppToolDesktopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appToolDesktop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppToolDesktopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppToolDesktopResponse) ProtoMessage() {}

func (x *AppToolDesktopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_appToolDesktop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppToolDesktopResponse.ProtoReflect.Descriptor instead.
func (*AppToolDesktopResponse) Descriptor() ([]byte, []int) {
	return file_appToolDesktop_proto_rawDescGZIP(), []int{2}
}

func (x *AppToolDesktopResponse) GetEntity() *AppToolDesktop {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *AppToolDesktopResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *AppToolDesktopResponse) GetItems() []*AppToolDesktop {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *AppToolDesktopResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *AppToolDesktopResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_appToolDesktop_proto protoreflect.FileDescriptor

var file_appToolDesktop_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65, 0x73, 0x6b, 0x74, 0x6f, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x61, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65,
	0x73, 0x6b, 0x74, 0x6f, 0x70, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c,
	0x49, 0x64, 0x22, 0xe6, 0x01, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65,
	0x73, 0x6b, 0x74, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b,
	0x61, 0x70, 0x70, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x61, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x6f, 0x6f, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70,
	0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x74, 0x6f, 0x6f, 0x6c, 0x22, 0xe6, 0x01, 0x0a, 0x16,
	0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65, 0x73, 0x6b, 0x74, 0x6f, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65, 0x73, 0x6b, 0x74, 0x6f, 0x70,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2e, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x44,
	0x65, 0x73, 0x6b, 0x74, 0x6f, 0x70, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x32, 0x90, 0x03, 0x0a, 0x15, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c,
	0x44, 0x65, 0x73, 0x6b, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65, 0x73, 0x6b, 0x74,
	0x6f, 0x70, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70,
	0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65, 0x73, 0x6b, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54,
	0x6f, 0x6f, 0x6c, 0x44, 0x65, 0x73, 0x6b, 0x74, 0x6f, 0x70, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a,
	0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f,
	0x6f, 0x6c, 0x44, 0x65, 0x73, 0x6b, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65,
	0x73, 0x6b, 0x74, 0x6f, 0x70, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65, 0x73,
	0x6b, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65, 0x73, 0x6b, 0x74,
	0x6f, 0x70, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65, 0x73, 0x6b, 0x74, 0x6f,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0a, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65, 0x73, 0x6b,
	0x74, 0x6f, 0x70, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41,
	0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65, 0x73, 0x6b, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_appToolDesktop_proto_rawDescOnce sync.Once
	file_appToolDesktop_proto_rawDescData = file_appToolDesktop_proto_rawDesc
)

func file_appToolDesktop_proto_rawDescGZIP() []byte {
	file_appToolDesktop_proto_rawDescOnce.Do(func() {
		file_appToolDesktop_proto_rawDescData = protoimpl.X.CompressGZIP(file_appToolDesktop_proto_rawDescData)
	})
	return file_appToolDesktop_proto_rawDescData
}

var file_appToolDesktop_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_appToolDesktop_proto_goTypes = []interface{}{
	(*AppToolDesktopWhere)(nil),    // 0: services.AppToolDesktopWhere
	(*AppToolDesktop)(nil),         // 1: services.AppToolDesktop
	(*AppToolDesktopResponse)(nil), // 2: services.AppToolDesktopResponse
	(*AppTool)(nil),                // 3: services.AppTool
	(*common.Pager)(nil),           // 4: common.Pager
	(*common.Error)(nil),           // 5: common.Error
	(*common.Info)(nil),            // 6: common.Info
}
var file_appToolDesktop_proto_depIdxs = []int32{
	3,  // 0: services.AppToolDesktop.tool:type_name -> services.AppTool
	1,  // 1: services.AppToolDesktopResponse.entity:type_name -> services.AppToolDesktop
	4,  // 2: services.AppToolDesktopResponse.pager:type_name -> common.Pager
	1,  // 3: services.AppToolDesktopResponse.items:type_name -> services.AppToolDesktop
	5,  // 4: services.AppToolDesktopResponse.error:type_name -> common.Error
	6,  // 5: services.AppToolDesktopResponse.info:type_name -> common.Info
	1,  // 6: services.AppToolDesktopService.Create:input_type -> services.AppToolDesktop
	0,  // 7: services.AppToolDesktopService.Delete:input_type -> services.AppToolDesktopWhere
	0,  // 8: services.AppToolDesktopService.List:input_type -> services.AppToolDesktopWhere
	0,  // 9: services.AppToolDesktopService.Search:input_type -> services.AppToolDesktopWhere
	1,  // 10: services.AppToolDesktopService.ChangeSort:input_type -> services.AppToolDesktop
	2,  // 11: services.AppToolDesktopService.Create:output_type -> services.AppToolDesktopResponse
	2,  // 12: services.AppToolDesktopService.Delete:output_type -> services.AppToolDesktopResponse
	2,  // 13: services.AppToolDesktopService.List:output_type -> services.AppToolDesktopResponse
	2,  // 14: services.AppToolDesktopService.Search:output_type -> services.AppToolDesktopResponse
	2,  // 15: services.AppToolDesktopService.ChangeSort:output_type -> services.AppToolDesktopResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_appToolDesktop_proto_init() }
func file_appToolDesktop_proto_init() {
	if File_appToolDesktop_proto != nil {
		return
	}
	file_appTool_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_appToolDesktop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppToolDesktopWhere); i {
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
		file_appToolDesktop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppToolDesktop); i {
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
		file_appToolDesktop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppToolDesktopResponse); i {
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
			RawDescriptor: file_appToolDesktop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_appToolDesktop_proto_goTypes,
		DependencyIndexes: file_appToolDesktop_proto_depIdxs,
		MessageInfos:      file_appToolDesktop_proto_msgTypes,
	}.Build()
	File_appToolDesktop_proto = out.File
	file_appToolDesktop_proto_rawDesc = nil
	file_appToolDesktop_proto_goTypes = nil
	file_appToolDesktop_proto_depIdxs = nil
}
