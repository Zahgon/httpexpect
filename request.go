package httpexpect

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Request provides methods to incrementally build http.Request object,
// send it, and receive response.
type Request struct {
	mu sync.Mutex

	config Config
	chain  *chain

	redirectPolicy RedirectPolicy
	maxRedirects   int

	retryPolicy   RetryPolicy
	retryPolicyFn func(*http.Response, error) bool
	maxRetries    int
	minRetryDelay time.Duration
	maxRetryDelay time.Duration
	sleepFn       func(d time.Duration) <-chan time.Time

	timeout time.Duration

	httpReq *http.Request
	path    string

	query        url.Values
	queryEncoder QueryEncoder

	form        url.Values
	formBuf     *bytes.Buffer
	multipart   *multipart.Writer
	multipartFn func(w io.Writer) *multipart.Writer

	bodySetter   string
	typeSetter   string
	forceType    bool
	expectCalled bool

	wsUpgrade bool

	transformers []func(*http.Request)
	matchers     []func(*Response)
}

// Deprecated: use NewRequestC instead.
func NewRequest(config Config, method, path string, pathargs ...interface{}) *Request {
	_ = "STUB: not implemented"
	return nil
}

// NewRequestC returns a new Request instance.
//
// Requirements for config are same as for WithConfig function.
//
// method defines the HTTP method (GET, POST, PUT, etc.). path defines url path.
//
// Simple interpolation is allowed for {named} parameters in path:
//   - if pathargs is given, it's used to substitute first len(pathargs) parameters,
//     regardless of their names
//   - if WithPath() or WithPathObject() is called, it's used to substitute given
//     parameters by name
//
// For example:
//
//	req := NewRequestC(config, "POST", "/repos/{user}/{repo}", "gavv", "httpexpect")
//	// path will be "/repos/gavv/httpexpect"
//
// Or:
//
//	req := NewRequestC(config, "POST", "/repos/{user}/{repo}")
//	req.WithPath("user", "gavv")
//	req.WithPath("repo", "httpexpect")
//	// path will be "/repos/gavv/httpexpect"
//
// After interpolation, path is urlencoded and appended to Config.BaseURL,
// separated by slash. If BaseURL ends with a slash and path (after interpolation)
// starts with a slash, only single slash is inserted.
func NewRequestC(config Config, method, path string, pathargs ...interface{}) *Request {
	_ = "STUB: not implemented"
	return nil
}

func newRequest(
	parent *chain, config Config, method, path string, pathargs ...interface{},
) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) initPath(opChain *chain, path string, pathargs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (r *Request) initReq(opChain *chain, method string) { _ = "STUB: not implemented"; return }

// Alias is similar to Value.Alias.
func (r *Request) Alias(name string) *Request { _ = "STUB: not implemented"; return nil }

// WithName sets convenient request name.
// This name will be included in assertion reports for this request.
// It does not affect assertion chain path, inlike Alias.
//
// Example:
//
//	req := NewRequestC(config, "POST", "/api/login")
//	req.WithName("Login Request")
func (r *Request) WithName(name string) *Request { _ = "STUB: not implemented"; return nil }

// WithReporter sets reporter to be used for this request.
//
// The new reporter overwrites AssertionHandler.
// The new AssertionHandler is DefaultAssertionHandler with specified reporter,
// existing Config.Formatter and nil Logger.
// It will be used to report formatted fatal failure messages.
//
// Example:
//
//	req := NewRequestC(config, "GET", "http://example.com/path")
//	req.WithReporter(t)
func (r *Request) WithReporter(reporter Reporter) *Request { _ = "STUB: not implemented"; return nil }

// WithAssertionHandler sets assertion handler to be used for this request.
//
// The new handler overwrites assertion handler that will be used
// by Request and its children (Response, Body, etc.).
// It will be used to format and report test Failure or Success.
//
// Example:
//
//	req := NewRequestC(config, "GET", "http://example.com/path")
//	req.WithAssertionHandler(&DefaultAssertionHandler{
//		Reporter:  reporter,
//		Formatter: formatter,
//	})
func (r *Request) WithAssertionHandler(handler AssertionHandler) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WithMatcher attaches a matcher to the request.
// All attached matchers are invoked in the Expect method for a newly
// created Response.
//
// Example:
//
//	req := NewRequestC(config, "GET", "/path")
//	req.WithMatcher(func (resp *httpexpect.Response) {
//		resp.Header("API-Version").NotEmpty()
//	})
func (r *Request) WithMatcher(matcher func(*Response)) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WithTransformer attaches a transform to the Request.
// All attachhed transforms are invoked in the Expect methods for
// http.Request struct, after it's encoded and before it's sent.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithTransformer(func(r *http.Request) { r.Header.Add("foo", "bar") })
func (r *Request) WithTransformer(transformer func(*http.Request)) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WithClient sets client.
//
// The new client overwrites Config.Client. It will be used once to send the
// request and receive a response.
//
// Example:
//
//	req := NewRequestC(config, "GET", "/path")
//	req.WithClient(&http.Client{
//	  Transport: &http.Transport{
//		DisableCompression: true,
//	  },
//	})
func (r *Request) WithClient(client Client) *Request { _ = "STUB: not implemented"; return nil }

// WithHandler configures client to invoke the given handler directly.
//
// If Config.Client is http.Client, then only its Transport field is overwritten
// because the client may contain some state shared among requests like a cookie
// jar. Otherwise, the whole client is overwritten with a new client.
//
// Example:
//
//	req := NewRequestC(config, "GET", "/path")
//	req.WithHandler(myServer.someHandler)
func (r *Request) WithHandler(handler http.Handler) *Request { _ = "STUB: not implemented"; return nil }

// WithContext sets the context.
//
// Config.Context will be overwritten.
//
// Any retries will stop after one is cancelled.
// If the intended behavior is to continue any further retries, use WithTimeout.
//
// Example:
//
//	ctx, _ = context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
//	req := NewRequestC(config, "GET", "/path")
//	req.WithContext(ctx)
//	req.Expect().Status(http.StatusOK)
func (r *Request) WithContext(ctx context.Context) *Request { _ = "STUB: not implemented"; return nil }

// WithTimeout sets a timeout duration for the request.
//
// Will attach to the request a context.WithTimeout around the Config.Context
// or any context set WithContext. If these are nil, the new context will be
// created on top of a context.Background().
//
// Any retries will continue after one is cancelled.
// If the intended behavior is to stop any further retries, use WithContext or
// Config.Context.
//
// Example:
//
//	req := NewRequestC(config, "GET", "/path")
//	req.WithTimeout(time.Duration(3)*time.Second)
//	req.Expect().Status(http.StatusOK)
func (r *Request) WithTimeout(timeout time.Duration) *Request {
	_ = "STUB: not implemented"
	return nil
}

// RedirectPolicy defines how redirection responses are handled.
//
// Status codes 307, 308 require resending body. They are followed only if
// redirect policy is FollowAllRedirects.
//
// Status codes 301, 302, 303 don't require resending body. On such redirect,
// http.Client automatically switches HTTP method to GET, if it's not GET or
// HEAD already. These redirects are followed if redirect policy is either
// FollowAllRedirects or FollowRedirectsWithoutBody.
//
// Default redirect policy is FollowRedirectsWithoutBody.
type RedirectPolicy int

const (
	// indicates that WithRedirectPolicy was not called
	defaultRedirectPolicy RedirectPolicy = iota

	// DontFollowRedirects forbids following any redirects.
	// Redirection response is returned to the user and can be inspected.
	DontFollowRedirects

	// FollowAllRedirects allows following any redirects, including those
	// which require resending body.
	FollowAllRedirects

	// FollowRedirectsWithoutBody allows following only redirects which
	// don't require resending body.
	// If redirect requires resending body, it's not followed, and redirection
	// response is returned instead.
	FollowRedirectsWithoutBody
)

// WithRedirectPolicy sets policy for redirection response handling.
//
// How redirect is handled depends on both response status code and
// redirect policy. See comments for RedirectPolicy for details.
//
// Default redirect policy is defined by Client implementation.
// Default behavior of http.Client corresponds to FollowRedirectsWithoutBody.
//
// This method can be used only if Client interface points to
// *http.Client struct, since we rely on it in redirect handling.
//
// Example:
//
//	req1 := NewRequestC(config, "POST", "/path")
//	req1.WithRedirectPolicy(FollowAllRedirects)
//	req1.Expect().Status(http.StatusOK)
//
//	req2 := NewRequestC(config, "POST", "/path")
//	req2.WithRedirectPolicy(DontFollowRedirects)
//	req2.Expect().Status(http.StatusPermanentRedirect)
func (r *Request) WithRedirectPolicy(policy RedirectPolicy) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WithMaxRedirects sets maximum number of redirects to follow.
//
// If the number of redirects exceedes this limit, request is failed.
//
// Default limit is defined by Client implementation.
// Default behavior of http.Client corresponds to maximum of 10-1 redirects.
//
// This method can be used only if Client interface points to
// *http.Client struct, since we rely on it in redirect handling.
//
// Example:
//
//	req1 := NewRequestC(config, "POST", "/path")
//	req1.WithMaxRedirects(1)
//	req1.Expect().Status(http.StatusOK)
func (r *Request) WithMaxRedirects(maxRedirects int) *Request {
	_ = "STUB: not implemented"
	return nil
}

// RetryPolicy defines how failed requests are retried.
//
// Whether a request is retried depends on error type (if any), response
// status code (if any), and retry policy.
type RetryPolicy int

const (
	// indicates that WithRetryPolicy was not called
	defaultRetryPolicy RetryPolicy = iota

	// DontRetry disables retrying at all.
	DontRetry

	// Deprecated: use RetryTimeoutErrors instead.
	RetryTemporaryNetworkErrors

	// Deprecated: use RetryTimeoutAndServerErrors instead.
	RetryTemporaryNetworkAndServerErrors

	// RetryTimeoutErrors enables retrying of timeout errors.
	// Retry happens if Client returns net.Error and its Timeout() method
	// returns true.
	RetryTimeoutErrors

	// RetryTimeoutAndServerErrors enables retrying of network timeout errors,
	// as well as 5xx status codes.
	RetryTimeoutAndServerErrors

	// RetryAllErrors enables retrying of any error or 4xx/5xx status code.
	RetryAllErrors
)

// WithRetryPolicy sets policy for retries.
//
// Whether a request is retried depends on error type (if any), response
// status code (if any), and retry policy.
//
// How much retry attempts happens is defined by WithMaxRetries().
// How much to wait between attempts is defined by WithRetryDelay().
//
// Default retry policy is RetryTimeoutAndServerErrors, but
// default maximum number of retries is zero, so no retries happen
// unless WithMaxRetries() is called.
//
// Example:
//
//	req := NewRequestC(config, "POST", "/path")
//	req.WithRetryPolicy(RetryAllErrors)
//	req.Expect().Status(http.StatusOK)
func (r *Request) WithRetryPolicy(policy RetryPolicy) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WithRetryPolicyFunc sets a function to replace built-in policies
// with user-defined policy.
//
// The function expects you to return true to perform a retry. And false to
// not perform a retry.
//
// Example:
//
//	req := NewRequestC(config, "POST", "/path")
//	req.WithRetryPolicyFunc(func(res *http.Response, err error) bool {
//		return resp.StatusCode == http.StatusTeapot
//	})
func (r *Request) WithRetryPolicyFunc(
	fn func(res *http.Response, err error) bool,
) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WithMaxRetries sets maximum number of retry attempts.
//
// After first request failure, additional retry attempts may happen,
// depending on the retry policy.
//
// Setting this to zero disables retries, i.e. only one request is sent.
// Setting this to N enables retries, and up to N+1 requests may be sent.
//
// Default number of retries is zero, i.e. retries are disabled.
//
// Example:
//
//	req := NewRequestC(config, "POST", "/path")
//	req.WithMaxRetries(1)
//	req.Expect().Status(http.StatusOK)
func (r *Request) WithMaxRetries(maxRetries int) *Request { _ = "STUB: not implemented"; return nil }

// WithRetryDelay sets minimum and maximum delay between retries.
//
// If multiple retry attempts happen, delay between attempts starts from
// minDelay and then grows exponentionally until it reaches maxDelay.
//
// Default delay range is [50ms; 5s].
//
// Example:
//
//	req := NewRequestC(config, "POST", "/path")
//	req.WithRetryDelay(time.Second, time.Minute)
//	req.Expect().Status(http.StatusOK)
func (r *Request) WithRetryDelay(minDelay, maxDelay time.Duration) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WithWebsocketUpgrade enables upgrades the connection to websocket.
//
// At least the following fields are added to the request header:
//
//	Upgrade: websocket
//	Connection: Upgrade
//
// The actual set of header fields is define by the protocol implementation
// in the gorilla/websocket package.
//
// The user should then call the Response.Websocket() method which returns
// the Websocket instance. This instance can be used to send messages to the
// server, to inspect the received messages, and to close the websocket.
//
// Example:
//
//	req := NewRequestC(config, "GET", "/path")
//	req.WithWebsocketUpgrade()
//	ws := req.Expect().Status(http.StatusSwitchingProtocols).Websocket()
//	defer ws.Disconnect()
func (r *Request) WithWebsocketUpgrade() *Request { _ = "STUB: not implemented"; return nil }

// WithWebsocketDialer sets the custom websocket dialer.
//
// The new dialer overwrites Config.WebsocketDialer. It will be used once to establish
// the WebSocket connection and receive a response of handshake result.
//
// Example:
//
//	req := NewRequestC(config, "GET", "/path")
//	req.WithWebsocketUpgrade()
//	req.WithWebsocketDialer(&websocket.Dialer{
//	  EnableCompression: false,
//	})
//	ws := req.Expect().Status(http.StatusSwitchingProtocols).Websocket()
//	defer ws.Disconnect()
func (r *Request) WithWebsocketDialer(dialer WebsocketDialer) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WithPath substitutes named parameters in url path.
//
// value is converted to string using fmt.Sprint(). If there is no named
// parameter '{key}' in url path, failure is reported.
//
// Named parameters are case-insensitive.
//
// Example:
//
//	req := NewRequestC(config, "POST", "/repos/{user}/{repo}")
//	req.WithPath("user", "gavv")
//	req.WithPath("repo", "httpexpect")
//	// path will be "/repos/gavv/httpexpect"
func (r *Request) WithPath(key string, value interface{}) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WithPathObject substitutes multiple named parameters in url path.
//
// object should be map or struct. If object is struct, it's converted
// to map using https://github.com/fatih/structs. Structs may contain
// "path" struct tag, similar to "json" struct tag for json.Marshal().
//
// Each map value is converted to string using fmt.Sprint(). If there
// is no named parameter for some map '{key}' in url path, failure is
// reported.
//
// Named parameters are case-insensitive.
//
// Example:
//
//	type MyPath struct {
//		Login string `path:"user"`
//		Repo  string
//	}
//
//	req := NewRequestC(config, "POST", "/repos/{user}/{repo}")
//	req.WithPathObject(MyPath{"gavv", "httpexpect"})
//	// path will be "/repos/gavv/httpexpect"
//
//	req := NewRequestC(config, "POST", "/repos/{user}/{repo}")
//	req.WithPathObject(map[string]string{"user": "gavv", "repo": "httpexpect"})
//	// path will be "/repos/gavv/httpexpect"
func (r *Request) WithPathObject(object interface{}) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) withPath(opChain *chain, key string, value interface{}) {
	_ = "STUB: not implemented"
	return
}

