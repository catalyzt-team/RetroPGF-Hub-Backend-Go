// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: modules/users/usersPb/usersPb.proto

package RetroPGF_Hub_Backend_Go

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

// Structures
type GetUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetUserInfoReq) Reset() {
	*x = GetUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_users_usersPb_usersPb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReq) ProtoMessage() {}

func (x *GetUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_users_usersPb_usersPb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoReq.ProtoReflect.Descriptor instead.
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return file_modules_users_usersPb_usersPb_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserInfoReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Source    string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Profile   string `protobuf:"bytes,4,opt,name=profile,proto3" json:"profile,omitempty"`
	UserName  string `protobuf:"bytes,5,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	FirstName string `protobuf:"bytes,6,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,7,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *GetUserInfoRes) Reset() {
	*x = GetUserInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_users_usersPb_usersPb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoRes) ProtoMessage() {}

func (x *GetUserInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_modules_users_usersPb_usersPb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoRes.ProtoReflect.Descriptor instead.
func (*GetUserInfoRes) Descriptor() ([]byte, []int) {
	return file_modules_users_usersPb_usersPb_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserInfoRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserInfoRes) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUserInfoRes) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *GetUserInfoRes) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *GetUserInfoRes) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetUserInfoRes) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetUserInfoRes) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type GetManyUserInfoForProjectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsersId []string `protobuf:"bytes,1,rep,name=usersId,proto3" json:"usersId,omitempty"`
}

func (x *GetManyUserInfoForProjectReq) Reset() {
	*x = GetManyUserInfoForProjectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_users_usersPb_usersPb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManyUserInfoForProjectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManyUserInfoForProjectReq) ProtoMessage() {}

func (x *GetManyUserInfoForProjectReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_users_usersPb_usersPb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManyUserInfoForProjectReq.ProtoReflect.Descriptor instead.
func (*GetManyUserInfoForProjectReq) Descriptor() ([]byte, []int) {
	return file_modules_users_usersPb_usersPb_proto_rawDescGZIP(), []int{2}
}

func (x *GetManyUserInfoForProjectReq) GetUsersId() []string {
	if x != nil {
		return x.UsersId
	}
	return nil
}

type GetManyUserInfoForProjectRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsersProfile []*UserProfile `protobuf:"bytes,1,rep,name=usersProfile,proto3" json:"usersProfile,omitempty"`
}

func (x *GetManyUserInfoForProjectRes) Reset() {
	*x = GetManyUserInfoForProjectRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_users_usersPb_usersPb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManyUserInfoForProjectRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManyUserInfoForProjectRes) ProtoMessage() {}

func (x *GetManyUserInfoForProjectRes) ProtoReflect() protoreflect.Message {
	mi := &file_modules_users_usersPb_usersPb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManyUserInfoForProjectRes.ProtoReflect.Descriptor instead.
func (*GetManyUserInfoForProjectRes) Descriptor() ([]byte, []int) {
	return file_modules_users_usersPb_usersPb_proto_rawDescGZIP(), []int{3}
}

func (x *GetManyUserInfoForProjectRes) GetUsersProfile() []*UserProfile {
	if x != nil {
		return x.UsersProfile
	}
	return nil
}

type UserProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Source    string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Profile   string `protobuf:"bytes,4,opt,name=profile,proto3" json:"profile,omitempty"`
	UserName  string `protobuf:"bytes,5,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	FirstName string `protobuf:"bytes,6,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,7,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_users_usersPb_usersPb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_modules_users_usersPb_usersPb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_modules_users_usersPb_usersPb_proto_rawDescGZIP(), []int{4}
}

func (x *UserProfile) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserProfile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserProfile) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *UserProfile) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *UserProfile) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserProfile) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserProfile) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

var File_modules_users_usersPb_usersPb_proto protoreflect.FileDescriptor

var file_modules_users_usersPb_usersPb_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x50, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x50, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0xc9, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x1c, 0x47,
	0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x6f,
	0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x73, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0xc6, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x32, 0xa2, 0x01, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x73, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x59, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x6e, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x6f, 0x72,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x52, 0x65, 0x74, 0x72, 0x6f, 0x50, 0x47,
	0x46, 0x2d, 0x48, 0x75, 0x62, 0x2f, 0x52, 0x65, 0x74, 0x72, 0x6f, 0x50, 0x47, 0x46, 0x2d, 0x48,
	0x75, 0x62, 0x2d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x47, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modules_users_usersPb_usersPb_proto_rawDescOnce sync.Once
	file_modules_users_usersPb_usersPb_proto_rawDescData = file_modules_users_usersPb_usersPb_proto_rawDesc
)

func file_modules_users_usersPb_usersPb_proto_rawDescGZIP() []byte {
	file_modules_users_usersPb_usersPb_proto_rawDescOnce.Do(func() {
		file_modules_users_usersPb_usersPb_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_users_usersPb_usersPb_proto_rawDescData)
	})
	return file_modules_users_usersPb_usersPb_proto_rawDescData
}

var file_modules_users_usersPb_usersPb_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_modules_users_usersPb_usersPb_proto_goTypes = []interface{}{
	(*GetUserInfoReq)(nil),               // 0: GetUserInfoReq
	(*GetUserInfoRes)(nil),               // 1: GetUserInfoRes
	(*GetManyUserInfoForProjectReq)(nil), // 2: GetManyUserInfoForProjectReq
	(*GetManyUserInfoForProjectRes)(nil), // 3: GetManyUserInfoForProjectRes
	(*UserProfile)(nil),                  // 4: UserProfile
}
var file_modules_users_usersPb_usersPb_proto_depIdxs = []int32{
	4, // 0: GetManyUserInfoForProjectRes.usersProfile:type_name -> UserProfile
	0, // 1: UsersGrpcService.GetUserInfoById:input_type -> GetUserInfoReq
	2, // 2: UsersGrpcService.GetManyUserInfoForProject:input_type -> GetManyUserInfoForProjectReq
	1, // 3: UsersGrpcService.GetUserInfoById:output_type -> GetUserInfoRes
	3, // 4: UsersGrpcService.GetManyUserInfoForProject:output_type -> GetManyUserInfoForProjectRes
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_modules_users_usersPb_usersPb_proto_init() }
func file_modules_users_usersPb_usersPb_proto_init() {
	if File_modules_users_usersPb_usersPb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modules_users_usersPb_usersPb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoReq); i {
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
		file_modules_users_usersPb_usersPb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoRes); i {
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
		file_modules_users_usersPb_usersPb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetManyUserInfoForProjectReq); i {
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
		file_modules_users_usersPb_usersPb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetManyUserInfoForProjectRes); i {
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
		file_modules_users_usersPb_usersPb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProfile); i {
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
			RawDescriptor: file_modules_users_usersPb_usersPb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_users_usersPb_usersPb_proto_goTypes,
		DependencyIndexes: file_modules_users_usersPb_usersPb_proto_depIdxs,
		MessageInfos:      file_modules_users_usersPb_usersPb_proto_msgTypes,
	}.Build()
	File_modules_users_usersPb_usersPb_proto = out.File
	file_modules_users_usersPb_usersPb_proto_rawDesc = nil
	file_modules_users_usersPb_usersPb_proto_goTypes = nil
	file_modules_users_usersPb_usersPb_proto_depIdxs = nil
}
