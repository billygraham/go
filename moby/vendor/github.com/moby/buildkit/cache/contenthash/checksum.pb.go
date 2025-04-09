// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.11.4
// source: github.com/moby/buildkit/cache/contenthash/checksum.proto

package contenthash

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

type CacheRecordType int32

const (
	CacheRecordType_FILE       CacheRecordType = 0
	CacheRecordType_DIR        CacheRecordType = 1
	CacheRecordType_DIR_HEADER CacheRecordType = 2
	CacheRecordType_SYMLINK    CacheRecordType = 3
)

// Enum value maps for CacheRecordType.
var (
	CacheRecordType_name = map[int32]string{
		0: "FILE",
		1: "DIR",
		2: "DIR_HEADER",
		3: "SYMLINK",
	}
	CacheRecordType_value = map[string]int32{
		"FILE":       0,
		"DIR":        1,
		"DIR_HEADER": 2,
		"SYMLINK":    3,
	}
)

func (x CacheRecordType) Enum() *CacheRecordType {
	p := new(CacheRecordType)
	*p = x
	return p
}

func (x CacheRecordType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CacheRecordType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_moby_buildkit_cache_contenthash_checksum_proto_enumTypes[0].Descriptor()
}

func (CacheRecordType) Type() protoreflect.EnumType {
	return &file_github_com_moby_buildkit_cache_contenthash_checksum_proto_enumTypes[0]
}

func (x CacheRecordType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CacheRecordType.Descriptor instead.
func (CacheRecordType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_moby_buildkit_cache_contenthash_checksum_proto_rawDescGZIP(), []int{0}
}

type CacheRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Digest   string          `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	Type     CacheRecordType `protobuf:"varint,2,opt,name=type,proto3,enum=contenthash.CacheRecordType" json:"type,omitempty"`
	Linkname string          `protobuf:"bytes,3,opt,name=linkname,proto3" json:"linkname,omitempty"`
}

func (x *CacheRecord) Reset() {
	*x = CacheRecord{}
	mi := &file_github_com_moby_buildkit_cache_contenthash_checksum_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CacheRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheRecord) ProtoMessage() {}

func (x *CacheRecord) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_moby_buildkit_cache_contenthash_checksum_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheRecord.ProtoReflect.Descriptor instead.
func (*CacheRecord) Descriptor() ([]byte, []int) {
	return file_github_com_moby_buildkit_cache_contenthash_checksum_proto_rawDescGZIP(), []int{0}
}

func (x *CacheRecord) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *CacheRecord) GetType() CacheRecordType {
	if x != nil {
		return x.Type
	}
	return CacheRecordType_FILE
}

func (x *CacheRecord) GetLinkname() string {
	if x != nil {
		return x.Linkname
	}
	return ""
}

type CacheRecordWithPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path   string       `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Record *CacheRecord `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *CacheRecordWithPath) Reset() {
	*x = CacheRecordWithPath{}
	mi := &file_github_com_moby_buildkit_cache_contenthash_checksum_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CacheRecordWithPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheRecordWithPath) ProtoMessage() {}

func (x *CacheRecordWithPath) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_moby_buildkit_cache_contenthash_checksum_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheRecordWithPath.ProtoReflect.Descriptor instead.
func (*CacheRecordWithPath) Descriptor() ([]byte, []int) {
	return file_github_com_moby_buildkit_cache_contenthash_checksum_proto_rawDescGZIP(), []int{1}
}

func (x *CacheRecordWithPath) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CacheRecordWithPath) GetRecord() *CacheRecord {
	if x != nil {
		return x.Record
	}
	return nil
}

type CacheRecords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paths []*CacheRecordWithPath `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (x *CacheRecords) Reset() {
	*x = CacheRecords{}
	mi := &file_github_com_moby_buildkit_cache_contenthash_checksum_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CacheRecords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheRecords) ProtoMessage() {}

func (x *CacheRecords) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_moby_buildkit_cache_contenthash_checksum_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheRecords.ProtoReflect.Descriptor instead.
func (*CacheRecords) Descriptor() ([]byte, []int) {
	return file_github_com_moby_buildkit_cache_contenthash_checksum_proto_rawDescGZIP(), []int{2}
}

func (x *CacheRecords) GetPaths() []*CacheRecordWithPath {
	if x != nil {
		return x.Paths
	}
	return nil
}

var File_github_com_moby_buildkit_cache_contenthash_checksum_proto protoreflect.FileDescriptor

var file_github_com_moby_buildkit_cache_contenthash_checksum_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x62,
	0x79, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6b, 0x69, 0x74, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x73, 0x68, 0x2f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x73, 0x68, 0x22, 0x73, 0x0a, 0x0b, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x73, 0x68, 0x2e, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5b, 0x0a,
	0x13, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x57, 0x69, 0x74, 0x68,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x68, 0x61, 0x73, 0x68, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x46, 0x0a, 0x0c, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x05, 0x70, 0x61,
	0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x68, 0x61, 0x73, 0x68, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x57, 0x69, 0x74, 0x68, 0x50, 0x61, 0x74, 0x68, 0x52, 0x05, 0x70, 0x61, 0x74,
	0x68, 0x73, 0x2a, 0x41, 0x0a, 0x0f, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x44, 0x49, 0x52, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x49, 0x52, 0x5f,
	0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x59, 0x4d, 0x4c,
	0x49, 0x4e, 0x4b, 0x10, 0x03, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x62, 0x79, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6b, 0x69,
	0x74, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x68,
	0x61, 0x73, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_moby_buildkit_cache_contenthash_checksum_proto_rawDescOnce sync.Once
	file_github_com_moby_buildkit_cache_contenthash_checksum_proto_rawDescData = file_github_com_moby_buildkit_cache_contenthash_checksum_proto_rawDesc
)

