// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: roomPriceCalendarService.proto

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

type RoomPriceCalendarWhere struct {
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

func (x *RoomPriceCalendarWhere) Reset() {
	*x = RoomPriceCalendarWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomPriceCalendarService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomPriceCalendarWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomPriceCalendarWhere) ProtoMessage() {}

func (x *RoomPriceCalendarWhere) ProtoReflect() protoreflect.Message {
	mi := &file_roomPriceCalendarService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomPriceCalendarWhere.ProtoReflect.Descriptor instead.
func (*RoomPriceCalendarWhere) Descriptor() ([]byte, []int) {
	return file_roomPriceCalendarService_proto_rawDescGZIP(), []int{0}
}

func (x *RoomPriceCalendarWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *RoomPriceCalendarWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *RoomPriceCalendarWhere) GetRoomTypeId() int64 {
	if x != nil {
		return x.RoomTypeId
	}
	return 0
}

func (x *RoomPriceCalendarWhere) GetRoomPricePlanId() int64 {
	if x != nil {
		return x.RoomPricePlanId
	}
	return 0
}

func (x *RoomPriceCalendarWhere) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *RoomPriceCalendarWhere) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *RoomPriceCalendarWhere) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type RoomPriceCalendar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CanSalesCount int32   `protobuf:"varint,1,opt,name=can_sales_count,json=canSalesCount,proto3" json:"can_sales_count"` // 剩余房量
	Price         float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price"`                                       // 标准价
	MinPrice      float32 `protobuf:"fixed32,3,opt,name=min_price,json=minPrice,proto3" json:"min_price"`                 // 最低价
	MaxPrice      float32 `protobuf:"fixed32,4,opt,name=max_price,json=maxPrice,proto3" json:"max_price"`                 // 最高价
	Date          string  `protobuf:"bytes,5,opt,name=date,proto3" json:"date"`                                           // 日期
}

func (x *RoomPriceCalendar) Reset() {
	*x = RoomPriceCalendar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomPriceCalendarService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomPriceCalendar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomPriceCalendar) ProtoMessage() {}

func (x *RoomPriceCalendar) ProtoReflect() protoreflect.Message {
	mi := &file_roomPriceCalendarService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomPriceCalendar.ProtoReflect.Descriptor instead.
func (*RoomPriceCalendar) Descriptor() ([]byte, []int) {
	return file_roomPriceCalendarService_proto_rawDescGZIP(), []int{1}
}

func (x *RoomPriceCalendar) GetCanSalesCount() int32 {
	if x != nil {
		return x.CanSalesCount
	}
	return 0
}

func (x *RoomPriceCalendar) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *RoomPriceCalendar) GetMinPrice() float32 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *RoomPriceCalendar) GetMaxPrice() float32 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

func (x *RoomPriceCalendar) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type RoomPriceCalendarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *common.Error        `protobuf:"bytes,1,opt,name=error,proto3" json:"error"`
	Info   *common.Info         `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	Pager  *common.Pager        `protobuf:"bytes,3,opt,name=pager,proto3" json:"pager"`
	Entity *RoomPriceCalendar   `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity"`
	Items  []*RoomPriceCalendar `protobuf:"bytes,5,rep,name=items,proto3" json:"items"`
}

func (x *RoomPriceCalendarResponse) Reset() {
	*x = RoomPriceCalendarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomPriceCalendarService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomPriceCalendarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomPriceCalendarResponse) ProtoMessage() {}

func (x *RoomPriceCalendarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_roomPriceCalendarService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomPriceCalendarResponse.ProtoReflect.Descriptor instead.
func (*RoomPriceCalendarResponse) Descriptor() ([]byte, []int) {
	return file_roomPriceCalendarService_proto_rawDescGZIP(), []int{2}
}

func (x *RoomPriceCalendarResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *RoomPriceCalendarResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *RoomPriceCalendarResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *RoomPriceCalendarResponse) GetEntity() *RoomPriceCalendar {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *RoomPriceCalendarResponse) GetItems() []*RoomPriceCalendar {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_roomPriceCalendarService_proto protoreflect.FileDescriptor

var file_roomPriceCalendarService_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01,
	0x0a, 0x16, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a,
	0x12, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x72, 0x6f, 0x6f, 0x6d, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x11, 0x52, 0x6f, 0x6f,
	0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x12, 0x26,
	0x0a, 0x0f, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x61, 0x6e, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x61,
	0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0xef, 0x01, 0x0a, 0x19, 0x52,
	0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12,
	0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6c,
	0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x69, 0x0a, 0x18,
	0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x57, 0x68, 0x65,
	0x72, 0x65, 0x1a, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_roomPriceCalendarService_proto_rawDescOnce sync.Once
	file_roomPriceCalendarService_proto_rawDescData = file_roomPriceCalendarService_proto_rawDesc
)

func file_roomPriceCalendarService_proto_rawDescGZIP() []byte {
	file_roomPriceCalendarService_proto_rawDescOnce.Do(func() {
		file_roomPriceCalendarService_proto_rawDescData = protoimpl.X.CompressGZIP(file_roomPriceCalendarService_proto_rawDescData)
	})
	return file_roomPriceCalendarService_proto_rawDescData
}

var file_roomPriceCalendarService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_roomPriceCalendarService_proto_goTypes = []interface{}{
	(*RoomPriceCalendarWhere)(nil),    // 0: services.RoomPriceCalendarWhere
	(*RoomPriceCalendar)(nil),         // 1: services.RoomPriceCalendar
	(*RoomPriceCalendarResponse)(nil), // 2: services.RoomPriceCalendarResponse
	(*common.Error)(nil),              // 3: common.Error
	(*common.Info)(nil),               // 4: common.Info
	(*common.Pager)(nil),              // 5: common.Pager
}
var file_roomPriceCalendarService_proto_depIdxs = []int32{
	3, // 0: services.RoomPriceCalendarResponse.error:type_name -> common.Error
	4, // 1: services.RoomPriceCalendarResponse.info:type_name -> common.Info
	5, // 2: services.RoomPriceCalendarResponse.pager:type_name -> common.Pager
	1, // 3: services.RoomPriceCalendarResponse.entity:type_name -> services.RoomPriceCalendar
	1, // 4: services.RoomPriceCalendarResponse.items:type_name -> services.RoomPriceCalendar
	0, // 5: services.RoomPriceCalendarService.List:input_type -> services.RoomPriceCalendarWhere
	2, // 6: services.RoomPriceCalendarService.List:output_type -> services.RoomPriceCalendarResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_roomPriceCalendarService_proto_init() }
func file_roomPriceCalendarService_proto_init() {
	if File_roomPriceCalendarService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_roomPriceCalendarService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomPriceCalendarWhere); i {
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
		file_roomPriceCalendarService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomPriceCalendar); i {
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
		file_roomPriceCalendarService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomPriceCalendarResponse); i {
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
			RawDescriptor: file_roomPriceCalendarService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_roomPriceCalendarService_proto_goTypes,
		DependencyIndexes: file_roomPriceCalendarService_proto_depIdxs,
		MessageInfos:      file_roomPriceCalendarService_proto_msgTypes,
	}.Build()
	File_roomPriceCalendarService_proto = out.File
	file_roomPriceCalendarService_proto_rawDesc = nil
	file_roomPriceCalendarService_proto_goTypes = nil
	file_roomPriceCalendarService_proto_depIdxs = nil
}
