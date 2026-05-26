package httpexpect

import (
	"bufio"
	"net"
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/valyala/fasthttp"
)

// NewWebsocketDialer produces new websocket.Dialer which dials to bound
// http.Handler without creating a real net.Conn.
func NewWebsocketDialer(handler http.Handler) *websocket.Dialer {
	_ = "STUB: not implemented"
	return nil
}

// NewFastWebsocketDialer produces new websocket.Dialer which dials to bound
// fasthttp.RequestHandler without creating a real net.Conn.
func NewFastWebsocketDialer(handler fasthttp.RequestHandler) *websocket.Dialer {
	_ = "STUB: not implemented"
	return nil
}

type handlerConn struct {
	net.Conn          // returned from dialer
	backConn net.Conn // passed to the background goroutine

	wg sync.WaitGroup
}

func newHandlerConn() *handlerConn { _ = "STUB: not implemented"; return nil }

func (hc *handlerConn) Close() error { _ = "STUB: not implemented"; return nil }

// wait the background goroutine

func (hc *handlerConn) runHandler(handler http.Handler) { _ = "STUB: not implemented"; return }

func (hc *handlerConn) runFastHandler(handler fasthttp.RequestHandler) {
	_ = "STUB: not implemented"
	return
}

// hijackRecorder it similar to httptest.ResponseRecorder,
// but with Hijack capabilities.
//
// Original idea is stolen from https://github.com/posener/wstest
type hijackRecorder struct {
	httptest.ResponseRecorder
	conn net.Conn
}

// Hijack the connection for caller.
//
// Implements http.Hijacker interface.
func (r *hijackRecorder) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

// WriteHeader write HTTP header to the client and closes the connection
//
// Implements http.ResponseWriter interface.
func (r *hijackRecorder) WriteHeader(code int) { _ = "STUB: not implemented"; return }
