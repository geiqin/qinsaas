// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: integralService.proto

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

type IntegralWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged       int32   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize    int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting     string  `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Keywords    string  `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Id          int64   `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	Ids         []int64 `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	CustomerId  int64   `protobuf:"varint,7,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CustomerIds []int64 `protobuf:"varint,8,rep,packed,name=customer_ids,json=customerIds,proto3" json:"customer_ids,omitempty"`
	Points      int32   `protobuf:"varint,9,opt,name=points,proto3" json:"points,omitempty"`
	Memo        string  `protobuf:"bytes,10,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *IntegralWhere) Reset() {
	*x = IntegralWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegralWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegralWhere) ProtoMessage() {}

func (x *IntegralWhere) ProtoReflect() protoreflect.Message {
	mi := &file_integralService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegralWhere.ProtoReflect.Descriptor instead.
func (*IntegralWhere) Descriptor() ([]byte, []int) {
	return file_integralService_proto_rawDescGZIP(), []int{0}
}

func (x *IntegralWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *IntegralWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *IntegralWhere) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *IntegralWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *IntegralWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IntegralWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *IntegralWhere) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *IntegralWhere) GetCustomerIds() []int64 {
	if x != nil {
		return x.CustomerIds
	}
	return nil
}

func (x *IntegralWhere) GetPoints() int32 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *IntegralWhere) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

type Integral struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId    int64  `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Points        int64  `protobuf:"varint,3,opt,name=points,proto3" json:"points,omitempty"`
	TotalPoints   int64  `protobuf:"varint,4,opt,name=total_points,json=totalPoints,proto3" json:"total_points,omitempty"`
	UsedPoints    int64  `protobuf:"varint,5,opt,name=used_points,json=usedPoints,proto3" json:"used_points,omitempty"`
	ExpiredPoints int64  `protobuf:"varint,6,opt,name=expired_points,json=expiredPoints,proto3" json:"expired_points,omitempty"`
	FrozenPoints  int64  `protobuf:"varint,7,opt,name=frozen_points,json=frozenPoints,proto3" json:"frozen_points,omitempty"`
	ProtectPoints int64  `protobuf:"varint,8,opt,name=protect_points,json=protectPoints,proto3" json:"protect_points,omitempty"`
	CreatedAt     string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Integral) Reset() {
	*x = Integral{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Integral) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Integral) ProtoMessage() {}

func (x *Integral) ProtoReflect() protoreflect.Message {
	mi := &file_integralService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Integral.ProtoReflect.Descriptor instead.
func (*Integral) Descriptor() ([]byte, []int) {
	return file_integralService_proto_rawDescGZIP(), []int{1}
}

func (x *Integral) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Integral) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Integral) GetPoints() int64 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *Integral) GetTotalPoints() int64 {
	if x != nil {
		return x.TotalPoints
	}
	return 0
}

func (x *Integral) GetUsedPoints() int64 {
	if x != nil {
		return x.UsedPoints
	}
	return 0
}

func (x *Integral) GetExpiredPoints() int64 {
	if x != nil {
		return x.ExpiredPoints
	}
	return 0
}

func (x *Integral) GetFrozenPoints() int64 {
	if x != nil {
		return x.FrozenPoints
	}
	return 0
}

func (x *Integral) GetProtectPoints() int64 {
	if x != nil {
		return x.ProtectPoints
	}
	return 0
}

func (x *Integral) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Integral) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type IntegralResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Integral     `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*Integral   `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *IntegralResponse) Reset() {
	*x = IntegralResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegralResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegralResponse) ProtoMessage() {}

func (x *IntegralResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integralService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegralResponse.ProtoReflect.Descriptor instead.
func (*IntegralResponse) Descriptor() ([]byte, []int) {
	return file_integralService_proto_rawDescGZIP(), []int{2}
}

func (x *IntegralResponse) GetEntity() *Integral {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *IntegralResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *IntegralResponse) GetItems() []*Integral {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *IntegralResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *IntegralResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_integralService_proto protoreflect.FileDescriptor

var file_integralService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x02, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x6c, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x22, 0xc8, 0x02, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73,
	0x65, 0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x75, 0x73, 0x65, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x7a, 0x65,
	0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd4, 0x01, 0x0a,
	0x10, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x32, 0x8c, 0x02, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x6c, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x6c, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x03, 0x49, 0x6e, 0x63, 0x12, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x6c, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x03, 0x44, 0x65, 0x63, 0x12, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x57,
	0x68, 0x65, 0x72, 0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x32, 0x51, 0x0a, 0x11, 0x4d, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x6c, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_integralService_proto_rawDescOnce sync.Once
	file_integralService_proto_rawDescData = file_integralService_proto_rawDesc
)

func file_integralService_proto_rawDescGZIP() []byte {
	file_integralService_proto_rawDescOnce.Do(func() {
		file_integralService_proto_rawDescData = protoimpl.X.CompressGZIP(file_integralService_proto_rawDescData)
	})
	return file_integralService_proto_rawDescData
}

var file_integralService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_integralService_proto_goTypes = []interface{}{
	(*IntegralWhere)(nil),    // 0: services.IntegralWhere
	(*Integral)(nil),         // 1: services.Integral
	(*IntegralResponse)(nil), // 2: services.IntegralResponse
	(*common.Pager)(nil),     // 3: common.Pager
	(*common.Error)(nil),     // 4: common.Error
	(*common.Info)(nil),      // 5: common.Info
}
var file_integralService_proto_depIdxs = []int32{
	1,  // 0: services.IntegralResponse.entity:type_name -> services.Integral
	3,  // 1: services.IntegralResponse.pager:type_name -> common.Pager
	1,  // 2: services.IntegralResponse.items:type_name -> services.Integral
	4,  // 3: services.IntegralResponse.error:type_name -> common.Error
	5,  // 4: services.IntegralResponse.info:type_name -> common.Info
	0,  // 5: services.IntegralService.Search:input_type -> services.IntegralWhere
	0,  // 6: services.IntegralService.Get:input_type -> services.IntegralWhere
	0,  // 7: services.IntegralService.Inc:input_type -> services.IntegralWhere
	0,  // 8: services.IntegralService.Dec:input_type -> services.IntegralWhere
	0,  // 9: services.MyIntegralService.Get:input_type -> services.IntegralWhere
	2,  // 10: services.IntegralService.Search:output_type -> services.IntegralResponse
	2,  // 11: services.IntegralService.Get:output_type -> services.IntegralResponse
	2,  // 12: services.IntegralService.Inc:output_type -> services.IntegralResponse
	2,  // 13: services.IntegralService.Dec:output_type -> services.IntegralResponse
	2,  // 14: services.MyIntegralService.Get:output_type -> services.IntegralResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_integralService_proto_init() }
func file_integralService_proto_init() {
	if File_integralService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_integralService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegralWhere); i {
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
		file_integralService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Integral); i {
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
		file_integralService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegralResponse); i {
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
			RawDescriptor: file_integralService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_integralService_proto_goTypes,
		DependencyIndexes: file_integralService_proto_depIdxs,
		MessageInfos:      file_integralService_proto_msgTypes,
	}.Build()
	File_integralService_proto = out.File
	file_integralService_proto_rawDesc = nil
	file_integralService_proto_goTypes = nil
	file_integralService_proto_depIdxs = nil
}
