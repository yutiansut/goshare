// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MessageType int32

const (
	// 订阅行情
	MessageType_REQ_SUBSCRIBE_MARKET_DATA MessageType = 0
	// 订阅结果
	MessageType_RSP_SUBSCRIBE_MARKET_DATA MessageType = 1
	// 有新行情
	MessageType_RTN_MARKET_DATA_UPDATE MessageType = 2
	// 订阅行情
	MessageType_REQ_UNSUBSCRIBE_MARKET_DATA MessageType = 3
	// 订阅结果
	MessageType_RSP_UNSUBSCRIBE_MARKET_DATA MessageType = 4
	// 心跳
	MessageType_HEATBEAT MessageType = 5
	// 上传tick
	MessageType_UPLOAD_TICK MessageType = 6
)

var MessageType_name = map[int32]string{
	0: "REQ_SUBSCRIBE_MARKET_DATA",
	1: "RSP_SUBSCRIBE_MARKET_DATA",
	2: "RTN_MARKET_DATA_UPDATE",
	3: "REQ_UNSUBSCRIBE_MARKET_DATA",
	4: "RSP_UNSUBSCRIBE_MARKET_DATA",
	5: "HEATBEAT",
	6: "UPLOAD_TICK",
}
var MessageType_value = map[string]int32{
	"REQ_SUBSCRIBE_MARKET_DATA":   0,
	"RSP_SUBSCRIBE_MARKET_DATA":   1,
	"RTN_MARKET_DATA_UPDATE":      2,
	"REQ_UNSUBSCRIBE_MARKET_DATA": 3,
	"RSP_UNSUBSCRIBE_MARKET_DATA": 4,
	"HEATBEAT":                    5,
	"UPLOAD_TICK":                 6,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

// Message 消息
type Message struct {
	Type   MessageType `protobuf:"varint,1,opt,name=type,enum=pb.MessageType" json:"type,omitempty"`
	Target string      `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	Data   []byte      `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *Message) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_REQ_SUBSCRIBE_MARKET_DATA
}

func (m *Message) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// MessageList 消息list
type MessageList struct {
	List []*Message `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *MessageList) Reset()                    { *m = MessageList{} }
func (m *MessageList) String() string            { return proto.CompactTextString(m) }
func (*MessageList) ProtoMessage()               {}
func (*MessageList) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *MessageList) GetList() []*Message {
	if m != nil {
		return m.List
	}
	return nil
}

type StreamMessage struct {
	Type   MessageType `protobuf:"varint,1,opt,name=type,enum=pb.MessageType" json:"type,omitempty"`
	Target string      `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	Data   string      `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *StreamMessage) Reset()                    { *m = StreamMessage{} }
func (m *StreamMessage) String() string            { return proto.CompactTextString(m) }
func (*StreamMessage) ProtoMessage()               {}
func (*StreamMessage) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *StreamMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_REQ_SUBSCRIBE_MARKET_DATA
}

func (m *StreamMessage) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *StreamMessage) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "pb.Message")
	proto.RegisterType((*MessageList)(nil), "pb.MessageList")
	proto.RegisterType((*StreamMessage)(nil), "pb.StreamMessage")
	proto.RegisterEnum("pb.MessageType", MessageType_name, MessageType_value)
}

func init() { proto.RegisterFile("message.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x4f, 0x4f, 0x83, 0x30,
	0x18, 0xc6, 0x2d, 0x20, 0xba, 0x97, 0xcd, 0x91, 0xf7, 0xb0, 0xa0, 0xc6, 0x8c, 0xe0, 0x85, 0x78,
	0xe0, 0x30, 0x3f, 0x41, 0x19, 0x4d, 0x5c, 0xf6, 0x47, 0x2c, 0xe5, 0xe2, 0x05, 0x21, 0x36, 0xcb,
	0x12, 0x97, 0x11, 0xe8, 0x65, 0xdf, 0xd0, 0x8f, 0x65, 0xd6, 0x11, 0xb3, 0x83, 0xbb, 0x79, 0x6b,
	0x9f, 0xdf, 0xd3, 0x5f, 0x9e, 0xa4, 0x30, 0xd8, 0xca, 0xb6, 0x2d, 0xd7, 0x32, 0xaa, 0x9b, 0x9d,
	0xda, 0xa1, 0x51, 0x57, 0xc1, 0x3b, 0x5c, 0x2d, 0x8f, 0x21, 0x3e, 0x82, 0xa5, 0xf6, 0xb5, 0xf4,
	0x88, 0x4f, 0xc2, 0x9b, 0xc9, 0x30, 0xaa, 0xab, 0xa8, 0x43, 0x62, 0x5f, 0x4b, 0xae, 0x21, 0x8e,
	0xc0, 0x56, 0x65, 0xb3, 0x96, 0xca, 0x33, 0x7c, 0x12, 0xf6, 0x78, 0x77, 0x43, 0x04, 0xeb, 0xb3,
	0x54, 0xa5, 0x67, 0xfa, 0x24, 0xec, 0x73, 0x7d, 0x0e, 0x22, 0x70, 0x3a, 0xc1, 0x62, 0xd3, 0x2a,
	0x1c, 0x83, 0xf5, 0xb5, 0x69, 0x95, 0x47, 0x7c, 0x33, 0x74, 0x26, 0xce, 0x89, 0x9f, 0x6b, 0x10,
	0x7c, 0xc0, 0x20, 0x53, 0x8d, 0x2c, 0xb7, 0xff, 0xbe, 0xa8, 0x77, 0x5c, 0xf4, 0xf4, 0x4d, 0x7e,
	0x27, 0x1d, 0x0c, 0xf8, 0x00, 0xb7, 0x9c, 0xbd, 0x15, 0x59, 0x1e, 0x67, 0x53, 0x3e, 0x8b, 0x59,
	0xb1, 0xa4, 0x7c, 0xce, 0x44, 0x91, 0x50, 0x41, 0xdd, 0x0b, 0x8d, 0xb3, 0xf4, 0x0c, 0x26, 0x78,
	0x07, 0x23, 0x2e, 0x56, 0xa7, 0x61, 0x91, 0xa7, 0x09, 0x15, 0xcc, 0x35, 0x70, 0x0c, 0xf7, 0x07,
	0x73, 0xbe, 0xfa, 0xfb, 0xb1, 0xa9, 0x0b, 0x59, 0x7a, 0xb6, 0x60, 0x61, 0x1f, 0xae, 0x5f, 0x18,
	0x15, 0x31, 0xa3, 0xc2, 0xbd, 0xc4, 0x21, 0x38, 0x79, 0xba, 0x78, 0xa5, 0x49, 0x21, 0x66, 0xd3,
	0xb9, 0x6b, 0x57, 0xb6, 0xfe, 0xc3, 0xe7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x1f, 0xb5,
	0x31, 0xd4, 0x01, 0x00, 0x00,
}
