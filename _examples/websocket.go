package examples

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

// WsHttpHandler is a simple http.Handler that implements WebSocket echo server.
func WsHTTPHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// WsFastHandler is a simple fasthttp.RequestHandler that implements
// WebSocket echo server.
func WsFastHTTPHandler(ctx *fasthttp.RequestCtx) { _ = "STUB: not implemented"; return }
