// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: appToolPersonalService.proto

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

type AppToolPersonalWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged      int32   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize   int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Keywords   string  `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Id         int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Ids        []int64 `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	StoreId    int64   `protobuf:"varint,6,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	UserId     int64   `protobuf:"varint,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AppToolId  int32   `protobuf:"varint,8,opt,name=app_tool_id,json=appToolId,proto3" json:"app_tool_id,omitempty"`
	AppToolIds []int32 `protobuf:"varint,9,rep,packed,name=app_tool_ids,json=appToolIds,proto3" json:"app_tool_ids,omitempty"`
}

func (x *AppToolPersonalWhere) Reset() {
	*x = AppToolPersonalWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appToolPersonalService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppToolPersonalWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppToolPersonalWhere) ProtoMessage() {}

func (x *AppToolPersonalWhere) ProtoReflect() protoreflect.Message {
	mi := &file_appToolPersonalService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppToolPersonalWhere.ProtoReflect.Descriptor instead.
func (*AppToolPersonalWhere) Descriptor() ([]byte, []int) {
	return file_appToolPersonalService_proto_rawDescGZIP(), []int{0}
}

func (x *AppToolPersonalWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *AppToolPersonalWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AppToolPersonalWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *AppToolPersonalWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppToolPersonalWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *AppToolPersonalWhere) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *AppToolPersonalWhere) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AppToolPersonalWhere) GetAppToolId() int32 {
	if x != nil {
		return x.AppToolId
	}
	return 0
}

func (x *AppToolPersonalWhere) GetAppToolIds() []int32 {
	if x != nil {
		return x.AppToolIds
	}
	return nil
}

type AppToolPersonal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//    int64 store_id = 2;
	//    int64 user_id = 3;
	AppToolId int32    `protobuf:"varint,4,opt,name=app_tool_id,json=appToolId,proto3" json:"app_tool_id,omitempty"`
	Sorting   int32    `protobuf:"varint,5,opt,name=sorting,proto3" json:"sorting,omitempty"`
	CreatedAt string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Tool      *AppTool `protobuf:"bytes,8,opt,name=tool,proto3" json:"tool,omitempty"`
}

func (x *AppToolPersonal) Reset() {
	*x = AppToolPersonal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appToolPersonalService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppToolPersonal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppToolPersonal) ProtoMessage() {}

func (x *AppToolPersonal) ProtoReflect() protoreflect.Message {
	mi := &file_appToolPersonalService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppToolPersonal.ProtoReflect.Descriptor instead.
func (*AppToolPersonal) Descriptor() ([]byte, []int) {
	return file_appToolPersonalService_proto_rawDescGZIP(), []int{1}
}

func (x *AppToolPersonal) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppToolPersonal) GetAppToolId() int32 {
	if x != nil {
		return x.AppToolId
	}
	return 0
}

func (x *AppToolPersonal) GetSorting() int32 {
	if x != nil {
		return x.Sorting
	}
	return 0
}

func (x *AppToolPersonal) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AppToolPersonal) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *AppToolPersonal) GetTool() *AppTool {
	if x != nil {
		return x.Tool
	}
	return nil
}

type AppToolPersonalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *AppToolPersonal   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*AppToolPersonal `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *AppToolPersonalResponse) Reset() {
	*x = AppToolPersonalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appToolPersonalService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppToolPersonalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppToolPersonalResponse) ProtoMessage() {}

func (x *AppToolPersonalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_appToolPersonalService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppToolPersonalResponse.ProtoReflect.Descriptor instead.
func (*AppToolPersonalResponse) Descriptor() ([]byte, []int) {
	return file_appToolPersonalService_proto_rawDescGZIP(), []int{2}
}

func (x *AppToolPersonalResponse) GetEntity() *AppToolPersonal {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *AppToolPersonalResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *AppToolPersonalResponse) GetItems() []*AppToolPersonal {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *AppToolPersonalResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *AppToolPersonalResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_appToolPersonalService_proto protoreflect.FileDescriptor

var file_appToolPersonalService_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x61, 0x70, 0x70,
	0x54, 0x6f, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xfd, 0x01, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x49, 0x64,
	0x73, 0x22, 0xc0, 0x01, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x74, 0x6f, 0x6f,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x70, 0x70, 0x54,
	0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x25, 0x0a,
	0x04, 0x74, 0x6f, 0x6f, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x04,
	0x74, 0x6f, 0x6f, 0x6c, 0x22, 0xe9, 0x01, 0x0a, 0x17, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54,
	0x6f, 0x6f, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x6c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x32, 0xb2, 0x01, 0x0a, 0x16, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x53,
	0x61, 0x76, 0x65, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41,
	0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x57, 0x68,
	0x65, 0x72, 0x65, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41,
	0x70, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54,
	0x6f, 0x6f, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x57, 0x68, 0x65, 0x72, 0x65,
	0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x54,
	0x6f, 0x6f, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_appToolPersonalService_proto_rawDescOnce sync.Once
	file_appToolPersonalService_proto_rawDescData = file_appToolPersonalService_proto_rawDesc
)

func file_appToolPersonalService_proto_rawDescGZIP() []byte {
	file_appToolPersonalService_proto_rawDescOnce.Do(func() {
		file_appToolPersonalService_proto_rawDescData = protoimpl.X.CompressGZIP(file_appToolPersonalService_proto_rawDescData)
	})
	return file_appToolPersonalService_proto_rawDescData
}

var file_appToolPersonalService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_appToolPersonalService_proto_goTypes = []interface{}{
	(*AppToolPersonalWhere)(nil),    // 0: services.AppToolPersonalWhere
	(*AppToolPersonal)(nil),         // 1: services.AppToolPersonal
	(*AppToolPersonalResponse)(nil), // 2: services.AppToolPersonalResponse
	(*AppTool)(nil),                 // 3: services.AppTool
	(*common.Pager)(nil),            // 4: common.Pager
	(*common.Error)(nil),            // 5: common.Error
	(*common.Info)(nil),             // 6: common.Info
}
var file_appToolPersonalService_proto_depIdxs = []int32{
	3, // 0: services.AppToolPersonal.tool:type_name -> services.AppTool
	1, // 1: services.AppToolPersonalResponse.entity:type_name -> services.AppToolPersonal
	4, // 2: services.AppToolPersonalResponse.pager:type_name -> common.Pager
	1, // 3: services.AppToolPersonalResponse.items:type_name -> services.AppToolPersonal
	5, // 4: services.AppToolPersonalResponse.error:type_name -> common.Error
	6, // 5: services.AppToolPersonalResponse.info:type_name -> common.Info
	0, // 6: services.AppToolPersonalService.Save:input_type -> services.AppToolPersonalWhere
	0, // 7: services.AppToolPersonalService.List:input_type -> services.AppToolPersonalWhere
	2, // 8: services.AppToolPersonalService.Save:output_type -> services.AppToolPersonalResponse
	2, // 9: services.AppToolPersonalService.List:output_type -> services.AppToolPersonalResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_appToolPersonalService_proto_init() }
func file_appToolPersonalService_proto_init() {
	if File_appToolPersonalService_proto != nil {
		return
	}
	file_appToolService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_appToolPersonalService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppToolPersonalWhere); i {
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
		file_appToolPersonalService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppToolPersonal); i {
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
		file_appToolPersonalService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppToolPersonalResponse); i {
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
			RawDescriptor: file_appToolPersonalService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_appToolPersonalService_proto_goTypes,
		DependencyIndexes: file_appToolPersonalService_proto_depIdxs,
		MessageInfos:      file_appToolPersonalService_proto_msgTypes,
	}.Build()
	File_appToolPersonalService_proto = out.File
	file_appToolPersonalService_proto_rawDesc = nil
	file_appToolPersonalService_proto_goTypes = nil
	file_appToolPersonalService_proto_depIdxs = nil
}