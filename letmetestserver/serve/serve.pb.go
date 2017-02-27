// Code generated by protoc-gen-gogo.
// source: serve.proto
// DO NOT EDIT!

/*
Package serve is a generated protocol buffer package.

It is generated from these files:
	serve.proto

It has these top-level messages:
	Artist
	Song
	Album
	EndLess
	Tree
*/
package serve

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Instrument int32

const (
	Instrument_Voice  Instrument = 0
	Instrument_Guitar Instrument = 1
	Instrument_Drum   Instrument = 2
)

var Instrument_name = map[int32]string{
	0: "Voice",
	1: "Guitar",
	2: "Drum",
}
var Instrument_value = map[string]int32{
	"Voice":  0,
	"Guitar": 1,
	"Drum":   2,
}

func (x Instrument) String() string {
	return proto.EnumName(Instrument_name, int32(x))
}
func (Instrument) EnumDescriptor() ([]byte, []int) { return fileDescriptorServe, []int{0} }

type Genre int32

const (
	Genre_Pop          Genre = 0
	Genre_Rock         Genre = 1
	Genre_Jazz         Genre = 2
	Genre_NintendoCore Genre = 3
	Genre_Indie        Genre = 4
	Genre_Punk         Genre = 5
	Genre_Dance        Genre = 6
)

var Genre_name = map[int32]string{
	0: "Pop",
	1: "Rock",
	2: "Jazz",
	3: "NintendoCore",
	4: "Indie",
	5: "Punk",
	6: "Dance",
}
var Genre_value = map[string]int32{
	"Pop":          0,
	"Rock":         1,
	"Jazz":         2,
	"NintendoCore": 3,
	"Indie":        4,
	"Punk":         5,
	"Dance":        6,
}

func (x Genre) String() string {
	return proto.EnumName(Genre_name, int32(x))
}
func (Genre) EnumDescriptor() ([]byte, []int) { return fileDescriptorServe, []int{1} }

type Artist struct {
	// Pick something original
	Name string     `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Role Instrument `protobuf:"varint,2,opt,name=Role,proto3,enum=serve.Instrument" json:"Role,omitempty"`
}

func (m *Artist) Reset()                    { *m = Artist{} }
func (m *Artist) String() string            { return proto.CompactTextString(m) }
func (*Artist) ProtoMessage()               {}
func (*Artist) Descriptor() ([]byte, []int) { return fileDescriptorServe, []int{0} }

func (m *Artist) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Artist) GetRole() Instrument {
	if m != nil {
		return m.Role
	}
	return Instrument_Voice
}

type Song struct {
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// 1,2,3,4...
	Track    uint64    `protobuf:"varint,2,opt,name=Track,proto3" json:"Track,omitempty"`
	Duration float64   `protobuf:"fixed64,3,opt,name=Duration,proto3" json:"Duration,omitempty"`
	Composer []*Artist `protobuf:"bytes,4,rep,name=Composer" json:"Composer,omitempty"`
}

func (m *Song) Reset()                    { *m = Song{} }
func (m *Song) String() string            { return proto.CompactTextString(m) }
func (*Song) ProtoMessage()               {}
func (*Song) Descriptor() ([]byte, []int) { return fileDescriptorServe, []int{1} }

func (m *Song) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Song) GetTrack() uint64 {
	if m != nil {
		return m.Track
	}
	return 0
}

func (m *Song) GetDuration() float64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Song) GetComposer() []*Artist {
	if m != nil {
		return m.Composer
	}
	return nil
}

type Album struct {
	// Untitled?
	Name  string  `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Song  []*Song `protobuf:"bytes,2,rep,name=Song" json:"Song,omitempty"`
	Genre Genre   `protobuf:"varint,3,opt,name=Genre,proto3,enum=serve.Genre" json:"Genre,omitempty"`
	// 2015
	Year string `protobuf:"bytes,4,opt,name=Year,proto3" json:"Year,omitempty"`
	// Uhm ja
	Producer []string `protobuf:"bytes,5,rep,name=Producer" json:"Producer,omitempty"`
	Mediocre bool     `protobuf:"varint,6,opt,name=Mediocre,proto3" json:"Mediocre,omitempty"`
	Rated    bool     `protobuf:"varint,7,opt,name=Rated,proto3" json:"Rated,omitempty"`
	Epilogue string   `protobuf:"bytes,8,opt,name=Epilogue,proto3" json:"Epilogue,omitempty"`
	Likes    []bool   `protobuf:"varint,9,rep,packed,name=Likes" json:"Likes,omitempty"`
}

