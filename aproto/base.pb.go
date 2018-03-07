// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: base.proto

/*
	Package aproto is a generated protocol buffer package.

	It is generated from these files:
		base.proto
		market_data.proto

	It has these top-level messages:
		Symbol
		Kline
		KlineSeries
		OrderBook
		MarketDataSnapshot
		MdsList
*/
package aproto

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ExchangeType int32

const (
	// / 上海期货交易所
	ExchangeType_SHFE ExchangeType = 0
	// / 郑州商品交易所
	ExchangeType_CZCE ExchangeType = 1
	// / 大连商品期货交易所
	ExchangeType_DCE ExchangeType = 2
	// / 中金所
	ExchangeType_CFFEX ExchangeType = 3
	// / 上海证券交易所
	ExchangeType_SSE ExchangeType = 4
	// / 深圳证券交易所
	ExchangeType_SZE ExchangeType = 5
	// / 香港证券交易所
	ExchangeType_HKG ExchangeType = 6
	// / 香港期货交易所
	ExchangeType_HKFE ExchangeType = 7
)

var ExchangeType_name = map[int32]string{
	0: "SHFE",
	1: "CZCE",
	2: "DCE",
	3: "CFFEX",
	4: "SSE",
	5: "SZE",
	6: "HKG",
	7: "HKFE",
}
var ExchangeType_value = map[string]int32{
	"SHFE":  0,
	"CZCE":  1,
	"DCE":   2,
	"CFFEX": 3,
	"SSE":   4,
	"SZE":   5,
	"HKG":   6,
	"HKFE":  7,
}

func (x ExchangeType) String() string {
	return proto.EnumName(ExchangeType_name, int32(x))
}
func (ExchangeType) EnumDescriptor() ([]byte, []int) { return fileDescriptorBase, []int{0} }

type PeriodType int32

const (
	// Tick
	PeriodType_TICK PeriodType = 0
	// 分钟
	PeriodType_M1 PeriodType = 1
	// 5分钟
	PeriodType_M3 PeriodType = 2
	// 10分钟
	PeriodType_M10 PeriodType = 3
	// 15分钟
	PeriodType_M15 PeriodType = 4
	// 小时
	PeriodType_H1 PeriodType = 5
	// 3小时
	PeriodType_H3 PeriodType = 6
	// 日线
	PeriodType_D1 PeriodType = 7
	// 周线
	PeriodType_W1 PeriodType = 8
	// 月线
	PeriodType_MON1 PeriodType = 9
)

var PeriodType_name = map[int32]string{
	0: "TICK",
	1: "M1",
	2: "M3",
	3: "M10",
	4: "M15",
	5: "H1",
	6: "H3",
	7: "D1",
	8: "W1",
	9: "MON1",
}
var PeriodType_value = map[string]int32{
	"TICK": 0,
	"M1":   1,
	"M3":   2,
	"M10":  3,
	"M15":  4,
	"H1":   5,
	"H3":   6,
	"D1":   7,
	"W1":   8,
	"MON1": 9,
}

func (x PeriodType) String() string {
	return proto.EnumName(PeriodType_name, int32(x))
}
func (PeriodType) EnumDescriptor() ([]byte, []int) { return fileDescriptorBase, []int{1} }

// 合约代码
type Symbol struct {
	Exchange ExchangeType `protobuf:"varint,1,opt,name=exchange,proto3,enum=aproto.ExchangeType" json:"exchange"`
	Code     string       `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`
}

func (m *Symbol) Reset()                    { *m = Symbol{} }
func (m *Symbol) String() string            { return proto.CompactTextString(m) }
func (*Symbol) ProtoMessage()               {}
func (*Symbol) Descriptor() ([]byte, []int) { return fileDescriptorBase, []int{0} }

func (m *Symbol) GetExchange() ExchangeType {
	if m != nil {
		return m.Exchange
	}
	return ExchangeType_SHFE
}

func (m *Symbol) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

