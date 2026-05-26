package examples

import (
	"net/http"
)

// GaeHandler creates http.Handler to run in the Google App Engine.
//
// Routes:
//
//	GET /ping   return "pong"
func GaeHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }
