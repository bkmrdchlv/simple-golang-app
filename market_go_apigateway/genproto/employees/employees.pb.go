// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: employees.proto

package employees

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

type Employees struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Lastname   string `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname"`
	Phone      string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone"`
	Login      string `protobuf:"bytes,5,opt,name=login,proto3" json:"login"`
	Password   string `protobuf:"bytes,6,opt,name=password,proto3" json:"password"`
	BranchId   string `protobuf:"bytes,7,opt,name=Branch_id,json=BranchId,proto3" json:"Branch_id"`
	MagazineId string `protobuf:"bytes,8,opt,name=Magazine_id,json=MagazineId,proto3" json:"Magazine_id"`
	UserType   string `protobuf:"bytes,9,opt,name=user_type,json=userType,proto3" json:"user_type"`
	Status     int32  `protobuf:"varint,10,opt,name=status,proto3" json:"status"`
}

func (x *Employees) Reset() {
	*x = Employees{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employees_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employees) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employees) ProtoMessage() {}

func (x *Employees) ProtoReflect() protoreflect.Message {
	mi := &file_employees_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employees.ProtoReflect.Descriptor instead.
func (*Employees) Descriptor() ([]byte, []int) {
	return file_employees_proto_rawDescGZIP(), []int{0}
}

func (x *Employees) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Employees) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Employees) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *Employees) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Employees) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Employees) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Employees) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *Employees) GetMagazineId() string {
	if x != nil {
		return x.MagazineId
	}
	return ""
}

func (x *Employees) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

func (x *Employees) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type Regist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Lastname   string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname"`
	Phone      string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone"`
	Login      string `protobuf:"bytes,4,opt,name=login,proto3" json:"login"`
	Password   string `protobuf:"bytes,5,opt,name=password,proto3" json:"password"`
	BranchId   string `protobuf:"bytes,6,opt,name=Branch_id,json=BranchId,proto3" json:"Branch_id"`
	MagazineId string `protobuf:"bytes,7,opt,name=Magazine_id,json=MagazineId,proto3" json:"Magazine_id"`
	UserType   string `protobuf:"bytes,8,opt,name=user_type,json=userType,proto3" json:"user_type"`
}

func (x *Regist) Reset() {
	*x = Regist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employees_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Regist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Regist) ProtoMessage() {}

func (x *Regist) ProtoReflect() protoreflect.Message {
	mi := &file_employees_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Regist.ProtoReflect.Descriptor instead.
func (*Regist) Descriptor() ([]byte, []int) {
	return file_employees_proto_rawDescGZIP(), []int{1}
}

func (x *Regist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Regist) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *Regist) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Regist) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Regist) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Regist) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *Regist) GetMagazineId() string {
	if x != nil {
		return x.MagazineId
	}
	return ""
}

func (x *Regist) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

type Loginer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password"`
}

func (x *Loginer) Reset() {
	*x = Loginer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employees_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Loginer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Loginer) ProtoMessage() {}

func (x *Loginer) ProtoReflect() protoreflect.Message {
	mi := &file_employees_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Loginer.ProtoReflect.Descriptor instead.
func (*Loginer) Descriptor() ([]byte, []int) {
	return file_employees_proto_rawDescGZIP(), []int{2}
}

func (x *Loginer) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Loginer) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type EmployeesCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Lastname   string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname"`
	Phone      string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone"`
	Login      string `protobuf:"bytes,4,opt,name=login,proto3" json:"login"`
	Password   string `protobuf:"bytes,5,opt,name=password,proto3" json:"password"`
	BranchId   string `protobuf:"bytes,6,opt,name=Branch_id,json=BranchId,proto3" json:"Branch_id"`
	MagazineId string `protobuf:"bytes,7,opt,name=Magazine_id,json=MagazineId,proto3" json:"Magazine_id"`
	UserType   string `protobuf:"bytes,8,opt,name=user_type,json=userType,proto3" json:"user_type"`
}

func (x *EmployeesCreate) Reset() {
	*x = EmployeesCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employees_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeesCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeesCreate) ProtoMessage() {}

func (x *EmployeesCreate) ProtoReflect() protoreflect.Message {
	mi := &file_employees_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeesCreate.ProtoReflect.Descriptor instead.
func (*EmployeesCreate) Descriptor() ([]byte, []int) {
	return file_employees_proto_rawDescGZIP(), []int{3}
}

func (x *EmployeesCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EmployeesCreate) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *EmployeesCreate) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *EmployeesCreate) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *EmployeesCreate) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *EmployeesCreate) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *EmployeesCreate) GetMagazineId() string {
	if x != nil {
		return x.MagazineId
	}
	return ""
}

