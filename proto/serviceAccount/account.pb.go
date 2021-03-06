// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/account.proto

package accountpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Account struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Account) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Account) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

type CreateAccountReq struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountReq) Reset()         { *m = CreateAccountReq{} }
func (m *CreateAccountReq) String() string { return proto.CompactTextString(m) }
func (*CreateAccountReq) ProtoMessage()    {}
func (*CreateAccountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{1}
}

func (m *CreateAccountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountReq.Unmarshal(m, b)
}
func (m *CreateAccountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountReq.Marshal(b, m, deterministic)
}
func (m *CreateAccountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountReq.Merge(m, src)
}
func (m *CreateAccountReq) XXX_Size() int {
	return xxx_messageInfo_CreateAccountReq.Size(m)
}
func (m *CreateAccountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountReq proto.InternalMessageInfo

func (m *CreateAccountReq) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type CreateAccountRes struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountRes) Reset()         { *m = CreateAccountRes{} }
func (m *CreateAccountRes) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRes) ProtoMessage()    {}
func (*CreateAccountRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{2}
}

func (m *CreateAccountRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountRes.Unmarshal(m, b)
}
func (m *CreateAccountRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountRes.Marshal(b, m, deterministic)
}
func (m *CreateAccountRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountRes.Merge(m, src)
}
func (m *CreateAccountRes) XXX_Size() int {
	return xxx_messageInfo_CreateAccountRes.Size(m)
}
func (m *CreateAccountRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountRes proto.InternalMessageInfo

func (m *CreateAccountRes) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type UpdateAccountReq struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAccountReq) Reset()         { *m = UpdateAccountReq{} }
func (m *UpdateAccountReq) String() string { return proto.CompactTextString(m) }
func (*UpdateAccountReq) ProtoMessage()    {}
func (*UpdateAccountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{3}
}

func (m *UpdateAccountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAccountReq.Unmarshal(m, b)
}
func (m *UpdateAccountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAccountReq.Marshal(b, m, deterministic)
}
func (m *UpdateAccountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAccountReq.Merge(m, src)
}
func (m *UpdateAccountReq) XXX_Size() int {
	return xxx_messageInfo_UpdateAccountReq.Size(m)
}
func (m *UpdateAccountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAccountReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAccountReq proto.InternalMessageInfo

func (m *UpdateAccountReq) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type UpdateAccountRes struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAccountRes) Reset()         { *m = UpdateAccountRes{} }
func (m *UpdateAccountRes) String() string { return proto.CompactTextString(m) }
func (*UpdateAccountRes) ProtoMessage()    {}
func (*UpdateAccountRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{4}
}

func (m *UpdateAccountRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAccountRes.Unmarshal(m, b)
}
func (m *UpdateAccountRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAccountRes.Marshal(b, m, deterministic)
}
func (m *UpdateAccountRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAccountRes.Merge(m, src)
}
func (m *UpdateAccountRes) XXX_Size() int {
	return xxx_messageInfo_UpdateAccountRes.Size(m)
}
func (m *UpdateAccountRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAccountRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAccountRes proto.InternalMessageInfo

func (m *UpdateAccountRes) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type ReadAccountReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAccountReq) Reset()         { *m = ReadAccountReq{} }
func (m *ReadAccountReq) String() string { return proto.CompactTextString(m) }
func (*ReadAccountReq) ProtoMessage()    {}
func (*ReadAccountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{5}
}

