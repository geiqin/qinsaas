// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: metaFieldService.proto

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

type MetaFieldWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id       int32 `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	// @inject_tag: gorm:"-"
	Ids         []int32 `protobuf:"varint,4,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
	Keywords    string  `protobuf:"bytes,5,opt,name=keywords,proto3" json:"keywords"`
	MetaClassId int32   `protobuf:"varint,6,opt,name=meta_class_id,json=metaClassId,proto3" json:"meta_class_id"`
	// @inject_tag: gorm:"-"
	MetaClassIds []int32 `protobuf:"varint,7,rep,packed,name=meta_class_ids,json=metaClassIds,proto3" json:"meta_class_ids" gorm:"-"`
}

func (x *MetaFieldWhere) Reset() {
	*x = MetaFieldWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metaFieldService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaFieldWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaFieldWhere) ProtoMessage() {}

func (x *MetaFieldWhere) ProtoReflect() protoreflect.Message {
	mi := &file_metaFieldService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaFieldWhere.ProtoReflect.Descriptor instead.
func (*MetaFieldWhere) Descriptor() ([]byte, []int) {
	return file_metaFieldService_proto_rawDescGZIP(), []int{0}
}

func (x *MetaFieldWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *MetaFieldWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *MetaFieldWhere) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MetaFieldWhere) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *MetaFieldWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *MetaFieldWhere) GetMetaClassId() int32 {
	if x != nil {
		return x.MetaClassId
	}
	return 0
}

func (x *MetaFieldWhere) GetMetaClassIds() []int32 {
	if x != nil {
		return x.MetaClassIds
	}
	return nil
}

type MetaField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	MetaClassId int32  `protobuf:"varint,2,opt,name=meta_class_id,json=metaClassId,proto3" json:"meta_class_id"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name"`
	Type        string `protobuf:"bytes,5,opt,name=type,proto3" json:"type"`
	Required    bool   `protobuf:"varint,6,opt,name=required,proto3" json:"required"`
	Rule        string `protobuf:"bytes,7,opt,name=rule,proto3" json:"rule"`
	Tip         string `protobuf:"bytes,8,opt,name=tip,proto3" json:"tip"`
	Sorting     int32  `protobuf:"varint,9,opt,name=sorting,proto3" json:"sorting"`
	Data        string `protobuf:"bytes,10,opt,name=data,proto3" json:"data"`
	ShowInList  bool   `protobuf:"varint,11,opt,name=show_in_list,json=showInList,proto3" json:"show_in_list"`
	IsQuery     bool   `protobuf:"varint,12,opt,name=is_query,json=isQuery,proto3" json:"is_query"`
	QueryMethod string `protobuf:"bytes,13,opt,name=query_method,json=queryMethod,proto3" json:"query_method"`
	CreatedAt   string `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt   string `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *MetaField) Reset() {
	*x = MetaField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metaFieldService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaField) ProtoMessage() {}

func (x *MetaField) ProtoReflect() protoreflect.Message {
	mi := &file_metaFieldService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaField.ProtoReflect.Descriptor instead.
func (*MetaField) Descriptor() ([]byte, []int) {
	return file_metaFieldService_proto_rawDescGZIP(), []int{1}
}

func (x *MetaField) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MetaField) GetMetaClassId() int32 {
	if x != nil {
		return x.MetaClassId
	}
	return 0
}

func (x *MetaField) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetaField) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *MetaField) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MetaField) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *MetaField) GetRule() string {
	if x != nil {
		return x.Rule
	}
	return ""
}

func (x *MetaField) GetTip() string {
	if x != nil {
		return x.Tip
	}
	return ""
}

func (x *MetaField) GetSorting() int32 {
	if x != nil {
		return x.Sorting
	}
	return 0
}

func (x *MetaField) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *MetaField) GetShowInList() bool {
	if x != nil {
		return x.ShowInList
	}
	return false
}

func (x *MetaField) GetIsQuery() bool {
	if x != nil {
		return x.IsQuery
	}
	return false
}

func (x *MetaField) GetQueryMethod() string {
	if x != nil {
		return x.QueryMethod
	}
	return ""
}

func (x *MetaField) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MetaField) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type MetaFieldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *common.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	Pager  *common.Pager `protobuf:"bytes,3,opt,name=pager,proto3" json:"pager"`
	Entity *MetaField    `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity"`
	Items  []*MetaField  `protobuf:"bytes,5,rep,name=items,proto3" json:"items"`
}

