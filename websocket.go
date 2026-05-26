package httpexpect

import (
	"time"

	"github.com/gorilla/websocket"
)

var (
	noDuration   = time.Duration(0)
	infiniteTime = time.Time{}
)

// WebsocketConn is used by Websocket to communicate with actual WebSocket connection.
type WebsocketConn interface {
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
	Close() error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
	Subprotocol() string
}

// Websocket provides methods to read from, write into and close WebSocket
// connection.
type Websocket struct {
	noCopy noCopy
	config Config
	chain  *chain

	conn WebsocketConn

	readTimeout  time.Duration
	writeTimeout time.Duration

	isClosed bool
}

// Deprecated: use NewWebsocketC instead.
func NewWebsocket(config Config, conn WebsocketConn) *Websocket {
	_ = "STUB: not implemented"
	return nil
}

// NewWebsocketC returns a new Websocket instance.
//
// Requirements for config are same as for WithConfig function.
func NewWebsocketC(config Config, conn WebsocketConn) *Websocket {
	_ = "STUB: not implemented"
	return nil
}

func newWebsocket(parent *chain, config Config, conn WebsocketConn) *Websocket {
	_ = "STUB: not implemented"
	return nil
}

// Conn returns underlying WebsocketConn object.
// This is the value originally passed to NewConnection.
func (ws *Websocket) Conn() WebsocketConn {
	_ = "STUB: not implemented"

	// Deprecated: use Conn instead.
	return *new(WebsocketConn)
}

func (ws *Websocket) Raw() *websocket.Conn { _ = "STUB: not implemented"; return nil }

// Alias is similar to Value.Alias.
func (ws *Websocket) Alias(name string) *Websocket { _ = "STUB: not implemented"; return nil }

// WithReadTimeout sets timeout duration for WebSocket connection reads.
//
// By default no timeout is used.
func (ws *Websocket) WithReadTimeout(timeout time.Duration) *Websocket {
	_ = "STUB: not implemented"
	return nil
}

// WithoutReadTimeout removes timeout for WebSocket connection reads.
func (ws *Websocket) WithoutReadTimeout() *Websocket { _ = "STUB: not implemented"; return nil }

// WithWriteTimeout sets timeout duration for WebSocket connection writes.
//
// By default no timeout is used.
func (ws *Websocket) WithWriteTimeout(timeout time.Duration) *Websocket {
	_ = "STUB: not implemented"
	return nil
}

// WithoutWriteTimeout removes timeout for WebSocket connection writes.
//
// If not used then DefaultWebsocketTimeout will be used.
func (ws *Websocket) WithoutWriteTimeout() *Websocket { _ = "STUB: not implemented"; return nil }

// Subprotocol returns a new String instance with negotiated protocol
// for the connection.
func (ws *Websocket) Subprotocol() *String { _ = "STUB: not implemented"; return nil }

// Expect reads next message from WebSocket connection and
// returns a new WebsocketMessage instance.
//
// Example:
//
//	msg := conn.Expect()
//	msg.JSON().Object().HasValue("message", "hi")
func (ws *Websocket) Expect() *WebsocketMessage { _ = "STUB: not implemented"; return nil }

// Disconnect closes the underlying WebSocket connection without sending or
// waiting for a close message.
//
// It's okay to call this function multiple times.
//
// It's recommended to always call this function after connection usage is over
// to ensure that no resource leaks will happen.
//
// Example:
//
//	conn := resp.Connection()
//	defer conn.Disconnect()
func (ws *Websocket) Disconnect() *Websocket { _ = "STUB: not implemented"; return nil }

// Close cleanly closes the underlying WebSocket connection
// by sending an empty close message and then waiting (with timeout)
// for the server to close the connection.
//
// WebSocket close code may be optionally specified.
// If not, then "1000 - Normal Closure" will be used.
//
// WebSocket close codes are defined in RFC 6455, section 11.7.
// See also https://godoc.org/github.com/gorilla/websocket#pkg-constants
//
// It's okay to call this function multiple times.
//
// Example:
//
//	conn := resp.Connection()
//	conn.Close(websocket.CloseUnsupportedData)
func (ws *Websocket) Close(code ...int) *Websocket { _ = "STUB: not implemented"; return nil }

