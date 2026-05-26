package httpexpect

import (
	"net/http"
)

// Cookie provides methods to inspect attached http.Cookie value.
type Cookie struct {
	noCopy noCopy
	chain  *chain
	value  *http.Cookie
}

// NewCookie returns a new Cookie instance.
//
// If reporter is nil, the function panics.
// If value is nil, failure is reported.
//
// Example:
//
//	cookie := NewCookie(t, &http.Cookie{...})
//
//	cookie.Domain().IsEqual("example.com")
//	cookie.Path().IsEqual("/")
//	cookie.Expires().InRange(time.Now(), time.Now().Add(time.Hour * 24))
func NewCookie(reporter Reporter, value *http.Cookie) *Cookie {
	_ = "STUB: not implemented"
	return nil
}

// NewCookieC returns a new Cookie instance with config.
//
// Requirements for config are same as for WithConfig function.
// If value is nil, failure is reported.
//
// See NewCookie for usage example.
func NewCookieC(config Config, value *http.Cookie) *Cookie { _ = "STUB: not implemented"; return nil }

func newCookie(parent *chain, val *http.Cookie) *Cookie { _ = "STUB: not implemented"; return nil }

// Raw returns underlying http.Cookie value attached to Cookie.
// This is the value originally passed to NewCookie.
//
// Example:
//
//	cookie := NewCookie(t, c)
//	assert.Equal(t, c, cookie.Raw())
func (c *Cookie) Raw() *http.Cookie {
	_ = "STUB: not implemented"

	// Alias is similar to Value.Alias.
	return nil
}

func (c *Cookie) Alias(name string) *Cookie { _ = "STUB: not implemented"; return nil }

// Name returns a new String instance with cookie name.
//
// Example:
//
//	cookie := NewCookie(t, &http.Cookie{...})
//	cookie.Name().IsEqual("session")
func (c *Cookie) Name() *String { _ = "STUB: not implemented"; return nil }

// Value returns a new String instance with cookie value.
//
// Example:
//
//	cookie := NewCookie(t, &http.Cookie{...})
//	cookie.Value().IsEqual("gH6z7Y")
func (c *Cookie) Value() *String { _ = "STUB: not implemented"; return nil }

// Domain returns a new String instance with cookie domain.
//
// Example:
//
//	cookie := NewCookie(t, &http.Cookie{...})
//	cookie.Domain().IsEqual("example.com")
func (c *Cookie) Domain() *String { _ = "STUB: not implemented"; return nil }

// Path returns a new String instance with cookie path.
//
// Example:
//
//	cookie := NewCookie(t, &http.Cookie{...})
//	cookie.Path().IsEqual("/foo")
func (c *Cookie) Path() *String { _ = "STUB: not implemented"; return nil }

// Expires returns a new DateTime instance with cookie expiration date.
//
// Example:
//
//	cookie := NewCookie(t, &http.Cookie{...})
//	cookie.Expires().InRange(time.Now(), time.Now().Add(time.Hour * 24))
func (c *Cookie) Expires() *DateTime { _ = "STUB: not implemented"; return nil }

// ContainsMaxAge succeeds if cookie has Max-Age field.
//
// In particular, if Max-Age is present and is zero (which means delete
// cookie now), method succeeds.
//
// Example:
//
//	cookie := NewCookie(t, &http.Cookie{...})
//	cookie.ContainsMaxAge()
func (c *Cookie) ContainsMaxAge() *Cookie { _ = "STUB: not implemented"; return nil }

// NotContainsMaxAge succeeds if cookie does not have Max-Age field.
//
// In particular, if Max-Age is present and is zero (which means delete
// cookie now), method fails.
//
// Example:
//
//	cookie := NewCookie(t, &http.Cookie{...})
//	cookie.NotContainsMaxAge()
func (c *Cookie) NotContainsMaxAge() *Cookie { _ = "STUB: not implemented"; return nil }

// Deprecated: use ContainsMaxAge instead.
func (c *Cookie) HasMaxAge() *Cookie { _ = "STUB: not implemented"; return nil }

// Deprecated: use NotContainsMaxAge instead.
func (c *Cookie) NotHasMaxAge() *Cookie { _ = "STUB: not implemented"; return nil }

// Deprecated: use ContainsMaxAge instead.
func (c *Cookie) HaveMaxAge() *Cookie { _ = "STUB: not implemented"; return nil }

// Deprecated: use NotContainsMaxAge instead.
func (c *Cookie) NotHaveMaxAge() *Cookie { _ = "STUB: not implemented"; return nil }

// MaxAge returns a new Duration instance with cookie Max-Age field.
//
// If Max-Age is not present, method fails.
//
// If Max-Age is present and is zero (which means delete cookie now),
// methods succeeds and the returned Duration is equal to zero.
//
// Example:
//
//	cookie := NewCookie(t, &http.Cookie{...})
//	cookie.ContainsMaxAge()
//	cookie.MaxAge().InRange(time.Minute, time.Minute*10)
func (c *Cookie) MaxAge() *Duration { _ = "STUB: not implemented"; return nil }

// zero value means not present
// TODO: after removing Duration.IsSet, add failure here (breaking change)

// negative value means present and zero
