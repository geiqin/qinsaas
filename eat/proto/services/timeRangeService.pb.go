// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: timeRangeService.proto

package services

import (
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

type TimeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TimeConfigId int64  `protobuf:"varint,2,opt,name=time_config_id,json=timeConfigId,proto3" json:"time_config_id,omitempty"`
	StartTime    string `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime      string `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Weeks        string `protobuf:"bytes,5,opt,name=weeks,proto3" json:"weeks,omitempty"`
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timeRangeService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_timeRangeService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_timeRangeService_proto_rawDescGZIP(), []int{0}
}

func (x *TimeRange) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TimeRange) GetTimeConfigId() int64 {
	if x != nil {
		return x.TimeConfigId
	}
	return 0
}

func (x *TimeRange) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *TimeRange) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *TimeRange) GetWeeks() string {
	if x != nil {
		return x.Weeks
	}
	return ""
}

var File_timeRangeService_proto protoreflect.FileDescriptor

var file_timeRangeService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x22, 0x91, 0x01, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x24, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_timeRangeService_proto_rawDescOnce sync.Once
	file_timeRangeService_proto_rawDescData = file_timeRangeService_proto_rawDesc
)

func file_timeRangeService_proto_rawDescGZIP() []byte {
	file_timeRangeService_proto_rawDescOnce.Do(func() {
		file_timeRangeService_proto_rawDescData = protoimpl.X.CompressGZIP(file_timeRangeService_proto_rawDescData)
	})
	return file_timeRangeService_proto_rawDescData
}

var file_timeRangeService_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_timeRangeService_proto_goTypes = []interface{}{
	(*TimeRange)(nil), // 0: services.TimeRange
}
var file_timeRangeService_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_timeRangeService_proto_init() }
func file_timeRangeService_proto_init() {
	if File_timeRangeService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_timeRangeService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRange); i {
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
			RawDescriptor: file_timeRangeService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_timeRangeService_proto_goTypes,
		DependencyIndexes: file_timeRangeService_proto_depIdxs,
		MessageInfos:      file_timeRangeService_proto_msgTypes,
	}.Build()
	File_timeRangeService_proto = out.File
	file_timeRangeService_proto_rawDesc = nil
	file_timeRangeService_proto_goTypes = nil
	file_timeRangeService_proto_depIdxs = nil
}
