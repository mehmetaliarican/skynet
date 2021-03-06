// Code generated by protoc-gen-go. DO NOT EDIT.
// source: player.proto

package playerpb

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

// Player mesaj tipinin tanimi
type Player struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PlayerId             string   `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Fullname             string   `protobuf:"bytes,3,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Position             string   `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
	Bio                  string   `protobuf:"bytes,5,opt,name=bio,proto3" json:"bio,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{0}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Player) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *Player) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *Player) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Player) GetBio() string {
	if m != nil {
		return m.Bio
	}
	return ""
}

// Operasyonlarin kullandigi request ve response mesajlarina ait tanimlamalar
type AddPlayerReq struct {
	Plyr                 *Player  `protobuf:"bytes,1,opt,name=plyr,proto3" json:"plyr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPlayerReq) Reset()         { *m = AddPlayerReq{} }
func (m *AddPlayerReq) String() string { return proto.CompactTextString(m) }
func (*AddPlayerReq) ProtoMessage()    {}
func (*AddPlayerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{1}
}

func (m *AddPlayerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPlayerReq.Unmarshal(m, b)
}
func (m *AddPlayerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPlayerReq.Marshal(b, m, deterministic)
}
func (m *AddPlayerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPlayerReq.Merge(m, src)
}
func (m *AddPlayerReq) XXX_Size() int {
	return xxx_messageInfo_AddPlayerReq.Size(m)
}
func (m *AddPlayerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPlayerReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddPlayerReq proto.InternalMessageInfo

func (m *AddPlayerReq) GetPlyr() *Player {
	if m != nil {
		return m.Plyr
	}
	return nil
}

type AddPlayerRes struct {
	Plyr                 *Player  `protobuf:"bytes,1,opt,name=plyr,proto3" json:"plyr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPlayerRes) Reset()         { *m = AddPlayerRes{} }
func (m *AddPlayerRes) String() string { return proto.CompactTextString(m) }
func (*AddPlayerRes) ProtoMessage()    {}
func (*AddPlayerRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{2}
}

func (m *AddPlayerRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPlayerRes.Unmarshal(m, b)
}
func (m *AddPlayerRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPlayerRes.Marshal(b, m, deterministic)
}
func (m *AddPlayerRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPlayerRes.Merge(m, src)
}
func (m *AddPlayerRes) XXX_Size() int {
	return xxx_messageInfo_AddPlayerRes.Size(m)
}
func (m *AddPlayerRes) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPlayerRes.DiscardUnknown(m)
}

var xxx_messageInfo_AddPlayerRes proto.InternalMessageInfo

func (m *AddPlayerRes) GetPlyr() *Player {
	if m != nil {
		return m.Plyr
	}
	return nil
}

type EditPlayerReq struct {
	Plyr                 *Player  `protobuf:"bytes,1,opt,name=plyr,proto3" json:"plyr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditPlayerReq) Reset()         { *m = EditPlayerReq{} }
func (m *EditPlayerReq) String() string { return proto.CompactTextString(m) }
func (*EditPlayerReq) ProtoMessage()    {}
func (*EditPlayerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{3}
}

func (m *EditPlayerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditPlayerReq.Unmarshal(m, b)
}
func (m *EditPlayerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditPlayerReq.Marshal(b, m, deterministic)
}
func (m *EditPlayerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditPlayerReq.Merge(m, src)
}
func (m *EditPlayerReq) XXX_Size() int {
	return xxx_messageInfo_EditPlayerReq.Size(m)
}
func (m *EditPlayerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EditPlayerReq.DiscardUnknown(m)
}

var xxx_messageInfo_EditPlayerReq proto.InternalMessageInfo

func (m *EditPlayerReq) GetPlyr() *Player {
	if m != nil {
		return m.Plyr
	}
	return nil
}

type EditPlayerRes struct {
	Plyr                 *Player  `protobuf:"bytes,1,opt,name=plyr,proto3" json:"plyr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditPlayerRes) Reset()         { *m = EditPlayerRes{} }
func (m *EditPlayerRes) String() string { return proto.CompactTextString(m) }
func (*EditPlayerRes) ProtoMessage()    {}
func (*EditPlayerRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{4}
}

func (m *EditPlayerRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditPlayerRes.Unmarshal(m, b)
}
func (m *EditPlayerRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditPlayerRes.Marshal(b, m, deterministic)
}
func (m *EditPlayerRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditPlayerRes.Merge(m, src)
}
func (m *EditPlayerRes) XXX_Size() int {
	return xxx_messageInfo_EditPlayerRes.Size(m)
}
func (m *EditPlayerRes) XXX_DiscardUnknown() {
	xxx_messageInfo_EditPlayerRes.DiscardUnknown(m)
}

var xxx_messageInfo_EditPlayerRes proto.InternalMessageInfo

func (m *EditPlayerRes) GetPlyr() *Player {
	if m != nil {
		return m.Plyr
	}
	return nil
}

type RemovePlayerReq struct {
	PlayerId             string   `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemovePlayerReq) Reset()         { *m = RemovePlayerReq{} }
func (m *RemovePlayerReq) String() string { return proto.CompactTextString(m) }
func (*RemovePlayerReq) ProtoMessage()    {}
func (*RemovePlayerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{5}
}

func (m *RemovePlayerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemovePlayerReq.Unmarshal(m, b)
}
func (m *RemovePlayerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemovePlayerReq.Marshal(b, m, deterministic)
}
func (m *RemovePlayerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemovePlayerReq.Merge(m, src)
}
func (m *RemovePlayerReq) XXX_Size() int {
	return xxx_messageInfo_RemovePlayerReq.Size(m)
}
func (m *RemovePlayerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RemovePlayerReq.DiscardUnknown(m)
}

var xxx_messageInfo_RemovePlayerReq proto.InternalMessageInfo

func (m *RemovePlayerReq) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

type RemovePlayerRes struct {
	Removed              bool     `protobuf:"varint,1,opt,name=removed,proto3" json:"removed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemovePlayerRes) Reset()         { *m = RemovePlayerRes{} }
func (m *RemovePlayerRes) String() string { return proto.CompactTextString(m) }
func (*RemovePlayerRes) ProtoMessage()    {}
func (*RemovePlayerRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{6}
}

func (m *RemovePlayerRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemovePlayerRes.Unmarshal(m, b)
}
func (m *RemovePlayerRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemovePlayerRes.Marshal(b, m, deterministic)
}
func (m *RemovePlayerRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemovePlayerRes.Merge(m, src)
}
func (m *RemovePlayerRes) XXX_Size() int {
	return xxx_messageInfo_RemovePlayerRes.Size(m)
}
func (m *RemovePlayerRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RemovePlayerRes.DiscardUnknown(m)
}

var xxx_messageInfo_RemovePlayerRes proto.InternalMessageInfo

func (m *RemovePlayerRes) GetRemoved() bool {
	if m != nil {
		return m.Removed
	}
	return false
}

type GetPlayerReq struct {
	PlayerId             string   `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlayerReq) Reset()         { *m = GetPlayerReq{} }
func (m *GetPlayerReq) String() string { return proto.CompactTextString(m) }
func (*GetPlayerReq) ProtoMessage()    {}
func (*GetPlayerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{7}
}

func (m *GetPlayerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlayerReq.Unmarshal(m, b)
}
func (m *GetPlayerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlayerReq.Marshal(b, m, deterministic)
}
func (m *GetPlayerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerReq.Merge(m, src)
}
func (m *GetPlayerReq) XXX_Size() int {
	return xxx_messageInfo_GetPlayerReq.Size(m)
}
func (m *GetPlayerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerReq proto.InternalMessageInfo

func (m *GetPlayerReq) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

type GetPlayerRes struct {
	Plyr                 *Player  `protobuf:"bytes,1,opt,name=plyr,proto3" json:"plyr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlayerRes) Reset()         { *m = GetPlayerRes{} }
func (m *GetPlayerRes) String() string { return proto.CompactTextString(m) }
func (*GetPlayerRes) ProtoMessage()    {}
func (*GetPlayerRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{8}
}

func (m *GetPlayerRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlayerRes.Unmarshal(m, b)
}
func (m *GetPlayerRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlayerRes.Marshal(b, m, deterministic)
}
func (m *GetPlayerRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerRes.Merge(m, src)
}
func (m *GetPlayerRes) XXX_Size() int {
	return xxx_messageInfo_GetPlayerRes.Size(m)
}
func (m *GetPlayerRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerRes proto.InternalMessageInfo

func (m *GetPlayerRes) GetPlyr() *Player {
	if m != nil {
		return m.Plyr
	}
	return nil
}

type GetPlayerListReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlayerListReq) Reset()         { *m = GetPlayerListReq{} }
func (m *GetPlayerListReq) String() string { return proto.CompactTextString(m) }
func (*GetPlayerListReq) ProtoMessage()    {}
func (*GetPlayerListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{9}
}

func (m *GetPlayerListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlayerListReq.Unmarshal(m, b)
}
func (m *GetPlayerListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlayerListReq.Marshal(b, m, deterministic)
}
func (m *GetPlayerListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerListReq.Merge(m, src)
}
func (m *GetPlayerListReq) XXX_Size() int {
	return xxx_messageInfo_GetPlayerListReq.Size(m)
}
func (m *GetPlayerListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerListReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerListReq proto.InternalMessageInfo

type GetPlayerListRes struct {
	Plyr                 *Player  `protobuf:"bytes,1,opt,name=plyr,proto3" json:"plyr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlayerListRes) Reset()         { *m = GetPlayerListRes{} }
func (m *GetPlayerListRes) String() string { return proto.CompactTextString(m) }
func (*GetPlayerListRes) ProtoMessage()    {}
func (*GetPlayerListRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{10}
}

func (m *GetPlayerListRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlayerListRes.Unmarshal(m, b)
}
func (m *GetPlayerListRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlayerListRes.Marshal(b, m, deterministic)
}
func (m *GetPlayerListRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlayerListRes.Merge(m, src)
}
func (m *GetPlayerListRes) XXX_Size() int {
	return xxx_messageInfo_GetPlayerListRes.Size(m)
}
func (m *GetPlayerListRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlayerListRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlayerListRes proto.InternalMessageInfo

func (m *GetPlayerListRes) GetPlyr() *Player {
	if m != nil {
		return m.Plyr
	}
	return nil
}

func init() {
	proto.RegisterType((*Player)(nil), "player.Player")
	proto.RegisterType((*AddPlayerReq)(nil), "player.AddPlayerReq")
	proto.RegisterType((*AddPlayerRes)(nil), "player.AddPlayerRes")
	proto.RegisterType((*EditPlayerReq)(nil), "player.EditPlayerReq")
	proto.RegisterType((*EditPlayerRes)(nil), "player.EditPlayerRes")
	proto.RegisterType((*RemovePlayerReq)(nil), "player.RemovePlayerReq")
	proto.RegisterType((*RemovePlayerRes)(nil), "player.RemovePlayerRes")
	proto.RegisterType((*GetPlayerReq)(nil), "player.GetPlayerReq")
	proto.RegisterType((*GetPlayerRes)(nil), "player.GetPlayerRes")
	proto.RegisterType((*GetPlayerListReq)(nil), "player.GetPlayerListReq")
	proto.RegisterType((*GetPlayerListRes)(nil), "player.GetPlayerListRes")
}

func init() { proto.RegisterFile("player.proto", fileDescriptor_41d803d1b635d5c6) }

var fileDescriptor_41d803d1b635d5c6 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x51, 0x4b, 0xf3, 0x30,
	0x14, 0xa5, 0xdd, 0xbe, 0x7d, 0xdb, 0xb5, 0x9b, 0xe3, 0xa2, 0x18, 0xe2, 0x8b, 0xe4, 0x49, 0x18,
	0x0c, 0xe9, 0x40, 0xc1, 0x27, 0x15, 0x86, 0x08, 0x3e, 0x48, 0x7d, 0xf3, 0x45, 0x36, 0x13, 0x21,
	0xd0, 0x2d, 0x5d, 0x53, 0x07, 0x7b, 0xf2, 0xdf, 0xf9, 0xbb, 0x64, 0x49, 0xba, 0xad, 0x5d, 0x07,
	0xf5, 0x2d, 0xf7, 0xdc, 0x73, 0xee, 0x3d, 0x4d, 0x4e, 0x21, 0x48, 0xe2, 0xc9, 0x4a, 0xa4, 0xc3,
	0x24, 0x55, 0x99, 0xc2, 0x96, 0xad, 0xd8, 0x37, 0xb4, 0x5e, 0xcc, 0x09, 0x7b, 0xe0, 0x4b, 0x4e,
	0xbc, 0x0b, 0xef, 0xb2, 0x13, 0xf9, 0x92, 0xe3, 0x39, 0x74, 0x2c, 0xe7, 0x5d, 0x72, 0xe2, 0x1b,
	0xb8, 0x6d, 0x81, 0x27, 0x8e, 0x14, 0xda, 0x9f, 0x5f, 0x71, 0x3c, 0x9f, 0xcc, 0x04, 0x69, 0xd8,
	0x5e, 0x5e, 0xaf, 0x7b, 0x89, 0xd2, 0x32, 0x93, 0x6a, 0x4e, 0x9a, 0x4e, 0xe7, 0x6a, 0xec, 0x43,
	0x63, 0x2a, 0x15, 0xf9, 0x67, 0xe0, 0xf5, 0x91, 0x85, 0x10, 0xdc, 0x73, 0x6e, 0x3d, 0x44, 0x62,
	0x81, 0x0c, 0x9a, 0x49, 0xbc, 0x4a, 0x8d, 0x91, 0xa3, 0xb0, 0x37, 0x74, 0xae, 0x1d, 0xc1, 0xf4,
	0x4a, 0x1a, 0x5d, 0x4b, 0x33, 0x82, 0xee, 0x98, 0xcb, 0xec, 0x6f, 0x8b, 0x4a, 0xa2, 0x7a, 0x9b,
	0x86, 0x70, 0x1c, 0x89, 0x99, 0x5a, 0x8a, 0xed, 0xae, 0xc2, 0x5d, 0x7a, 0xc5, 0xbb, 0x64, 0x83,
	0x32, 0x5f, 0x23, 0x81, 0xff, 0xa9, 0x81, 0x2c, 0xbb, 0x1d, 0xe5, 0x25, 0x1b, 0x40, 0xf0, 0x28,
	0xb2, 0x9a, 0x93, 0xc3, 0x02, 0xb9, 0x9e, 0x7b, 0x84, 0xfe, 0x46, 0xf3, 0x2c, 0x75, 0x16, 0x89,
	0x05, 0xbb, 0xde, 0xc3, 0x6a, 0xcd, 0x0a, 0x7f, 0x7c, 0xe8, 0x5a, 0xe0, 0x55, 0xa4, 0x4b, 0xf9,
	0x21, 0xf0, 0x06, 0x3a, 0x9b, 0x49, 0x78, 0x92, 0x8b, 0x76, 0xbf, 0x88, 0x56, 0xa1, 0x1a, 0xc7,
	0xd0, 0x2d, 0x58, 0x40, 0xb2, 0x47, 0x73, 0x6e, 0xe9, 0xa1, 0x8e, 0xbe, 0xf2, 0xd6, 0xfb, 0x37,
	0xc9, 0xd9, 0xee, 0xdf, 0x0d, 0x20, 0xad, 0x42, 0x35, 0xde, 0x02, 0x6c, 0x93, 0x80, 0xa7, 0x39,
	0xa7, 0x10, 0x29, 0x5a, 0x09, 0x6b, 0xbc, 0x83, 0x60, 0xf7, 0x81, 0xf1, 0x2c, 0xa7, 0x95, 0x62,
	0x42, 0x0f, 0x34, 0xf4, 0x03, 0xbc, 0xb9, 0x47, 0x4d, 0xa6, 0xd3, 0x96, 0xf9, 0x81, 0x47, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x40, 0x12, 0x19, 0xd0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlayerServiceClient is the client API for PlayerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlayerServiceClient interface {
	GetPlayer(ctx context.Context, in *GetPlayerReq, opts ...grpc.CallOption) (*GetPlayerRes, error)
	GetPlayerList(ctx context.Context, in *GetPlayerListReq, opts ...grpc.CallOption) (PlayerService_GetPlayerListClient, error)
	AddPlayer(ctx context.Context, in *AddPlayerReq, opts ...grpc.CallOption) (*AddPlayerRes, error)
	EditPlayer(ctx context.Context, in *EditPlayerReq, opts ...grpc.CallOption) (*EditPlayerRes, error)
	RemovePlayer(ctx context.Context, in *RemovePlayerReq, opts ...grpc.CallOption) (*RemovePlayerRes, error)
}

type playerServiceClient struct {
	cc *grpc.ClientConn
}

func NewPlayerServiceClient(cc *grpc.ClientConn) PlayerServiceClient {
	return &playerServiceClient{cc}
}

func (c *playerServiceClient) GetPlayer(ctx context.Context, in *GetPlayerReq, opts ...grpc.CallOption) (*GetPlayerRes, error) {
	out := new(GetPlayerRes)
	err := c.cc.Invoke(ctx, "/player.PlayerService/GetPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerServiceClient) GetPlayerList(ctx context.Context, in *GetPlayerListReq, opts ...grpc.CallOption) (PlayerService_GetPlayerListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PlayerService_serviceDesc.Streams[0], "/player.PlayerService/GetPlayerList", opts...)
	if err != nil {
		return nil, err
	}
	x := &playerServiceGetPlayerListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PlayerService_GetPlayerListClient interface {
	Recv() (*GetPlayerListRes, error)
	grpc.ClientStream
}

type playerServiceGetPlayerListClient struct {
	grpc.ClientStream
}

func (x *playerServiceGetPlayerListClient) Recv() (*GetPlayerListRes, error) {
	m := new(GetPlayerListRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *playerServiceClient) AddPlayer(ctx context.Context, in *AddPlayerReq, opts ...grpc.CallOption) (*AddPlayerRes, error) {
	out := new(AddPlayerRes)
	err := c.cc.Invoke(ctx, "/player.PlayerService/AddPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerServiceClient) EditPlayer(ctx context.Context, in *EditPlayerReq, opts ...grpc.CallOption) (*EditPlayerRes, error) {
	out := new(EditPlayerRes)
	err := c.cc.Invoke(ctx, "/player.PlayerService/EditPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerServiceClient) RemovePlayer(ctx context.Context, in *RemovePlayerReq, opts ...grpc.CallOption) (*RemovePlayerRes, error) {
	out := new(RemovePlayerRes)
	err := c.cc.Invoke(ctx, "/player.PlayerService/RemovePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerServiceServer is the server API for PlayerService service.
type PlayerServiceServer interface {
	GetPlayer(context.Context, *GetPlayerReq) (*GetPlayerRes, error)
	GetPlayerList(*GetPlayerListReq, PlayerService_GetPlayerListServer) error
	AddPlayer(context.Context, *AddPlayerReq) (*AddPlayerRes, error)
	EditPlayer(context.Context, *EditPlayerReq) (*EditPlayerRes, error)
	RemovePlayer(context.Context, *RemovePlayerReq) (*RemovePlayerRes, error)
}

// UnimplementedPlayerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPlayerServiceServer struct {
}

func (*UnimplementedPlayerServiceServer) GetPlayer(ctx context.Context, req *GetPlayerReq) (*GetPlayerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayer not implemented")
}
func (*UnimplementedPlayerServiceServer) GetPlayerList(req *GetPlayerListReq, srv PlayerService_GetPlayerListServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPlayerList not implemented")
}
func (*UnimplementedPlayerServiceServer) AddPlayer(ctx context.Context, req *AddPlayerReq) (*AddPlayerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPlayer not implemented")
}
func (*UnimplementedPlayerServiceServer) EditPlayer(ctx context.Context, req *EditPlayerReq) (*EditPlayerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditPlayer not implemented")
}
func (*UnimplementedPlayerServiceServer) RemovePlayer(ctx context.Context, req *RemovePlayerReq) (*RemovePlayerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePlayer not implemented")
}

func RegisterPlayerServiceServer(s *grpc.Server, srv PlayerServiceServer) {
	s.RegisterService(&_PlayerService_serviceDesc, srv)
}

func _PlayerService_GetPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).GetPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/player.PlayerService/GetPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).GetPlayer(ctx, req.(*GetPlayerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerService_GetPlayerList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetPlayerListReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PlayerServiceServer).GetPlayerList(m, &playerServiceGetPlayerListServer{stream})
}

type PlayerService_GetPlayerListServer interface {
	Send(*GetPlayerListRes) error
	grpc.ServerStream
}

type playerServiceGetPlayerListServer struct {
	grpc.ServerStream
}

func (x *playerServiceGetPlayerListServer) Send(m *GetPlayerListRes) error {
	return x.ServerStream.SendMsg(m)
}

func _PlayerService_AddPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPlayerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).AddPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/player.PlayerService/AddPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).AddPlayer(ctx, req.(*AddPlayerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerService_EditPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditPlayerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).EditPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/player.PlayerService/EditPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).EditPlayer(ctx, req.(*EditPlayerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerService_RemovePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePlayerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).RemovePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/player.PlayerService/RemovePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).RemovePlayer(ctx, req.(*RemovePlayerReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlayerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "player.PlayerService",
	HandlerType: (*PlayerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayer",
			Handler:    _PlayerService_GetPlayer_Handler,
		},
		{
			MethodName: "AddPlayer",
			Handler:    _PlayerService_AddPlayer_Handler,
		},
		{
			MethodName: "EditPlayer",
			Handler:    _PlayerService_EditPlayer_Handler,
		},
		{
			MethodName: "RemovePlayer",
			Handler:    _PlayerService_RemovePlayer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPlayerList",
			Handler:       _PlayerService_GetPlayerList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "player.proto",
}