// CloseWithBytes cleanly closes the underlying WebSocket connection
// by sending given slice of bytes as a close message and then waiting
// (with timeout) for the server to close the connection.
//
// WebSocket close code may be optionally specified.
// If not, then "1000 - Normal Closure" will be used.
//
// WebSocket close codes are defined in RFC 6455, section 11.7.
// See also https://godoc.org/github.com/gorilla/websocket#pkg-constants
//
// It's okay to call this function multiple times.
//
// Example:
//
//	conn := resp.Connection()
//	conn.CloseWithBytes([]byte("bye!"), websocket.CloseGoingAway)
func (ws *Websocket) CloseWithBytes(b []byte, code ...int) *Websocket {
	_ = "STUB: not implemented"
	return nil
}

// CloseWithJSON cleanly closes the underlying WebSocket connection
// by sending given object (marshaled using json.Marshal()) as a close message
// and then waiting (with timeout) for the server to close the connection.
//
// WebSocket close code may be optionally specified.
// If not, then "1000 - Normal Closure" will be used.
//
// WebSocket close codes are defined in RFC 6455, section 11.7.
// See also https://godoc.org/github.com/gorilla/websocket#pkg-constants
//
// It's okay to call this function multiple times.
//
// Example:
//
//	type MyJSON struct {
//	  Foo int `json:"foo"`
//	}
//
//	conn := resp.Connection()
//	conn.CloseWithJSON(MyJSON{Foo: 123}, websocket.CloseUnsupportedData)
func (ws *Websocket) CloseWithJSON(
	object interface{}, code ...int,
) *Websocket {
	_ = "STUB: not implemented"
	return nil
}

// CloseWithText cleanly closes the underlying WebSocket connection
// by sending given text as a close message and then waiting (with timeout)
// for the server to close the connection.
//
// WebSocket close code may be optionally specified.
// If not, then "1000 - Normal Closure" will be used.
//
// WebSocket close codes are defined in RFC 6455, section 11.7.
// See also https://godoc.org/github.com/gorilla/websocket#pkg-constants
//
// It's okay to call this function multiple times.
//
// Example:
//
//	conn := resp.Connection()
//	conn.CloseWithText("bye!")
func (ws *Websocket) CloseWithText(s string, code ...int) *Websocket {
	_ = "STUB: not implemented"
	return nil
}

// WriteMessage writes to the underlying WebSocket connection a message
// of given type with given content.
// Additionally, WebSocket close code may be specified for close messages.
//
// WebSocket message types are defined in RFC 6455, section 11.8.
// See also https://godoc.org/github.com/gorilla/websocket#pkg-constants
//
// WebSocket close codes are defined in RFC 6455, section 11.7.
// See also https://godoc.org/github.com/gorilla/websocket#pkg-constants
//
// Example:
//
//	conn := resp.Connection()
//	conn.WriteMessage(websocket.CloseMessage, []byte("Namárië..."))
func (ws *Websocket) WriteMessage(typ int, content []byte, closeCode ...int) *Websocket {
	_ = "STUB: not implemented"
	return nil
}

// WriteBytesBinary is a shorthand for c.WriteMessage(websocket.BinaryMessage, b).
func (ws *Websocket) WriteBytesBinary(b []byte) *Websocket { _ = "STUB: not implemented"; return nil }

// WriteBytesText is a shorthand for c.WriteMessage(websocket.TextMessage, b).
func (ws *Websocket) WriteBytesText(b []byte) *Websocket { _ = "STUB: not implemented"; return nil }

// WriteText is a shorthand for
// c.WriteMessage(websocket.TextMessage, []byte(s)).
func (ws *Websocket) WriteText(s string) *Websocket { _ = "STUB: not implemented"; return nil }

// WriteJSON writes to the underlying WebSocket connection given object,
// marshaled using json.Marshal().
func (ws *Websocket) WriteJSON(object interface{}) *Websocket {
	_ = "STUB: not implemented"
	return nil
}

func (ws *Websocket) checkUnusable(opChain *chain, where string) bool {
	_ = "STUB: not implemented"
	return false
}

func (ws *Websocket) readMessage(opChain *chain) *WebsocketMessage {
	_ = "STUB: not implemented"
	return nil
}

func (ws *Websocket) writeMessage(
	opChain *chain, typ int, content []byte, closeCode ...int,
) {
	_ = "STUB: not implemented"
	return
}

func (ws *Websocket) setReadDeadline(opChain *chain) bool { _ = "STUB: not implemented"; return false }

func (ws *Websocket) setWriteDeadline(opChain *chain) bool { _ = "STUB: not implemented"; return false }

func (ws *Websocket) printRead(typ int, content []byte, closeCode int) {
	_ = "STUB: not implemented"
	return
}

func (ws *Websocket) printWrite(typ int, content []byte, closeCode int) {
	_ = "STUB: not implemented"
	return
}
