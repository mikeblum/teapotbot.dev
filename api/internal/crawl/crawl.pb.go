// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: crawl.proto

package crawl

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ALPN protocols per [MDN Docs]
// [MDN Docs]: https://developer.mozilla.org/en-US/docs/Glossary/ALPN
type Protocol int32

const (
	Protocol_HTTP1 Protocol = 0 // default to http/1.1
	Protocol_HTTP2 Protocol = 1 // http/2
)

// Enum value maps for Protocol.
var (
	Protocol_name = map[int32]string{
		0: "HTTP1",
		1: "HTTP2",
	}
	Protocol_value = map[string]int32{
		"HTTP1": 0,
		"HTTP2": 1,
	}
)

func (x Protocol) Enum() *Protocol {
	p := new(Protocol)
	*p = x
	return p
}

func (x Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_crawl_proto_enumTypes[0].Descriptor()
}

func (Protocol) Type() protoreflect.EnumType {
	return &file_crawl_proto_enumTypes[0]
}

func (x Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Protocol.Descriptor instead.
func (Protocol) EnumDescriptor() ([]byte, []int) {
	return file_crawl_proto_rawDescGZIP(), []int{0}
}

// See [HAProxy L4 vs L7]
// [HAProxy L4 vs L7]: https://www.haproxy.com/blog/layer-4-and-layer-7-proxy-mode/
type Scheme int32

const (
	Scheme_TCP   Scheme = 0 // default to TCP - Layer 4
	Scheme_HTTP  Scheme = 1 // Layer 7
	Scheme_HTTPS Scheme = 2
)

// Enum value maps for Scheme.
var (
	Scheme_name = map[int32]string{
		0: "TCP",
		1: "HTTP",
		2: "HTTPS",
	}
	Scheme_value = map[string]int32{
		"TCP":   0,
		"HTTP":  1,
		"HTTPS": 2,
	}
)

func (x Scheme) Enum() *Scheme {
	p := new(Scheme)
	*p = x
	return p
}

func (x Scheme) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Scheme) Descriptor() protoreflect.EnumDescriptor {
	return file_crawl_proto_enumTypes[1].Descriptor()
}

func (Scheme) Type() protoreflect.EnumType {
	return &file_crawl_proto_enumTypes[1]
}

func (x Scheme) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Scheme.Descriptor instead.
func (Scheme) EnumDescriptor() ([]byte, []int) {
	return file_crawl_proto_rawDescGZIP(), []int{1}
}

// See [IANA HTTP Methods]
// Unsafe methods such as CONNECT, DELETE, PATCH, POST, and PUT are omitted
// [IANA HTTP Methods]: https://www.iana.org/assignments/http-methods/http-methods.xhtml
type Method int32

const (
	Method_GET     Method = 0 // default to GET - safe: yes; idempotent: yes; reference: [RFC9110, Section 9.3.1]
	Method_HEAD    Method = 1 // safe: yes; idempotent: yes; reference: [RFC9110, Section 9.3.2]
	Method_OPTIONS Method = 2 // safe: yes; idempotent: yes; reference: [RFC9110, Section 9.3.7]
	Method_TRACE   Method = 4 // safe: yes; idempotent: yes; reference: [RFC9110, Section 9.3.8]
)

// Enum value maps for Method.
var (
	Method_name = map[int32]string{
		0: "GET",
		1: "HEAD",
		2: "OPTIONS",
		4: "TRACE",
	}
	Method_value = map[string]int32{
		"GET":     0,
		"HEAD":    1,
		"OPTIONS": 2,
		"TRACE":   4,
	}
)

func (x Method) Enum() *Method {
	p := new(Method)
	*p = x
	return p
}

func (x Method) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Method) Descriptor() protoreflect.EnumDescriptor {
	return file_crawl_proto_enumTypes[2].Descriptor()
}

func (Method) Type() protoreflect.EnumType {
	return &file_crawl_proto_enumTypes[2]
}