func (x *MetaFieldResponse) Reset() {
	*x = MetaFieldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metaFieldService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaFieldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaFieldResponse) ProtoMessage() {}

func (x *MetaFieldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metaFieldService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaFieldResponse.ProtoReflect.Descriptor instead.
func (*MetaFieldResponse) Descriptor() ([]byte, []int) {
	return file_metaFieldService_proto_rawDescGZIP(), []int{2}
}

func (x *MetaFieldResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *MetaFieldResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *MetaFieldResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *MetaFieldResponse) GetEntity() *MetaField {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *MetaFieldResponse) GetItems() []*MetaField {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_metaFieldService_proto protoreflect.FileDescriptor

var file_metaFieldService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x65, 0x74,
	0x61, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x49, 0x64, 0x73, 0x22, 0x98, 0x03, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x75, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x69, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x20, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x69, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x21, 0x0a,
	0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd7,
	0x01, 0x0a, 0x11, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x29, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x89, 0x03, 0x0a, 0x10, 0x4d, 0x65, 0x74,
	0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0x1b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metaFieldService_proto_rawDescOnce sync.Once
	file_metaFieldService_proto_rawDescData = file_metaFieldService_proto_rawDesc
)

func file_metaFieldService_proto_rawDescGZIP() []byte {
	file_metaFieldService_proto_rawDescOnce.Do(func() {
		file_metaFieldService_proto_rawDescData = protoimpl.X.CompressGZIP(file_metaFieldService_proto_rawDescData)
	})
	return file_metaFieldService_proto_rawDescData
}

var file_metaFieldService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_metaFieldService_proto_goTypes = []interface{}{
	(*MetaFieldWhere)(nil),    // 0: services.MetaFieldWhere
	(*MetaField)(nil),         // 1: services.MetaField
	(*MetaFieldResponse)(nil), // 2: services.MetaFieldResponse
	(*common.Error)(nil),      // 3: common.Error
	(*common.Info)(nil),       // 4: common.Info
	(*common.Pager)(nil),      // 5: common.Pager
}
var file_metaFieldService_proto_depIdxs = []int32{
	3,  // 0: services.MetaFieldResponse.error:type_name -> common.Error
	4,  // 1: services.MetaFieldResponse.info:type_name -> common.Info
	5,  // 2: services.MetaFieldResponse.pager:type_name -> common.Pager
	1,  // 3: services.MetaFieldResponse.entity:type_name -> services.MetaField
	1,  // 4: services.MetaFieldResponse.items:type_name -> services.MetaField
	1,  // 5: services.MetaFieldService.Create:input_type -> services.MetaField
	1,  // 6: services.MetaFieldService.Update:input_type -> services.MetaField
	0,  // 7: services.MetaFieldService.Search:input_type -> services.MetaFieldWhere
	0,  // 8: services.MetaFieldService.List:input_type -> services.MetaFieldWhere
	0,  // 9: services.MetaFieldService.Delete:input_type -> services.MetaFieldWhere
	0,  // 10: services.MetaFieldService.Get:input_type -> services.MetaFieldWhere
	2,  // 11: services.MetaFieldService.Create:output_type -> services.MetaFieldResponse
	2,  // 12: services.MetaFieldService.Update:output_type -> services.MetaFieldResponse
	2,  // 13: services.MetaFieldService.Search:output_type -> services.MetaFieldResponse
	2,  // 14: services.MetaFieldService.List:output_type -> services.MetaFieldResponse
	2,  // 15: services.MetaFieldService.Delete:output_type -> services.MetaFieldResponse
	2,  // 16: services.MetaFieldService.Get:output_type -> services.MetaFieldResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_metaFieldService_proto_init() }
func file_metaFieldService_proto_init() {
	if File_metaFieldService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metaFieldService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaFieldWhere); i {
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
		file_metaFieldService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaField); i {
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
		file_metaFieldService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaFieldResponse); i {
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
			RawDescriptor: file_metaFieldService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metaFieldService_proto_goTypes,
		DependencyIndexes: file_metaFieldService_proto_depIdxs,
		MessageInfos:      file_metaFieldService_proto_msgTypes,
	}.Build()
	File_metaFieldService_proto = out.File
	file_metaFieldService_proto_rawDesc = nil
	file_metaFieldService_proto_goTypes = nil
	file_metaFieldService_proto_depIdxs = nil
}