// WithQuery adds query parameter to request URL.
//
// value is converted to string using fmt.Sprint() and urlencoded.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithQuery("a", 123)
//	req.WithQuery("b", "foo")
//	// URL is now http://example.com/path?a=123&b=foo
func (r *Request) WithQuery(key string, value interface{}) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WithQueryObject adds multiple query parameters to request URL.
//
// object is converted to query string using github.com/google/go-querystring
// if it's a struct or pointer to struct, or github.com/ajg/form otherwise.
//
// Various object types are supported. Structs may contain "url" struct tag,
// similar to "json" struct tag for json.Marshal().
//
// You can force usage of specific encoder (google/go-querystring, ajg/form)
// or its variation using WithQueryEncoder() method.
//
// Example:
//
//	type MyURL struct {
//		A int    `url:"a"`
//		B string `url:"b"`
//	}
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithQueryObject(MyURL{A: 123, B: "foo"})
//	// URL is now http://example.com/path?a=123&b=foo
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithQueryObject(map[string]interface{}{"a": 123, "b": "foo"})
//	// URL is now http://example.com/path?a=123&b=foo
func (r *Request) WithQueryObject(object interface{}) *Request {
	_ = "STUB: not implemented"
	return nil
}

// can't happen

// QueryEncoder defines how to encode object into query in WithQueryObject().
// If not set, appropriate encoder is selected automatically:
//   - for structs and pointer to structs, QueryEncoderGoogle is used
//   - otherwise QueryEncoderAjg is used
type QueryEncoder int

