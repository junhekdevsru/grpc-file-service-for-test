// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/file_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Загрузка файла, клиент отправляет чанки
type UploadFileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      string                 `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"` // только в первом чанке обязательно
	Content       []byte                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`   // содержимое файла
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	mi := &file_proto_file_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_file_service_proto_rawDescGZIP(), []int{0}
}

func (x *UploadFileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *UploadFileRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

// Ответ после загрузки
type UploadFileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadFileResponse) Reset() {
	*x = UploadFileResponse{}
	mi := &file_proto_file_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResponse) ProtoMessage() {}

func (x *UploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResponse.ProtoReflect.Descriptor instead.
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_file_service_proto_rawDescGZIP(), []int{1}
}

func (x *UploadFileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Запрос на скачивание файла
type DownloadFileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      string                 `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadFileRequest) Reset() {
	*x = DownloadFileRequest{}
	mi := &file_proto_file_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileRequest) ProtoMessage() {}

func (x *DownloadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileRequest.ProtoReflect.Descriptor instead.
func (*DownloadFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_file_service_proto_rawDescGZIP(), []int{2}
}

func (x *DownloadFileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

// Ответ чанком файла
type DownloadFileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       []byte                 `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadFileResponse) Reset() {
	*x = DownloadFileResponse{}
	mi := &file_proto_file_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileResponse) ProtoMessage() {}

func (x *DownloadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileResponse.ProtoReflect.Descriptor instead.
func (*DownloadFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_file_service_proto_rawDescGZIP(), []int{3}
}

func (x *DownloadFileResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

// Запрос списка файлов (пустой)
type ListFilesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFilesRequest) Reset() {
	*x = ListFilesRequest{}
	mi := &file_proto_file_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFilesRequest) ProtoMessage() {}

func (x *ListFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFilesRequest.ProtoReflect.Descriptor instead.
func (*ListFilesRequest) Descriptor() ([]byte, []int) {
	return file_proto_file_service_proto_rawDescGZIP(), []int{4}
}

// Ответ со списком файлов
type ListFilesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Files         []*FileInfo            `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFilesResponse) Reset() {
	*x = ListFilesResponse{}
	mi := &file_proto_file_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFilesResponse) ProtoMessage() {}

func (x *ListFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFilesResponse.ProtoReflect.Descriptor instead.
func (*ListFilesResponse) Descriptor() ([]byte, []int) {
	return file_proto_file_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListFilesResponse) GetFiles() []*FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

// Информация о файле
type FileInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // можно использовать формат RFC3339
	UpdatedAt     string                 `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	mi := &file_proto_file_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_proto_file_service_proto_rawDescGZIP(), []int{6}
}

func (x *FileInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *FileInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_proto_file_service_proto protoreflect.FileDescriptor

const file_proto_file_service_proto_rawDesc = "" +
	"\n" +
	"\x18proto/file_service.proto\x12\x04file\"I\n" +
	"\x11UploadFileRequest\x12\x1a\n" +
	"\bfilename\x18\x01 \x01(\tR\bfilename\x12\x18\n" +
	"\acontent\x18\x02 \x01(\fR\acontent\".\n" +
	"\x12UploadFileResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"1\n" +
	"\x13DownloadFileRequest\x12\x1a\n" +
	"\bfilename\x18\x01 \x01(\tR\bfilename\"0\n" +
	"\x14DownloadFileResponse\x12\x18\n" +
	"\acontent\x18\x01 \x01(\fR\acontent\"\x12\n" +
	"\x10ListFilesRequest\"9\n" +
	"\x11ListFilesResponse\x12$\n" +
	"\x05files\x18\x01 \x03(\v2\x0e.file.FileInfoR\x05files\"\\\n" +
	"\bFileInfo\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1d\n" +
	"\n" +
	"created_at\x18\x02 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\x03 \x01(\tR\tupdatedAt2\xd7\x01\n" +
	"\vFileService\x12A\n" +
	"\n" +
	"UploadFile\x12\x17.file.UploadFileRequest\x1a\x18.file.UploadFileResponse(\x01\x12G\n" +
	"\fDownloadFile\x12\x19.file.DownloadFileRequest\x1a\x1a.file.DownloadFileResponse0\x01\x12<\n" +
	"\tListFiles\x12\x16.file.ListFilesRequest\x1a\x17.file.ListFilesResponseB\x1fZ\x1dgrpc-file-service/proto;protob\x06proto3"

var (
	file_proto_file_service_proto_rawDescOnce sync.Once
	file_proto_file_service_proto_rawDescData []byte
)

func file_proto_file_service_proto_rawDescGZIP() []byte {
	file_proto_file_service_proto_rawDescOnce.Do(func() {
		file_proto_file_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_file_service_proto_rawDesc), len(file_proto_file_service_proto_rawDesc)))
	})
	return file_proto_file_service_proto_rawDescData
}

var file_proto_file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_file_service_proto_goTypes = []any{
	(*UploadFileRequest)(nil),    // 0: file.UploadFileRequest
	(*UploadFileResponse)(nil),   // 1: file.UploadFileResponse
	(*DownloadFileRequest)(nil),  // 2: file.DownloadFileRequest
	(*DownloadFileResponse)(nil), // 3: file.DownloadFileResponse
	(*ListFilesRequest)(nil),     // 4: file.ListFilesRequest
	(*ListFilesResponse)(nil),    // 5: file.ListFilesResponse
	(*FileInfo)(nil),             // 6: file.FileInfo
}
var file_proto_file_service_proto_depIdxs = []int32{
	6, // 0: file.ListFilesResponse.files:type_name -> file.FileInfo
	0, // 1: file.FileService.UploadFile:input_type -> file.UploadFileRequest
	2, // 2: file.FileService.DownloadFile:input_type -> file.DownloadFileRequest
	4, // 3: file.FileService.ListFiles:input_type -> file.ListFilesRequest
	1, // 4: file.FileService.UploadFile:output_type -> file.UploadFileResponse
	3, // 5: file.FileService.DownloadFile:output_type -> file.DownloadFileResponse
	5, // 6: file.FileService.ListFiles:output_type -> file.ListFilesResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_file_service_proto_init() }
func file_proto_file_service_proto_init() {
	if File_proto_file_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_file_service_proto_rawDesc), len(file_proto_file_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_file_service_proto_goTypes,
		DependencyIndexes: file_proto_file_service_proto_depIdxs,
		MessageInfos:      file_proto_file_service_proto_msgTypes,
	}.Build()
	File_proto_file_service_proto = out.File
	file_proto_file_service_proto_goTypes = nil
	file_proto_file_service_proto_depIdxs = nil
}
