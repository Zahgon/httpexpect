package httpexpect

import (
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/gorilla/websocket"
)

// Response provides methods to inspect attached http.Response object.
type Response struct {
	noCopy noCopy
	config Config
	chain  *chain

	httpResp  *http.Response
	websocket *websocket.Conn
	rtt       *time.Duration

	content       []byte
	contentState  contentState
	contentMethod string

	cookies []*http.Cookie
}

type contentState int

const (
	// We didn't try to retrieve response content yet
	contentPending contentState = iota
	// We successfully retrieved response content
	contentRetreived
	// We tried to retrieve response content and failed
	contentFailed
	// We transferred body reader to user and will not use it by ourselves
	contentHijacked
)

// NewResponse returns a new Response instance.
//
// If reporter is nil, the function panics.
// If response is nil, failure is reported.
//
// If rtt is given, it defines response round-trip time to be reported
// by response.RoundTripTime().
func NewResponse(
	reporter Reporter, response *http.Response, rtt ...time.Duration,
) *Response {
	_ = "STUB: not implemented"
	return nil
}

// NewResponse returns a new Response instance with config.
//
// Requirements for config are same as for WithConfig function.
// If response is nil, failure is reported.
//
// If rtt is given, it defines response round-trip time to be reported
// by response.RoundTripTime().
func NewResponseC(
	config Config, response *http.Response, rtt ...time.Duration,
) *Response {
	_ = "STUB: not implemented"
	return nil
}

type responseOpts struct {
	config    Config
	chain     *chain
	httpResp  *http.Response
	websocket *websocket.Conn
	rtt       []time.Duration
}

func newResponse(opts responseOpts) *Response { _ = "STUB: not implemented"; return nil }

func (r *Response) getContent(opChain *chain, method string) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Raw returns underlying http.Response object.
// This is the value originally passed to NewResponse.
func (r *Response) Raw() *http.Response {
	_ = "STUB: not implemented"

	// Alias is similar to Value.Alias.
	return nil
}

func (r *Response) Alias(name string) *Response { _ = "STUB: not implemented"; return nil }

// RoundTripTime returns a new Duration instance with response round-trip time.
//
// The returned duration is the time interval starting just before request is
// sent and ending right after response is received (handshake finished for
// WebSocket request), retrieved from a monotonic clock source.
//
// Example:
//
//	resp := NewResponse(t, response, time.Duration(10000000))
//	resp.RoundTripTime().IsLt(10 * time.Millisecond)
func (r *Response) RoundTripTime() *Duration { _ = "STUB: not implemented"; return nil }

// Deprecated: use RoundTripTime instead.
func (r *Response) Duration() *Number { _ = "STUB: not implemented"; return nil }

// Status succeeds if response contains given status code.
//
// Example:
//
//	resp := NewResponse(t, response)
//	resp.Status(http.StatusOK)
func (r *Response) Status(status int) *Response { _ = "STUB: not implemented"; return nil }

// StatusRange is enum for response status ranges.
type StatusRange int

const (
	// Status1xx defines "Informational" status codes.
	Status1xx StatusRange = 100

	// Status2xx defines "Success" status codes.
	Status2xx StatusRange = 200

	// Status3xx defines "Redirection" status codes.
	Status3xx StatusRange = 300

	// Status4xx defines "Client Error" status codes.
	Status4xx StatusRange = 400

	// Status5xx defines "Server Error" status codes.
	Status5xx StatusRange = 500
)

// StatusRange succeeds if response status belongs to given range.
//
// Supported ranges:
//   - Status1xx - Informational
//   - Status2xx - Success
//   - Status3xx - Redirection
//   - Status4xx - Client Error
//   - Status5xx - Server Error
//
// See https://en.wikipedia.org/wiki/List_of_HTTP_status_codes.
//
// Example:
//
//	resp := NewResponse(t, response)
//	resp.StatusRange(Status2xx)
func (r *Response) StatusRange(rn StatusRange) *Response { _ = "STUB: not implemented"; return nil }