func (x Method) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Method.Descriptor instead.
func (Method) EnumDescriptor() ([]byte, []int) {
	return file_crawl_proto_rawDescGZIP(), []int{2}
}

// [scheme:][//[userinfo@]host][/]path[?query] as defined in [net/url]
// fragments are omitted per [url#ParseRequestURI]
//
// [net/url]: https://pkg.go.dev/net/url#URL
// [url#ParseRequestURI]: https://pkg.go.dev/net/url#ParseRequestURI
type Url struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tcp or http or https
	Scheme Scheme `protobuf:"varint,1,opt,name=scheme,proto3,enum=teapotbot.Scheme" json:"scheme,omitempty"`
	// host or host:port;
	// [IANA Port Numbers]: https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.xhtml
	Host  string  `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Path  *string `protobuf:"bytes,3,opt,name=path,proto3,oneof" json:"path,omitempty"`
	Query *string `protobuf:"bytes,4,opt,name=query,proto3,oneof" json:"query,omitempty"`
}

func (x *Url) Reset() {
	*x = Url{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Url) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Url) ProtoMessage() {}

func (x *Url) ProtoReflect() protoreflect.Message {
	mi := &file_crawl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Url.ProtoReflect.Descriptor instead.
func (*Url) Descriptor() ([]byte, []int) {
	return file_crawl_proto_rawDescGZIP(), []int{0}
}

func (x *Url) GetScheme() Scheme {
	if x != nil {
		return x.Scheme
	}
	return Scheme_TCP
}

func (x *Url) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Url) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *Url) GetQuery() string {
	if x != nil && x.Query != nil {
		return *x.Query
	}
	return ""
}

type CrawlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DescriptorSet *descriptorpb.FileDescriptorSet `protobuf:"bytes,1,opt,name=descriptor_set,json=descriptorSet,proto3" json:"descriptor_set,omitempty"`
	Protocol      Protocol                        `protobuf:"varint,2,opt,name=protocol,proto3,enum=teapotbot.Protocol" json:"protocol,omitempty"`
	Method        Method                          `protobuf:"varint,3,opt,name=method,proto3,enum=teapotbot.Method" json:"method,omitempty"`
	Url           *Url                            `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	//
	// [net/http Header]: https://pkg.go.dev/net/http#Header
	Headers []*structpb.Struct `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty"`
	// UTC timestamp in ISO-8601 formatted as time.RFC3339
	// example:
	// `fmt.Println(time.Now().Format(time.RFC3339))`
	Created *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	// Cancel request if timeout (in seconds or nanoseconds) is exceeded
	Timeout *durationpb.Duration `protobuf:"bytes,7,opt,name=timeout,proto3,oneof" json:"timeout,omitempty"`
}

func (x *CrawlRequest) Reset() {
	*x = CrawlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlRequest) ProtoMessage() {}

