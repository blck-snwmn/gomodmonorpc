// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: modules/proto/order/v1/ec.proto

package v1

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

type OrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId string              `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items  []*OrderRequestItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_proto_order_v1_ec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_proto_order_v1_ec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequest.ProtoReflect.Descriptor instead.
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_modules_proto_order_v1_ec_proto_rawDescGZIP(), []int{0}
}

func (x *OrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderRequest) GetItems() []*OrderRequestItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_proto_order_v1_ec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_modules_proto_order_v1_ec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_modules_proto_order_v1_ec_proto_rawDescGZIP(), []int{1}
}

func (x *OrderResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type OrderRequestItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity int64  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *OrderRequestItem) Reset() {
	*x = OrderRequestItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_proto_order_v1_ec_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRequestItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequestItem) ProtoMessage() {}

func (x *OrderRequestItem) ProtoReflect() protoreflect.Message {
	mi := &file_modules_proto_order_v1_ec_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequestItem.ProtoReflect.Descriptor instead.
func (*OrderRequestItem) Descriptor() ([]byte, []int) {
	return file_modules_proto_order_v1_ec_proto_rawDescGZIP(), []int{0, 0}
}

func (x *OrderRequestItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderRequestItem) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_modules_proto_order_v1_ec_proto protoreflect.FileDescriptor

var file_modules_proto_order_v1_ec_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x22, 0xa4, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x1a, 0x32, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x27, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0x54, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x44, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modules_proto_order_v1_ec_proto_rawDescOnce sync.Once
	file_modules_proto_order_v1_ec_proto_rawDescData = file_modules_proto_order_v1_ec_proto_rawDesc
)

func file_modules_proto_order_v1_ec_proto_rawDescGZIP() []byte {
	file_modules_proto_order_v1_ec_proto_rawDescOnce.Do(func() {
		file_modules_proto_order_v1_ec_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_proto_order_v1_ec_proto_rawDescData)
	})
	return file_modules_proto_order_v1_ec_proto_rawDescData
}

var file_modules_proto_order_v1_ec_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_modules_proto_order_v1_ec_proto_goTypes = []interface{}{
	(*OrderRequest)(nil),     // 0: proto.order.v1.OrderRequest
	(*OrderResponse)(nil),    // 1: proto.order.v1.OrderResponse
	(*OrderRequestItem)(nil), // 2: proto.order.v1.OrderRequest.item
}
var file_modules_proto_order_v1_ec_proto_depIdxs = []int32{
	2, // 0: proto.order.v1.OrderRequest.items:type_name -> proto.order.v1.OrderRequest.item
	0, // 1: proto.order.v1.OrderService.Order:input_type -> proto.order.v1.OrderRequest
	1, // 2: proto.order.v1.OrderService.Order:output_type -> proto.order.v1.OrderResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_modules_proto_order_v1_ec_proto_init() }
func file_modules_proto_order_v1_ec_proto_init() {
	if File_modules_proto_order_v1_ec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modules_proto_order_v1_ec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRequest); i {
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
		file_modules_proto_order_v1_ec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResponse); i {
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
		file_modules_proto_order_v1_ec_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRequestItem); i {
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
			RawDescriptor: file_modules_proto_order_v1_ec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_proto_order_v1_ec_proto_goTypes,
		DependencyIndexes: file_modules_proto_order_v1_ec_proto_depIdxs,
		MessageInfos:      file_modules_proto_order_v1_ec_proto_msgTypes,
	}.Build()
	File_modules_proto_order_v1_ec_proto = out.File
	file_modules_proto_order_v1_ec_proto_rawDesc = nil
	file_modules_proto_order_v1_ec_proto_goTypes = nil
	file_modules_proto_order_v1_ec_proto_depIdxs = nil
}
