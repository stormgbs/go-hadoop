// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RpcHeader.proto

package hadoop_common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// *
// RpcKind determine the rpcEngine and the serialization of the rpc request
type RpcKindProto int32

const (
	RpcKindProto_RPC_BUILTIN         RpcKindProto = 0
	RpcKindProto_RPC_WRITABLE        RpcKindProto = 1
	RpcKindProto_RPC_PROTOCOL_BUFFER RpcKindProto = 2
)

var RpcKindProto_name = map[int32]string{
	0: "RPC_BUILTIN",
	1: "RPC_WRITABLE",
	2: "RPC_PROTOCOL_BUFFER",
}
var RpcKindProto_value = map[string]int32{
	"RPC_BUILTIN":         0,
	"RPC_WRITABLE":        1,
	"RPC_PROTOCOL_BUFFER": 2,
}

func (x RpcKindProto) Enum() *RpcKindProto {
	p := new(RpcKindProto)
	*p = x
	return p
}
func (x RpcKindProto) String() string {
	return proto.EnumName(RpcKindProto_name, int32(x))
}
func (x *RpcKindProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcKindProto_value, data, "RpcKindProto")
	if err != nil {
		return err
	}
	*x = RpcKindProto(value)
	return nil
}
func (RpcKindProto) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

type RpcRequestHeaderProto_OperationProto int32

const (
	RpcRequestHeaderProto_RPC_FINAL_PACKET        RpcRequestHeaderProto_OperationProto = 0
	RpcRequestHeaderProto_RPC_CONTINUATION_PACKET RpcRequestHeaderProto_OperationProto = 1
	RpcRequestHeaderProto_RPC_CLOSE_CONNECTION    RpcRequestHeaderProto_OperationProto = 2
)

var RpcRequestHeaderProto_OperationProto_name = map[int32]string{
	0: "RPC_FINAL_PACKET",
	1: "RPC_CONTINUATION_PACKET",
	2: "RPC_CLOSE_CONNECTION",
}
var RpcRequestHeaderProto_OperationProto_value = map[string]int32{
	"RPC_FINAL_PACKET":        0,
	"RPC_CONTINUATION_PACKET": 1,
	"RPC_CLOSE_CONNECTION":    2,
}

func (x RpcRequestHeaderProto_OperationProto) Enum() *RpcRequestHeaderProto_OperationProto {
	p := new(RpcRequestHeaderProto_OperationProto)
	*p = x
	return p
}
func (x RpcRequestHeaderProto_OperationProto) String() string {
	return proto.EnumName(RpcRequestHeaderProto_OperationProto_name, int32(x))
}
func (x *RpcRequestHeaderProto_OperationProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcRequestHeaderProto_OperationProto_value, data, "RpcRequestHeaderProto_OperationProto")
	if err != nil {
		return err
	}
	*x = RpcRequestHeaderProto_OperationProto(value)
	return nil
}
func (RpcRequestHeaderProto_OperationProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor9, []int{1, 0}
}

type RpcResponseHeaderProto_RpcStatusProto int32

const (
	RpcResponseHeaderProto_SUCCESS RpcResponseHeaderProto_RpcStatusProto = 0
	RpcResponseHeaderProto_ERROR   RpcResponseHeaderProto_RpcStatusProto = 1
	RpcResponseHeaderProto_FATAL   RpcResponseHeaderProto_RpcStatusProto = 2
)

var RpcResponseHeaderProto_RpcStatusProto_name = map[int32]string{
	0: "SUCCESS",
	1: "ERROR",
	2: "FATAL",
}
var RpcResponseHeaderProto_RpcStatusProto_value = map[string]int32{
	"SUCCESS": 0,
	"ERROR":   1,
	"FATAL":   2,
}

func (x RpcResponseHeaderProto_RpcStatusProto) Enum() *RpcResponseHeaderProto_RpcStatusProto {
	p := new(RpcResponseHeaderProto_RpcStatusProto)
	*p = x
	return p
}
func (x RpcResponseHeaderProto_RpcStatusProto) String() string {
	return proto.EnumName(RpcResponseHeaderProto_RpcStatusProto_name, int32(x))
}
func (x *RpcResponseHeaderProto_RpcStatusProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcResponseHeaderProto_RpcStatusProto_value, data, "RpcResponseHeaderProto_RpcStatusProto")
	if err != nil {
		return err
	}
	*x = RpcResponseHeaderProto_RpcStatusProto(value)
	return nil
}
func (RpcResponseHeaderProto_RpcStatusProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor9, []int{2, 0}
}

