// Code generated by protoc-gen-go. DO NOT EDIT.
// source: goshare/trade.proto

package goshare

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

// 成交类型
type TradeType int32

const (
	// 普通成交
	TradeType_TT_NORMAL TradeType = 0
	// 期权执行
	TradeType_TT_OPTIONS_EXECUTION TradeType = 1
	// OTC成交
	TradeType_TT_OTC TradeType = 2
	// 期转现衍生成交
	TradeType_TT_EFP_DERIVED TradeType = 3
	// 组合衍生成交
	TradeType_TT_COMBINATION_DERIVED TradeType = 4
	// 交割平仓
	TradeType_TT_DELIVERY_CLOSE TradeType = 5
	// 结算衍生
	TradeType_TT_SETTLEMENT_DERIVED TradeType = 6
	// 分红股票
	TradeType_TT_PLACEMENT_DERIVED TradeType = 7
	// 回购
	TradeType_TT_REPURCHASE TradeType = 8
)

var TradeType_name = map[int32]string{
	0: "TT_NORMAL",
	1: "TT_OPTIONS_EXECUTION",
	2: "TT_OTC",
	3: "TT_EFP_DERIVED",
	4: "TT_COMBINATION_DERIVED",
	5: "TT_DELIVERY_CLOSE",
	6: "TT_SETTLEMENT_DERIVED",
	7: "TT_PLACEMENT_DERIVED",
	8: "TT_REPURCHASE",
}

var TradeType_value = map[string]int32{
	"TT_NORMAL":              0,
	"TT_OPTIONS_EXECUTION":   1,
	"TT_OTC":                 2,
	"TT_EFP_DERIVED":         3,
	"TT_COMBINATION_DERIVED": 4,
	"TT_DELIVERY_CLOSE":      5,
	"TT_SETTLEMENT_DERIVED":  6,
	"TT_PLACEMENT_DERIVED":   7,
	"TT_REPURCHASE":          8,
}

func (x TradeType) String() string {
	return proto.EnumName(TradeType_name, int32(x))
}

func (TradeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e39337ecc2a5265d, []int{0}
}

