package httpexpect

// String provides methods to inspect attached string value
// (Go representation of JSON string).
type String struct {
	noCopy noCopy
	chain  *chain
	value  string
}

// NewString returns a new String instance.
//
// If reporter is nil, the function panics.
//
// Example:
//
//	str := NewString(t, "Hello")
func NewString(reporter Reporter, value string) *String { _ = "STUB: not implemented"; return nil }

// NewStringC returns a new String instance with config.
//
// Requirements for config are same as for WithConfig function.
//
// Example:
//
//	str := NewStringC(config, "Hello")
func NewStringC(config Config, value string) *String { _ = "STUB: not implemented"; return nil }

func newString(parent *chain, val string) *String { _ = "STUB: not implemented"; return nil }

// Raw returns underlying value attached to String.
// This is the value originally passed to NewString.
//
// Example:
//
//	str := NewString(t, "Hello")
//	assert.Equal(t, "Hello", str.Raw())
func (s *String) Raw() string {
	_ = "STUB: not implemented"

	// Decode unmarshals the underlying value attached to the String to a target variable.
	// target should be one of these:
	//
	//   - pointer to an empty interface
	//   - pointer to a string
	//
	// Example:
	//
	//	value := NewString(t, "foo")
	//
	//	var target string
	//	value.Decode(&target)
	//
	//	assert.Equal(t, "foo", target)
	return ""
}

func (s *String) Decode(target interface{}) *String { _ = "STUB: not implemented"; return nil }

// Alias is similar to Value.Alias.
func (s *String) Alias(name string) *String { _ = "STUB: not implemented"; return nil }

// Path is similar to Value.Path.
func (s *String) Path(path string) *Value { _ = "STUB: not implemented"; return nil }

// Schema is similar to Value.Schema.
func (s *String) Schema(schema interface{}) *String { _ = "STUB: not implemented"; return nil }

// Length returns a new Number instance with string length.
//
// Example:
//
//	str := NewString(t, "Hello")
//	str.Length().IsEqual(5)
func (s *String) Length() *Number { _ = "STUB: not implemented"; return nil }

// IsEmpty succeeds if string is empty.
//
// Example:
//
//	str := NewString(t, "")
//	str.IsEmpty()
func (s *String) IsEmpty() *String { _ = "STUB: not implemented"; return nil }

