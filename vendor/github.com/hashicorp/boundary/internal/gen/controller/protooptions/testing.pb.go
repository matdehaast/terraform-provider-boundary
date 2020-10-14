// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: controller/custom_options/v1/testing.proto

package protooptions

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

type TestBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstField                 string `protobuf:"bytes,1,opt,name=first_field,json=firstField,proto3" json:"first_field,omitempty"`
	StrangeFormatField         string `protobuf:"bytes,2,opt,name=StrangeFormatField,proto3" json:"StrangeFormatField,omitempty"`
	FieldWithDifferentJsonName string `protobuf:"bytes,3,opt,name=field_with_different_json_name,json=json_name,proto3" json:"field_with_different_json_name,omitempty"`
	ExtraField                 string `protobuf:"bytes,4,opt,name=extra_field,json=extraField,proto3" json:"extra_field,omitempty"`
}

func (x *TestBase) Reset() {
	*x = TestBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_custom_options_v1_testing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestBase) ProtoMessage() {}

func (x *TestBase) ProtoReflect() protoreflect.Message {
	mi := &file_controller_custom_options_v1_testing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestBase.ProtoReflect.Descriptor instead.
func (*TestBase) Descriptor() ([]byte, []int) {
	return file_controller_custom_options_v1_testing_proto_rawDescGZIP(), []int{0}
}

func (x *TestBase) GetFirstField() string {
	if x != nil {
		return x.FirstField
	}
	return ""
}

func (x *TestBase) GetStrangeFormatField() string {
	if x != nil {
		return x.StrangeFormatField
	}
	return ""
}

func (x *TestBase) GetFieldWithDifferentJsonName() string {
	if x != nil {
		return x.FieldWithDifferentJsonName
	}
	return ""
}

func (x *TestBase) GetExtraField() string {
	if x != nil {
		return x.ExtraField
	}
	return ""
}

type TestBaseSplit1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstField                 string `protobuf:"bytes,1,opt,name=first_field,json=firstField,proto3" json:"first_field,omitempty"`
	FieldWithDifferentJsonName string `protobuf:"bytes,2,opt,name=field_with_different_json_name,json=json_name,proto3" json:"field_with_different_json_name,omitempty"`
	ExtraField                 string `protobuf:"bytes,3,opt,name=extra_field,json=extraField,proto3" json:"extra_field,omitempty"`
}

func (x *TestBaseSplit1) Reset() {
	*x = TestBaseSplit1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_custom_options_v1_testing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestBaseSplit1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestBaseSplit1) ProtoMessage() {}

func (x *TestBaseSplit1) ProtoReflect() protoreflect.Message {
	mi := &file_controller_custom_options_v1_testing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestBaseSplit1.ProtoReflect.Descriptor instead.
func (*TestBaseSplit1) Descriptor() ([]byte, []int) {
	return file_controller_custom_options_v1_testing_proto_rawDescGZIP(), []int{1}
}

func (x *TestBaseSplit1) GetFirstField() string {
	if x != nil {
		return x.FirstField
	}
	return ""
}

func (x *TestBaseSplit1) GetFieldWithDifferentJsonName() string {
	if x != nil {
		return x.FieldWithDifferentJsonName
	}
	return ""
}

func (x *TestBaseSplit1) GetExtraField() string {
	if x != nil {
		return x.ExtraField
	}
	return ""
}

type TestBaseSplit2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SplitField1 string `protobuf:"bytes,1,opt,name=split_field1,json=splitField1,proto3" json:"split_field1,omitempty"`
}

func (x *TestBaseSplit2) Reset() {
	*x = TestBaseSplit2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_custom_options_v1_testing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestBaseSplit2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestBaseSplit2) ProtoMessage() {}

func (x *TestBaseSplit2) ProtoReflect() protoreflect.Message {
	mi := &file_controller_custom_options_v1_testing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestBaseSplit2.ProtoReflect.Descriptor instead.
func (*TestBaseSplit2) Descriptor() ([]byte, []int) {
	return file_controller_custom_options_v1_testing_proto_rawDescGZIP(), []int{2}
}

func (x *TestBaseSplit2) GetSplitField1() string {
	if x != nil {
		return x.SplitField1
	}
	return ""
}

type TestProperlyNamedFields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtherFirstField   string `protobuf:"bytes,1,opt,name=other_first_field,json=otherFirstField,proto3" json:"other_first_field,omitempty"`
	OtherSecondField  string `protobuf:"bytes,2,opt,name=other_second_field,json=otherSecondField,proto3" json:"other_second_field,omitempty"`
	OtherThirdField   string `protobuf:"bytes,3,opt,name=other_third_field,proto3" json:"other_third_field,omitempty"`
	AnotherExtraField string `protobuf:"bytes,4,opt,name=another_extra_field,json=anotherExtraField,proto3" json:"another_extra_field,omitempty"`
}

