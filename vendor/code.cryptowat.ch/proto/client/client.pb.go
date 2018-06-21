// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client/client.proto

/*
Package ProtobufClient is a generated protocol buffer package.

It is generated from these files:
	client/client.proto

It has these top-level messages:
	ClientMessage
	ClientIdentificationMessage
	WebAuthenticationMessage
	APIAuthenticationMessage
	ClientSessionMessage
	ClientSubscribeMessage
	ClientUnsubscribeMessage
*/
package ProtobufClient

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

type APIAuthenticationMessage_Source int32

const (
	APIAuthenticationMessage_UNKNOWN        APIAuthenticationMessage_Source = 0
	APIAuthenticationMessage_GOLANG_SDK     APIAuthenticationMessage_Source = 1
	APIAuthenticationMessage_JAVASCRIPT_SDK APIAuthenticationMessage_Source = 2
)

var APIAuthenticationMessage_Source_name = map[int32]string{
	0: "UNKNOWN",
	1: "GOLANG_SDK",
	2: "JAVASCRIPT_SDK",
}
var APIAuthenticationMessage_Source_value = map[string]int32{
	"UNKNOWN":        0,
	"GOLANG_SDK":     1,
	"JAVASCRIPT_SDK": 2,
}

func (x APIAuthenticationMessage_Source) String() string {
	return proto.EnumName(APIAuthenticationMessage_Source_name, int32(x))
}
func (APIAuthenticationMessage_Source) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

// ClientMessage is a wrapper message used to describe the supplied client message
// and pass metadata about the client.
type ClientMessage struct {
	// Types that are valid to be assigned to Body:
	//	*ClientMessage_Identification
	//	*ClientMessage_Subscribe
	//	*ClientMessage_Unsubscribe
	//	*ClientMessage_WebAuthentication
	//	*ClientMessage_ApiAuthentication
	Body isClientMessage_Body `protobuf_oneof:"body"`
}

func (m *ClientMessage) Reset()                    { *m = ClientMessage{} }
func (m *ClientMessage) String() string            { return proto.CompactTextString(m) }
func (*ClientMessage) ProtoMessage()               {}
func (*ClientMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isClientMessage_Body interface {
	isClientMessage_Body()
}

type ClientMessage_Identification struct {
	Identification *ClientIdentificationMessage `protobuf:"bytes,1,opt,name=identification,oneof"`
}
type ClientMessage_Subscribe struct {
	Subscribe *ClientSubscribeMessage `protobuf:"bytes,2,opt,name=subscribe,oneof"`
}
type ClientMessage_Unsubscribe struct {
	Unsubscribe *ClientUnsubscribeMessage `protobuf:"bytes,3,opt,name=unsubscribe,oneof"`
}
type ClientMessage_WebAuthentication struct {
	WebAuthentication *WebAuthenticationMessage `protobuf:"bytes,4,opt,name=webAuthentication,oneof"`
}
type ClientMessage_ApiAuthentication struct {
	ApiAuthentication *APIAuthenticationMessage `protobuf:"bytes,5,opt,name=apiAuthentication,oneof"`
}

func (*ClientMessage_Identification) isClientMessage_Body()    {}
func (*ClientMessage_Subscribe) isClientMessage_Body()         {}
func (*ClientMessage_Unsubscribe) isClientMessage_Body()       {}
func (*ClientMessage_WebAuthentication) isClientMessage_Body() {}
func (*ClientMessage_ApiAuthentication) isClientMessage_Body() {}

func (m *ClientMessage) GetBody() isClientMessage_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *ClientMessage) GetIdentification() *ClientIdentificationMessage {
	if x, ok := m.GetBody().(*ClientMessage_Identification); ok {
		return x.Identification
	}
	return nil
}

func (m *ClientMessage) GetSubscribe() *ClientSubscribeMessage {
	if x, ok := m.GetBody().(*ClientMessage_Subscribe); ok {
		return x.Subscribe
	}
	return nil
}

