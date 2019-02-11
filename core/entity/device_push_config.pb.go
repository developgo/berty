// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: entity/device_push_config.proto

package entity

import (
	fmt "fmt"
	io "io"
	math "math"
	time "time"

	_ "berty.tech/core/api/protobuf/graphql"
	push "berty.tech/core/push"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Local unencrypted push configuration
type DevicePushConfig struct {
	ID                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	CreatedAt            time.Time           `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	UpdatedAt            time.Time           `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,stdtime" json:"updated_at"`
	DeviceID             string              `protobuf:"bytes,4,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	PushType             push.DevicePushType `protobuf:"varint,5,opt,name=push_type,json=pushType,proto3,enum=berty.push.DevicePushType" json:"push_type,omitempty"`
	PushID               []byte              `protobuf:"bytes,6,opt,name=push_id,json=pushId,proto3" json:"push_id,omitempty"`
	RelayPubkey          string              `protobuf:"bytes,7,opt,name=relay_pubkey,json=relayPubkey,proto3" json:"relay_pubkey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DevicePushConfig) Reset()         { *m = DevicePushConfig{} }
func (m *DevicePushConfig) String() string { return proto.CompactTextString(m) }
func (*DevicePushConfig) ProtoMessage()    {}
func (*DevicePushConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_469c6c4d4fab077f, []int{0}
}
func (m *DevicePushConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DevicePushConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DevicePushConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DevicePushConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevicePushConfig.Merge(m, src)
}
func (m *DevicePushConfig) XXX_Size() int {
	return m.Size()
}
func (m *DevicePushConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DevicePushConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DevicePushConfig proto.InternalMessageInfo

func (m *DevicePushConfig) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *DevicePushConfig) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *DevicePushConfig) GetUpdatedAt() time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return time.Time{}
}

func (m *DevicePushConfig) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *DevicePushConfig) GetPushType() push.DevicePushType {
	if m != nil {
		return m.PushType
	}
	return push.DevicePushType_UnknownDevicePushType
}

func (m *DevicePushConfig) GetPushID() []byte {
	if m != nil {
		return m.PushID
	}
	return nil
}

func (m *DevicePushConfig) GetRelayPubkey() string {
	if m != nil {
		return m.RelayPubkey
	}
	return ""
}

func init() {
	proto.RegisterType((*DevicePushConfig)(nil), "berty.entity.DevicePushConfig")
}

func init() { proto.RegisterFile("entity/device_push_config.proto", fileDescriptor_469c6c4d4fab077f) }

