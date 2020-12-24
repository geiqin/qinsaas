// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: statisticsService.proto

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

type StatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string `protobuf:"bytes,1,opt,name=type,proto3" json:"type"`
	Days        int32  `protobuf:"varint,2,opt,name=days,proto3" json:"days"`
	Paged       int32  `protobuf:"varint,3,opt,name=paged,proto3" json:"paged"`
	PageSize    int32  `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	CustomerId  int64  `protobuf:"varint,5,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	StartAt     string `protobuf:"bytes,6,opt,name=start_at,json=startAt,proto3" json:"start_at"`
	EndAt       string `protobuf:"bytes,7,opt,name=end_at,json=endAt,proto3" json:"end_at"`
	StatisticAt string `protobuf:"bytes,8,opt,name=statistic_at,json=statisticAt,proto3" json:"statistic_at"`
}

func (x *StatsRequest) Reset() {
	*x = StatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statisticsService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsRequest) ProtoMessage() {}

func (x *StatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_statisticsService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsRequest.ProtoReflect.Descriptor instead.
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return file_statisticsService_proto_rawDescGZIP(), []int{0}
}

func (x *StatsRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *StatsRequest) GetDays() int32 {
	if x != nil {
		return x.Days
	}
	return 0
}

func (x *StatsRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *StatsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *StatsRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *StatsRequest) GetStartAt() string {
	if x != nil {
		return x.StartAt
	}
	return ""
}

func (x *StatsRequest) GetEndAt() string {
	if x != nil {
		return x.EndAt
	}
	return ""
}

func (x *StatsRequest) GetStatisticAt() string {
	if x != nil {
		return x.StatisticAt
	}
	return ""
}

var File_statisticsService_proto protoreflect.FileDescriptor

var file_statisticsService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x6e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x41, 0x74, 0x32, 0x69, 0x0a, 0x13, 0x4d, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x09, 0x55, 0x6e,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0xb0,
	0x05, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x61, 0x69,
	0x6c, 0x79, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x61, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a,
	0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x16, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x09,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x61, 0x79, 0x73, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x09, 0x55, 0x6e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x64, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4a, 0x0a, 0x0d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0d,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a,
	0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x6e, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x64, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x05, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_statisticsService_proto_rawDescOnce sync.Once
	file_statisticsService_proto_rawDescData = file_statisticsService_proto_rawDesc
)

func file_statisticsService_proto_rawDescGZIP() []byte {
	file_statisticsService_proto_rawDescOnce.Do(func() {
		file_statisticsService_proto_rawDescData = protoimpl.X.CompressGZIP(file_statisticsService_proto_rawDescData)
	})
	return file_statisticsService_proto_rawDescData
}

var file_statisticsService_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_statisticsService_proto_goTypes = []interface{}{
	(*StatsRequest)(nil),                      // 0: services.StatsRequest
	(*common.Empty)(nil),                      // 1: common.Empty
	(*CustomerOrderHandledStatsResponse)(nil), // 2: services.CustomerOrderHandledStatsResponse
	(*OrderDayStatsResponse)(nil),             // 3: services.OrderDayStatsResponse
	(*OrderStatsResponse)(nil),                // 4: services.OrderStatsResponse
	(*OrderHandledStatsResponse)(nil),         // 5: services.OrderHandledStatsResponse
	(*CustomerGoodsResponse)(nil),             // 6: services.CustomerGoodsResponse
	(*CustomerOrderStatsResponse)(nil),        // 7: services.CustomerOrderStatsResponse
}
var file_statisticsService_proto_depIdxs = []int32{
	0,  // 0: services.MyStatisticsService.Unhandled:input_type -> services.StatsRequest
	0,  // 1: services.StatisticsService.OrderDaily:input_type -> services.StatsRequest
	0,  // 2: services.StatisticsService.OrderTotal:input_type -> services.StatsRequest
	0,  // 3: services.StatisticsService.TodayTotal:input_type -> services.StatsRequest
	0,  // 4: services.StatisticsService.OrderDays:input_type -> services.StatsRequest
	0,  // 5: services.StatisticsService.Unhandled:input_type -> services.StatsRequest
	0,  // 6: services.StatisticsService.CustomerGoods:input_type -> services.StatsRequest
	0,  // 7: services.StatisticsService.CustomerOrder:input_type -> services.StatsRequest
	0,  // 8: services.StatisticsService.CustomerUnhandled:input_type -> services.StatsRequest
	1,  // 9: services.StatisticsService.Reset:input_type -> common.Empty
	2,  // 10: services.MyStatisticsService.Unhandled:output_type -> services.CustomerOrderHandledStatsResponse
	3,  // 11: services.StatisticsService.OrderDaily:output_type -> services.OrderDayStatsResponse
	4,  // 12: services.StatisticsService.OrderTotal:output_type -> services.OrderStatsResponse
	3,  // 13: services.StatisticsService.TodayTotal:output_type -> services.OrderDayStatsResponse
	3,  // 14: services.StatisticsService.OrderDays:output_type -> services.OrderDayStatsResponse
	5,  // 15: services.StatisticsService.Unhandled:output_type -> services.OrderHandledStatsResponse
	6,  // 16: services.StatisticsService.CustomerGoods:output_type -> services.CustomerGoodsResponse
	7,  // 17: services.StatisticsService.CustomerOrder:output_type -> services.CustomerOrderStatsResponse
	2,  // 18: services.StatisticsService.CustomerUnhandled:output_type -> services.CustomerOrderHandledStatsResponse
	4,  // 19: services.StatisticsService.Reset:output_type -> services.OrderStatsResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_statisticsService_proto_init() }
func file_statisticsService_proto_init() {
	if File_statisticsService_proto != nil {
		return
	}
	file_orderStatsService_proto_init()
	file_orderUnhandledStatsService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_statisticsService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsRequest); i {
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
			RawDescriptor: file_statisticsService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_statisticsService_proto_goTypes,
		DependencyIndexes: file_statisticsService_proto_depIdxs,
		MessageInfos:      file_statisticsService_proto_msgTypes,
	}.Build()
	File_statisticsService_proto = out.File
	file_statisticsService_proto_rawDesc = nil
	file_statisticsService_proto_goTypes = nil
	file_statisticsService_proto_depIdxs = nil
}
