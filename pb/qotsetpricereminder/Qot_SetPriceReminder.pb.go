// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: Qot_SetPriceReminder.proto

package qotsetpricereminder

import (
	_ "github.com/futuopen/ftapi4go/pb/common"
	qotcommon "github.com/futuopen/ftapi4go/pb/qotcommon"
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

type SetPriceReminderOp int32

const (
	SetPriceReminderOp_SetPriceReminderOp_Unknown SetPriceReminderOp = 0
	SetPriceReminderOp_SetPriceReminderOp_Add     SetPriceReminderOp = 1 //新增
	SetPriceReminderOp_SetPriceReminderOp_Del     SetPriceReminderOp = 2 //删除
	SetPriceReminderOp_SetPriceReminderOp_Enable  SetPriceReminderOp = 3 //启用
	SetPriceReminderOp_SetPriceReminderOp_Disable SetPriceReminderOp = 4 //禁用
	SetPriceReminderOp_SetPriceReminderOp_Modify  SetPriceReminderOp = 5 //修改
	SetPriceReminderOp_SetPriceReminderOp_DelAll  SetPriceReminderOp = 6 //删除该支股票下所有到价提醒
)

// Enum value maps for SetPriceReminderOp.
var (
	SetPriceReminderOp_name = map[int32]string{
		0: "SetPriceReminderOp_Unknown",
		1: "SetPriceReminderOp_Add",
		2: "SetPriceReminderOp_Del",
		3: "SetPriceReminderOp_Enable",
		4: "SetPriceReminderOp_Disable",
		5: "SetPriceReminderOp_Modify",
		6: "SetPriceReminderOp_DelAll",
	}
	SetPriceReminderOp_value = map[string]int32{
		"SetPriceReminderOp_Unknown": 0,
		"SetPriceReminderOp_Add":     1,
		"SetPriceReminderOp_Del":     2,
		"SetPriceReminderOp_Enable":  3,
		"SetPriceReminderOp_Disable": 4,
		"SetPriceReminderOp_Modify":  5,
		"SetPriceReminderOp_DelAll":  6,
	}
)

func (x SetPriceReminderOp) Enum() *SetPriceReminderOp {
	p := new(SetPriceReminderOp)
	*p = x
	return p
}

func (x SetPriceReminderOp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetPriceReminderOp) Descriptor() protoreflect.EnumDescriptor {
	return file_Qot_SetPriceReminder_proto_enumTypes[0].Descriptor()
}

func (SetPriceReminderOp) Type() protoreflect.EnumType {
	return &file_Qot_SetPriceReminder_proto_enumTypes[0]
}

func (x SetPriceReminderOp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SetPriceReminderOp) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SetPriceReminderOp(num)
	return nil
}

// Deprecated: Use SetPriceReminderOp.Descriptor instead.
func (SetPriceReminderOp) EnumDescriptor() ([]byte, []int) {
	return file_Qot_SetPriceReminder_proto_rawDescGZIP(), []int{0}
}

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Security *qotcommon.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"` // 股票
	Op       *int32              `protobuf:"varint,2,req,name=op" json:"op,omitempty"`            // SetPriceReminderOp，操作类型
	Key      *int64              `protobuf:"varint,3,opt,name=key" json:"key,omitempty"`          // 到价提醒的标识，GetPriceReminder协议可获得，用于指定要操作的到价提醒项，对于新增的情况不需要填
	Type     *int32              `protobuf:"varint,4,opt,name=type" json:"type,omitempty"`        // Qot_Common::PriceReminderType，提醒类型，删除、启用、禁用的情况下会忽略该字段
	Freq     *int32              `protobuf:"varint,7,opt,name=freq" json:"freq,omitempty"`        // Qot_Common::PriceReminderFreq，提醒频率类型，删除、启用、禁用的情况下会忽略该字段
	Value    *float64            `protobuf:"fixed64,5,opt,name=value" json:"value,omitempty"`     // 提醒值，删除、启用、禁用的情况下会忽略该字段（精确到小数点后 3 位，超出部分会被舍弃）
	Note     *string             `protobuf:"bytes,6,opt,name=note" json:"note,omitempty"`         // 用户设置到价提醒时的标注，仅支持 20 个以内的中文字符，删除、启用、禁用的情况下会忽略该字段
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_SetPriceReminder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_SetPriceReminder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S.ProtoReflect.Descriptor instead.
func (*C2S) Descriptor() ([]byte, []int) {
	return file_Qot_SetPriceReminder_proto_rawDescGZIP(), []int{0}
}

func (x *C2S) GetSecurity() *qotcommon.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *C2S) GetOp() int32 {
	if x != nil && x.Op != nil {
		return *x.Op
	}
	return 0
}

func (x *C2S) GetKey() int64 {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return 0
}

func (x *C2S) GetType() int32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *C2S) GetFreq() int32 {
	if x != nil && x.Freq != nil {
		return *x.Freq
	}
	return 0
}

func (x *C2S) GetValue() float64 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

