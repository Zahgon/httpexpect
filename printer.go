package httpexpect

import (
	"net/http"
	"time"
)

// Printer is used to print requests and responses.
// CompactPrinter, DebugPrinter, and CurlPrinter implement this interface.
type Printer interface {
	// Request is called before request is sent.
	// It is allowed to read and close request body, or ignore it.
	Request(*http.Request)

	// Response is called after response is received.
	// It is allowed to read and close response body, or ignore it.
	Response(*http.Response, time.Duration)
}

// WebsocketPrinter is used to print writes and reads of WebSocket connection.
//
// If WebSocket connection is used, all Printers that also implement WebsocketPrinter
// are invoked on every WebSocket message read or written.
//
// DebugPrinter implements this interface.
type WebsocketPrinter interface {
	Printer

	// WebsocketWrite is called before writes to WebSocket connection.
	WebsocketWrite(typ int, content []byte, closeCode int)

	// WebsocketRead is called after reads from WebSocket connection.
	WebsocketRead(typ int, content []byte, closeCode int)
}

// CompactPrinter implements Printer.
// Prints requests in compact form. Does not print responses.
type CompactPrinter struct {
	logger Logger
}

// NewCompactPrinter returns a new CompactPrinter given a logger.
func NewCompactPrinter(logger Logger) CompactPrinter {
	_ = "STUB: not implemented"
	return *new(CompactPrinter)
}

// Request implements Printer.Request.
func (p CompactPrinter) Request(req *http.Request) { _ = "STUB: not implemented"; return }

// Response implements Printer.Response.
func (CompactPrinter) Response(*http.Response, time.Duration) {
	_ = "STUB: not implemented"

	// CurlPrinter implements Printer.
	// Uses http2curl to dump requests as curl commands that can be inserted
	// into terminal.
	return
}

type CurlPrinter struct {
	logger Logger
}

// NewCurlPrinter returns a new CurlPrinter given a logger.
func NewCurlPrinter(logger Logger) CurlPrinter { _ = "STUB: not implemented"; return *new(CurlPrinter) }

// Request implements Printer.Request.
func (p CurlPrinter) Request(req *http.Request) { _ = "STUB: not implemented"; return }

// Response implements Printer.Response.
func (CurlPrinter) Response(*http.Response, time.Duration) {
	_ = "STUB: not implemented"

	// DebugPrinter implements Printer and WebsocketPrinter.
	// Uses net/http/httputil to dump both requests and responses.
	// Also prints all websocket messages.
	return
}

type DebugPrinter struct {
	logger Logger
	body   bool
}

// NewDebugPrinter returns a new DebugPrinter given a logger and body
// flag. If body is true, request and response body is also printed.
func NewDebugPrinter(logger Logger, body bool) DebugPrinter {
	_ = "STUB: not implemented"
	return *new(DebugPrinter)
}

// Request implements Printer.Request.
func (p DebugPrinter) Request(req *http.Request) { _ = "STUB: not implemented"; return }

// Response implements Printer.Response.
func (p DebugPrinter) Response(resp *http.Response, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

// WebsocketWrite implements WebsocketPrinter.WebsocketWrite.
func (p DebugPrinter) WebsocketWrite(typ int, content []byte, closeCode int) {
	_ = "STUB: not implemented"
	return
}

// WebsocketRead implements WebsocketPrinter.WebsocketRead.
func (p DebugPrinter) WebsocketRead(typ int, content []byte, closeCode int) {
	_ = "STUB: not implemented"
	return
}