func (m *Album) Reset()                    { *m = Album{} }
func (m *Album) String() string            { return proto.CompactTextString(m) }
func (*Album) ProtoMessage()               {}
func (*Album) Descriptor() ([]byte, []int) { return fileDescriptorServe, []int{2} }

func (m *Album) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Album) GetSong() []*Song {
	if m != nil {
		return m.Song
	}
	return nil
}

func (m *Album) GetGenre() Genre {
	if m != nil {
		return m.Genre
	}
	return Genre_Pop
}

func (m *Album) GetYear() string {
	if m != nil {
		return m.Year
	}
	return ""
}

func (m *Album) GetProducer() []string {
	if m != nil {
		return m.Producer
	}
	return nil
}

func (m *Album) GetMediocre() bool {
	if m != nil {
		return m.Mediocre
	}
	return false
}

func (m *Album) GetRated() bool {
	if m != nil {
		return m.Rated
	}
	return false
}

func (m *Album) GetEpilogue() string {
	if m != nil {
		return m.Epilogue
	}
	return ""
}

func (m *Album) GetLikes() []bool {
	if m != nil {
		return m.Likes
	}
	return nil
}

type EndLess struct {
	Tree *Tree `protobuf:"bytes,1,opt,name=Tree" json:"Tree,omitempty"`
}

func (m *EndLess) Reset()                    { *m = EndLess{} }
func (m *EndLess) String() string            { return proto.CompactTextString(m) }
func (*EndLess) ProtoMessage()               {}
func (*EndLess) Descriptor() ([]byte, []int) { return fileDescriptorServe, []int{3} }

func (m *EndLess) GetTree() *Tree {
	if m != nil {
		return m.Tree
	}
	return nil
}

type Tree struct {
	// Types that are valid to be assigned to Stuff:
	//	*Tree_ValueString
	//	*Tree_ValueNum
	Stuff isTree_Stuff `protobuf_oneof:"stuff"`
	Left  *Tree        `protobuf:"bytes,2,opt,name=Left" json:"Left,omitempty"`
	Right *Tree        `protobuf:"bytes,3,opt,name=Right" json:"Right,omitempty"`
}

func (m *Tree) Reset()                    { *m = Tree{} }
func (m *Tree) String() string            { return proto.CompactTextString(m) }
func (*Tree) ProtoMessage()               {}
func (*Tree) Descriptor() ([]byte, []int) { return fileDescriptorServe, []int{4} }

func (m *Tree) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type isTree_Stuff interface {
	isTree_Stuff()
}

type Tree_ValueString struct {
	ValueString string `protobuf:"bytes,1,opt,name=ValueString,proto3,oneof"`
}
type Tree_ValueNum struct {
	ValueNum uint64 `protobuf:"varint,4,opt,name=ValueNum,proto3,oneof"`
}

func (*Tree_ValueString) isTree_Stuff() {}
func (*Tree_ValueNum) isTree_Stuff()    {}

func (m *Tree) GetStuff() isTree_Stuff {
	if m != nil {
		return m.Stuff
	}
	return nil
}

func (m *Tree) GetValueString() string {
	if x, ok := m.GetStuff().(*Tree_ValueString); ok {
		return x.ValueString
	}
	return ""
}

func (m *Tree) GetValueNum() uint64 {
	if x, ok := m.GetStuff().(*Tree_ValueNum); ok {
		return x.ValueNum
	}
	return 0
}

func (m *Tree) GetLeft() *Tree {
	if m != nil {
		return m.Left
	}
	return nil
}

func (m *Tree) GetRight() *Tree {
	if m != nil {
		return m.Right
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Tree) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _Tree_OneofMarshaler, _Tree_OneofUnmarshaler, []interface{}{
		(*Tree_ValueString)(nil),
		(*Tree_ValueNum)(nil),
	}
}

func _Tree_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Tree)
	// stuff
	switch x := m.Stuff.(type) {
	case *Tree_ValueString:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.ValueString)
	case *Tree_ValueNum:
		_ = b.EncodeVarint(4<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.ValueNum))
	case nil:
	default:
		return fmt.Errorf("Tree.Stuff has unexpected type %T", x)
	}
	return nil
}

func _Tree_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Tree)
	switch tag {
	case 1: // stuff.ValueString
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Stuff = &Tree_ValueString{x}
		return true, err
	case 4: // stuff.ValueNum
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Stuff = &Tree_ValueNum{x}
		return true, err
	default:
		return false, nil
	}
}