func (m *ClientMessage) GetUnsubscribe() *ClientUnsubscribeMessage {
	if x, ok := m.GetBody().(*ClientMessage_Unsubscribe); ok {
		return x.Unsubscribe
	}
	return nil
}

func (m *ClientMessage) GetWebAuthentication() *WebAuthenticationMessage {
	if x, ok := m.GetBody().(*ClientMessage_WebAuthentication); ok {
		return x.WebAuthentication
	}
	return nil
}

func (m *ClientMessage) GetApiAuthentication() *APIAuthenticationMessage {
	if x, ok := m.GetBody().(*ClientMessage_ApiAuthentication); ok {
		return x.ApiAuthentication
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ClientMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ClientMessage_OneofMarshaler, _ClientMessage_OneofUnmarshaler, _ClientMessage_OneofSizer, []interface{}{
		(*ClientMessage_Identification)(nil),
		(*ClientMessage_Subscribe)(nil),
		(*ClientMessage_Unsubscribe)(nil),
		(*ClientMessage_WebAuthentication)(nil),
		(*ClientMessage_ApiAuthentication)(nil),
	}
}

func _ClientMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ClientMessage)
	// body
	switch x := m.Body.(type) {
	case *ClientMessage_Identification:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Identification); err != nil {
			return err
		}
	case *ClientMessage_Subscribe:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Subscribe); err != nil {
			return err
		}
	case *ClientMessage_Unsubscribe:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Unsubscribe); err != nil {
			return err
		}
	case *ClientMessage_WebAuthentication:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.WebAuthentication); err != nil {
			return err
		}
	case *ClientMessage_ApiAuthentication:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ApiAuthentication); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ClientMessage.Body has unexpected type %T", x)
	}
	return nil
}

func _ClientMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ClientMessage)
	switch tag {
	case 1: // body.identification
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClientIdentificationMessage)
		err := b.DecodeMessage(msg)
		m.Body = &ClientMessage_Identification{msg}
		return true, err
	case 2: // body.subscribe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClientSubscribeMessage)
		err := b.DecodeMessage(msg)
		m.Body = &ClientMessage_Subscribe{msg}
		return true, err
	case 3: // body.unsubscribe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClientUnsubscribeMessage)
		err := b.DecodeMessage(msg)
		m.Body = &ClientMessage_Unsubscribe{msg}
		return true, err
	case 4: // body.webAuthentication
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WebAuthenticationMessage)
		err := b.DecodeMessage(msg)
		m.Body = &ClientMessage_WebAuthentication{msg}
		return true, err
	case 5: // body.apiAuthentication
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(APIAuthenticationMessage)
		err := b.DecodeMessage(msg)
		m.Body = &ClientMessage_ApiAuthentication{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ClientMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ClientMessage)
	// body
	switch x := m.Body.(type) {
	case *ClientMessage_Identification:
		s := proto.Size(x.Identification)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ClientMessage_Subscribe:
		s := proto.Size(x.Subscribe)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ClientMessage_Unsubscribe:
		s := proto.Size(x.Unsubscribe)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ClientMessage_WebAuthentication:
		s := proto.Size(x.WebAuthentication)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ClientMessage_ApiAuthentication:
		s := proto.Size(x.ApiAuthentication)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// ClientIdentificationMessage is the first message sent is from the client to
// the server, identifying itself and the subscriptions it desires.
type ClientIdentificationMessage struct {
	Useragent     string   `protobuf:"bytes,1,opt,name=useragent" json:"useragent,omitempty"`
	Revision      string   `protobuf:"bytes,2,opt,name=revision" json:"revision,omitempty"`
	Integration   string   `protobuf:"bytes,3,opt,name=integration" json:"integration,omitempty"`
	Locale        string   `protobuf:"bytes,4,opt,name=locale" json:"locale,omitempty"`
	Subscriptions []string `protobuf:"bytes,5,rep,name=subscriptions" json:"subscriptions,omitempty"`
}

func (m *ClientIdentificationMessage) Reset()                    { *m = ClientIdentificationMessage{} }
func (m *ClientIdentificationMessage) String() string            { return proto.CompactTextString(m) }
func (*ClientIdentificationMessage) ProtoMessage()               {}
func (*ClientIdentificationMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClientIdentificationMessage) GetUseragent() string {
	if m != nil {
		return m.Useragent
	}
	return ""
}

func (m *ClientIdentificationMessage) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *ClientIdentificationMessage) GetIntegration() string {
	if m != nil {
		return m.Integration
	}
	return ""
}

