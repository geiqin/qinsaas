// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: roomkeepService.proto

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

type RoomKeepWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged      int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize   int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	RoomTypeId int64  `protobuf:"varint,3,opt,name=room_type_id,json=roomTypeId,proto3" json:"room_type_id"`
	RoomId     int64  `protobuf:"varint,4,opt,name=room_id,json=roomId,proto3" json:"room_id"`
	Date       string `protobuf:"bytes,5,opt,name=date,proto3" json:"date"`
	StartDate  string `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	EndDate    string `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date"`
	// @inject_tag: gorm:"-"
	RoomTypeIds []int64 `protobuf:"varint,8,rep,packed,name=room_type_ids,json=roomTypeIds,proto3" json:"room_type_ids" gorm:"-"`
}

func (x *RoomKeepWhere) Reset() {
	*x = RoomKeepWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomkeepService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomKeepWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomKeepWhere) ProtoMessage() {}

func (x *RoomKeepWhere) ProtoReflect() protoreflect.Message {
	mi := &file_roomkeepService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomKeepWhere.ProtoReflect.Descriptor instead.
func (*RoomKeepWhere) Descriptor() ([]byte, []int) {
	return file_roomkeepService_proto_rawDescGZIP(), []int{0}
}

func (x *RoomKeepWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *RoomKeepWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *RoomKeepWhere) GetRoomTypeId() int64 {
	if x != nil {
		return x.RoomTypeId
	}
	return 0
}

func (x *RoomKeepWhere) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *RoomKeepWhere) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *RoomKeepWhere) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *RoomKeepWhere) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *RoomKeepWhere) GetRoomTypeIds() []int64 {
	if x != nil {
		return x.RoomTypeIds
	}
	return nil
}

type RoomKeep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	RoomTypeId  int64  `protobuf:"varint,2,opt,name=room_type_id,json=roomTypeId,proto3" json:"room_type_id"`
	RoomId      int64  `protobuf:"varint,3,opt,name=room_id,json=roomId,proto3" json:"room_id"`
	Date        string `protobuf:"bytes,4,opt,name=date,proto3" json:"date"`
	StartDate   string `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	EndDate     string `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date"`
	RepeatType  int32  `protobuf:"varint,7,opt,name=repeat_type,json=repeatType,proto3" json:"repeat_type"`   // 重复类型: 1-按天, 2-按周
	RepeatWeeks string `protobuf:"bytes,8,opt,name=repeat_weeks,json=repeatWeeks,proto3" json:"repeat_weeks"` // 重复的星期, 用逗号隔开, 如: 0(周天),1,2,3,4,5,6
	Status      int32  `protobuf:"varint,9,opt,name=status,proto3" json:"status"`
	Memo        string `protobuf:"bytes,10,opt,name=memo,proto3" json:"memo"`
	CreatedAt   string `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt   string `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	// @inject_tag: gorm:"-"
	RoomTypeIds []int64 `protobuf:"varint,13,rep,packed,name=room_type_ids,json=roomTypeIds,proto3" json:"room_type_ids" gorm:"-"`
}

func (x *RoomKeep) Reset() {
	*x = RoomKeep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomkeepService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomKeep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomKeep) ProtoMessage() {}

func (x *RoomKeep) ProtoReflect() protoreflect.Message {
	mi := &file_roomkeepService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomKeep.ProtoReflect.Descriptor instead.
func (*RoomKeep) Descriptor() ([]byte, []int) {
	return file_roomkeepService_proto_rawDescGZIP(), []int{1}
}

func (x *RoomKeep) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RoomKeep) GetRoomTypeId() int64 {
	if x != nil {
		return x.RoomTypeId
	}
	return 0
}

func (x *RoomKeep) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *RoomKeep) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *RoomKeep) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *RoomKeep) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *RoomKeep) GetRepeatType() int32 {
	if x != nil {
		return x.RepeatType
	}
	return 0
}

func (x *RoomKeep) GetRepeatWeeks() string {
	if x != nil {
		return x.RepeatWeeks
	}
	return ""
}

func (x *RoomKeep) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RoomKeep) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *RoomKeep) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RoomKeep) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *RoomKeep) GetRoomTypeIds() []int64 {
	if x != nil {
		return x.RoomTypeIds
	}
	return nil
}

type RoomKeepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *common.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	Pager  *common.Pager `protobuf:"bytes,3,opt,name=pager,proto3" json:"pager"`
	Entity *RoomKeep     `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity"`
	Items  []*RoomKeep   `protobuf:"bytes,5,rep,name=items,proto3" json:"items"`
}

