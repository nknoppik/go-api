// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: api/models/v1/object_models.proto

package v1

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

type ObjectGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DatasetId   string                 `protobuf:"bytes,4,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	Labels      []*Label               `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty"`
	Metadata    []*Metadata            `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty"`
	Status      Status                 `protobuf:"varint,7,opt,name=status,proto3,enum=api.models.v1.Status" json:"status,omitempty"`
	Objects     []*Object              `protobuf:"bytes,8,rep,name=objects,proto3" json:"objects,omitempty"`
	Generated   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=generated,proto3" json:"generated,omitempty"`
	Created     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *ObjectGroup) Reset() {
	*x = ObjectGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_v1_object_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectGroup) ProtoMessage() {}

func (x *ObjectGroup) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_v1_object_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectGroup.ProtoReflect.Descriptor instead.
func (*ObjectGroup) Descriptor() ([]byte, []int) {
	return file_api_models_v1_object_models_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ObjectGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjectGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ObjectGroup) GetDatasetId() string {
	if x != nil {
		return x.DatasetId
	}
	return ""
}

func (x *ObjectGroup) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ObjectGroup) GetMetadata() []*Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ObjectGroup) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_INITIATING
}

func (x *ObjectGroup) GetObjects() []*Object {
	if x != nil {
		return x.Objects
	}
	return nil
}

func (x *ObjectGroup) GetGenerated() *timestamppb.Timestamp {
	if x != nil {
		return x.Generated
	}
	return nil
}

func (x *ObjectGroup) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`             //ID of the entity
	Filename      string                 `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"` // Filename: Name of the original file e.g.: mydata.json
	Filetype      string                 `protobuf:"bytes,3,opt,name=filetype,proto3" json:"filetype,omitempty"` // Filetype: Type of the stored file, e.g.: json, gbff,...
	Labels        []*Label               `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	Metadata      []*Metadata            `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty"`
	Created       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`                          // When the data object was created
	Location      *Location              `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`                        // Location: Location of the data
	Origin        *Origin                `protobuf:"bytes,8,opt,name=origin,proto3" json:"origin,omitempty"`                            // Origin: Source of the dataset
	ContentLen    int64                  `protobuf:"varint,9,opt,name=content_len,json=contentLen,proto3" json:"content_len,omitempty"` // ContentLen: Lenght of the stored dataset
	UploadId      string                 `protobuf:"bytes,10,opt,name=upload_id,json=uploadId,proto3" json:"upload_id,omitempty"`
	Generated     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=generated,proto3" json:"generated,omitempty"`
	ObjectGroupId string                 `protobuf:"bytes,12,opt,name=object_group_id,json=objectGroupId,proto3" json:"object_group_id,omitempty"`
	DatasetId     string                 `protobuf:"bytes,13,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	ProjectId     string                 `protobuf:"bytes,14,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_v1_object_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_v1_object_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_api_models_v1_object_models_proto_rawDescGZIP(), []int{1}
}

func (x *Object) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Object) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Object) GetFiletype() string {
	if x != nil {
		return x.Filetype
	}
	return ""
}

func (x *Object) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Object) GetMetadata() []*Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Object) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Object) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Object) GetOrigin() *Origin {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *Object) GetContentLen() int64 {
	if x != nil {
		return x.ContentLen
	}
	return 0
}

func (x *Object) GetUploadId() string {
	if x != nil {
		return x.UploadId
	}
	return ""
}

func (x *Object) GetGenerated() *timestamppb.Timestamp {
	if x != nil {
		return x.Generated
	}
	return nil
}

func (x *Object) GetObjectGroupId() string {
	if x != nil {
		return x.ObjectGroupId
	}
	return ""
}

func (x *Object) GetDatasetId() string {
	if x != nil {
		return x.DatasetId
	}
	return ""
}

func (x *Object) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

var File_api_models_v1_object_models_proto protoreflect.FileDescriptor

var file_api_models_v1_object_models_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x03, 0x0a, 0x0b, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2d,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a,
	0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x38,
	0x0a, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0xab,
	0x04, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x2c, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x33, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2d, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x42, 0x76, 0x0a, 0x32,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x53, 0x63, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x42, 0x2e, 0x6a, 0x61, 0x76, 0x61,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0x0c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x42, 0x2f,
	0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_models_v1_object_models_proto_rawDescOnce sync.Once
	file_api_models_v1_object_models_proto_rawDescData = file_api_models_v1_object_models_proto_rawDesc
)

func file_api_models_v1_object_models_proto_rawDescGZIP() []byte {
	file_api_models_v1_object_models_proto_rawDescOnce.Do(func() {
		file_api_models_v1_object_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_models_v1_object_models_proto_rawDescData)
	})
	return file_api_models_v1_object_models_proto_rawDescData
}

var file_api_models_v1_object_models_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_models_v1_object_models_proto_goTypes = []interface{}{
	(*ObjectGroup)(nil),           // 0: api.models.v1.ObjectGroup
	(*Object)(nil),                // 1: api.models.v1.Object
	(*Label)(nil),                 // 2: api.models.v1.Label
	(*Metadata)(nil),              // 3: api.models.v1.Metadata
	(Status)(0),                   // 4: api.models.v1.Status
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*Location)(nil),              // 6: api.models.v1.Location
	(*Origin)(nil),                // 7: api.models.v1.Origin
}
var file_api_models_v1_object_models_proto_depIdxs = []int32{
	2,  // 0: api.models.v1.ObjectGroup.labels:type_name -> api.models.v1.Label
	3,  // 1: api.models.v1.ObjectGroup.metadata:type_name -> api.models.v1.Metadata
	4,  // 2: api.models.v1.ObjectGroup.status:type_name -> api.models.v1.Status
	1,  // 3: api.models.v1.ObjectGroup.objects:type_name -> api.models.v1.Object
	5,  // 4: api.models.v1.ObjectGroup.generated:type_name -> google.protobuf.Timestamp
	5,  // 5: api.models.v1.ObjectGroup.created:type_name -> google.protobuf.Timestamp
	2,  // 6: api.models.v1.Object.labels:type_name -> api.models.v1.Label
	3,  // 7: api.models.v1.Object.metadata:type_name -> api.models.v1.Metadata
	5,  // 8: api.models.v1.Object.created:type_name -> google.protobuf.Timestamp
	6,  // 9: api.models.v1.Object.location:type_name -> api.models.v1.Location
	7,  // 10: api.models.v1.Object.origin:type_name -> api.models.v1.Origin
	5,  // 11: api.models.v1.Object.generated:type_name -> google.protobuf.Timestamp
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_api_models_v1_object_models_proto_init() }
func file_api_models_v1_object_models_proto_init() {
	if File_api_models_v1_object_models_proto != nil {
		return
	}
	file_api_models_v1_common_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_models_v1_object_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectGroup); i {
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
		file_api_models_v1_object_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
			RawDescriptor: file_api_models_v1_object_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_models_v1_object_models_proto_goTypes,
		DependencyIndexes: file_api_models_v1_object_models_proto_depIdxs,
		MessageInfos:      file_api_models_v1_object_models_proto_msgTypes,
	}.Build()
	File_api_models_v1_object_models_proto = out.File
	file_api_models_v1_object_models_proto_rawDesc = nil
	file_api_models_v1_object_models_proto_goTypes = nil
	file_api_models_v1_object_models_proto_depIdxs = nil
}