func (x *TestProperlyNamedFields) Reset() {
	*x = TestProperlyNamedFields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_custom_options_v1_testing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestProperlyNamedFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestProperlyNamedFields) ProtoMessage() {}

func (x *TestProperlyNamedFields) ProtoReflect() protoreflect.Message {
	mi := &file_controller_custom_options_v1_testing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestProperlyNamedFields.ProtoReflect.Descriptor instead.
func (*TestProperlyNamedFields) Descriptor() ([]byte, []int) {
	return file_controller_custom_options_v1_testing_proto_rawDescGZIP(), []int{3}
}

func (x *TestProperlyNamedFields) GetOtherFirstField() string {
	if x != nil {
		return x.OtherFirstField
	}
	return ""
}

func (x *TestProperlyNamedFields) GetOtherSecondField() string {
	if x != nil {
		return x.OtherSecondField
	}
	return ""
}

func (x *TestProperlyNamedFields) GetOtherThirdField() string {
	if x != nil {
		return x.OtherThirdField
	}
	return ""
}

func (x *TestProperlyNamedFields) GetAnotherExtraField() string {
	if x != nil {
		return x.AnotherExtraField
	}
	return ""
}

type TestNameDoesntMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstDoesntMap    string `protobuf:"bytes,1,opt,name=first_doesnt_map,json=firstDoesntMap,proto3" json:"first_doesnt_map,omitempty"`
	SecondDoesntMap   string `protobuf:"bytes,2,opt,name=second_doesnt_map,json=secondDoesntMap,proto3" json:"second_doesnt_map,omitempty"`
	ThirdDoesntMap    string `protobuf:"bytes,3,opt,name=third_doesnt_map,json=thirdDoesntMap,proto3" json:"third_doesnt_map,omitempty"`
	AnotherExtraField string `protobuf:"bytes,4,opt,name=another_extra_field,json=anotherExtraField,proto3" json:"another_extra_field,omitempty"`
}

func (x *TestNameDoesntMap) Reset() {
	*x = TestNameDoesntMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_custom_options_v1_testing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestNameDoesntMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestNameDoesntMap) ProtoMessage() {}

func (x *TestNameDoesntMap) ProtoReflect() protoreflect.Message {
	mi := &file_controller_custom_options_v1_testing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestNameDoesntMap.ProtoReflect.Descriptor instead.
func (*TestNameDoesntMap) Descriptor() ([]byte, []int) {
	return file_controller_custom_options_v1_testing_proto_rawDescGZIP(), []int{4}
}

func (x *TestNameDoesntMap) GetFirstDoesntMap() string {
	if x != nil {
		return x.FirstDoesntMap
	}
	return ""
}

func (x *TestNameDoesntMap) GetSecondDoesntMap() string {
	if x != nil {
		return x.SecondDoesntMap
	}
	return ""
}

func (x *TestNameDoesntMap) GetThirdDoesntMap() string {
	if x != nil {
		return x.ThirdDoesntMap
	}
	return ""
}

func (x *TestNameDoesntMap) GetAnotherExtraField() string {
	if x != nil {
		return x.AnotherExtraField
	}
	return ""
}

type TestNotEnoughFields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtherFirstField  string `protobuf:"bytes,1,opt,name=other_first_field,json=otherFirstField,proto3" json:"other_first_field,omitempty"`
	OtherSecondField string `protobuf:"bytes,2,opt,name=other_second_field,json=otherSecondField,proto3" json:"other_second_field,omitempty"`
}

func (x *TestNotEnoughFields) Reset() {
	*x = TestNotEnoughFields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_custom_options_v1_testing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestNotEnoughFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestNotEnoughFields) ProtoMessage() {}

func (x *TestNotEnoughFields) ProtoReflect() protoreflect.Message {
	mi := &file_controller_custom_options_v1_testing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestNotEnoughFields.ProtoReflect.Descriptor instead.
func (*TestNotEnoughFields) Descriptor() ([]byte, []int) {
	return file_controller_custom_options_v1_testing_proto_rawDescGZIP(), []int{5}
}

func (x *TestNotEnoughFields) GetOtherFirstField() string {
	if x != nil {
		return x.OtherFirstField
	}
	return ""
}

func (x *TestNotEnoughFields) GetOtherSecondField() string {
	if x != nil {
		return x.OtherSecondField
	}
	return ""
}

type TestManyToOneMappings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtherFirstField  string `protobuf:"bytes,1,opt,name=other_first_field,json=otherFirstField,proto3" json:"other_first_field,omitempty"`
	OtherSecondField string `protobuf:"bytes,2,opt,name=other_second_field,json=otherSecondField,proto3" json:"other_second_field,omitempty"`
	OtherThirdField  string `protobuf:"bytes,3,opt,name=other_third_field,json=otherThirdField,proto3" json:"other_third_field,omitempty"`
	// this field maps to the first_field like 'other_first_field'
	ExtraField string `protobuf:"bytes,4,opt,name=extra_field,json=extraField,proto3" json:"extra_field,omitempty"`
}

