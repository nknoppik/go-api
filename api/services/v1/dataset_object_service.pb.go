// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: api/services/v1/dataset_object_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_services_v1_dataset_object_service_proto protoreflect.FileDescriptor

var file_api_services_v1_dataset_object_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x33, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xfc, 0x05, 0x0a, 0x15, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x91, 0x01, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0xa5, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x26, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x67, 0x65, 0x74,
	0x12, 0x8f, 0x01, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x12, 0x8c, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x42, 0x84, 0x01, 0x0a, 0x34, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44,
	0x42, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44,
	0x42, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_services_v1_dataset_object_service_proto_goTypes = []interface{}{
	(*CreateObjectGroupRequest)(nil),       // 0: api.services.v1.CreateObjectGroupRequest
	(*CreateObjectGroupBatchRequest)(nil),  // 1: api.services.v1.CreateObjectGroupBatchRequest
	(*GetObjectGroupRequest)(nil),          // 2: api.services.v1.GetObjectGroupRequest
	(*FinishObjectUploadRequest)(nil),      // 3: api.services.v1.FinishObjectUploadRequest
	(*DeleteObjectGroupRequest)(nil),       // 4: api.services.v1.DeleteObjectGroupRequest
	(*CreateObjectGroupResponse)(nil),      // 5: api.services.v1.CreateObjectGroupResponse
	(*CreateObjectGroupBatchResponse)(nil), // 6: api.services.v1.CreateObjectGroupBatchResponse
	(*GetObjectGroupResponse)(nil),         // 7: api.services.v1.GetObjectGroupResponse
	(*FinishObjectUploadResponse)(nil),     // 8: api.services.v1.FinishObjectUploadResponse
	(*DeleteObjectGroupResponse)(nil),      // 9: api.services.v1.DeleteObjectGroupResponse
}
var file_api_services_v1_dataset_object_service_proto_depIdxs = []int32{
	0, // 0: api.services.v1.DatasetObjectsService.CreateObjectGroup:input_type -> api.services.v1.CreateObjectGroupRequest
	1, // 1: api.services.v1.DatasetObjectsService.CreateObjectGroupBatch:input_type -> api.services.v1.CreateObjectGroupBatchRequest
	2, // 2: api.services.v1.DatasetObjectsService.GetObjectGroup:input_type -> api.services.v1.GetObjectGroupRequest
	3, // 3: api.services.v1.DatasetObjectsService.FinishObjectUpload:input_type -> api.services.v1.FinishObjectUploadRequest
	4, // 4: api.services.v1.DatasetObjectsService.DeleteObjectGroup:input_type -> api.services.v1.DeleteObjectGroupRequest
	5, // 5: api.services.v1.DatasetObjectsService.CreateObjectGroup:output_type -> api.services.v1.CreateObjectGroupResponse
	6, // 6: api.services.v1.DatasetObjectsService.CreateObjectGroupBatch:output_type -> api.services.v1.CreateObjectGroupBatchResponse
	7, // 7: api.services.v1.DatasetObjectsService.GetObjectGroup:output_type -> api.services.v1.GetObjectGroupResponse
	8, // 8: api.services.v1.DatasetObjectsService.FinishObjectUpload:output_type -> api.services.v1.FinishObjectUploadResponse
	9, // 9: api.services.v1.DatasetObjectsService.DeleteObjectGroup:output_type -> api.services.v1.DeleteObjectGroupResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_services_v1_dataset_object_service_proto_init() }
func file_api_services_v1_dataset_object_service_proto_init() {
	if File_api_services_v1_dataset_object_service_proto != nil {
		return
	}
	file_api_services_v1_dataset_object_service_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_services_v1_dataset_object_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_services_v1_dataset_object_service_proto_goTypes,
		DependencyIndexes: file_api_services_v1_dataset_object_service_proto_depIdxs,
	}.Build()
	File_api_services_v1_dataset_object_service_proto = out.File
	file_api_services_v1_dataset_object_service_proto_rawDesc = nil
	file_api_services_v1_dataset_object_service_proto_goTypes = nil
	file_api_services_v1_dataset_object_service_proto_depIdxs = nil
}
