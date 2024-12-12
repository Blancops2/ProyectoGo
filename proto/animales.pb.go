// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: proto/animales.proto

package proto

import (
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

type AnimalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AnimalRequest) Reset() {
	*x = AnimalRequest{}
	mi := &file_proto_animales_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnimalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimalRequest) ProtoMessage() {}

func (x *AnimalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_animales_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimalRequest.ProtoReflect.Descriptor instead.
func (*AnimalRequest) Descriptor() ([]byte, []int) {
	return file_proto_animales_proto_rawDescGZIP(), []int{0}
}

func (x *AnimalRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AnimalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	NumPatas int32  `protobuf:"varint,3,opt,name=NumPatas,proto3" json:"NumPatas,omitempty"`
}

func (x *AnimalResponse) Reset() {
	*x = AnimalResponse{}
	mi := &file_proto_animales_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnimalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimalResponse) ProtoMessage() {}

func (x *AnimalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_animales_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimalResponse.ProtoReflect.Descriptor instead.
func (*AnimalResponse) Descriptor() ([]byte, []int) {
	return file_proto_animales_proto_rawDescGZIP(), []int{1}
}

func (x *AnimalResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AnimalResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AnimalResponse) GetNumPatas() int32 {
	if x != nil {
		return x.NumPatas
	}
	return 0
}

type NewAnimalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	NumPatas int32  `protobuf:"varint,3,opt,name=NumPatas,proto3" json:"NumPatas,omitempty"`
}

func (x *NewAnimalRequest) Reset() {
	*x = NewAnimalRequest{}
	mi := &file_proto_animales_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewAnimalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAnimalRequest) ProtoMessage() {}

func (x *NewAnimalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_animales_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAnimalRequest.ProtoReflect.Descriptor instead.
func (*NewAnimalRequest) Descriptor() ([]byte, []int) {
	return file_proto_animales_proto_rawDescGZIP(), []int{2}
}

func (x *NewAnimalRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewAnimalRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *NewAnimalRequest) GetNumPatas() int32 {
	if x != nil {
		return x.NumPatas
	}
	return 0
}

type AddAnimalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *AddAnimalResponse) Reset() {
	*x = AddAnimalResponse{}
	mi := &file_proto_animales_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddAnimalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAnimalResponse) ProtoMessage() {}

func (x *AddAnimalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_animales_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAnimalResponse.ProtoReflect.Descriptor instead.
func (*AddAnimalResponse) Descriptor() ([]byte, []int) {
	return file_proto_animales_proto_rawDescGZIP(), []int{3}
}

func (x *AddAnimalResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type AnimalTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *AnimalTypeRequest) Reset() {
	*x = AnimalTypeRequest{}
	mi := &file_proto_animales_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnimalTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimalTypeRequest) ProtoMessage() {}

func (x *AnimalTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_animales_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimalTypeRequest.ProtoReflect.Descriptor instead.
func (*AnimalTypeRequest) Descriptor() ([]byte, []int) {
	return file_proto_animales_proto_rawDescGZIP(), []int{4}
}

func (x *AnimalTypeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_animales_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_animales_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_animales_proto_rawDescGZIP(), []int{5}
}

var File_proto_animales_proto protoreflect.FileDescriptor

var file_proto_animales_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x65, 0x73,
	0x22, 0x23, 0x0a, 0x0d, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x54, 0x0a, 0x0e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x4e, 0x75, 0x6d, 0x50, 0x61, 0x74, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x4e, 0x75, 0x6d, 0x50, 0x61, 0x74, 0x61, 0x73, 0x22, 0x56, 0x0a, 0x10, 0x4e,
	0x65, 0x77, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x75, 0x6d, 0x50, 0x61,
	0x74, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4e, 0x75, 0x6d, 0x50, 0x61,
	0x74, 0x61, 0x73, 0x22, 0x29, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x27,
	0x0a, 0x11, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x32, 0xb1, 0x02, 0x0a, 0x0f, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x65, 0x73, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61,
	0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c,
	0x65, 0x73, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x41, 0x6e, 0x69, 0x6d,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e,
	0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18,
	0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x48, 0x0a, 0x0b, 0x41, 0x64,
	0x64, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x61, 0x6e, 0x69, 0x6d,
	0x61, 0x6c, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x65, 0x73,
	0x2e, 0x41, 0x64, 0x64, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x28, 0x01, 0x12, 0x4e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61,
	0x6c, 0x65, 0x73, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x2e, 0x61, 0x6e, 0x69, 0x6d,
	0x61, 0x6c, 0x65, 0x73, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x65,
	0x73, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x19, 0x5a, 0x17, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x65, 0x73,
	0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_animales_proto_rawDescOnce sync.Once
	file_proto_animales_proto_rawDescData = file_proto_animales_proto_rawDesc
)

func file_proto_animales_proto_rawDescGZIP() []byte {
	file_proto_animales_proto_rawDescOnce.Do(func() {
		file_proto_animales_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_animales_proto_rawDescData)
	})
	return file_proto_animales_proto_rawDescData
}

var file_proto_animales_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_animales_proto_goTypes = []any{
	(*AnimalRequest)(nil),     // 0: animales.AnimalRequest
	(*AnimalResponse)(nil),    // 1: animales.AnimalResponse
	(*NewAnimalRequest)(nil),  // 2: animales.NewAnimalRequest
	(*AddAnimalResponse)(nil), // 3: animales.AddAnimalResponse
	(*AnimalTypeRequest)(nil), // 4: animales.AnimalTypeRequest
	(*Empty)(nil),             // 5: animales.Empty
}
var file_proto_animales_proto_depIdxs = []int32{
	0, // 0: animales.animalesservice.GetAnimalesInfo:input_type -> animales.AnimalRequest
	5, // 1: animales.animalesservice.GetAnimalesList:input_type -> animales.Empty
	2, // 2: animales.animalesservice.AddAnimales:input_type -> animales.NewAnimalRequest
	4, // 3: animales.animalesservice.GetAnimalesByType:input_type -> animales.AnimalTypeRequest
	1, // 4: animales.animalesservice.GetAnimalesInfo:output_type -> animales.AnimalResponse
	1, // 5: animales.animalesservice.GetAnimalesList:output_type -> animales.AnimalResponse
	3, // 6: animales.animalesservice.AddAnimales:output_type -> animales.AddAnimalResponse
	1, // 7: animales.animalesservice.GetAnimalesByType:output_type -> animales.AnimalResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_animales_proto_init() }
func file_proto_animales_proto_init() {
	if File_proto_animales_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_animales_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_animales_proto_goTypes,
		DependencyIndexes: file_proto_animales_proto_depIdxs,
		MessageInfos:      file_proto_animales_proto_msgTypes,
	}.Build()
	File_proto_animales_proto = out.File
	file_proto_animales_proto_rawDesc = nil
	file_proto_animales_proto_goTypes = nil
	file_proto_animales_proto_depIdxs = nil
}