func (x *CrawlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crawl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawlRequest.ProtoReflect.Descriptor instead.
func (*CrawlRequest) Descriptor() ([]byte, []int) {
	return file_crawl_proto_rawDescGZIP(), []int{1}
}

func (x *CrawlRequest) GetDescriptorSet() *descriptorpb.FileDescriptorSet {
	if x != nil {
		return x.DescriptorSet
	}
	return nil
}

func (x *CrawlRequest) GetProtocol() Protocol {
	if x != nil {
		return x.Protocol
	}
	return Protocol_HTTP1
}

func (x *CrawlRequest) GetMethod() Method {
	if x != nil {
		return x.Method
	}
	return Method_GET
}

func (x *CrawlRequest) GetUrl() *Url {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *CrawlRequest) GetHeaders() []*structpb.Struct {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *CrawlRequest) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *CrawlRequest) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type CrawlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DescriptorSet *descriptorpb.FileDescriptorSet `protobuf:"bytes,1,opt,name=descriptor_set,json=descriptorSet,proto3" json:"descriptor_set,omitempty"`
	// HTTP status code per [net/http status.go]
	// [net/http status.go]:  https://go.dev/src/net/http/status.go
	StatusCode int32 `protobuf:"varint,4,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	// HTTP status text per [net/http status.go]
	// [net/http status.go]:  https://go.dev/src/net/http/status.go
	StatusText string `protobuf:"bytes,5,opt,name=status_text,json=statusText,proto3" json:"status_text,omitempty"`
	// Types that are assignable to Data:
	//
	//	*CrawlResponse_Plain
	//	*CrawlResponse_Json
	Data isCrawlResponse_Data `protobuf_oneof:"data"`
	// UTC timestamp in ISO-8601 formatted as time.RFC3339
	// updated will be nil until the request is serviced
	Updated *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated,proto3,oneof" json:"updated,omitempty"`
	// if CrawlRequest.timeout is exceeded note the wall clock time
	Cancelled *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=cancelled,proto3,oneof" json:"cancelled,omitempty"`
	Timeout   *durationpb.Duration   `protobuf:"bytes,10,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *CrawlResponse) Reset() {
	*x = CrawlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlResponse) ProtoMessage() {}

func (x *CrawlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crawl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawlResponse.ProtoReflect.Descriptor instead.
func (*CrawlResponse) Descriptor() ([]byte, []int) {
	return file_crawl_proto_rawDescGZIP(), []int{2}
}

func (x *CrawlResponse) GetDescriptorSet() *descriptorpb.FileDescriptorSet {
	if x != nil {
		return x.DescriptorSet
	}
	return nil
}

func (x *CrawlResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CrawlResponse) GetStatusText() string {
	if x != nil {
		return x.StatusText
	}
	return ""
}

func (m *CrawlResponse) GetData() isCrawlResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *CrawlResponse) GetPlain() *anypb.Any {
	if x, ok := x.GetData().(*CrawlResponse_Plain); ok {
		return x.Plain
	}
	return nil
}

func (x *CrawlResponse) GetJson() *structpb.Struct {
	if x, ok := x.GetData().(*CrawlResponse_Json); ok {
		return x.Json
	}
	return nil
}

func (x *CrawlResponse) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *CrawlResponse) GetCancelled() *timestamppb.Timestamp {
	if x != nil {
		return x.Cancelled
	}
	return nil
}

func (x *CrawlResponse) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type isCrawlResponse_Data interface {
	isCrawlResponse_Data()
}

type CrawlResponse_Plain struct {
	// net/http.Response Body payload
	// `body, err := ioutil.ReadAll(resp.Body)`
	Plain *anypb.Any `protobuf:"bytes,6,opt,name=plain,proto3,oneof"`
}

type CrawlResponse_Json struct {
	// JSON-style structured metadata
	Json *structpb.Struct `protobuf:"bytes,7,opt,name=json,proto3,oneof"`
}

func (*CrawlResponse_Plain) isCrawlResponse_Data() {}

func (*CrawlResponse_Json) isCrawlResponse_Data() {}

var File_crawl_proto protoreflect.FileDescriptor

var file_crawl_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74,
	0x65, 0x61, 0x70, 0x6f, 0x74, 0x62, 0x6f, 0x74, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x29, 0x0a, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74,
	0x65, 0x61, 0x70, 0x6f, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x88, 0x01, 0x01, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x22, 0x86, 0x03, 0x0a, 0x0c, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x52,
	0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x12, 0x2f,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x13, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x6f, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x29, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x11, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x6f, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x20, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x65, 0x61, 0x70, 0x6f, 0x74,
	0x62, 0x6f, 0x74, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x31, 0x0a, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0xca, 0x03, 0x0a, 0x0d,
	0x43, 0x72, 0x61, 0x77, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a,
	0x0e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x65, 0x78, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x70, 0x6c,
	0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48,
	0x00, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x12, 0x2d, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48,
	0x00, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x02, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x2a, 0x20, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x54, 0x54, 0x50, 0x31, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x48, 0x54, 0x54, 0x50, 0x32, 0x10, 0x01, 0x2a, 0x26, 0x0a, 0x06, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x43, 0x50, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x48, 0x54, 0x54, 0x50, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x54, 0x54, 0x50, 0x53,
	0x10, 0x02, 0x2a, 0x33, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x07, 0x0a, 0x03,
	0x47, 0x45, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x45, 0x41, 0x44, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x54, 0x52, 0x41, 0x43, 0x45, 0x10, 0x04, 0x42, 0x14, 0x5a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crawl_proto_rawDescOnce sync.Once
	file_crawl_proto_rawDescData = file_crawl_proto_rawDesc
)

func file_crawl_proto_rawDescGZIP() []byte {
	file_crawl_proto_rawDescOnce.Do(func() {
		file_crawl_proto_rawDescData = protoimpl.X.CompressGZIP(file_crawl_proto_rawDescData)
	})
	return file_crawl_proto_rawDescData
}

var file_crawl_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_crawl_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_crawl_proto_goTypes = []interface{}{
	(Protocol)(0),                          // 0: teapotbot.Protocol
	(Scheme)(0),                            // 1: teapotbot.Scheme
	(Method)(0),                            // 2: teapotbot.Method
	(*Url)(nil),                            // 3: teapotbot.Url
	(*CrawlRequest)(nil),                   // 4: teapotbot.CrawlRequest
	(*CrawlResponse)(nil),                  // 5: teapotbot.CrawlResponse
	(*descriptorpb.FileDescriptorSet)(nil), // 6: google.protobuf.FileDescriptorSet
	(*structpb.Struct)(nil),                // 7: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil),          // 8: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),            // 9: google.protobuf.Duration
	(*anypb.Any)(nil),                      // 10: google.protobuf.Any
}
var file_crawl_proto_depIdxs = []int32{
	1,  // 0: teapotbot.Url.scheme:type_name -> teapotbot.Scheme
	6,  // 1: teapotbot.CrawlRequest.descriptor_set:type_name -> google.protobuf.FileDescriptorSet
	0,  // 2: teapotbot.CrawlRequest.protocol:type_name -> teapotbot.Protocol
	2,  // 3: teapotbot.CrawlRequest.method:type_name -> teapotbot.Method
	3,  // 4: teapotbot.CrawlRequest.url:type_name -> teapotbot.Url
	7,  // 5: teapotbot.CrawlRequest.headers:type_name -> google.protobuf.Struct
	8,  // 6: teapotbot.CrawlRequest.created:type_name -> google.protobuf.Timestamp
	9,  // 7: teapotbot.CrawlRequest.timeout:type_name -> google.protobuf.Duration
	6,  // 8: teapotbot.CrawlResponse.descriptor_set:type_name -> google.protobuf.FileDescriptorSet
	10, // 9: teapotbot.CrawlResponse.plain:type_name -> google.protobuf.Any
	7,  // 10: teapotbot.CrawlResponse.json:type_name -> google.protobuf.Struct
	8,  // 11: teapotbot.CrawlResponse.updated:type_name -> google.protobuf.Timestamp
	8,  // 12: teapotbot.CrawlResponse.cancelled:type_name -> google.protobuf.Timestamp
	9,  // 13: teapotbot.CrawlResponse.timeout:type_name -> google.protobuf.Duration
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_crawl_proto_init() }
func file_crawl_proto_init() {
	if File_crawl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crawl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Url); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_crawl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrawlRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_crawl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrawlResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_crawl_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_crawl_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_crawl_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CrawlResponse_Plain)(nil),
		(*CrawlResponse_Json)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_crawl_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crawl_proto_goTypes,
		DependencyIndexes: file_crawl_proto_depIdxs,
		EnumInfos:         file_crawl_proto_enumTypes,
		MessageInfos:      file_crawl_proto_msgTypes,
	}.Build()
	File_crawl_proto = out.File
	file_crawl_proto_rawDesc = nil
	file_crawl_proto_goTypes = nil
	file_crawl_proto_depIdxs = nil
}