const (
	// indicates that WithQueryEncoder was not called
	defaultQueryEncoder QueryEncoder = iota

	// QueryEncoderGoogle encodes query using google/go-querystring package.
	// Query object should be a struct annotated with `url` tags.
	// See https://github.com/google/go-querystring.
	QueryEncoderGoogle

	// QueryEncoderForm encodes query using ajg/form package.
	// Query object can be arbitrary type including maps and structs
	// annotated with `form` tags.
	// See https://github.com/ajg/form.
	QueryEncoderForm

	// QueryEncoderFormKeepZeros is same as QueryEncoderForm, but with KeepZeros
	// flag enabled. With this flag, encoder keeps zero (default) values in their
	// literal form when encoding, and returns the former; by default zero values
	// are not kept, but are rather encoded as the empty string.
	QueryEncoderFormKeepZeros
)

// WithQueryEncoder forces use of specific encoder or encoder mode when
// an object is converted to query string in WithQueryObject().
//
// See documentation for QueryEncoder enum for evailable encoders.
//
// In particular, you can use it to force ajg/form encoder for structs instead
// of google/go-querystring, or to use KeepZeros mode of the ajg/form encoder.
//
// Example:
//
//	type MyURL struct {
//		A int    `form:"a"`
//		B string `form:"b"`
//	}
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithQueryEncoder(QueryEncoderForm)
//	req.WithQueryObject(MyURL{A: 123, B: "foo"})
//	// URL is now http://example.com/path?a=123&b=foo
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithQueryEncoder(QueryEncoderFormKeepZeros)
//	req.WithQueryObject(map[string]interface{}{"a": 0, "b": 0})
//	// URL is now http://example.com/path?a=0&b=0
func (r *Request) WithQueryEncoder(encoder QueryEncoder) *Request {
	_ = "STUB: not implemented"
	return nil
}

