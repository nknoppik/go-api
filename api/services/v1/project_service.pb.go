// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: api/services/v1/project_service.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

var File_api_services_v1_project_service_proto protoreflect.FileDescriptor

var file_api_services_v1_project_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcd, 0x09, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22,
	0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x94, 0x01, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x64, 0x64, 0x75, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x81, 0x01, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x61, 0x70, 0x69, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x9b, 0x01,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x73, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x12, 0x7e, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x27,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x73, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x72, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x83, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x50,
	0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x2a, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x70, 0x69, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0xd8, 0x06, 0x0a, 0x34, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x44, 0x42, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x42,
	0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x92, 0x41, 0xd4, 0x05, 0x12, 0x20, 0x0a, 0x19, 0x53,
	0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x42, 0x20,
	0x52, 0x45, 0x53, 0x54, 0x20, 0x41, 0x50, 0x49, 0x32, 0x03, 0x30, 0x2e, 0x31, 0x2a, 0x02, 0x02,
	0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x50, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x49, 0x0a, 0x47,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20,
	0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20,
	0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x3b, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x34,
	0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65,
	0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x12, 0x06, 0x0a, 0x04,
	0x9a, 0x02, 0x01, 0x07, 0x52, 0x57, 0x0a, 0x03, 0x34, 0x31, 0x38, 0x12, 0x50, 0x0a, 0x0d, 0x49,
	0x27, 0x6d, 0x20, 0x61, 0x20, 0x74, 0x65, 0x61, 0x70, 0x6f, 0x74, 0x2e, 0x12, 0x3f, 0x0a, 0x3d,
	0x1a, 0x3b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70,
	0x62, 0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x9b, 0x02,
	0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x93, 0x02, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x3f, 0x1a, 0x3d, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0xbf, 0x01, 0x0a, 0x10, 0x58, 0x2d,
	0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x49, 0x64, 0x12, 0xaa,
	0x01, 0x0a, 0x2b, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x20,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x06,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x32, 0x26, 0x22, 0x32,
	0x34, 0x33, 0x38, 0x61, 0x63, 0x33, 0x63, 0x2d, 0x33, 0x37, 0x65, 0x62, 0x2d, 0x34, 0x39, 0x30,
	0x32, 0x2d, 0x61, 0x64, 0x65, 0x66, 0x2d, 0x65, 0x64, 0x31, 0x36, 0x62, 0x34, 0x34, 0x33, 0x31,
	0x30, 0x33, 0x30, 0x22, 0x6a, 0x45, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b,
	0x38, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x34, 0x7d, 0x2d, 0x34,
	0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x38, 0x39, 0x41,
	0x42, 0x5d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x30,
	0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x31, 0x32, 0x7d, 0x24, 0x5a, 0x5d, 0x0a, 0x2e, 0x0a,
	0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1d,
	0x08, 0x02, 0x1a, 0x17, 0x47, 0x72, 0x70, 0x63, 0x2d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6b, 0x65, 0x79, 0x20, 0x02, 0x0a, 0x2b, 0x0a,
	0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1d, 0x08, 0x02, 0x1a,
	0x17, 0x47, 0x72, 0x70, 0x63, 0x2d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x41,
	0x50, 0x49, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x20, 0x02, 0x62, 0x23, 0x0a, 0x11, 0x0a, 0x0d,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x0a,
	0x0e, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_services_v1_project_service_proto_goTypes = []interface{}{
	(*CreateProjectRequest)(nil),       // 0: api.services.v1.CreateProjectRequest
	(*AddUserToProjectRequest)(nil),    // 1: api.services.v1.AddUserToProjectRequest
	(*CreateAPITokenRequest)(nil),      // 2: api.services.v1.CreateAPITokenRequest
	(*GetProjectDatasetsRequest)(nil),  // 3: api.services.v1.GetProjectDatasetsRequest
	(*GetUserProjectsRequest)(nil),     // 4: api.services.v1.GetUserProjectsRequest
	(*GetProjectRequest)(nil),          // 5: api.services.v1.GetProjectRequest
	(*GetAPITokenRequest)(nil),         // 6: api.services.v1.GetAPITokenRequest
	(*DeleteProjectRequest)(nil),       // 7: api.services.v1.DeleteProjectRequest
	(*DeleteAPITokenRequest)(nil),      // 8: api.services.v1.DeleteAPITokenRequest
	(*CreateProjectResponse)(nil),      // 9: api.services.v1.CreateProjectResponse
	(*AddUserToProjectResponse)(nil),   // 10: api.services.v1.AddUserToProjectResponse
	(*CreateAPITokenResponse)(nil),     // 11: api.services.v1.CreateAPITokenResponse
	(*GetProjectDatasetsResponse)(nil), // 12: api.services.v1.GetProjectDatasetsResponse
	(*GetUserProjectsResponse)(nil),    // 13: api.services.v1.GetUserProjectsResponse
	(*GetProjectResponse)(nil),         // 14: api.services.v1.GetProjectResponse
	(*GetAPITokenResponse)(nil),        // 15: api.services.v1.GetAPITokenResponse
	(*DeleteProjectResponse)(nil),      // 16: api.services.v1.DeleteProjectResponse
	(*DeleteAPITokenResponse)(nil),     // 17: api.services.v1.DeleteAPITokenResponse
}
var file_api_services_v1_project_service_proto_depIdxs = []int32{
	0,  // 0: api.services.v1.ProjectService.CreateProject:input_type -> api.services.v1.CreateProjectRequest
	1,  // 1: api.services.v1.ProjectService.AddUserToProject:input_type -> api.services.v1.AddUserToProjectRequest
	2,  // 2: api.services.v1.ProjectService.CreateAPIToken:input_type -> api.services.v1.CreateAPITokenRequest
	3,  // 3: api.services.v1.ProjectService.GetProjectDatasets:input_type -> api.services.v1.GetProjectDatasetsRequest
	4,  // 4: api.services.v1.ProjectService.GetUserProjects:input_type -> api.services.v1.GetUserProjectsRequest
	5,  // 5: api.services.v1.ProjectService.GetProject:input_type -> api.services.v1.GetProjectRequest
	6,  // 6: api.services.v1.ProjectService.GetAPIToken:input_type -> api.services.v1.GetAPITokenRequest
	7,  // 7: api.services.v1.ProjectService.DeleteProject:input_type -> api.services.v1.DeleteProjectRequest
	8,  // 8: api.services.v1.ProjectService.DeleteAPIToken:input_type -> api.services.v1.DeleteAPITokenRequest
	9,  // 9: api.services.v1.ProjectService.CreateProject:output_type -> api.services.v1.CreateProjectResponse
	10, // 10: api.services.v1.ProjectService.AddUserToProject:output_type -> api.services.v1.AddUserToProjectResponse
	11, // 11: api.services.v1.ProjectService.CreateAPIToken:output_type -> api.services.v1.CreateAPITokenResponse
	12, // 12: api.services.v1.ProjectService.GetProjectDatasets:output_type -> api.services.v1.GetProjectDatasetsResponse
	13, // 13: api.services.v1.ProjectService.GetUserProjects:output_type -> api.services.v1.GetUserProjectsResponse
	14, // 14: api.services.v1.ProjectService.GetProject:output_type -> api.services.v1.GetProjectResponse
	15, // 15: api.services.v1.ProjectService.GetAPIToken:output_type -> api.services.v1.GetAPITokenResponse
	16, // 16: api.services.v1.ProjectService.DeleteProject:output_type -> api.services.v1.DeleteProjectResponse
	17, // 17: api.services.v1.ProjectService.DeleteAPIToken:output_type -> api.services.v1.DeleteAPITokenResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_services_v1_project_service_proto_init() }
func file_api_services_v1_project_service_proto_init() {
	if File_api_services_v1_project_service_proto != nil {
		return
	}
	file_api_services_v1_project_service_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_services_v1_project_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_services_v1_project_service_proto_goTypes,
		DependencyIndexes: file_api_services_v1_project_service_proto_depIdxs,
	}.Build()
	File_api_services_v1_project_service_proto = out.File
	file_api_services_v1_project_service_proto_rawDesc = nil
	file_api_services_v1_project_service_proto_goTypes = nil
	file_api_services_v1_project_service_proto_depIdxs = nil
}