func (x *TestManyToOneMappings) Reset() {
	*x = TestManyToOneMappings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_custom_options_v1_testing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestManyToOneMappings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestManyToOneMappings) ProtoMessage() {}

func (x *TestManyToOneMappings) ProtoReflect() protoreflect.Message {
	mi := &file_controller_custom_options_v1_testing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestManyToOneMappings.ProtoReflect.Descriptor instead.
func (*TestManyToOneMappings) Descriptor() ([]byte, []int) {
	return file_controller_custom_options_v1_testing_proto_rawDescGZIP(), []int{6}
}

func (x *TestManyToOneMappings) GetOtherFirstField() string {
	if x != nil {
		return x.OtherFirstField
	}
	return ""
}

func (x *TestManyToOneMappings) GetOtherSecondField() string {
	if x != nil {
		return x.OtherSecondField
	}
	return ""
}

func (x *TestManyToOneMappings) GetOtherThirdField() string {
	if x != nil {
		return x.OtherThirdField
	}
	return ""
}

func (x *TestManyToOneMappings) GetExtraField() string {
	if x != nil {
		return x.ExtraField
	}
	return ""
}

var File_controller_custom_options_v1_testing_proto protoreflect.FileDescriptor

var file_controller_custom_options_v1_testing_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x42,
	0x61, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xc2, 0xdd, 0x29, 0x1e, 0x0a, 0x0b,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0f, 0x4f, 0x74, 0x68,
	0x65, 0x72, 0x46, 0x69, 0x72, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x62, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x32, 0xc2, 0xdd, 0x29, 0x2e, 0x0a, 0x18, 0x73, 0x74, 0x72, 0x61,
	0x6e, 0x67, 0x6c, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x12, 0x53, 0x74, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x4e, 0x0a, 0x1e,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xc2, 0xdd, 0x29, 0x17, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x33, 0x12, 0x0d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x33, 0x52, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0xc6, 0x01,
	0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x42, 0x61, 0x73, 0x65, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x31,
	0x12, 0x43, 0x0a, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xc2, 0xdd, 0x29, 0x1e, 0x0a, 0x0b, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0f, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x46,
	0x69, 0x72, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x4e, 0x0a, 0x1e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x77,
	0x69, 0x74, 0x68, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6a, 0x73,
	0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xc2,
	0xdd, 0x29, 0x17, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x12, 0x0d, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x33, 0x52, 0x09, 0x6a, 0x73, 0x6f, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x67, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x42, 0x61,
	0x73, 0x65, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x32, 0x12, 0x55, 0x0a, 0x0c, 0x73, 0x70, 0x6c, 0x69,
	0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x32,
	0xc2, 0xdd, 0x29, 0x2e, 0x0a, 0x18, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x67, 0x6c, 0x79, 0x5f, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x52, 0x0b, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x22,
	0xc6, 0x02, 0x0a, 0x17, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x6c, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x4e, 0x0a, 0x11, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xc2, 0xdd, 0x29, 0x1e, 0x0a, 0x0f, 0x4f, 0x74,
	0x68, 0x65, 0x72, 0x46, 0x69, 0x72, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0b, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0f, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x46, 0x69, 0x72, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x60, 0x0a, 0x12, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x32, 0xc2, 0xdd, 0x29, 0x2e, 0x0a, 0x12, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x18, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x67, 0x6c, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x10, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x49, 0x0a,
	0x11, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xc2, 0xdd, 0x29, 0x17, 0x0a, 0x0d,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x33, 0x12, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x33, 0x52, 0x11, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x6e, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0xcd, 0x02, 0x0a, 0x11, 0x54, 0x65, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x44, 0x6f, 0x65, 0x73, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x4d,
	0x0a, 0x10, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x64, 0x6f, 0x65, 0x73, 0x6e, 0x74, 0x5f, 0x6d,
	0x61, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xc2, 0xdd, 0x29, 0x1f, 0x0a, 0x10,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x64, 0x6f, 0x65, 0x73, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x70,
	0x12, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0e, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x44, 0x6f, 0x65, 0x73, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x57, 0x0a,
	0x11, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x64, 0x6f, 0x65, 0x73, 0x6e, 0x74, 0x5f, 0x6d,
	0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xc2, 0xdd, 0x29, 0x27, 0x0a, 0x11,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x64, 0x6f, 0x65, 0x73, 0x6e, 0x74, 0x5f, 0x6d, 0x61,
	0x70, 0x12, 0x12, 0x53, 0x74, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x44, 0x6f, 0x65,
	0x73, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x60, 0x0a, 0x10, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f,
	0x64, 0x6f, 0x65, 0x73, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x36, 0xc2, 0xdd, 0x29, 0x32, 0x0a, 0x10, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x64, 0x6f,
	0x65, 0x73, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x12, 0x1e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6a,
	0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x0e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x44,
	0x6f, 0x65, 0x73, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x6e, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0xc7, 0x01, 0x0a, 0x13, 0x54, 0x65, 0x73,
	0x74, 0x4e, 0x6f, 0x74, 0x45, 0x6e, 0x6f, 0x75, 0x67, 0x68, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x12, 0x4e, 0x0a, 0x11, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xc2, 0xdd, 0x29,
	0x1e, 0x0a, 0x0f, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x69, 0x72, 0x73, 0x74, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x52,
	0x0f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x69, 0x72, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x60, 0x0a, 0x12, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x32, 0xc2, 0xdd,
	0x29, 0x2e, 0x0a, 0x12, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x18, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x67, 0x6c, 0x79,
	0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x10, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x22, 0xd2, 0x02, 0x0a, 0x15, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54,
	0x6f, 0x4f, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x4e, 0x0a, 0x11,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xc2, 0xdd, 0x29, 0x1e, 0x0a, 0x0f, 0x4f,
	0x74, 0x68, 0x65, 0x72, 0x46, 0x69, 0x72, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0b,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0f, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x46, 0x69, 0x72, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x60, 0x0a, 0x12,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x32, 0xc2, 0xdd, 0x29, 0x2e, 0x0a, 0x12,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x18, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x67, 0x6c, 0x79, 0x5f, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x10, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x47,
	0x0a, 0x11, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xc2, 0xdd, 0x29, 0x17, 0x0a,
	0x0d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x33, 0x12, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x52, 0x0f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x54, 0x68, 0x69,
	0x72, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3e, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xc2, 0xdd,
	0x29, 0x19, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0b,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0a, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_controller_custom_options_v1_testing_proto_rawDescOnce sync.Once
	file_controller_custom_options_v1_testing_proto_rawDescData = file_controller_custom_options_v1_testing_proto_rawDesc
)