type RpcResponseHeaderProto_RpcErrorCodeProto int32

const (
	// Non-fatal Rpc error - connection left open for future rpc calls
	RpcResponseHeaderProto_ERROR_APPLICATION          RpcResponseHeaderProto_RpcErrorCodeProto = 1
	RpcResponseHeaderProto_ERROR_NO_SUCH_METHOD       RpcResponseHeaderProto_RpcErrorCodeProto = 2
	RpcResponseHeaderProto_ERROR_NO_SUCH_PROTOCOL     RpcResponseHeaderProto_RpcErrorCodeProto = 3
	RpcResponseHeaderProto_ERROR_RPC_SERVER           RpcResponseHeaderProto_RpcErrorCodeProto = 4
	RpcResponseHeaderProto_ERROR_SERIALIZING_RESPONSE RpcResponseHeaderProto_RpcErrorCodeProto = 5
	RpcResponseHeaderProto_ERROR_RPC_VERSION_MISMATCH RpcResponseHeaderProto_RpcErrorCodeProto = 6
	// Fatal Server side Rpc error - connection closed
	RpcResponseHeaderProto_FATAL_UNKNOWN                   RpcResponseHeaderProto_RpcErrorCodeProto = 10
	RpcResponseHeaderProto_FATAL_UNSUPPORTED_SERIALIZATION RpcResponseHeaderProto_RpcErrorCodeProto = 11
	RpcResponseHeaderProto_FATAL_INVALID_RPC_HEADER        RpcResponseHeaderProto_RpcErrorCodeProto = 12
	RpcResponseHeaderProto_FATAL_DESERIALIZING_REQUEST     RpcResponseHeaderProto_RpcErrorCodeProto = 13
	RpcResponseHeaderProto_FATAL_VERSION_MISMATCH          RpcResponseHeaderProto_RpcErrorCodeProto = 14
	RpcResponseHeaderProto_FATAL_UNAUTHORIZED              RpcResponseHeaderProto_RpcErrorCodeProto = 15
)

var RpcResponseHeaderProto_RpcErrorCodeProto_name = map[int32]string{
	1:  "ERROR_APPLICATION",
	2:  "ERROR_NO_SUCH_METHOD",
	3:  "ERROR_NO_SUCH_PROTOCOL",
	4:  "ERROR_RPC_SERVER",
	5:  "ERROR_SERIALIZING_RESPONSE",
	6:  "ERROR_RPC_VERSION_MISMATCH",
	10: "FATAL_UNKNOWN",
	11: "FATAL_UNSUPPORTED_SERIALIZATION",
	12: "FATAL_INVALID_RPC_HEADER",
	13: "FATAL_DESERIALIZING_REQUEST",
	14: "FATAL_VERSION_MISMATCH",
	15: "FATAL_UNAUTHORIZED",
}
var RpcResponseHeaderProto_RpcErrorCodeProto_value = map[string]int32{
	"ERROR_APPLICATION":               1,
	"ERROR_NO_SUCH_METHOD":            2,
	"ERROR_NO_SUCH_PROTOCOL":          3,
	"ERROR_RPC_SERVER":                4,
	"ERROR_SERIALIZING_RESPONSE":      5,
	"ERROR_RPC_VERSION_MISMATCH":      6,
	"FATAL_UNKNOWN":                   10,
	"FATAL_UNSUPPORTED_SERIALIZATION": 11,
	"FATAL_INVALID_RPC_HEADER":        12,
	"FATAL_DESERIALIZING_REQUEST":     13,
	"FATAL_VERSION_MISMATCH":          14,
	"FATAL_UNAUTHORIZED":              15,
}

func (x RpcResponseHeaderProto_RpcErrorCodeProto) Enum() *RpcResponseHeaderProto_RpcErrorCodeProto {
	p := new(RpcResponseHeaderProto_RpcErrorCodeProto)
	*p = x
	return p
}
func (x RpcResponseHeaderProto_RpcErrorCodeProto) String() string {
	return proto.EnumName(RpcResponseHeaderProto_RpcErrorCodeProto_name, int32(x))
}
func (x *RpcResponseHeaderProto_RpcErrorCodeProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcResponseHeaderProto_RpcErrorCodeProto_value, data, "RpcResponseHeaderProto_RpcErrorCodeProto")
	if err != nil {
		return err
	}
	*x = RpcResponseHeaderProto_RpcErrorCodeProto(value)
	return nil
}
func (RpcResponseHeaderProto_RpcErrorCodeProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor9, []int{2, 1}
}

