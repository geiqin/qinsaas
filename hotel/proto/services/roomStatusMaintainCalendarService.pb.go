// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: roomStatusMaintainCalendarService.proto

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

type RoomStatusMaintainCalendarWhere struct {
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
}

func (x *RoomStatusMaintainCalendarWhere) Reset() {
	*x = RoomStatusMaintainCalendarWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomStatusMaintainCalendarService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomStatusMaintainCalendarWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomStatusMaintainCalendarWhere) ProtoMessage() {}

func (x *RoomStatusMaintainCalendarWhere) ProtoReflect() protoreflect.Message {
	mi := &file_roomStatusMaintainCalendarService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomStatusMaintainCalendarWhere.ProtoReflect.Descriptor instead.
func (*RoomStatusMaintainCalendarWhere) Descriptor() ([]byte, []int) {
	return file_roomStatusMaintainCalendarService_proto_rawDescGZIP(), []int{0}
}

func (x *RoomStatusMaintainCalendarWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *RoomStatusMaintainCalendarWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *RoomStatusMaintainCalendarWhere) GetRoomTypeId() int64 {
	if x != nil {
		return x.RoomTypeId
	}
	return 0
}

func (x *RoomStatusMaintainCalendarWhere) GetRoomPricePlanId() int64 {
	if x != nil {
		return x.RoomPricePlanId
	}
	return 0
}

func (x *RoomStatusMaintainCalendarWhere) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *RoomStatusMaintainCalendarWhere) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *RoomStatusMaintainCalendarWhere) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type RoomStatusMaintainCalendar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomTypeId   int64  `protobuf:"varint,1,opt,name=room_type_id,json=roomTypeId,proto3" json:"room_type_id"`
	RoomTypeName string `protobuf:"bytes,2,opt,name=room_type_name,json=roomTypeName,proto3" json:"room_type_name"`
	// @inject_tag: gorm:"-"
	RoomPricePlanList []*RoomPricePlanByMaintainList `protobuf:"bytes,3,rep,name=room_price_plan_list,json=roomPricePlanList,proto3" json:"room_price_plan_list" gorm:"-"`
}

func (x *RoomStatusMaintainCalendar) Reset() {
	*x = RoomStatusMaintainCalendar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomStatusMaintainCalendarService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomStatusMaintainCalendar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomStatusMaintainCalendar) ProtoMessage() {}

func (x *RoomStatusMaintainCalendar) ProtoReflect() protoreflect.Message {
	mi := &file_roomStatusMaintainCalendarService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomStatusMaintainCalendar.ProtoReflect.Descriptor instead.
func (*RoomStatusMaintainCalendar) Descriptor() ([]byte, []int) {
	return file_roomStatusMaintainCalendarService_proto_rawDescGZIP(), []int{1}
}

func (x *RoomStatusMaintainCalendar) GetRoomTypeId() int64 {
	if x != nil {
		return x.RoomTypeId
	}
	return 0
}

func (x *RoomStatusMaintainCalendar) GetRoomTypeName() string {
	if x != nil {
		return x.RoomTypeName
	}
	return ""
}

func (x *RoomStatusMaintainCalendar) GetRoomPricePlanList() []*RoomPricePlanByMaintainList {
	if x != nil {
		return x.RoomPricePlanList
	}
	return nil
}

type RoomPricePlanByMaintainList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomPricePlanId   int64  `protobuf:"varint,1,opt,name=room_price_plan_id,json=roomPricePlanId,proto3" json:"room_price_plan_id"`
	RoomPricePlanName string `protobuf:"bytes,2,opt,name=room_price_plan_name,json=roomPricePlanName,proto3" json:"room_price_plan_name"`
	// @inject_tag: gorm:"-"
	RoomStatusList []*RoomStatusMaintainList `protobuf:"bytes,3,rep,name=room_status_list,json=roomStatusList,proto3" json:"room_status_list" gorm:"-"`
}

func (x *RoomPricePlanByMaintainList) Reset() {
	*x = RoomPricePlanByMaintainList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomStatusMaintainCalendarService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomPricePlanByMaintainList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomPricePlanByMaintainList) ProtoMessage() {}

