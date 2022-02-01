// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: activity-log/api/v1/activity.proto

package api_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RetrieveRequest) Reset() {
	*x = RetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_log_api_v1_activity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveRequest) ProtoMessage() {}

func (x *RetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activity_log_api_v1_activity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveRequest.ProtoReflect.Descriptor instead.
func (*RetrieveRequest) Descriptor() ([]byte, []int) {
	return file_activity_log_api_v1_activity_proto_rawDescGZIP(), []int{0}
}

func (x *RetrieveRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type InsertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InsertResponse) Reset() {
	*x = InsertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_log_api_v1_activity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertResponse) ProtoMessage() {}

func (x *InsertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_activity_log_api_v1_activity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertResponse.ProtoReflect.Descriptor instead.
func (*InsertResponse) Descriptor() ([]byte, []int) {
	return file_activity_log_api_v1_activity_proto_rawDescGZIP(), []int{1}
}

func (x *InsertResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_log_api_v1_activity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activity_log_api_v1_activity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_activity_log_api_v1_activity_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Time        *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_log_api_v1_activity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_activity_log_api_v1_activity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_activity_log_api_v1_activity_proto_rawDescGZIP(), []int{3}
}

func (x *Activity) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Activity) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Activity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Activities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activities []*Activity `protobuf:"bytes,1,rep,name=activities,proto3" json:"activities,omitempty"`
}

func (x *Activities) Reset() {
	*x = Activities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_log_api_v1_activity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activities) ProtoMessage() {}

func (x *Activities) ProtoReflect() protoreflect.Message {
	mi := &file_activity_log_api_v1_activity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activities.ProtoReflect.Descriptor instead.
func (*Activities) Descriptor() ([]byte, []int) {
	return file_activity_log_api_v1_activity_proto_rawDescGZIP(), []int{4}
}

func (x *Activities) GetActivities() []*Activity {
	if x != nil {
		return x.Activities
	}
	return nil
}

type ActivityQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ActivityQuery) Reset() {
	*x = ActivityQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_log_api_v1_activity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityQuery) ProtoMessage() {}

func (x *ActivityQuery) ProtoReflect() protoreflect.Message {
	mi := &file_activity_log_api_v1_activity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityQuery.ProtoReflect.Descriptor instead.
func (*ActivityQuery) Descriptor() ([]byte, []int) {
	return file_activity_log_api_v1_activity_proto_rawDescGZIP(), []int{5}
}

func (x *ActivityQuery) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_activity_log_api_v1_activity_proto protoreflect.FileDescriptor

var file_activity_log_api_v1_activity_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2d, 0x6c, 0x6f, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a,
	0x0f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x20, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x25, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x6c, 0x0a, 0x08, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x27, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x32, 0xb0, 0x01, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x4c, 0x6f,
	0x67, 0x12, 0x34, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x1a, 0x16, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x31, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x64, 0x61, 0x6d, 0x67, 0x6f, 0x72, 0x64, 0x6f, 0x6e, 0x62, 0x65, 0x6c, 0x6c,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2d, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x5f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_activity_log_api_v1_activity_proto_rawDescOnce sync.Once
	file_activity_log_api_v1_activity_proto_rawDescData = file_activity_log_api_v1_activity_proto_rawDesc
)

func file_activity_log_api_v1_activity_proto_rawDescGZIP() []byte {
	file_activity_log_api_v1_activity_proto_rawDescOnce.Do(func() {
		file_activity_log_api_v1_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_activity_log_api_v1_activity_proto_rawDescData)
	})
	return file_activity_log_api_v1_activity_proto_rawDescData
}

var file_activity_log_api_v1_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_activity_log_api_v1_activity_proto_goTypes = []interface{}{
	(*RetrieveRequest)(nil),       // 0: api.v1.RetrieveRequest
	(*InsertResponse)(nil),        // 1: api.v1.InsertResponse
	(*ListRequest)(nil),           // 2: api.v1.ListRequest
	(*Activity)(nil),              // 3: api.v1.Activity
	(*Activities)(nil),            // 4: api.v1.Activities
	(*ActivityQuery)(nil),         // 5: api.v1.ActivityQuery
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_activity_log_api_v1_activity_proto_depIdxs = []int32{
	6, // 0: api.v1.Activity.time:type_name -> google.protobuf.Timestamp
	3, // 1: api.v1.Activities.activities:type_name -> api.v1.Activity
	3, // 2: api.v1.Activity_Log.Insert:input_type -> api.v1.Activity
	0, // 3: api.v1.Activity_Log.Retrieve:input_type -> api.v1.RetrieveRequest
	2, // 4: api.v1.Activity_Log.List:input_type -> api.v1.ListRequest
	1, // 5: api.v1.Activity_Log.Insert:output_type -> api.v1.InsertResponse
	3, // 6: api.v1.Activity_Log.Retrieve:output_type -> api.v1.Activity
	4, // 7: api.v1.Activity_Log.List:output_type -> api.v1.Activities
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_activity_log_api_v1_activity_proto_init() }
func file_activity_log_api_v1_activity_proto_init() {
	if File_activity_log_api_v1_activity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_activity_log_api_v1_activity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveRequest); i {
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
		file_activity_log_api_v1_activity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertResponse); i {
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
		file_activity_log_api_v1_activity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_activity_log_api_v1_activity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
		file_activity_log_api_v1_activity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activities); i {
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
		file_activity_log_api_v1_activity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityQuery); i {
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
			RawDescriptor: file_activity_log_api_v1_activity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_activity_log_api_v1_activity_proto_goTypes,
		DependencyIndexes: file_activity_log_api_v1_activity_proto_depIdxs,
		MessageInfos:      file_activity_log_api_v1_activity_proto_msgTypes,
	}.Build()
	File_activity_log_api_v1_activity_proto = out.File
	file_activity_log_api_v1_activity_proto_rawDesc = nil
	file_activity_log_api_v1_activity_proto_goTypes = nil
	file_activity_log_api_v1_activity_proto_depIdxs = nil
}
