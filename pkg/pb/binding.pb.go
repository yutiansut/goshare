// Code generated by protoc-gen-go. DO NOT EDIT.
// source: binding.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	binding.proto
	branch.proto
	common.proto
	ctp.proto
	dcenter.proto
	goshare.proto
	market_data.proto
	message.proto
	order.proto
	others.proto
	position.proto
	resource.proto
	strategy.proto
	tahub.proto
	trade_report.proto
	trading.proto
	ucenter.proto
	user.proto

It has these top-level messages:
	BranchTreeNode
	Branch
	Symbol
	SymbolList
	ProductID
	BrokerRoute
	BrokerRouteList
	InstrumentInfo
	SimpleTimePeriod
	MarketTimeRule
	ProductInfo
	TradingInstrument
	TradingInstrumentList
	ProductInfoList
	ReqSetTradingInstrumentList
	ReqSetProductInfoList
	ReqUpdateTIOpenDate
	ReqUpdateTIOpenDateList
	AccountMoneySummary
	AccountMoneySummaryList
	MoneyTransferRecord
	MTRList
	SSEStockOption
	OptionMonth
	OptionMonthList
	EmptyRequest
	EmptyResponse
	CommonRequest
	CommonResponse
	CTPOrderID
	CTPOrderField
	CTPOrderList
	CTPOrderLink
	CTPCancelOrderRequest
	CTPOnRtnCancelOrder
	CTPTradeField
	CTPTradeReportList
	CTPPositionSummation
	CTPPosition
	CTPPositionList
	CTPCloseTradeRecord
	CTPCloseTradeRecordList
	CTPMarginRate
	CTPMarginItem
	CTPCommissionRateItem
	CTPCommissionRate
	CTPTradingUnit
	CTPForceCloseTimeRule
	CTPTradingUnitOption
	MarginCheckRule
	SettlementUnitOption
	CTPTradingAccount
	CTPTradingAccountSnapshot
	CTPRspInfo
	CtpReqQryAccountRegister
	CtpRspQryAccountRegister
	CtpReqAuthencate
	CtpRspAuthencate
	CtpReqUserLogin
	CtpRspUserLogin
	CtpReqConnect
	CtpOnFrontConnected
	CtpOnFrontDisconnected
	CtpReqSettlementInfoConfirm
	CtpOnRspSettlementInfoConfirm
	CtpReqQryInvestor
	CtpRspQryInvestor
	CtpReqQryTransferBank
	CtpRspQryTransferBank
	CtpReqTransfer
	CtpRspTransfer
	CTPReqSubscribeMarketData
	CTPRspSusbcribeMarketData
	CTPReqUnSusbibeMarketData
	CTPRspUnSusbibeMarketData
	CTPOnRspOrderInsert
	CTPInstrumentField
	CTPOnRspQryInstrument
	ReqGetTradingInstrumentList
	RspGetTradingInstrumentList
	ReqGetTradingInstrument
	RspGetTradingInstrument
	ReqSetTradingInstrument
	RspSetTradingInstrument
	ReqSubscribe
	RspSubscribe
	ReqUnSubscribe
	RspUnSubscribe
	ReqSaveKline
	RspSaveKline
	ReqCombineSubscribe
	RspCombineSubscribe
	SymbolCacheSummary
	CacheSummary
	ReqGetDCenterInfo
	RspGetDCenterInfo
	ReqSSEStockOptionList
	RspSSEStockOptionList
	OrderBook
	MarketDataSnapshot
	MdsList
	OptionTMarket
	SimpleTickForTQuote
	OptionTQuoteItem
	OptionTQuoteItemList
	Kline
	KlineSeries
	ReqSubscribeMarketData
	RspSubscribeMarketData
	RtnMarketDataUpdate
	TickSeries
	SimpleTick
	SimpleTickSeries
	Message
	MessageList
	StreamMessage
	OrderID
	Order
	OrderList
	JointOrder
	OrderLink
	CancelOrderRequest
	OnRtnCancelOrder
	ConditionOrder
	TradeCmd
	DemoOrder
	NetInAmountDetail
	RealtimeMoneyTrendItem
	RealtimeMoneyTrendItemList
	PositionSummation
	Position
	PositionList
	ReqForceClosePosition
	Strategy
	StrategyList
	ReqDeleteStrategy
	ReqGetStrategy
	ReqGetAccountList
	RspGetAccountList
	ReqAddAccount
	RspAddAccount
	ReqDeleteAccount
	RspDeleteAccount
	ReqGetTradingRouteList
	RspGetTradingRouteList
	ReqCreateTradingRoute
	RspCreateTradingRoute
	ReqDeleteTradingRoute
	RspDeleteTradingRoute
	TradeReport
	TradeReportList
	TradeReportSession
	TradingRoute
	TradingRouteList
	TradingAccount
	TradingAccountList
	TradingAccountSummary
	RspTradingAccountLogin
	RspOrderInsert
	ReqCreateBranch
	ReqDeleteBranch
	ReqCreateUser
	ReqDeleteUser
	ReqCheckAPIPermission
	RspCheckAPIPermission
	ReqCheckResourcePermission
	RspCheckResourcePermission
	ReqUserLogin
	RspUserLogin
	User
	UserList
	Permission
	PermissionTreeNode
	PermissionList
	UserRole
	UserSession
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BindingMessageType int32