func selectQueryEncoder(object interface{}) QueryEncoder {
	_ = "STUB: not implemented"
	return *new(QueryEncoder)
}

// Use google/go-querystring.

// Use ajg/form.

// WithQueryString parses given query string and adds it to request URL.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithQuery("a", 11)
//	req.WithQueryString("b=22&c=33")
//	// URL is now http://example.com/path?a=11&bb=22&c=33
func (r *Request) WithQueryString(query string) *Request { _ = "STUB: not implemented"; return nil }

// WithURL sets request URL.
//
// This URL overwrites Config.BaseURL. Request path passed to request constructor
// is appended to this URL, separated by slash if necessary.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "/path")
//	req.WithURL("http://example.com")
//	// URL is now http://example.com/path
func (r *Request) WithURL(urlStr string) *Request { _ = "STUB: not implemented"; return nil }

// WithHeaders adds given headers to request.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithHeaders(map[string]string{
//		"Content-Type": "application/json",
//	})
func (r *Request) WithHeaders(headers map[string]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WithHeader adds given single header to request.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithHeader("Content-Type", "application/json")
func (r *Request) WithHeader(k, v string) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) withHeader(k, v string) { _ = "STUB: not implemented"; return }

// WithCookies adds given cookies to request.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithCookies(map[string]string{
//		"foo": "aa",
//		"bar": "bb",
//	})
func (r *Request) WithCookies(cookies map[string]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WithCookie adds given single cookie to request.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithCookie("name", "value")
func (r *Request) WithCookie(k, v string) *Request { _ = "STUB: not implemented"; return nil }

// WithBasicAuth sets the request's Authorization header to use HTTP
// Basic Authentication with the provided username and password.
//
// With HTTP Basic Authentication the provided username and password
// are not encrypted.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithBasicAuth("john", "secret")
func (r *Request) WithBasicAuth(username, password string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WithHost sets request host to given string.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithHost("example.com")
func (r *Request) WithHost(host string) *Request { _ = "STUB: not implemented"; return nil }

// WithProto sets HTTP protocol version.
//
// proto should have form of "HTTP/{major}.{minor}", e.g. "HTTP/1.1".
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithProto("HTTP/2.0")
func (r *Request) WithProto(proto string) *Request { _ = "STUB: not implemented"; return nil }

// WithChunked enables chunked encoding and sets request body reader.
//
// Expect() will read all available data from given reader. Content-Length
// is not set, and "chunked" Transfer-Encoding is used.
//
// If protocol version is not at least HTTP/1.1 (required for chunked
// encoding), failure is reported.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/upload")
//	fh, _ := os.Open("data")
//	defer fh.Close()
//	req.WithHeader("Content-Type", "application/octet-stream")
//	req.WithChunked(fh)
func (r *Request) WithChunked(reader io.Reader) *Request { _ = "STUB: not implemented"; return nil }

// WithBytes sets request body to given slice of bytes.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithHeader("Content-Type", "application/json")
//	req.WithBytes([]byte(`{"foo": 123}`))
func (r *Request) WithBytes(b []byte) *Request { _ = "STUB: not implemented"; return nil }

// WithText sets Content-Type header to "text/plain; charset=utf-8" and
// sets body to given string.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithText("hello, world!")
func (r *Request) WithText(s string) *Request { _ = "STUB: not implemented"; return nil }

// WithJSON sets Content-Type header to "application/json; charset=utf-8"
// and sets body to object, marshaled using json.Marshal().
//
// Example:
//
//	type MyJSON struct {
//		Foo int `json:"foo"`
//	}
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithJSON(MyJSON{Foo: 123})
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithJSON(map[string]interface{}{"foo": 123})
func (r *Request) WithJSON(object interface{}) *Request { _ = "STUB: not implemented"; return nil }

// WithForm sets Content-Type header to "application/x-www-form-urlencoded"
// or (if WithMultipart() was called) "multipart/form-data", converts given
// object to url.Values using github.com/ajg/form, and adds it to request body.
//
// Various object types are supported, including maps and structs. Structs may
// contain "form" struct tag, similar to "json" struct tag for json.Marshal().
// See https://github.com/ajg/form for details.
//
// Multiple WithForm(), WithFormField(), and WithFile() calls may be combined.
// If WithMultipart() is called, it should be called first.
//
// Example:
//
//	type MyForm struct {
//		Foo int `form:"foo"`
//	}
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithForm(MyForm{Foo: 123})
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithForm(map[string]interface{}{"foo": 123})
func (r *Request) WithForm(object interface{}) *Request { _ = "STUB: not implemented"; return nil }

// WithFormField sets Content-Type header to "application/x-www-form-urlencoded"
// or (if WithMultipart() was called) "multipart/form-data", converts given
// value to string using fmt.Sprint(), and adds it to request body.
//
// Multiple WithForm(), WithFormField(), and WithFile() calls may be combined.
// If WithMultipart() is called, it should be called first.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithFormField("foo", 123).
//		WithFormField("bar", 456)
func (r *Request) WithFormField(key string, value interface{}) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WithFile sets Content-Type header to "multipart/form-data", reads given
// file and adds its contents to request body.
//
// If reader is given, it's used to read file contents. Otherwise, os.Open()
// is used to read a file with given path.
//
// Multiple WithForm(), WithFormField(), and WithFile() calls may be combined.
// WithMultipart() should be called before WithFile(), otherwise WithFile()
// fails.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithFile("avatar", "./john.png")
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	fh, _ := os.Open("./john.png")
//	req.WithMultipart().
//		WithFile("avatar", "john.png", fh)
//	fh.Close()
func (r *Request) WithFile(key, path string, reader ...io.Reader) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WithFileBytes is like WithFile, but uses given slice of bytes as the
// file contents.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	fh, _ := os.Open("./john.png")
//	b, _ := io.ReadAll(fh)
//	req.WithMultipart().
//		WithFileBytes("avatar", "john.png", b)
//	fh.Close()
func (r *Request) WithFileBytes(key, path string, data []byte) *Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) withFile(
	opChain *chain, method, key, path string, reader ...io.Reader,
) {
	_ = "STUB: not implemented"
	return
}

// WithMultipart sets Content-Type header to "multipart/form-data".
//
// After this call, WithForm() and WithFormField() switch to multipart
// form instead of urlencoded form.
//
// If WithMultipart() is called, it should be called before WithForm(),
// WithFormField(), and WithFile().
//
// WithFile() always requires WithMultipart() to be called first.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithMultipart().
//		WithForm(map[string]interface{}{"foo": 123})
func (r *Request) WithMultipart() *Request { _ = "STUB: not implemented"; return nil }

// Expect constructs http.Request, sends it, receives http.Response, and
// returns a new Response instance.
//
// Request is sent using Client interface, or WebsocketDialer in case of
// WebSocket request.
//
// After calling Expect, there should not be any more calls of Expect or
// other WithXXX methods on the same Request instance.
//
// Example:
//
//	req := NewRequestC(config, "PUT", "http://example.com/path")
//	req.WithJSON(map[string]interface{}{"foo": 123})
//	resp := req.Expect()
//	resp.Status(http.StatusOK)
func (r *Request) Expect() *Response { _ = "STUB: not implemented"; return nil }

func (r *Request) expect(opChain *chain) *Response { _ = "STUB: not implemented"; return nil }

// after return from prepare(), all subsequent calls to WithXXX and Expect will
// abort early due to checkOrder(); so we can safely proceed without a lock

func (r *Request) prepare(opChain *chain) bool { _ = "STUB: not implemented"; return false }

func (r *Request) execute(opChain *chain) *Response { _ = "STUB: not implemented"; return nil }

func (r *Request) encodeRequest(opChain *chain) bool { _ = "STUB: not implemented"; return false }

var websocketErr = `webocket request can not have body:
  body was set by %s
  webocket was enabled by WithWebsocketUpgrade()`

func (r *Request) encodeWebsocketRequest(opChain *chain) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Request) sendRequest(opChain *chain) (*http.Response, time.Duration) {
	_ = "STUB: not implemented"
	return nil, *new(time.Duration)
}

