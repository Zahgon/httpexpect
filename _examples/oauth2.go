package examples

import (
	"net/http"

	"github.com/go-oauth2/oauth2/v4/server"
)

var (
	ClientID     = "aaa"
	ClientSecret = "bbb"
)

func OAuth2Handler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func validateToken(f http.HandlerFunc, srv *server.Server) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