func init() {
	proto.RegisterType((*Artist)(nil), "serve.Artist")
	proto.RegisterType((*Song)(nil), "serve.Song")
	proto.RegisterType((*Album)(nil), "serve.Album")
	proto.RegisterType((*EndLess)(nil), "serve.EndLess")
	proto.RegisterType((*Tree)(nil), "serve.Tree")
	proto.RegisterEnum("serve.Instrument", Instrument_name, Instrument_value)
	proto.RegisterEnum("serve.Genre", Genre_name, Genre_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Label service

type LabelClient interface {
	Produce(ctx context.Context, in *Album, opts ...grpc.CallOption) (*Album, error)
	Loop(ctx context.Context, in *EndLess, opts ...grpc.CallOption) (*EndLess, error)
}

type labelClient struct {
	cc *grpc.ClientConn
}

func NewLabelClient(cc *grpc.ClientConn) LabelClient {
	return &labelClient{cc}
}

func (c *labelClient) Produce(ctx context.Context, in *Album, opts ...grpc.CallOption) (*Album, error) {
	out := new(Album)
	err := grpc.Invoke(ctx, "/serve.Label/Produce", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelClient) Loop(ctx context.Context, in *EndLess, opts ...grpc.CallOption) (*EndLess, error) {
	out := new(EndLess)
	err := grpc.Invoke(ctx, "/serve.Label/Loop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Label service

type LabelServer interface {
	Produce(context.Context, *Album) (*Album, error)
	Loop(context.Context, *EndLess) (*EndLess, error)
}

func RegisterLabelServer(s *grpc.Server, srv LabelServer) {
	s.RegisterService(&_Label_serviceDesc, srv)
}

func _Label_Produce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Album)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServer).Produce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serve.Label/Produce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServer).Produce(ctx, req.(*Album))
	}
	return interceptor(ctx, in, info, handler)
}

func _Label_Loop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndLess)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServer).Loop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serve.Label/Loop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServer).Loop(ctx, req.(*EndLess))
	}
	return interceptor(ctx, in, info, handler)
}

var _Label_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serve.Label",
	HandlerType: (*LabelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Produce",
			Handler:    _Label_Produce_Handler,
		},
		{
			MethodName: "Loop",
			Handler:    _Label_Loop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serve.proto",
}

func init() { proto.RegisterFile("serve.proto", fileDescriptorServe) }

