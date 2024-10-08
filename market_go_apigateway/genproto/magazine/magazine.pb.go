// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: magazine.proto

package magazine

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

type Magazine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	BranchId string `protobuf:"bytes,3,opt,name=Branch_id,json=BranchId,proto3" json:"Branch_id"`
}

func (x *Magazine) Reset() {
	*x = Magazine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_magazine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Magazine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Magazine) ProtoMessage() {}

func (x *Magazine) ProtoReflect() protoreflect.Message {
	mi := &file_magazine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Magazine.ProtoReflect.Descriptor instead.
func (*Magazine) Descriptor() ([]byte, []int) {
	return file_magazine_proto_rawDescGZIP(), []int{0}
}

func (x *Magazine) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Magazine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Magazine) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

type MagazineCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	BranchId string `protobuf:"bytes,2,opt,name=Branch_id,json=BranchId,proto3" json:"Branch_id"`
}

func (x *MagazineCreate) Reset() {
	*x = MagazineCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_magazine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MagazineCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MagazineCreate) ProtoMessage() {}

func (x *MagazineCreate) ProtoReflect() protoreflect.Message {
	mi := &file_magazine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MagazineCreate.ProtoReflect.Descriptor instead.
func (*MagazineCreate) Descriptor() ([]byte, []int) {
	return file_magazine_proto_rawDescGZIP(), []int{1}
}

func (x *MagazineCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MagazineCreate) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

type GetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search"`
}

func (x *GetListRequest) Reset() {
	*x = GetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_magazine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRequest) ProtoMessage() {}

func (x *GetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_magazine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListRequest.ProtoReflect.Descriptor instead.
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return file_magazine_proto_rawDescGZIP(), []int{2}
}

func (x *GetListRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count     int64       `protobuf:"varint,1,opt,name=count,proto3" json:"count"`
	Magazines []*Magazine `protobuf:"bytes,2,rep,name=magazines,proto3" json:"magazines"`
}

func (x *GetListResponse) Reset() {
	*x = GetListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_magazine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse) ProtoMessage() {}

func (x *GetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_magazine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListResponse.ProtoReflect.Descriptor instead.
func (*GetListResponse) Descriptor() ([]byte, []int) {
	return file_magazine_proto_rawDescGZIP(), []int{3}
}

func (x *GetListResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListResponse) GetMagazines() []*Magazine {
	if x != nil {
		return x.Magazines
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_magazine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_magazine_proto_msgTypes[4]
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
	return file_magazine_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Pkey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *Pkey) Reset() {
	*x = Pkey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_magazine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pkey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pkey) ProtoMessage() {}

func (x *Pkey) ProtoReflect() protoreflect.Message {
	mi := &file_magazine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pkey.ProtoReflect.Descriptor instead.
func (*Pkey) Descriptor() ([]byte, []int) {
	return file_magazine_proto_rawDescGZIP(), []int{5}
}

func (x *Pkey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_magazine_proto protoreflect.FileDescriptor

var file_magazine_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x22, 0x4b, 0x0a, 0x08, 0x4d, 0x61,
	0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x0e, 0x4d, 0x61, 0x67, 0x61, 0x7a,
	0x69, 0x6e, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x22, 0x59, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x09, 0x6d,
	0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x61, 0x67, 0x61, 0x7a, 0x69,
	0x6e, 0x65, 0x52, 0x09, 0x6d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x24, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x16, 0x0a, 0x04, 0x50, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xa1, 0x02, 0x0a, 0x0f,
	0x4d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x18, 0x2e, 0x6d, 0x61, 0x67, 0x61,
	0x7a, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x2f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x2e, 0x6d, 0x61,
	0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x6b, 0x65, 0x79, 0x1a, 0x12, 0x2e, 0x6d, 0x61,
	0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x6d, 0x61,
	0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x12, 0x2e, 0x6d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x6d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65,
	0x2e, 0x50, 0x6b, 0x65, 0x79, 0x1a, 0x12, 0x2e, 0x6d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x6d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65,
	0x2e, 0x4d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x1a, 0x12, 0x2e, 0x6d, 0x61, 0x67, 0x61,
	0x7a, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x13, 0x5a, 0x11, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x67, 0x61,
	0x7a, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_magazine_proto_rawDescOnce sync.Once
	file_magazine_proto_rawDescData = file_magazine_proto_rawDesc
)

func file_magazine_proto_rawDescGZIP() []byte {
	file_magazine_proto_rawDescOnce.Do(func() {
		file_magazine_proto_rawDescData = protoimpl.X.CompressGZIP(file_magazine_proto_rawDescData)
	})
	return file_magazine_proto_rawDescData
}

var file_magazine_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_magazine_proto_goTypes = []interface{}{
	(*Magazine)(nil),        // 0: magazine.Magazine
	(*MagazineCreate)(nil),  // 1: magazine.MagazineCreate
	(*GetListRequest)(nil),  // 2: magazine.GetListRequest
	(*GetListResponse)(nil), // 3: magazine.GetListResponse
	(*Response)(nil),        // 4: magazine.Response
	(*Pkey)(nil),            // 5: magazine.Pkey
}
var file_magazine_proto_depIdxs = []int32{
	0, // 0: magazine.GetListResponse.magazines:type_name -> magazine.Magazine
	2, // 1: magazine.MagazineService.GetAll:input_type -> magazine.GetListRequest
	5, // 2: magazine.MagazineService.GetById:input_type -> magazine.Pkey
	1, // 3: magazine.MagazineService.Create:input_type -> magazine.MagazineCreate
	5, // 4: magazine.MagazineService.Delete:input_type -> magazine.Pkey
	0, // 5: magazine.MagazineService.Update:input_type -> magazine.Magazine
	3, // 6: magazine.MagazineService.GetAll:output_type -> magazine.GetListResponse
	0, // 7: magazine.MagazineService.GetById:output_type -> magazine.Magazine
	4, // 8: magazine.MagazineService.Create:output_type -> magazine.Response
	4, // 9: magazine.MagazineService.Delete:output_type -> magazine.Response
	4, // 10: magazine.MagazineService.Update:output_type -> magazine.Response
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_magazine_proto_init() }
func file_magazine_proto_init() {
	if File_magazine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_magazine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Magazine); i {
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
		file_magazine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MagazineCreate); i {
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
		file_magazine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListRequest); i {
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
		file_magazine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListResponse); i {
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
		file_magazine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_magazine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pkey); i {
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
			RawDescriptor: file_magazine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_magazine_proto_goTypes,
		DependencyIndexes: file_magazine_proto_depIdxs,
		MessageInfos:      file_magazine_proto_msgTypes,
	}.Build()
	File_magazine_proto = out.File
	file_magazine_proto_rawDesc = nil
	file_magazine_proto_goTypes = nil
	file_magazine_proto_depIdxs = nil
}