const (
	BindingMessageType_CTP_CREATE_API                  BindingMessageType = 0
	BindingMessageType_CTP_RSP_CREATE_API              BindingMessageType = 1
	BindingMessageType_CTP_REQ_CALL_INIT               BindingMessageType = 2
	BindingMessageType_CTP_RSP_CALL_INIT               BindingMessageType = 3
	BindingMessageType_CTP_REQ_USER_LOGIN              BindingMessageType = 4
	BindingMessageType_CTP_RSP_USER_LOGIN              BindingMessageType = 5
	BindingMessageType_CTP_REQ_SUBSCRIBE_MARKET_DATA   BindingMessageType = 6
	BindingMessageType_CTP_RSP_SUBSCRIBE_MARKET_DATA   BindingMessageType = 7
	BindingMessageType_CTP_ON_FRONT_CONNECTED          BindingMessageType = 8
	BindingMessageType_CTP_ON_FRONT_DISCONNECTED       BindingMessageType = 9
	BindingMessageType_CTP_ON_RTN_DEPTH_MARKET_DATA    BindingMessageType = 10
	BindingMessageType_CTP_ON_RTN_ORDER                BindingMessageType = 11
	BindingMessageType_CTP_ON_RTN_TRADE                BindingMessageType = 12
	BindingMessageType_CTP_DELETE_API                  BindingMessageType = 13
	BindingMessageType_CTP_REQ_INSERT_ORDER            BindingMessageType = 14
	BindingMessageType_CTP_REQ_CANCEL_ORDER            BindingMessageType = 15
	BindingMessageType_CTP_REQ_QRY_TRADING_ACCOUNT     BindingMessageType = 16
	BindingMessageType_CTP_REQ_QRY_POSITION_DETAIL     BindingMessageType = 17
	BindingMessageType_CTP_REQ_QRY_COMMISSION_RATE     BindingMessageType = 18
	BindingMessageType_CTP_REQ_QRY_MARGIN_RATE         BindingMessageType = 19
	BindingMessageType_CTP_REQ_CONNECT                 BindingMessageType = 20
	BindingMessageType_CTP_REQ_AUTHENTICATE            BindingMessageType = 21
	BindingMessageType_CTP_REQ_SETTLEMENT_INFO_CONFIRM BindingMessageType = 22
	BindingMessageType_CTP_REQ_QRY_INVESTOR            BindingMessageType = 23
	BindingMessageType_CTP_REQ_QRY_TRANSFER_BANK       BindingMessageType = 24
	BindingMessageType_CTP_REQ_TRANSFER                BindingMessageType = 25
	BindingMessageType_CTP_REQ_QRY_ACCOUNT_REGISTER    BindingMessageType = 26
	BindingMessageType_CTP_REQ_UNSUBSCRIBE_MARKET_DATA BindingMessageType = 27
	BindingMessageType_CTP_RSP_UNSUBSCRIBE_MARKET_DATA BindingMessageType = 28
	BindingMessageType_CTP_REQ_DISCONNECT              BindingMessageType = 29
	BindingMessageType_CTP_ON_RSP_ORDER_INSERT         BindingMessageType = 30
	BindingMessageType_CTP_REQ_QRY_INSTRUMENT          BindingMessageType = 31
	BindingMessageType_CTP_ON_RSP_QRY_INSTRUMENT       BindingMessageType = 32
)

