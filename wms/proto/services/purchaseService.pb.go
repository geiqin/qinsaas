// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: purchaseService.proto

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

type PurchaseWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged       int32   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize    int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Id          int64   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Ids         []int64 `protobuf:"varint,4,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Type        string  `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Status      int32   `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Ok          bool    `protobuf:"varint,7,opt,name=ok,proto3" json:"ok,omitempty"`
	Failure     string  `protobuf:"bytes,8,opt,name=failure,proto3" json:"failure,omitempty"`
	PurchaseId  int64   `protobuf:"varint,9,opt,name=purchase_id,json=purchaseId,proto3" json:"purchase_id,omitempty"`
	PurchaseIds []int64 `protobuf:"varint,10,rep,packed,name=purchase_ids,json=purchaseIds,proto3" json:"purchase_ids,omitempty"`
}

func (x *PurchaseWhere) Reset() {
	*x = PurchaseWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchaseService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseWhere) ProtoMessage() {}

func (x *PurchaseWhere) ProtoReflect() protoreflect.Message {
	mi := &file_purchaseService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseWhere.ProtoReflect.Descriptor instead.
func (*PurchaseWhere) Descriptor() ([]byte, []int) {
	return file_purchaseService_proto_rawDescGZIP(), []int{0}
}

func (x *PurchaseWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *PurchaseWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PurchaseWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PurchaseWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *PurchaseWhere) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PurchaseWhere) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PurchaseWhere) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *PurchaseWhere) GetFailure() string {
	if x != nil {
		return x.Failure
	}
	return ""
}

func (x *PurchaseWhere) GetPurchaseId() int64 {
	if x != nil {
		return x.PurchaseId
	}
	return 0
}

func (x *PurchaseWhere) GetPurchaseIds() []int64 {
	if x != nil {
		return x.PurchaseIds
	}
	return nil
}

type Purchase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PurchaseSn    string            `protobuf:"bytes,2,opt,name=purchase_sn,json=purchaseSn,proto3" json:"purchase_sn,omitempty"`
	Type          string            `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	SourceNo      string            `protobuf:"bytes,4,opt,name=source_no,json=sourceNo,proto3" json:"source_no,omitempty"`
	UserId        int64             `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Memo          string            `protobuf:"bytes,6,opt,name=memo,proto3" json:"memo,omitempty"`
	RowTotal      int32             `protobuf:"varint,7,opt,name=row_total,json=rowTotal,proto3" json:"row_total,omitempty"`
	QuantityTotal int32             `protobuf:"varint,8,opt,name=quantity_total,json=quantityTotal,proto3" json:"quantity_total,omitempty"`
	HandledAt     string            `protobuf:"bytes,9,opt,name=handled_at,json=handledAt,proto3" json:"handled_at,omitempty"`
	Status        int32             `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	Failure       string            `protobuf:"bytes,11,opt,name=failure,proto3" json:"failure,omitempty"`
	CreatedAt     string            `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string            `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Details       []*PurchaseDetail `protobuf:"bytes,14,rep,name=details,proto3" json:"details,omitempty"`
	Ok            bool              `protobuf:"varint,15,opt,name=ok,proto3" json:"ok,omitempty"` //确定操作
}

func (x *Purchase) Reset() {
	*x = Purchase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchaseService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Purchase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Purchase) ProtoMessage() {}

func (x *Purchase) ProtoReflect() protoreflect.Message {
	mi := &file_purchaseService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Purchase.ProtoReflect.Descriptor instead.
func (*Purchase) Descriptor() ([]byte, []int) {
	return file_purchaseService_proto_rawDescGZIP(), []int{1}
}

func (x *Purchase) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Purchase) GetPurchaseSn() string {
	if x != nil {
		return x.PurchaseSn
	}
	return ""
}

func (x *Purchase) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Purchase) GetSourceNo() string {
	if x != nil {
		return x.SourceNo
	}
	return ""
}

