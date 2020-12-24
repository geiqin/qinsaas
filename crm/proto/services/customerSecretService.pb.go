// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: customerSecretService.proto

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

type CustomerSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CustomerId int64  `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	Password   string `protobuf:"bytes,3,opt,name=password,proto3" json:"password"`
	Locked     bool   `protobuf:"varint,4,opt,name=locked,proto3" json:"locked"`
}

func (x *CustomerSecret) Reset() {
	*x = CustomerSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customerSecretService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerSecret) ProtoMessage() {}

func (x *CustomerSecret) ProtoReflect() protoreflect.Message {
	mi := &file_customerSecretService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerSecret.ProtoReflect.Descriptor instead.
func (*CustomerSecret) Descriptor() ([]byte, []int) {
	return file_customerSecretService_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerSecret) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CustomerSecret) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CustomerSecret) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CustomerSecret) GetLocked() bool {
	if x != nil {
		return x.Locked
	}
	return false
}

//
type CustomerSecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *CustomerSecret   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*CustomerSecret `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error     `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info      `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *CustomerSecretResponse) Reset() {
	*x = CustomerSecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customerSecretService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerSecretResponse) ProtoMessage() {}

func (x *CustomerSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customerSecretService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerSecretResponse.ProtoReflect.Descriptor instead.
func (*CustomerSecretResponse) Descriptor() ([]byte, []int) {
	return file_customerSecretService_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerSecretResponse) GetEntity() *CustomerSecret {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CustomerSecretResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CustomerSecretResponse) GetItems() []*CustomerSecret {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CustomerSecretResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CustomerSecretResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_customerSecretService_proto protoreflect.FileDescriptor

var file_customerSecretService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x0e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x22, 0xe6, 0x01, 0x0a, 0x16, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customerSecretService_proto_rawDescOnce sync.Once
	file_customerSecretService_proto_rawDescData = file_customerSecretService_proto_rawDesc
)

func file_customerSecretService_proto_rawDescGZIP() []byte {
	file_customerSecretService_proto_rawDescOnce.Do(func() {
		file_customerSecretService_proto_rawDescData = protoimpl.X.CompressGZIP(file_customerSecretService_proto_rawDescData)
	})
	return file_customerSecretService_proto_rawDescData
}

var file_customerSecretService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customerSecretService_proto_goTypes = []interface{}{
	(*CustomerSecret)(nil),         // 0: services.CustomerSecret
	(*CustomerSecretResponse)(nil), // 1: services.CustomerSecretResponse
	(*common.Pager)(nil),           // 2: common.Pager
	(*common.Error)(nil),           // 3: common.Error
	(*common.Info)(nil),            // 4: common.Info
}
var file_customerSecretService_proto_depIdxs = []int32{
	0, // 0: services.CustomerSecretResponse.entity:type_name -> services.CustomerSecret
	2, // 1: services.CustomerSecretResponse.pager:type_name -> common.Pager
	0, // 2: services.CustomerSecretResponse.items:type_name -> services.CustomerSecret
	3, // 3: services.CustomerSecretResponse.error:type_name -> common.Error
	4, // 4: services.CustomerSecretResponse.info:type_name -> common.Info
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_customerSecretService_proto_init() }
func file_customerSecretService_proto_init() {
	if File_customerSecretService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customerSecretService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerSecret); i {
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
		file_customerSecretService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerSecretResponse); i {
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
			RawDescriptor: file_customerSecretService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customerSecretService_proto_goTypes,
		DependencyIndexes: file_customerSecretService_proto_depIdxs,
		MessageInfos:      file_customerSecretService_proto_msgTypes,
	}.Build()
	File_customerSecretService_proto = out.File
	file_customerSecretService_proto_rawDesc = nil
	file_customerSecretService_proto_goTypes = nil
	file_customerSecretService_proto_depIdxs = nil
}
