// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: markets/asset.proto

package ProtobufMarkets

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AssetUpdateMessage struct {
	Asset int32 `protobuf:"varint,1,opt,name=asset,proto3" json:"asset,omitempty"`
	// Types that are valid to be assigned to Update:
	//	*AssetUpdateMessage_UsdVolumeUpdate
	Update               isAssetUpdateMessage_Update `protobuf_oneof:"Update"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *AssetUpdateMessage) Reset()         { *m = AssetUpdateMessage{} }
func (m *AssetUpdateMessage) String() string { return proto.CompactTextString(m) }
func (*AssetUpdateMessage) ProtoMessage()    {}
func (*AssetUpdateMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_asset_75334e71434da020, []int{0}
}
func (m *AssetUpdateMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetUpdateMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetUpdateMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AssetUpdateMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetUpdateMessage.Merge(dst, src)
}
func (m *AssetUpdateMessage) XXX_Size() int {
	return m.Size()
}
func (m *AssetUpdateMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetUpdateMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AssetUpdateMessage proto.InternalMessageInfo

type isAssetUpdateMessage_Update interface {
	isAssetUpdateMessage_Update()
	MarshalTo([]byte) (int, error)
	Size() int
}

type AssetUpdateMessage_UsdVolumeUpdate struct {
	UsdVolumeUpdate *AssetUSDVolumeUpdate `protobuf:"bytes,2,opt,name=usdVolumeUpdate,proto3,oneof"`
}

func (*AssetUpdateMessage_UsdVolumeUpdate) isAssetUpdateMessage_Update() {}

func (m *AssetUpdateMessage) GetUpdate() isAssetUpdateMessage_Update {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *AssetUpdateMessage) GetAsset() int32 {
	if m != nil {
		return m.Asset
	}
	return 0
}

func (m *AssetUpdateMessage) GetUsdVolumeUpdate() *AssetUSDVolumeUpdate {
	if x, ok := m.GetUpdate().(*AssetUpdateMessage_UsdVolumeUpdate); ok {
		return x.UsdVolumeUpdate
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AssetUpdateMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AssetUpdateMessage_OneofMarshaler, _AssetUpdateMessage_OneofUnmarshaler, _AssetUpdateMessage_OneofSizer, []interface{}{
		(*AssetUpdateMessage_UsdVolumeUpdate)(nil),
	}
}

func _AssetUpdateMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AssetUpdateMessage)
	// Update
	switch x := m.Update.(type) {
	case *AssetUpdateMessage_UsdVolumeUpdate:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UsdVolumeUpdate); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AssetUpdateMessage.Update has unexpected type %T", x)
	}
	return nil
}

func _AssetUpdateMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AssetUpdateMessage)
	switch tag {
	case 2: // Update.usdVolumeUpdate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AssetUSDVolumeUpdate)
		err := b.DecodeMessage(msg)
		m.Update = &AssetUpdateMessage_UsdVolumeUpdate{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AssetUpdateMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AssetUpdateMessage)
	// Update
	switch x := m.Update.(type) {
	case *AssetUpdateMessage_UsdVolumeUpdate:
		s := proto.Size(x.UsdVolumeUpdate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type AssetUSDVolumeUpdate struct {
	Volume               string   `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetUSDVolumeUpdate) Reset()         { *m = AssetUSDVolumeUpdate{} }
