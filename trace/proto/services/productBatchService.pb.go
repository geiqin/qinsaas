// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: productBatchService.proto

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

type ProductBatchWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id       int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	// @inject_tag: gorm:"-"
	Ids        []int64 `protobuf:"varint,4,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
	Keywords   string  `protobuf:"bytes,5,opt,name=keywords,proto3" json:"keywords"`
	ProductId  int64   `protobuf:"varint,6,opt,name=product_id,json=productId,proto3" json:"product_id"`
	WorkflowId int64   `protobuf:"varint,7,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id"`
	Status     int32   `protobuf:"varint,8,opt,name=status,proto3" json:"status"`
}

func (x *ProductBatchWhere) Reset() {
	*x = ProductBatchWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_productBatchService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductBatchWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductBatchWhere) ProtoMessage() {}

func (x *ProductBatchWhere) ProtoReflect() protoreflect.Message {
	mi := &file_productBatchService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductBatchWhere.ProtoReflect.Descriptor instead.
func (*ProductBatchWhere) Descriptor() ([]byte, []int) {
	return file_productBatchService_proto_rawDescGZIP(), []int{0}
}

func (x *ProductBatchWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *ProductBatchWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ProductBatchWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductBatchWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ProductBatchWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *ProductBatchWhere) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *ProductBatchWhere) GetWorkflowId() int64 {
	if x != nil {
		return x.WorkflowId
	}
	return 0
}

func (x *ProductBatchWhere) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ProductBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	ProductId     int64     `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id"`
	WorkflowId    int64     `protobuf:"varint,3,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id"`
	BatchNo       string    `protobuf:"bytes,4,opt,name=batch_no,json=batchNo,proto3" json:"batch_no"`
	ProduceDate   string    `protobuf:"bytes,5,opt,name=produce_date,json=produceDate,proto3" json:"produce_date"`
	ExpireDate    string    `protobuf:"bytes,6,opt,name=expire_date,json=expireDate,proto3" json:"expire_date"`
	ShelfLifeUnit int32     `protobuf:"varint,7,opt,name=shelf_life_unit,json=shelfLifeUnit,proto3" json:"shelf_life_unit"`
	ShelfLife     int32     `protobuf:"varint,8,opt,name=shelf_life,json=shelfLife,proto3" json:"shelf_life"`
	Content       string    `protobuf:"bytes,9,opt,name=content,proto3" json:"content"`
	Memo          string    `protobuf:"bytes,10,opt,name=memo,proto3" json:"memo"`
	Status        int32     `protobuf:"varint,11,opt,name=status,proto3" json:"status"`
	CreatedAt     string    `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     string    `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Product       *Product  `protobuf:"bytes,14,opt,name=product,proto3" json:"product"`
	Workflow      *Workflow `protobuf:"bytes,15,opt,name=workflow,proto3" json:"workflow"`
}

func (x *ProductBatch) Reset() {
	*x = ProductBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_productBatchService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductBatch) ProtoMessage() {}

func (x *ProductBatch) ProtoReflect() protoreflect.Message {
	mi := &file_productBatchService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductBatch.ProtoReflect.Descriptor instead.
func (*ProductBatch) Descriptor() ([]byte, []int) {
	return file_productBatchService_proto_rawDescGZIP(), []int{1}
}

func (x *ProductBatch) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductBatch) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *ProductBatch) GetWorkflowId() int64 {
	if x != nil {
		return x.WorkflowId
	}
	return 0
}

func (x *ProductBatch) GetBatchNo() string {
	if x != nil {
		return x.BatchNo
	}
	return ""
}

func (x *ProductBatch) GetProduceDate() string {
	if x != nil {
		return x.ProduceDate
	}
	return ""
}

func (x *ProductBatch) GetExpireDate() string {
	if x != nil {
		return x.ExpireDate
	}
	return ""
}

func (x *ProductBatch) GetShelfLifeUnit() int32 {
	if x != nil {
		return x.ShelfLifeUnit
	}
	return 0
}

func (x *ProductBatch) GetShelfLife() int32 {
	if x != nil {
		return x.ShelfLife
	}
	return 0
}

func (x *ProductBatch) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ProductBatch) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *ProductBatch) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ProductBatch) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ProductBatch) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ProductBatch) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *ProductBatch) GetWorkflow() *Workflow {
	if x != nil {
		return x.Workflow
	}
	return nil
}

type ProductBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *common.Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error"`
	Info   *common.Info    `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	Pager  *common.Pager   `protobuf:"bytes,3,opt,name=pager,proto3" json:"pager"`
	Entity *ProductBatch   `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity"`
	Items  []*ProductBatch `protobuf:"bytes,5,rep,name=items,proto3" json:"items"`
}

func (x *ProductBatchResponse) Reset() {
	*x = ProductBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_productBatchService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductBatchResponse) ProtoMessage() {}

func (x *ProductBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_productBatchService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductBatchResponse.ProtoReflect.Descriptor instead.
func (*ProductBatchResponse) Descriptor() ([]byte, []int) {
	return file_productBatchService_proto_rawDescGZIP(), []int{2}
}

func (x *ProductBatchResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ProductBatchResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ProductBatchResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ProductBatchResponse) GetEntity() *ProductBatch {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *ProductBatchResponse) GetItems() []*ProductBatch {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_productBatchService_proto protoreflect.FileDescriptor

var file_productBatchService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0xe5, 0x03, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6e,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x5f, 0x6c, 0x69,
	0x66, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73,
	0x68, 0x65, 0x6c, 0x66, 0x4c, 0x69, 0x66, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x68, 0x65, 0x6c, 0x66, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x4c, 0x69, 0x66, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x2b, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x2e, 0x0a, 0x08,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0xe0, 0x01, 0x0a,
	0x14, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x12, 0x2e, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32,
	0xfa, 0x03, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x1e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x57, 0x68, 0x65,
	0x72, 0x65, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a,
	0x09, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x5e, 0x0a, 0x18,
	0x46, 0x72, 0x6f, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_productBatchService_proto_rawDescOnce sync.Once
	file_productBatchService_proto_rawDescData = file_productBatchService_proto_rawDesc
)

func file_productBatchService_proto_rawDescGZIP() []byte {
	file_productBatchService_proto_rawDescOnce.Do(func() {
		file_productBatchService_proto_rawDescData = protoimpl.X.CompressGZIP(file_productBatchService_proto_rawDescData)
	})
	return file_productBatchService_proto_rawDescData
}

var file_productBatchService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_productBatchService_proto_goTypes = []interface{}{
	(*ProductBatchWhere)(nil),    // 0: services.ProductBatchWhere
	(*ProductBatch)(nil),         // 1: services.ProductBatch
	(*ProductBatchResponse)(nil), // 2: services.ProductBatchResponse
	(*Product)(nil),              // 3: services.Product
	(*Workflow)(nil),             // 4: services.Workflow
	(*common.Error)(nil),         // 5: common.Error
	(*common.Info)(nil),          // 6: common.Info
	(*common.Pager)(nil),         // 7: common.Pager
}
var file_productBatchService_proto_depIdxs = []int32{
	3,  // 0: services.ProductBatch.product:type_name -> services.Product
	4,  // 1: services.ProductBatch.workflow:type_name -> services.Workflow
	5,  // 2: services.ProductBatchResponse.error:type_name -> common.Error
	6,  // 3: services.ProductBatchResponse.info:type_name -> common.Info
	7,  // 4: services.ProductBatchResponse.pager:type_name -> common.Pager
	1,  // 5: services.ProductBatchResponse.entity:type_name -> services.ProductBatch
	1,  // 6: services.ProductBatchResponse.items:type_name -> services.ProductBatch
	1,  // 7: services.ProductBatchService.Create:input_type -> services.ProductBatch
	1,  // 8: services.ProductBatchService.Update:input_type -> services.ProductBatch
	0,  // 9: services.ProductBatchService.Search:input_type -> services.ProductBatchWhere
	0,  // 10: services.ProductBatchService.Delete:input_type -> services.ProductBatchWhere
	0,  // 11: services.ProductBatchService.Get:input_type -> services.ProductBatchWhere
	0,  // 12: services.ProductBatchService.SetStatus:input_type -> services.ProductBatchWhere
	0,  // 13: services.ProductBatchService.List:input_type -> services.ProductBatchWhere
	0,  // 14: services.FrontProductBatchService.Get:input_type -> services.ProductBatchWhere
	2,  // 15: services.ProductBatchService.Create:output_type -> services.ProductBatchResponse
	2,  // 16: services.ProductBatchService.Update:output_type -> services.ProductBatchResponse
	2,  // 17: services.ProductBatchService.Search:output_type -> services.ProductBatchResponse
	2,  // 18: services.ProductBatchService.Delete:output_type -> services.ProductBatchResponse
	2,  // 19: services.ProductBatchService.Get:output_type -> services.ProductBatchResponse
	2,  // 20: services.ProductBatchService.SetStatus:output_type -> services.ProductBatchResponse
	2,  // 21: services.ProductBatchService.List:output_type -> services.ProductBatchResponse
	2,  // 22: services.FrontProductBatchService.Get:output_type -> services.ProductBatchResponse
	15, // [15:23] is the sub-list for method output_type
	7,  // [7:15] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_productBatchService_proto_init() }
func file_productBatchService_proto_init() {
	if File_productBatchService_proto != nil {
		return
	}
	file_productService_proto_init()
	file_workflowService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_productBatchService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductBatchWhere); i {
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
		file_productBatchService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductBatch); i {
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
		file_productBatchService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductBatchResponse); i {
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
			RawDescriptor: file_productBatchService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_productBatchService_proto_goTypes,
		DependencyIndexes: file_productBatchService_proto_depIdxs,
		MessageInfos:      file_productBatchService_proto_msgTypes,
	}.Build()
	File_productBatchService_proto = out.File
	file_productBatchService_proto_rawDesc = nil
	file_productBatchService_proto_goTypes = nil
	file_productBatchService_proto_depIdxs = nil
}
