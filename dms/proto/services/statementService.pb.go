// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: statementService.proto

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

type Statement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TargetId      int64   `protobuf:"varint,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	SourceId      int64   `protobuf:"varint,3,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	SourceType    string  `protobuf:"bytes,4,opt,name=source_type,json=sourceType,proto3" json:"source_type,omitempty"`
	SourceMethod  int32   `protobuf:"varint,5,opt,name=source_method,json=sourceMethod,proto3" json:"source_method,omitempty"`
	Money         float32 `protobuf:"fixed32,6,opt,name=money,proto3" json:"money,omitempty"`
	Type          string  `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Balance       float32 `protobuf:"fixed32,8,opt,name=balance,proto3" json:"balance,omitempty"`
	CreatedAt     string  `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string  `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DistributorId int64   `protobuf:"varint,11,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id,omitempty"`
}

func (x *Statement) Reset() {
	*x = Statement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statementService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statement) ProtoMessage() {}

func (x *Statement) ProtoReflect() protoreflect.Message {
	mi := &file_statementService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statement.ProtoReflect.Descriptor instead.
func (*Statement) Descriptor() ([]byte, []int) {
	return file_statementService_proto_rawDescGZIP(), []int{0}
}

func (x *Statement) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Statement) GetTargetId() int64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

func (x *Statement) GetSourceId() int64 {
	if x != nil {
		return x.SourceId
	}
	return 0
}

func (x *Statement) GetSourceType() string {
	if x != nil {
		return x.SourceType
	}
	return ""
}

func (x *Statement) GetSourceMethod() int32 {
	if x != nil {
		return x.SourceMethod
	}
	return 0
}

func (x *Statement) GetMoney() float32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *Statement) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Statement) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Statement) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Statement) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Statement) GetDistributorId() int64 {
	if x != nil {
		return x.DistributorId
	}
	return 0
}

type StatementWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Top      int32 `protobuf:"varint,3,opt,name=top,proto3" json:"top,omitempty"`
	//base params
	Id           int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Ids          []int64 `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Type         string  `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	TargetId     int64   `protobuf:"varint,7,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	SourceType   string  `protobuf:"bytes,8,opt,name=source_type,json=sourceType,proto3" json:"source_type,omitempty"`
	SourceMethod int32   `protobuf:"varint,9,opt,name=source_method,json=sourceMethod,proto3" json:"source_method,omitempty"`
}

func (x *StatementWhere) Reset() {
	*x = StatementWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statementService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatementWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatementWhere) ProtoMessage() {}

func (x *StatementWhere) ProtoReflect() protoreflect.Message {
	mi := &file_statementService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatementWhere.ProtoReflect.Descriptor instead.
func (*StatementWhere) Descriptor() ([]byte, []int) {
	return file_statementService_proto_rawDescGZIP(), []int{1}
}

func (x *StatementWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *StatementWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *StatementWhere) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *StatementWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StatementWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *StatementWhere) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *StatementWhere) GetTargetId() int64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

func (x *StatementWhere) GetSourceType() string {
	if x != nil {
		return x.SourceType
	}
	return ""
}

func (x *StatementWhere) GetSourceMethod() int32 {
	if x != nil {
		return x.SourceMethod
	}
	return 0
}

type StatementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Statement    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*Statement  `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *StatementResponse) Reset() {
	*x = StatementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statementService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatementResponse) ProtoMessage() {}

func (x *StatementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_statementService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatementResponse.ProtoReflect.Descriptor instead.
func (*StatementResponse) Descriptor() ([]byte, []int) {
	return file_statementService_proto_rawDescGZIP(), []int{2}
}

func (x *StatementResponse) GetEntity() *Statement {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *StatementResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *StatementResponse) GetItems() []*Statement {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *StatementResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *StatementResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_statementService_proto protoreflect.FileDescriptor

var file_statementService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x02, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0xee, 0x01, 0x0a,
	0x0e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x74, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0xd7, 0x01,
	0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x92, 0x01, 0x0a, 0x12, 0x4d, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0xc9, 0x01, 0x0a,
	0x10, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x68, 0x65, 0x72, 0x65,
	0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x37, 0x0a, 0x07, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_statementService_proto_rawDescOnce sync.Once
	file_statementService_proto_rawDescData = file_statementService_proto_rawDesc
)

func file_statementService_proto_rawDescGZIP() []byte {
	file_statementService_proto_rawDescOnce.Do(func() {
		file_statementService_proto_rawDescData = protoimpl.X.CompressGZIP(file_statementService_proto_rawDescData)
	})
	return file_statementService_proto_rawDescData
}

var file_statementService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_statementService_proto_goTypes = []interface{}{
	(*Statement)(nil),         // 0: services.Statement
	(*StatementWhere)(nil),    // 1: services.StatementWhere
	(*StatementResponse)(nil), // 2: services.StatementResponse
	(*common.Pager)(nil),      // 3: common.Pager
	(*common.Error)(nil),      // 4: common.Error
	(*common.Info)(nil),       // 5: common.Info
	(*common.Empty)(nil),      // 6: common.Empty
}
var file_statementService_proto_depIdxs = []int32{
	0,  // 0: services.StatementResponse.entity:type_name -> services.Statement
	3,  // 1: services.StatementResponse.pager:type_name -> common.Pager
	0,  // 2: services.StatementResponse.items:type_name -> services.Statement
	4,  // 3: services.StatementResponse.error:type_name -> common.Error
	5,  // 4: services.StatementResponse.info:type_name -> common.Info
	0,  // 5: services.MyStatementService.Get:input_type -> services.Statement
	1,  // 6: services.MyStatementService.Search:input_type -> services.StatementWhere
	0,  // 7: services.StatementService.Get:input_type -> services.Statement
	1,  // 8: services.StatementService.Search:input_type -> services.StatementWhere
	6,  // 9: services.StatementService.Upgrade:input_type -> common.Empty
	2,  // 10: services.MyStatementService.Get:output_type -> services.StatementResponse
	2,  // 11: services.MyStatementService.Search:output_type -> services.StatementResponse
	2,  // 12: services.StatementService.Get:output_type -> services.StatementResponse
	2,  // 13: services.StatementService.Search:output_type -> services.StatementResponse
	2,  // 14: services.StatementService.Upgrade:output_type -> services.StatementResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_statementService_proto_init() }
func file_statementService_proto_init() {
	if File_statementService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_statementService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Statement); i {
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
		file_statementService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatementWhere); i {
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
		file_statementService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatementResponse); i {
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
			RawDescriptor: file_statementService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_statementService_proto_goTypes,
		DependencyIndexes: file_statementService_proto_depIdxs,
		MessageInfos:      file_statementService_proto_msgTypes,
	}.Build()
	File_statementService_proto = out.File
	file_statementService_proto_rawDesc = nil
	file_statementService_proto_goTypes = nil
	file_statementService_proto_depIdxs = nil
}
