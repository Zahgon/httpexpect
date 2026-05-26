package examples

import (
	"github.com/valyala/fasthttp"
)

// FastHTTPHandler creates fasthttp.RequestHandler.
//
// Routes:
//
//	GET /ping   return "pong"
func FastHTTPHandler() fasthttp.RequestHandler {
	_ = "STUB: not implemented"
	return *new(fasthttp.RequestHandler)
}
