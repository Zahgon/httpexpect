package httpexpect

// Match provides methods to inspect attached regexp match results.
type Match struct {
	chain          *chain
	submatchValues []string
	submatchNames  map[string]int
}

// NewMatch returns a new Match instance.
//
// If reporter is nil, the function panics.
// Both submatchValues and submatchNames may be nil.
//
// Example:
//
//	s := "http://example.com/users/john"
//	r := regexp.MustCompile(`http://(?P<host>.+)/users/(?P<user>.+)`)
//
//	m := NewMatch(t, r.FindStringSubmatch(s), r.SubexpNames())
//
//	m.NotEmpty()
//	m.Length().IsEqual(3)
//
//	m.Submatch(0).IsEqual("http://example.com/users/john")
//	m.Submatch(1).IsEqual("example.com")
//	m.Submatch(2).IsEqual("john")
//
//	m.NamedSubmatch("host").IsEqual("example.com")
//	m.NamedSubmatch("user").IsEqual("john")
func NewMatch(reporter Reporter, submatchValues []string, submatchNames []string) *Match {
	_ = "STUB: not implemented"
	return nil
}

// NewMatchC returns a new Match instance with config.
//
// Requirements for config are same as for WithConfig function.
// Both submatches and names may be nil.
//
// See NewMatch for usage example.
func NewMatchC(config Config, submatchValues []string, submatchNames []string) *Match {
	_ = "STUB: not implemented"
	return nil
}

func newMatch(parent *chain, submatchValues []string, submatchNames []string) *Match {
	_ = "STUB: not implemented"
	return nil
}

// Raw returns underlying submatches attached to Match.
// This is the value originally passed to NewMatch.
//
// Example:
//
//	m := NewMatch(t, submatches, names)
//	assert.Equal(t, submatches, m.Raw())
func (m *Match) Raw() []string { _ = "STUB: not implemented"; return nil }

// Alias is similar to Value.Alias.
func (m *Match) Alias(name string) *Match { _ = "STUB: not implemented"; return nil }

// Length returns a new Number instance with number of submatches.
//
// Example:
//
//	m := NewMatch(t, submatches, names)
//	m.Length().IsEqual(len(submatches))
func (m *Match) Length() *Number { _ = "STUB: not implemented"; return nil }

// Submatch returns a new String instance with submatch for given index.
//
// Note that submatch with index 0 contains the whole match. If index is out
// of bounds, Submatch reports failure and returns empty (but non-nil) instance.
//
// Example:
//
//	s := "http://example.com/users/john"
//
//	r := regexp.MustCompile(`http://(.+)/users/(.+)`)
//	m := NewMatch(t, r.FindStringSubmatch(s), nil)
//
//	m.Submatch(0).IsEqual("http://example.com/users/john")
//	m.Submatch(1).IsEqual("example.com")
//	m.Submatch(2).IsEqual("john")
func (m *Match) Submatch(index int) *String { _ = "STUB: not implemented"; return nil }

// NamedSubmatch returns a new String instance with submatch for given name.
//
// If there is no submatch with given name, NamedSubmatch reports failure and returns
// empty (but non-nil) instance.
//
// Example:
//
//	s := "http://example.com/users/john"
//
//	r := regexp.MustCompile(`http://(?P<host>.+)/users/(?P<user>.+)`)
//	m := NewMatch(t, r.FindStringSubmatch(s), r.SubexpNames())
//
//	m.NamedSubmatch("host").IsEqual("example.com")
//	m.NamedSubmatch("user").IsEqual("john")
func (m *Match) NamedSubmatch(name string) *String { _ = "STUB: not implemented"; return nil }

// Deprecated: use Submatch instead.
func (m *Match) Index(index int) *String { _ = "STUB: not implemented"; return nil }

// Deprecated: use NamedSubmatch instead.
func (m *Match) Name(name string) *String { _ = "STUB: not implemented"; return nil }

// IsEmpty succeeds if submatches array is empty.
//
// Example:
//
//	m := NewMatch(t, submatches, names)
//	m.IsEmpty()
func (m *Match) IsEmpty() *Match { _ = "STUB: not implemented"; return nil }

// NotEmpty succeeds if submatches array is non-empty.
//
// Example:
//
//	m := NewMatch(t, submatches, names)
//	m.NotEmpty()
func (m *Match) NotEmpty() *Match { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsEmpty instead.
func (m *Match) Empty() *Match {
	_ = "STUB: not implemented"

	// HasSubmatches succeeds if submatches array, starting from index 1, is equal to
	// given array.
	//
	// Note that submatch with index 0 contains the whole match and is not
	// included into this check.
	//
	// Example:
	//
	//	s := "http://example.com/users/john"
	//	r := regexp.MustCompile(`http://(.+)/users/(.+)`)
	//	m := NewMatch(t, r.FindStringSubmatch(s), nil)
	//	m.HasSubmatches("example.com", "john")
	return nil
}

func (m *Match) HasSubmatches(submatchValues ...string) *Match {
	_ = "STUB: not implemented"
	return nil
}

// NotHasSubmatches succeeds if submatches array, starting from index 1, is not
// equal to given array.
//
// Note that submatch with index 0 contains the whole match and is not
// included into this check.
//
// Example:
//
//	s := "http://example.com/users/john"
//	r := regexp.MustCompile(`http://(.+)/users/(.+)`)
//	m := NewMatch(t, r.FindStringSubmatch(s), nil)
//	m.NotHasSubmatches("example.com", "bob")
func (m *Match) NotHasSubmatches(submatchValues ...string) *Match {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: use HasSubmatches instead.
func (m *Match) Values(submatchValues ...string) *Match { _ = "STUB: not implemented"; return nil }

// Deprecated: use NotHasSubmatches instead.
func (m *Match) NotValues(submatchValues ...string) *Match { _ = "STUB: not implemented"; return nil }

func (m *Match) getValues() []string { _ = "STUB: not implemented"; return nil }
