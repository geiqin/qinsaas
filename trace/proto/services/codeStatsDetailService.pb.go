// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: codeStatsDetailService.proto

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

type CodeStatsDetailWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id       int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	// @inject_tag: gorm:"-"
	Ids      []int64 `protobuf:"varint,4,rep,packed,name=ids,proto3" json:"ids"`
	CodingId int64   `protobuf:"varint,5,opt,name=coding_id,json=codingId,proto3" json:"coding_id"`
}

func (x *CodeStatsDetailWhere) Reset() {
	*x = CodeStatsDetailWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codeStatsDetailService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeStatsDetailWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeStatsDetailWhere) ProtoMessage() {}

func (x *CodeStatsDetailWhere) ProtoReflect() protoreflect.Message {
	mi := &file_codeStatsDetailService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeStatsDetailWhere.ProtoReflect.Descriptor instead.
func (*CodeStatsDetailWhere) Descriptor() ([]byte, []int) {
	return file_codeStatsDetailService_proto_rawDescGZIP(), []int{0}
}

func (x *CodeStatsDetailWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *CodeStatsDetailWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CodeStatsDetailWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CodeStatsDetailWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *CodeStatsDetailWhere) GetCodingId() int64 {
	if x != nil {
		return x.CodingId
	}
	return 0
}

type CodeStatsDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CodingId        int64   `protobuf:"varint,2,opt,name=coding_id,json=codingId,proto3" json:"coding_id"`
	Total           int64   `protobuf:"varint,3,opt,name=total,proto3" json:"total"`
	ScanTotal       int64   `protobuf:"varint,4,opt,name=scan_total,json=scanTotal,proto3" json:"scan_total"`
	InputTotal      int64   `protobuf:"varint,5,opt,name=input_total,json=inputTotal,proto3" json:"input_total"`
	ScanProportion  float32 `protobuf:"fixed32,6,opt,name=scan_proportion,json=scanProportion,proto3" json:"scan_proportion"`
	InputProportion float32 `protobuf:"fixed32,7,opt,name=input_proportion,json=inputProportion,proto3" json:"input_proportion"`
	FirstVerifyAt   string  `protobuf:"bytes,8,opt,name=first_verify_at,json=firstVerifyAt,proto3" json:"first_verify_at"`
	CreatedAt       string  `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt       string  `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *CodeStatsDetail) Reset() {
	*x = CodeStatsDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codeStatsDetailService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeStatsDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeStatsDetail) ProtoMessage() {}

func (x *CodeStatsDetail) ProtoReflect() protoreflect.Message {
	mi := &file_codeStatsDetailService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeStatsDetail.ProtoReflect.Descriptor instead.
func (*CodeStatsDetail) Descriptor() ([]byte, []int) {
	return file_codeStatsDetailService_proto_rawDescGZIP(), []int{1}
}

func (x *CodeStatsDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CodeStatsDetail) GetCodingId() int64 {
	if x != nil {
		return x.CodingId
	}
	return 0
}

func (x *CodeStatsDetail) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CodeStatsDetail) GetScanTotal() int64 {
	if x != nil {
		return x.ScanTotal
	}
	return 0
}

func (x *CodeStatsDetail) GetInputTotal() int64 {
	if x != nil {
		return x.InputTotal
	}
	return 0
}

func (x *CodeStatsDetail) GetScanProportion() float32 {
	if x != nil {
		return x.ScanProportion
	}
	return 0
}

func (x *CodeStatsDetail) GetInputProportion() float32 {
	if x != nil {
		return x.InputProportion
	}
	return 0
}

func (x *CodeStatsDetail) GetFirstVerifyAt() string {
	if x != nil {
		return x.FirstVerifyAt
	}
	return ""
}

func (x *CodeStatsDetail) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CodeStatsDetail) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CodeStatsDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *common.Error      `protobuf:"bytes,1,opt,name=error,proto3" json:"error"`
	Info   *common.Info       `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	Pager  *common.Pager      `protobuf:"bytes,3,opt,name=pager,proto3" json:"pager"`
	Entity *CodeStatsDetail   `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity"`
	Items  []*CodeStatsDetail `protobuf:"bytes,5,rep,name=items,proto3" json:"items"`
}