func (m *ReadAccountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAccountReq.Unmarshal(m, b)
}
func (m *ReadAccountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAccountReq.Marshal(b, m, deterministic)
}
func (m *ReadAccountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAccountReq.Merge(m, src)
}
func (m *ReadAccountReq) XXX_Size() int {
	return xxx_messageInfo_ReadAccountReq.Size(m)
}
func (m *ReadAccountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAccountReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAccountReq proto.InternalMessageInfo

func (m *ReadAccountReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadAccountRes struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAccountRes) Reset()         { *m = ReadAccountRes{} }
func (m *ReadAccountRes) String() string { return proto.CompactTextString(m) }
func (*ReadAccountRes) ProtoMessage()    {}
func (*ReadAccountRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{6}
}

func (m *ReadAccountRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAccountRes.Unmarshal(m, b)
}
func (m *ReadAccountRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAccountRes.Marshal(b, m, deterministic)
}
func (m *ReadAccountRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAccountRes.Merge(m, src)
}
func (m *ReadAccountRes) XXX_Size() int {
	return xxx_messageInfo_ReadAccountRes.Size(m)
}
func (m *ReadAccountRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAccountRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAccountRes proto.InternalMessageInfo

func (m *ReadAccountRes) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type DeleteAccountReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAccountReq) Reset()         { *m = DeleteAccountReq{} }
func (m *DeleteAccountReq) String() string { return proto.CompactTextString(m) }
func (*DeleteAccountReq) ProtoMessage()    {}
func (*DeleteAccountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{7}
}

func (m *DeleteAccountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAccountReq.Unmarshal(m, b)
}
func (m *DeleteAccountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAccountReq.Marshal(b, m, deterministic)
}
func (m *DeleteAccountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAccountReq.Merge(m, src)
}
func (m *DeleteAccountReq) XXX_Size() int {
	return xxx_messageInfo_DeleteAccountReq.Size(m)
}
func (m *DeleteAccountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAccountReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAccountReq proto.InternalMessageInfo

func (m *DeleteAccountReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteAccountRes struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAccountRes) Reset()         { *m = DeleteAccountRes{} }
func (m *DeleteAccountRes) String() string { return proto.CompactTextString(m) }
func (*DeleteAccountRes) ProtoMessage()    {}
func (*DeleteAccountRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{8}
}

func (m *DeleteAccountRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAccountRes.Unmarshal(m, b)
}
func (m *DeleteAccountRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAccountRes.Marshal(b, m, deterministic)
}
func (m *DeleteAccountRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAccountRes.Merge(m, src)
}
func (m *DeleteAccountRes) XXX_Size() int {
	return xxx_messageInfo_DeleteAccountRes.Size(m)
}
func (m *DeleteAccountRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAccountRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAccountRes proto.InternalMessageInfo

func (m *DeleteAccountRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Account)(nil), "proto.Account")
	proto.RegisterType((*CreateAccountReq)(nil), "proto.CreateAccountReq")
	proto.RegisterType((*CreateAccountRes)(nil), "proto.CreateAccountRes")
	proto.RegisterType((*UpdateAccountReq)(nil), "proto.UpdateAccountReq")
	proto.RegisterType((*UpdateAccountRes)(nil), "proto.UpdateAccountRes")
	proto.RegisterType((*ReadAccountReq)(nil), "proto.ReadAccountReq")
	proto.RegisterType((*ReadAccountRes)(nil), "proto.ReadAccountRes")
	proto.RegisterType((*DeleteAccountReq)(nil), "proto.DeleteAccountReq")
	proto.RegisterType((*DeleteAccountRes)(nil), "proto.DeleteAccountRes")
}

func init() { proto.RegisterFile("proto/account.proto", fileDescriptor_477cbf5ae5b46edf) }

var fileDescriptor_477cbf5ae5b46edf = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x4d, 0xb1, 0x8a, 0x0e, 0x8a, 0x75, 0x8c, 0x91, 0x70, 0xaa, 0x9c, 0x7a, 0x30, 0x35, 0xa9,
	0x37, 0xeb, 0xa5, 0xea, 0xd9, 0x03, 0xc6, 0x8b, 0x17, 0x03, 0xec, 0x24, 0x62, 0x2c, 0x20, 0x0b,
	0x7e, 0x8a, 0xdf, 0x6b, 0x98, 0x2e, 0x4d, 0x77, 0x2d, 0x07, 0xd2, 0x13, 0xcc, 0x7b, 0x6f, 0xe6,
	0xcd, 0xbe, 0x81, 0xf3, 0xa2, 0xcc, 0xab, 0xfc, 0x26, 0x4a, 0x92, 0xbc, 0xce, 0xaa, 0x29, 0x57,
	0xb8, 0xcf, 0x9f, 0xe0, 0x13, 0xec, 0xc5, 0x0a, 0x47, 0x17, 0xac, 0x54, 0x78, 0x83, 0xf1, 0x60,
	0x72, 0x14, 0x5a, 0xa9, 0x40, 0x84, 0x61, 0x16, 0x2d, 0xc9, 0xb3, 0x18, 0xe1, 0x7f, 0xf4, 0xc0,
	0x8e, 0x84, 0x28, 0x49, 0x4a, 0x6f, 0x8f, 0xe1, 0xb6, 0xc4, 0x2b, 0x38, 0x2e, 0x3e, 0xf2, 0x8c,
	0xde, 0xb3, 0x7a, 0x19, 0x53, 0xe9, 0x0d, 0x99, 0x76, 0x18, 0x7b, 0x66, 0x28, 0xb8, 0x87, 0xd1,
	0x63, 0x49, 0x51, 0x45, 0xca, 0x31, 0xa4, 0x6f, 0x9c, 0x80, 0xad, 0xf6, 0x62, 0x67, 0x67, 0xe6,
	0xae, 0xf6, 0x9b, 0xb6, 0x9a, 0x96, 0xde, 0xd2, 0x2d, 0xfb, 0x75, 0xbf, 0x16, 0x62, 0x07, 0x6f,
	0xa3, 0xbb, 0x8f, 0xf7, 0x18, 0xdc, 0x90, 0x22, 0xb1, 0xe1, 0x6c, 0x44, 0x1d, 0xdc, 0x19, 0x8a,
	0x3e, 0xd3, 0x03, 0x18, 0x3d, 0xd1, 0x17, 0x69, 0x2f, 0x33, 0xe7, 0x5f, 0xff, 0xd3, 0xc8, 0xe6,
	0x94, 0xb2, 0x4e, 0x92, 0xe6, 0x94, 0x8d, 0xf0, 0x30, 0x6c, 0xcb, 0xd9, 0xaf, 0x05, 0xae, 0x12,
	0xbe, 0x50, 0xf9, 0x93, 0x26, 0x84, 0x0b, 0x38, 0xd1, 0xc2, 0xc7, 0x4b, 0xb5, 0x8e, 0x79, 0x50,
	0xbf, 0x83, 0x90, 0x38, 0x07, 0x67, 0xe3, 0x8d, 0x78, 0xa1, 0x74, 0x7a, 0x32, 0xfe, 0x56, 0x58,
	0x36, 0xfe, 0xda, 0x01, 0xd6, 0xfe, 0xe6, 0x51, 0xfd, 0x0e, 0x82, 0x47, 0x68, 0x19, 0xac, 0x47,
	0x98, 0xe9, 0xf9, 0x1d, 0x84, 0x7c, 0x38, 0x7b, 0x3b, 0x65, 0x66, 0xae, 0xb2, 0x2f, 0xe2, 0xf8,
	0x80, 0x81, 0xdb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x1c, 0x87, 0x7f, 0x64, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountReq, opts ...grpc.CallOption) (*CreateAccountRes, error)
	ReadAccount(ctx context.Context, in *ReadAccountReq, opts ...grpc.CallOption) (*ReadAccountRes, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountReq, opts ...grpc.CallOption) (*UpdateAccountRes, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountReq, opts ...grpc.CallOption) (*DeleteAccountRes, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *CreateAccountReq, opts ...grpc.CallOption) (*CreateAccountRes, error) {
	out := new(CreateAccountRes)
	err := c.cc.Invoke(ctx, "/proto.AccountService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ReadAccount(ctx context.Context, in *ReadAccountReq, opts ...grpc.CallOption) (*ReadAccountRes, error) {
	out := new(ReadAccountRes)
	err := c.cc.Invoke(ctx, "/proto.AccountService/ReadAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateAccount(ctx context.Context, in *UpdateAccountReq, opts ...grpc.CallOption) (*UpdateAccountRes, error) {
	out := new(UpdateAccountRes)
	err := c.cc.Invoke(ctx, "/proto.AccountService/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteAccount(ctx context.Context, in *DeleteAccountReq, opts ...grpc.CallOption) (*DeleteAccountRes, error) {
	out := new(DeleteAccountRes)
	err := c.cc.Invoke(ctx, "/proto.AccountService/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	CreateAccount(context.Context, *CreateAccountReq) (*CreateAccountRes, error)
	ReadAccount(context.Context, *ReadAccountReq) (*ReadAccountRes, error)
	UpdateAccount(context.Context, *UpdateAccountReq) (*UpdateAccountRes, error)
	DeleteAccount(context.Context, *DeleteAccountReq) (*DeleteAccountRes, error)
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccount(ctx, req.(*CreateAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_ReadAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).ReadAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountService/ReadAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).ReadAccount(ctx, req.(*ReadAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountService/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UpdateAccount(ctx, req.(*UpdateAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountService/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteAccount(ctx, req.(*DeleteAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountService_CreateAccount_Handler,
		},
		{
			MethodName: "ReadAccount",
			Handler:    _AccountService_ReadAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _AccountService_UpdateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AccountService_DeleteAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/account.proto",
}
