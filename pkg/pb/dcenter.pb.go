// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dcenter.proto

package pb

import proto "github.com/golang/protobuf/proto"
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

type ReqGetTradingInstrumentList struct {
	Exchange ExchangeType `protobuf:"varint,1,opt,name=exchange,enum=pb.ExchangeType" json:"exchange,omitempty"`
}

func (m *ReqGetTradingInstrumentList) Reset()                    { *m = ReqGetTradingInstrumentList{} }
func (m *ReqGetTradingInstrumentList) String() string            { return proto.CompactTextString(m) }
func (*ReqGetTradingInstrumentList) ProtoMessage()               {}
func (*ReqGetTradingInstrumentList) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ReqGetTradingInstrumentList) GetExchange() ExchangeType {
	if m != nil {
		return m.Exchange
	}
	return ExchangeType_SHFE
}

type RspGetTradingInstrumentList struct {
}

func (m *RspGetTradingInstrumentList) Reset()                    { *m = RspGetTradingInstrumentList{} }
func (m *RspGetTradingInstrumentList) String() string            { return proto.CompactTextString(m) }
func (*RspGetTradingInstrumentList) ProtoMessage()               {}
func (*RspGetTradingInstrumentList) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

type ReqGetTradingInstrument struct {
	Symbol *Symbol `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
}

func (m *ReqGetTradingInstrument) Reset()                    { *m = ReqGetTradingInstrument{} }
func (m *ReqGetTradingInstrument) String() string            { return proto.CompactTextString(m) }
func (*ReqGetTradingInstrument) ProtoMessage()               {}
func (*ReqGetTradingInstrument) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *ReqGetTradingInstrument) GetSymbol() *Symbol {
	if m != nil {
		return m.Symbol
	}
	return nil
}

type RspGetTradingInstrument struct {
}

func (m *RspGetTradingInstrument) Reset()                    { *m = RspGetTradingInstrument{} }
func (m *RspGetTradingInstrument) String() string            { return proto.CompactTextString(m) }
func (*RspGetTradingInstrument) ProtoMessage()               {}
func (*RspGetTradingInstrument) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

type ReqSetTradingInstrument struct {
	List []*TradingInstrument `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *ReqSetTradingInstrument) Reset()                    { *m = ReqSetTradingInstrument{} }
func (m *ReqSetTradingInstrument) String() string            { return proto.CompactTextString(m) }
func (*ReqSetTradingInstrument) ProtoMessage()               {}
func (*ReqSetTradingInstrument) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *ReqSetTradingInstrument) GetList() []*TradingInstrument {
	if m != nil {
		return m.List
	}
	return nil
}

type RspSetTradingInstrument struct {
}

func (m *RspSetTradingInstrument) Reset()                    { *m = RspSetTradingInstrument{} }
func (m *RspSetTradingInstrument) String() string            { return proto.CompactTextString(m) }
func (*RspSetTradingInstrument) ProtoMessage()               {}
func (*RspSetTradingInstrument) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

type ReqSubscribe struct {
	List       []*Symbol `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
	BinaryData bool      `protobuf:"varint,2,opt,name=binary_data,json=binaryData" json:"binary_data,omitempty"`
}

func (m *ReqSubscribe) Reset()                    { *m = ReqSubscribe{} }
func (m *ReqSubscribe) String() string            { return proto.CompactTextString(m) }
func (*ReqSubscribe) ProtoMessage()               {}
func (*ReqSubscribe) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *ReqSubscribe) GetList() []*Symbol {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReqSubscribe) GetBinaryData() bool {
	if m != nil {
		return m.BinaryData
	}
	return false
}

type RspSubscribe struct {
}

func (m *RspSubscribe) Reset()                    { *m = RspSubscribe{} }
func (m *RspSubscribe) String() string            { return proto.CompactTextString(m) }
func (*RspSubscribe) ProtoMessage()               {}
func (*RspSubscribe) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

type ReqUnSubscribe struct {
	List []*Symbol `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *ReqUnSubscribe) Reset()                    { *m = ReqUnSubscribe{} }
