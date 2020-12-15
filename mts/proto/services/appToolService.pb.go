// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: appToolService.proto

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

type AppToolWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged         int32   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize      int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Keywords      string  `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Id            int32   `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Ids           []int32 `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	AppToolTypeId int32   `protobuf:"varint,6,opt,name=app_tool_type_id,json=appToolTypeId,proto3" json:"app_tool_type_id,omitempty"`
	Disabled      bool    `protobuf:"varint,7,opt,name=disabled,proto3" json:"disabled,omitempty"`
}

func (x *AppToolWhere) Reset() {
	*x = AppToolWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appToolService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppToolWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppToolWhere) ProtoMessage() {}

func (x *AppToolWhere) ProtoReflect() protoreflect.Message {
	mi := &file_appToolService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppToolWhere.ProtoReflect.Descriptor instead.
func (*AppToolWhere) Descriptor() ([]byte, []int) {
	return file_appToolService_proto_rawDescGZIP(), []int{0}
}

func (x *AppToolWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *AppToolWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AppToolWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *AppToolWhere) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppToolWhere) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *AppToolWhere) GetAppToolTypeId() int32 {
	if x != nil {
		return x.AppToolTypeId
	}
	return 0
}

func (x *AppToolWhere) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

// 应用信息
type AppTool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AppToolTypeId int32  `protobuf:"varint,2,opt,name=app_tool_type_id,json=appToolTypeId,proto3" json:"app_tool_type_id,omitempty"`
	Slug          string `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Name          string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Mode          string `protobuf:"bytes,5,opt,name=mode,proto3" json:"mode,omitempty"`
	Icon          string `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	Url           string `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	Memo          string `protobuf:"bytes,8,opt,name=memo,proto3" json:"memo,omitempty"`
	Disabled      bool   `protobuf:"varint,9,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Tags          string `protobuf:"bytes,10,opt,name=tags,proto3" json:"tags,omitempty"`
	CreatedAt     string `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *AppTool) Reset() {
	*x = AppTool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appToolService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppTool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppTool) ProtoMessage() {}

func (x *AppTool) ProtoReflect() protoreflect.Message {
	mi := &file_appToolService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppTool.ProtoReflect.Descriptor instead.
func (*AppTool) Descriptor() ([]byte, []int) {
	return file_appToolService_proto_rawDescGZIP(), []int{1}
}

func (x *AppTool) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppTool) GetAppToolTypeId() int32 {
	if x != nil {
		return x.AppToolTypeId
	}
	return 0
}

func (x *AppTool) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *AppTool) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppTool) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *AppTool) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *AppTool) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AppTool) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *AppTool) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *AppTool) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *AppTool) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AppTool) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

//
type AppToolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *AppTool      `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*AppTool    `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *AppToolResponse) Reset() {
	*x = AppToolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appToolService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppToolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppToolResponse) ProtoMessage() {}

func (x *AppToolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_appToolService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppToolResponse.ProtoReflect.Descriptor instead.
func (*AppToolResponse) Descriptor() ([]byte, []int) {
	return file_appToolService_proto_rawDescGZIP(), []int{2}
}

func (x *AppToolResponse) GetEntity() *AppTool {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *AppToolResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *AppToolResponse) GetItems() []*AppTool {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *AppToolResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *AppToolResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_appToolService_proto protoreflect.FileDescriptor

var file_appToolService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x57,
	0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x27, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x5f, 0x74, 0x6f, 0x6f,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x61, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0xa6, 0x02, 0x0a, 0x07, 0x41,
	0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x5f, 0x74, 0x6f,
	0x6f, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x61, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xd1, 0x01, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xb6, 0x03, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x54,
	0x6f, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f,
	0x6c, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70,
	0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x57, 0x68, 0x65, 0x72, 0x65,
	0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54,
	0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x57,
	0x68, 0x65, 0x72, 0x65, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x16, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x57, 0x68,
	0x65, 0x72, 0x65, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41,
	0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3e, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x57, 0x68,
	0x65, 0x72, 0x65, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41,
	0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_appToolService_proto_rawDescOnce sync.Once
	file_appToolService_proto_rawDescData = file_appToolService_proto_rawDesc
)

func file_appToolService_proto_rawDescGZIP() []byte {
	file_appToolService_proto_rawDescOnce.Do(func() {
		file_appToolService_proto_rawDescData = protoimpl.X.CompressGZIP(file_appToolService_proto_rawDescData)
	})
	return file_appToolService_proto_rawDescData
}

var file_appToolService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_appToolService_proto_goTypes = []interface{}{
	(*AppToolWhere)(nil),    // 0: services.AppToolWhere
	(*AppTool)(nil),         // 1: services.AppTool
	(*AppToolResponse)(nil), // 2: services.AppToolResponse
	(*common.Pager)(nil),    // 3: common.Pager
	(*common.Error)(nil),    // 4: common.Error
	(*common.Info)(nil),     // 5: common.Info
}
var file_appToolService_proto_depIdxs = []int32{
	1,  // 0: services.AppToolResponse.entity:type_name -> services.AppTool
	3,  // 1: services.AppToolResponse.pager:type_name -> common.Pager
	1,  // 2: services.AppToolResponse.items:type_name -> services.AppTool
	4,  // 3: services.AppToolResponse.error:type_name -> common.Error
	5,  // 4: services.AppToolResponse.info:type_name -> common.Info
	1,  // 5: services.AppToolService.Create:input_type -> services.AppTool
	1,  // 6: services.AppToolService.Update:input_type -> services.AppTool
	0,  // 7: services.AppToolService.Delete:input_type -> services.AppToolWhere
	1,  // 8: services.AppToolService.Get:input_type -> services.AppTool
	0,  // 9: services.AppToolService.List:input_type -> services.AppToolWhere
	0,  // 10: services.AppToolService.Search:input_type -> services.AppToolWhere
	0,  // 11: services.AppToolService.Disable:input_type -> services.AppToolWhere
	2,  // 12: services.AppToolService.Create:output_type -> services.AppToolResponse
	2,  // 13: services.AppToolService.Update:output_type -> services.AppToolResponse
	2,  // 14: services.AppToolService.Delete:output_type -> services.AppToolResponse
	2,  // 15: services.AppToolService.Get:output_type -> services.AppToolResponse
	2,  // 16: services.AppToolService.List:output_type -> services.AppToolResponse
	2,  // 17: services.AppToolService.Search:output_type -> services.AppToolResponse
	2,  // 18: services.AppToolService.Disable:output_type -> services.AppToolResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_appToolService_proto_init() }
func file_appToolService_proto_init() {
	if File_appToolService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_appToolService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppToolWhere); i {
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
		file_appToolService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppTool); i {
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
		file_appToolService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppToolResponse); i {
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
			RawDescriptor: file_appToolService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_appToolService_proto_goTypes,
		DependencyIndexes: file_appToolService_proto_depIdxs,
		MessageInfos:      file_appToolService_proto_msgTypes,
	}.Build()
	File_appToolService_proto = out.File
	file_appToolService_proto_rawDesc = nil
	file_appToolService_proto_goTypes = nil
	file_appToolService_proto_depIdxs = nil
}