// NotEmpty succeeds if string is non-empty.
//
// Example:
//
//	str := NewString(t, "Hello")
//	str.NotEmpty()
func (s *String) NotEmpty() *String { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsEmpty instead.
func (s *String) Empty() *String {
	_ = "STUB: not implemented"

	// IsEqual succeeds if string is equal to given Go string.
	//
	// Example:
	//
	//	str := NewString(t, "Hello")
	//	str.IsEqual("Hello")
	return nil
}

func (s *String) IsEqual(value string) *String { _ = "STUB: not implemented"; return nil }

// NotEqual succeeds if string is not equal to given Go string.
//
// Example:
//
//	str := NewString(t, "Hello")
//	str.NotEqual("Goodbye")
func (s *String) NotEqual(value string) *String { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsEqual instead.
func (s *String) Equal(value string) *String { _ = "STUB: not implemented"; return nil }

// IsEqualFold succeeds if string is equal to given Go string after applying Unicode
// case-folding (so it's a case-insensitive match).
//
// Example:
//
//	str := NewString(t, "Hello")
//	str.IsEqualFold("hELLo")
func (s *String) IsEqualFold(value string) *String { _ = "STUB: not implemented"; return nil }

// NotEqualFold succeeds if string is not equal to given Go string after applying
// Unicode case-folding (so it's a case-insensitive match).
//
// Example:
//
//	str := NewString(t, "Hello")
//	str.NotEqualFold("gOODBYe")
func (s *String) NotEqualFold(value string) *String { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsEqualFold instead.
func (s *String) EqualFold(value string) *String { _ = "STUB: not implemented"; return nil }

// InList succeeds if the string is equal to one of the values from given
// list of strings.
//
// Example:
//
//	str := NewString(t, "Hello")
//	str.InList("Hello", "Goodbye")
func (s *String) InList(values ...string) *String { _ = "STUB: not implemented"; return nil }

// NotInList succeeds if the string is not equal to any of the values from
// given list of strings.
//
// Example:
//
//	str := NewString(t, "Hello")
//	str.NotInList("Sayonara", "Goodbye")
func (s *String) NotInList(values ...string) *String { _ = "STUB: not implemented"; return nil }

// InListFold succeeds if the string is equal to one of the values from given
// list of strings after applying Unicode case-folding (so it's a case-insensitive match).
//
// Example:
//
//	str := NewString(t, "Hello")
//	str.InListFold("hEllo", "Goodbye")
func (s *String) InListFold(values ...string) *String { _ = "STUB: not implemented"; return nil }

// NotInListFold succeeds if the string is not equal to any of the values from given
// list of strings after applying Unicode case-folding (so it's a case-insensitive match).
//
// Example:
//
//	str := NewString(t, "Hello")
//	str.NotInListFold("Bye", "Goodbye")
func (s *String) NotInListFold(values ...string) *String { _ = "STUB: not implemented"; return nil }

// Contains succeeds if string contains given Go string as a substring.
//
// Example:
//
//	str := NewString(t, "Hello")
//	str.Contains("ell")
func (s *String) Contains(value string) *String { _ = "STUB: not implemented"; return nil }

// NotContains succeeds if string doesn't contain Go string as a substring.
//
// Example:
//
//	str := NewString(t, "Hello")
//	str.NotContains("bye")
func (s *String) NotContains(value string) *String { _ = "STUB: not implemented"; return nil }

// ContainsFold succeeds if string contains given Go string as a substring after
// applying Unicode case-folding (so it's a case-insensitive match).
//
// Example:
//
//	str := NewString(t, "Hello")
//	str.ContainsFold("ELL")
func (s *String) ContainsFold(value string) *String { _ = "STUB: not implemented"; return nil }

// NotContainsFold succeeds if string doesn't contain given Go string as a substring
// after applying Unicode case-folding (so it's a case-insensitive match).
//
// Example:
//
//	str := NewString(t, "Hello")
//	str.NotContainsFold("BYE")
func (s *String) NotContainsFold(value string) *String { _ = "STUB: not implemented"; return nil }

// HasPrefix succeeds if string has given Go string as prefix
//
// Example:
//
//	str := NewString(t, "Hello World")
//	str.HasPrefix("Hello")
func (s *String) HasPrefix(value string) *String { _ = "STUB: not implemented"; return nil }

// NotHasPrefix succeeds if string doesn't have given Go string as prefix
//
// Example:
//
//	str := NewString(t, "Hello World")
//	str.NotHasPrefix("Bye")
func (s *String) NotHasPrefix(value string) *String { _ = "STUB: not implemented"; return nil }

// HasSuffix succeeds if string has given Go string as suffix
//
// Example:
//
//	str := NewString(t, "Hello World")
//	str.HasSuffix("World")
func (s *String) HasSuffix(value string) *String { _ = "STUB: not implemented"; return nil }

// NotHasSuffix succeeds if string doesn't have given Go string as suffix
//
// Example:
//
//	str := NewString(t, "Hello World")
//	str.NotHasSuffix("Hello")
func (s *String) NotHasSuffix(value string) *String { _ = "STUB: not implemented"; return nil }

// HasPrefixFold succeeds if string has given Go string as prefix
// after applying Unicode case-folding (so it's a case-insensitive match).
//
// Example:
//
//	str := NewString(t, "Hello World")
//	str.HasPrefixFold("hello")
func (s *String) HasPrefixFold(value string) *String { _ = "STUB: not implemented"; return nil }

// NotHasPrefixFold succeeds if string doesn't have given Go string as prefix
// after applying Unicode case-folding (so it's a case-insensitive match).
//
// Example:
//
//	str := NewString(t, "Hello World")
//	str.NotHasPrefixFold("Bye")
func (s *String) NotHasPrefixFold(value string) *String { _ = "STUB: not implemented"; return nil }

// HasSuffixFold succeeds if string has given Go string as suffix
// after applying Unicode case-folding (so it's a case-insensitive match).
//
// Example:
//
//	str := NewString(t, "Hello World")
//	str.HasSuffixFold("world")
func (s *String) HasSuffixFold(value string) *String { _ = "STUB: not implemented"; return nil }

// NotHasSuffixFold succeeds if string doesn't have given Go string as suffix
// after applying Unicode case-folding (so it's a case-insensitive match).
//
// Example:
//
//	str := NewString(t, "Hello World")
//	str.NotHasSuffixFold("Bye")
func (s *String) NotHasSuffixFold(value string) *String { _ = "STUB: not implemented"; return nil }

// Match matches the string with given regexp and returns a new Match instance
// with found submatches.
//
// If regexp is invalid or string doesn't match regexp, Match fails and returns
// empty (but non-nil) instance. regexp.Compile is used to construct regexp, and
// Regexp.FindStringSubmatch is used to construct matches.
//
// Example:
//
//	s := NewString(t, "http://example.com/users/john")
//	m := s.Match(`http://(?P<host>.+)/users/(?P<user>.+)`)
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
func (s *String) Match(re string) *Match { _ = "STUB: not implemented"; return nil }

// NotMatch succeeds if the string doesn't match to given regexp.
//
// regexp.Compile is used to construct regexp, and Regexp.MatchString
// is used to perform match.
//
// Example:
//
//	s := NewString(t, "a")
//	s.NotMatch(`[^a]`)
func (s *String) NotMatch(re string) *String { _ = "STUB: not implemented"; return nil }

// MatchAll find all matches in string for given regexp and returns a list
// of found matches.
//
// If regexp is invalid or string doesn't match regexp, MatchAll fails and
// returns empty (but non-nil) slice. regexp.Compile is used to construct
// regexp, and Regexp.FindAllStringSubmatch is used to find matches.
//
// Example:
//
//	s := NewString(t,
//	   "http://example.com/users/john http://example.com/users/bob")
//
//	m := s.MatchAll(`http://(?P<host>\S+)/users/(?P<user>\S+)`)
//
//	m[0].NamedSubmatch("user").IsEqual("john")
//	m[1].NamedSubmatch("user").IsEqual("bob")
func (s *String) MatchAll(re string) []Match { _ = "STUB: not implemented"; return nil }

// IsASCII succeeds if all string characters belongs to ASCII.
//
// Example:
//
//	str := NewString(t, "Hello")
//	str.IsASCII()
func (s *String) IsASCII() *String { _ = "STUB: not implemented"; return nil }

// NotASCII succeeds if at least one string character does not belong to ASCII.
//
// Example:
//
//	str := NewString(t, "こんにちは")
//	str.NotASCII()
func (s *String) NotASCII() *String { _ = "STUB: not implemented"; return nil }

// Deprecated: use NotASCII instead.
func (s *String) NotIsASCII() *String { _ = "STUB: not implemented"; return nil }

// AsNumber parses float from string and returns a new Number instance
// with result.
//
// If base is 10 or omitted, uses strconv.ParseFloat.
// Otherwise, uses strconv.ParseInt or strconv.ParseUint with given base.
//
// Example:
//
//	str := NewString(t, "100")
//	str.AsNumber().IsEqual(100)
//
// Specifying base:
//
//	str.AsNumber(10).IsEqual(100)
//	str.AsNumber(16).IsEqual(256)
func (s *String) AsNumber(base ...int) *Number { _ = "STUB: not implemented"; return nil }

// AsBoolean parses true/false value string and returns a new Boolean instance
// with result.
//
// Accepts string values "true", "True", "false", "False".
//
// Example:
//
//	str := NewString(t, "true")
//	str.AsBoolean().IsTrue()
func (s *String) AsBoolean() *Boolean { _ = "STUB: not implemented"; return nil }

// AsDateTime parses date/time from string and returns a new DateTime instance
// with result.
//
// If format is given, AsDateTime() uses time.Parse() with every given format.
// Otherwise, it uses the list of predefined common formats.
//
// If the string can't be parsed with any format, AsDateTime reports failure
// and returns empty (but non-nil) instance.
//
// Example:
//
//	str := NewString(t, "Tue, 15 Nov 1994 08:12:31 GMT")
//	str.AsDateTime().IsLt(time.Now())
//
//	str := NewString(t, "15 Nov 94 08:12 GMT")
//	str.AsDateTime(time.RFC822).IsLt(time.Now())
func (s *String) AsDateTime(format ...string) *DateTime { _ = "STUB: not implemented"; return nil }

type datetimeFormat struct {
	layout string
	name   string
}

func (f datetimeFormat) String() string { _ = "STUB: not implemented"; return "" }

// Deprecated: use AsNumber instead.
func (s *String) Number() *Number { _ = "STUB: not implemented"; return nil }

// Deprecated: use AsDateTime instead.
func (s *String) DateTime(layout ...string) *DateTime { _ = "STUB: not implemented"; return nil }