func (m *ReqUnSubscribe) String() string            { return proto.CompactTextString(m) }
func (*ReqUnSubscribe) ProtoMessage()               {}
func (*ReqUnSubscribe) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *ReqUnSubscribe) GetList() []*Symbol {
	if m != nil {
		return m.List
	}
	return nil
}

type RspUnSubscribe struct {
}

func (m *RspUnSubscribe) Reset()                    { *m = RspUnSubscribe{} }
func (m *RspUnSubscribe) String() string            { return proto.CompactTextString(m) }
func (*RspUnSubscribe) ProtoMessage()               {}
func (*RspUnSubscribe) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

type ReqSaveKline struct {
	Series *KlineSeries `protobuf:"bytes,1,opt,name=series" json:"series,omitempty"`
}

func (m *ReqSaveKline) Reset()                    { *m = ReqSaveKline{} }
func (m *ReqSaveKline) String() string            { return proto.CompactTextString(m) }
func (*ReqSaveKline) ProtoMessage()               {}
func (*ReqSaveKline) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *ReqSaveKline) GetSeries() *KlineSeries {
	if m != nil {
		return m.Series
	}
	return nil
}

type RspSaveKline struct {
}

func (m *RspSaveKline) Reset()                    { *m = RspSaveKline{} }
func (m *RspSaveKline) String() string            { return proto.CompactTextString(m) }
func (*RspSaveKline) ProtoMessage()               {}
func (*RspSaveKline) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

type ReqCombineSubscribe struct {
	Symbol     *Symbol      `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	PeriodList []PeriodType `protobuf:"varint,2,rep,packed,name=period_list,json=periodList,enum=pb.PeriodType" json:"period_list,omitempty"`
}

func (m *ReqCombineSubscribe) Reset()                    { *m = ReqCombineSubscribe{} }
func (m *ReqCombineSubscribe) String() string            { return proto.CompactTextString(m) }
func (*ReqCombineSubscribe) ProtoMessage()               {}
func (*ReqCombineSubscribe) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

func (m *ReqCombineSubscribe) GetSymbol() *Symbol {
	if m != nil {
		return m.Symbol
	}
	return nil
}

func (m *ReqCombineSubscribe) GetPeriodList() []PeriodType {
	if m != nil {
		return m.PeriodList
	}
	return nil
}

type RspCombineSubscribe struct {
	// 合约
	Symbol *Symbol `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	// 历史
	History []*KlineSeries `protobuf:"bytes,2,rep,name=history" json:"history,omitempty"`
	// 最新行情
	Tick *MarketDataSnapshot `protobuf:"bytes,3,opt,name=tick" json:"tick,omitempty"`
	// 最新K线,按请求时的周期顺序依次返回
	Klines []*Kline `protobuf:"bytes,4,rep,name=klines" json:"klines,omitempty"`
}

func (m *RspCombineSubscribe) Reset()                    { *m = RspCombineSubscribe{} }
func (m *RspCombineSubscribe) String() string            { return proto.CompactTextString(m) }
func (*RspCombineSubscribe) ProtoMessage()               {}
func (*RspCombineSubscribe) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{13} }

func (m *RspCombineSubscribe) GetSymbol() *Symbol {
	if m != nil {
		return m.Symbol
	}
	return nil
}

func (m *RspCombineSubscribe) GetHistory() []*KlineSeries {
	if m != nil {
		return m.History
	}
	return nil
}

func (m *RspCombineSubscribe) GetTick() *MarketDataSnapshot {
	if m != nil {
		return m.Tick
	}
	return nil
}

func (m *RspCombineSubscribe) GetKlines() []*Kline {
	if m != nil {
		return m.Klines
	}
	return nil
}