func file_controller_custom_options_v1_testing_proto_rawDescGZIP() []byte {
	file_controller_custom_options_v1_testing_proto_rawDescOnce.Do(func() {
		file_controller_custom_options_v1_testing_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_custom_options_v1_testing_proto_rawDescData)
	})
	return file_controller_custom_options_v1_testing_proto_rawDescData
}

var file_controller_custom_options_v1_testing_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_controller_custom_options_v1_testing_proto_goTypes = []interface{}{
	(*TestBase)(nil),                // 0: controller.custom_options.v1.TestBase
	(*TestBaseSplit1)(nil),          // 1: controller.custom_options.v1.TestBaseSplit1
	(*TestBaseSplit2)(nil),          // 2: controller.custom_options.v1.TestBaseSplit2
	(*TestProperlyNamedFields)(nil), // 3: controller.custom_options.v1.TestProperlyNamedFields
	(*TestNameDoesntMap)(nil),       // 4: controller.custom_options.v1.TestNameDoesntMap
	(*TestNotEnoughFields)(nil),     // 5: controller.custom_options.v1.TestNotEnoughFields
	(*TestManyToOneMappings)(nil),   // 6: controller.custom_options.v1.TestManyToOneMappings
}
var file_controller_custom_options_v1_testing_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_controller_custom_options_v1_testing_proto_init() }
func file_controller_custom_options_v1_testing_proto_init() {
	if File_controller_custom_options_v1_testing_proto != nil {
		return
	}
	file_controller_custom_options_v1_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_controller_custom_options_v1_testing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestBase); i {
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
		file_controller_custom_options_v1_testing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestBaseSplit1); i {
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
		file_controller_custom_options_v1_testing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestBaseSplit2); i {
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
		file_controller_custom_options_v1_testing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestProperlyNamedFields); i {
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
		file_controller_custom_options_v1_testing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestNameDoesntMap); i {
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
		file_controller_custom_options_v1_testing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestNotEnoughFields); i {
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
		file_controller_custom_options_v1_testing_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestManyToOneMappings); i {
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
			RawDescriptor: file_controller_custom_options_v1_testing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_custom_options_v1_testing_proto_goTypes,
		DependencyIndexes: file_controller_custom_options_v1_testing_proto_depIdxs,
		MessageInfos:      file_controller_custom_options_v1_testing_proto_msgTypes,
	}.Build()
	File_controller_custom_options_v1_testing_proto = out.File
	file_controller_custom_options_v1_testing_proto_rawDesc = nil
	file_controller_custom_options_v1_testing_proto_goTypes = nil
	file_controller_custom_options_v1_testing_proto_depIdxs = nil
}
