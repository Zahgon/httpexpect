package httpexpect

import (
	"net/http"
)

// NewCookieJar returns a new http.CookieJar.
//
// Returned jar is implemented in net/http/cookiejar. PublicSuffixList is
// implemented in golang.org/x/net/publicsuffix.
//
// Note that this jar ignores cookies when request url is empty.
func NewCookieJar() http.CookieJar { _ = "STUB: not implemented"; return *new(http.CookieJar) }

// Deprecated: use NewCookieJar instead.
func NewJar() http.CookieJar { _ = "STUB: not implemented"; return *new(http.CookieJar) }
