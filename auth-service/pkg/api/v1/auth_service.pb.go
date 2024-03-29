// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateUserResponse_ErrorCode int32

const (
	CreateUserResponse_BAD_REQUEST      CreateUserResponse_ErrorCode = 0
	CreateUserResponse_INTERNAL_ERROR   CreateUserResponse_ErrorCode = 1
	CreateUserResponse_INVALID_EMAIL    CreateUserResponse_ErrorCode = 2
	CreateUserResponse_INVALID_USERNAME CreateUserResponse_ErrorCode = 3
	CreateUserResponse_USERNAME_TAKEN   CreateUserResponse_ErrorCode = 4
	CreateUserResponse_EMAIL_IN_USE     CreateUserResponse_ErrorCode = 5
	CreateUserResponse_BLANK_USERNAME   CreateUserResponse_ErrorCode = 6
	CreateUserResponse_BLANK_PASSWORD   CreateUserResponse_ErrorCode = 7
)

var CreateUserResponse_ErrorCode_name = map[int32]string{
	0: "BAD_REQUEST",
	1: "INTERNAL_ERROR",
	2: "INVALID_EMAIL",
	3: "INVALID_USERNAME",
	4: "USERNAME_TAKEN",
	5: "EMAIL_IN_USE",
	6: "BLANK_USERNAME",
	7: "BLANK_PASSWORD",
}

var CreateUserResponse_ErrorCode_value = map[string]int32{
	"BAD_REQUEST":      0,
	"INTERNAL_ERROR":   1,
	"INVALID_EMAIL":    2,
	"INVALID_USERNAME": 3,
	"USERNAME_TAKEN":   4,
	"EMAIL_IN_USE":     5,
	"BLANK_USERNAME":   6,
	"BLANK_PASSWORD":   7,
}

func (x CreateUserResponse_ErrorCode) String() string {
	return proto.EnumName(CreateUserResponse_ErrorCode_name, int32(x))
}

func (CreateUserResponse_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{1, 0}
}

type SetUserValidationResponse_ErrorCode int32

const (
	SetUserValidationResponse_BAD_REQUEST            SetUserValidationResponse_ErrorCode = 0
	SetUserValidationResponse_INTERNAL_ERROR         SetUserValidationResponse_ErrorCode = 1
	SetUserValidationResponse_USER_DOES_NOT_EXIST    SetUserValidationResponse_ErrorCode = 2
	SetUserValidationResponse_USER_ALREADY_VALIDATED SetUserValidationResponse_ErrorCode = 3
)

var SetUserValidationResponse_ErrorCode_name = map[int32]string{
	0: "BAD_REQUEST",
	1: "INTERNAL_ERROR",
	2: "USER_DOES_NOT_EXIST",
	3: "USER_ALREADY_VALIDATED",
}

var SetUserValidationResponse_ErrorCode_value = map[string]int32{
	"BAD_REQUEST":            0,
	"INTERNAL_ERROR":         1,
	"USER_DOES_NOT_EXIST":    2,
	"USER_ALREADY_VALIDATED": 3,
}

func (x SetUserValidationResponse_ErrorCode) String() string {
	return proto.EnumName(SetUserValidationResponse_ErrorCode_name, int32(x))
}

func (SetUserValidationResponse_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{3, 0}
}

type LoginResponse_ErrorCode int32

const (
	LoginResponse_BAD_REQUEST                    LoginResponse_ErrorCode = 0
	LoginResponse_INTERNAL_ERROR                 LoginResponse_ErrorCode = 1
	LoginResponse_INCORRECT_USERNAME_OR_PASSWORD LoginResponse_ErrorCode = 2
	LoginResponse_BLANK_USERNAME                 LoginResponse_ErrorCode = 3
	LoginResponse_BLANK_PASSWORD                 LoginResponse_ErrorCode = 4
)

var LoginResponse_ErrorCode_name = map[int32]string{
	0: "BAD_REQUEST",
	1: "INTERNAL_ERROR",
	2: "INCORRECT_USERNAME_OR_PASSWORD",
	3: "BLANK_USERNAME",
	4: "BLANK_PASSWORD",
}

