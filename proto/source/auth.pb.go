// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package hayaku

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type RegisterRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type RegisterResponse struct {
	// Types that are valid to be assigned to Response:
	//	*RegisterResponse_Fail
	//	*RegisterResponse_Success
	Response             isRegisterResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

type isRegisterResponse_Response interface {
	isRegisterResponse_Response()
}

type RegisterResponse_Fail struct {
	Fail *Error `protobuf:"bytes,1,opt,name=fail,proto3,oneof"`
}

type RegisterResponse_Success struct {
	Success *Success `protobuf:"bytes,2,opt,name=success,proto3,oneof"`
}

func (*RegisterResponse_Fail) isRegisterResponse_Response() {}

func (*RegisterResponse_Success) isRegisterResponse_Response() {}

func (m *RegisterResponse) GetResponse() isRegisterResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *RegisterResponse) GetFail() *Error {
	if x, ok := m.GetResponse().(*RegisterResponse_Fail); ok {
		return x.Fail
	}
	return nil
}

func (m *RegisterResponse) GetSuccess() *Success {
	if x, ok := m.GetResponse().(*RegisterResponse_Success); ok {
		return x.Success
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RegisterResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RegisterResponse_Fail)(nil),
		(*RegisterResponse_Success)(nil),
	}
}

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
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
	// Types that are valid to be assigned to Response:
	//	*LoginResponse_Fail
	//	*LoginResponse_Success
	Response             isLoginResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
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

type isLoginResponse_Response interface {
	isLoginResponse_Response()
}

type LoginResponse_Fail struct {
	Fail *Error `protobuf:"bytes,1,opt,name=fail,proto3,oneof"`
}

type LoginResponse_Success struct {
	Success *Success `protobuf:"bytes,2,opt,name=success,proto3,oneof"`
}

func (*LoginResponse_Fail) isLoginResponse_Response() {}

func (*LoginResponse_Success) isLoginResponse_Response() {}

func (m *LoginResponse) GetResponse() isLoginResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *LoginResponse) GetFail() *Error {
	if x, ok := m.GetResponse().(*LoginResponse_Fail); ok {
		return x.Fail
	}
	return nil
}

func (m *LoginResponse) GetSuccess() *Success {
	if x, ok := m.GetResponse().(*LoginResponse_Success); ok {
		return x.Success
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LoginResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LoginResponse_Fail)(nil),
		(*LoginResponse_Success)(nil),
	}
}

type VerifyRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *VerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyRequest.Unmarshal(m, b)
}
func (m *VerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyRequest.Marshal(b, m, deterministic)
}
func (m *VerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRequest.Merge(m, src)
}
func (m *VerifyRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyRequest.Size(m)
}
func (m *VerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRequest proto.InternalMessageInfo

func (m *VerifyRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyResponse struct {
	Allow bool `protobuf:"varint,1,opt,name=allow,proto3" json:"allow,omitempty"`
	// Types that are valid to be assigned to Response:
	//	*VerifyResponse_Fail
	//	*VerifyResponse_Success
	Response             isVerifyResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *VerifyResponse) Reset()         { *m = VerifyResponse{} }
func (m *VerifyResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyResponse) ProtoMessage()    {}
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{5}
}

func (m *VerifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyResponse.Unmarshal(m, b)
}
func (m *VerifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyResponse.Marshal(b, m, deterministic)
}
func (m *VerifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyResponse.Merge(m, src)
}
func (m *VerifyResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyResponse.Size(m)
}
func (m *VerifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyResponse proto.InternalMessageInfo

func (m *VerifyResponse) GetAllow() bool {
	if m != nil {
		return m.Allow
	}
	return false
}

type isVerifyResponse_Response interface {
	isVerifyResponse_Response()
}

type VerifyResponse_Fail struct {
	Fail *Error `protobuf:"bytes,2,opt,name=fail,proto3,oneof"`
}

type VerifyResponse_Success struct {
	Success *Success `protobuf:"bytes,3,opt,name=success,proto3,oneof"`
}

func (*VerifyResponse_Fail) isVerifyResponse_Response() {}

func (*VerifyResponse_Success) isVerifyResponse_Response() {}

func (m *VerifyResponse) GetResponse() isVerifyResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *VerifyResponse) GetFail() *Error {
	if x, ok := m.GetResponse().(*VerifyResponse_Fail); ok {
		return x.Fail
	}
	return nil
}

func (m *VerifyResponse) GetSuccess() *Success {
	if x, ok := m.GetResponse().(*VerifyResponse_Success); ok {
		return x.Success
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*VerifyResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*VerifyResponse_Fail)(nil),
		(*VerifyResponse_Success)(nil),
	}
}

