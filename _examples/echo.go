package examples

import (
	"net/http"
)

// EchoHandler creates http.Handler using echo framework.
//
// Routes:
//
//	GET /login             authenticate user and return JWT token
//	GET /restricted/hello  return "hello, world!" (requires authentication)
func EchoHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

// create token

// generate encoded token and send it as response