// StatusList succeeds if response matches with any given status code list
//
// Example:
//
//	resp := NewResponse(t, response)
//	resp.StatusList(http.StatusForbidden, http.StatusUnauthorized)
func (r *Response) StatusList(values ...int) *Response { _ = "STUB: not implemented"; return nil }

func statusCodeText(code int) string { _ = "STUB: not implemented"; return "" }

func statusRangeText(code int) string { _ = "STUB: not implemented"; return "" }

func statusListText(values []int) []interface{} { _ = "STUB: not implemented"; return nil }

// Headers returns a new Object instance with response header map.
//
// Example:
//
//	resp := NewResponse(t, response)
//	resp.Headers().Value("Content-Type").String().IsEqual("application-json")
func (r *Response) Headers() *Object { _ = "STUB: not implemented"; return nil }

// Header returns a new String instance with given header field.
//
// Example:
//
//	resp := NewResponse(t, response)
//	resp.Header("Content-Type").IsEqual("application-json")
//	resp.Header("Date").AsDateTime().IsLe(time.Now())
func (r *Response) Header(header string) *String { _ = "STUB: not implemented"; return nil }

// Cookies returns a new Array instance with all cookie names set by this response.
// Returned Array contains a String value for every cookie name.
//
// Note that this returns only cookies set by Set-Cookie headers of this response.
// It doesn't return session cookies from previous responses, which may be stored
// in a cookie jar.
//
// Example:
//
//	resp := NewResponse(t, response)
//	resp.Cookies().Contains("session")
func (r *Response) Cookies() *Array { _ = "STUB: not implemented"; return nil }

// Cookie returns a new Cookie instance with specified cookie from response.
//
// Note that this returns only cookies set by Set-Cookie headers of this response.
// It doesn't return session cookies from previous responses, which may be stored
// in a cookie jar.
//
// Example:
//
//	resp := NewResponse(t, response)
//	resp.Cookie("session").Domain().IsEqual("example.com")
func (r *Response) Cookie(name string) *Cookie { _ = "STUB: not implemented"; return nil }

// Websocket returns Websocket instance for interaction with WebSocket server.
//
// May be called only if the WithWebsocketUpgrade was called on the request.
// That is responsibility of the caller to explicitly disconnect websocket after use.
//
// Example:
//
//	req := NewRequestC(config, "GET", "/path")
//	req.WithWebsocketUpgrade()
//	ws := req.Expect().Websocket()
//	defer ws.Disconnect()
func (r *Response) Websocket() *Websocket { _ = "STUB: not implemented"; return nil }

// Reader returns the body reader from the response.
//
// This method is mutually exclusive with methods that read entire
// response body, like Text, Body, JSON, etc. It can be used when
// you need to parse body manually or retrieve infinite responses.
//
// Example:
//
//	resp := NewResponse(t, response)
//	reader := resp.Reader()
func (r *Response) Reader() io.ReadCloser { _ = "STUB: not implemented"; return *new(io.ReadCloser) }

// Body returns a new String instance with response body.
//
// Example:
//
//	resp := NewResponse(t, response)
//	resp.Body().NotEmpty()
//	resp.Body().Length().IsEqual(100)
func (r *Response) Body() *String { _ = "STUB: not implemented"; return nil }

// NoContent succeeds if response contains empty Content-Type header and
// empty body.
func (r *Response) NoContent() *Response { _ = "STUB: not implemented"; return nil }

// HasContentType succeeds if response contains Content-Type header with given
// media type and charset.
//
// If charset is omitted, and mediaType is non-empty, Content-Type header
// should contain empty or utf-8 charset.
//
// If charset is omitted, and mediaType is also empty, Content-Type header
// should contain no charset.
func (r *Response) HasContentType(mediaType string, charset ...string) *Response {
	_ = "STUB: not implemented"
	return nil
}

// HasContentEncoding succeeds if response has exactly given Content-Encoding list.
// Common values are empty, "gzip", "compress", "deflate", "identity" and "br".
func (r *Response) HasContentEncoding(encoding ...string) *Response {
	_ = "STUB: not implemented"
	return nil
}

