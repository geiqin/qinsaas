// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: codeDayStatsService.proto

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

type CodeDayStatsWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id       int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	// @inject_tag: gorm:"-"
	Ids       []int64 `protobuf:"varint,4,rep,packed,name=ids,proto3" json:"ids"`
	StartDate string  `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	EndDate   string  `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date"`
}

func (x *CodeDayStatsWhere) Reset() {
	*x = CodeDayStatsWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codeDayStatsService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeDayStatsWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeDayStatsWhere) ProtoMessage() {}

func (x *CodeDayStatsWhere) ProtoReflect() protoreflect.Message {
	mi := &file_codeDayStatsService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeDayStatsWhere.ProtoReflect.Descriptor instead.
func (*CodeDayStatsWhere) Descriptor() ([]byte, []int) {
	return file_codeDayStatsService_proto_rawDescGZIP(), []int{0}
}

func (x *CodeDayStatsWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *CodeDayStatsWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CodeDayStatsWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CodeDayStatsWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *CodeDayStatsWhere) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *CodeDayStatsWhere) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type CodeDayStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Total           int64   `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
	ScanTotal       int64   `protobuf:"varint,3,opt,name=scan_total,json=scanTotal,proto3" json:"scan_total"`
	InputTotal      int64   `protobuf:"varint,4,opt,name=input_total,json=inputTotal,proto3" json:"input_total"`
	ScanProportion  float32 `protobuf:"fixed32,5,opt,name=scan_proportion,json=scanProportion,proto3" json:"scan_proportion"`
	InputProportion float32 `protobuf:"fixed32,6,opt,name=input_proportion,json=inputProportion,proto3" json:"input_proportion"`
	StatsDate       string  `protobuf:"bytes,7,opt,name=stats_date,json=statsDate,proto3" json:"stats_date"`
	CreatedAt       string  `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt       string  `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *CodeDayStats) Reset() {
	*x = CodeDayStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codeDayStatsService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeDayStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeDayStats) ProtoMessage() {}

func (x *CodeDayStats) ProtoReflect() protoreflect.Message {
	mi := &file_codeDayStatsService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeDayStats.ProtoReflect.Descriptor instead.
func (*CodeDayStats) Descriptor() ([]byte, []int) {
	return file_codeDayStatsService_proto_rawDescGZIP(), []int{1}
}

func (x *CodeDayStats) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CodeDayStats) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CodeDayStats) GetScanTotal() int64 {
	if x != nil {
		return x.ScanTotal
	}
	return 0
}

func (x *CodeDayStats) GetInputTotal() int64 {
	if x != nil {
		return x.InputTotal
	}
	return 0
}

func (x *CodeDayStats) GetScanProportion() float32 {
	if x != nil {
		return x.ScanProportion
	}
	return 0
}

func (x *CodeDayStats) GetInputProportion() float32 {
	if x != nil {
		return x.InputProportion
	}
	return 0
}

func (x *CodeDayStats) GetStatsDate() string {
	if x != nil {
		return x.StatsDate
	}
	return ""
}

func (x *CodeDayStats) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CodeDayStats) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CodeDayStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *common.Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error"`
	Info   *common.Info    `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	Pager  *common.Pager   `protobuf:"bytes,3,opt,name=pager,proto3" json:"pager"`
	Entity *CodeDayStats   `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity"`
	Items  []*CodeDayStats `protobuf:"bytes,5,rep,name=items,proto3" json:"items"`
}

func (x *CodeDayStatsResponse) Reset() {
	*x = CodeDayStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codeDayStatsService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeDayStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeDayStatsResponse) ProtoMessage() {}

func (x *CodeDayStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_codeDayStatsService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeDayStatsResponse.ProtoReflect.Descriptor instead.
func (*CodeDayStatsResponse) Descriptor() ([]byte, []int) {
	return file_codeDayStatsService_proto_rawDescGZIP(), []int{2}
}

func (x *CodeDayStatsResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CodeDayStatsResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *CodeDayStatsResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CodeDayStatsResponse) GetEntity() *CodeDayStats {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CodeDayStatsResponse) GetItems() []*CodeDayStats {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_codeDayStatsService_proto protoreflect.FileDescriptor

var file_codeDayStatsService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x64,
	0x65, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xa5, 0x02,
	0x0a, 0x0c, 0x43, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x63, 0x61, 0x6e, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x73,
	0x63, 0x61, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a,
	0x10, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xe0, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x64, 0x65, 0x44, 0x61,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xe5, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x64,
	0x65, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x42, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57,
	0x68, 0x65, 0x72, 0x65, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x44, 0x61,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57, 0x68, 0x65, 0x72, 0x65,
	0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_codeDayStatsService_proto_rawDescOnce sync.Once
	file_codeDayStatsService_proto_rawDescData = file_codeDayStatsService_proto_rawDesc
)

func file_codeDayStatsService_proto_rawDescGZIP() []byte {
	file_codeDayStatsService_proto_rawDescOnce.Do(func() {
		file_codeDayStatsService_proto_rawDescData = protoimpl.X.CompressGZIP(file_codeDayStatsService_proto_rawDescData)
	})
	return file_codeDayStatsService_proto_rawDescData
}

var file_codeDayStatsService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_codeDayStatsService_proto_goTypes = []interface{}{
	(*CodeDayStatsWhere)(nil),    // 0: services.CodeDayStatsWhere
	(*CodeDayStats)(nil),         // 1: services.CodeDayStats
	(*CodeDayStatsResponse)(nil), // 2: services.CodeDayStatsResponse
	(*common.Error)(nil),         // 3: common.Error
	(*common.Info)(nil),          // 4: common.Info
	(*common.Pager)(nil),         // 5: common.Pager
}
var file_codeDayStatsService_proto_depIdxs = []int32{
	3, // 0: services.CodeDayStatsResponse.error:type_name -> common.Error
	4, // 1: services.CodeDayStatsResponse.info:type_name -> common.Info
	5, // 2: services.CodeDayStatsResponse.pager:type_name -> common.Pager
	1, // 3: services.CodeDayStatsResponse.entity:type_name -> services.CodeDayStats
	1, // 4: services.CodeDayStatsResponse.items:type_name -> services.CodeDayStats
	0, // 5: services.CodeDayStatsService.Get:input_type -> services.CodeDayStatsWhere
	0, // 6: services.CodeDayStatsService.Search:input_type -> services.CodeDayStatsWhere
	0, // 7: services.CodeDayStatsService.List:input_type -> services.CodeDayStatsWhere
	2, // 8: services.CodeDayStatsService.Get:output_type -> services.CodeDayStatsResponse
	2, // 9: services.CodeDayStatsService.Search:output_type -> services.CodeDayStatsResponse
	2, // 10: services.CodeDayStatsService.List:output_type -> services.CodeDayStatsResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_codeDayStatsService_proto_init() }
func file_codeDayStatsService_proto_init() {
	if File_codeDayStatsService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_codeDayStatsService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeDayStatsWhere); i {
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
		file_codeDayStatsService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeDayStats); i {
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
		file_codeDayStatsService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeDayStatsResponse); i {
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
			RawDescriptor: file_codeDayStatsService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_codeDayStatsService_proto_goTypes,
		DependencyIndexes: file_codeDayStatsService_proto_depIdxs,
		MessageInfos:      file_codeDayStatsService_proto_msgTypes,
	}.Build()
	File_codeDayStatsService_proto = out.File
	file_codeDayStatsService_proto_rawDesc = nil
	file_codeDayStatsService_proto_goTypes = nil
	file_codeDayStatsService_proto_depIdxs = nil
}
