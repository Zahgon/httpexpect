package httpexpect

// Value provides methods to inspect attached interface{} object
// (Go representation of arbitrary JSON value) and cast it to
// concrete type.
type Value struct {
	chain *chain
	value interface{}
}

// NewValue returns a new Value instance.
//
// If reporter is nil, the function panics.
// Value may be nil.
//
// Example:
//
//	value := NewValue(t, map[string]interface{}{"foo": 123})
//	value.IsObject()
//
//	value := NewValue(t, []interface{}{"foo", 123})
//	value.IsArray()
//
//	value := NewValue(t, "foo")
//	value.IsString()
//
//	value := NewValue(t, 123)
//	value.IsNumber()
//
//	value := NewValue(t, true)
//	value.IsBoolean()
//
//	value := NewValue(t, nil)
//	value.IsNull()
func NewValue(reporter Reporter, value interface{}) *Value { _ = "STUB: not implemented"; return nil }

// NewValueC returns a new Value instance with config.
//
// Requirements for config are same as for WithConfig function.
// Value may be nil.
//
// See NewValue for usage example.
func NewValueC(config Config, value interface{}) *Value { _ = "STUB: not implemented"; return nil }

func newValue(parent *chain, val interface{}) *Value { _ = "STUB: not implemented"; return nil }

// Raw returns underlying value attached to Value.
// This is the value originally passed to NewValue, converted to canonical form.
//
// Example:
//
//	value := NewValue(t, "foo")
//	assert.Equal(t, "foo", number.Raw().(string))
func (v *Value) Raw() interface{} {
	_ = "STUB: not implemented"

	// Decode unmarshals the underlying value attached to the Object to a target variable
	// target should be pointer to any type.
	//
	// Example:
	//
	//	type S struct {
	//		Foo int             `json:"foo"`
	//		Bar []interface{}   `json:"bar"`
	//		Baz struct{ A int } `json:"baz"`
	//	}
	//
	//	m := map[string]interface{}{
	//		"foo": 123,
	//		"bar": []interface{}{"123", 456.0},
	//		"baz": struct{ A int }{123},
	//	}
	//
	//	value = NewValue(reporter,m)
	//
	//	var target S
	//	value.Decode(&target)
	//
	//	assert.Equal(t, S{123, []interface{}{"123", 456.0}, struct{ A int }{123}, target})
	return nil
}

func (v *Value) Decode(target interface{}) *Value { _ = "STUB: not implemented"; return nil }

// Alias returns a new Value object with alias.
// When a test of Value object with alias is failed,
// an assertion is displayed as a chain starting from the alias.
//
// Example:
//
//	// In this example, GET /example responds "foo"
//	foo := e.GET("/example").Expect().Status(http.StatusOK).JSON().Object()
//
//	// When a test is failed, an assertion without alias is
//	// Request("GET").Expect().JSON().Object().IsEqual()
//	foo.IsEqual("bar")
//
//	// Set Alias
//	fooWithAlias := e.GET("/example").
//		Expect().
//		Status(http.StatusOK).JSON().Object().Alias("foo")
//
//	// When a test is failed, an assertion with alias is
//	// foo.IsEqual()
//	fooWithAlias.IsEqual("bar")
func (v *Value) Alias(name string) *Value { _ = "STUB: not implemented"; return nil }

// Path returns a new Value object for child object(s) matching given
// JSONPath expression.
//
// JSONPath is a simple XPath-like query language.
// See http://goessner.net/articles/JsonPath/.
//
// We currently use https://github.com/yalp/jsonpath, which implements
// only a subset of JSONPath, yet useful for simple queries. It doesn't
// support filters and requires double quotes for strings.
//
// Example 1:
//
//	json := `{"users": [{"name": "john"}, {"name": "bob"}]}`
//	value := NewValue(t, json)
//
//	value.Path("$.users[0].name").String().IsEqual("john")
//	value.Path("$.users[1].name").String().IsEqual("bob")
//
// Example 2:
//
//	json := `{"yfGH2a": {"user": "john"}, "f7GsDd": {"user": "john"}}`
//	value := NewValue(t, json)
//
//	for _, user := range value.Path("$..user").Array().Iter() {
//		user.String().IsEqual("john")
//	}
func (v *Value) Path(path string) *Value { _ = "STUB: not implemented"; return nil }