type Error struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code                 int32    `protobuf:"zigzag32,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{6}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type Success struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Success) Reset()         { *m = Success{} }
func (m *Success) String() string { return proto.CompactTextString(m) }
func (*Success) ProtoMessage()    {}
func (*Success) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{7}
}

func (m *Success) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Success.Unmarshal(m, b)
}
func (m *Success) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Success.Marshal(b, m, deterministic)
}
func (m *Success) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Success.Merge(m, src)
}
func (m *Success) XXX_Size() int {
	return xxx_messageInfo_Success.Size(m)
}
func (m *Success) XXX_DiscardUnknown() {
	xxx_messageInfo_Success.DiscardUnknown(m)
}

var xxx_messageInfo_Success proto.InternalMessageInfo

func (m *Success) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Success) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type MetricsRequest struct {
	Params               []string `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsRequest) Reset()         { *m = MetricsRequest{} }
func (m *MetricsRequest) String() string { return proto.CompactTextString(m) }
func (*MetricsRequest) ProtoMessage()    {}
func (*MetricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{8}
}

func (m *MetricsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsRequest.Unmarshal(m, b)
}
func (m *MetricsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsRequest.Marshal(b, m, deterministic)
}
func (m *MetricsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsRequest.Merge(m, src)
}
func (m *MetricsRequest) XXX_Size() int {
	return xxx_messageInfo_MetricsRequest.Size(m)
}
func (m *MetricsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsRequest proto.InternalMessageInfo

func (m *MetricsRequest) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

type MetricsFrame struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsFrame) Reset()         { *m = MetricsFrame{} }
func (m *MetricsFrame) String() string { return proto.CompactTextString(m) }
func (*MetricsFrame) ProtoMessage()    {}
func (*MetricsFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{9}
}

