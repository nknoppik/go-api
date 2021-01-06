// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: api/services/ProjectServiceModels.proto

package services

import (
	models "github.com/ScienceObjectsDB/go-api/models"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type CreateProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *CreateProjectRequest) Reset() {
	*x = CreateProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_ProjectServiceModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectRequest) ProtoMessage() {}

func (x *CreateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_ProjectServiceModels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectRequest.ProtoReflect.Descriptor instead.
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return file_api_services_ProjectServiceModels_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProjectRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AddUserToProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    string         `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Scope     []models.Right `protobuf:"varint,2,rep,packed,name=Scope,proto3,enum=Right" json:"Scope,omitempty"`
	ProjectID string         `protobuf:"bytes,3,opt,name=ProjectID,proto3" json:"ProjectID,omitempty"`
}

func (x *AddUserToProjectRequest) Reset() {
	*x = AddUserToProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_ProjectServiceModels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserToProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserToProjectRequest) ProtoMessage() {}

func (x *AddUserToProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_ProjectServiceModels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserToProjectRequest.ProtoReflect.Descriptor instead.
func (*AddUserToProjectRequest) Descriptor() ([]byte, []int) {
	return file_api_services_ProjectServiceModels_proto_rawDescGZIP(), []int{1}
}

func (x *AddUserToProjectRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *AddUserToProjectRequest) GetScope() []models.Right {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *AddUserToProjectRequest) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

type ProjectEntryList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*models.ProjectEntry `protobuf:"bytes,1,rep,name=Projects,proto3" json:"Projects,omitempty"`
}

func (x *ProjectEntryList) Reset() {
	*x = ProjectEntryList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_ProjectServiceModels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectEntryList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectEntryList) ProtoMessage() {}

func (x *ProjectEntryList) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_ProjectServiceModels_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectEntryList.ProtoReflect.Descriptor instead.
func (*ProjectEntryList) Descriptor() ([]byte, []int) {
	return file_api_services_ProjectServiceModels_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectEntryList) GetProjects() []*models.ProjectEntry {
	if x != nil {
		return x.Projects
	}
	return nil
}

type DatasetList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datasets []*models.DatasetEntry `protobuf:"bytes,1,rep,name=Datasets,proto3" json:"Datasets,omitempty"`
}

func (x *DatasetList) Reset() {
	*x = DatasetList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_ProjectServiceModels_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetList) ProtoMessage() {}

func (x *DatasetList) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_ProjectServiceModels_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetList.ProtoReflect.Descriptor instead.
func (*DatasetList) Descriptor() ([]byte, []int) {
	return file_api_services_ProjectServiceModels_proto_rawDescGZIP(), []int{3}
}

func (x *DatasetList) GetDatasets() []*models.DatasetEntry {
	if x != nil {
		return x.Datasets
	}
	return nil
}

var File_api_services_ProjectServiceModels_proto protoreflect.FileDescriptor

var file_api_services_ProjectServiceModels_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3c, 0x92, 0x41, 0x39, 0x0a, 0x37, 0x2a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x32, 0x1f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x74, 0x6f, 0x20,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x22, 0xab, 0x01, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x06, 0x2e, 0x52, 0x69, 0x67, 0x68, 0x74, 0x52,
	0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x44, 0x3a, 0x3c, 0x92, 0x41, 0x39, 0x0a, 0x37, 0x2a, 0x17, 0x41, 0x64, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x32, 0x1c, 0x41, 0x64, 0x64, 0x73, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x74, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x3a, 0x35, 0x92, 0x41, 0x32, 0x0a, 0x30, 0x2a, 0x10, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x1c, 0x4c, 0x69, 0x73, 0x74,
	0x73, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x20, 0x6f,
	0x66, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x22, 0x60, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x3a, 0x26, 0x92, 0x41, 0x23, 0x0a, 0x21, 0x2a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x12, 0x41, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f,
	0x66, 0x20, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x42, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_services_ProjectServiceModels_proto_rawDescOnce sync.Once
	file_api_services_ProjectServiceModels_proto_rawDescData = file_api_services_ProjectServiceModels_proto_rawDesc
)

func file_api_services_ProjectServiceModels_proto_rawDescGZIP() []byte {
	file_api_services_ProjectServiceModels_proto_rawDescOnce.Do(func() {
		file_api_services_ProjectServiceModels_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_services_ProjectServiceModels_proto_rawDescData)
	})
	return file_api_services_ProjectServiceModels_proto_rawDescData
}

var file_api_services_ProjectServiceModels_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_services_ProjectServiceModels_proto_goTypes = []interface{}{
	(*CreateProjectRequest)(nil),    // 0: CreateProjectRequest
	(*AddUserToProjectRequest)(nil), // 1: AddUserToProjectRequest
	(*ProjectEntryList)(nil),        // 2: ProjectEntryList
	(*DatasetList)(nil),             // 3: DatasetList
	(models.Right)(0),               // 4: Right
	(*models.ProjectEntry)(nil),     // 5: ProjectEntry
	(*models.DatasetEntry)(nil),     // 6: DatasetEntry
}
var file_api_services_ProjectServiceModels_proto_depIdxs = []int32{
	4, // 0: AddUserToProjectRequest.Scope:type_name -> Right
	5, // 1: ProjectEntryList.Projects:type_name -> ProjectEntry
	6, // 2: DatasetList.Datasets:type_name -> DatasetEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_services_ProjectServiceModels_proto_init() }
func file_api_services_ProjectServiceModels_proto_init() {
	if File_api_services_ProjectServiceModels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_services_ProjectServiceModels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectRequest); i {
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
		file_api_services_ProjectServiceModels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserToProjectRequest); i {
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
		file_api_services_ProjectServiceModels_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectEntryList); i {
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
		file_api_services_ProjectServiceModels_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasetList); i {
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
			RawDescriptor: file_api_services_ProjectServiceModels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_services_ProjectServiceModels_proto_goTypes,
		DependencyIndexes: file_api_services_ProjectServiceModels_proto_depIdxs,
		MessageInfos:      file_api_services_ProjectServiceModels_proto_msgTypes,
	}.Build()
	File_api_services_ProjectServiceModels_proto = out.File
	file_api_services_ProjectServiceModels_proto_rawDesc = nil
	file_api_services_ProjectServiceModels_proto_goTypes = nil
	file_api_services_ProjectServiceModels_proto_depIdxs = nil
}
