// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: wxappRoomAssistantService.proto

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

type WxappRoomAssistantWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	RoomId   int64  `protobuf:"varint,4,opt,name=room_id,json=roomId,proto3" json:"room_id"`
	Username string `protobuf:"bytes,5,opt,name=username,proto3" json:"username"`
	Nickname string `protobuf:"bytes,6,opt,name=nickname,proto3" json:"nickname"`
	// @inject_tag: gorm:"-"
	Users []*WxappRoomAssistant `protobuf:"bytes,7,rep,name=users,proto3" json:"users" gorm:"-"`
}

func (x *WxappRoomAssistantWhere) Reset() {
	*x = WxappRoomAssistantWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxappRoomAssistantService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WxappRoomAssistantWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WxappRoomAssistantWhere) ProtoMessage() {}

func (x *WxappRoomAssistantWhere) ProtoReflect() protoreflect.Message {
	mi := &file_wxappRoomAssistantService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WxappRoomAssistantWhere.ProtoReflect.Descriptor instead.
func (*WxappRoomAssistantWhere) Descriptor() ([]byte, []int) {
	return file_wxappRoomAssistantService_proto_rawDescGZIP(), []int{0}
}

func (x *WxappRoomAssistantWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *WxappRoomAssistantWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *WxappRoomAssistantWhere) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *WxappRoomAssistantWhere) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *WxappRoomAssistantWhere) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *WxappRoomAssistantWhere) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *WxappRoomAssistantWhere) GetUsers() []*WxappRoomAssistant {
	if x != nil {
		return x.Users
	}
	return nil
}

type WxappRoomAssistant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username"`
	Nickname  string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname"`
	Headimg   string `protobuf:"bytes,3,opt,name=headimg,proto3" json:"headimg"`
	Openid    string `protobuf:"bytes,4,opt,name=openid,proto3" json:"openid"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *WxappRoomAssistant) Reset() {
	*x = WxappRoomAssistant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxappRoomAssistantService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WxappRoomAssistant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WxappRoomAssistant) ProtoMessage() {}

func (x *WxappRoomAssistant) ProtoReflect() protoreflect.Message {
	mi := &file_wxappRoomAssistantService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WxappRoomAssistant.ProtoReflect.Descriptor instead.
func (*WxappRoomAssistant) Descriptor() ([]byte, []int) {
	return file_wxappRoomAssistantService_proto_rawDescGZIP(), []int{1}
}

func (x *WxappRoomAssistant) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *WxappRoomAssistant) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *WxappRoomAssistant) GetHeadimg() string {
	if x != nil {
		return x.Headimg
	}
	return ""
}

func (x *WxappRoomAssistant) GetOpenid() string {
	if x != nil {
		return x.Openid
	}
	return ""
}

func (x *WxappRoomAssistant) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type WxappRoomAssistantList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count"`
	MaxCount int32 `protobuf:"varint,2,opt,name=max_count,json=maxCount,proto3" json:"max_count"`
	// @inject_tag: gorm:"-"
	List []*WxappRoomAssistant `protobuf:"bytes,3,rep,name=list,proto3" json:"list" gorm:"-"`
}

func (x *WxappRoomAssistantList) Reset() {
	*x = WxappRoomAssistantList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxappRoomAssistantService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WxappRoomAssistantList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WxappRoomAssistantList) ProtoMessage() {}

func (x *WxappRoomAssistantList) ProtoReflect() protoreflect.Message {
	mi := &file_wxappRoomAssistantService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WxappRoomAssistantList.ProtoReflect.Descriptor instead.
func (*WxappRoomAssistantList) Descriptor() ([]byte, []int) {
	return file_wxappRoomAssistantService_proto_rawDescGZIP(), []int{2}
}

func (x *WxappRoomAssistantList) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *WxappRoomAssistantList) GetMaxCount() int32 {
	if x != nil {
		return x.MaxCount
	}
	return 0
}

func (x *WxappRoomAssistantList) GetList() []*WxappRoomAssistant {
	if x != nil {
		return x.List
	}
	return nil
}

type WxappRoomAssistantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *WxappRoomAssistant   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager         `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*WxappRoomAssistant `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error         `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info          `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *WxappRoomAssistantResponse) Reset() {
	*x = WxappRoomAssistantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxappRoomAssistantService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WxappRoomAssistantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WxappRoomAssistantResponse) ProtoMessage() {}