// 成交
type Trade struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange"`
	Symbol               string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol"`
	Product              string   `protobuf:"bytes,3,opt,name=product,proto3" json:"product"`
	Direction            int32    `protobuf:"varint,5,opt,name=direction,proto3" json:"direction"`
	Offset               int32    `protobuf:"varint,6,opt,name=offset,proto3" json:"offset"`
	Price                float64  `protobuf:"fixed64,7,opt,name=price,proto3" json:"price"`
	Volume               int32    `protobuf:"varint,8,opt,name=volume,proto3" json:"volume"`
	UserId               string   `protobuf:"bytes,9,opt,name=user_id,json=userId,proto3" json:"userId"`
	TaId                 string   `protobuf:"bytes,10,opt,name=ta_id,json=taId,proto3" json:"taId"`
	BuId                 string   `protobuf:"bytes,11,opt,name=bu_id,json=buId,proto3" json:"buId"`
	TradeId              string   `protobuf:"bytes,12,opt,name=trade_id,json=tradeId,proto3" json:"tradeId"`
	TradedTime           int64    `protobuf:"varint,13,opt,name=traded_time,json=tradedTime,proto3" json:"tradedTime"`
	TradedTradingDay     int32    `protobuf:"varint,14,opt,name=traded_trading_day,json=tradedTradingDay,proto3" json:"tradedTradingDay"`
	FrontId              int32    `protobuf:"varint,15,opt,name=front_id,json=frontId,proto3" json:"frontId"`
	SessionId            int32    `protobuf:"varint,16,opt,name=session_id,json=sessionId,proto3" json:"sessionId"`
	OrderRef             string   `protobuf:"bytes,17,opt,name=order_ref,json=orderRef,proto3" json:"orderRef"`
	UserName             string   `protobuf:"bytes,18,opt,name=user_name,json=userName,proto3" json:"userName"`
	Branch               string   `protobuf:"bytes,19,opt,name=branch,proto3" json:"branch"`
	BranchName           string   `protobuf:"bytes,20,opt,name=branch_name,json=branchName,proto3" json:"branchName"`
	TradeType            int32    `protobuf:"varint,21,opt,name=trade_type,json=tradeType,proto3" json:"tradeType"`
	ExchangeOrderId      string   `protobuf:"bytes,22,opt,name=exchange_order_id,json=exchangeOrderId,proto3" json:"exchangeOrderId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Trade) Reset()         { *m = Trade{} }
func (m *Trade) String() string { return proto.CompactTextString(m) }
func (*Trade) ProtoMessage()    {}
func (*Trade) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39337ecc2a5265d, []int{0}
}

func (m *Trade) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trade.Unmarshal(m, b)
}
func (m *Trade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trade.Marshal(b, m, deterministic)
}
func (m *Trade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trade.Merge(m, src)
}
func (m *Trade) XXX_Size() int {
	return xxx_messageInfo_Trade.Size(m)
}
func (m *Trade) XXX_DiscardUnknown() {
	xxx_messageInfo_Trade.DiscardUnknown(m)
}

var xxx_messageInfo_Trade proto.InternalMessageInfo

func (m *Trade) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *Trade) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Trade) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *Trade) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *Trade) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Trade) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Trade) GetVolume() int32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Trade) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Trade) GetTaId() string {
	if m != nil {
		return m.TaId
	}
	return ""
}

func (m *Trade) GetBuId() string {
	if m != nil {
		return m.BuId
	}
	return ""
}

func (m *Trade) GetTradeId() string {
	if m != nil {
		return m.TradeId
	}
	return ""
}

func (m *Trade) GetTradedTime() int64 {
	if m != nil {
		return m.TradedTime
	}
	return 0
}

func (m *Trade) GetTradedTradingDay() int32 {
	if m != nil {
		return m.TradedTradingDay
	}
	return 0
}

func (m *Trade) GetFrontId() int32 {
	if m != nil {
		return m.FrontId
	}
	return 0
}

func (m *Trade) GetSessionId() int32 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *Trade) GetOrderRef() string {
	if m != nil {
		return m.OrderRef
	}
	return ""
}

func (m *Trade) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Trade) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *Trade) GetBranchName() string {
	if m != nil {
		return m.BranchName
	}
	return ""
}

func (m *Trade) GetTradeType() int32 {
	if m != nil {
		return m.TradeType
	}
	return 0
}

func (m *Trade) GetExchangeOrderId() string {
	if m != nil {
		return m.ExchangeOrderId
	}
	return ""
}

func init() {
	proto.RegisterEnum("goshare.TradeType", TradeType_name, TradeType_value)
	proto.RegisterType((*Trade)(nil), "goshare.Trade")
}

func init() { proto.RegisterFile("goshare/trade.proto", fileDescriptor_e39337ecc2a5265d) }

var fileDescriptor_e39337ecc2a5265d = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x93, 0xdd, 0x6e, 0x1a, 0x3d,
	0x10, 0x86, 0xbf, 0x4d, 0xc2, 0xdf, 0xe4, 0x4b, 0xb2, 0x4c, 0x7e, 0xea, 0xa4, 0x8d, 0x8a, 0x7a,
	0x84, 0xa2, 0x28, 0x1c, 0xf4, 0x0a, 0x08, 0xb8, 0xea, 0x4a, 0x04, 0xd0, 0xe2, 0x44, 0x6d, 0x4f,
	0x2c, 0x2f, 0x36, 0xb0, 0x6a, 0x76, 0xbd, 0xf2, 0x2e, 0x55, 0xb9, 0xca, 0xde, 0x41, 0xaf, 0xa5,
	0xb2, 0x0d, 0x54, 0x3d, 0x9b, 0xf7, 0x79, 0x67, 0xec, 0x99, 0xd9, 0x35, 0x9c, 0x2f, 0x75, 0xb9,
	0x12, 0x46, 0xf5, 0x2a, 0x23, 0xa4, 0x7a, 0x28, 0x8c, 0xae, 0x34, 0x36, 0xb6, 0xf0, 0xc3, 0xef,
	0x23, 0xa8, 0x31, 0x6b, 0xe0, 0x0d, 0x34, 0xd5, 0xcf, 0xf9, 0x4a, 0xe4, 0x4b, 0x45, 0x82, 0x4e,
	0xd0, 0x6d, 0xc5, 0x7b, 0x8d, 0x57, 0x50, 0x2f, 0x37, 0x59, 0xa2, 0x5f, 0xc9, 0x81, 0x73, 0xb6,
	0x0a, 0x09, 0x34, 0x0a, 0xa3, 0xe5, 0x7a, 0x5e, 0x91, 0x43, 0x67, 0xec, 0x24, 0xbe, 0x83, 0x96,
	0x4c, 0x8d, 0x9a, 0x57, 0xa9, 0xce, 0x49, 0xad, 0x13, 0x74, 0x6b, 0xf1, 0x5f, 0x60, 0xcf, 0xd3,
	0x8b, 0x45, 0xa9, 0x2a, 0x52, 0x77, 0xd6, 0x56, 0xe1, 0x05, 0xd4, 0x0a, 0x93, 0xce, 0x15, 0x69,
	0x74, 0x82, 0x6e, 0x10, 0x7b, 0x61, 0xb3, 0x7f, 0xe8, 0xd7, 0x75, 0xa6, 0x48, 0xd3, 0x67, 0x7b,
	0x85, 0x6f, 0xa0, 0xb1, 0x2e, 0x95, 0xe1, 0xa9, 0x24, 0x2d, 0xdf, 0x96, 0x95, 0x91, 0xc4, 0x73,
	0xa8, 0x55, 0xc2, 0x62, 0x70, 0xf8, 0xa8, 0x12, 0x1e, 0x26, 0x6b, 0x0b, 0x8f, 0x3d, 0x4c, 0xd6,
	0x91, 0xc4, 0x6b, 0x68, 0xba, 0xb5, 0x58, 0xfe, 0xbf, 0x9f, 0xc0, 0xe9, 0x48, 0xe2, 0x7b, 0x38,
	0x76, 0xa1, 0xe4, 0x55, 0x9a, 0x29, 0x72, 0xd2, 0x09, 0xba, 0x87, 0x31, 0x78, 0xc4, 0xd2, 0x4c,
	0xe1, 0x3d, 0xe0, 0x2e, 0xc1, 0x08, 0x99, 0xe6, 0x4b, 0x2e, 0xc5, 0x86, 0x9c, 0xba, 0x16, 0xc3,
	0x6d, 0x9e, 0x37, 0x86, 0x62, 0x63, 0x6f, 0x5a, 0x18, 0x9d, 0x57, 0xf6, 0xa6, 0x33, 0x97, 0xd3,
	0x70, 0x3a, 0x92, 0x78, 0x0b, 0x50, 0xaa, 0xb2, 0x4c, 0x75, 0x6e, 0xcd, 0xd0, 0x2f, 0x6b, 0x4b,
	0x22, 0x89, 0x6f, 0xa1, 0xa5, 0x8d, 0x54, 0x86, 0x1b, 0xb5, 0x20, 0x6d, 0xff, 0x65, 0x1c, 0x88,
	0xd5, 0xc2, 0x9a, 0x6e, 0x07, 0xb9, 0xc8, 0x14, 0x41, 0x6f, 0x5a, 0x30, 0x16, 0x99, 0x5b, 0x5c,
	0x62, 0x44, 0x3e, 0x5f, 0x91, 0x73, 0xbf, 0x1f, 0xaf, 0xec, 0x68, 0x3e, 0xf2, 0x65, 0x17, 0xce,
	0x04, 0x8f, 0x5c, 0xe1, 0x2d, 0xf8, 0x41, 0x79, 0xb5, 0x29, 0x14, 0xb9, 0xf4, 0x1d, 0x39, 0xc2,
	0x36, 0x85, 0xc2, 0x3b, 0x68, 0xef, 0x7e, 0x0d, 0xee, 0x5b, 0x4b, 0x25, 0xb9, 0x72, 0xa7, 0x9c,
	0xed, 0x8c, 0x89, 0xe5, 0x91, 0xbc, 0xfb, 0x15, 0x40, 0x8b, 0xed, 0x2b, 0x4f, 0xa0, 0xc5, 0x18,
	0x1f, 0x4f, 0xe2, 0xa7, 0xfe, 0x28, 0xfc, 0x0f, 0x09, 0x5c, 0x30, 0xc6, 0x27, 0x53, 0x16, 0x4d,
	0xc6, 0x33, 0x4e, 0xbf, 0xd0, 0xc1, 0xb3, 0x0d, 0xc3, 0x00, 0x01, 0xea, 0xd6, 0x61, 0x83, 0xf0,
	0x00, 0x11, 0x4e, 0x19, 0xe3, 0xf4, 0xd3, 0x94, 0x0f, 0x69, 0x1c, 0xbd, 0xd0, 0x61, 0x78, 0x88,
	0x37, 0x70, 0xc5, 0x18, 0x1f, 0x4c, 0x9e, 0x1e, 0xa3, 0x71, 0xdf, 0xd6, 0xec, 0xbd, 0x23, 0xbc,
	0x84, 0x36, 0x63, 0x7c, 0x48, 0x47, 0xd1, 0x0b, 0x8d, 0xbf, 0xf2, 0xc1, 0x68, 0x32, 0xa3, 0x61,
	0x0d, 0xaf, 0xe1, 0x92, 0x31, 0x3e, 0xa3, 0x8c, 0x8d, 0xe8, 0x13, 0x1d, 0xb3, 0x7d, 0x45, 0x7d,
	0xdb, 0xc7, 0x74, 0xd4, 0x1f, 0xfc, 0xeb, 0x34, 0xb0, 0x0d, 0x27, 0x8c, 0xf1, 0x98, 0x4e, 0x9f,
	0xe3, 0xc1, 0xe7, 0xfe, 0x8c, 0x86, 0xcd, 0xc7, 0x87, 0x6f, 0xf7, 0xcb, 0xb4, 0x5a, 0xad, 0x93,
	0x87, 0xb9, 0xce, 0x7a, 0x59, 0x9a, 0x2b, 0x23, 0x5e, 0x8d, 0x2a, 0x7b, 0xbb, 0x87, 0x56, 0x7c,
	0x5f, 0xf6, 0x8a, 0x64, 0x27, 0x93, 0xba, 0x7b, 0x72, 0x1f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x02, 0x9c, 0x1c, 0xad, 0x89, 0x03, 0x00, 0x00,
}