// K线
type Kline struct {
	// 时间
	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time"`
	// 开盘
	Open float64 `protobuf:"fixed64,2,opt,name=open,proto3" json:"open"`
	// 最高
	High float64 `protobuf:"fixed64,3,opt,name=high,proto3" json:"high"`
	// 最低
	Low float64 `protobuf:"fixed64,4,opt,name=low,proto3" json:"low"`
	// 收盘
	Close float64 `protobuf:"fixed64,5,opt,name=close,proto3" json:"close"`
	// 成交量
	Volume float64 `protobuf:"fixed64,6,opt,name=volume,proto3" json:"volume"`
	// 持仓量
	Position float64 `protobuf:"fixed64,7,opt,name=position,proto3" json:"position"`
}

func (m *Kline) Reset()                    { *m = Kline{} }
func (m *Kline) String() string            { return proto.CompactTextString(m) }
func (*Kline) ProtoMessage()               {}
func (*Kline) Descriptor() ([]byte, []int) { return fileDescriptorBase, []int{1} }

func (m *Kline) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Kline) GetOpen() float64 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *Kline) GetHigh() float64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *Kline) GetLow() float64 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *Kline) GetClose() float64 {
	if m != nil {
		return m.Close
	}
	return 0
}

func (m *Kline) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Kline) GetPosition() float64 {
	if m != nil {
		return m.Position
	}
	return 0
}

// KlineSeries K线序列
type KlineSeries struct {
	Symbol Symbol     `protobuf:"bytes,1,opt,name=symbol" json:"symbol"`
	Period PeriodType `protobuf:"varint,2,opt,name=period,proto3,enum=aproto.PeriodType" json:"period"`
	List   []Kline    `protobuf:"bytes,3,rep,name=list" json:"list"`
}

func (m *KlineSeries) Reset()                    { *m = KlineSeries{} }
func (m *KlineSeries) String() string            { return proto.CompactTextString(m) }
func (*KlineSeries) ProtoMessage()               {}
func (*KlineSeries) Descriptor() ([]byte, []int) { return fileDescriptorBase, []int{2} }

func (m *KlineSeries) GetSymbol() Symbol {
	if m != nil {
		return m.Symbol
	}
	return Symbol{}
}

func (m *KlineSeries) GetPeriod() PeriodType {
	if m != nil {
		return m.Period
	}
	return PeriodType_TICK
}

func (m *KlineSeries) GetList() []Kline {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*Symbol)(nil), "aproto.Symbol")
	proto.RegisterType((*Kline)(nil), "aproto.Kline")
	proto.RegisterType((*KlineSeries)(nil), "aproto.KlineSeries")
	proto.RegisterEnum("aproto.ExchangeType", ExchangeType_name, ExchangeType_value)
	proto.RegisterEnum("aproto.PeriodType", PeriodType_name, PeriodType_value)
}
func (m *Symbol) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Symbol) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Exchange != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBase(dAtA, i, uint64(m.Exchange))
	}
	if len(m.Code) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBase(dAtA, i, uint64(len(m.Code)))
		i += copy(dAtA[i:], m.Code)
	}
	return i, nil
}

func (m *Kline) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Kline) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Time != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBase(dAtA, i, uint64(m.Time))
	}
	if m.Open != 0 {
		dAtA[i] = 0x11
		i++
		i = encodeFixed64Base(dAtA, i, uint64(math.Float64bits(float64(m.Open))))
	}
	if m.High != 0 {
		dAtA[i] = 0x19
		i++
		i = encodeFixed64Base(dAtA, i, uint64(math.Float64bits(float64(m.High))))
	}
	if m.Low != 0 {
		dAtA[i] = 0x21
		i++
		i = encodeFixed64Base(dAtA, i, uint64(math.Float64bits(float64(m.Low))))
	}
	if m.Close != 0 {
		dAtA[i] = 0x29
		i++
		i = encodeFixed64Base(dAtA, i, uint64(math.Float64bits(float64(m.Close))))
	}
	if m.Volume != 0 {
		dAtA[i] = 0x31
		i++
		i = encodeFixed64Base(dAtA, i, uint64(math.Float64bits(float64(m.Volume))))
	}
	if m.Position != 0 {
		dAtA[i] = 0x39
		i++
		i = encodeFixed64Base(dAtA, i, uint64(math.Float64bits(float64(m.Position))))
	}
	return i, nil
}