func (m *ClientIdentificationMessage) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *ClientIdentificationMessage) GetSubscriptions() []string {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

type WebAuthenticationMessage struct {
	Identification *ClientIdentificationMessage `protobuf:"bytes,1,opt,name=identification" json:"identification,omitempty"`
	Token          string                       `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	Nonce          string                       `protobuf:"bytes,3,opt,name=nonce" json:"nonce,omitempty"`
	AccessList     []string                     `protobuf:"bytes,4,rep,name=access_list,json=accessList" json:"access_list,omitempty"`
}

func (m *WebAuthenticationMessage) Reset()                    { *m = WebAuthenticationMessage{} }
func (m *WebAuthenticationMessage) String() string            { return proto.CompactTextString(m) }
func (*WebAuthenticationMessage) ProtoMessage()               {}
func (*WebAuthenticationMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WebAuthenticationMessage) GetIdentification() *ClientIdentificationMessage {
	if m != nil {
		return m.Identification
	}
	return nil
}

func (m *WebAuthenticationMessage) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *WebAuthenticationMessage) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *WebAuthenticationMessage) GetAccessList() []string {
	if m != nil {
		return m.AccessList
	}
	return nil
}

type APIAuthenticationMessage struct {
	Token         string                          `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Nonce         string                          `protobuf:"bytes,2,opt,name=nonce" json:"nonce,omitempty"`
	ApiKey        string                          `protobuf:"bytes,3,opt,name=api_key,json=apiKey" json:"api_key,omitempty"`
	Source        APIAuthenticationMessage_Source `protobuf:"varint,4,opt,name=source,enum=ProtobufClient.APIAuthenticationMessage_Source" json:"source,omitempty"`
	Version       string                          `protobuf:"bytes,5,opt,name=version" json:"version,omitempty"`
	Subscriptions []string                        `protobuf:"bytes,6,rep,name=subscriptions" json:"subscriptions,omitempty"`
}

func (m *APIAuthenticationMessage) Reset()                    { *m = APIAuthenticationMessage{} }
func (m *APIAuthenticationMessage) String() string            { return proto.CompactTextString(m) }
func (*APIAuthenticationMessage) ProtoMessage()               {}
func (*APIAuthenticationMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *APIAuthenticationMessage) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *APIAuthenticationMessage) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *APIAuthenticationMessage) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *APIAuthenticationMessage) GetSource() APIAuthenticationMessage_Source {
	if m != nil {
		return m.Source
	}
	return APIAuthenticationMessage_UNKNOWN
}