type RpcSaslProto_SaslState int32

const (
	RpcSaslProto_SUCCESS   RpcSaslProto_SaslState = 0
	RpcSaslProto_NEGOTIATE RpcSaslProto_SaslState = 1
	RpcSaslProto_INITIATE  RpcSaslProto_SaslState = 2
	RpcSaslProto_CHALLENGE RpcSaslProto_SaslState = 3
	RpcSaslProto_RESPONSE  RpcSaslProto_SaslState = 4
	RpcSaslProto_WRAP      RpcSaslProto_SaslState = 5
)

var RpcSaslProto_SaslState_name = map[int32]string{
	0: "SUCCESS",
	1: "NEGOTIATE",
	2: "INITIATE",
	3: "CHALLENGE",
	4: "RESPONSE",
	5: "WRAP",
}
var RpcSaslProto_SaslState_value = map[string]int32{
	"SUCCESS":   0,
	"NEGOTIATE": 1,
	"INITIATE":  2,
	"CHALLENGE": 3,
	"RESPONSE":  4,
	"WRAP":      5,
}

func (x RpcSaslProto_SaslState) Enum() *RpcSaslProto_SaslState {
	p := new(RpcSaslProto_SaslState)
	*p = x
	return p
}
func (x RpcSaslProto_SaslState) String() string {
	return proto.EnumName(RpcSaslProto_SaslState_name, int32(x))
}
func (x *RpcSaslProto_SaslState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcSaslProto_SaslState_value, data, "RpcSaslProto_SaslState")
	if err != nil {
		return err
	}
	*x = RpcSaslProto_SaslState(value)
	return nil
}
func (RpcSaslProto_SaslState) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{3, 0} }

