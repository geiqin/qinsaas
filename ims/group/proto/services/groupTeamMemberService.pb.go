// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: groupTeamMemberService.proto

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

type GroupTeamMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	GroupTeamId         int64  `protobuf:"varint,2,opt,name=group_team_id,json=groupTeamId,proto3" json:"group_team_id"`
	OrderId             int64  `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	IsPaid              bool   `protobuf:"varint,4,opt,name=is_paid,json=isPaid,proto3" json:"is_paid"`
	CustomerId          int64  `protobuf:"varint,5,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	CustomerDisplayName string `protobuf:"bytes,6,opt,name=customer_display_name,json=customerDisplayName,proto3" json:"customer_display_name"`
	CustomerAvatarUrl   string `protobuf:"bytes,7,opt,name=customer_avatar_url,json=customerAvatarUrl,proto3" json:"customer_avatar_url"`
	IsLeader            bool   `protobuf:"varint,8,opt,name=is_leader,json=isLeader,proto3" json:"is_leader"`
	IsSimulate          bool   `protobuf:"varint,9,opt,name=is_simulate,json=isSimulate,proto3" json:"is_simulate"`
	CreatedAt           string `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt           string `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *GroupTeamMember) Reset() {
	*x = GroupTeamMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupTeamMemberService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupTeamMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupTeamMember) ProtoMessage() {}

func (x *GroupTeamMember) ProtoReflect() protoreflect.Message {
	mi := &file_groupTeamMemberService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupTeamMember.ProtoReflect.Descriptor instead.
func (*GroupTeamMember) Descriptor() ([]byte, []int) {
	return file_groupTeamMemberService_proto_rawDescGZIP(), []int{0}
}

func (x *GroupTeamMember) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GroupTeamMember) GetGroupTeamId() int64 {
	if x != nil {
		return x.GroupTeamId
	}
	return 0
}

func (x *GroupTeamMember) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *GroupTeamMember) GetIsPaid() bool {
	if x != nil {
		return x.IsPaid
	}
	return false
}

func (x *GroupTeamMember) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *GroupTeamMember) GetCustomerDisplayName() string {
	if x != nil {
		return x.CustomerDisplayName
	}
	return ""
}

func (x *GroupTeamMember) GetCustomerAvatarUrl() string {
	if x != nil {
		return x.CustomerAvatarUrl
	}
	return ""
}

func (x *GroupTeamMember) GetIsLeader() bool {
	if x != nil {
		return x.IsLeader
	}
	return false
}

func (x *GroupTeamMember) GetIsSimulate() bool {
	if x != nil {
		return x.IsSimulate
	}
	return false
}

func (x *GroupTeamMember) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GroupTeamMember) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_groupTeamMemberService_proto protoreflect.FileDescriptor

var file_groupTeamMemberService_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0xfa, 0x02, 0x0a, 0x0f, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0d,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x73, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73,
	0x50, 0x61, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x73, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x53,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_groupTeamMemberService_proto_rawDescOnce sync.Once
	file_groupTeamMemberService_proto_rawDescData = file_groupTeamMemberService_proto_rawDesc
)

func file_groupTeamMemberService_proto_rawDescGZIP() []byte {
	file_groupTeamMemberService_proto_rawDescOnce.Do(func() {
		file_groupTeamMemberService_proto_rawDescData = protoimpl.X.CompressGZIP(file_groupTeamMemberService_proto_rawDescData)
	})
	return file_groupTeamMemberService_proto_rawDescData
}

var file_groupTeamMemberService_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_groupTeamMemberService_proto_goTypes = []interface{}{
	(*GroupTeamMember)(nil), // 0: services.GroupTeamMember
}
var file_groupTeamMemberService_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_groupTeamMemberService_proto_init() }
func file_groupTeamMemberService_proto_init() {
	if File_groupTeamMemberService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_groupTeamMemberService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupTeamMember); i {
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
			RawDescriptor: file_groupTeamMemberService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_groupTeamMemberService_proto_goTypes,
		DependencyIndexes: file_groupTeamMemberService_proto_depIdxs,
		MessageInfos:      file_groupTeamMemberService_proto_msgTypes,
	}.Build()
	File_groupTeamMemberService_proto = out.File
	file_groupTeamMemberService_proto_rawDesc = nil
	file_groupTeamMemberService_proto_goTypes = nil
	file_groupTeamMemberService_proto_depIdxs = nil
}