func (m *APIAuthenticationMessage) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *APIAuthenticationMessage) GetSubscriptions() []string {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

type ClientSessionMessage struct {
	// Types that are valid to be assigned to SessionConfig:
	//	*ClientSessionMessage_Session_
	//	*ClientSessionMessage_AnonymousTradingSession_
	SessionConfig isClientSessionMessage_SessionConfig `protobuf_oneof:"SessionConfig"`
}

func (m *ClientSessionMessage) Reset()                    { *m = ClientSessionMessage{} }
func (m *ClientSessionMessage) String() string            { return proto.CompactTextString(m) }
func (*ClientSessionMessage) ProtoMessage()               {}
func (*ClientSessionMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isClientSessionMessage_SessionConfig interface {
	isClientSessionMessage_SessionConfig()
}

type ClientSessionMessage_Session_ struct {
	Session *ClientSessionMessage_Session `protobuf:"bytes,1,opt,name=session,oneof"`
}
type ClientSessionMessage_AnonymousTradingSession_ struct {
	AnonymousTradingSession *ClientSessionMessage_AnonymousTradingSession `protobuf:"bytes,2,opt,name=anonymousTradingSession,oneof"`
}

func (*ClientSessionMessage_Session_) isClientSessionMessage_SessionConfig()                 {}
func (*ClientSessionMessage_AnonymousTradingSession_) isClientSessionMessage_SessionConfig() {}

func (m *ClientSessionMessage) GetSessionConfig() isClientSessionMessage_SessionConfig {
	if m != nil {
		return m.SessionConfig
	}
	return nil
}

func (m *ClientSessionMessage) GetSession() *ClientSessionMessage_Session {
	if x, ok := m.GetSessionConfig().(*ClientSessionMessage_Session_); ok {
		return x.Session
	}
	return nil
}

func (m *ClientSessionMessage) GetAnonymousTradingSession() *ClientSessionMessage_AnonymousTradingSession {
	if x, ok := m.GetSessionConfig().(*ClientSessionMessage_AnonymousTradingSession_); ok {
		return x.AnonymousTradingSession
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ClientSessionMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ClientSessionMessage_OneofMarshaler, _ClientSessionMessage_OneofUnmarshaler, _ClientSessionMessage_OneofSizer, []interface{}{
		(*ClientSessionMessage_Session_)(nil),
		(*ClientSessionMessage_AnonymousTradingSession_)(nil),
	}
}

func _ClientSessionMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ClientSessionMessage)
	// SessionConfig
	switch x := m.SessionConfig.(type) {
	case *ClientSessionMessage_Session_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Session); err != nil {
			return err
		}
	case *ClientSessionMessage_AnonymousTradingSession_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AnonymousTradingSession); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ClientSessionMessage.SessionConfig has unexpected type %T", x)
	}
	return nil
}

