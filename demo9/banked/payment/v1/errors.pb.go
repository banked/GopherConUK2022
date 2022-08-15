// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: banked/payment/v1/errors.proto

package paymentv1

import (
	v1 "github.com/banked/gopherconuk2022/demo8/banked/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ErrInsuffcientFunds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentId   string                 `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	Amount      *v1.Amount             `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Provider    string                 `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	AttemptedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=attempted_at,json=attemptedAt,proto3" json:"attempted_at,omitempty"`
}

func (x *ErrInsuffcientFunds) Reset() {
	*x = ErrInsuffcientFunds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banked_payment_v1_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrInsuffcientFunds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrInsuffcientFunds) ProtoMessage() {}

func (x *ErrInsuffcientFunds) ProtoReflect() protoreflect.Message {
	mi := &file_banked_payment_v1_errors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrInsuffcientFunds.ProtoReflect.Descriptor instead.
func (*ErrInsuffcientFunds) Descriptor() ([]byte, []int) {
	return file_banked_payment_v1_errors_proto_rawDescGZIP(), []int{0}
}

func (x *ErrInsuffcientFunds) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

func (x *ErrInsuffcientFunds) GetAmount() *v1.Amount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *ErrInsuffcientFunds) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *ErrInsuffcientFunds) GetAttemptedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AttemptedAt
	}
	return nil
}

var File_banked_payment_v1_errors_proto protoreflect.FileDescriptor

var file_banked_payment_v1_errors_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x62, 0x61, 0x6e, 0x6b, 0x65, 0x64, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x62, 0x61, 0x6e, 0x6b, 0x65, 0x64, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x16, 0x62, 0x61, 0x6e, 0x6b, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x13, 0x45, 0x72, 0x72, 0x49, 0x6e, 0x73,
	0x75, 0x66, 0x66, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x27, 0x0a,
	0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x09, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x65, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x3d, 0x0a,
	0x0c, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0b, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0xcf, 0x01, 0x0a,
	0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x65, 0x64, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x63,
	0x6f, 0x6e, 0x75, 0x6b, 0x32, 0x30, 0x32, 0x32, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x38, 0x2f, 0x62,
	0x61, 0x6e, 0x6b, 0x65, 0x64, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x3b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x50, 0x58,
	0xaa, 0x02, 0x11, 0x42, 0x61, 0x6e, 0x6b, 0x65, 0x64, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x42, 0x61, 0x6e, 0x6b, 0x65, 0x64, 0x5c, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x42, 0x61, 0x6e, 0x6b, 0x65,
	0x64, 0x5c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x42, 0x61, 0x6e, 0x6b, 0x65,
	0x64, 0x3a, 0x3a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_banked_payment_v1_errors_proto_rawDescOnce sync.Once
	file_banked_payment_v1_errors_proto_rawDescData = file_banked_payment_v1_errors_proto_rawDesc
)

func file_banked_payment_v1_errors_proto_rawDescGZIP() []byte {
	file_banked_payment_v1_errors_proto_rawDescOnce.Do(func() {
		file_banked_payment_v1_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_banked_payment_v1_errors_proto_rawDescData)
	})
	return file_banked_payment_v1_errors_proto_rawDescData
}

var file_banked_payment_v1_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_banked_payment_v1_errors_proto_goTypes = []interface{}{
	(*ErrInsuffcientFunds)(nil),   // 0: banked.payment.v1.ErrInsuffcientFunds
	(*v1.Amount)(nil),             // 1: banked.v1.Amount
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_banked_payment_v1_errors_proto_depIdxs = []int32{
	1, // 0: banked.payment.v1.ErrInsuffcientFunds.amount:type_name -> banked.v1.Amount
	2, // 1: banked.payment.v1.ErrInsuffcientFunds.attempted_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_banked_payment_v1_errors_proto_init() }
func file_banked_payment_v1_errors_proto_init() {
	if File_banked_payment_v1_errors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_banked_payment_v1_errors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrInsuffcientFunds); i {
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
			RawDescriptor: file_banked_payment_v1_errors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_banked_payment_v1_errors_proto_goTypes,
		DependencyIndexes: file_banked_payment_v1_errors_proto_depIdxs,
		MessageInfos:      file_banked_payment_v1_errors_proto_msgTypes,
	}.Build()
	File_banked_payment_v1_errors_proto = out.File
	file_banked_payment_v1_errors_proto_rawDesc = nil
	file_banked_payment_v1_errors_proto_goTypes = nil
	file_banked_payment_v1_errors_proto_depIdxs = nil
}
