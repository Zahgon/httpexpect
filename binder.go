package httpexpect

import (
	"crypto/tls"
	"net"
	"net/http"

	"github.com/valyala/fasthttp"
)

// Binder implements networkless http.RoundTripper attached directly to
// http.Handler.
//
// Binder emulates network communication by invoking given http.Handler
// directly. It passes httptest.ResponseRecorder as http.ResponseWriter
// to the handler, and then constructs http.Response from recorded data.
type Binder struct {
	// HTTP handler invoked for every request.
	Handler http.Handler
	// TLS connection state used for https:// requests.
	TLS *tls.ConnectionState
}

// NewBinder returns a new Binder given a http.Handler.
//
// Example:
//
//	client := &http.Client{
//		Transport: NewBinder(handler),
//	}
func NewBinder(handler http.Handler) Binder { _ = "STUB: not implemented"; return *new(Binder) }

// RoundTrip implements http.RoundTripper.RoundTrip.
func (binder Binder) RoundTrip(origReq *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FastBinder implements networkless http.RoundTripper attached directly
// to fasthttp.RequestHandler.
//
// FastBinder emulates network communication by invoking given fasthttp.RequestHandler
// directly. It converts http.Request to fasthttp.Request, invokes handler, and then
// converts fasthttp.Response to http.Response.
type FastBinder struct {
	// FastHTTP handler invoked for every request.
	Handler fasthttp.RequestHandler
	// TLS connection state used for https:// requests.
	TLS *tls.ConnectionState
	// If non-nil, fasthttp.RequestCtx.Logger() will print messages to it.
	Logger Logger
}

// NewFastBinder returns a new FastBinder given a fasthttp.RequestHandler.
//
// Example:
//
//	client := &http.Client{
//		Transport: NewFastBinder(fasthandler),
//	}
func NewFastBinder(handler fasthttp.RequestHandler) FastBinder {
	_ = "STUB: not implemented"
	return *new(FastBinder)
}

// RoundTrip implements http.RoundTripper.RoundTrip.
func (binder FastBinder) RoundTrip(stdreq *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func std2fast(stdreq *http.Request) *fasthttp.Request { _ = "STUB: not implemented"; return nil }

func fast2std(stdreq *http.Request, fastresp *fasthttp.Response) *http.Response {
	_ = "STUB: not implemented"
	return nil
}

type fastLogger struct {
	logger Logger
}

func (f fastLogger) Printf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

type connNonTLS struct {
	net.Conn
}

func (connNonTLS) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (connNonTLS) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

type connTLS struct {
	connNonTLS
	state *tls.ConnectionState
}

func (c connTLS) Handshake() error { _ = "STUB: not implemented"; return nil }

func (c connTLS) ConnectionState() tls.ConnectionState {
	_ = "STUB: not implemented"
	return *new(tls.ConnectionState)
}
