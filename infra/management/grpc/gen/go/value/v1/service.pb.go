// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: value/v1/service.proto

package valuev1

import (
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

var File_value_v1_service_proto protoreflect.FileDescriptor

var file_value_v1_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x13, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd1, 0x02, 0x0a, 0x0c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x17, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x42, 0x79, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x42, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x42, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3e, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x18, 0x2e, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb4, 0x01, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x72, 0x69, 0x62, 0x61, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x6b, 0x75, 0x6c, 0x74, 0x65, 0x6d, 0x2d, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x14, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_value_v1_service_proto_goTypes = []any{
	(*CreateRequest)(nil),       // 0: value.v1.CreateRequest
	(*ReadByGroupRequest)(nil),  // 1: value.v1.ReadByGroupRequest
	(*ReadAllRequest)(nil),      // 2: value.v1.ReadAllRequest
	(*UpdateRequest)(nil),       // 3: value.v1.UpdateRequest
	(*DeleteRequest)(nil),       // 4: value.v1.DeleteRequest
	(*CreateResponse)(nil),      // 5: value.v1.CreateResponse
	(*ReadByGroupResponse)(nil), // 6: value.v1.ReadByGroupResponse
	(*ReadAllResponse)(nil),     // 7: value.v1.ReadAllResponse
	(*UpdateResponse)(nil),      // 8: value.v1.UpdateResponse
	(*DeleteResponse)(nil),      // 9: value.v1.DeleteResponse
}
var file_value_v1_service_proto_depIdxs = []int32{
	0, // 0: value.v1.ValueService.Create:input_type -> value.v1.CreateRequest
	1, // 1: value.v1.ValueService.ReadByBatch:input_type -> value.v1.ReadByGroupRequest
	2, // 2: value.v1.ValueService.ReadAll:input_type -> value.v1.ReadAllRequest
	3, // 3: value.v1.ValueService.Update:input_type -> value.v1.UpdateRequest
	4, // 4: value.v1.ValueService.Delete:input_type -> value.v1.DeleteRequest
	5, // 5: value.v1.ValueService.Create:output_type -> value.v1.CreateResponse
	6, // 6: value.v1.ValueService.ReadByBatch:output_type -> value.v1.ReadByGroupResponse
	7, // 7: value.v1.ValueService.ReadAll:output_type -> value.v1.ReadAllResponse
	8, // 8: value.v1.ValueService.Update:output_type -> value.v1.UpdateResponse
	9, // 9: value.v1.ValueService.Delete:output_type -> value.v1.DeleteResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_value_v1_service_proto_init() }
func file_value_v1_service_proto_init() {
	if File_value_v1_service_proto != nil {
		return
	}
	file_value_v1_data_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_value_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_value_v1_service_proto_goTypes,
		DependencyIndexes: file_value_v1_service_proto_depIdxs,
	}.Build()
	File_value_v1_service_proto = out.File
	file_value_v1_service_proto_rawDesc = nil
	file_value_v1_service_proto_goTypes = nil
	file_value_v1_service_proto_depIdxs = nil
}