var LoginResponse_ErrorCode_value = map[string]int32{
	"BAD_REQUEST":                    0,
	"INTERNAL_ERROR":                 1,
	"INCORRECT_USERNAME_OR_PASSWORD": 2,
	"BLANK_USERNAME":                 3,
	"BLANK_PASSWORD":                 4,
}

func (x LoginResponse_ErrorCode) String() string {
	return proto.EnumName(LoginResponse_ErrorCode_name, int32(x))
}

func (LoginResponse_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{5, 0}
}

type CreateUserRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{0}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateUserResponse struct {
	Api                  string                       `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Success              bool                         `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	ErrorCode            CreateUserResponse_ErrorCode `protobuf:"varint,3,opt,name=error_code,json=errorCode,proto3,enum=auth.v1.CreateUserResponse_ErrorCode" json:"error_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{1}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateUserResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CreateUserResponse) GetErrorCode() CreateUserResponse_ErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return CreateUserResponse_BAD_REQUEST
}

type SetUserValidationRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	IsValidated          bool     `protobuf:"varint,3,opt,name=is_validated,json=isValidated,proto3" json:"is_validated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetUserValidationRequest) Reset()         { *m = SetUserValidationRequest{} }
func (m *SetUserValidationRequest) String() string { return proto.CompactTextString(m) }
func (*SetUserValidationRequest) ProtoMessage()    {}
func (*SetUserValidationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{2}
}

func (m *SetUserValidationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetUserValidationRequest.Unmarshal(m, b)
}
func (m *SetUserValidationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetUserValidationRequest.Marshal(b, m, deterministic)
}
func (m *SetUserValidationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetUserValidationRequest.Merge(m, src)
}
func (m *SetUserValidationRequest) XXX_Size() int {
	return xxx_messageInfo_SetUserValidationRequest.Size(m)
}
func (m *SetUserValidationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetUserValidationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetUserValidationRequest proto.InternalMessageInfo

func (m *SetUserValidationRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *SetUserValidationRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SetUserValidationRequest) GetIsValidated() bool {
	if m != nil {
		return m.IsValidated
	}
	return false
}

type SetUserValidationResponse struct {
	Api                  string                              `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Success              bool                                `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	ErrorCode            SetUserValidationResponse_ErrorCode `protobuf:"varint,3,opt,name=error_code,json=errorCode,proto3,enum=auth.v1.SetUserValidationResponse_ErrorCode" json:"error_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *SetUserValidationResponse) Reset()         { *m = SetUserValidationResponse{} }
func (m *SetUserValidationResponse) String() string { return proto.CompactTextString(m) }
func (*SetUserValidationResponse) ProtoMessage()    {}
func (*SetUserValidationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{3}
}

func (m *SetUserValidationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetUserValidationResponse.Unmarshal(m, b)
}
func (m *SetUserValidationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetUserValidationResponse.Marshal(b, m, deterministic)
}
func (m *SetUserValidationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetUserValidationResponse.Merge(m, src)
}
func (m *SetUserValidationResponse) XXX_Size() int {
	return xxx_messageInfo_SetUserValidationResponse.Size(m)
}
func (m *SetUserValidationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetUserValidationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetUserValidationResponse proto.InternalMessageInfo

func (m *SetUserValidationResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *SetUserValidationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SetUserValidationResponse) GetErrorCode() SetUserValidationResponse_ErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return SetUserValidationResponse_BAD_REQUEST
}

type LoginRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{4}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Api                  string                  `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Success              bool                    `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Token                string                  `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	ErrorCode            LoginResponse_ErrorCode `protobuf:"varint,4,opt,name=error_code,json=errorCode,proto3,enum=auth.v1.LoginResponse_ErrorCode" json:"error_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{5}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *LoginResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginResponse) GetErrorCode() LoginResponse_ErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return LoginResponse_BAD_REQUEST
}

func init() {
	proto.RegisterEnum("auth.v1.CreateUserResponse_ErrorCode", CreateUserResponse_ErrorCode_name, CreateUserResponse_ErrorCode_value)
	proto.RegisterEnum("auth.v1.SetUserValidationResponse_ErrorCode", SetUserValidationResponse_ErrorCode_name, SetUserValidationResponse_ErrorCode_value)
	proto.RegisterEnum("auth.v1.LoginResponse_ErrorCode", LoginResponse_ErrorCode_name, LoginResponse_ErrorCode_value)
	proto.RegisterType((*CreateUserRequest)(nil), "auth.v1.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "auth.v1.CreateUserResponse")
	proto.RegisterType((*SetUserValidationRequest)(nil), "auth.v1.SetUserValidationRequest")
	proto.RegisterType((*SetUserValidationResponse)(nil), "auth.v1.SetUserValidationResponse")
	proto.RegisterType((*LoginRequest)(nil), "auth.v1.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "auth.v1.LoginResponse")
}

func init() { proto.RegisterFile("auth_service.proto", fileDescriptor_0f39bb026ca10b68) }

var fileDescriptor_0f39bb026ca10b68 = []byte{
	// 777 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x72, 0xe3, 0x44,
	0x10, 0x46, 0x92, 0xf3, 0xd7, 0x9b, 0x04, 0xed, 0x24, 0x04, 0xc7, 0x50, 0x94, 0x50, 0x15, 0x55,
	0x94, 0x6b, 0x23, 0x6d, 0x42, 0x4e, 0xe6, 0x00, 0x8a, 0xad, 0x83, 0x2b, 0x5e, 0x99, 0x1d, 0xdb,
	0x61, 0xd9, 0x8b, 0x4a, 0x2b, 0x37, 0xb6, 0x58, 0x47, 0x23, 0x34, 0x23, 0x87, 0xa2, 0x38, 0xf1,
	0x04, 0x14, 0xdc, 0x78, 0x01, 0xae, 0x5c, 0x78, 0x12, 0x6e, 0x1c, 0x29, 0xde, 0x82, 0x0b, 0x35,
	0xe3, 0xdf, 0x4d, 0xb4, 0x54, 0x6d, 0x4e, 0x9e, 0xee, 0xfe, 0xba, 0x3f, 0xeb, 0xfb, 0x7a, 0x06,
	0x48, 0x54, 0x88, 0x71, 0xc8, 0x31, 0x9f, 0x26, 0x31, 0x3a, 0x59, 0xce, 0x04, 0x23, 0x5b, 0x32,
	0xe7, 0x4c, 0x4f, 0x6b, 0xef, 0x8f, 0x18, 0x1b, 0x4d, 0xd0, 0x8d, 0xb2, 0xc4, 0x8d, 0xd2, 0x94,
	0x89, 0x48, 0x24, 0x2c, 0xe5, 0x33, 0x58, 0xed, 0x91, 0xfa, 0x89, 0x4f, 0x46, 0x98, 0x9e, 0xf0,
	0x9b, 0x68, 0x34, 0xc2, 0xdc, 0x65, 0x99, 0x42, 0xdc, 0x45, 0xdb, 0x1c, 0x1e, 0x36, 0x73, 0x8c,
	0x04, 0x0e, 0x38, 0xe6, 0x14, 0xbf, 0x2d, 0x90, 0x0b, 0x62, 0x82, 0x11, 0x65, 0x49, 0x55, 0xb3,
	0xb4, 0x8f, 0x77, 0xa8, 0x3c, 0x92, 0x43, 0xd8, 0xc0, 0xeb, 0x28, 0x99, 0x54, 0x75, 0x95, 0x9b,
	0x05, 0xa4, 0x06, 0xdb, 0x05, 0xc7, 0x3c, 0x8d, 0xae, 0xb1, 0x6a, 0xa8, 0xc2, 0x32, 0x96, 0xb5,
	0x2c, 0xe2, 0xfc, 0x86, 0xe5, 0xc3, 0x6a, 0x65, 0x56, 0x5b, 0xc4, 0xf6, 0xef, 0x3a, 0x90, 0x75,
	0x56, 0x9e, 0xb1, 0x94, 0x63, 0x09, 0x6d, 0x15, 0xb6, 0x78, 0x11, 0xc7, 0xc8, 0xb9, 0x22, 0xde,
	0xa6, 0x8b, 0x90, 0xb4, 0x00, 0x30, 0xcf, 0x59, 0x1e, 0xc6, 0x6c, 0x38, 0x23, 0xdf, 0x3f, 0xfb,
	0xc8, 0x99, 0x2b, 0xe4, 0xdc, 0x1d, 0xee, 0xf8, 0x12, 0xdd, 0x64, 0x43, 0xa4, 0x3b, 0xb8, 0x38,
	0xda, 0xbf, 0x69, 0xb0, 0xb3, 0x2c, 0x90, 0xb7, 0xe1, 0xc1, 0x85, 0xd7, 0x0a, 0xa9, 0xff, 0x74,
	0xe0, 0xf7, 0xfa, 0xe6, 0x5b, 0x84, 0xc0, 0x7e, 0x3b, 0xe8, 0xfb, 0x34, 0xf0, 0x3a, 0xa1, 0x4f,
	0x69, 0x97, 0x9a, 0x1a, 0x79, 0x08, 0x7b, 0xed, 0xe0, 0xca, 0xeb, 0xb4, 0x5b, 0xa1, 0xff, 0xc4,
	0x6b, 0x77, 0x4c, 0x9d, 0x1c, 0x82, 0xb9, 0x48, 0x0d, 0x7a, 0x12, 0xfd, 0xc4, 0x37, 0x0d, 0xd9,
	0xbc, 0x88, 0xc2, 0xbe, 0x77, 0xe9, 0x07, 0x66, 0x85, 0x98, 0xb0, 0xab, 0x9a, 0xc2, 0x76, 0x20,
	0xa1, 0xe6, 0x86, 0x44, 0x5d, 0x74, 0xbc, 0xe0, 0x72, 0xd5, 0xb9, 0xb9, 0xca, 0x7d, 0xe1, 0xf5,
	0x7a, 0x5f, 0x76, 0x69, 0xcb, 0xdc, 0xb2, 0x5f, 0x42, 0xb5, 0x87, 0x42, 0x7e, 0xd1, 0x55, 0x34,
	0x49, 0x86, 0xca, 0xc3, 0xd7, 0xdb, 0xb5, 0x6e, 0x8c, 0x7e, 0xcb, 0x98, 0x0f, 0x61, 0x37, 0xe1,
	0xe1, 0x74, 0x36, 0x05, 0x87, 0x4a, 0xbb, 0x6d, 0xfa, 0x20, 0xe1, 0x57, 0x8b, 0x94, 0xfd, 0xaf,
	0x06, 0xc7, 0x25, 0x6c, 0xf7, 0xb0, 0xe9, 0xb2, 0xc4, 0xa6, 0x47, 0x4b, 0x9b, 0x5e, 0xcb, 0x51,
	0xee, 0x16, 0xbe, 0xb1, 0x59, 0xef, 0xc2, 0x81, 0xd4, 0x35, 0x6c, 0x75, 0xfd, 0x5e, 0x18, 0x74,
	0xfb, 0xa1, 0xff, 0xac, 0xdd, 0xeb, 0x9b, 0x3a, 0xa9, 0xc1, 0x91, 0x2a, 0x78, 0x1d, 0xea, 0x7b,
	0xad, 0xaf, 0x42, 0xe5, 0x9e, 0xd7, 0xf7, 0x5b, 0xa6, 0x61, 0x3f, 0x83, 0xdd, 0x0e, 0x1b, 0x25,
	0xf7, 0x94, 0x77, 0x7d, 0xef, 0x8d, 0x5b, 0x7b, 0xff, 0x93, 0x0e, 0x7b, 0xf3, 0xd1, 0xf7, 0xd0,
	0xf2, 0x10, 0x36, 0x04, 0x7b, 0x89, 0xe9, 0x7c, 0xec, 0x2c, 0x20, 0x9f, 0xbd, 0xa2, 0x70, 0x45,
	0x29, 0x6c, 0x2d, 0x15, 0x7e, 0x85, 0xad, 0x5c, 0xd5, 0x1f, 0xde, 0x58, 0x55, 0x1b, 0x3e, 0x68,
	0x07, 0xcd, 0x2e, 0xa5, 0x7e, 0xb3, 0xbf, 0xdc, 0xdb, 0xb0, 0x4b, 0x57, 0xfb, 0xaa, 0x97, 0xec,
	0xb5, 0x51, 0xb2, 0xd7, 0x95, 0xb3, 0x3f, 0x34, 0xa8, 0x78, 0x85, 0x18, 0x93, 0xe7, 0x00, 0xab,
	0x5b, 0x4b, 0x6a, 0xa5, 0x57, 0x59, 0xf9, 0x51, 0x7b, 0xef, 0x7f, 0xae, 0xb9, 0x7d, 0xf0, 0xe3,
	0x9f, 0xff, 0xfc, 0xa2, 0xef, 0xd9, 0xdb, 0xee, 0xf4, 0xd4, 0x95, 0xb8, 0x86, 0x56, 0x27, 0x4f,
	0x61, 0x43, 0x09, 0x41, 0xde, 0xb9, 0x2d, 0xcc, 0x6c, 0xe2, 0x51, 0xb9, 0x5e, 0xf6, 0xb1, 0x1a,
	0x76, 0x60, 0xef, 0x2f, 0x86, 0xb9, 0x13, 0x59, 0x6f, 0x68, 0xf5, 0x8b, 0xbf, 0xb5, 0x9f, 0xbd,
	0xbf, 0x34, 0x22, 0xe0, 0x48, 0xfe, 0x7b, 0x4c, 0x45, 0x12, 0xab, 0x0d, 0xb6, 0xe6, 0x6f, 0xb6,
	0x3d, 0x80, 0x5d, 0xd9, 0x71, 0x32, 0x8f, 0x49, 0x7d, 0x2c, 0x44, 0xc6, 0x1b, 0xae, 0x3b, 0x4a,
	0xc4, 0xb8, 0x78, 0xe1, 0xc4, 0xec, 0xda, 0xfd, 0x1e, 0x73, 0x76, 0xf2, 0x75, 0xce, 0xb8, 0x70,
	0xd7, 0xb1, 0xb5, 0x63, 0x59, 0x08, 0x55, 0xe1, 0x73, 0xf5, 0x4c, 0xa7, 0xf2, 0xb9, 0x95, 0x2d,
	0x67, 0xc6, 0xa9, 0xf3, 0xb8, 0xae, 0x69, 0x67, 0x66, 0x94, 0x65, 0x93, 0x39, 0xa7, 0xfb, 0x0d,
	0x67, 0x69, 0xe3, 0x4e, 0x86, 0x7e, 0x0a, 0xc6, 0xf9, 0xe3, 0x73, 0x72, 0x0e, 0x75, 0x8a, 0xa2,
	0xc8, 0x53, 0x1c, 0x5a, 0x37, 0x63, 0x4c, 0x2d, 0x31, 0x46, 0x2b, 0x47, 0xce, 0x8a, 0x3c, 0x46,
	0x6b, 0xc8, 0x90, 0x5b, 0x29, 0x13, 0x16, 0x7e, 0x97, 0x70, 0xe1, 0x90, 0x4d, 0xa8, 0xfc, 0xaa,
	0x6b, 0x5b, 0xcf, 0xf5, 0xe9, 0xe9, 0x8b, 0x4d, 0x45, 0xfe, 0xc9, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x3e, 0x15, 0x19, 0x16, 0x8e, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/auth.v1.Auth/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth.v1.Auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedAuthServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.v1.Auth/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.v1.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.v1.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Auth_CreateUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_service.proto",
}