func file_github_com_moby_buildkit_cache_contenthash_checksum_proto_rawDescGZIP() []byte {
	file_github_com_moby_buildkit_cache_contenthash_checksum_proto_rawDescOnce.Do(func() {
		file_github_com_moby_buildkit_cache_contenthash_checksum_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_moby_buildkit_cache_contenthash_checksum_proto_rawDescData)
	})
	return file_github_com_moby_buildkit_cache_contenthash_checksum_proto_rawDescData
}

var file_github_com_moby_buildkit_cache_contenthash_checksum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_moby_buildkit_cache_contenthash_checksum_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_moby_buildkit_cache_contenthash_checksum_proto_goTypes = []any{
	(CacheRecordType)(0),        // 0: contenthash.CacheRecordType
	(*CacheRecord)(nil),         // 1: contenthash.CacheRecord
	(*CacheRecordWithPath)(nil), // 2: contenthash.CacheRecordWithPath
	(*CacheRecords)(nil),        // 3: contenthash.CacheRecords
}
var file_github_com_moby_buildkit_cache_contenthash_checksum_proto_depIdxs = []int32{
	0, // 0: contenthash.CacheRecord.type:type_name -> contenthash.CacheRecordType
	1, // 1: contenthash.CacheRecordWithPath.record:type_name -> contenthash.CacheRecord
	2, // 2: contenthash.CacheRecords.paths:type_name -> contenthash.CacheRecordWithPath
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_moby_buildkit_cache_contenthash_checksum_proto_init() }
func file_github_com_moby_buildkit_cache_contenthash_checksum_proto_init() {
	if File_github_com_moby_buildkit_cache_contenthash_checksum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_moby_buildkit_cache_contenthash_checksum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_moby_buildkit_cache_contenthash_checksum_proto_goTypes,
		DependencyIndexes: file_github_com_moby_buildkit_cache_contenthash_checksum_proto_depIdxs,
		EnumInfos:         file_github_com_moby_buildkit_cache_contenthash_checksum_proto_enumTypes,
		MessageInfos:      file_github_com_moby_buildkit_cache_contenthash_checksum_proto_msgTypes,
	}.Build()
	File_github_com_moby_buildkit_cache_contenthash_checksum_proto = out.File
	file_github_com_moby_buildkit_cache_contenthash_checksum_proto_rawDesc = nil
	file_github_com_moby_buildkit_cache_contenthash_checksum_proto_goTypes = nil
	file_github_com_moby_buildkit_cache_contenthash_checksum_proto_depIdxs = nil
}
