// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: roomPriceService.proto

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

type RoomPriceWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged           int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize        int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	RoomTypeId      int64  `protobuf:"varint,3,opt,name=room_type_id,json=roomTypeId,proto3" json:"room_type_id"`
	RoomPricePlanId int64  `protobuf:"varint,4,opt,name=room_price_plan_id,json=roomPricePlanId,proto3" json:"room_price_plan_id"`
	Date            string `protobuf:"bytes,5,opt,name=date,proto3" json:"date"`
	StartDate       string `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	EndDate         string `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date"`
	// @inject_tag: gorm:"-"
	RoomPricePlanIds []int64 `protobuf:"varint,9,rep,packed,name=room_price_plan_ids,json=roomPricePlanIds,proto3" json:"room_price_plan_ids" gorm:"-"`
	IsValid          bool    `protobuf:"varint,10,opt,name=is_valid,json=isValid,proto3" json:"is_valid"`
}

func (x *RoomPriceWhere) Reset() {
	*x = RoomPriceWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomPriceService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomPriceWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomPriceWhere) ProtoMessage() {}

func (x *RoomPriceWhere) ProtoReflect() protoreflect.Message {
	mi := &file_roomPriceService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomPriceWhere.ProtoReflect.Descriptor instead.
func (*RoomPriceWhere) Descriptor() ([]byte, []int) {
	return file_roomPriceService_proto_rawDescGZIP(), []int{0}
}

func (x *RoomPriceWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *RoomPriceWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *RoomPriceWhere) GetRoomTypeId() int64 {
	if x != nil {
		return x.RoomTypeId
	}
	return 0
}

func (x *RoomPriceWhere) GetRoomPricePlanId() int64 {
	if x != nil {
		return x.RoomPricePlanId
	}
	return 0
}

func (x *RoomPriceWhere) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *RoomPriceWhere) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *RoomPriceWhere) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *RoomPriceWhere) GetRoomPricePlanIds() []int64 {
	if x != nil {
		return x.RoomPricePlanIds
	}
	return nil
}

func (x *RoomPriceWhere) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

type RoomPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	RoomTypeId      int64               `protobuf:"varint,2,opt,name=room_type_id,json=roomTypeId,proto3" json:"room_type_id"`
	RoomPricePlanId int64               `protobuf:"varint,3,opt,name=room_price_plan_id,json=roomPricePlanId,proto3" json:"room_price_plan_id"`
	StartDate       string              `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	EndDate         string              `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date"`
	RepeatType      int32               `protobuf:"varint,6,opt,name=repeat_type,json=repeatType,proto3" json:"repeat_type"`
	RepeatWeeks     string              `protobuf:"bytes,7,opt,name=repeat_weeks,json=repeatWeeks,proto3" json:"repeat_weeks"`
	Price           float32             `protobuf:"fixed32,8,opt,name=price,proto3" json:"price"`
	CreatedAt       string              `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt       string              `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Vipcards        []*RoomPriceVipcard `protobuf:"bytes,12,rep,name=vipcards,proto3" json:"vipcards"`
	MinPrice        float32             `protobuf:"fixed32,13,opt,name=min_price,json=minPrice,proto3" json:"min_price"`
	MaxPrice        float32             `protobuf:"fixed32,14,opt,name=max_price,json=maxPrice,proto3" json:"max_price"`
}

func (x *RoomPrice) Reset() {
	*x = RoomPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomPriceService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomPrice) ProtoMessage() {}

func (x *RoomPrice) ProtoReflect() protoreflect.Message {
	mi := &file_roomPriceService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomPrice.ProtoReflect.Descriptor instead.
func (*RoomPrice) Descriptor() ([]byte, []int) {
	return file_roomPriceService_proto_rawDescGZIP(), []int{1}
}

func (x *RoomPrice) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RoomPrice) GetRoomTypeId() int64 {
	if x != nil {
		return x.RoomTypeId
	}
	return 0
}

func (x *RoomPrice) GetRoomPricePlanId() int64 {
	if x != nil {
		return x.RoomPricePlanId
	}
	return 0
}

func (x *RoomPrice) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *RoomPrice) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *RoomPrice) GetRepeatType() int32 {
	if x != nil {
		return x.RepeatType
	}
	return 0
}

func (x *RoomPrice) GetRepeatWeeks() string {
	if x != nil {
		return x.RepeatWeeks
	}
	return ""
}

func (x *RoomPrice) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *RoomPrice) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RoomPrice) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *RoomPrice) GetVipcards() []*RoomPriceVipcard {
	if x != nil {
		return x.Vipcards
	}
	return nil
}

func (x *RoomPrice) GetMinPrice() float32 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *RoomPrice) GetMaxPrice() float32 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

type RoomPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *common.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	Pager  *common.Pager `protobuf:"bytes,3,opt,name=pager,proto3" json:"pager"`
	Entity *RoomPrice    `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity"`
	Items  []*RoomPrice  `protobuf:"bytes,5,rep,name=items,proto3" json:"items"`
}

func (x *RoomPriceResponse) Reset() {
	*x = RoomPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomPriceService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomPriceResponse) ProtoMessage() {}

func (x *RoomPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_roomPriceService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomPriceResponse.ProtoReflect.Descriptor instead.
func (*RoomPriceResponse) Descriptor() ([]byte, []int) {
	return file_roomPriceService_proto_rawDescGZIP(), []int{2}
}

func (x *RoomPriceResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *RoomPriceResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *RoomPriceResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *RoomPriceResponse) GetEntity() *RoomPrice {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *RoomPriceResponse) GetItems() []*RoomPrice {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_roomPriceService_proto protoreflect.FileDescriptor

var file_roomPriceService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x72, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x56, 0x69, 0x70, 0x63, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x02, 0x0a, 0x0e, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x12,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x72, 0x6f, 0x6f, 0x6d, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x13, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x10, 0x72, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x22, 0xae, 0x03, 0x0a, 0x09, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x2b, 0x0a, 0x12, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f,
	0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x72,
	0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x36, 0x0a, 0x08, 0x76, 0x69, 0x70, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x56, 0x69, 0x70, 0x63, 0x61, 0x72, 0x64, 0x52, 0x08,
	0x76, 0x69, 0x70, 0x63, 0x61, 0x72, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x69, 0x6e,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x22, 0xd7, 0x01, 0x0a, 0x11, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12,
	0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x8a, 0x01, 0x0a,
	0x10, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x57, 0x68, 0x65,
	0x72, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x38, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_roomPriceService_proto_rawDescOnce sync.Once
	file_roomPriceService_proto_rawDescData = file_roomPriceService_proto_rawDesc
)

func file_roomPriceService_proto_rawDescGZIP() []byte {
	file_roomPriceService_proto_rawDescOnce.Do(func() {
		file_roomPriceService_proto_rawDescData = protoimpl.X.CompressGZIP(file_roomPriceService_proto_rawDescData)
	})
	return file_roomPriceService_proto_rawDescData
}

var file_roomPriceService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_roomPriceService_proto_goTypes = []interface{}{
	(*RoomPriceWhere)(nil),    // 0: services.RoomPriceWhere
	(*RoomPrice)(nil),         // 1: services.RoomPrice
	(*RoomPriceResponse)(nil), // 2: services.RoomPriceResponse
	(*RoomPriceVipcard)(nil),  // 3: services.RoomPriceVipcard
	(*common.Error)(nil),      // 4: common.Error
	(*common.Info)(nil),       // 5: common.Info
	(*common.Pager)(nil),      // 6: common.Pager
}
var file_roomPriceService_proto_depIdxs = []int32{
	3, // 0: services.RoomPrice.vipcards:type_name -> services.RoomPriceVipcard
	4, // 1: services.RoomPriceResponse.error:type_name -> common.Error
	5, // 2: services.RoomPriceResponse.info:type_name -> common.Info
	6, // 3: services.RoomPriceResponse.pager:type_name -> common.Pager
	1, // 4: services.RoomPriceResponse.entity:type_name -> services.RoomPrice
	1, // 5: services.RoomPriceResponse.items:type_name -> services.RoomPrice
	0, // 6: services.RoomPriceService.Get:input_type -> services.RoomPriceWhere
	1, // 7: services.RoomPriceService.Save:input_type -> services.RoomPrice
	2, // 8: services.RoomPriceService.Get:output_type -> services.RoomPriceResponse
	2, // 9: services.RoomPriceService.Save:output_type -> services.RoomPriceResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_roomPriceService_proto_init() }
func file_roomPriceService_proto_init() {
	if File_roomPriceService_proto != nil {
		return
	}
	file_roomPriceVipcardService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_roomPriceService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomPriceWhere); i {
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
		file_roomPriceService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomPrice); i {
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
		file_roomPriceService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomPriceResponse); i {
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
			RawDescriptor: file_roomPriceService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_roomPriceService_proto_goTypes,
		DependencyIndexes: file_roomPriceService_proto_depIdxs,
		MessageInfos:      file_roomPriceService_proto_msgTypes,
	}.Build()
	File_roomPriceService_proto = out.File
	file_roomPriceService_proto_rawDesc = nil
	file_roomPriceService_proto_goTypes = nil
	file_roomPriceService_proto_depIdxs = nil
}