func (m *MetricsFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsFrame.Unmarshal(m, b)
}
func (m *MetricsFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsFrame.Marshal(b, m, deterministic)
}
func (m *MetricsFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsFrame.Merge(m, src)
}
func (m *MetricsFrame) XXX_Size() int {
	return xxx_messageInfo_MetricsFrame.Size(m)
}
func (m *MetricsFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsFrame.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsFrame proto.InternalMessageInfo

func (m *MetricsFrame) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "hayaku.register_request")
	proto.RegisterType((*RegisterResponse)(nil), "hayaku.register_response")
	proto.RegisterType((*LoginRequest)(nil), "hayaku.login_request")
	proto.RegisterType((*LoginResponse)(nil), "hayaku.login_response")
	proto.RegisterType((*VerifyRequest)(nil), "hayaku.verify_request")
	proto.RegisterType((*VerifyResponse)(nil), "hayaku.verify_response")
	proto.RegisterType((*Error)(nil), "hayaku.error")
	proto.RegisterType((*Success)(nil), "hayaku.success")
	proto.RegisterType((*MetricsRequest)(nil), "hayaku.metrics_request")
	proto.RegisterType((*MetricsFrame)(nil), "hayaku.metrics_frame")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xde, 0xa4, 0x4d, 0x36, 0xfb, 0x96, 0x6e, 0x77, 0x87, 0xb6, 0xc6, 0x9c, 0x4a, 0x04, 0x69,
	0x11, 0x8a, 0x54, 0x14, 0x7a, 0x12, 0xa4, 0x60, 0x0f, 0x7a, 0xc9, 0xc1, 0xab, 0x8e, 0xe9, 0x6b,
	0x1b, 0x9a, 0x64, 0xd2, 0x99, 0xc4, 0xd2, 0x9b, 0x17, 0xff, 0x6f, 0xc9, 0x64, 0x66, 0x68, 0xe3,
	0x0f, 0x14, 0xc1, 0x5b, 0xbe, 0xf9, 0xbe, 0x79, 0xef, 0x7b, 0x3f, 0x26, 0x00, 0xb4, 0x2a, 0x77,
	0xb3, 0x82, 0xb3, 0x92, 0x11, 0x77, 0x47, 0x4f, 0x74, 0x5f, 0x85, 0x9f, 0xe0, 0x9e, 0xe3, 0x36,
	0x11, 0x25, 0xf2, 0x8f, 0x1c, 0x0f, 0x15, 0x8a, 0x92, 0x04, 0xe0, 0x55, 0x02, 0x79, 0x4e, 0x33,
	0xf4, 0xad, 0xb1, 0x35, 0xb9, 0x89, 0x0c, 0xae, 0xb9, 0x82, 0x0a, 0x71, 0x64, 0x7c, 0xed, 0xdb,
	0x0d, 0xa7, 0x31, 0x19, 0x80, 0x83, 0x19, 0x4d, 0x52, 0xbf, 0x23, 0x89, 0x06, 0x84, 0x07, 0x78,
	0x38, 0xcb, 0x20, 0x0a, 0x96, 0x0b, 0x24, 0x4f, 0xa0, 0xbb, 0xa9, 0x95, 0x75, 0xf8, 0xdb, 0x79,
	0x6f, 0xd6, 0xb8, 0x99, 0x21, 0xe7, 0x8c, 0xaf, 0xae, 0x22, 0x49, 0x92, 0x67, 0x70, 0x2d, 0xaa,
	0x38, 0x46, 0x21, 0x64, 0xaa, 0xdb, 0x79, 0x5f, 0xeb, 0xd4, 0xf1, 0xea, 0x2a, 0xd2, 0x8a, 0x37,
	0x00, 0x9e, 0x8e, 0x1e, 0xbe, 0x85, 0x5e, 0xca, 0xb6, 0x49, 0xfe, 0xaf, 0x15, 0x85, 0x39, 0xdc,
	0xe9, 0x40, 0xff, 0xc5, 0xf8, 0x53, 0xb8, 0xfb, 0x82, 0x3c, 0xd9, 0x9c, 0x8c, 0xf3, 0x01, 0x38,
	0x25, 0xdb, 0x63, 0xae, 0x6c, 0x37, 0x20, 0xfc, 0x66, 0x41, 0xdf, 0x08, 0x95, 0xb3, 0x01, 0x38,
	0x34, 0x4d, 0xd9, 0x51, 0x2a, 0xbd, 0xa8, 0x01, 0xc6, 0xaf, 0xfd, 0x87, 0x7e, 0x3b, 0x7f, 0xe5,
	0xf7, 0x25, 0x38, 0x32, 0x12, 0xf1, 0xe1, 0x3a, 0x43, 0x21, 0xe8, 0x56, 0xf7, 0x57, 0x43, 0x42,
	0xa0, 0x1b, 0xb3, 0x35, 0x4a, 0x03, 0x0f, 0x91, 0xfc, 0x0e, 0x17, 0x26, 0xdf, 0x6f, 0x2e, 0x9a,
	0xca, 0xed, 0xf3, 0xca, 0xa7, 0xd0, 0xcf, 0xb0, 0xe4, 0x49, 0x2c, 0x4c, 0x8b, 0x46, 0xe0, 0x16,
	0x94, 0xd3, 0x4c, 0xf8, 0xd6, 0xb8, 0x33, 0xb9, 0x89, 0x14, 0x0a, 0xa7, 0xd0, 0xd3, 0xd2, 0x0d,
	0xaf, 0x27, 0xfd, 0xcb, 0x5c, 0xf3, 0xaf, 0x36, 0x74, 0xeb, 0xc7, 0x41, 0x5e, 0x83, 0x17, 0xa9,
	0x65, 0x25, 0xbe, 0x6e, 0x42, 0xfb, 0x81, 0x04, 0x8f, 0x7f, 0xc2, 0xa8, 0x29, 0xbc, 0x02, 0xe7,
	0x5d, 0xbd, 0x31, 0x64, 0xa8, 0x35, 0x17, 0x9b, 0x18, 0x8c, 0xda, 0xc7, 0xea, 0xde, 0x02, 0xdc,
	0x0f, 0x72, 0xa0, 0xc4, 0x28, 0x2e, 0x37, 0x21, 0x78, 0xf4, 0xc3, 0xb9, 0xba, 0xba, 0x84, 0xfb,
	0xf7, 0xb2, 0xce, 0x65, 0x22, 0x0a, 0x5a, 0xc6, 0x3b, 0xe4, 0xc4, 0x88, 0x5b, 0xcd, 0x0a, 0x86,
	0x6d, 0x42, 0xb6, 0xe6, 0xb9, 0xf5, 0xd9, 0x95, 0xff, 0x85, 0x17, 0xdf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x34, 0xe8, 0x32, 0x54, 0x25, 0x04, 0x00, 0x00,
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
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
	MetricDispatcher(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (Auth_MetricDispatcherClient, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/hayaku.auth/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/hayaku.auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/hayaku.auth/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) MetricDispatcher(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (Auth_MetricDispatcherClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Auth_serviceDesc.Streams[0], "/hayaku.auth/MetricDispatcher", opts...)
	if err != nil {
		return nil, err
	}
	x := &authMetricDispatcherClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Auth_MetricDispatcherClient interface {
	Recv() (*MetricsFrame, error)
	grpc.ClientStream
}

type authMetricDispatcherClient struct {
	grpc.ClientStream
}

func (x *authMetricDispatcherClient) Recv() (*MetricsFrame, error) {
	m := new(MetricsFrame)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
	MetricDispatcher(*MetricsRequest, Auth_MetricDispatcherServer) error
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedAuthServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServer) Verify(ctx context.Context, req *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (*UnimplementedAuthServer) MetricDispatcher(req *MetricsRequest, srv Auth_MetricDispatcherServer) error {
	return status.Errorf(codes.Unimplemented, "method MetricDispatcher not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hayaku.auth/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Register(ctx, req.(*RegisterRequest))
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
		FullMethod: "/hayaku.auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hayaku.auth/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_MetricDispatcher_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MetricsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AuthServer).MetricDispatcher(m, &authMetricDispatcherServer{stream})
}

type Auth_MetricDispatcherServer interface {
	Send(*MetricsFrame) error
	grpc.ServerStream
}

type authMetricDispatcherServer struct {
	grpc.ServerStream
}

func (x *authMetricDispatcherServer) Send(m *MetricsFrame) error {
	return x.ServerStream.SendMsg(m)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hayaku.auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Auth_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Auth_Verify_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MetricDispatcher",
			Handler:       _Auth_MetricDispatcher_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "auth.proto",
}