// Schema succeeds if value matches given JSON Schema.
//
// JSON Schema specifies a JSON-based format to define the structure of
// JSON data. See http://json-schema.org/.
// We use https://github.com/xeipuuv/gojsonschema implementation.
//
// schema should be one of the following:
//   - go value that can be json.Marshal-ed to a valid schema
//   - type convertible to string containing valid schema
//   - type convertible to string containing valid http:// or file:// URI,
//     pointing to reachable and valid schema
//
// Example 1:
//
//	 schema := `{
//	   "type": "object",
//	   "properties": {
//		  "foo": {
//			  "type": "string"
//		  },
//		  "bar": {
//			  "type": "integer"
//		  }
//	  },
//	  "require": ["foo", "bar"]
//	}`
//
//	value := NewValue(t, map[string]interface{}{
//		"foo": "a",
//		"bar": 1,
//	})
//
//	value.Schema(schema)
//
// Example 2:
//
//	value := NewValue(t, data)
//	value.Schema("http://example.com/schema.json")
func (v *Value) Schema(schema interface{}) *Value { _ = "STUB: not implemented"; return nil }

// Object returns a new Object attached to underlying value.
//
// If underlying value is not an object (map[string]interface{}), failure is reported
// and empty (but non-nil) value is returned.
//
// Example:
//
//	value := NewValue(t, map[string]interface{}{"foo": 123})
//	value.Object().ContainsKey("foo")
func (v *Value) Object() *Object { _ = "STUB: not implemented"; return nil }

// Array returns a new Array attached to underlying value.
//
// If underlying value is not an array ([]interface{}), failure is reported and empty
// (but non-nil) value is returned.
//
// Example:
//
//	value := NewValue(t, []interface{}{"foo", 123})
//	value.Array().ConsistsOf("foo", 123)
func (v *Value) Array() *Array { _ = "STUB: not implemented"; return nil }

// String returns a new String attached to underlying value.
//
// If underlying value is not a string, failure is reported and empty (but non-nil)
// value is returned.
//
// Example:
//
//	value := NewValue(t, "foo")
//	value.String().IsEqualFold("FOO")
func (v *Value) String() *String { _ = "STUB: not implemented"; return nil }

// Number returns a new Number attached to underlying value.
//
// If underlying value is not a number (numeric type convertible to float64), failure
// is reported and empty (but non-nil) value is returned.
//
// Example:
//
//	value := NewValue(t, 123)
//	value.Number().InRange(100, 200)
func (v *Value) Number() *Number { _ = "STUB: not implemented"; return nil }

// Boolean returns a new Boolean attached to underlying value.
//
// If underlying value is not a bool, failure is reported and empty (but non-nil)
// value is returned.
//
// Example:
//
//	value := NewValue(t, true)
//	value.Boolean().IsTrue()
func (v *Value) Boolean() *Boolean { _ = "STUB: not implemented"; return nil }

// IsNull succeeds if value is nil.
//
// Note that non-nil interface{} that points to nil value (e.g. nil slice or map)
// is also treated as null value. Empty (non-nil) slice or map, empty string, and
// zero number are not treated as null value.
//
// Example:
//
//	value := NewValue(t, nil)
//	value.IsNull()
//
//	value := NewValue(t, []interface{}(nil))
//	value.IsNull()
func (v *Value) IsNull() *Value { _ = "STUB: not implemented"; return nil }