// *
// Used to pass through the information necessary to continue
// a trace after an RPC is made. All we need is the traceid
// (so we know the overarching trace this message is a part of), and
// the id of the current span when this message was sent, so we know
// what span caused the new span we will create when this message is received.
type RPCTraceInfoProto struct {
	TraceId          *int64 `protobuf:"varint,1,opt,name=traceId" json:"traceId,omitempty"`
	ParentId         *int64 `protobuf:"varint,2,opt,name=parentId" json:"parentId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RPCTraceInfoProto) Reset()                    { *m = RPCTraceInfoProto{} }
func (m *RPCTraceInfoProto) String() string            { return proto.CompactTextString(m) }
func (*RPCTraceInfoProto) ProtoMessage()               {}
func (*RPCTraceInfoProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *RPCTraceInfoProto) GetTraceId() int64 {
	if m != nil && m.TraceId != nil {
		return *m.TraceId
	}
	return 0
}

func (m *RPCTraceInfoProto) GetParentId() int64 {
	if m != nil && m.ParentId != nil {
		return *m.ParentId
	}
	return 0
}

type RpcRequestHeaderProto struct {
	RpcKind  *RpcKindProto                         `protobuf:"varint,1,opt,name=rpcKind,enum=hadoop.common.RpcKindProto" json:"rpcKind,omitempty"`
	RpcOp    *RpcRequestHeaderProto_OperationProto `protobuf:"varint,2,opt,name=rpcOp,enum=hadoop.common.RpcRequestHeaderProto_OperationProto" json:"rpcOp,omitempty"`
	CallId   *int32                                `protobuf:"zigzag32,3,req,name=callId" json:"callId,omitempty"`
	ClientId []byte                                `protobuf:"bytes,4,req,name=clientId" json:"clientId,omitempty"`
	// clientId + callId uniquely identifies a request
	// retry count, 1 means this is the first retry
	RetryCount       *int32             `protobuf:"zigzag32,5,opt,name=retryCount,def=-1" json:"retryCount,omitempty"`
	TraceInfo        *RPCTraceInfoProto `protobuf:"bytes,6,opt,name=traceInfo" json:"traceInfo,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *RpcRequestHeaderProto) Reset()                    { *m = RpcRequestHeaderProto{} }
func (m *RpcRequestHeaderProto) String() string            { return proto.CompactTextString(m) }
func (*RpcRequestHeaderProto) ProtoMessage()               {}
func (*RpcRequestHeaderProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

const Default_RpcRequestHeaderProto_RetryCount int32 = -1

func (m *RpcRequestHeaderProto) GetRpcKind() RpcKindProto {
	if m != nil && m.RpcKind != nil {
		return *m.RpcKind
	}
	return RpcKindProto_RPC_BUILTIN
}

func (m *RpcRequestHeaderProto) GetRpcOp() RpcRequestHeaderProto_OperationProto {
	if m != nil && m.RpcOp != nil {
		return *m.RpcOp
	}
	return RpcRequestHeaderProto_RPC_FINAL_PACKET
}

func (m *RpcRequestHeaderProto) GetCallId() int32 {
	if m != nil && m.CallId != nil {
		return *m.CallId
	}
	return 0
}

func (m *RpcRequestHeaderProto) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *RpcRequestHeaderProto) GetRetryCount() int32 {
	if m != nil && m.RetryCount != nil {
		return *m.RetryCount
	}
	return Default_RpcRequestHeaderProto_RetryCount
}

func (m *RpcRequestHeaderProto) GetTraceInfo() *RPCTraceInfoProto {
	if m != nil {
		return m.TraceInfo
	}
	return nil
}

// *
// Rpc Response Header
// +------------------------------------------------------------------+
// | Rpc total response length in bytes (4 bytes int)                 |
// |  (sum of next two parts)                                         |
// +------------------------------------------------------------------+
// | RpcResponseHeaderProto - serialized delimited ie has len         |
// +------------------------------------------------------------------+
// | if request is successful:                                        |
// |   - RpcResponse -  The actual rpc response  bytes follow         |
// |     the response header                                          |
// |     This response is serialized based on RpcKindProto            |
// | if request fails :                                               |
// |   The rpc response header contains the necessary info            |
// +------------------------------------------------------------------+
//
// Note that rpc response header is also used when connection setup fails.
// Ie the response looks like a rpc response with a fake callId.
type RpcResponseHeaderProto struct {
	CallId              *uint32                                   `protobuf:"varint,1,req,name=callId" json:"callId,omitempty"`
	Status              *RpcResponseHeaderProto_RpcStatusProto    `protobuf:"varint,2,req,name=status,enum=hadoop.common.RpcResponseHeaderProto_RpcStatusProto" json:"status,omitempty"`
	ServerIpcVersionNum *uint32                                   `protobuf:"varint,3,opt,name=serverIpcVersionNum" json:"serverIpcVersionNum,omitempty"`
	ExceptionClassName  *string                                   `protobuf:"bytes,4,opt,name=exceptionClassName" json:"exceptionClassName,omitempty"`
	ErrorMsg            *string                                   `protobuf:"bytes,5,opt,name=errorMsg" json:"errorMsg,omitempty"`
	ErrorDetail         *RpcResponseHeaderProto_RpcErrorCodeProto `protobuf:"varint,6,opt,name=errorDetail,enum=hadoop.common.RpcResponseHeaderProto_RpcErrorCodeProto" json:"errorDetail,omitempty"`
	ClientId            []byte                                    `protobuf:"bytes,7,opt,name=clientId" json:"clientId,omitempty"`
	RetryCount          *int32                                    `protobuf:"zigzag32,8,opt,name=retryCount,def=-1" json:"retryCount,omitempty"`
	XXX_unrecognized    []byte                                    `json:"-"`
}

func (m *RpcResponseHeaderProto) Reset()                    { *m = RpcResponseHeaderProto{} }
func (m *RpcResponseHeaderProto) String() string            { return proto.CompactTextString(m) }
func (*RpcResponseHeaderProto) ProtoMessage()               {}
func (*RpcResponseHeaderProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

const Default_RpcResponseHeaderProto_RetryCount int32 = -1

func (m *RpcResponseHeaderProto) GetCallId() uint32 {
	if m != nil && m.CallId != nil {
		return *m.CallId
	}
	return 0
}

func (m *RpcResponseHeaderProto) GetStatus() RpcResponseHeaderProto_RpcStatusProto {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return RpcResponseHeaderProto_SUCCESS
}

func (m *RpcResponseHeaderProto) GetServerIpcVersionNum() uint32 {
	if m != nil && m.ServerIpcVersionNum != nil {
		return *m.ServerIpcVersionNum
	}
	return 0
}

func (m *RpcResponseHeaderProto) GetExceptionClassName() string {
	if m != nil && m.ExceptionClassName != nil {
		return *m.ExceptionClassName
	}
	return ""
}

func (m *RpcResponseHeaderProto) GetErrorMsg() string {
	if m != nil && m.ErrorMsg != nil {
		return *m.ErrorMsg
	}
	return ""
}

func (m *RpcResponseHeaderProto) GetErrorDetail() RpcResponseHeaderProto_RpcErrorCodeProto {
	if m != nil && m.ErrorDetail != nil {
		return *m.ErrorDetail
	}
	return RpcResponseHeaderProto_ERROR_APPLICATION
}

func (m *RpcResponseHeaderProto) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *RpcResponseHeaderProto) GetRetryCount() int32 {
	if m != nil && m.RetryCount != nil {
		return *m.RetryCount
	}
	return Default_RpcResponseHeaderProto_RetryCount
}

type RpcSaslProto struct {
	Version          *uint32                  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	State            *RpcSaslProto_SaslState  `protobuf:"varint,2,req,name=state,enum=hadoop.common.RpcSaslProto_SaslState" json:"state,omitempty"`
	Token            []byte                   `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	Auths            []*RpcSaslProto_SaslAuth `protobuf:"bytes,4,rep,name=auths" json:"auths,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *RpcSaslProto) Reset()                    { *m = RpcSaslProto{} }
func (m *RpcSaslProto) String() string            { return proto.CompactTextString(m) }
func (*RpcSaslProto) ProtoMessage()               {}
func (*RpcSaslProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *RpcSaslProto) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *RpcSaslProto) GetState() RpcSaslProto_SaslState {
	if m != nil && m.State != nil {
		return *m.State
	}
	return RpcSaslProto_SUCCESS
}

func (m *RpcSaslProto) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *RpcSaslProto) GetAuths() []*RpcSaslProto_SaslAuth {
	if m != nil {
		return m.Auths
	}
	return nil
}

