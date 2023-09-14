// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/storage/stores/indexshipper/compactor/client/grpc/grpc.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GetDeleteRequestsRequest struct {
}

func (m *GetDeleteRequestsRequest) Reset()      { *m = GetDeleteRequestsRequest{} }
func (*GetDeleteRequestsRequest) ProtoMessage() {}
func (*GetDeleteRequestsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4e385854eeb93d, []int{0}
}
func (m *GetDeleteRequestsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetDeleteRequestsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetDeleteRequestsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetDeleteRequestsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeleteRequestsRequest.Merge(m, src)
}
func (m *GetDeleteRequestsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetDeleteRequestsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeleteRequestsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeleteRequestsRequest proto.InternalMessageInfo

type GetDeleteRequestsResponse struct {
	DeleteRequests []*DeleteRequest `protobuf:"bytes,1,rep,name=deleteRequests,proto3" json:"deleteRequests,omitempty"`
}

func (m *GetDeleteRequestsResponse) Reset()      { *m = GetDeleteRequestsResponse{} }
func (*GetDeleteRequestsResponse) ProtoMessage() {}
func (*GetDeleteRequestsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4e385854eeb93d, []int{1}
}
func (m *GetDeleteRequestsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetDeleteRequestsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetDeleteRequestsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetDeleteRequestsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeleteRequestsResponse.Merge(m, src)
}
func (m *GetDeleteRequestsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetDeleteRequestsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeleteRequestsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeleteRequestsResponse proto.InternalMessageInfo

func (m *GetDeleteRequestsResponse) GetDeleteRequests() []*DeleteRequest {
	if m != nil {
		return m.DeleteRequests
	}
	return nil
}

type DeleteRequest struct {
	RequestID string `protobuf:"bytes,1,opt,name=requestID,proto3" json:"requestID,omitempty"`
	StartTime int64  `protobuf:"varint,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime   int64  `protobuf:"varint,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Query     string `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	Status    string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt int64  `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (m *DeleteRequest) Reset()      { *m = DeleteRequest{} }
func (*DeleteRequest) ProtoMessage() {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4e385854eeb93d, []int{2}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return m.Size()
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetRequestID() string {
	if m != nil {
		return m.RequestID
	}
	return ""
}

func (m *DeleteRequest) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *DeleteRequest) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *DeleteRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *DeleteRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DeleteRequest) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type GetCacheGenNumbersRequest struct {
}

func (m *GetCacheGenNumbersRequest) Reset()      { *m = GetCacheGenNumbersRequest{} }
func (*GetCacheGenNumbersRequest) ProtoMessage() {}
func (*GetCacheGenNumbersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4e385854eeb93d, []int{3}
}
func (m *GetCacheGenNumbersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetCacheGenNumbersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetCacheGenNumbersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetCacheGenNumbersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCacheGenNumbersRequest.Merge(m, src)
}
func (m *GetCacheGenNumbersRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetCacheGenNumbersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCacheGenNumbersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCacheGenNumbersRequest proto.InternalMessageInfo

type GetCacheGenNumbersResponse struct {
	ResultsCacheGen string `protobuf:"bytes,1,opt,name=resultsCacheGen,proto3" json:"resultsCacheGen,omitempty"`
}

func (m *GetCacheGenNumbersResponse) Reset()      { *m = GetCacheGenNumbersResponse{} }
func (*GetCacheGenNumbersResponse) ProtoMessage() {}
func (*GetCacheGenNumbersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4e385854eeb93d, []int{4}
}
func (m *GetCacheGenNumbersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetCacheGenNumbersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetCacheGenNumbersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetCacheGenNumbersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCacheGenNumbersResponse.Merge(m, src)
}
func (m *GetCacheGenNumbersResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetCacheGenNumbersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCacheGenNumbersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCacheGenNumbersResponse proto.InternalMessageInfo

func (m *GetCacheGenNumbersResponse) GetResultsCacheGen() string {
	if m != nil {
		return m.ResultsCacheGen
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDeleteRequestsRequest)(nil), "grpc.GetDeleteRequestsRequest")
	proto.RegisterType((*GetDeleteRequestsResponse)(nil), "grpc.GetDeleteRequestsResponse")
	proto.RegisterType((*DeleteRequest)(nil), "grpc.DeleteRequest")
	proto.RegisterType((*GetCacheGenNumbersRequest)(nil), "grpc.GetCacheGenNumbersRequest")
	proto.RegisterType((*GetCacheGenNumbersResponse)(nil), "grpc.GetCacheGenNumbersResponse")
}

func init() {
	proto.RegisterFile("pkg/storage/stores/indexshipper/compactor/client/grpc/grpc.proto", fileDescriptor_af4e385854eeb93d)
}

var fileDescriptor_af4e385854eeb93d = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xbb, 0x4e, 0xe3, 0x40,
	0x14, 0xf5, 0x6c, 0x1e, 0xab, 0xcc, 0x6a, 0x77, 0xc5, 0x80, 0xd0, 0x60, 0xd0, 0x60, 0xb9, 0x72,
	0x15, 0x4b, 0x81, 0x8e, 0x06, 0x48, 0x44, 0x44, 0x43, 0x61, 0x45, 0x82, 0xd6, 0xb1, 0xaf, 0x12,
	0x8b, 0xc4, 0x76, 0x66, 0xc6, 0x12, 0x74, 0x7c, 0x02, 0x9f, 0x41, 0x4b, 0xcb, 0x17, 0x50, 0xa6,
	0x4c, 0x49, 0x9c, 0x86, 0x32, 0x9f, 0x80, 0x32, 0x76, 0x12, 0xe5, 0xd5, 0x78, 0xe6, 0x9e, 0x73,
	0xe6, 0x3e, 0xce, 0x35, 0xbe, 0x8c, 0x1f, 0x3b, 0xb6, 0x90, 0x11, 0x77, 0x3b, 0xa0, 0x4e, 0x10,
	0x76, 0x10, 0xfa, 0xf0, 0x24, 0xba, 0x41, 0x1c, 0x03, 0xb7, 0xbd, 0xa8, 0x1f, 0xbb, 0x9e, 0x8c,
	0xb8, 0xed, 0xf5, 0x02, 0x08, 0xa5, 0xdd, 0xe1, 0xb1, 0xa7, 0x3e, 0xd5, 0x98, 0x47, 0x32, 0x22,
	0xc5, 0xd9, 0xdd, 0xd4, 0x31, 0x6d, 0x82, 0x6c, 0x40, 0x0f, 0x24, 0x38, 0x30, 0x48, 0x40, 0x48,
	0x91, 0x9f, 0xe6, 0x03, 0x3e, 0xda, 0xc2, 0x89, 0x38, 0x0a, 0x05, 0x90, 0x0b, 0xfc, 0xcf, 0x5f,
	0x61, 0x28, 0x32, 0x0a, 0xd6, 0x9f, 0xda, 0x7e, 0x55, 0xd5, 0x58, 0x79, 0xe5, 0xac, 0x49, 0xcd,
	0x77, 0x84, 0xff, 0xae, 0x28, 0xc8, 0x09, 0xae, 0xf0, 0xec, 0x7a, 0xdb, 0xa0, 0xc8, 0x40, 0x56,
	0xc5, 0x59, 0x02, 0x33, 0x56, 0x48, 0x97, 0xcb, 0x56, 0xd0, 0x07, 0xfa, 0xcb, 0x40, 0x56, 0xc1,
	0x59, 0x02, 0x84, 0xe2, 0xdf, 0x10, 0xfa, 0x8a, 0x2b, 0x28, 0x6e, 0x1e, 0x92, 0x03, 0x5c, 0x1a,
	0x24, 0xc0, 0x9f, 0x69, 0x51, 0x65, 0xcc, 0x02, 0x72, 0x88, 0xcb, 0x42, 0xba, 0x32, 0x11, 0xb4,
	0xa4, 0xe0, 0x3c, 0x9a, 0x55, 0xf1, 0x38, 0xb8, 0x12, 0xfc, 0x2b, 0x49, 0xcb, 0x59, 0x95, 0x05,
	0x60, 0x1e, 0x2b, 0x37, 0xea, 0xae, 0xd7, 0x85, 0x26, 0x84, 0x77, 0x49, 0xbf, 0x0d, 0x7c, 0x61,
	0xd5, 0x0d, 0xd6, 0xb7, 0x91, 0xb9, 0x57, 0x16, 0xfe, 0xcf, 0x41, 0x24, 0x3d, 0x29, 0xe6, 0x8a,
	0x7c, 0xc4, 0x75, 0xb8, 0xf6, 0x81, 0x70, 0xa5, 0x3e, 0xdf, 0x1c, 0x69, 0xe1, 0xbd, 0x8d, 0x05,
	0x10, 0x96, 0x19, 0xbc, 0x6b, 0x6b, 0xfa, 0xe9, 0x4e, 0x3e, 0xef, 0xe6, 0x1e, 0x93, 0xcd, 0x5e,
	0xc9, 0xf2, 0xd9, 0xf6, 0x11, 0x75, 0x63, 0xb7, 0x20, 0x4b, 0x7c, 0x7d, 0x3e, 0x1c, 0x33, 0x6d,
	0x34, 0x66, 0xda, 0x74, 0xcc, 0xd0, 0x4b, 0xca, 0xd0, 0x5b, 0xca, 0xd0, 0x67, 0xca, 0xd0, 0x30,
	0x65, 0xe8, 0x2b, 0x65, 0xe8, 0x3b, 0x65, 0xda, 0x34, 0x65, 0xe8, 0x75, 0xc2, 0xb4, 0xe1, 0x84,
	0x69, 0xa3, 0x09, 0xd3, 0xda, 0x65, 0xf5, 0x3b, 0x9e, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x37,
	0x5f, 0x43, 0x54, 0xd2, 0x02, 0x00, 0x00,
}

func (this *GetDeleteRequestsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetDeleteRequestsRequest)
	if !ok {
		that2, ok := that.(GetDeleteRequestsRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *GetDeleteRequestsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetDeleteRequestsResponse)
	if !ok {
		that2, ok := that.(GetDeleteRequestsResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.DeleteRequests) != len(that1.DeleteRequests) {
		return false
	}
	for i := range this.DeleteRequests {
		if !this.DeleteRequests[i].Equal(that1.DeleteRequests[i]) {
			return false
		}
	}
	return true
}
func (this *DeleteRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DeleteRequest)
	if !ok {
		that2, ok := that.(DeleteRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.RequestID != that1.RequestID {
		return false
	}
	if this.StartTime != that1.StartTime {
		return false
	}
	if this.EndTime != that1.EndTime {
		return false
	}
	if this.Query != that1.Query {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if this.CreatedAt != that1.CreatedAt {
		return false
	}
	return true
}
func (this *GetCacheGenNumbersRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetCacheGenNumbersRequest)
	if !ok {
		that2, ok := that.(GetCacheGenNumbersRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *GetCacheGenNumbersResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetCacheGenNumbersResponse)
	if !ok {
		that2, ok := that.(GetCacheGenNumbersResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ResultsCacheGen != that1.ResultsCacheGen {
		return false
	}
	return true
}
func (this *GetDeleteRequestsRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&grpc.GetDeleteRequestsRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetDeleteRequestsResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&grpc.GetDeleteRequestsResponse{")
	if this.DeleteRequests != nil {
		s = append(s, "DeleteRequests: "+fmt.Sprintf("%#v", this.DeleteRequests)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DeleteRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&grpc.DeleteRequest{")
	s = append(s, "RequestID: "+fmt.Sprintf("%#v", this.RequestID)+",\n")
	s = append(s, "StartTime: "+fmt.Sprintf("%#v", this.StartTime)+",\n")
	s = append(s, "EndTime: "+fmt.Sprintf("%#v", this.EndTime)+",\n")
	s = append(s, "Query: "+fmt.Sprintf("%#v", this.Query)+",\n")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "CreatedAt: "+fmt.Sprintf("%#v", this.CreatedAt)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetCacheGenNumbersRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&grpc.GetCacheGenNumbersRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetCacheGenNumbersResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&grpc.GetCacheGenNumbersResponse{")
	s = append(s, "ResultsCacheGen: "+fmt.Sprintf("%#v", this.ResultsCacheGen)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGrpc(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CompactorClient is the client API for Compactor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompactorClient interface {
	GetDeleteRequests(ctx context.Context, in *GetDeleteRequestsRequest, opts ...grpc.CallOption) (*GetDeleteRequestsResponse, error)
	GetCacheGenNumbers(ctx context.Context, in *GetCacheGenNumbersRequest, opts ...grpc.CallOption) (*GetCacheGenNumbersResponse, error)
}

type compactorClient struct {
	cc *grpc.ClientConn
}

func NewCompactorClient(cc *grpc.ClientConn) CompactorClient {
	return &compactorClient{cc}
}

func (c *compactorClient) GetDeleteRequests(ctx context.Context, in *GetDeleteRequestsRequest, opts ...grpc.CallOption) (*GetDeleteRequestsResponse, error) {
	out := new(GetDeleteRequestsResponse)
	err := c.cc.Invoke(ctx, "/grpc.Compactor/GetDeleteRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactorClient) GetCacheGenNumbers(ctx context.Context, in *GetCacheGenNumbersRequest, opts ...grpc.CallOption) (*GetCacheGenNumbersResponse, error) {
	out := new(GetCacheGenNumbersResponse)
	err := c.cc.Invoke(ctx, "/grpc.Compactor/GetCacheGenNumbers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompactorServer is the server API for Compactor service.
type CompactorServer interface {
	GetDeleteRequests(context.Context, *GetDeleteRequestsRequest) (*GetDeleteRequestsResponse, error)
	GetCacheGenNumbers(context.Context, *GetCacheGenNumbersRequest) (*GetCacheGenNumbersResponse, error)
}

// UnimplementedCompactorServer can be embedded to have forward compatible implementations.
type UnimplementedCompactorServer struct {
}

func (*UnimplementedCompactorServer) GetDeleteRequests(ctx context.Context, req *GetDeleteRequestsRequest) (*GetDeleteRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeleteRequests not implemented")
}
func (*UnimplementedCompactorServer) GetCacheGenNumbers(ctx context.Context, req *GetCacheGenNumbersRequest) (*GetCacheGenNumbersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCacheGenNumbers not implemented")
}

func RegisterCompactorServer(s *grpc.Server, srv CompactorServer) {
	s.RegisterService(&_Compactor_serviceDesc, srv)
}

func _Compactor_GetDeleteRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeleteRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactorServer).GetDeleteRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Compactor/GetDeleteRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactorServer).GetDeleteRequests(ctx, req.(*GetDeleteRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Compactor_GetCacheGenNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCacheGenNumbersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactorServer).GetCacheGenNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Compactor/GetCacheGenNumbers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactorServer).GetCacheGenNumbers(ctx, req.(*GetCacheGenNumbersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Compactor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Compactor",
	HandlerType: (*CompactorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeleteRequests",
			Handler:    _Compactor_GetDeleteRequests_Handler,
		},
		{
			MethodName: "GetCacheGenNumbers",
			Handler:    _Compactor_GetCacheGenNumbers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/storage/stores/indexshipper/compactor/client/grpc/grpc.proto",
}

func (m *GetDeleteRequestsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetDeleteRequestsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetDeleteRequestsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetDeleteRequestsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetDeleteRequestsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetDeleteRequestsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DeleteRequests) > 0 {
		for iNdEx := len(m.DeleteRequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DeleteRequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGrpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DeleteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeleteRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeleteRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAt != 0 {
		i = encodeVarintGrpc(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintGrpc(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Query) > 0 {
		i -= len(m.Query)
		copy(dAtA[i:], m.Query)
		i = encodeVarintGrpc(dAtA, i, uint64(len(m.Query)))
		i--
		dAtA[i] = 0x22
	}
	if m.EndTime != 0 {
		i = encodeVarintGrpc(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x18
	}
	if m.StartTime != 0 {
		i = encodeVarintGrpc(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RequestID) > 0 {
		i -= len(m.RequestID)
		copy(dAtA[i:], m.RequestID)
		i = encodeVarintGrpc(dAtA, i, uint64(len(m.RequestID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetCacheGenNumbersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetCacheGenNumbersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetCacheGenNumbersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetCacheGenNumbersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetCacheGenNumbersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetCacheGenNumbersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ResultsCacheGen) > 0 {
		i -= len(m.ResultsCacheGen)
		copy(dAtA[i:], m.ResultsCacheGen)
		i = encodeVarintGrpc(dAtA, i, uint64(len(m.ResultsCacheGen)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGrpc(dAtA []byte, offset int, v uint64) int {
	offset -= sovGrpc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetDeleteRequestsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetDeleteRequestsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DeleteRequests) > 0 {
		for _, e := range m.DeleteRequests {
			l = e.Size()
			n += 1 + l + sovGrpc(uint64(l))
		}
	}
	return n
}

func (m *DeleteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequestID)
	if l > 0 {
		n += 1 + l + sovGrpc(uint64(l))
	}
	if m.StartTime != 0 {
		n += 1 + sovGrpc(uint64(m.StartTime))
	}
	if m.EndTime != 0 {
		n += 1 + sovGrpc(uint64(m.EndTime))
	}
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovGrpc(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovGrpc(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovGrpc(uint64(m.CreatedAt))
	}
	return n
}

func (m *GetCacheGenNumbersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetCacheGenNumbersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ResultsCacheGen)
	if l > 0 {
		n += 1 + l + sovGrpc(uint64(l))
	}
	return n
}

func sovGrpc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGrpc(x uint64) (n int) {
	return sovGrpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetDeleteRequestsRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetDeleteRequestsRequest{`,
		`}`,
	}, "")
	return s
}
func (this *GetDeleteRequestsResponse) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForDeleteRequests := "[]*DeleteRequest{"
	for _, f := range this.DeleteRequests {
		repeatedStringForDeleteRequests += strings.Replace(f.String(), "DeleteRequest", "DeleteRequest", 1) + ","
	}
	repeatedStringForDeleteRequests += "}"
	s := strings.Join([]string{`&GetDeleteRequestsResponse{`,
		`DeleteRequests:` + repeatedStringForDeleteRequests + `,`,
		`}`,
	}, "")
	return s
}
func (this *DeleteRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DeleteRequest{`,
		`RequestID:` + fmt.Sprintf("%v", this.RequestID) + `,`,
		`StartTime:` + fmt.Sprintf("%v", this.StartTime) + `,`,
		`EndTime:` + fmt.Sprintf("%v", this.EndTime) + `,`,
		`Query:` + fmt.Sprintf("%v", this.Query) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`CreatedAt:` + fmt.Sprintf("%v", this.CreatedAt) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetCacheGenNumbersRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetCacheGenNumbersRequest{`,
		`}`,
	}, "")
	return s
}
func (this *GetCacheGenNumbersResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetCacheGenNumbersResponse{`,
		`ResultsCacheGen:` + fmt.Sprintf("%v", this.ResultsCacheGen) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGrpc(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GetDeleteRequestsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetDeleteRequestsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetDeleteRequestsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetDeleteRequestsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetDeleteRequestsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetDeleteRequestsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeleteRequests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeleteRequests = append(m.DeleteRequests, &DeleteRequest{})
			if err := m.DeleteRequests[len(m.DeleteRequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeleteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeleteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetCacheGenNumbersRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetCacheGenNumbersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCacheGenNumbersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetCacheGenNumbersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetCacheGenNumbersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCacheGenNumbersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultsCacheGen", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResultsCacheGen = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGrpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGrpc
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGrpc
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthGrpc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGrpc
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGrpc(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthGrpc
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGrpc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGrpc   = fmt.Errorf("proto: integer overflow")
)