// NotNull succeeds if value is not nil.
//
// Note that non-nil interface{} that points to nil value (e.g. nil slice or map)
// is also treated as null value. Empty (non-nil) slice or map, empty string, and
// zero number are not treated as null value.
//
// Example:
//
//	value := NewValue(t, "")
//	value.NotNull()
//
//	value := NewValue(t, make([]interface{}, 0))
//	value.NotNull()
func (v *Value) NotNull() *Value { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsNull instead.
func (v *Value) Null() *Value {
	_ = "STUB: not implemented"

	// IsObject succeeds if the underlying value is an object.
	//
	// If underlying value is not an object (map[string]interface{}), failure is reported.
	//
	// Example:
	//
	//	value := NewValue(t, map[string]interface{}{"foo": 123})
	//	value.IsObject()
	return nil
}

func (v *Value) IsObject() *Value { _ = "STUB: not implemented"; return nil }

// NotObject succeeds if the underlying value is not an object.
//
// If underlying value is an object (map[string]interface{}), failure is reported.
//
// Example:
//
//	value := NewValue(t, nil)
//	value.NotObject()
func (v *Value) NotObject() *Value { _ = "STUB: not implemented"; return nil }

// IsArray succeeds if the underlying value is an array.
//
// If underlying value is not an array ([]interface{}), failure is reported.
//
// Example:
//
//	value := NewValue(t, []interface{}{"foo", "123"})
//	value.IsArray()
func (v *Value) IsArray() *Value { _ = "STUB: not implemented"; return nil }

// NotArray succeeds if the underlying value is not an array.
//
// If underlying value is an array ([]interface{}), failure is reported.
//
// Example:
//
//	value := NewValue(t, nil)
//	value.NotArray()
func (v *Value) NotArray() *Value { _ = "STUB: not implemented"; return nil }

// IsString succeeds if the underlying value is a string.
//
// If underlying value is not a string, failure is reported.
//
// Example:
//
//	value := NewValue(t, "foo")
//	value.IsString()
func (v *Value) IsString() *Value { _ = "STUB: not implemented"; return nil }

// NotString succeeds if the underlying value is not a string.
//
// If underlying value is a string, failure is reported.
//
// Example:
//
//	value := NewValue(t, nil)
//	value.NotString()
func (v *Value) NotString() *Value { _ = "STUB: not implemented"; return nil }

// IsNumber succeeds if the underlying value is a number.
//
// If underlying value is not a number (numeric type convertible to float64),
// failure is reported
//
// Example:
//
//	value := NewValue(t, 123)
//	value.IsNumber()
func (v *Value) IsNumber() *Value { _ = "STUB: not implemented"; return nil }

// NotNumber succeeds if the underlying value is a not a number.
//
// If underlying value is a number (numeric type convertible to float64),
// failure is reported
//
// Example:
//
//	value := NewValue(t, nil)
//	value.NotNumber()
func (v *Value) NotNumber() *Value { _ = "STUB: not implemented"; return nil }

// IsBoolean succeeds if the underlying value is a boolean.
//
// If underlying value is not a boolean, failure is reported.
//
// Example:
//
//	value := NewValue(t, true)
//	value.IsBoolean()
func (v *Value) IsBoolean() *Value { _ = "STUB: not implemented"; return nil }

// NotBoolean succeeds if the underlying value is not a boolean.
//
// If underlying value is a boolean, failure is reported.
//
// Example:
//
//	value := NewValue(t, nil)
//	value.NotBoolean()
func (v *Value) NotBoolean() *Value { _ = "STUB: not implemented"; return nil }

// IsEqual succeeds if value is equal to another value (e.g. map, slice, string, etc).
// Before comparison, both values are converted to canonical form.
//
// Example:
//
//	value := NewValue(t, "foo")
//	value.IsEqual("foo")
func (v *Value) IsEqual(value interface{}) *Value { _ = "STUB: not implemented"; return nil }

// NotEqual succeeds if value is not equal to another value (e.g. map, slice,
// string, etc). Before comparison, both values are converted to canonical form.
//
// Example:
//
//	value := NewValue(t, "foo")
//	value.NorEqual("bar")
func (v *Value) NotEqual(value interface{}) *Value { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsEqual instead.
func (v *Value) Equal(value interface{}) *Value { _ = "STUB: not implemented"; return nil }

// InList succeeds if whole value is equal to one of the values from given
// list of values (e.g. map, slice, string, etc). Before comparison, all
// values are converted to canonical form.
//
// If at least one value has wrong type, failure is reported.
//
// Example:
//
//	value := NewValue(t, "foo")
//	value.InList("foo", 123)
func (v *Value) InList(values ...interface{}) *Value { _ = "STUB: not implemented"; return nil }

// continue loop to check that all values are correct

// NotInList succeeds if the whole value is not equal to any of the values from
// given list of values (e.g. map, slice, string, etc).
// Before comparison, all values are converted to canonical form.
//
// If at least one value has wrong type, failure is reported.
//
// Example:
//
//	value := NewValue(t, "foo")
//	value.NotInList("bar", 123)
func (v *Value) NotInList(values ...interface{}) *Value { _ = "STUB: not implemented"; return nil }