func (m *AssetUSDVolumeUpdate) String() string { return proto.CompactTextString(m) }
func (*AssetUSDVolumeUpdate) ProtoMessage()    {}
func (*AssetUSDVolumeUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_asset_75334e71434da020, []int{1}
}
func (m *AssetUSDVolumeUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetUSDVolumeUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetUSDVolumeUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AssetUSDVolumeUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetUSDVolumeUpdate.Merge(dst, src)
}
func (m *AssetUSDVolumeUpdate) XXX_Size() int {
	return m.Size()
}
func (m *AssetUSDVolumeUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetUSDVolumeUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_AssetUSDVolumeUpdate proto.InternalMessageInfo

func (m *AssetUSDVolumeUpdate) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

func init() {
	proto.RegisterType((*AssetUpdateMessage)(nil), "ProtobufMarkets.AssetUpdateMessage")
	proto.RegisterType((*AssetUSDVolumeUpdate)(nil), "ProtobufMarkets.AssetUSDVolumeUpdate")
}
func (m *AssetUpdateMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetUpdateMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Asset != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAsset(dAtA, i, uint64(m.Asset))
	}
	if m.Update != nil {
		nn1, err := m.Update.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *AssetUpdateMessage_UsdVolumeUpdate) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.UsdVolumeUpdate != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAsset(dAtA, i, uint64(m.UsdVolumeUpdate.Size()))
		n2, err := m.UsdVolumeUpdate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *AssetUSDVolumeUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetUSDVolumeUpdate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Volume) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Volume)))
		i += copy(dAtA[i:], m.Volume)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAsset(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AssetUpdateMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Asset != 0 {
		n += 1 + sovAsset(uint64(m.Asset))
	}
	if m.Update != nil {
		n += m.Update.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AssetUpdateMessage_UsdVolumeUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UsdVolumeUpdate != nil {
		l = m.UsdVolumeUpdate.Size()
		n += 1 + l + sovAsset(uint64(l))
	}
	return n
}
func (m *AssetUSDVolumeUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Volume)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAsset(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAsset(x uint64) (n int) {
	return sovAsset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AssetUpdateMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAsset
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
			return fmt.Errorf("proto: AssetUpdateMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetUpdateMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			m.Asset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Asset |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsdVolumeUpdate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &AssetUSDVolumeUpdate{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Update = &AssetUpdateMessage_UsdVolumeUpdate{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAsset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAsset
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
func (m *AssetUSDVolumeUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAsset
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
			return fmt.Errorf("proto: AssetUSDVolumeUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetUSDVolumeUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Volume", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Volume = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAsset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAsset
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
func skipAsset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAsset
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
					return 0, ErrIntOverflowAsset
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
					return 0, ErrIntOverflowAsset
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
				return 0, ErrInvalidLengthAsset
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAsset
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
				next, err := skipAsset(dAtA[start:])
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
	ErrInvalidLengthAsset = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAsset   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("markets/asset.proto", fileDescriptor_asset_75334e71434da020) }

var fileDescriptor_asset_75334e71434da020 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x4d, 0x2c, 0xca,
	0x4e, 0x2d, 0x29, 0xd6, 0x4f, 0x2c, 0x2e, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x0f, 0x00, 0x51, 0x49, 0xa5, 0x69, 0xbe, 0x10, 0x49, 0xa5, 0x76, 0x46, 0x2e, 0x21, 0x47,
	0x90, 0x82, 0xd0, 0x82, 0x94, 0xc4, 0x92, 0x54, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x21,
	0x11, 0x2e, 0x56, 0xb0, 0x36, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x08, 0x47, 0x28, 0x90,
	0x8b, 0xbf, 0xb4, 0x38, 0x25, 0x2c, 0x3f, 0xa7, 0x34, 0x37, 0x15, 0xa2, 0x5e, 0x82, 0x49, 0x81,
	0x51, 0x83, 0xdb, 0x48, 0x55, 0x0f, 0xcd, 0x5c, 0x3d, 0x88, 0x99, 0xc1, 0x2e, 0xc8, 0x8a, 0x3d,
	0x18, 0x82, 0xd0, 0xf5, 0x3b, 0x71, 0x70, 0xb1, 0x41, 0x58, 0x4a, 0x7a, 0x5c, 0x22, 0xd8, 0x34,
	0x09, 0x89, 0x71, 0xb1, 0x95, 0x81, 0xf9, 0x60, 0xb7, 0x70, 0x06, 0x41, 0x79, 0x4e, 0x02, 0x27,
	0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c,
	0x49, 0x6c, 0x60, 0x3f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x58, 0x1a, 0x7d, 0x08, 0xfa,
	0x00, 0x00, 0x00,
}