func (x *RoomKeepResponse) Reset() {
	*x = RoomKeepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomkeepService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomKeepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomKeepResponse) ProtoMessage() {}

func (x *RoomKeepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_roomkeepService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomKeepResponse.ProtoReflect.Descriptor instead.
func (*RoomKeepResponse) Descriptor() ([]byte, []int) {
	return file_roomkeepService_proto_rawDescGZIP(), []int{2}
}

func (x *RoomKeepResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *RoomKeepResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *RoomKeepResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *RoomKeepResponse) GetEntity() *RoomKeep {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *RoomKeepResponse) GetItems() []*RoomKeep {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_roomkeepService_proto protoreflect.FileDescriptor

var file_roomkeepService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x6f, 0x6f, 0x6d, 0x6b, 0x65, 0x65, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x0d, 0x52, 0x6f, 0x6f, 0x6d, 0x4b, 0x65, 0x65,
	0x70, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x6f, 0x6f, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x64, 0x73, 0x22, 0xf5, 0x02, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x4b,
	0x65, 0x65, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x0b, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x73, 0x22, 0xd4,
	0x01, 0x0a, 0x10, 0x52, 0x6f, 0x6f, 0x6d, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x2a, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4b,
	0x65, 0x65, 0x70, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xc5, 0x01, 0x0a, 0x0f, 0x52, 0x6f, 0x6f, 0x6d, 0x4b, 0x65,
	0x65, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x53, 0x61, 0x76,
	0x65, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x4b, 0x65, 0x65, 0x70, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x09, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x61, 0x76, 0x65, 0x12, 0x12,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4b, 0x65,
	0x65, 0x70, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4b, 0x65, 0x65, 0x70, 0x57, 0x68, 0x65, 0x72,
	0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_roomkeepService_proto_rawDescOnce sync.Once
	file_roomkeepService_proto_rawDescData = file_roomkeepService_proto_rawDesc
)

func file_roomkeepService_proto_rawDescGZIP() []byte {
	file_roomkeepService_proto_rawDescOnce.Do(func() {
		file_roomkeepService_proto_rawDescData = protoimpl.X.CompressGZIP(file_roomkeepService_proto_rawDescData)
	})
	return file_roomkeepService_proto_rawDescData
}

var file_roomkeepService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_roomkeepService_proto_goTypes = []interface{}{
	(*RoomKeepWhere)(nil),    // 0: services.RoomKeepWhere
	(*RoomKeep)(nil),         // 1: services.RoomKeep
	(*RoomKeepResponse)(nil), // 2: services.RoomKeepResponse
	(*common.Error)(nil),     // 3: common.Error
	(*common.Info)(nil),      // 4: common.Info
	(*common.Pager)(nil),     // 5: common.Pager
}
var file_roomkeepService_proto_depIdxs = []int32{
	3, // 0: services.RoomKeepResponse.error:type_name -> common.Error
	4, // 1: services.RoomKeepResponse.info:type_name -> common.Info
	5, // 2: services.RoomKeepResponse.pager:type_name -> common.Pager
	1, // 3: services.RoomKeepResponse.entity:type_name -> services.RoomKeep
	1, // 4: services.RoomKeepResponse.items:type_name -> services.RoomKeep
	1, // 5: services.RoomKeepService.Save:input_type -> services.RoomKeep
	1, // 6: services.RoomKeepService.BatchSave:input_type -> services.RoomKeep
	0, // 7: services.RoomKeepService.Cancel:input_type -> services.RoomKeepWhere
	2, // 8: services.RoomKeepService.Save:output_type -> services.RoomKeepResponse
	2, // 9: services.RoomKeepService.BatchSave:output_type -> services.RoomKeepResponse
	2, // 10: services.RoomKeepService.Cancel:output_type -> services.RoomKeepResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_roomkeepService_proto_init() }
func file_roomkeepService_proto_init() {
	if File_roomkeepService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_roomkeepService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomKeepWhere); i {
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
		file_roomkeepService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomKeep); i {
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
		file_roomkeepService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomKeepResponse); i {
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
			RawDescriptor: file_roomkeepService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_roomkeepService_proto_goTypes,
		DependencyIndexes: file_roomkeepService_proto_depIdxs,
		MessageInfos:      file_roomkeepService_proto_msgTypes,
	}.Build()
	File_roomkeepService_proto = out.File
	file_roomkeepService_proto_rawDesc = nil
	file_roomkeepService_proto_goTypes = nil
	file_roomkeepService_proto_depIdxs = nil
}