func (r *Request) sendWebsocketRequest(opChain *chain) (
	*http.Response, *websocket.Conn, time.Duration,
) {
	_ = "STUB: not implemented"
	return nil, nil, *new(time.Duration)
}

func (r *Request) retryRequest(reqFunc func() (*http.Response, error)) (
	*http.Response, time.Duration, error,
) {
	_ = "STUB: not implemented"
	return nil, *new(time.Duration), nil
}

// Make a copy to avoid accidental modification of request.
// In particular, httputil.DumpRequest reads sets request body into a buffer
// and set req.Body to a wrapper that will re-read body from buffer.
// It breaks our bodyWrapper logic as we don't expect that someone will
// replace bodyWrapper with something else.

// Make a copy to avoid accidental modification of request.
// See comment above.

func (r *Request) shouldRetry(resp *http.Response, err error) bool {
	_ = "STUB: not implemented"
	return false
}

// set by WithRetryPolicyFunc

// Deprecated

//nolint

// set by WithRetryPolicy

// can't happen

func (r *Request) setupRedirects(opChain *chain) { _ = "STUB: not implemented"; return }

var typeErr = `ambiguous request "Content-Type" header values:
  first set by %s:
    %q
  then replaced by %s:
    %q`

func (r *Request) setType(
	opChain *chain, newSetter, newType string, overwrite bool,
) {
	_ = "STUB: not implemented"
	return
}

var bodyErr = `ambiguous request body contents:
  first set by %s
  then replaced by %s`

func (r *Request) setBody(
	opChain *chain, setter string, reader io.Reader, length int, overwrite bool,
) {
	_ = "STUB: not implemented"
	return
}

func (r *Request) checkOrder(opChain *chain, funcCall string) bool {
	_ = "STUB: not implemented"
	return false
}

func concatPaths(a, b string) string { _ = "STUB: not implemented"; return "" }

func mustWrite(w io.Writer, s string) { _ = "STUB: not implemented"; return }