var fileDescriptor_469c6c4d4fab077f = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x31, 0x8e, 0x9c, 0x30,
	0x14, 0x5d, 0xcf, 0x26, 0xb3, 0x33, 0x9e, 0x51, 0x36, 0xb2, 0xa2, 0x88, 0x4c, 0x81, 0x11, 0x69,
	0x48, 0x03, 0xd1, 0xa4, 0x88, 0x94, 0x6e, 0x59, 0x1a, 0xba, 0x15, 0xda, 0x2a, 0x0d, 0x02, 0xec,
	0x05, 0x6b, 0x87, 0xb5, 0xe3, 0x31, 0x91, 0x7c, 0x86, 0x34, 0x29, 0x73, 0xa4, 0x2d, 0x73, 0x02,
	0xb2, 0x22, 0x17, 0x88, 0x72, 0x82, 0x08, 0x1b, 0x34, 0x75, 0x1a, 0xf8, 0xbc, 0xff, 0xdf, 0xfb,
	0xcf, 0xcf, 0x40, 0x4c, 0x1f, 0x14, 0x53, 0x3a, 0x22, 0xf4, 0x2b, 0xab, 0x68, 0x2e, 0xba, 0x63,
	0x93, 0x57, 0xfc, 0xe1, 0x8e, 0xd5, 0xa1, 0x90, 0x5c, 0x71, 0xb4, 0x2d, 0xa9, 0x54, 0x3a, 0xb4,
	0x63, 0x3b, 0xbf, 0x10, 0x2c, 0x32, 0x8d, 0xb2, 0xbb, 0x8b, 0x6a, 0x59, 0x88, 0xe6, 0xcb, 0x61,
	0x7e, 0x5b, 0xc6, 0xee, 0x72, 0x14, 0x89, 0xc6, 0xc7, 0x04, 0xbc, 0xaa, 0x79, 0xcd, 0x4d, 0x19,
	0x8d, 0xd5, 0x84, 0xe2, 0x9a, 0xf3, 0xfa, 0x40, 0x4f, 0x6a, 0x8a, 0xb5, 0xf4, 0xa8, 0x8a, 0x56,
	0xd8, 0x01, 0xff, 0xdb, 0x39, 0x7c, 0x99, 0x18, 0x5b, 0x37, 0xdd, 0xb1, 0xb9, 0x36, 0xa6, 0xd0,
	0x7b, 0xb8, 0x60, 0xc4, 0x01, 0x1e, 0x08, 0xd6, 0xb1, 0xf7, 0xf4, 0xe7, 0x0d, 0x18, 0x7a, 0xbc,
	0x48, 0x93, 0xbf, 0x3d, 0x46, 0x35, 0x97, 0xed, 0x27, 0x5f, 0x48, 0xd6, 0x16, 0x52, 0xe7, 0xf7,
	0x54, 0xfb, 0xd9, 0x82, 0x11, 0x74, 0x0d, 0x61, 0x25, 0x69, 0xa1, 0x28, 0xc9, 0x0b, 0xe5, 0x2c,
	0x3c, 0x10, 0x6c, 0xf6, 0xbb, 0xd0, 0x2e, 0x0f, 0xe7, 0xe5, 0xe1, 0xed, 0xbc, 0x3c, 0x5e, 0x3d,
	0xf6, 0xf8, 0xec, 0xfb, 0x2f, 0x0c, 0xb2, 0xf5, 0xc4, 0xbb, 0x52, 0xa3, 0x48, 0x27, 0xc8, 0x2c,
	0x72, 0xfe, 0x3f, 0x22, 0x13, 0xef, 0x4a, 0xa1, 0x77, 0x70, 0x3d, 0xc5, 0xcc, 0x88, 0xf3, 0xcc,
	0x1c, 0x61, 0x3b, 0xf4, 0x78, 0x65, 0x0f, 0x99, 0x26, 0xd9, 0xca, 0xb6, 0x53, 0x82, 0x3e, 0xc2,
	0xb5, 0xb9, 0x0a, 0xa5, 0x05, 0x75, 0x9e, 0x7b, 0x20, 0x78, 0xb1, 0xdf, 0x85, 0xf6, 0x26, 0x4c,
	0xb0, 0xa7, 0x5c, 0x6e, 0xb5, 0xa0, 0xd9, 0x4a, 0x4c, 0x15, 0x7a, 0x0b, 0x2f, 0x0c, 0x91, 0x11,
	0x67, 0xe9, 0x81, 0x60, 0x1b, 0xc3, 0xa1, 0xc7, 0xcb, 0x71, 0x30, 0x4d, 0xb2, 0xe5, 0xd8, 0x4a,
	0x09, 0xda, 0xc3, 0xad, 0xa4, 0x87, 0x42, 0xe7, 0xa2, 0x2b, 0xef, 0xa9, 0x76, 0x2e, 0x8c, 0x97,
	0xcb, 0xa1, 0xc7, 0x9b, 0x6c, 0xc4, 0x6f, 0x0c, 0x9c, 0x6d, 0xe4, 0xe9, 0x23, 0x0e, 0x1e, 0x07,
	0x17, 0xfc, 0x1c, 0x5c, 0xf0, 0x34, 0xb8, 0xe0, 0xc7, 0x6f, 0xf7, 0xec, 0xf3, 0x6b, 0xeb, 0x47,
	0xd1, 0xaa, 0x89, 0x2a, 0x2e, 0x69, 0x64, 0xff, 0x91, 0x72, 0x69, 0xf2, 0xf8, 0xf0, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0xb1, 0xee, 0xdf, 0xb0, 0x5b, 0x02, 0x00, 0x00,
}

func (m *DevicePushConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DevicePushConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDevicePushConfig(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintDevicePushConfig(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)))
	n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x1a
	i++
	i = encodeVarintDevicePushConfig(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdatedAt)))
	n2, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.UpdatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.DeviceID) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintDevicePushConfig(dAtA, i, uint64(len(m.DeviceID)))
		i += copy(dAtA[i:], m.DeviceID)
	}
	if m.PushType != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintDevicePushConfig(dAtA, i, uint64(m.PushType))
	}
	if len(m.PushID) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintDevicePushConfig(dAtA, i, uint64(len(m.PushID)))
		i += copy(dAtA[i:], m.PushID)
	}
	if len(m.RelayPubkey) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintDevicePushConfig(dAtA, i, uint64(len(m.RelayPubkey)))
		i += copy(dAtA[i:], m.RelayPubkey)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintDevicePushConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DevicePushConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovDevicePushConfig(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovDevicePushConfig(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdatedAt)
	n += 1 + l + sovDevicePushConfig(uint64(l))
	l = len(m.DeviceID)
	if l > 0 {
		n += 1 + l + sovDevicePushConfig(uint64(l))
	}
	if m.PushType != 0 {
		n += 1 + sovDevicePushConfig(uint64(m.PushType))
	}
	l = len(m.PushID)
	if l > 0 {
		n += 1 + l + sovDevicePushConfig(uint64(l))
	}
	l = len(m.RelayPubkey)
	if l > 0 {
		n += 1 + l + sovDevicePushConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDevicePushConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDevicePushConfig(x uint64) (n int) {
	return sovDevicePushConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DevicePushConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDevicePushConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DevicePushConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DevicePushConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDevicePushConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDevicePushConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDevicePushConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDevicePushConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PushType", wireType)
			}
			m.PushType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PushType |= (push.DevicePushType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PushID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDevicePushConfig
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PushID = append(m.PushID[:0], dAtA[iNdEx:postIndex]...)
			if m.PushID == nil {
				m.PushID = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayPubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDevicePushConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RelayPubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDevicePushConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDevicePushConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDevicePushConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDevicePushConfig
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
					return 0, ErrIntOverflowDevicePushConfig
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
					return 0, ErrIntOverflowDevicePushConfig
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDevicePushConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDevicePushConfig
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
				next, err := skipDevicePushConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthDevicePushConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDevicePushConfig   = fmt.Errorf("proto: integer overflow")
)