func (x *Purchase) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Purchase) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Purchase) GetRowTotal() int32 {
	if x != nil {
		return x.RowTotal
	}
	return 0
}

func (x *Purchase) GetQuantityTotal() int32 {
	if x != nil {
		return x.QuantityTotal
	}
	return 0
}

func (x *Purchase) GetHandledAt() string {
	if x != nil {
		return x.HandledAt
	}
	return ""
}

func (x *Purchase) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Purchase) GetFailure() string {
	if x != nil {
		return x.Failure
	}
	return ""
}

func (x *Purchase) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Purchase) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Purchase) GetDetails() []*PurchaseDetail {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Purchase) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type PurchaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Purchase     `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*Purchase   `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *PurchaseResponse) Reset() {
	*x = PurchaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchaseService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseResponse) ProtoMessage() {}

func (x *PurchaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_purchaseService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseResponse.ProtoReflect.Descriptor instead.
func (*PurchaseResponse) Descriptor() ([]byte, []int) {
	return file_purchaseService_proto_rawDescGZIP(), []int{2}
}

func (x *PurchaseResponse) GetEntity() *Purchase {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PurchaseResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *PurchaseResponse) GetItems() []*Purchase {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PurchaseResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *PurchaseResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_purchaseService_proto protoreflect.FileDescriptor

var file_purchaseService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xfe, 0x01, 0x0a, 0x0d, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x57, 0x68,
	0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x6f, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49,
	0x64, 0x73, 0x22, 0xb0, 0x03, 0x0a, 0x08, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x53, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e,
	0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65,
	0x6d, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x6f, 0x77, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x72, 0x6f, 0x77, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x32, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x07, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0xd4, 0x01, 0x0a, 0x10, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x92, 0x04, 0x0a,
	0x11, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x1a, 0x1a, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1a, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a,
	0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a,
	0x07, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x57, 0x68, 0x65, 0x72,
	0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x37, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x07, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x20, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x32, 0x93, 0x04, 0x0a, 0x12, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x4f, 0x75,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x57, 0x68,
	0x65, 0x72, 0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x40, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x57, 0x68, 0x65, 0x72,
	0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x57, 0x68,
	0x65, 0x72, 0x65, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_purchaseService_proto_rawDescOnce sync.Once
	file_purchaseService_proto_rawDescData = file_purchaseService_proto_rawDesc
)

func file_purchaseService_proto_rawDescGZIP() []byte {
	file_purchaseService_proto_rawDescOnce.Do(func() {
		file_purchaseService_proto_rawDescData = protoimpl.X.CompressGZIP(file_purchaseService_proto_rawDescData)
	})
	return file_purchaseService_proto_rawDescData
}

var file_purchaseService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_purchaseService_proto_goTypes = []interface{}{
	(*PurchaseWhere)(nil),          // 0: services.PurchaseWhere
	(*Purchase)(nil),               // 1: services.Purchase
	(*PurchaseResponse)(nil),       // 2: services.PurchaseResponse
	(*PurchaseDetail)(nil),         // 3: services.PurchaseDetail
	(*common.Pager)(nil),           // 4: common.Pager
	(*common.Error)(nil),           // 5: common.Error
	(*common.Info)(nil),            // 6: common.Info
	(*PurchaseDetailResponse)(nil), // 7: services.PurchaseDetailResponse
}
var file_purchaseService_proto_depIdxs = []int32{
	3,  // 0: services.Purchase.details:type_name -> services.PurchaseDetail
	1,  // 1: services.PurchaseResponse.entity:type_name -> services.Purchase
	4,  // 2: services.PurchaseResponse.pager:type_name -> common.Pager
	1,  // 3: services.PurchaseResponse.items:type_name -> services.Purchase
	5,  // 4: services.PurchaseResponse.error:type_name -> common.Error
	6,  // 5: services.PurchaseResponse.info:type_name -> common.Info
	1,  // 6: services.PurchaseInService.Create:input_type -> services.Purchase
	1,  // 7: services.PurchaseInService.Update:input_type -> services.Purchase
	0,  // 8: services.PurchaseInService.Delete:input_type -> services.PurchaseWhere
	0,  // 9: services.PurchaseInService.Confirm:input_type -> services.PurchaseWhere
	0,  // 10: services.PurchaseInService.Approve:input_type -> services.PurchaseWhere
	1,  // 11: services.PurchaseInService.Get:input_type -> services.Purchase
	0,  // 12: services.PurchaseInService.Search:input_type -> services.PurchaseWhere
	0,  // 13: services.PurchaseInService.Details:input_type -> services.PurchaseWhere
	1,  // 14: services.PurchaseOutService.Create:input_type -> services.Purchase
	1,  // 15: services.PurchaseOutService.Update:input_type -> services.Purchase
	0,  // 16: services.PurchaseOutService.Delete:input_type -> services.PurchaseWhere
	0,  // 17: services.PurchaseOutService.Confirm:input_type -> services.PurchaseWhere
	0,  // 18: services.PurchaseOutService.Approve:input_type -> services.PurchaseWhere
	1,  // 19: services.PurchaseOutService.Get:input_type -> services.Purchase
	0,  // 20: services.PurchaseOutService.Search:input_type -> services.PurchaseWhere
	0,  // 21: services.PurchaseOutService.Details:input_type -> services.PurchaseWhere
	2,  // 22: services.PurchaseInService.Create:output_type -> services.PurchaseResponse
	2,  // 23: services.PurchaseInService.Update:output_type -> services.PurchaseResponse
	2,  // 24: services.PurchaseInService.Delete:output_type -> services.PurchaseResponse
	2,  // 25: services.PurchaseInService.Confirm:output_type -> services.PurchaseResponse
	2,  // 26: services.PurchaseInService.Approve:output_type -> services.PurchaseResponse
	2,  // 27: services.PurchaseInService.Get:output_type -> services.PurchaseResponse
	2,  // 28: services.PurchaseInService.Search:output_type -> services.PurchaseResponse
	7,  // 29: services.PurchaseInService.Details:output_type -> services.PurchaseDetailResponse
	2,  // 30: services.PurchaseOutService.Create:output_type -> services.PurchaseResponse
	2,  // 31: services.PurchaseOutService.Update:output_type -> services.PurchaseResponse
	2,  // 32: services.PurchaseOutService.Delete:output_type -> services.PurchaseResponse
	2,  // 33: services.PurchaseOutService.Confirm:output_type -> services.PurchaseResponse
	2,  // 34: services.PurchaseOutService.Approve:output_type -> services.PurchaseResponse
	2,  // 35: services.PurchaseOutService.Get:output_type -> services.PurchaseResponse
	2,  // 36: services.PurchaseOutService.Search:output_type -> services.PurchaseResponse
	7,  // 37: services.PurchaseOutService.Details:output_type -> services.PurchaseDetailResponse
	22, // [22:38] is the sub-list for method output_type
	6,  // [6:22] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_purchaseService_proto_init() }
func file_purchaseService_proto_init() {
	if File_purchaseService_proto != nil {
		return
	}
	file_purchaseDetailService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_purchaseService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseWhere); i {
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
		file_purchaseService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Purchase); i {
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
		file_purchaseService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseResponse); i {
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
			RawDescriptor: file_purchaseService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_purchaseService_proto_goTypes,
		DependencyIndexes: file_purchaseService_proto_depIdxs,
		MessageInfos:      file_purchaseService_proto_msgTypes,
	}.Build()
	File_purchaseService_proto = out.File
	file_purchaseService_proto_rawDesc = nil
	file_purchaseService_proto_goTypes = nil
	file_purchaseService_proto_depIdxs = nil
}
