// Code generated by protoc-gen-go. DO NOT EDIT.
// source: GetUserMappingsProtocol.proto

package hadoop_common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// *
//  Get groups for user request.
type GetGroupsForUserRequestProto struct {
	User             *string `protobuf:"bytes,1,req,name=user" json:"user,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetGroupsForUserRequestProto) Reset()                    { *m = GetGroupsForUserRequestProto{} }
func (m *GetGroupsForUserRequestProto) String() string            { return proto.CompactTextString(m) }
func (*GetGroupsForUserRequestProto) ProtoMessage()               {}
func (*GetGroupsForUserRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *GetGroupsForUserRequestProto) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return ""
}

// *
// Response for get groups.
type GetGroupsForUserResponseProto struct {
	Groups           []string `protobuf:"bytes,1,rep,name=groups" json:"groups,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GetGroupsForUserResponseProto) Reset()                    { *m = GetGroupsForUserResponseProto{} }
func (m *GetGroupsForUserResponseProto) String() string            { return proto.CompactTextString(m) }
func (*GetGroupsForUserResponseProto) ProtoMessage()               {}
func (*GetGroupsForUserResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *GetGroupsForUserResponseProto) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func init() {
	proto.RegisterType((*GetGroupsForUserRequestProto)(nil), "hadoop.common.GetGroupsForUserRequestProto")
	proto.RegisterType((*GetGroupsForUserResponseProto)(nil), "hadoop.common.GetGroupsForUserResponseProto")
}

func init() { proto.RegisterFile("GetUserMappingsProtocol.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x75, 0x4f, 0x2d, 0x09,
	0x2d, 0x4e, 0x2d, 0xf2, 0x4d, 0x2c, 0x28, 0xc8, 0xcc, 0x4b, 0x2f, 0x0e, 0x28, 0xca, 0x2f, 0xc9,
	0x4f, 0xce, 0xcf, 0xd1, 0x2b, 0x00, 0x31, 0x84, 0x78, 0x33, 0x12, 0x53, 0xf2, 0xf3, 0x0b, 0xf4,
	0x92, 0xf3, 0x73, 0x73, 0xf3, 0xf3, 0x94, 0x8c, 0xb8, 0x64, 0xdc, 0x53, 0x4b, 0xdc, 0x8b, 0xf2,
	0x4b, 0x0b, 0x8a, 0xdd, 0xf2, 0x8b, 0x40, 0x1a, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xc0,
	0xfa, 0x84, 0x84, 0xb8, 0x58, 0x4a, 0x8b, 0x53, 0x8b, 0x24, 0x18, 0x15, 0x98, 0x34, 0x38, 0x83,
	0xc0, 0x6c, 0x25, 0x73, 0xb0, 0x1d, 0x68, 0x7a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x21, 0x9a,
	0xc4, 0xb8, 0xd8, 0xd2, 0xc1, 0xb2, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x51,
	0x3f, 0x23, 0x97, 0x1c, 0x0e, 0xd7, 0x05, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xe5, 0x72,
	0x09, 0xa4, 0xa3, 0x99, 0x2d, 0xa4, 0xad, 0x87, 0xe2, 0x66, 0x3d, 0x7c, 0x0e, 0x96, 0xd2, 0x21,
	0xa8, 0x18, 0xc9, 0xa5, 0x4e, 0x2e, 0x5c, 0xb2, 0xf9, 0x45, 0xe9, 0x7a, 0x89, 0x05, 0x89, 0xc9,
	0x19, 0xa9, 0x30, 0x9d, 0x25, 0xf9, 0xf9, 0x39, 0xc5, 0x90, 0xe0, 0x72, 0xc2, 0x15, 0x9a, 0x60,
	0xba, 0xb8, 0x83, 0x91, 0x71, 0x01, 0x23, 0x23, 0x20, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x38, 0xcc,
	0x6f, 0x73, 0x01, 0x00, 0x00,
}