type SymbolCacheSummary struct {
	Symbol          *Symbol `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	KlineLen        int32   `protobuf:"varint,2,opt,name=kline_len,json=klineLen" json:"kline_len,omitempty"`
	TickLen         int32   `protobuf:"varint,3,opt,name=tick_len,json=tickLen" json:"tick_len,omitempty"`
	SubscriberCount int32   `protobuf:"varint,4,opt,name=subscriber_count,json=subscriberCount" json:"subscriber_count,omitempty"`
}

func (m *SymbolCacheSummary) Reset()                    { *m = SymbolCacheSummary{} }
func (m *SymbolCacheSummary) String() string            { return proto.CompactTextString(m) }
func (*SymbolCacheSummary) ProtoMessage()               {}
func (*SymbolCacheSummary) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{14} }

func (m *SymbolCacheSummary) GetSymbol() *Symbol {
	if m != nil {
		return m.Symbol
	}
	return nil
}

func (m *SymbolCacheSummary) GetKlineLen() int32 {
	if m != nil {
		return m.KlineLen
	}
	return 0
}

func (m *SymbolCacheSummary) GetTickLen() int32 {
	if m != nil {
		return m.TickLen
	}
	return 0
}

func (m *SymbolCacheSummary) GetSubscriberCount() int32 {
	if m != nil {
		return m.SubscriberCount
	}
	return 0
}

type CacheSummary struct {
	KsMapSize            int32                 `protobuf:"varint,1,opt,name=ks_map_size,json=ksMapSize" json:"ks_map_size,omitempty"`
	TotalKlineLen        int32                 `protobuf:"varint,2,opt,name=total_kline_len,json=totalKlineLen" json:"total_kline_len,omitempty"`
	TotalTickLen         int32                 `protobuf:"varint,3,opt,name=total_tick_len,json=totalTickLen" json:"total_tick_len,omitempty"`
	TotalSubscriberCount int32                 `protobuf:"varint,4,opt,name=total_subscriber_count,json=totalSubscriberCount" json:"total_subscriber_count,omitempty"`
	SymbolList           []*SymbolCacheSummary `protobuf:"bytes,5,rep,name=symbol_list,json=symbolList" json:"symbol_list,omitempty"`
}

func (m *CacheSummary) Reset()                    { *m = CacheSummary{} }
func (m *CacheSummary) String() string            { return proto.CompactTextString(m) }
func (*CacheSummary) ProtoMessage()               {}
func (*CacheSummary) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{15} }

func (m *CacheSummary) GetKsMapSize() int32 {
	if m != nil {
		return m.KsMapSize
	}
	return 0
}

func (m *CacheSummary) GetTotalKlineLen() int32 {
	if m != nil {
		return m.TotalKlineLen
	}
	return 0
}

func (m *CacheSummary) GetTotalTickLen() int32 {
	if m != nil {
		return m.TotalTickLen
	}
	return 0
}

func (m *CacheSummary) GetTotalSubscriberCount() int32 {
	if m != nil {
		return m.TotalSubscriberCount
	}
	return 0
}

func (m *CacheSummary) GetSymbolList() []*SymbolCacheSummary {
	if m != nil {
		return m.SymbolList
	}
	return nil
}

type ReqGetDCenterInfo struct {
}

func (m *ReqGetDCenterInfo) Reset()                    { *m = ReqGetDCenterInfo{} }
func (m *ReqGetDCenterInfo) String() string            { return proto.CompactTextString(m) }
func (*ReqGetDCenterInfo) ProtoMessage()               {}
func (*ReqGetDCenterInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{16} }

type RspGetDCenterInfo struct {
	CacheSummary *CacheSummary `protobuf:"bytes,1,opt,name=cache_summary,json=cacheSummary" json:"cache_summary,omitempty"`
}

func (m *RspGetDCenterInfo) Reset()                    { *m = RspGetDCenterInfo{} }
func (m *RspGetDCenterInfo) String() string            { return proto.CompactTextString(m) }
func (*RspGetDCenterInfo) ProtoMessage()               {}
func (*RspGetDCenterInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{17} }

func (m *RspGetDCenterInfo) GetCacheSummary() *CacheSummary {
	if m != nil {
		return m.CacheSummary
	}
	return nil
}

func init() {
	proto.RegisterType((*ReqGetTradingInstrumentList)(nil), "pb.ReqGetTradingInstrumentList")
	proto.RegisterType((*RspGetTradingInstrumentList)(nil), "pb.RspGetTradingInstrumentList")
	proto.RegisterType((*ReqGetTradingInstrument)(nil), "pb.ReqGetTradingInstrument")
	proto.RegisterType((*RspGetTradingInstrument)(nil), "pb.RspGetTradingInstrument")
	proto.RegisterType((*ReqSetTradingInstrument)(nil), "pb.ReqSetTradingInstrument")
	proto.RegisterType((*RspSetTradingInstrument)(nil), "pb.RspSetTradingInstrument")
	proto.RegisterType((*ReqSubscribe)(nil), "pb.ReqSubscribe")
	proto.RegisterType((*RspSubscribe)(nil), "pb.RspSubscribe")
	proto.RegisterType((*ReqUnSubscribe)(nil), "pb.ReqUnSubscribe")
	proto.RegisterType((*RspUnSubscribe)(nil), "pb.RspUnSubscribe")
	proto.RegisterType((*ReqSaveKline)(nil), "pb.ReqSaveKline")
	proto.RegisterType((*RspSaveKline)(nil), "pb.RspSaveKline")
	proto.RegisterType((*ReqCombineSubscribe)(nil), "pb.ReqCombineSubscribe")
	proto.RegisterType((*RspCombineSubscribe)(nil), "pb.RspCombineSubscribe")
	proto.RegisterType((*SymbolCacheSummary)(nil), "pb.SymbolCacheSummary")
	proto.RegisterType((*CacheSummary)(nil), "pb.CacheSummary")
	proto.RegisterType((*ReqGetDCenterInfo)(nil), "pb.ReqGetDCenterInfo")
	proto.RegisterType((*RspGetDCenterInfo)(nil), "pb.RspGetDCenterInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DCenter service

type DCenterClient interface {
	// 保存K线数据
	// rpc SaveKline(ReqSaveKline) returns (RspSaveKline);
	// 订阅行情
	Subscribe(ctx context.Context, in *ReqSubscribe, opts ...grpc.CallOption) (DCenter_SubscribeClient, error)
	// 推送tick行情
	// rpc UpdateTick(MarketDataSnapshot) returns (EmptyResponse);
	// 复杂订阅.同时返回K线
	CombineSubscribe(ctx context.Context, in *ReqCombineSubscribe, opts ...grpc.CallOption) (DCenter_CombineSubscribeClient, error)
	// 数据库情况
	GetDCenterInfo(ctx context.Context, in *ReqGetDCenterInfo, opts ...grpc.CallOption) (*RspGetDCenterInfo, error)
	// 合约
	// rpc SetTradingInstrument(ReqSetTradingInstrument)
	//     returns (RspSetTradingInstrument);
	// 读取合约
	GetTradingInstrument(ctx context.Context, in *ReqGetTradingInstrument, opts ...grpc.CallOption) (*RspGetTradingInstrument, error)
	// 全部合约
	GetTradingInstrumentList(ctx context.Context, in *ReqGetTradingInstrumentList, opts ...grpc.CallOption) (*RspGetTradingInstrumentList, error)
}

type dCenterClient struct {
	cc *grpc.ClientConn
}

func NewDCenterClient(cc *grpc.ClientConn) DCenterClient {
	return &dCenterClient{cc}
}

func (c *dCenterClient) Subscribe(ctx context.Context, in *ReqSubscribe, opts ...grpc.CallOption) (DCenter_SubscribeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DCenter_serviceDesc.Streams[0], c.cc, "/pb.DCenter/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &dCenterSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DCenter_SubscribeClient interface {
	Recv() (*MarketDataSnapshot, error)
	grpc.ClientStream
}

type dCenterSubscribeClient struct {
	grpc.ClientStream
}

func (x *dCenterSubscribeClient) Recv() (*MarketDataSnapshot, error) {
	m := new(MarketDataSnapshot)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dCenterClient) CombineSubscribe(ctx context.Context, in *ReqCombineSubscribe, opts ...grpc.CallOption) (DCenter_CombineSubscribeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DCenter_serviceDesc.Streams[1], c.cc, "/pb.DCenter/CombineSubscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &dCenterCombineSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DCenter_CombineSubscribeClient interface {
	Recv() (*RspCombineSubscribe, error)
	grpc.ClientStream
}

type dCenterCombineSubscribeClient struct {
	grpc.ClientStream
}

func (x *dCenterCombineSubscribeClient) Recv() (*RspCombineSubscribe, error) {
	m := new(RspCombineSubscribe)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dCenterClient) GetDCenterInfo(ctx context.Context, in *ReqGetDCenterInfo, opts ...grpc.CallOption) (*RspGetDCenterInfo, error) {
	out := new(RspGetDCenterInfo)
	err := grpc.Invoke(ctx, "/pb.DCenter/GetDCenterInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCenterClient) GetTradingInstrument(ctx context.Context, in *ReqGetTradingInstrument, opts ...grpc.CallOption) (*RspGetTradingInstrument, error) {
	out := new(RspGetTradingInstrument)
	err := grpc.Invoke(ctx, "/pb.DCenter/GetTradingInstrument", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCenterClient) GetTradingInstrumentList(ctx context.Context, in *ReqGetTradingInstrumentList, opts ...grpc.CallOption) (*RspGetTradingInstrumentList, error) {
	out := new(RspGetTradingInstrumentList)
	err := grpc.Invoke(ctx, "/pb.DCenter/GetTradingInstrumentList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DCenter service

type DCenterServer interface {
	// 保存K线数据
	// rpc SaveKline(ReqSaveKline) returns (RspSaveKline);
	// 订阅行情
	Subscribe(*ReqSubscribe, DCenter_SubscribeServer) error
	// 推送tick行情
	// rpc UpdateTick(MarketDataSnapshot) returns (EmptyResponse);
	// 复杂订阅.同时返回K线
	CombineSubscribe(*ReqCombineSubscribe, DCenter_CombineSubscribeServer) error
	// 数据库情况
	GetDCenterInfo(context.Context, *ReqGetDCenterInfo) (*RspGetDCenterInfo, error)
	// 合约
	// rpc SetTradingInstrument(ReqSetTradingInstrument)
	//     returns (RspSetTradingInstrument);
	// 读取合约
	GetTradingInstrument(context.Context, *ReqGetTradingInstrument) (*RspGetTradingInstrument, error)
	// 全部合约
	GetTradingInstrumentList(context.Context, *ReqGetTradingInstrumentList) (*RspGetTradingInstrumentList, error)
}

func RegisterDCenterServer(s *grpc.Server, srv DCenterServer) {
	s.RegisterService(&_DCenter_serviceDesc, srv)
}

func _DCenter_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReqSubscribe)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DCenterServer).Subscribe(m, &dCenterSubscribeServer{stream})
}

type DCenter_SubscribeServer interface {
	Send(*MarketDataSnapshot) error
	grpc.ServerStream
}

type dCenterSubscribeServer struct {
	grpc.ServerStream
}

func (x *dCenterSubscribeServer) Send(m *MarketDataSnapshot) error {
	return x.ServerStream.SendMsg(m)
}

func _DCenter_CombineSubscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReqCombineSubscribe)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DCenterServer).CombineSubscribe(m, &dCenterCombineSubscribeServer{stream})
}

type DCenter_CombineSubscribeServer interface {
	Send(*RspCombineSubscribe) error
	grpc.ServerStream
}

type dCenterCombineSubscribeServer struct {
	grpc.ServerStream
}

func (x *dCenterCombineSubscribeServer) Send(m *RspCombineSubscribe) error {
	return x.ServerStream.SendMsg(m)
}

func _DCenter_GetDCenterInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetDCenterInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DCenterServer).GetDCenterInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DCenter/GetDCenterInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DCenterServer).GetDCenterInfo(ctx, req.(*ReqGetDCenterInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _DCenter_GetTradingInstrument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetTradingInstrument)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DCenterServer).GetTradingInstrument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DCenter/GetTradingInstrument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DCenterServer).GetTradingInstrument(ctx, req.(*ReqGetTradingInstrument))
	}
	return interceptor(ctx, in, info, handler)
}

func _DCenter_GetTradingInstrumentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetTradingInstrumentList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DCenterServer).GetTradingInstrumentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DCenter/GetTradingInstrumentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DCenterServer).GetTradingInstrumentList(ctx, req.(*ReqGetTradingInstrumentList))
	}
	return interceptor(ctx, in, info, handler)
}

var _DCenter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DCenter",
	HandlerType: (*DCenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDCenterInfo",
			Handler:    _DCenter_GetDCenterInfo_Handler,
		},
		{
			MethodName: "GetTradingInstrument",
			Handler:    _DCenter_GetTradingInstrument_Handler,
		},
		{
			MethodName: "GetTradingInstrumentList",
			Handler:    _DCenter_GetTradingInstrumentList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _DCenter_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CombineSubscribe",
			Handler:       _DCenter_CombineSubscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dcenter.proto",
}

func init() { proto.RegisterFile("dcenter.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 719 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xdd, 0x4e, 0xdb, 0x48,
	0x14, 0xc7, 0xe5, 0x24, 0x84, 0xe4, 0x24, 0x98, 0x30, 0x7c, 0x85, 0x44, 0x0b, 0xd9, 0xd1, 0x6a,
	0x37, 0xac, 0x56, 0x2c, 0xca, 0xee, 0x8a, 0xab, 0xed, 0x4d, 0x68, 0x2b, 0x0a, 0xa8, 0x68, 0x42,
	0xef, 0x2a, 0x59, 0x63, 0x67, 0x4a, 0xdc, 0xc4, 0x1f, 0x78, 0x26, 0x55, 0xc3, 0xab, 0xf4, 0x31,
	0xfa, 0x58, 0x7d, 0x83, 0x5e, 0x55, 0x73, 0xc6, 0x21, 0xdf, 0x08, 0xf5, 0xd2, 0xff, 0xf3, 0x9f,
	0xdf, 0xf9, 0x7b, 0x7c, 0x66, 0x0c, 0x1b, 0x5d, 0x4f, 0x84, 0x4a, 0x24, 0x27, 0x71, 0x12, 0xa9,
	0x88, 0x64, 0x62, 0xb7, 0x56, 0xf6, 0xa2, 0x20, 0x88, 0x42, 0xa3, 0xd4, 0xb6, 0x02, 0x9e, 0xf4,
	0x85, 0x72, 0xba, 0x5c, 0x71, 0x23, 0xd1, 0x4b, 0xa8, 0x33, 0x71, 0xff, 0x5a, 0xa8, 0xdb, 0x84,
	0x77, 0xfd, 0xf0, 0xee, 0x22, 0x94, 0x2a, 0x19, 0x06, 0x22, 0x54, 0x57, 0xbe, 0x54, 0xe4, 0x2f,
	0x28, 0x88, 0xcf, 0x5e, 0x8f, 0x87, 0x77, 0xa2, 0x6a, 0x35, 0xac, 0xa6, 0xdd, 0xaa, 0x9c, 0xc4,
	0xee, 0xc9, 0xcb, 0x54, 0xbb, 0x1d, 0xc5, 0x82, 0x3d, 0x3a, 0xe8, 0x2f, 0x50, 0x67, 0x32, 0x5e,
	0x05, 0xa3, 0xff, 0xc3, 0xfe, 0x8a, 0x5e, 0x84, 0x42, 0x5e, 0x8e, 0x02, 0x37, 0x1a, 0x60, 0x97,
	0x52, 0x0b, 0x74, 0x97, 0x0e, 0x2a, 0x2c, 0xad, 0xd0, 0x03, 0xd8, 0x5f, 0x41, 0xa7, 0xe7, 0x48,
	0xee, 0x2c, 0x23, 0x1f, 0x43, 0x6e, 0xe0, 0x4b, 0x55, 0xb5, 0x1a, 0xd9, 0x66, 0xa9, 0xb5, 0xab,
	0xb9, 0x0b, 0x26, 0x86, 0x96, 0xb4, 0xc1, 0x32, 0x0a, 0x7d, 0x0b, 0x65, 0xdd, 0x60, 0xe8, 0x4a,
	0x2f, 0xf1, 0x5d, 0x41, 0x0e, 0x67, 0xa8, 0xd3, 0x69, 0x51, 0x27, 0x47, 0x50, 0x72, 0xfd, 0x90,
	0x27, 0x23, 0xdc, 0xeb, 0x6a, 0xa6, 0x61, 0x35, 0x0b, 0x0c, 0x8c, 0x74, 0xce, 0x15, 0xa7, 0x36,
	0x94, 0x75, 0xaf, 0x31, 0x90, 0x9e, 0x82, 0xcd, 0xc4, 0xfd, 0xbb, 0xf0, 0xd9, 0x2d, 0x68, 0x05,
	0x6c, 0x26, 0xe3, 0xa9, 0x15, 0xf4, 0xcc, 0x84, 0xe4, 0x9f, 0xc4, 0xe5, 0xc0, 0x0f, 0x05, 0xf9,
	0x03, 0xf2, 0x52, 0x24, 0xbe, 0x90, 0xe9, 0xa6, 0x6e, 0x6a, 0x06, 0x96, 0x3a, 0x28, 0xb3, 0xb4,
	0x3c, 0x0e, 0x33, 0x5e, 0x48, 0x3f, 0xc2, 0x36, 0x13, 0xf7, 0xed, 0x28, 0x70, 0xb5, 0xf7, 0x31,
	0xd1, 0x33, 0x3e, 0x12, 0xf9, 0x1b, 0x4a, 0xb1, 0x48, 0xfc, 0xa8, 0xeb, 0x60, 0xf8, 0x4c, 0x23,
	0xdb, 0xb4, 0x5b, 0xb6, 0x36, 0xde, 0xa0, 0x8c, 0x13, 0x03, 0xc6, 0x82, 0x43, 0xf1, 0xd5, 0x82,
	0x6d, 0x26, 0xe3, 0x9f, 0x6a, 0x76, 0x0c, 0xeb, 0x3d, 0x5f, 0xaa, 0x28, 0x19, 0x61, 0xa3, 0x25,
	0x6f, 0x38, 0xae, 0x93, 0x3f, 0x21, 0xa7, 0x7c, 0xaf, 0x5f, 0xcd, 0x22, 0x6c, 0x4f, 0xfb, 0xae,
	0xf1, 0x30, 0xe8, 0xaf, 0xd1, 0x09, 0x79, 0x2c, 0x7b, 0x91, 0x62, 0xe8, 0x21, 0xbf, 0x42, 0xbe,
	0xaf, 0x19, 0xb2, 0x9a, 0x43, 0x6a, 0xf1, 0x91, 0xca, 0xd2, 0x02, 0xfd, 0x62, 0x01, 0x31, 0x61,
	0xda, 0xdc, 0xeb, 0x89, 0xce, 0x30, 0x08, 0x78, 0x32, 0x7a, 0x56, 0xe8, 0x3a, 0x14, 0x11, 0xe2,
	0x0c, 0x44, 0x88, 0x83, 0xb1, 0xc6, 0x0a, 0x28, 0x5c, 0x89, 0x90, 0x1c, 0x40, 0x41, 0x47, 0xc0,
	0x5a, 0x16, 0x6b, 0xeb, 0xfa, 0x59, 0x97, 0x8e, 0xa1, 0x22, 0xc7, 0xbb, 0x93, 0x38, 0x5e, 0x34,
	0x0c, 0x55, 0x35, 0x87, 0x96, 0xcd, 0x89, 0xde, 0xd6, 0x32, 0xfd, 0x66, 0x41, 0x79, 0x26, 0xd7,
	0x21, 0x94, 0xfa, 0xd2, 0x09, 0x78, 0xec, 0x48, 0xff, 0xc1, 0x9c, 0xe4, 0x35, 0x56, 0xec, 0xcb,
	0x6b, 0x1e, 0x77, 0xfc, 0x07, 0x41, 0x7e, 0x87, 0x4d, 0x15, 0x29, 0x3e, 0x70, 0xe6, 0x93, 0x6d,
	0xa0, 0x7c, 0x39, 0x8e, 0xf7, 0x1b, 0xd8, 0xc6, 0x37, 0x17, 0xb2, 0x8c, 0xea, 0x6d, 0x9a, 0xf4,
	0x5f, 0xd8, 0x33, 0xae, 0x15, 0x79, 0x77, 0xb0, 0xda, 0x99, 0x0d, 0x4d, 0xce, 0xa0, 0x64, 0x76,
	0xc8, 0x4c, 0xce, 0x1a, 0x6e, 0xfd, 0xde, 0x64, 0x03, 0xa7, 0x5f, 0x88, 0x81, 0xb1, 0xe2, 0x04,
	0x6d, 0xc3, 0x96, 0xb9, 0x56, 0xce, 0xdb, 0x78, 0xfd, 0x5d, 0x84, 0x1f, 0x22, 0xfa, 0x06, 0xb6,
	0xcc, 0x65, 0x31, 0x25, 0x92, 0xff, 0x60, 0xc3, 0xd3, 0x14, 0x47, 0x1a, 0x4c, 0xfa, 0x95, 0xf0,
	0x4a, 0x9b, 0xc1, 0x97, 0xbd, 0xa9, 0xa7, 0xd6, 0xf7, 0x0c, 0xac, 0xa7, 0x18, 0x72, 0x06, 0xc5,
	0xc9, 0x8c, 0xe2, 0xc2, 0xe9, 0x7b, 0xa1, 0xb6, 0x62, 0xb0, 0x4e, 0x2d, 0xf2, 0x0a, 0x2a, 0x0b,
	0x33, 0xbe, 0x9f, 0xae, 0x9f, 0x2f, 0xd4, 0x4c, 0x61, 0xf1, 0x54, 0x9c, 0x5a, 0xe4, 0x05, 0xd8,
	0x73, 0x6f, 0xb5, 0x9b, 0x52, 0x66, 0xe5, 0xda, 0x6e, 0xca, 0x98, 0x73, 0xdf, 0xc0, 0xce, 0xd2,
	0x1b, 0xb8, 0x3e, 0xa1, 0x2c, 0x14, 0x6b, 0xf5, 0x09, 0x6b, 0x71, 0xe5, 0x7b, 0xa8, 0xae, 0xfc,
	0x7f, 0x1c, 0x3d, 0x41, 0xd5, 0x86, 0xda, 0xd1, 0x13, 0x64, 0x6d, 0x70, 0xf3, 0xf8, 0x9f, 0xfa,
	0xe7, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0xe7, 0x37, 0xcb, 0xdd, 0x06, 0x00, 0x00,
}
