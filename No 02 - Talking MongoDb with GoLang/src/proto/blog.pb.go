// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog.proto

package blogpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Blog struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId             string   `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{0}
}

func (m *Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog.Unmarshal(m, b)
}
func (m *Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog.Marshal(b, m, deterministic)
}
func (m *Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog.Merge(m, src)
}
func (m *Blog) XXX_Size() int {
	return xxx_messageInfo_Blog.Size(m)
}
func (m *Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog proto.InternalMessageInfo

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Blog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CreateBlogReq struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogReq) Reset()         { *m = CreateBlogReq{} }
func (m *CreateBlogReq) String() string { return proto.CompactTextString(m) }
func (*CreateBlogReq) ProtoMessage()    {}
func (*CreateBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{1}
}

func (m *CreateBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogReq.Unmarshal(m, b)
}
func (m *CreateBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogReq.Marshal(b, m, deterministic)
}
func (m *CreateBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogReq.Merge(m, src)
}
func (m *CreateBlogReq) XXX_Size() int {
	return xxx_messageInfo_CreateBlogReq.Size(m)
}
func (m *CreateBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogReq proto.InternalMessageInfo

func (m *CreateBlogReq) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type CreateBlogRes struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogRes) Reset()         { *m = CreateBlogRes{} }
func (m *CreateBlogRes) String() string { return proto.CompactTextString(m) }
func (*CreateBlogRes) ProtoMessage()    {}
func (*CreateBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{2}
}

func (m *CreateBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogRes.Unmarshal(m, b)
}
func (m *CreateBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogRes.Marshal(b, m, deterministic)
}
func (m *CreateBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogRes.Merge(m, src)
}
func (m *CreateBlogRes) XXX_Size() int {
	return xxx_messageInfo_CreateBlogRes.Size(m)
}
func (m *CreateBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogRes proto.InternalMessageInfo

func (m *CreateBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type ReadBlogReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogReq) Reset()         { *m = ReadBlogReq{} }
func (m *ReadBlogReq) String() string { return proto.CompactTextString(m) }
func (*ReadBlogReq) ProtoMessage()    {}
func (*ReadBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{3}
}

func (m *ReadBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogReq.Unmarshal(m, b)
}
func (m *ReadBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogReq.Marshal(b, m, deterministic)
}
func (m *ReadBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogReq.Merge(m, src)
}
func (m *ReadBlogReq) XXX_Size() int {
	return xxx_messageInfo_ReadBlogReq.Size(m)
}
func (m *ReadBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogReq proto.InternalMessageInfo

func (m *ReadBlogReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadBlogRes struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogRes) Reset()         { *m = ReadBlogRes{} }
func (m *ReadBlogRes) String() string { return proto.CompactTextString(m) }
func (*ReadBlogRes) ProtoMessage()    {}
func (*ReadBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{4}
}

func (m *ReadBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogRes.Unmarshal(m, b)
}
func (m *ReadBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogRes.Marshal(b, m, deterministic)
}
func (m *ReadBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogRes.Merge(m, src)
}
func (m *ReadBlogRes) XXX_Size() int {
	return xxx_messageInfo_ReadBlogRes.Size(m)
}
func (m *ReadBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogRes proto.InternalMessageInfo

func (m *ReadBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogReq struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlogReq) Reset()         { *m = UpdateBlogReq{} }
func (m *UpdateBlogReq) String() string { return proto.CompactTextString(m) }
func (*UpdateBlogReq) ProtoMessage()    {}
func (*UpdateBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{5}
}

func (m *UpdateBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlogReq.Unmarshal(m, b)
}
func (m *UpdateBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlogReq.Marshal(b, m, deterministic)
}
func (m *UpdateBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlogReq.Merge(m, src)
}
func (m *UpdateBlogReq) XXX_Size() int {
	return xxx_messageInfo_UpdateBlogReq.Size(m)
}
func (m *UpdateBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlogReq proto.InternalMessageInfo

func (m *UpdateBlogReq) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogRes struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlogRes) Reset()         { *m = UpdateBlogRes{} }
func (m *UpdateBlogRes) String() string { return proto.CompactTextString(m) }
func (*UpdateBlogRes) ProtoMessage()    {}
func (*UpdateBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{6}
}

func (m *UpdateBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlogRes.Unmarshal(m, b)
}
func (m *UpdateBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlogRes.Marshal(b, m, deterministic)
}
func (m *UpdateBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlogRes.Merge(m, src)
}
func (m *UpdateBlogRes) XXX_Size() int {
	return xxx_messageInfo_UpdateBlogRes.Size(m)
}
func (m *UpdateBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlogRes proto.InternalMessageInfo

func (m *UpdateBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type DeleteBlogReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlogReq) Reset()         { *m = DeleteBlogReq{} }
func (m *DeleteBlogReq) String() string { return proto.CompactTextString(m) }
func (*DeleteBlogReq) ProtoMessage()    {}
func (*DeleteBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{7}
}

func (m *DeleteBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlogReq.Unmarshal(m, b)
}
func (m *DeleteBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlogReq.Marshal(b, m, deterministic)
}
func (m *DeleteBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlogReq.Merge(m, src)
}
func (m *DeleteBlogReq) XXX_Size() int {
	return xxx_messageInfo_DeleteBlogReq.Size(m)
}
func (m *DeleteBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlogReq proto.InternalMessageInfo

func (m *DeleteBlogReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteBlogRes struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlogRes) Reset()         { *m = DeleteBlogRes{} }
func (m *DeleteBlogRes) String() string { return proto.CompactTextString(m) }
func (*DeleteBlogRes) ProtoMessage()    {}
func (*DeleteBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{8}
}

func (m *DeleteBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlogRes.Unmarshal(m, b)
}
func (m *DeleteBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlogRes.Marshal(b, m, deterministic)
}
func (m *DeleteBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlogRes.Merge(m, src)
}
func (m *DeleteBlogRes) XXX_Size() int {
	return xxx_messageInfo_DeleteBlogRes.Size(m)
}
func (m *DeleteBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlogRes proto.InternalMessageInfo

func (m *DeleteBlogRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListBlogsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlogsReq) Reset()         { *m = ListBlogsReq{} }
func (m *ListBlogsReq) String() string { return proto.CompactTextString(m) }
func (*ListBlogsReq) ProtoMessage()    {}
func (*ListBlogsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{9}
}

func (m *ListBlogsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlogsReq.Unmarshal(m, b)
}
func (m *ListBlogsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlogsReq.Marshal(b, m, deterministic)
}
func (m *ListBlogsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlogsReq.Merge(m, src)
}
func (m *ListBlogsReq) XXX_Size() int {
	return xxx_messageInfo_ListBlogsReq.Size(m)
}
func (m *ListBlogsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlogsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlogsReq proto.InternalMessageInfo

type ListBlogsRes struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlogsRes) Reset()         { *m = ListBlogsRes{} }
func (m *ListBlogsRes) String() string { return proto.CompactTextString(m) }
func (*ListBlogsRes) ProtoMessage()    {}
func (*ListBlogsRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{10}
}

func (m *ListBlogsRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlogsRes.Unmarshal(m, b)
}
func (m *ListBlogsRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlogsRes.Marshal(b, m, deterministic)
}
func (m *ListBlogsRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlogsRes.Merge(m, src)
}
func (m *ListBlogsRes) XXX_Size() int {
	return xxx_messageInfo_ListBlogsRes.Size(m)
}
func (m *ListBlogsRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlogsRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlogsRes proto.InternalMessageInfo

func (m *ListBlogsRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

func init() {
	proto.RegisterType((*Blog)(nil), "blog.Blog")
	proto.RegisterType((*CreateBlogReq)(nil), "blog.CreateBlogReq")
	proto.RegisterType((*CreateBlogRes)(nil), "blog.CreateBlogRes")
	proto.RegisterType((*ReadBlogReq)(nil), "blog.ReadBlogReq")
	proto.RegisterType((*ReadBlogRes)(nil), "blog.ReadBlogRes")
	proto.RegisterType((*UpdateBlogReq)(nil), "blog.UpdateBlogReq")
	proto.RegisterType((*UpdateBlogRes)(nil), "blog.UpdateBlogRes")
	proto.RegisterType((*DeleteBlogReq)(nil), "blog.DeleteBlogReq")
	proto.RegisterType((*DeleteBlogRes)(nil), "blog.DeleteBlogRes")
	proto.RegisterType((*ListBlogsReq)(nil), "blog.ListBlogsReq")
	proto.RegisterType((*ListBlogsRes)(nil), "blog.ListBlogsRes")
}

func init() { proto.RegisterFile("blog.proto", fileDescriptor_6745b25902462fb1) }

var fileDescriptor_6745b25902462fb1 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x4f, 0xc2, 0x40,
	0x14, 0x0c, 0x15, 0xb1, 0x3c, 0x84, 0xc4, 0xa7, 0x87, 0xa6, 0xc6, 0x8f, 0xf4, 0xa4, 0x07, 0x91,
	0x60, 0xf4, 0x07, 0xa0, 0x17, 0x13, 0x4f, 0x35, 0x5e, 0xbc, 0x18, 0xe8, 0xbe, 0xe0, 0x26, 0x0d,
	0x0b, 0x7d, 0x8b, 0x3f, 0xc7, 0xdf, 0x6a, 0x76, 0x37, 0x6b, 0x5b, 0x20, 0xa9, 0xde, 0x3a, 0xb3,
	0x33, 0x9d, 0xe9, 0x6c, 0x0a, 0x30, 0xcb, 0xd5, 0x7c, 0xb8, 0x2c, 0x94, 0x56, 0xd8, 0x36, 0xcf,
	0x49, 0x06, 0xed, 0x49, 0xae, 0xe6, 0x38, 0x80, 0x40, 0x8a, 0xa8, 0x75, 0xd9, 0xba, 0xea, 0xa6,
	0x81, 0x14, 0x78, 0x0a, 0xdd, 0xe9, 0x5a, 0x7f, 0xaa, 0xe2, 0x43, 0x8a, 0x28, 0xb0, 0x74, 0xe8,
	0x88, 0x67, 0x81, 0x27, 0xb0, 0xaf, 0xa5, 0xce, 0x29, 0xda, 0xb3, 0x07, 0x0e, 0x60, 0x04, 0x07,
	0x99, 0x5a, 0x68, 0x5a, 0xe8, 0xa8, 0x6d, 0x79, 0x0f, 0x93, 0x5b, 0xe8, 0x3f, 0x16, 0x34, 0xd5,
	0x64, 0xa2, 0x52, 0x5a, 0xe1, 0x39, 0xd8, 0x74, 0x9b, 0xd7, 0x1b, 0xc3, 0xd0, 0xd6, 0xb2, 0x87,
	0xae, 0xd5, 0x86, 0x81, 0x1b, 0x0d, 0x67, 0xd0, 0x4b, 0x69, 0x2a, 0xfc, 0xfb, 0x37, 0xbe, 0x26,
	0xb9, 0xa9, 0x1e, 0xf3, 0x5f, 0xe2, 0xdf, 0x96, 0xe2, 0x7f, 0x7d, 0xab, 0x86, 0xe6, 0x84, 0x0b,
	0xe8, 0x3f, 0x51, 0x4e, 0x65, 0xc2, 0x66, 0xe3, 0xeb, 0xba, 0x80, 0xcd, 0xba, 0xbc, 0xce, 0x32,
	0x62, 0xb6, 0xaa, 0x30, 0xf5, 0x30, 0x19, 0xc0, 0xe1, 0x8b, 0x64, 0x6d, 0x84, 0x9c, 0xd2, 0x2a,
	0x19, 0xd6, 0x70, 0x63, 0x97, 0xf1, 0x77, 0x00, 0x3d, 0x03, 0x5f, 0xa9, 0xf8, 0x92, 0x19, 0xe1,
	0x03, 0x40, 0x39, 0x3e, 0x1e, 0x3b, 0x7d, 0xed, 0xfe, 0xe2, 0x1d, 0x24, 0xe3, 0x08, 0x42, 0x3f,
	0x32, 0x1e, 0x39, 0x41, 0xe5, 0x4e, 0xe2, 0x2d, 0x8a, 0x4d, 0x52, 0x39, 0x9b, 0x4f, 0xaa, 0x2d,
	0x1f, 0xef, 0x20, 0xad, 0xaf, 0x1c, 0xc7, 0xfb, 0x6a, 0x7b, 0xc6, 0x3b, 0x48, 0xc6, 0x7b, 0xe8,
	0xfe, 0x2e, 0x83, 0xe8, 0x14, 0xd5, 0xe9, 0xe2, 0x6d, 0x8e, 0x47, 0xad, 0x49, 0xf8, 0xde, 0x31,
	0xf4, 0x72, 0x36, 0xeb, 0xd8, 0x5f, 0xe7, 0xee, 0x27, 0x00, 0x00, 0xff, 0xff, 0x81, 0xb6, 0x8a,
	0xa3, 0x48, 0x03, 0x00, 0x00,
}