func (x *CodeStatsDetailResponse) Reset() {
	*x = CodeStatsDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codeStatsDetailService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeStatsDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeStatsDetailResponse) ProtoMessage() {}

func (x *CodeStatsDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_codeStatsDetailService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeStatsDetailResponse.ProtoReflect.Descriptor instead.
func (*CodeStatsDetailResponse) Descriptor() ([]byte, []int) {
	return file_codeStatsDetailService_proto_rawDescGZIP(), []int{2}
}

func (x *CodeStatsDetailResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CodeStatsDetailResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *CodeStatsDetailResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CodeStatsDetailResponse) GetEntity() *CodeStatsDetail {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CodeStatsDetailResponse) GetItems() []*CodeStatsDetail {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_codeStatsDetailService_proto protoreflect.FileDescriptor

var file_codeStatsDetailService_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x14,
	0x43, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x57,
	0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0xce, 0x02, 0x0a, 0x0f, 0x43, 0x6f, 0x64, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x73, 0x63, 0x61, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x27, 0x0a,
	0x0f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x73, 0x63, 0x61, 0x6e, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xe9, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x64, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x31, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x32, 0xfa, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x57, 0x68, 0x65,
	0x72, 0x65, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x21, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_codeStatsDetailService_proto_rawDescOnce sync.Once
	file_codeStatsDetailService_proto_rawDescData = file_codeStatsDetailService_proto_rawDesc
)

func file_codeStatsDetailService_proto_rawDescGZIP() []byte {
	file_codeStatsDetailService_proto_rawDescOnce.Do(func() {
		file_codeStatsDetailService_proto_rawDescData = protoimpl.X.CompressGZIP(file_codeStatsDetailService_proto_rawDescData)
	})
	return file_codeStatsDetailService_proto_rawDescData
}

var file_codeStatsDetailService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_codeStatsDetailService_proto_goTypes = []interface{}{
	(*CodeStatsDetailWhere)(nil),    // 0: services.CodeStatsDetailWhere
	(*CodeStatsDetail)(nil),         // 1: services.CodeStatsDetail
	(*CodeStatsDetailResponse)(nil), // 2: services.CodeStatsDetailResponse
	(*common.Error)(nil),            // 3: common.Error
	(*common.Info)(nil),             // 4: common.Info
	(*common.Pager)(nil),            // 5: common.Pager
}
var file_codeStatsDetailService_proto_depIdxs = []int32{
	3, // 0: services.CodeStatsDetailResponse.error:type_name -> common.Error
	4, // 1: services.CodeStatsDetailResponse.info:type_name -> common.Info
	5, // 2: services.CodeStatsDetailResponse.pager:type_name -> common.Pager
	1, // 3: services.CodeStatsDetailResponse.entity:type_name -> services.CodeStatsDetail
	1, // 4: services.CodeStatsDetailResponse.items:type_name -> services.CodeStatsDetail
	0, // 5: services.CodeStatsDetailService.Get:input_type -> services.CodeStatsDetailWhere
	0, // 6: services.CodeStatsDetailService.Search:input_type -> services.CodeStatsDetailWhere
	0, // 7: services.CodeStatsDetailService.List:input_type -> services.CodeStatsDetailWhere
	2, // 8: services.CodeStatsDetailService.Get:output_type -> services.CodeStatsDetailResponse
	2, // 9: services.CodeStatsDetailService.Search:output_type -> services.CodeStatsDetailResponse
	2, // 10: services.CodeStatsDetailService.List:output_type -> services.CodeStatsDetailResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_codeStatsDetailService_proto_init() }
func file_codeStatsDetailService_proto_init() {
	if File_codeStatsDetailService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_codeStatsDetailService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeStatsDetailWhere); i {
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
		file_codeStatsDetailService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeStatsDetail); i {
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
		file_codeStatsDetailService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeStatsDetailResponse); i {
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
			RawDescriptor: file_codeStatsDetailService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_codeStatsDetailService_proto_goTypes,
		DependencyIndexes: file_codeStatsDetailService_proto_depIdxs,
		MessageInfos:      file_codeStatsDetailService_proto_msgTypes,
	}.Build()
	File_codeStatsDetailService_proto = out.File
	file_codeStatsDetailService_proto_rawDesc = nil
	file_codeStatsDetailService_proto_goTypes = nil
	file_codeStatsDetailService_proto_depIdxs = nil
}