func (x *EmployeesCreate) GetUserType() string {
	if x != nil {
		return x.UserType
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
		mi := &file_employees_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRequest) ProtoMessage() {}

func (x *GetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employees_proto_msgTypes[4]
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
	return file_employees_proto_rawDescGZIP(), []int{4}
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

	Count     int64        `protobuf:"varint,1,opt,name=count,proto3" json:"count"`
	Employees []*Employees `protobuf:"bytes,2,rep,name=employees,proto3" json:"employees"`
}

func (x *GetListResponse) Reset() {
	*x = GetListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employees_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse) ProtoMessage() {}

func (x *GetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_employees_proto_msgTypes[5]
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
	return file_employees_proto_rawDescGZIP(), []int{5}
}

func (x *GetListResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListResponse) GetEmployees() []*Employees {
	if x != nil {
		return x.Employees
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
		mi := &file_employees_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_employees_proto_msgTypes[6]
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
	return file_employees_proto_rawDescGZIP(), []int{6}
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
		mi := &file_employees_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pkey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pkey) ProtoMessage() {}

func (x *Pkey) ProtoReflect() protoreflect.Message {
	mi := &file_employees_proto_msgTypes[7]
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
	return file_employees_proto_rawDescGZIP(), []int{7}
}

func (x *Pkey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_employees_proto protoreflect.FileDescriptor

var file_employees_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x22, 0x86, 0x02, 0x0a,
	0x09, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x4d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xdb, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x4d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x61, 0x67, 0x61,
	0x7a, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x3b, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0xe4, 0x01, 0x0a, 0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x4d, 0x61, 0x67,
	0x61, 0x7a, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x4d, 0x61, 0x67, 0x61, 0x7a, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x56, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22,
	0x5b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x73, 0x52, 0x09, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x22, 0x24, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x16, 0x0a, 0x04, 0x50, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x99, 0x03, 0x0a, 0x10, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x41, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x19, 0x2e, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x32, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0f, 0x2e,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x2e, 0x50, 0x6b, 0x65, 0x79, 0x1a, 0x14,
	0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1a, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x13, 0x2e, 0x65,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x2e, 0x50, 0x6b, 0x65, 0x79, 0x1a, 0x13,
	0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x14, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x73, 0x1a, 0x13, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x08,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x32, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65, 0x72, 0x1a,
	0x13, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_employees_proto_rawDescOnce sync.Once
	file_employees_proto_rawDescData = file_employees_proto_rawDesc
)

func file_employees_proto_rawDescGZIP() []byte {
	file_employees_proto_rawDescOnce.Do(func() {
		file_employees_proto_rawDescData = protoimpl.X.CompressGZIP(file_employees_proto_rawDescData)
	})
	return file_employees_proto_rawDescData
}

var file_employees_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_employees_proto_goTypes = []interface{}{
	(*Employees)(nil),       // 0: employees.Employees
	(*Regist)(nil),          // 1: employees.Regist
	(*Loginer)(nil),         // 2: employees.Loginer
	(*EmployeesCreate)(nil), // 3: employees.EmployeesCreate
	(*GetListRequest)(nil),  // 4: employees.GetListRequest
	(*GetListResponse)(nil), // 5: employees.GetListResponse
	(*Response)(nil),        // 6: employees.Response
	(*Pkey)(nil),            // 7: employees.Pkey
}
var file_employees_proto_depIdxs = []int32{
	0, // 0: employees.GetListResponse.employees:type_name -> employees.Employees
	4, // 1: employees.EmployeesService.GetAll:input_type -> employees.GetListRequest
	7, // 2: employees.EmployeesService.GetById:input_type -> employees.Pkey
	3, // 3: employees.EmployeesService.Create:input_type -> employees.EmployeesCreate
	7, // 4: employees.EmployeesService.Delete:input_type -> employees.Pkey
	0, // 5: employees.EmployeesService.Update:input_type -> employees.Employees
	1, // 6: employees.EmployeesService.Register:input_type -> employees.Regist
	2, // 7: employees.EmployeesService.Login:input_type -> employees.Loginer
	5, // 8: employees.EmployeesService.GetAll:output_type -> employees.GetListResponse
	0, // 9: employees.EmployeesService.GetById:output_type -> employees.Employees
	6, // 10: employees.EmployeesService.Create:output_type -> employees.Response
	6, // 11: employees.EmployeesService.Delete:output_type -> employees.Response
	6, // 12: employees.EmployeesService.Update:output_type -> employees.Response
	6, // 13: employees.EmployeesService.Register:output_type -> employees.Response
	6, // 14: employees.EmployeesService.Login:output_type -> employees.Response
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_employees_proto_init() }
func file_employees_proto_init() {
	if File_employees_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_employees_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employees); i {
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
		file_employees_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Regist); i {
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
		file_employees_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Loginer); i {
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
		file_employees_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeesCreate); i {
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
		file_employees_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_employees_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_employees_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_employees_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_employees_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_employees_proto_goTypes,
		DependencyIndexes: file_employees_proto_depIdxs,
		MessageInfos:      file_employees_proto_msgTypes,
	}.Build()
	File_employees_proto = out.File
	file_employees_proto_rawDesc = nil
	file_employees_proto_goTypes = nil
	file_employees_proto_depIdxs = nil
}
