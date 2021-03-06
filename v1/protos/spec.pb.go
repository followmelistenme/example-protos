// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: v1/spec.proto

package protos

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

//kafka event
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash      string                 `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	EventTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=EventTime,proto3" json:"EventTime,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_v1_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_v1_spec_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Event) GetEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

type Event1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash1      string                 `protobuf:"bytes,1,opt,name=Hash1,proto3" json:"Hash1,omitempty"`
	EventTime1 *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=EventTime1,proto3" json:"EventTime1,omitempty"`
}

func (x *Event1) Reset() {
	*x = Event1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_spec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event1) ProtoMessage() {}

func (x *Event1) ProtoReflect() protoreflect.Message {
	mi := &file_v1_spec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event1.ProtoReflect.Descriptor instead.
func (*Event1) Descriptor() ([]byte, []int) {
	return file_v1_spec_proto_rawDescGZIP(), []int{1}
}

func (x *Event1) GetHash1() string {
	if x != nil {
		return x.Hash1
	}
	return ""
}

func (x *Event1) GetEventTime1() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTime1
	}
	return nil
}

var File_v1_spec_proto protoreflect.FileDescriptor

var file_v1_spec_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x55, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x38, 0x0a,
	0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x5a, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x31, 0x12, 0x14, 0x0a, 0x05, 0x48, 0x61, 0x73, 0x68, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x48, 0x61, 0x73, 0x68, 0x31, 0x12, 0x3a, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_spec_proto_rawDescOnce sync.Once
	file_v1_spec_proto_rawDescData = file_v1_spec_proto_rawDesc
)

func file_v1_spec_proto_rawDescGZIP() []byte {
	file_v1_spec_proto_rawDescOnce.Do(func() {
		file_v1_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_spec_proto_rawDescData)
	})
	return file_v1_spec_proto_rawDescData
}

var file_v1_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_spec_proto_goTypes = []interface{}{
	(*Event)(nil),                 // 0: Event
	(*Event1)(nil),                // 1: Event1
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_v1_spec_proto_depIdxs = []int32{
	2, // 0: Event.EventTime:type_name -> google.protobuf.Timestamp
	2, // 1: Event1.EventTime1:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_spec_proto_init() }
func file_v1_spec_proto_init() {
	if File_v1_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_spec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_v1_spec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event1); i {
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
			RawDescriptor: file_v1_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_spec_proto_goTypes,
		DependencyIndexes: file_v1_spec_proto_depIdxs,
		MessageInfos:      file_v1_spec_proto_msgTypes,
	}.Build()
	File_v1_spec_proto = out.File
	file_v1_spec_proto_rawDesc = nil
	file_v1_spec_proto_goTypes = nil
	file_v1_spec_proto_depIdxs = nil
}
