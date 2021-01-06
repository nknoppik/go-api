// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: api/services/DatasetServiceModels.proto

package services

import (
	models "github.com/ScienceObjectsDB/go-api/models"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// Dataset related Models
type CreateDatasetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatasetName string `protobuf:"bytes,1,opt,name=DatasetName,proto3" json:"DatasetName,omitempty"` // Name of the dataset
	Datatype    string `protobuf:"bytes,2,opt,name=Datatype,proto3" json:"Datatype,omitempty"`       //Datatype of the dataset, e.g. json, gbff, fasta
	ProjectID   string `protobuf:"bytes,3,opt,name=ProjectID,proto3" json:"ProjectID,omitempty"`     //ProjectID of the project the dataset is associated with
}

func (x *CreateDatasetRequest) Reset() {
	*x = CreateDatasetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_DatasetServiceModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDatasetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatasetRequest) ProtoMessage() {}

func (x *CreateDatasetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_DatasetServiceModels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatasetRequest.ProtoReflect.Descriptor instead.
func (*CreateDatasetRequest) Descriptor() ([]byte, []int) {
	return file_api_services_DatasetServiceModels_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDatasetRequest) GetDatasetName() string {
	if x != nil {
		return x.DatasetName
	}
	return ""
}

func (x *CreateDatasetRequest) GetDatatype() string {
	if x != nil {
		return x.Datatype
	}
	return ""
}

func (x *CreateDatasetRequest) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

type DatasetVersionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatasetVersions []*models.DatasetVersionEntry `protobuf:"bytes,1,rep,name=DatasetVersions,proto3" json:"DatasetVersions,omitempty"`
}

func (x *DatasetVersionList) Reset() {
	*x = DatasetVersionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_DatasetServiceModels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetVersionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetVersionList) ProtoMessage() {}

func (x *DatasetVersionList) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_DatasetServiceModels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetVersionList.ProtoReflect.Descriptor instead.
func (*DatasetVersionList) Descriptor() ([]byte, []int) {
	return file_api_services_DatasetServiceModels_proto_rawDescGZIP(), []int{1}
}

func (x *DatasetVersionList) GetDatasetVersions() []*models.DatasetVersionEntry {
	if x != nil {
		return x.DatasetVersions
	}
	return nil
}

type ReleaseDatasetVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                               string                     `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	DatasetID                          string                     `protobuf:"bytes,2,opt,name=DatasetID,proto3" json:"DatasetID,omitempty"`
	Version                            *models.Version            `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`
	AdditionalMetadata                 map[string]*_struct.Struct `protobuf:"bytes,4,rep,name=AdditionalMetadata,proto3" json:"AdditionalMetadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`                                 // Additional metadata for the dataset version
	AdditionalMetadataMessageRef       map[string]string          `protobuf:"bytes,5,rep,name=AdditionalMetadataMessageRef,proto3" json:"AdditionalMetadataMessageRef,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`             // Message reference for the metadata
	AdditionalObjectMetadataMessageRef map[string]string          `protobuf:"bytes,6,rep,name=AdditionalObjectMetadataMessageRef,proto3" json:"AdditionalObjectMetadataMessageRef,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Message reference for the metadata of the objects associated with this DatasetVersion
	ObjectGroupIDs                     []string                   `protobuf:"bytes,7,rep,name=ObjectGroupIDs,proto3" json:"ObjectGroupIDs,omitempty"`
}

func (x *ReleaseDatasetVersionRequest) Reset() {
	*x = ReleaseDatasetVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_DatasetServiceModels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseDatasetVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseDatasetVersionRequest) ProtoMessage() {}

func (x *ReleaseDatasetVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_DatasetServiceModels_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseDatasetVersionRequest.ProtoReflect.Descriptor instead.
func (*ReleaseDatasetVersionRequest) Descriptor() ([]byte, []int) {
	return file_api_services_DatasetServiceModels_proto_rawDescGZIP(), []int{2}
}

func (x *ReleaseDatasetVersionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReleaseDatasetVersionRequest) GetDatasetID() string {
	if x != nil {
		return x.DatasetID
	}
	return ""
}

func (x *ReleaseDatasetVersionRequest) GetVersion() *models.Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *ReleaseDatasetVersionRequest) GetAdditionalMetadata() map[string]*_struct.Struct {
	if x != nil {
		return x.AdditionalMetadata
	}
	return nil
}

func (x *ReleaseDatasetVersionRequest) GetAdditionalMetadataMessageRef() map[string]string {
	if x != nil {
		return x.AdditionalMetadataMessageRef
	}
	return nil
}

func (x *ReleaseDatasetVersionRequest) GetAdditionalObjectMetadataMessageRef() map[string]string {
	if x != nil {
		return x.AdditionalObjectMetadataMessageRef
	}
	return nil
}

func (x *ReleaseDatasetVersionRequest) GetObjectGroupIDs() []string {
	if x != nil {
		return x.ObjectGroupIDs
	}
	return nil
}

var File_api_services_DatasetServiceModels_proto protoreflect.FileDescriptor

var file_api_services_DatasetServiceModels_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x02, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x3a,
	0xa1, 0x01, 0x92, 0x41, 0x9d, 0x01, 0x0a, 0x4e, 0x2a, 0x0f, 0x49, 0x6e, 0x69, 0x74, 0x4c, 0x6f,
	0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x32, 0x3b, 0x44, 0x61, 0x74, 0x61, 0x20,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x20, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x20, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x20, 0x69, 0x6e, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x42, 0x69, 0x6f,
	0x44, 0x61, 0x74, 0x61, 0x44, 0x42, 0x32, 0x4b, 0x12, 0x49, 0x7b, 0x20, 0x22, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x74, 0x65, 0x73, 0x74,
	0x22, 0x2c, 0x20, 0x22, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x20, 0x22,
	0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x2c, 0x20, 0x22, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x44, 0x22, 0x3a, 0x20, 0x22, 0x74, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x22, 0x20, 0x7d, 0x22, 0x8a, 0x01, 0x0a, 0x12, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0f, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x34, 0x92, 0x41, 0x31, 0x0a,
	0x2f, 0x2a, 0x12, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x19, 0x41, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66,
	0x20, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0xe9, 0x06, 0x0a, 0x1c, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x65, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x41, 0x64, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x83,
	0x01, 0x0a, 0x1c, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x1c, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x66, 0x12, 0x95, 0x01, 0x0a, 0x22, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x45, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x22, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x12, 0x26, 0x0a, 0x0e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x44, 0x73, 0x1a, 0x5e, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4f, 0x0a, 0x21, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x55, 0x0a, 0x27, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x3e, 0x92, 0x41,
	0x3b, 0x0a, 0x39, 0x2a, 0x1c, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x32, 0x19, 0x41, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x63, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x42, 0x2f, 0x67, 0x6f, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_services_DatasetServiceModels_proto_rawDescOnce sync.Once
	file_api_services_DatasetServiceModels_proto_rawDescData = file_api_services_DatasetServiceModels_proto_rawDesc
)

func file_api_services_DatasetServiceModels_proto_rawDescGZIP() []byte {
	file_api_services_DatasetServiceModels_proto_rawDescOnce.Do(func() {
		file_api_services_DatasetServiceModels_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_services_DatasetServiceModels_proto_rawDescData)
	})
	return file_api_services_DatasetServiceModels_proto_rawDescData
}

var file_api_services_DatasetServiceModels_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_services_DatasetServiceModels_proto_goTypes = []interface{}{
	(*CreateDatasetRequest)(nil),         // 0: CreateDatasetRequest
	(*DatasetVersionList)(nil),           // 1: DatasetVersionList
	(*ReleaseDatasetVersionRequest)(nil), // 2: ReleaseDatasetVersionRequest
	nil,                                  // 3: ReleaseDatasetVersionRequest.AdditionalMetadataEntry
	nil,                                  // 4: ReleaseDatasetVersionRequest.AdditionalMetadataMessageRefEntry
	nil,                                  // 5: ReleaseDatasetVersionRequest.AdditionalObjectMetadataMessageRefEntry
	(*models.DatasetVersionEntry)(nil),   // 6: DatasetVersionEntry
	(*models.Version)(nil),               // 7: Version
	(*_struct.Struct)(nil),               // 8: google.protobuf.Struct
}
var file_api_services_DatasetServiceModels_proto_depIdxs = []int32{
	6, // 0: DatasetVersionList.DatasetVersions:type_name -> DatasetVersionEntry
	7, // 1: ReleaseDatasetVersionRequest.Version:type_name -> Version
	3, // 2: ReleaseDatasetVersionRequest.AdditionalMetadata:type_name -> ReleaseDatasetVersionRequest.AdditionalMetadataEntry
	4, // 3: ReleaseDatasetVersionRequest.AdditionalMetadataMessageRef:type_name -> ReleaseDatasetVersionRequest.AdditionalMetadataMessageRefEntry
	5, // 4: ReleaseDatasetVersionRequest.AdditionalObjectMetadataMessageRef:type_name -> ReleaseDatasetVersionRequest.AdditionalObjectMetadataMessageRefEntry
	8, // 5: ReleaseDatasetVersionRequest.AdditionalMetadataEntry.value:type_name -> google.protobuf.Struct
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_services_DatasetServiceModels_proto_init() }
func file_api_services_DatasetServiceModels_proto_init() {
	if File_api_services_DatasetServiceModels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_services_DatasetServiceModels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDatasetRequest); i {
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
		file_api_services_DatasetServiceModels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasetVersionList); i {
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
		file_api_services_DatasetServiceModels_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseDatasetVersionRequest); i {
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
			RawDescriptor: file_api_services_DatasetServiceModels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_services_DatasetServiceModels_proto_goTypes,
		DependencyIndexes: file_api_services_DatasetServiceModels_proto_depIdxs,
		MessageInfos:      file_api_services_DatasetServiceModels_proto_msgTypes,
	}.Build()
	File_api_services_DatasetServiceModels_proto = out.File
	file_api_services_DatasetServiceModels_proto_rawDesc = nil
	file_api_services_DatasetServiceModels_proto_goTypes = nil
	file_api_services_DatasetServiceModels_proto_depIdxs = nil
}