func (m *KlineSeries) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KlineSeries) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintBase(dAtA, i, uint64(m.Symbol.Size()))
	n1, err := m.Symbol.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.Period != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBase(dAtA, i, uint64(m.Period))
	}
	if len(m.List) > 0 {
		for _, msg := range m.List {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintBase(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Base(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Base(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintBase(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Symbol) Size() (n int) {
	var l int
	_ = l
	if m.Exchange != 0 {
		n += 1 + sovBase(uint64(m.Exchange))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	return n
}

func (m *Kline) Size() (n int) {
	var l int
	_ = l
	if m.Time != 0 {
		n += 1 + sovBase(uint64(m.Time))
	}
	if m.Open != 0 {
		n += 9
	}
	if m.High != 0 {
		n += 9
	}
	if m.Low != 0 {
		n += 9
	}
	if m.Close != 0 {
		n += 9
	}
	if m.Volume != 0 {
		n += 9
	}
	if m.Position != 0 {
		n += 9
	}
	return n
}

func (m *KlineSeries) Size() (n int) {
	var l int
	_ = l
	l = m.Symbol.Size()
	n += 1 + l + sovBase(uint64(l))
	if m.Period != 0 {
		n += 1 + sovBase(uint64(m.Period))
	}
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovBase(uint64(l))
		}
	}
	return n
}

func sovBase(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBase(x uint64) (n int) {
	return sovBase(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Symbol) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
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
			return fmt.Errorf("proto: Symbol: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Symbol: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exchange", wireType)
			}
			m.Exchange = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Exchange |= (ExchangeType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBase
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
func (m *Kline) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
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
			return fmt.Errorf("proto: Kline: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Kline: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Open", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Open = float64(math.Float64frombits(v))
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field High", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.High = float64(math.Float64frombits(v))
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Low", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Low = float64(math.Float64frombits(v))
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Close", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Close = float64(math.Float64frombits(v))
		case 6:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Volume", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Volume = float64(math.Float64frombits(v))
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Position = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBase
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
func (m *KlineSeries) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
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
			return fmt.Errorf("proto: KlineSeries: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KlineSeries: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Symbol.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= (PeriodType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, Kline{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBase
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
func skipBase(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBase
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
					return 0, ErrIntOverflowBase
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
					return 0, ErrIntOverflowBase
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
				return 0, ErrInvalidLengthBase
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBase
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
				next, err := skipBase(dAtA[start:])
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
	ErrInvalidLengthBase = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBase   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("base.proto", fileDescriptorBase) }

var fileDescriptorBase = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcf, 0x8e, 0x12, 0x41,
	0x10, 0xc6, 0x69, 0x66, 0xa6, 0x81, 0x42, 0x49, 0xa5, 0xb2, 0x31, 0x13, 0x0f, 0x04, 0x39, 0x18,
	0xc2, 0x81, 0xec, 0x40, 0x7c, 0x01, 0x67, 0x1b, 0x31, 0x84, 0xd5, 0xcc, 0x6c, 0xa2, 0x59, 0x4f,
	0xc0, 0x76, 0x96, 0x4e, 0x06, 0x7a, 0xc2, 0xe0, 0x9f, 0xbd, 0xfb, 0x00, 0xbe, 0x81, 0xaf, 0xe3,
	0xd1, 0x47, 0x30, 0xf8, 0x22, 0xa6, 0x6b, 0x06, 0xdc, 0x53, 0x7d, 0xf5, 0xab, 0xaf, 0xaa, 0xab,
	0x0b, 0x60, 0xb5, 0x2c, 0xf4, 0x28, 0xdf, 0xdb, 0x83, 0x25, 0xb9, 0xe4, 0xd8, 0xbf, 0x06, 0x99,
	0x3e, 0x6c, 0x57, 0x36, 0xa3, 0x4b, 0x68, 0xea, 0x6f, 0xeb, 0xcd, 0x72, 0x77, 0xaf, 0x43, 0xd1,
	0x13, 0x83, 0xce, 0xf8, 0x62, 0x54, 0x9a, 0x46, 0xaa, 0xe2, 0x37, 0x0f, 0xb9, 0x4e, 0xce, 0x2e,
	0x22, 0xf0, 0xd7, 0xf6, 0x4e, 0x87, 0xf5, 0x9e, 0x18, 0xb4, 0x12, 0xd6, 0xfd, 0x9f, 0x02, 0x82,
	0x79, 0x66, 0x76, 0x5c, 0x3d, 0x98, 0x6d, 0x39, 0xcb, 0x4b, 0x58, 0x3b, 0x66, 0x73, 0xbd, 0xe3,
	0x0e, 0x91, 0xb0, 0x76, 0x6c, 0x63, 0xee, 0x37, 0xa1, 0x57, 0x32, 0xa7, 0x09, 0xc1, 0xcb, 0xec,
	0xd7, 0xd0, 0x67, 0xe4, 0x24, 0x5d, 0x40, 0xb0, 0xce, 0x6c, 0xa1, 0xc3, 0x80, 0x59, 0x99, 0xd0,
	0x33, 0x90, 0x5f, 0x6c, 0xf6, 0x79, 0xab, 0x43, 0xc9, 0xb8, 0xca, 0xe8, 0x39, 0x34, 0x73, 0x5b,
	0x98, 0x83, 0xb1, 0xbb, 0xb0, 0xc1, 0x95, 0x73, 0xde, 0xff, 0x2e, 0xa0, 0xcd, 0x1b, 0xa6, 0x7a,
	0x6f, 0x74, 0x41, 0x2f, 0x41, 0x16, 0x7c, 0x01, 0xde, 0xb4, 0x3d, 0xee, 0x9c, 0x7e, 0x5d, 0xde,
	0x25, 0xa9, 0xaa, 0x34, 0x04, 0x99, 0xeb, 0xbd, 0xb1, 0x77, 0xbc, 0x7d, 0x67, 0x4c, 0x27, 0xdf,
	0x7b, 0xa6, 0x7c, 0x9b, 0xca, 0x41, 0x2f, 0xc0, 0xcf, 0x4c, 0x71, 0x08, 0xbd, 0x9e, 0x37, 0x68,
	0x8f, 0x9f, 0x9e, 0x9c, 0xfc, 0x6c, 0xc2, 0xa5, 0xe1, 0x27, 0x78, 0xf2, 0xf8, 0xac, 0xd4, 0x04,
	0x3f, 0x9d, 0x4d, 0x15, 0xd6, 0x9c, 0x8a, 0x6f, 0x63, 0x85, 0x82, 0x1a, 0xe0, 0x5d, 0xc5, 0x0a,
	0xeb, 0xd4, 0x82, 0x20, 0x9e, 0x4e, 0xd5, 0x47, 0xf4, 0x1c, 0x4b, 0x53, 0x85, 0x3e, 0x8b, 0x5b,
	0x85, 0x81, 0x13, 0xb3, 0xf9, 0x1b, 0x94, 0xae, 0x71, 0x36, 0x9f, 0x2a, 0x6c, 0x0c, 0x57, 0x00,
	0xff, 0xb7, 0x72, 0xfc, 0xe6, 0x6d, 0x3c, 0xc7, 0x1a, 0x49, 0xa8, 0x2f, 0x22, 0x14, 0x1c, 0x27,
	0x58, 0x77, 0xad, 0x8b, 0xe8, 0xb2, 0x9c, 0xba, 0x88, 0x5e, 0xa1, 0xef, 0x2a, 0xb3, 0x08, 0x03,
	0x8e, 0x13, 0x94, 0x2e, 0x5e, 0x45, 0xd8, 0x70, 0xf1, 0x43, 0x84, 0x4d, 0x37, 0x6b, 0xf1, 0xee,
	0x3a, 0xc2, 0xd6, 0x6b, 0xfc, 0x75, 0xec, 0x8a, 0xdf, 0xc7, 0xae, 0xf8, 0x73, 0xec, 0x8a, 0x1f,
	0x7f, 0xbb, 0xb5, 0x95, 0xe4, 0x5f, 0x4e, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x18, 0x21,
	0xed, 0x68, 0x02, 0x00, 0x00,
}