func (x *RoomPricePlanByMaintainList) ProtoReflect() protoreflect.Message {
	mi := &file_roomStatusMaintainCalendarService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomPricePlanByMaintainList.ProtoReflect.Descriptor instead.
func (*RoomPricePlanByMaintainList) Descriptor() ([]byte, []int) {
	return file_roomStatusMaintainCalendarService_proto_rawDescGZIP(), []int{2}
}

func (x *RoomPricePlanByMaintainList) GetRoomPricePlanId() int64 {
	if x != nil {
		return x.RoomPricePlanId
	}
	return 0
}

func (x *RoomPricePlanByMaintainList) GetRoomPricePlanName() string {
	if x != nil {
		return x.RoomPricePlanName
	}
	return ""
}

func (x *RoomPricePlanByMaintainList) GetRoomStatusList() []*RoomStatusMaintainList {
	if x != nil {
		return x.RoomStatusList
	}
	return nil
}

type RoomStatusMaintainList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CanSalesCount int32   `protobuf:"varint,1,opt,name=can_sales_count,json=canSalesCount,proto3" json:"can_sales_count"` // 剩余房量
	SalesCount    int32   `protobuf:"varint,2,opt,name=sales_count,json=salesCount,proto3" json:"sales_count"`            // 已售房量
	Price         float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price"`                                       // 价格
	MinPrice      float32 `protobuf:"fixed32,4,opt,name=min_price,json=minPrice,proto3" json:"min_price"`                 // 最低价
	MaxPrice      float32 `protobuf:"fixed32,5,opt,name=max_price,json=maxPrice,proto3" json:"max_price"`                 // 最高价
	Status        int32   `protobuf:"varint,6,opt,name=status,proto3" json:"status"`                                      // 状态: 1-开房,2-关房
	Date          string  `protobuf:"bytes,7,opt,name=date,proto3" json:"date"`                                           // 日期
}

func (x *RoomStatusMaintainList) Reset() {
	*x = RoomStatusMaintainList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomStatusMaintainCalendarService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomStatusMaintainList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomStatusMaintainList) ProtoMessage() {}

func (x *RoomStatusMaintainList) ProtoReflect() protoreflect.Message {
	mi := &file_roomStatusMaintainCalendarService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomStatusMaintainList.ProtoReflect.Descriptor instead.
func (*RoomStatusMaintainList) Descriptor() ([]byte, []int) {
	return file_roomStatusMaintainCalendarService_proto_rawDescGZIP(), []int{3}
}

func (x *RoomStatusMaintainList) GetCanSalesCount() int32 {
	if x != nil {
		return x.CanSalesCount
	}
	return 0
}

func (x *RoomStatusMaintainList) GetSalesCount() int32 {
	if x != nil {
		return x.SalesCount
	}
	return 0
}

func (x *RoomStatusMaintainList) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *RoomStatusMaintainList) GetMinPrice() float32 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *RoomStatusMaintainList) GetMaxPrice() float32 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

func (x *RoomStatusMaintainList) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RoomStatusMaintainList) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type RoomStatusMaintainCalendarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *common.Error                 `protobuf:"bytes,1,opt,name=error,proto3" json:"error"`
	Info   *common.Info                  `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	Pager  *common.Pager                 `protobuf:"bytes,3,opt,name=pager,proto3" json:"pager"`
	Entity *RoomStatusMaintainCalendar   `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity"`
	Items  []*RoomStatusMaintainCalendar `protobuf:"bytes,5,rep,name=items,proto3" json:"items"`
}

func (x *RoomStatusMaintainCalendarResponse) Reset() {
	*x = RoomStatusMaintainCalendarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomStatusMaintainCalendarService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomStatusMaintainCalendarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomStatusMaintainCalendarResponse) ProtoMessage() {}

func (x *RoomStatusMaintainCalendarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_roomStatusMaintainCalendarService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomStatusMaintainCalendarResponse.ProtoReflect.Descriptor instead.
func (*RoomStatusMaintainCalendarResponse) Descriptor() ([]byte, []int) {
	return file_roomStatusMaintainCalendarService_proto_rawDescGZIP(), []int{4}
}

func (x *RoomStatusMaintainCalendarResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *RoomStatusMaintainCalendarResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *RoomStatusMaintainCalendarResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *RoomStatusMaintainCalendarResponse) GetEntity() *RoomStatusMaintainCalendar {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *RoomStatusMaintainCalendarResponse) GetItems() []*RoomStatusMaintainCalendar {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_roomStatusMaintainCalendarService_proto protoreflect.FileDescriptor

var file_roomStatusMaintainCalendarService_proto_rawDesc = []byte{
	0x0a, 0x27, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x61, 0x69, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01, 0x0a, 0x1f, 0x52, 0x6f, 0x6f, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x6c,
	0x65, 0x6e, 0x64, 0x61, 0x72, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a,
	0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12,
	0x2b, 0x0a, 0x12, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6c,
	0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x72, 0x6f, 0x6f,
	0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xbc, 0x01, 0x0a, 0x1a, 0x52,
	0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x56, 0x0a, 0x14, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f,
	0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x42, 0x79, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x11, 0x72, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x50, 0x6c, 0x61, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xc7, 0x01, 0x0a, 0x1b, 0x52, 0x6f,
	0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x42, 0x79, 0x4d, 0x61, 0x69,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x12, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x72, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x14, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x10, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x0e, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0xdd, 0x01, 0x0a, 0x16, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26,
	0x0a, 0x0f, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x61, 0x6e, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x61, 0x6c,
	0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61,
	0x78, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d,
	0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x22, 0x8a, 0x02, 0x0a, 0x22, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x61, 0x69, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52,
	0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x32, 0x84, 0x01, 0x0a, 0x21, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x2c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_roomStatusMaintainCalendarService_proto_rawDescOnce sync.Once
	file_roomStatusMaintainCalendarService_proto_rawDescData = file_roomStatusMaintainCalendarService_proto_rawDesc
)

func file_roomStatusMaintainCalendarService_proto_rawDescGZIP() []byte {
	file_roomStatusMaintainCalendarService_proto_rawDescOnce.Do(func() {
		file_roomStatusMaintainCalendarService_proto_rawDescData = protoimpl.X.CompressGZIP(file_roomStatusMaintainCalendarService_proto_rawDescData)
	})
	return file_roomStatusMaintainCalendarService_proto_rawDescData
}

var file_roomStatusMaintainCalendarService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_roomStatusMaintainCalendarService_proto_goTypes = []interface{}{
	(*RoomStatusMaintainCalendarWhere)(nil),    // 0: services.RoomStatusMaintainCalendarWhere
	(*RoomStatusMaintainCalendar)(nil),         // 1: services.RoomStatusMaintainCalendar
	(*RoomPricePlanByMaintainList)(nil),        // 2: services.RoomPricePlanByMaintainList
	(*RoomStatusMaintainList)(nil),             // 3: services.RoomStatusMaintainList
	(*RoomStatusMaintainCalendarResponse)(nil), // 4: services.RoomStatusMaintainCalendarResponse
	(*common.Error)(nil),                       // 5: common.Error
	(*common.Info)(nil),                        // 6: common.Info
	(*common.Pager)(nil),                       // 7: common.Pager
}
var file_roomStatusMaintainCalendarService_proto_depIdxs = []int32{
	2, // 0: services.RoomStatusMaintainCalendar.room_price_plan_list:type_name -> services.RoomPricePlanByMaintainList
	3, // 1: services.RoomPricePlanByMaintainList.room_status_list:type_name -> services.RoomStatusMaintainList
	5, // 2: services.RoomStatusMaintainCalendarResponse.error:type_name -> common.Error
	6, // 3: services.RoomStatusMaintainCalendarResponse.info:type_name -> common.Info
	7, // 4: services.RoomStatusMaintainCalendarResponse.pager:type_name -> common.Pager
	1, // 5: services.RoomStatusMaintainCalendarResponse.entity:type_name -> services.RoomStatusMaintainCalendar
	1, // 6: services.RoomStatusMaintainCalendarResponse.items:type_name -> services.RoomStatusMaintainCalendar
	0, // 7: services.RoomStatusMaintainCalendarService.List:input_type -> services.RoomStatusMaintainCalendarWhere
	4, // 8: services.RoomStatusMaintainCalendarService.List:output_type -> services.RoomStatusMaintainCalendarResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_roomStatusMaintainCalendarService_proto_init() }
func file_roomStatusMaintainCalendarService_proto_init() {
	if File_roomStatusMaintainCalendarService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_roomStatusMaintainCalendarService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomStatusMaintainCalendarWhere); i {
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
		file_roomStatusMaintainCalendarService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomStatusMaintainCalendar); i {
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
		file_roomStatusMaintainCalendarService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomPricePlanByMaintainList); i {
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
		file_roomStatusMaintainCalendarService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomStatusMaintainList); i {
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
		file_roomStatusMaintainCalendarService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomStatusMaintainCalendarResponse); i {
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
			RawDescriptor: file_roomStatusMaintainCalendarService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_roomStatusMaintainCalendarService_proto_goTypes,
		DependencyIndexes: file_roomStatusMaintainCalendarService_proto_depIdxs,
		MessageInfos:      file_roomStatusMaintainCalendarService_proto_msgTypes,
	}.Build()
	File_roomStatusMaintainCalendarService_proto = out.File
	file_roomStatusMaintainCalendarService_proto_rawDesc = nil
	file_roomStatusMaintainCalendarService_proto_goTypes = nil
	file_roomStatusMaintainCalendarService_proto_depIdxs = nil
}