func (x *C2S) GetNote() string {
	if x != nil && x.Note != nil {
		return *x.Note
	}
	return ""
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *int64 `protobuf:"varint,1,req,name=key" json:"key,omitempty"` //设置成功的情况下返回对应的key，不成功返回0
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_SetPriceReminder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_SetPriceReminder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C.ProtoReflect.Descriptor instead.
func (*S2C) Descriptor() ([]byte, []int) {
	return file_Qot_SetPriceReminder_proto_rawDescGZIP(), []int{1}
}

func (x *S2C) GetKey() int64 {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return 0
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_SetPriceReminder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_SetPriceReminder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_Qot_SetPriceReminder_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetC2S() *C2S {
	if x != nil {
		return x.C2S
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"` //RetType，返回结果
	RetMsg  *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C     *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

// Default values for Response fields.
const (
	Default_Response_RetType = int32(-400)
)

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_SetPriceReminder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_SetPriceReminder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_Qot_SetPriceReminder_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetRetType() int32 {
	if x != nil && x.RetType != nil {
		return *x.RetType
	}
	return Default_Response_RetType
}

func (x *Response) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *Response) GetErrCode() int32 {
	if x != nil && x.ErrCode != nil {
		return *x.ErrCode
	}
	return 0
}

func (x *Response) GetS2C() *S2C {
	if x != nil {
		return x.S2C
	}
	return nil
}

var File_Qot_SetPriceReminder_proto protoreflect.FileDescriptor

var file_Qot_SetPriceReminder_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x51, 0x6f, 0x74, 0x5f, 0x53, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x51, 0x6f,
	0x74, 0x5f, 0x53, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64,
	0x65, 0x72, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x03, 0x43, 0x32, 0x53, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51,
	0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x70, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x65, 0x71, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x66, 0x72, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x6f, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65,
	0x22, 0x17, 0x0a, 0x03, 0x53, 0x32, 0x43, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x36, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x03, 0x63, 0x32, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x53, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x32, 0x53, 0x52, 0x03, 0x63, 0x32,
	0x73, 0x22, 0x89, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x3a,
	0x04, 0x2d, 0x34, 0x30, 0x30, 0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x2b, 0x0a, 0x03, 0x73, 0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x51, 0x6f, 0x74, 0x5f, 0x53, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x2a, 0xe9, 0x01,
	0x0a, 0x12, 0x53, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64,
	0x65, 0x72, 0x4f, 0x70, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x4f, 0x70, 0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x4f, 0x70, 0x5f, 0x41, 0x64, 0x64, 0x10, 0x01,
	0x12, 0x1a, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x4f, 0x70, 0x5f, 0x44, 0x65, 0x6c, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19,
	0x53, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72,
	0x4f, 0x70, 0x5f, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x53,
	0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x4f,
	0x70, 0x5f, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x53,
	0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x4f,
	0x70, 0x5f, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x65,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x4f, 0x70,
	0x5f, 0x44, 0x65, 0x6c, 0x41, 0x6c, 0x6c, 0x10, 0x06, 0x42, 0x4a, 0x0a, 0x13, 0x63, 0x6f, 0x6d,
	0x2e, 0x66, 0x75, 0x74, 0x75, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75, 0x74,
	0x75, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x66, 0x74, 0x61, 0x70, 0x69, 0x34, 0x67, 0x6f, 0x2f, 0x70,
	0x62, 0x2f, 0x71, 0x6f, 0x74, 0x73, 0x65, 0x74, 0x70, 0x72, 0x69, 0x63, 0x65, 0x72, 0x65, 0x6d,
	0x69, 0x6e, 0x64, 0x65, 0x72,
}

var (
	file_Qot_SetPriceReminder_proto_rawDescOnce sync.Once
	file_Qot_SetPriceReminder_proto_rawDescData = file_Qot_SetPriceReminder_proto_rawDesc
)

func file_Qot_SetPriceReminder_proto_rawDescGZIP() []byte {
	file_Qot_SetPriceReminder_proto_rawDescOnce.Do(func() {
		file_Qot_SetPriceReminder_proto_rawDescData = protoimpl.X.CompressGZIP(file_Qot_SetPriceReminder_proto_rawDescData)
	})
	return file_Qot_SetPriceReminder_proto_rawDescData
}

var file_Qot_SetPriceReminder_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Qot_SetPriceReminder_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Qot_SetPriceReminder_proto_goTypes = []interface{}{
	(SetPriceReminderOp)(0),    // 0: Qot_SetPriceReminder.SetPriceReminderOp
	(*C2S)(nil),                // 1: Qot_SetPriceReminder.C2S
	(*S2C)(nil),                // 2: Qot_SetPriceReminder.S2C
	(*Request)(nil),            // 3: Qot_SetPriceReminder.Request
	(*Response)(nil),           // 4: Qot_SetPriceReminder.Response
	(*qotcommon.Security)(nil), // 5: Qot_Common.Security
}
var file_Qot_SetPriceReminder_proto_depIdxs = []int32{
	5, // 0: Qot_SetPriceReminder.C2S.security:type_name -> Qot_Common.Security
	1, // 1: Qot_SetPriceReminder.Request.c2s:type_name -> Qot_SetPriceReminder.C2S
	2, // 2: Qot_SetPriceReminder.Response.s2c:type_name -> Qot_SetPriceReminder.S2C
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Qot_SetPriceReminder_proto_init() }
func file_Qot_SetPriceReminder_proto_init() {
	if File_Qot_SetPriceReminder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Qot_SetPriceReminder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S); i {
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
		file_Qot_SetPriceReminder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C); i {
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
		file_Qot_SetPriceReminder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_Qot_SetPriceReminder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_Qot_SetPriceReminder_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Qot_SetPriceReminder_proto_goTypes,
		DependencyIndexes: file_Qot_SetPriceReminder_proto_depIdxs,
		EnumInfos:         file_Qot_SetPriceReminder_proto_enumTypes,
		MessageInfos:      file_Qot_SetPriceReminder_proto_msgTypes,
	}.Build()
	File_Qot_SetPriceReminder_proto = out.File
	file_Qot_SetPriceReminder_proto_rawDesc = nil
	file_Qot_SetPriceReminder_proto_goTypes = nil
	file_Qot_SetPriceReminder_proto_depIdxs = nil
}