func (x *WxappRoomAssistantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wxappRoomAssistantService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WxappRoomAssistantResponse.ProtoReflect.Descriptor instead.
func (*WxappRoomAssistantResponse) Descriptor() ([]byte, []int) {
	return file_wxappRoomAssistantService_proto_rawDescGZIP(), []int{3}
}

func (x *WxappRoomAssistantResponse) GetEntity() *WxappRoomAssistant {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *WxappRoomAssistantResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *WxappRoomAssistantResponse) GetItems() []*WxappRoomAssistant {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *WxappRoomAssistantResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *WxappRoomAssistantResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type WxappRoomAssistantListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *WxappRoomAssistantList `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Error  *common.Error           `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
	Info   *common.Info            `protobuf:"bytes,3,opt,name=info,proto3" json:"info"`
}

func (x *WxappRoomAssistantListResponse) Reset() {
	*x = WxappRoomAssistantListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxappRoomAssistantService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WxappRoomAssistantListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WxappRoomAssistantListResponse) ProtoMessage() {}

func (x *WxappRoomAssistantListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wxappRoomAssistantService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WxappRoomAssistantListResponse.ProtoReflect.Descriptor instead.
func (*WxappRoomAssistantListResponse) Descriptor() ([]byte, []int) {
	return file_wxappRoomAssistantService_proto_rawDescGZIP(), []int{4}
}

func (x *WxappRoomAssistantListResponse) GetEntity() *WxappRoomAssistantList {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *WxappRoomAssistantListResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *WxappRoomAssistantListResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_wxappRoomAssistantService_proto protoreflect.FileDescriptor

var file_wxappRoomAssistantService_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x77, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb,
	0x01, 0x0a, 0x17, 0x57, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x73, 0x73, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x57, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x73, 0x73, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x9d, 0x01, 0x0a,
	0x12, 0x57, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68,
	0x65, 0x61, 0x64, 0x69, 0x6d, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x69, 0x6d, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7d, 0x0a, 0x16,
	0x57, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x57, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x73, 0x73, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xf2, 0x01, 0x0a, 0x1a,
	0x57, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41,
	0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x57, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x22, 0xa1, 0x01, 0x0a, 0x1e, 0x57, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57,
	0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x32, 0xf1, 0x02, 0x0a, 0x19, 0x57, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f,
	0x6f, 0x6d, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x53, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f,
	0x6d, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a,
	0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x78, 0x61, 0x70, 0x70,
	0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x78, 0x61,
	0x70, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x57,
	0x68, 0x65, 0x72, 0x65, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x57, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x57, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x55, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x28, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x78, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x6f,
	0x6d, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wxappRoomAssistantService_proto_rawDescOnce sync.Once
	file_wxappRoomAssistantService_proto_rawDescData = file_wxappRoomAssistantService_proto_rawDesc
)

func file_wxappRoomAssistantService_proto_rawDescGZIP() []byte {
	file_wxappRoomAssistantService_proto_rawDescOnce.Do(func() {
		file_wxappRoomAssistantService_proto_rawDescData = protoimpl.X.CompressGZIP(file_wxappRoomAssistantService_proto_rawDescData)
	})
	return file_wxappRoomAssistantService_proto_rawDescData
}

var file_wxappRoomAssistantService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_wxappRoomAssistantService_proto_goTypes = []interface{}{
	(*WxappRoomAssistantWhere)(nil),        // 0: services.WxappRoomAssistantWhere
	(*WxappRoomAssistant)(nil),             // 1: services.WxappRoomAssistant
	(*WxappRoomAssistantList)(nil),         // 2: services.WxappRoomAssistantList
	(*WxappRoomAssistantResponse)(nil),     // 3: services.WxappRoomAssistantResponse
	(*WxappRoomAssistantListResponse)(nil), // 4: services.WxappRoomAssistantListResponse
	(*common.Pager)(nil),                   // 5: common.Pager
	(*common.Error)(nil),                   // 6: common.Error
	(*common.Info)(nil),                    // 7: common.Info
}
var file_wxappRoomAssistantService_proto_depIdxs = []int32{
	1,  // 0: services.WxappRoomAssistantWhere.users:type_name -> services.WxappRoomAssistant
	1,  // 1: services.WxappRoomAssistantList.list:type_name -> services.WxappRoomAssistant
	1,  // 2: services.WxappRoomAssistantResponse.entity:type_name -> services.WxappRoomAssistant
	5,  // 3: services.WxappRoomAssistantResponse.pager:type_name -> common.Pager
	1,  // 4: services.WxappRoomAssistantResponse.items:type_name -> services.WxappRoomAssistant
	6,  // 5: services.WxappRoomAssistantResponse.error:type_name -> common.Error
	7,  // 6: services.WxappRoomAssistantResponse.info:type_name -> common.Info
	2,  // 7: services.WxappRoomAssistantListResponse.entity:type_name -> services.WxappRoomAssistantList
	6,  // 8: services.WxappRoomAssistantListResponse.error:type_name -> common.Error
	7,  // 9: services.WxappRoomAssistantListResponse.info:type_name -> common.Info
	0,  // 10: services.WxappRoomAssistantService.Create:input_type -> services.WxappRoomAssistantWhere
	0,  // 11: services.WxappRoomAssistantService.Update:input_type -> services.WxappRoomAssistantWhere
	0,  // 12: services.WxappRoomAssistantService.Delete:input_type -> services.WxappRoomAssistantWhere
	0,  // 13: services.WxappRoomAssistantService.List:input_type -> services.WxappRoomAssistantWhere
	3,  // 14: services.WxappRoomAssistantService.Create:output_type -> services.WxappRoomAssistantResponse
	3,  // 15: services.WxappRoomAssistantService.Update:output_type -> services.WxappRoomAssistantResponse
	3,  // 16: services.WxappRoomAssistantService.Delete:output_type -> services.WxappRoomAssistantResponse
	4,  // 17: services.WxappRoomAssistantService.List:output_type -> services.WxappRoomAssistantListResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_wxappRoomAssistantService_proto_init() }
func file_wxappRoomAssistantService_proto_init() {
	if File_wxappRoomAssistantService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wxappRoomAssistantService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WxappRoomAssistantWhere); i {
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
		file_wxappRoomAssistantService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WxappRoomAssistant); i {
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
		file_wxappRoomAssistantService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WxappRoomAssistantList); i {
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
		file_wxappRoomAssistantService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WxappRoomAssistantResponse); i {
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
		file_wxappRoomAssistantService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WxappRoomAssistantListResponse); i {
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
			RawDescriptor: file_wxappRoomAssistantService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wxappRoomAssistantService_proto_goTypes,
		DependencyIndexes: file_wxappRoomAssistantService_proto_depIdxs,
		MessageInfos:      file_wxappRoomAssistantService_proto_msgTypes,
	}.Build()
	File_wxappRoomAssistantService_proto = out.File
	file_wxappRoomAssistantService_proto_rawDesc = nil
	file_wxappRoomAssistantService_proto_goTypes = nil
	file_wxappRoomAssistantService_proto_depIdxs = nil
}