// HasTransferEncoding succeeds if response contains given Transfer-Encoding list.
// Common values are empty, "chunked" and "identity".
func (r *Response) HasTransferEncoding(encoding ...string) *Response {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: use HasContentType instead.
func (r *Response) ContentType(mediaType string, charset ...string) *Response {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: use HasContentEncoding instead.
func (r *Response) ContentEncoding(encoding ...string) *Response {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: use HasTransferEncoding instead.
func (r *Response) TransferEncoding(encoding ...string) *Response {
	_ = "STUB: not implemented"
	return nil
}

// ContentOpts define parameters for matching the response content parameters.
type ContentOpts struct {
	// The media type Content-Type part, e.g. "application/json"
	MediaType string
	// The character set Content-Type part, e.g. "utf-8"
	Charset string
}

// Text returns a new String instance with response body.
//
// Text succeeds if response contains "text/plain" Content-Type header
// with empty or "utf-8" charset.
//
// Example:
//
//	resp := NewResponse(t, response)
//	resp.Text().IsEqual("hello, world!")
//	resp.Text(ContentOpts{
//	  MediaType: "text/plain",
//	}).IsEqual("hello, world!")
func (r *Response) Text(options ...ContentOpts) *String { _ = "STUB: not implemented"; return nil }

// Form returns a new Object instance with form decoded from response body.
//
// Form succeeds if response contains "application/x-www-form-urlencoded"
// Content-Type header and if form may be decoded from response body.
// Decoding is performed using https://github.com/ajg/form.
//
// Example:
//
//	resp := NewResponse(t, response)
//	resp.Form().Value("foo").IsEqual("bar")
//	resp.Form(ContentOpts{
//	  MediaType: "application/x-www-form-urlencoded",
//	}).Value("foo").IsEqual("bar")
func (r *Response) Form(options ...ContentOpts) *Object { _ = "STUB: not implemented"; return nil }

func (r *Response) getForm(
	opChain *chain, method string, options ...ContentOpts,
) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// JSON returns a new Value instance with JSON decoded from response body.
//
// JSON succeeds if response contains "application/json" Content-Type header
// with empty or "utf-8" charset and if JSON may be decoded from response body.
//
// Example:
//
//	resp := NewResponse(t, response)
//	resp.JSON().Array().ConsistsOf("foo", "bar")
//	resp.JSON(ContentOpts{
//	  MediaType: "application/json",
//	}).Array.ConsistsOf("foo", "bar")
func (r *Response) JSON(options ...ContentOpts) *Value { _ = "STUB: not implemented"; return nil }

func (r *Response) getJSON(
	opChain *chain, method string, options ...ContentOpts,
) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// JSONP returns a new Value instance with JSONP decoded from response body.
//
// JSONP succeeds if response contains "application/javascript" Content-Type
// header with empty or "utf-8" charset and response body of the following form:
//
//	callback(<valid json>);
//
// or:
//
//	callback(<valid json>)
//
// Whitespaces are allowed.
//
// Example:
//
//	resp := NewResponse(t, response)
//	resp.JSONP("myCallback").Array().ConsistsOf("foo", "bar")
//	resp.JSONP("myCallback", ContentOpts{
//	  MediaType: "application/javascript",
//	}).Array().ConsistsOf("foo", "bar")
func (r *Response) JSONP(callback string, options ...ContentOpts) *Value {
	_ = "STUB: not implemented"
	return nil
}

var (
	jsonp = regexp.MustCompile(`^\s*([^\s(]+)\s*\((.*)\)\s*;*\s*$`)
)

func (r *Response) getJSONP(
	opChain *chain, method string, callback string, options ...ContentOpts,
) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (r *Response) checkContentOptions(
	opChain *chain, options []ContentOpts, expectedType string, expectedCharset ...string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Response) checkContentType(
	opChain *chain, expectedType string, expectedCharset ...string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Response) checkEqual(
	opChain *chain, what string, expected, actual interface{},
) bool {
	_ = "STUB: not implemented"
	return false
}

type errBodyReader struct {
	err error
}

func (r errBodyReader) Read(_ []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r errBodyReader) Close() error { _ = "STUB: not implemented"; return nil }