var fileDescriptorServe = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xdd, 0x6a, 0xdb, 0x4c,
	0x10, 0x8d, 0x2c, 0xc9, 0x3f, 0xe3, 0x7c, 0x66, 0xbf, 0xa5, 0x17, 0x8b, 0x6f, 0xaa, 0x0a, 0x52,
	0x54, 0x43, 0x73, 0xe1, 0x3e, 0x41, 0xb0, 0x43, 0x48, 0x71, 0x83, 0xd9, 0x06, 0x43, 0xef, 0xba,
	0x96, 0xa6, 0xee, 0x62, 0x7b, 0xd7, 0xac, 0x56, 0x2d, 0xe4, 0xb9, 0xfb, 0x00, 0x65, 0x77, 0x15,
	0x85, 0x9a, 0xdc, 0xe9, 0xcc, 0x99, 0x39, 0x73, 0x66, 0x66, 0x05, 0xe3, 0x1a, 0xcd, 0x2f, 0xbc,
	0x3e, 0x19, 0x6d, 0x35, 0x4d, 0x3d, 0xc8, 0x17, 0xd0, 0xbf, 0x31, 0x56, 0xd6, 0x96, 0x52, 0x48,
	0x1e, 0xc4, 0x11, 0x59, 0x94, 0x45, 0xc5, 0x88, 0xfb, 0x6f, 0x7a, 0x05, 0x09, 0xd7, 0x07, 0x64,
	0xbd, 0x2c, 0x2a, 0x26, 0xf3, 0xff, 0xaf, 0x83, 0xc0, 0xbd, 0xaa, 0xad, 0x69, 0x8e, 0xa8, 0x2c,
	0xf7, 0x74, 0xfe, 0x1b, 0x92, 0xaf, 0x5a, 0xed, 0x5e, 0x95, 0x78, 0x03, 0xe9, 0xa3, 0x11, 0xe5,
	0xde, 0x6b, 0x24, 0x3c, 0x00, 0x3a, 0x85, 0xe1, 0xb2, 0x31, 0xc2, 0x4a, 0xad, 0x58, 0x9c, 0x45,
	0x45, 0xc4, 0x3b, 0x4c, 0x3f, 0xc0, 0x70, 0xa1, 0x8f, 0x27, 0x5d, 0xa3, 0x61, 0x49, 0x16, 0x17,
	0xe3, 0xf9, 0x7f, 0x6d, 0xe3, 0xe0, 0x94, 0x77, 0x74, 0xfe, 0x27, 0x82, 0xf4, 0xe6, 0xb0, 0x6d,
	0x8e, 0xaf, 0xb6, 0x7e, 0x1b, 0x6c, 0xb1, 0x9e, 0x17, 0x19, 0xb7, 0x22, 0x2e, 0xc4, 0x83, 0xdf,
	0x1c, 0xd2, 0x3b, 0x54, 0x06, 0xbd, 0x85, 0xc9, 0xfc, 0xb2, 0xcd, 0xf0, 0x31, 0x1e, 0x28, 0x27,
	0xfc, 0x0d, 0x85, 0x73, 0xe2, 0x85, 0xdd, 0xb7, 0x73, 0xbf, 0x36, 0xba, 0x6a, 0x4a, 0x34, 0x2c,
	0xcd, 0xe2, 0x62, 0xc4, 0x3b, 0xec, 0xb8, 0x2f, 0x58, 0x49, 0x5d, 0x1a, 0x64, 0xfd, 0x2c, 0x2a,
	0x86, 0xbc, 0xc3, 0x6e, 0x17, 0x5c, 0x58, 0xac, 0xd8, 0xc0, 0x13, 0x01, 0xb8, 0x8a, 0xdb, 0x93,
	0x3c, 0xe8, 0x5d, 0x83, 0x6c, 0xe8, 0xbb, 0x74, 0xd8, 0x55, 0xac, 0xe4, 0x1e, 0x6b, 0x36, 0xca,
	0x62, 0x57, 0xe1, 0x41, 0x3e, 0x83, 0xc1, 0xad, 0xaa, 0x56, 0x58, 0xd7, 0x6e, 0xc6, 0x47, 0x83,
	0x61, 0xee, 0x97, 0x19, 0x5d, 0x88, 0x7b, 0x22, 0xff, 0x1e, 0x12, 0x9c, 0xd2, 0x46, 0x1c, 0x9a,
	0xe7, 0x0d, 0x05, 0xe0, 0xca, 0x57, 0xf8, 0xc3, 0xfa, 0xe3, 0x9c, 0x97, 0x3b, 0x82, 0xbe, 0x83,
	0x94, 0xcb, 0xdd, 0x4f, 0xeb, 0x57, 0x74, 0x96, 0x11, 0x98, 0xd9, 0x47, 0x80, 0x97, 0x17, 0x41,
	0x47, 0x90, 0x6e, 0xb4, 0x2c, 0x91, 0x5c, 0x50, 0x80, 0xfe, 0x5d, 0x23, 0xad, 0x30, 0x24, 0xa2,
	0x43, 0x48, 0x96, 0xa6, 0x39, 0x92, 0xde, 0x6c, 0xd3, 0x2e, 0x9d, 0x0e, 0x20, 0x5e, 0xeb, 0x13,
	0xb9, 0x70, 0x1c, 0xd7, 0xe5, 0x3e, 0x64, 0x7d, 0x16, 0x4f, 0x4f, 0xa4, 0x47, 0x09, 0x5c, 0x3e,
	0x48, 0x65, 0x51, 0x55, 0x7a, 0xa1, 0x0d, 0x92, 0xd8, 0x09, 0xdf, 0xab, 0x4a, 0x22, 0x49, 0x5c,
	0xda, 0xba, 0x51, 0x7b, 0x92, 0xba, 0xe0, 0x52, 0xa8, 0x12, 0x49, 0x7f, 0xbe, 0x81, 0x74, 0x25,
	0xb6, 0x78, 0xa0, 0x57, 0x30, 0x68, 0xaf, 0x41, 0x9f, 0x2f, 0xea, 0xdf, 0xc8, 0xf4, 0x1f, 0x44,
	0xdf, 0x43, 0xb2, 0xd2, 0xfa, 0x44, 0x27, 0x6d, 0xb4, 0xdd, 0xe8, 0xf4, 0x0c, 0x6f, 0xfb, 0xfe,
	0x7f, 0xf9, 0xf4, 0x37, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x5f, 0xb1, 0x60, 0x3e, 0x03, 0x00, 0x00,
}
