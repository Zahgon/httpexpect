package examples

import (
	"net/http"
)

// IrisHandler tests iris handler
func IrisHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

// causes StreamWriter to stop writing

// return nil to continue writing

// if we had to write here then the StreamWriter callback should
// return nil instead of EOF