var BindingMessageType_name = map[int32]string{
	0:  "CTP_CREATE_API",
	1:  "CTP_RSP_CREATE_API",
	2:  "CTP_REQ_CALL_INIT",
	3:  "CTP_RSP_CALL_INIT",
	4:  "CTP_REQ_USER_LOGIN",
	5:  "CTP_RSP_USER_LOGIN",
	6:  "CTP_REQ_SUBSCRIBE_MARKET_DATA",
	7:  "CTP_RSP_SUBSCRIBE_MARKET_DATA",
	8:  "CTP_ON_FRONT_CONNECTED",
	9:  "CTP_ON_FRONT_DISCONNECTED",
	10: "CTP_ON_RTN_DEPTH_MARKET_DATA",
	11: "CTP_ON_RTN_ORDER",
	12: "CTP_ON_RTN_TRADE",
	13: "CTP_DELETE_API",
	14: "CTP_REQ_INSERT_ORDER",
	15: "CTP_REQ_CANCEL_ORDER",
	16: "CTP_REQ_QRY_TRADING_ACCOUNT",
	17: "CTP_REQ_QRY_POSITION_DETAIL",
	18: "CTP_REQ_QRY_COMMISSION_RATE",
	19: "CTP_REQ_QRY_MARGIN_RATE",
	20: "CTP_REQ_CONNECT",
	21: "CTP_REQ_AUTHENTICATE",
	22: "CTP_REQ_SETTLEMENT_INFO_CONFIRM",
	23: "CTP_REQ_QRY_INVESTOR",
	24: "CTP_REQ_QRY_TRANSFER_BANK",
	25: "CTP_REQ_TRANSFER",
	26: "CTP_REQ_QRY_ACCOUNT_REGISTER",
	27: "CTP_REQ_UNSUBSCRIBE_MARKET_DATA",
	28: "CTP_RSP_UNSUBSCRIBE_MARKET_DATA",
	29: "CTP_REQ_DISCONNECT",
	30: "CTP_ON_RSP_ORDER_INSERT",
	31: "CTP_REQ_QRY_INSTRUMENT",
	32: "CTP_ON_RSP_QRY_INSTRUMENT",
}
var BindingMessageType_value = map[string]int32{
	"CTP_CREATE_API":                  0,
	"CTP_RSP_CREATE_API":              1,
	"CTP_REQ_CALL_INIT":               2,
	"CTP_RSP_CALL_INIT":               3,
	"CTP_REQ_USER_LOGIN":              4,
	"CTP_RSP_USER_LOGIN":              5,
	"CTP_REQ_SUBSCRIBE_MARKET_DATA":   6,
	"CTP_RSP_SUBSCRIBE_MARKET_DATA":   7,
	"CTP_ON_FRONT_CONNECTED":          8,
	"CTP_ON_FRONT_DISCONNECTED":       9,
	"CTP_ON_RTN_DEPTH_MARKET_DATA":    10,
	"CTP_ON_RTN_ORDER":                11,
	"CTP_ON_RTN_TRADE":                12,
	"CTP_DELETE_API":                  13,
	"CTP_REQ_INSERT_ORDER":            14,
	"CTP_REQ_CANCEL_ORDER":            15,
	"CTP_REQ_QRY_TRADING_ACCOUNT":     16,
	"CTP_REQ_QRY_POSITION_DETAIL":     17,
	"CTP_REQ_QRY_COMMISSION_RATE":     18,
	"CTP_REQ_QRY_MARGIN_RATE":         19,
	"CTP_REQ_CONNECT":                 20,
	"CTP_REQ_AUTHENTICATE":            21,
	"CTP_REQ_SETTLEMENT_INFO_CONFIRM": 22,
	"CTP_REQ_QRY_INVESTOR":            23,
	"CTP_REQ_QRY_TRANSFER_BANK":       24,
	"CTP_REQ_TRANSFER":                25,
	"CTP_REQ_QRY_ACCOUNT_REGISTER":    26,
	"CTP_REQ_UNSUBSCRIBE_MARKET_DATA": 27,
	"CTP_RSP_UNSUBSCRIBE_MARKET_DATA": 28,
	"CTP_REQ_DISCONNECT":              29,
	"CTP_ON_RSP_ORDER_INSERT":         30,
	"CTP_REQ_QRY_INSTRUMENT":          31,
	"CTP_ON_RSP_QRY_INSTRUMENT":       32,
}

func (x BindingMessageType) String() string {
	return proto.EnumName(BindingMessageType_name, int32(x))
}
func (BindingMessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("pb.BindingMessageType", BindingMessageType_name, BindingMessageType_value)
}

func init() { proto.RegisterFile("binding.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x4f, 0x6f, 0x13, 0x3d,
	0x10, 0xc6, 0xdf, 0xb7, 0x40, 0x0b, 0x86, 0xb6, 0x53, 0xb7, 0x4d, 0xff, 0xa4, 0x21, 0x45, 0xdc,
	0x38, 0x70, 0xe1, 0x13, 0x38, 0xde, 0x49, 0x6a, 0x75, 0xd7, 0xde, 0x8c, 0x67, 0x91, 0x38, 0x59,
	0x44, 0x44, 0x55, 0x2f, 0x6d, 0x44, 0xb9, 0xf0, 0xbd, 0xf8, 0x80, 0xc8, 0xc9, 0x7a, 0xe3, 0x44,
	0xf4, 0xfa, 0x3c, 0xbf, 0x99, 0xf5, 0x3c, 0x33, 0x2b, 0xf6, 0x67, 0xf7, 0x0f, 0x3f, 0xee, 0x1f,
	0xee, 0x3e, 0x2f, 0x7e, 0x3e, 0xfe, 0x7a, 0x94, 0x3b, 0x8b, 0xd9, 0xa7, 0x3f, 0x7b, 0x42, 0x8e,
	0x56, 0x6a, 0x35, 0x7f, 0x7a, 0xfa, 0x7e, 0x37, 0xe7, 0xdf, 0x8b, 0xb9, 0x94, 0xe2, 0x40, 0x73,
	0x1d, 0x34, 0xa1, 0x62, 0x0c, 0xaa, 0x36, 0xf0, 0x9f, 0xec, 0x09, 0x19, 0x35, 0xf2, 0x1b, 0xfa,
	0xff, 0xf2, 0x54, 0x1c, 0x2d, 0x75, 0x9c, 0x06, 0xad, 0xca, 0x32, 0x18, 0x6b, 0x18, 0x76, 0x3a,
	0x39, 0xe2, 0x9d, 0xfc, 0xa2, 0xeb, 0x82, 0xd3, 0xd0, 0x78, 0xa4, 0x50, 0xba, 0x89, 0xb1, 0xf0,
	0x32, 0xef, 0x9e, 0xe9, 0xaf, 0xe4, 0x07, 0x31, 0x48, 0xbc, 0x6f, 0x46, 0x5e, 0x93, 0x19, 0x61,
	0xa8, 0x14, 0xdd, 0x22, 0x87, 0x42, 0xb1, 0x82, 0xdd, 0x0e, 0xf1, 0xf5, 0x33, 0xc8, 0x9e, 0xbc,
	0x14, 0xbd, 0x88, 0x38, 0x1b, 0xc6, 0xe4, 0x2c, 0x07, 0xed, 0xac, 0x45, 0xcd, 0x58, 0xc0, 0x6b,
	0x39, 0x10, 0x17, 0x1b, 0x5e, 0x61, 0xfc, 0xda, 0x7e, 0x23, 0xaf, 0xc5, 0x55, 0x6b, 0x13, 0xdb,
	0x50, 0x60, 0xcd, 0x37, 0x1b, 0xcd, 0x85, 0x3c, 0x11, 0x90, 0x11, 0x8e, 0x0a, 0x24, 0x78, 0xbb,
	0xa5, 0x32, 0xa9, 0x02, 0xe1, 0x5d, 0x0a, 0xb6, 0xc0, 0x12, 0xdb, 0x00, 0xf7, 0xe5, 0xb9, 0x38,
	0x49, 0x23, 0x1a, 0xeb, 0x91, 0xb8, 0xed, 0x71, 0x90, 0x3b, 0x5a, 0x59, 0x8d, 0x65, 0xeb, 0x1c,
	0xca, 0xa1, 0xe8, 0x27, 0x67, 0x4a, 0xdf, 0x96, 0xed, 0x8d, 0x9d, 0x04, 0xa5, 0xb5, 0x6b, 0x2c,
	0x03, 0x6c, 0x03, 0xb5, 0xf3, 0x86, 0x8d, 0x8b, 0x03, 0xb0, 0x32, 0x25, 0x1c, 0x6d, 0x03, 0xda,
	0x55, 0x95, 0xf1, 0x3e, 0x22, 0xa4, 0x18, 0x41, 0xca, 0xbe, 0x38, 0xcb, 0x81, 0x4a, 0xd1, 0xc4,
	0xb4, 0xe6, 0xb1, 0x3c, 0x16, 0x87, 0xdd, 0xcb, 0x56, 0x61, 0xc1, 0x49, 0xfe, 0x5c, 0xd5, 0xf0,
	0x0d, 0x5a, 0x36, 0x3a, 0xe2, 0xa7, 0xf2, 0xa3, 0x18, 0x76, 0x5b, 0x44, 0xe6, 0x12, 0x2b, 0xb4,
	0x1c, 0x8c, 0x1d, 0xbb, 0x58, 0x3e, 0x36, 0x54, 0x41, 0x2f, 0x2f, 0x8f, 0x1f, 0x34, 0xf6, 0x2b,
	0x7a, 0x76, 0x04, 0x67, 0x69, 0x45, 0xd9, 0xb4, 0xd6, 0x8f, 0x91, 0xc2, 0x48, 0xd9, 0x5b, 0x38,
	0x4f, 0x51, 0x47, 0x3b, 0x59, 0x70, 0x91, 0x16, 0x97, 0x8a, 0xda, 0x68, 0x02, 0xe1, 0xc4, 0x78,
	0x46, 0x82, 0xcb, 0xfc, 0x55, 0x8d, 0xfd, 0xf7, 0xe9, 0xf4, 0x3b, 0x28, 0x1e, 0xe6, 0x33, 0xd0,
	0x55, 0x7e, 0xd5, 0xeb, 0xf3, 0x81, 0x41, 0xca, 0x30, 0x86, 0xea, 0xeb, 0xd5, 0xf2, 0xda, 0x1d,
	0xc3, 0xfb, 0x74, 0x94, 0xeb, 0x79, 0x3d, 0x53, 0x13, 0x83, 0x81, 0x61, 0x76, 0x94, 0xb1, 0x70,
	0xcb, 0xbe, 0x9e, 0xed, 0x2e, 0xff, 0xe0, 0x2f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x82, 0x69,
	0x5f, 0xe0, 0xd2, 0x03, 0x00, 0x00,
}
