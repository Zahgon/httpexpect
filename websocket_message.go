package httpexpect

// WebsocketMessage provides methods to inspect message read from WebSocket connection.
type WebsocketMessage struct {
	noCopy noCopy
	chain  *chain

	typ       int
	content   []byte
	closeCode int
}

// NewWebsocketMessage returns a new WebsocketMessage instance.
//
// If reporter is nil, the function panics.
// Content may be nil.
//
// Example:
//
//	m := NewWebsocketMessage(t, websocket.TextMessage, []byte("content"), 0)
//	m.TextMessage()
func NewWebsocketMessage(
	reporter Reporter, typ int, content []byte, closeCode ...int,
) *WebsocketMessage {
	_ = "STUB: not implemented"
	return nil
}

// NewWebsocketMessageC returns a new WebsocketMessage instance with config.
//
// Requirements for config are same as for WithConfig function.
// Content may be nil.
//
// Example:
//
//	m := NewWebsocketMessageC(config, websocket.TextMessage, []byte("content"), 0)
//	m.TextMessage()
func NewWebsocketMessageC(
	config Config, typ int, content []byte, closeCode ...int,
) *WebsocketMessage {
	_ = "STUB: not implemented"
	return nil
}

func newWebsocketMessage(
	parent *chain, typ int, content []byte, closeCode ...int,
) *WebsocketMessage {
	_ = "STUB: not implemented"
	return nil
}

func newEmptyWebsocketMessage(parent *chain) *WebsocketMessage {
	_ = "STUB: not implemented"
	return nil
}

// Raw returns underlying type, content and close code of WebSocket message.
// Theses values are originally read from WebSocket connection.
func (wm *WebsocketMessage) Raw() (typ int, content []byte, closeCode int) {
	_ = "STUB: not implemented"
	return 0, nil, 0
}

// Alias is similar to Value.Alias.
func (wm *WebsocketMessage) Alias(name string) *WebsocketMessage {
	_ = "STUB: not implemented"
	return nil
}

// CloseMessage is a shorthand for m.Type(websocket.CloseMessage).
func (wm *WebsocketMessage) CloseMessage() *WebsocketMessage { _ = "STUB: not implemented"; return nil }

// NotCloseMessage is a shorthand for m.NotType(websocket.CloseMessage).
func (wm *WebsocketMessage) NotCloseMessage() *WebsocketMessage {
	_ = "STUB: not implemented"
	return nil
}

// BinaryMessage is a shorthand for m.Type(websocket.BinaryMessage).
func (wm *WebsocketMessage) BinaryMessage() *WebsocketMessage {
	_ = "STUB: not implemented"
	return nil
}

// NotBinaryMessage is a shorthand for m.NotType(websocket.BinaryMessage).
func (wm *WebsocketMessage) NotBinaryMessage() *WebsocketMessage {
	_ = "STUB: not implemented"
	return nil
}

// TextMessage is a shorthand for m.Type(websocket.TextMessage).
func (wm *WebsocketMessage) TextMessage() *WebsocketMessage { _ = "STUB: not implemented"; return nil }

// NotTextMessage is a shorthand for m.NotType(websocket.TextMessage).
func (wm *WebsocketMessage) NotTextMessage() *WebsocketMessage {
	_ = "STUB: not implemented"
	return nil
}

// Type succeeds if WebSocket message type is one of the given.
//
// WebSocket message types are defined in RFC 6455, section 11.8.
// See also https://godoc.org/github.com/gorilla/websocket#pkg-constants
//
// Example:
//
//	msg := conn.Expect()
//	msg.Type(websocket.TextMessage, websocket.BinaryMessage)
func (wm *WebsocketMessage) Type(types ...int) *WebsocketMessage {
	_ = "STUB: not implemented"
	return nil
}

// NotType succeeds if WebSocket message type is none of the given.
//
// WebSocket message types are defined in RFC 6455, section 11.8.
// See also https://godoc.org/github.com/gorilla/websocket#pkg-constants
//
// Example:
//
//	msg := conn.Expect()
//	msg.NotType(websocket.CloseMessage, websocket.BinaryMessage)
func (wm *WebsocketMessage) NotType(types ...int) *WebsocketMessage {
	_ = "STUB: not implemented"
	return nil
}

func (wm *WebsocketMessage) checkType(opChain *chain, types ...int) {
	_ = "STUB: not implemented"
	return
}

func (wm *WebsocketMessage) checkNotType(opChain *chain, types ...int) {
	_ = "STUB: not implemented"
	return
}

// Code succeeds if WebSocket close code is one of the given.
//
// Code fails if WebSocket message type is not "8 - Connection Close Frame".
//
// WebSocket close codes are defined in RFC 6455, section 11.7.
// See also https://godoc.org/github.com/gorilla/websocket#pkg-constants
//
// Example:
//
//	msg := conn.Expect().Closed()
//	msg.Code(websocket.CloseNormalClosure, websocket.CloseGoingAway)
func (wm *WebsocketMessage) Code(codes ...int) *WebsocketMessage {
	_ = "STUB: not implemented"
	return nil
}

// NotCode succeeds if WebSocket close code is none of the given.
//
// NotCode fails if WebSocket message type is not "8 - Connection Close Frame".
//
// WebSocket close codes are defined in RFC 6455, section 11.7.
// See also https://godoc.org/github.com/gorilla/websocket#pkg-constants
//
// Example:
//
//	msg := conn.Expect().Closed()
//	msg.NotCode(websocket.CloseAbnormalClosure, websocket.CloseNoStatusReceived)
func (wm *WebsocketMessage) NotCode(codes ...int) *WebsocketMessage {
	_ = "STUB: not implemented"
	return nil
}

func (wm *WebsocketMessage) checkCode(opChain *chain, codes ...int) {
	_ = "STUB: not implemented"
	return
}

func (wm *WebsocketMessage) checkNotCode(opChain *chain, codes ...int) {
	_ = "STUB: not implemented"
	return
}

// NoContent succeeds if WebSocket message has no content (is empty).
func (wm *WebsocketMessage) NoContent() *WebsocketMessage { _ = "STUB: not implemented"; return nil }

// Body returns a new String instance with WebSocket message content.
//
// Example:
//
//	msg := conn.Expect()
//	msg.Body().NotEmpty()
//	msg.Body().Length().IsEqual(100)
func (wm *WebsocketMessage) Body() *String { _ = "STUB: not implemented"; return nil }

// JSON returns a new Value instance with JSON contents of WebSocket message.
//
// JSON succeeds if JSON may be decoded from message content.
//
// Example:
//
//	msg := conn.Expect()
//	msg.JSON().Array().ConsistsOf("foo", "bar")
func (wm *WebsocketMessage) JSON() *Value { _ = "STUB: not implemented"; return nil }

type wsMessageType int

func (wmt wsMessageType) String() string { _ = "STUB: not implemented"; return "" }

type wsCloseCode int

// https://developer.mozilla.org/en-US/docs/Web/API/CloseEvent/code
func (wcc wsCloseCode) String() string { _ = "STUB: not implemented"; return "" }