type RpcSaslProto_SaslAuth struct {
	Method           *string `protobuf:"bytes,1,req,name=method" json:"method,omitempty"`
	Mechanism        *string `protobuf:"bytes,2,req,name=mechanism" json:"mechanism,omitempty"`
	Protocol         *string `protobuf:"bytes,3,opt,name=protocol" json:"protocol,omitempty"`
	ServerId         *string `protobuf:"bytes,4,opt,name=serverId" json:"serverId,omitempty"`
	Challenge        []byte  `protobuf:"bytes,5,opt,name=challenge" json:"challenge,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpcSaslProto_SaslAuth) Reset()                    { *m = RpcSaslProto_SaslAuth{} }
func (m *RpcSaslProto_SaslAuth) String() string            { return proto.CompactTextString(m) }
func (*RpcSaslProto_SaslAuth) ProtoMessage()               {}
func (*RpcSaslProto_SaslAuth) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3, 0} }

func (m *RpcSaslProto_SaslAuth) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *RpcSaslProto_SaslAuth) GetMechanism() string {
	if m != nil && m.Mechanism != nil {
		return *m.Mechanism
	}
	return ""
}

func (m *RpcSaslProto_SaslAuth) GetProtocol() string {
	if m != nil && m.Protocol != nil {
		return *m.Protocol
	}
	return ""
}

func (m *RpcSaslProto_SaslAuth) GetServerId() string {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return ""
}

func (m *RpcSaslProto_SaslAuth) GetChallenge() []byte {
	if m != nil {
		return m.Challenge
	}
	return nil
}

func init() {
	proto.RegisterType((*RPCTraceInfoProto)(nil), "hadoop.common.RPCTraceInfoProto")
	proto.RegisterType((*RpcRequestHeaderProto)(nil), "hadoop.common.RpcRequestHeaderProto")
	proto.RegisterType((*RpcResponseHeaderProto)(nil), "hadoop.common.RpcResponseHeaderProto")
	proto.RegisterType((*RpcSaslProto)(nil), "hadoop.common.RpcSaslProto")
	proto.RegisterType((*RpcSaslProto_SaslAuth)(nil), "hadoop.common.RpcSaslProto.SaslAuth")
	proto.RegisterEnum("hadoop.common.RpcKindProto", RpcKindProto_name, RpcKindProto_value)
	proto.RegisterEnum("hadoop.common.RpcRequestHeaderProto_OperationProto", RpcRequestHeaderProto_OperationProto_name, RpcRequestHeaderProto_OperationProto_value)
	proto.RegisterEnum("hadoop.common.RpcResponseHeaderProto_RpcStatusProto", RpcResponseHeaderProto_RpcStatusProto_name, RpcResponseHeaderProto_RpcStatusProto_value)
	proto.RegisterEnum("hadoop.common.RpcResponseHeaderProto_RpcErrorCodeProto", RpcResponseHeaderProto_RpcErrorCodeProto_name, RpcResponseHeaderProto_RpcErrorCodeProto_value)
	proto.RegisterEnum("hadoop.common.RpcSaslProto_SaslState", RpcSaslProto_SaslState_name, RpcSaslProto_SaslState_value)
}

func init() { proto.RegisterFile("RpcHeader.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 981 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xed, 0x6e, 0xe3, 0x44,
	0x17, 0x5e, 0x3b, 0x49, 0xdb, 0x9c, 0x7c, 0xac, 0x33, 0xdb, 0xed, 0x5a, 0xed, 0x6a, 0xd7, 0xca,
	0xfb, 0x22, 0x45, 0x48, 0x44, 0x90, 0x5d, 0x84, 0xb4, 0x48, 0x48, 0xae, 0x33, 0x6d, 0x86, 0xba,
	0xb6, 0x19, 0x3b, 0xad, 0x58, 0x81, 0x22, 0xe3, 0xcc, 0x36, 0x11, 0x89, 0x6d, 0x6c, 0x67, 0x05,
	0x77, 0xc1, 0x2f, 0x7e, 0x73, 0x05, 0x88, 0x5b, 0xe2, 0x4e, 0xd0, 0xcc, 0x24, 0x4e, 0xb3, 0xa9,
	0x80, 0x7f, 0x3e, 0xe7, 0x39, 0xf3, 0xcc, 0xf9, 0x78, 0x8e, 0x07, 0x1e, 0xd3, 0x34, 0x1a, 0xb1,
	0x70, 0xca, 0xb2, 0x7e, 0x9a, 0x25, 0x45, 0x82, 0x5a, 0xb3, 0x70, 0x9a, 0x24, 0x69, 0x3f, 0x4a,
	0x96, 0xcb, 0x24, 0xee, 0x12, 0xe8, 0x50, 0xcf, 0x0a, 0xb2, 0x30, 0x62, 0x24, 0x7e, 0x97, 0x78,
	0x22, 0x46, 0x87, 0xc3, 0x42, 0x78, 0xa6, 0xba, 0x62, 0x28, 0xbd, 0x0a, 0xdd, 0x98, 0xe8, 0x14,
	0x8e, 0xd2, 0x30, 0x63, 0x71, 0x41, 0xa6, 0xba, 0x2a, 0xa0, 0xd2, 0xee, 0xfe, 0x5a, 0x81, 0xa7,
	0x34, 0x8d, 0x28, 0xfb, 0x69, 0xc5, 0xf2, 0x42, 0x5e, 0x2a, 0xf9, 0x3e, 0x87, 0xc3, 0x2c, 0x8d,
	0xae, 0xe6, 0xb1, 0xe4, 0x6b, 0x0f, 0xce, 0xfa, 0x3b, 0x59, 0xf4, 0xa9, 0x44, 0x45, 0x34, 0xdd,
	0xc4, 0x22, 0x02, 0xb5, 0x2c, 0x8d, 0xdc, 0x54, 0xdc, 0xd4, 0x1e, 0xbc, 0xda, 0x3f, 0xb4, 0x7f,
	0x57, 0xdf, 0x4d, 0x59, 0x16, 0x16, 0xf3, 0x24, 0x96, 0x64, 0x92, 0x01, 0x9d, 0xc0, 0x41, 0x14,
	0x2e, 0x16, 0x64, 0xaa, 0x57, 0x0c, 0xb5, 0xd7, 0xa1, 0x6b, 0x8b, 0xd7, 0x13, 0x2d, 0xe6, 0xb2,
	0x9e, 0xaa, 0xa1, 0xf6, 0x9a, 0xb4, 0xb4, 0x51, 0x17, 0x20, 0x63, 0x45, 0xf6, 0x8b, 0x95, 0xac,
	0xe2, 0x42, 0xaf, 0x19, 0x4a, 0xaf, 0xf3, 0x46, 0xfd, 0xe4, 0x33, 0x7a, 0xcf, 0x8b, 0xbe, 0x82,
	0x7a, 0xb1, 0xe9, 0x9d, 0x7e, 0x60, 0x28, 0xbd, 0xc6, 0xc0, 0xf8, 0x30, 0xcd, 0x0f, 0xdb, 0x4b,
	0xb7, 0x47, 0xba, 0xdf, 0x43, 0x7b, 0x37, 0x61, 0x74, 0x0c, 0x1a, 0xf5, 0xac, 0xc9, 0x05, 0x71,
	0x4c, 0x7b, 0xe2, 0x99, 0xd6, 0x15, 0x0e, 0xb4, 0x47, 0xe8, 0x0c, 0x9e, 0x71, 0xaf, 0xe5, 0x3a,
	0x01, 0x71, 0xc6, 0x66, 0x40, 0x5c, 0x67, 0x03, 0x2a, 0x48, 0x87, 0x63, 0x01, 0xda, 0xae, 0x8f,
	0x79, 0x88, 0x83, 0x2d, 0x1e, 0xa0, 0xa9, 0xdd, 0x3f, 0x0f, 0xe0, 0x44, 0xb4, 0x29, 0x4f, 0x93,
	0x38, 0x67, 0xf7, 0x67, 0xb2, 0xed, 0x88, 0x62, 0xa8, 0xbd, 0x56, 0xd9, 0x11, 0x1b, 0x0e, 0xf2,
	0x22, 0x2c, 0x56, 0xb9, 0xae, 0x1a, 0x6a, 0xaf, 0x3d, 0x78, 0xfd, 0x50, 0xd7, 0xf7, 0xe8, 0xb8,
	0xdb, 0x17, 0xc7, 0x64, 0x89, 0x6b, 0x0e, 0xf4, 0x29, 0x3c, 0xc9, 0x59, 0xf6, 0x9e, 0x65, 0x24,
	0x8d, 0x6e, 0x58, 0x96, 0xcf, 0x93, 0xd8, 0x59, 0x2d, 0xf5, 0x8a, 0xa1, 0xf4, 0x5a, 0xf4, 0x21,
	0x08, 0xf5, 0x01, 0xb1, 0x9f, 0x23, 0x96, 0xf2, 0x8e, 0x58, 0x8b, 0x30, 0xcf, 0x9d, 0x70, 0xc9,
	0xf4, 0xaa, 0xa1, 0xf4, 0xea, 0xf4, 0x01, 0x84, 0x4f, 0x90, 0x65, 0x59, 0x92, 0x5d, 0xe7, 0x77,
	0x62, 0x46, 0x75, 0x5a, 0xda, 0xe8, 0x5b, 0x68, 0x88, 0xef, 0x21, 0x2b, 0xc2, 0xf9, 0x42, 0xcc,
	0xa7, 0x3d, 0xf8, 0xe2, 0x3f, 0x17, 0x84, 0xf9, 0x59, 0x2b, 0x99, 0x32, 0x59, 0xd3, 0x7d, 0xae,
	0x1d, 0xe1, 0x1c, 0x1a, 0xca, 0x3f, 0x08, 0xe7, 0xe8, 0x21, 0xe1, 0x74, 0x5f, 0x41, 0x7b, 0xb7,
	0x65, 0xa8, 0x01, 0x87, 0xfe, 0xd8, 0xb2, 0xb0, 0xef, 0x6b, 0x8f, 0x50, 0x1d, 0x6a, 0x98, 0x52,
	0x97, 0x6a, 0x0a, 0xff, 0xbc, 0x30, 0x03, 0xd3, 0xd6, 0xd4, 0xee, 0x5f, 0x2a, 0x74, 0xf6, 0xf2,
	0x42, 0x4f, 0xa1, 0x23, 0x62, 0x27, 0xa6, 0xe7, 0xd9, 0xc4, 0x12, 0xe2, 0x90, 0xaa, 0x90, 0x6e,
	0xc7, 0x9d, 0xf8, 0x63, 0x6b, 0x34, 0xb9, 0xc6, 0xc1, 0xc8, 0x1d, 0x6a, 0x2a, 0x3a, 0x85, 0x93,
	0x5d, 0xc4, 0xa3, 0x6e, 0xe0, 0x5a, 0xae, 0xad, 0x55, 0xb8, 0xfc, 0x24, 0xc6, 0x15, 0xe5, 0x63,
	0x7a, 0x83, 0xa9, 0x56, 0x45, 0x2f, 0xe0, 0x54, 0x7a, 0x7d, 0x4c, 0x89, 0x69, 0x93, 0xb7, 0xc4,
	0xb9, 0x9c, 0x50, 0xec, 0x7b, 0xae, 0xe3, 0x63, 0xad, 0xb6, 0xc5, 0xf9, 0xa9, 0x1b, 0x4c, 0x7d,
	0xae, 0xcf, 0x6b, 0xe2, 0x5f, 0x9b, 0x81, 0x35, 0xd2, 0x0e, 0x50, 0x07, 0x5a, 0xa2, 0x86, 0xc9,
	0xd8, 0xb9, 0x72, 0xdc, 0x5b, 0x47, 0x03, 0xf4, 0x3f, 0x78, 0xb9, 0x71, 0xf9, 0x63, 0xcf, 0x73,
	0x69, 0x80, 0x87, 0x25, 0xbd, 0xac, 0xa1, 0x81, 0x9e, 0x83, 0x2e, 0x83, 0x88, 0x73, 0x63, 0xda,
	0x64, 0x28, 0xf8, 0x47, 0xd8, 0x1c, 0x62, 0xaa, 0x35, 0xd1, 0x4b, 0x38, 0x93, 0xe8, 0x10, 0xef,
	0xe6, 0xf5, 0xcd, 0x18, 0xfb, 0x81, 0xd6, 0xe2, 0x85, 0xca, 0x80, 0xbd, 0x94, 0xda, 0xe8, 0x04,
	0xd0, 0xe6, 0x7e, 0x73, 0x1c, 0x8c, 0x5c, 0x4a, 0xde, 0xe2, 0xa1, 0xf6, 0xb8, 0xfb, 0x47, 0x05,
	0x9a, 0x7c, 0x32, 0x61, 0xbe, 0x28, 0x7f, 0x86, 0xef, 0xa5, 0x3c, 0xc5, 0xcf, 0xab, 0x45, 0x37,
	0x26, 0xfa, 0x12, 0x6a, 0x5c, 0xe6, 0x6c, 0xbd, 0x29, 0x1f, 0xed, 0x0b, 0xab, 0x64, 0xe9, 0xf3,
	0x2f, 0x3e, 0x6d, 0x46, 0xe5, 0x19, 0x74, 0x0c, 0xb5, 0x22, 0xf9, 0x91, 0xc5, 0x62, 0x17, 0x9a,
	0x54, 0x1a, 0xe8, 0x0d, 0xd4, 0xc2, 0x55, 0x31, 0xcb, 0xf5, 0xaa, 0x51, 0xe9, 0x35, 0x06, 0xff,
	0xff, 0x37, 0x4a, 0x73, 0x55, 0xcc, 0xa8, 0x3c, 0x72, 0xfa, 0x9b, 0x02, 0x47, 0x1b, 0x1f, 0x5f,
	0xef, 0x25, 0x2b, 0x66, 0x89, 0x5c, 0xef, 0x3a, 0x5d, 0x5b, 0xe8, 0x39, 0xd4, 0x97, 0x2c, 0x9a,
	0x85, 0xf1, 0x3c, 0x5f, 0x8a, 0xbc, 0xeb, 0x74, 0xeb, 0x10, 0xbf, 0x77, 0xce, 0x1d, 0x25, 0x0b,
	0x91, 0x57, 0x9d, 0x96, 0x36, 0xc7, 0xd6, 0xfb, 0x3a, 0x5d, 0xaf, 0x63, 0x69, 0x73, 0xd6, 0x68,
	0x16, 0x2e, 0x16, 0x2c, 0xbe, 0x63, 0x62, 0x0b, 0x9b, 0x74, 0xeb, 0xe8, 0x7e, 0x07, 0xf5, 0xb2,
	0xfc, 0x5d, 0x99, 0xb7, 0xa0, 0xee, 0xe0, 0x4b, 0x37, 0x20, 0x66, 0x80, 0x35, 0x05, 0x35, 0xe1,
	0x88, 0x38, 0x44, 0x5a, 0x2a, 0x07, 0xad, 0x91, 0x69, 0xdb, 0xd8, 0xb9, 0xc4, 0x5a, 0x85, 0x83,
	0xa5, 0xe2, 0xaa, 0xe8, 0x08, 0xaa, 0xb7, 0xd4, 0xf4, 0xb4, 0xda, 0xc7, 0x5f, 0x8b, 0x79, 0x95,
	0xcf, 0x07, 0x7a, 0x0c, 0x0d, 0xae, 0x92, 0xf3, 0x31, 0xb1, 0x03, 0xe2, 0x68, 0x8f, 0x90, 0x06,
	0x4d, 0xee, 0xb8, 0xa5, 0x24, 0x30, 0xcf, 0x6d, 0x7e, 0xcf, 0x33, 0x78, 0xc2, 0x3d, 0x1b, 0xd9,
	0x4f, 0xce, 0xc7, 0x17, 0x17, 0x98, 0x6a, 0xea, 0xf9, 0x6b, 0x78, 0x91, 0x64, 0x77, 0xfd, 0x30,
	0x0d, 0xa3, 0x19, 0xdb, 0xf4, 0x7e, 0x9e, 0x46, 0xf2, 0xed, 0xfc, 0x61, 0xf5, 0xee, 0x7c, 0xfb,
	0x9e, 0x8a, 0xdb, 0xf2, 0xdf, 0x15, 0xe5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0x57, 0x1d,
	0x7b, 0x64, 0x07, 0x00, 0x00,
}