func _ClientSessionMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ClientSessionMessage)
	switch tag {
	case 1: // SessionConfig.session
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClientSessionMessage_Session)
		err := b.DecodeMessage(msg)
		m.SessionConfig = &ClientSessionMessage_Session_{msg}
		return true, err
	case 2: // SessionConfig.anonymousTradingSession
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClientSessionMessage_AnonymousTradingSession)
		err := b.DecodeMessage(msg)
		m.SessionConfig = &ClientSessionMessage_AnonymousTradingSession_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ClientSessionMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ClientSessionMessage)
	// SessionConfig
	switch x := m.SessionConfig.(type) {
	case *ClientSessionMessage_Session_:
		s := proto.Size(x.Session)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ClientSessionMessage_AnonymousTradingSession_:
		s := proto.Size(x.AnonymousTradingSession)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ClientSessionMessage_Session struct {
	UserId   string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	Expires  int64  `protobuf:"varint,2,opt,name=expires" json:"expires,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	MfaToken string `protobuf:"bytes,4,opt,name=mfaToken" json:"mfaToken,omitempty"`
}

func (m *ClientSessionMessage_Session) Reset()                    { *m = ClientSessionMessage_Session{} }
func (m *ClientSessionMessage_Session) String() string            { return proto.CompactTextString(m) }
func (*ClientSessionMessage_Session) ProtoMessage()               {}
func (*ClientSessionMessage_Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *ClientSessionMessage_Session) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ClientSessionMessage_Session) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func (m *ClientSessionMessage_Session) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ClientSessionMessage_Session) GetMfaToken() string {
	if m != nil {
		return m.MfaToken
	}
	return ""
}

type ClientSessionMessage_AnonymousTradingSession struct {
	Exchange   string `protobuf:"bytes,1,opt,name=exchange" json:"exchange,omitempty"`
	Token      string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	Expiration int64  `protobuf:"varint,3,opt,name=expiration" json:"expiration,omitempty"`
}

func (m *ClientSessionMessage_AnonymousTradingSession) Reset() {
	*m = ClientSessionMessage_AnonymousTradingSession{}
}
func (m *ClientSessionMessage_AnonymousTradingSession) String() string {
	return proto.CompactTextString(m)
}
func (*ClientSessionMessage_AnonymousTradingSession) ProtoMessage() {}
func (*ClientSessionMessage_AnonymousTradingSession) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 1}
}

func (m *ClientSessionMessage_AnonymousTradingSession) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *ClientSessionMessage_AnonymousTradingSession) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ClientSessionMessage_AnonymousTradingSession) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

// SubscribeMessage informs the Stream Hub to subscribe the current web socket to
// the supplied channel.
type ClientSubscribeMessage struct {
	SubscriptionKeys []string `protobuf:"bytes,1,rep,name=subscriptionKeys" json:"subscriptionKeys,omitempty"`
}

func (m *ClientSubscribeMessage) Reset()                    { *m = ClientSubscribeMessage{} }
func (m *ClientSubscribeMessage) String() string            { return proto.CompactTextString(m) }
func (*ClientSubscribeMessage) ProtoMessage()               {}
func (*ClientSubscribeMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ClientSubscribeMessage) GetSubscriptionKeys() []string {
	if m != nil {
		return m.SubscriptionKeys
	}
	return nil
}

// UnsubscribeMessage informs the Stream Hub to unsubscribe the current web socket
// from the supplied channel.
type ClientUnsubscribeMessage struct {
	SubscriptionKeys []string `protobuf:"bytes,1,rep,name=subscriptionKeys" json:"subscriptionKeys,omitempty"`
}

func (m *ClientUnsubscribeMessage) Reset()                    { *m = ClientUnsubscribeMessage{} }
func (m *ClientUnsubscribeMessage) String() string            { return proto.CompactTextString(m) }
func (*ClientUnsubscribeMessage) ProtoMessage()               {}
func (*ClientUnsubscribeMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ClientUnsubscribeMessage) GetSubscriptionKeys() []string {
	if m != nil {
		return m.SubscriptionKeys
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientMessage)(nil), "ProtobufClient.ClientMessage")
	proto.RegisterType((*ClientIdentificationMessage)(nil), "ProtobufClient.ClientIdentificationMessage")
	proto.RegisterType((*WebAuthenticationMessage)(nil), "ProtobufClient.WebAuthenticationMessage")
	proto.RegisterType((*APIAuthenticationMessage)(nil), "ProtobufClient.APIAuthenticationMessage")
	proto.RegisterType((*ClientSessionMessage)(nil), "ProtobufClient.ClientSessionMessage")
	proto.RegisterType((*ClientSessionMessage_Session)(nil), "ProtobufClient.ClientSessionMessage.Session")
	proto.RegisterType((*ClientSessionMessage_AnonymousTradingSession)(nil), "ProtobufClient.ClientSessionMessage.AnonymousTradingSession")
	proto.RegisterType((*ClientSubscribeMessage)(nil), "ProtobufClient.ClientSubscribeMessage")
	proto.RegisterType((*ClientUnsubscribeMessage)(nil), "ProtobufClient.ClientUnsubscribeMessage")
	proto.RegisterEnum("ProtobufClient.APIAuthenticationMessage_Source", APIAuthenticationMessage_Source_name, APIAuthenticationMessage_Source_value)
}

func init() { proto.RegisterFile("client/client.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xdf, 0x4e, 0xdb, 0x3c,
	0x14, 0xc0, 0x9b, 0x06, 0x52, 0x72, 0x2a, 0xfa, 0xf1, 0x79, 0x08, 0xa2, 0x6e, 0xda, 0x50, 0x34,
	0x4d, 0x68, 0x9b, 0x8a, 0xc4, 0xae, 0x26, 0xed, 0x26, 0x80, 0xa0, 0x5d, 0x59, 0x41, 0x29, 0x8c,
	0xdd, 0x21, 0x27, 0x75, 0x83, 0x45, 0x71, 0xa2, 0x38, 0x61, 0xf4, 0x79, 0xf6, 0x10, 0xdb, 0xc5,
	0xee, 0xf7, 0x5a, 0x53, 0x1c, 0xa7, 0x49, 0xdb, 0x64, 0x62, 0xd2, 0xae, 0x9c, 0x73, 0x7c, 0xce,
	0xcf, 0xc7, 0xe7, 0x8f, 0x03, 0x4f, 0xdc, 0x09, 0x25, 0x2c, 0xda, 0x4b, 0x97, 0x4e, 0x10, 0xfa,
	0x91, 0x8f, 0x5a, 0xe7, 0xc9, 0xe2, 0xc4, 0xe3, 0x43, 0xa1, 0x35, 0x7f, 0xa8, 0xb0, 0x9e, 0x7e,
	0x7e, 0x22, 0x9c, 0x63, 0x8f, 0xa0, 0x4b, 0x68, 0xd1, 0x11, 0x61, 0x11, 0x1d, 0x53, 0x17, 0x47,
	0xd4, 0x67, 0x86, 0xb2, 0xa3, 0xec, 0x36, 0xf7, 0xdf, 0x74, 0xe6, 0x5d, 0x3b, 0xe9, 0xd2, 0x9b,
	0xb3, 0x95, 0x90, 0x6e, 0xcd, 0x5e, 0x80, 0xa0, 0x63, 0xd0, 0x79, 0xec, 0x70, 0x37, 0xa4, 0x0e,
	0x31, 0xea, 0x82, 0xf8, 0xaa, 0x9c, 0x38, 0xcc, 0xcc, 0x72, 0x58, 0xee, 0x8a, 0x4e, 0xa1, 0x19,
	0xb3, 0x9c, 0xa4, 0x0a, 0xd2, 0x6e, 0x39, 0xe9, 0x32, 0x37, 0xcc, 0x59, 0x45, 0x77, 0xf4, 0x05,
	0xfe, 0xff, 0x4a, 0x1c, 0x2b, 0x8e, 0x6e, 0x92, 0x60, 0xe5, 0x7d, 0x57, 0xca, 0x99, 0x57, 0x8b,
	0x86, 0x39, 0x73, 0x19, 0x92, 0x90, 0x71, 0x40, 0x17, 0xc8, 0xab, 0xe5, 0x64, 0xeb, 0xbc, 0x57,
	0x49, 0x5e, 0x82, 0x1c, 0x68, 0xb0, 0xe2, 0xf8, 0xa3, 0xa9, 0xf9, 0x5d, 0x81, 0xa7, 0x7f, 0xa8,
	0x01, 0x7a, 0x06, 0x7a, 0xcc, 0x49, 0x88, 0x3d, 0xc2, 0x22, 0x51, 0x43, 0xdd, 0xce, 0x15, 0xa8,
	0x0d, 0x6b, 0x21, 0xb9, 0xa7, 0x3c, 0x09, 0xab, 0x2e, 0x36, 0x67, 0x32, 0xda, 0x81, 0x26, 0x65,
	0x11, 0xf1, 0xc2, 0x34, 0x6a, 0x55, 0x6c, 0x17, 0x55, 0x68, 0x0b, 0xb4, 0x89, 0xef, 0xe2, 0x09,
	0x11, 0xc9, 0xd2, 0x6d, 0x29, 0xa1, 0x97, 0xb0, 0x2e, 0x93, 0x1b, 0x24, 0x76, 0xdc, 0x58, 0xdd,
	0x51, 0x77, 0x75, 0x7b, 0x5e, 0x69, 0xfe, 0x54, 0xc0, 0xa8, 0xca, 0x26, 0x1a, 0xfe, 0x83, 0xfe,
	0x5b, 0xea, 0xbe, 0x4d, 0x58, 0x8d, 0xfc, 0x5b, 0x92, 0x5d, 0x35, 0x15, 0x12, 0x2d, 0xf3, 0x99,
	0x4b, 0xe4, 0x0d, 0x53, 0x01, 0xbd, 0x80, 0x26, 0x76, 0x5d, 0xc2, 0xf9, 0xf5, 0x84, 0xf2, 0xc8,
	0x58, 0x11, 0x37, 0x80, 0x54, 0x75, 0x4a, 0x79, 0x64, 0x7e, 0xab, 0x83, 0x51, 0x55, 0xb2, 0xfc,
	0x24, 0xa5, 0xf4, 0xa4, 0x7a, 0xf1, 0xa4, 0x6d, 0x68, 0xe0, 0x80, 0x5e, 0xdf, 0x92, 0xa9, 0x8c,
	0x40, 0xc3, 0x01, 0xed, 0x93, 0x29, 0x3a, 0x01, 0x8d, 0xfb, 0x71, 0xe8, 0xa6, 0xe9, 0x6d, 0xed,
	0xef, 0x3d, 0xb6, 0x63, 0x3a, 0x43, 0xe1, 0x66, 0x4b, 0x77, 0x64, 0x40, 0xe3, 0x9e, 0x84, 0x3c,
	0xeb, 0x3d, 0xdd, 0xce, 0xc4, 0xe5, 0x4a, 0x69, 0x65, 0x95, 0x7a, 0x0f, 0x5a, 0x4a, 0x44, 0x4d,
	0x68, 0x5c, 0x0e, 0xfa, 0x83, 0xb3, 0xab, 0xc1, 0x46, 0x0d, 0xb5, 0x00, 0x4e, 0xce, 0x4e, 0xad,
	0xc1, 0xc9, 0xf5, 0xf0, 0xa8, 0xbf, 0xa1, 0x20, 0x04, 0xad, 0x8f, 0xd6, 0x67, 0x6b, 0x78, 0x68,
	0xf7, 0xce, 0x2f, 0x84, 0xae, 0x6e, 0xfe, 0x52, 0x61, 0x53, 0x0e, 0x34, 0xe1, 0xbc, 0x90, 0xa1,
	0x2e, 0x34, 0x78, 0xaa, 0x91, 0x95, 0x7d, 0x5b, 0xf1, 0x0e, 0xcc, 0xb9, 0x75, 0xa4, 0xd8, 0xad,
	0xd9, 0x99, 0x3b, 0x7a, 0x80, 0x6d, 0xcc, 0x7c, 0x36, 0xbd, 0xf3, 0x63, 0x7e, 0x11, 0xe2, 0x11,
	0x65, 0x9e, 0xb4, 0x92, 0x2f, 0xcc, 0x87, 0x47, 0x91, 0xad, 0x72, 0x46, 0xb7, 0x66, 0x57, 0xe1,
	0xdb, 0x77, 0xd0, 0x90, 0x9f, 0xc9, 0x28, 0x24, 0x53, 0xd5, 0x1b, 0xc9, 0x8a, 0x4b, 0x29, 0x49,
	0x3d, 0x79, 0x08, 0x68, 0x48, 0xb8, 0x08, 0x46, 0xb5, 0x33, 0x31, 0x6f, 0x11, 0xb5, 0xd8, 0x22,
	0x6d, 0x58, 0xbb, 0x1b, 0xe3, 0x0b, 0xb1, 0x91, 0x0e, 0xd5, 0x4c, 0x6e, 0xdf, 0xc2, 0x76, 0x45,
	0x90, 0x89, 0x1b, 0x79, 0x70, 0x6f, 0x30, 0xf3, 0x88, 0x0c, 0x60, 0x26, 0x57, 0x74, 0xfd, 0x73,
	0x00, 0x11, 0x49, 0x3e, 0xdc, 0xaa, 0x5d, 0xd0, 0x1c, 0xfc, 0x07, 0xeb, 0x12, 0x7e, 0xe8, 0xb3,
	0x31, 0xf5, 0xcc, 0x23, 0xd8, 0x2a, 0x7f, 0x99, 0xd1, 0x6b, 0xd8, 0x28, 0xf6, 0x4b, 0x9f, 0x4c,
	0xb9, 0xa1, 0x88, 0x3e, 0x5a, 0xd2, 0x9b, 0xc7, 0x60, 0x54, 0xbd, 0xca, 0x7f, 0xc3, 0x71, 0x34,
	0xf1, 0x23, 0x7b, 0xf7, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x8d, 0x0e, 0x8f, 0xdf, 0x06, 0x00,
	0x00,
}
