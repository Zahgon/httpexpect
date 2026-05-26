package httpexpect

// Object provides methods to inspect attached map[string]interface{} object
// (Go representation of JSON object).
type Object struct {
	noCopy noCopy
	chain  *chain
	value  map[string]interface{}
}

// NewObject returns a new Object instance.
//
// If reporter is nil, the function panics.
// If value is nil, failure is reported.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123})
func NewObject(reporter Reporter, value map[string]interface{}) *Object {
	_ = "STUB: not implemented"
	return nil
}

// NewObjectC returns a new Object instance with config.
//
// Requirements for config are same as for WithConfig function.
// If value is nil, failure is reported.
//
// Example:
//
//	object := NewObjectC(config, map[string]interface{}{"foo": 123})
func NewObjectC(config Config, value map[string]interface{}) *Object {
	_ = "STUB: not implemented"
	return nil
}

func newObject(parent *chain, val map[string]interface{}) *Object {
	_ = "STUB: not implemented"
	return nil
}

// Raw returns underlying value attached to Object.
// This is the value originally passed to NewObject, converted to canonical form.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123})
//	assert.Equal(t, map[string]interface{}{"foo": 123.0}, object.Raw())
func (o *Object) Raw() map[string]interface{} {
	_ = "STUB: not implemented"

	// Decode unmarshals the underlying value attached to the Object to a target variable
	// target should be one of this:
	//
	//   - pointer to an empty interface
	//   - pointer to a map
	//   - pointer to a struct
	//
	// Example:
	//
	//	type S struct{
	//		Foo int                    `json:"foo"`
	//		Bar []interface{}          `json:"bar"`
	//		Baz map[string]interface{} `json:"baz"`
	//		Bat struct{ A int }        `json:"bat"`
	//	}
	//
	//	m := map[string]interface{}{
	//		"foo": 123,
	//		"bar": []interface{}{"123", 234.0},
	//		"baz": map[string]interface{}{
	//			"a": "b",
	//		},
	//		"bat": struct{ A int }{123},
	//	}
	//
	//	value := NewObject(t, value)
	//
	//	var target S
	//	value.Decode(&target)
	//
	//	assert.Equal(t, S{123,[]interface{}{"123", 234.0},
	//		map[string]interface{}{"a": "b"}, struct{ A int }{123},
	//	}, target)
	return nil
}

func (o *Object) Decode(target interface{}) *Object { _ = "STUB: not implemented"; return nil }

// Alias is similar to Value.Alias.
func (o *Object) Alias(name string) *Object { _ = "STUB: not implemented"; return nil }

// Path is similar to Value.Path.
func (o *Object) Path(path string) *Value { _ = "STUB: not implemented"; return nil }

// Schema is similar to Value.Schema.
func (o *Object) Schema(schema interface{}) *Object { _ = "STUB: not implemented"; return nil }

// Length returns a new Number instance with value count.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123, "bar": 456})
//	object.Length().IsEqual(2)
func (o *Object) Length() *Number { _ = "STUB: not implemented"; return nil }

// Keys returns a new Array instance with object's keys.
// Keys are sorted in ascending order.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123, "bar": 456})
//	object.Keys().ContainsOnly("foo", "bar")
func (o *Object) Keys() *Array { _ = "STUB: not implemented"; return nil }

// Values returns a new Array instance with object's values.
// Values are sorted by keys ascending order.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123, "bar": 456})
//	object.Values().ContainsOnly(123, 456)
func (o *Object) Values() *Array { _ = "STUB: not implemented"; return nil }

// Value returns a new Value instance with value for given key.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123})
//	object.Value("foo").Number().IsEqual(123)
func (o *Object) Value(key string) *Value { _ = "STUB: not implemented"; return nil }

// HasValue succeeds if object's value for given key is equal to given value.
// Before comparison, both values are converted to canonical form.
//
// value should be map[string]interface{} or struct.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123})
//	object.HasValue("foo", 123)
func (o *Object) HasValue(key string, value interface{}) *Object {
	_ = "STUB: not implemented"
	return nil
}

// NotHasValue succeeds if object's value for given key is not equal to given
// value. Before comparison, both values are converted to canonical form.
//
// value should be map[string]interface{} or struct.
//
// If object doesn't contain any value for given key, failure is reported.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123})
//	object.NotHasValue("foo", "bad value")  // success
//	object.NotHasValue("bar", "bad value")  // failure! (key is missing)
func (o *Object) NotHasValue(key string, value interface{}) *Object {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: use HasValue instead.
func (o *Object) ValueEqual(key string, value interface{}) *Object {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: use NotHasValue instead.
func (o *Object) ValueNotEqual(key string, value interface{}) *Object {
	_ = "STUB: not implemented"
	return nil
}

// Iter returns a new map of Values attached to object elements.
//
// Example:
//
//	numbers := map[string]interface{}{"foo": 123, "bar": 456}
//	object := NewObject(t, numbers)
//
//	for key, value := range object.Iter() {
//		value.Number().IsEqual(numbers[key])
//	}
func (o *Object) Iter() map[string]Value { _ = "STUB: not implemented"; return nil }

// Every runs the passed function for all the key value pairs in the object.
//
// If assertion inside function fails, the original Object is marked failed.
//
// Every will execute the function for all values in the object irrespective
// of assertion failures for some values in the object.
//
// The function is invoked for key value pairs sorted by keys in ascending order.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123, "bar": 456})
//
//	object.Every(func(key string, value *httpexpect.Value) {
//	  value.String().NotEmpty()
//	})
func (o *Object) Every(fn func(key string, value *Value)) *Object {
	_ = "STUB: not implemented"
	return nil
}

// Filter accepts a function that returns a boolean. The function is ran
// over the object elements. If the function returns true, the element passes
// the filter and is added to the new object of filtered elements. If false,
// the value is skipped (or in other words filtered out). After iterating
// through all the elements of the original object, the new filtered object
// is returned.
//
// If there are any failed assertions in the filtering function, the
// element is omitted without causing test failure.
//
// The function is invoked for key value pairs sorted by keys in ascending order.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{
//		"foo": "bar",
//		"baz": 6,
//		"qux": "quux",
//	})
//	filteredObject := object.Filter(func(key string, value *httpexpect.Value) bool {
//		value.String().NotEmpty()		//fails on 6
//		return value.Raw() != "bar"		//fails on "bar"
//	})
//	filteredObject.IsEqual(map[string]interface{}{"qux":"quux"})	//succeeds
func (o *Object) Filter(fn func(key string, value *Value) bool) *Object {
	_ = "STUB: not implemented"
	return nil
}

// Transform runs the passed function on all the elements in the Object
// and returns a new object without effecting original object.
//
// The function is invoked for key value pairs sorted by keys in ascending order.
//
// Example:
//
//	object := NewObject(t, []interface{}{"x": "foo", "y": "bar"})
//	transformedObject := object.Transform(
//		func(key string, value interface{}) interface{} {
//			return strings.ToUpper(value.(string))
//		})
//	transformedObject.IsEqual([]interface{}{"x": "FOO", "y": "BAR"})
func (o *Object) Transform(fn func(key string, value interface{}) interface{}) *Object {
	_ = "STUB: not implemented"
	return nil
}

// Find accepts a function that returns a boolean, runs it over the object
// elements, and returns the first element on which it returned true.
//
// If there are any failed assertions in the predicate function, the
// element is skipped without causing test failure.
//
// If no elements were found, a failure is reported.
//
// The function is invoked for key value pairs sorted by keys in ascending order.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{
//		"a": 1,
//		"b": "foo",
//		"c": 101,
//		"d": "bar",
//		"e": 201,
//	})
//	foundValue := object.Find(func(key string, value *httpexpect.Value)  bool {
//		num := value.Number()      // skip if element is not a string
//		return num.Raw() > 100     // check element value
//	})
//	foundValue.IsEqual(101) // succeeds
func (o *Object) Find(fn func(key string, value *Value) bool) *Value {
	_ = "STUB: not implemented"
	return nil
}

// FindAll accepts a function that returns a boolean, runs it over the object
// elements, and returns all the elements on which it returned true.
//
// If there are any failed assertions in the predicate function, the
// element is skipped without causing test failure.
//
// If no elements were found, empty slice is returned without reporting error.
//
// The function is invoked for key value pairs sorted by keys in ascending order.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{
//		"a": 1,
//		"b": "foo",
//		"c": 101,
//		"d": "bar",
//		"e": 201,
//	})
//	foundValues := object.FindAll(func(key string, value *httpexpect.Value)  bool {
//		num := value.Number()      // skip if element is not a string
//		return num.Raw() > 100     // check element value
//	})
//
//	assert.Equal(t, len(foundValues), 2)
//	foundValues[0].IsEqual(101)
//	foundValues[1].IsEqual(201)
func (o *Object) FindAll(fn func(key string, value *Value) bool) []*Value {
	_ = "STUB: not implemented"
	return nil
}

// NotFind accepts a function that returns a boolean, runs it over the object
// elelements, and checks that it does not return true for any of the elements.
//
// If there are any failed assertions in the predicate function, the
// element is skipped without causing test failure.
//
// If the predicate function did not fail and returned true for at least
// one element, a failure is reported.
//
// The function is invoked for key value pairs sorted by keys in ascending order.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{
//		"a": 1,
//		"b": "foo",
//		"c": 2,
//		"d": "bar",
//	})
//	object.NotFind(func(key string, value *httpexpect.Value) bool {
//		num := value.Number()    // skip if element is not a number
//		return num.Raw() > 100   // check element value
//	}) // succeeds
func (o *Object) NotFind(fn func(key string, value *Value) bool) *Object {
	_ = "STUB: not implemented"
	return nil
}

// IsEmpty succeeds if object is empty.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{})
//	object.IsEmpty()
func (o *Object) IsEmpty() *Object { _ = "STUB: not implemented"; return nil }

// NotEmpty succeeds if object is non-empty.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123})
//	object.NotEmpty()
func (o *Object) NotEmpty() *Object { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsEmpty instead.
func (o *Object) Empty() *Object {
	_ = "STUB: not implemented"

	// IsEqual succeeds if object is equal to given value.
	// Before comparison, both object and value are converted to canonical form.
	//
	// value should be map[string]interface{} or struct.
	//
	// Example:
	//
	//	object := NewObject(t, map[string]interface{}{"foo": 123})
	//	object.IsEqual(map[string]interface{}{"foo": 123})
	return nil
}

func (o *Object) IsEqual(value interface{}) *Object { _ = "STUB: not implemented"; return nil }

// NotEqual succeeds if object is not equal to given value.
// Before comparison, both object and value are converted to canonical form.
//
// value should be map[string]interface{} or struct.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123})
//	object.IsEqual(map[string]interface{}{"bar": 123})
func (o *Object) NotEqual(value interface{}) *Object { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsEqual instead.
func (o *Object) Equal(value interface{}) *Object { _ = "STUB: not implemented"; return nil }

// InList succeeds if whole object is equal to one of the values from given list
// of objects. Before comparison, each value is converted to canonical form.
//
// Each value should be map[string]interface{} or struct. If at least one value
// has wrong type, failure is reported.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123})
//	object.InList(
//		map[string]interface{}{"foo": 123},
//		map[string]interface{}{"bar": 456},
//	)
func (o *Object) InList(values ...interface{}) *Object { _ = "STUB: not implemented"; return nil }

// continue loop to check that all values are correct

// NotInList succeeds if the whole object is not equal to any of the values
// from given list of objects. Before comparison, each value is converted to
// canonical form.
//
// Each value should be map[string]interface{} or struct. If at least one value
// has wrong type, failure is reported.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123})
//	object.NotInList(
//		map[string]interface{}{"bar": 456},
//		map[string]interface{}{"baz": 789},
//	)
func (o *Object) NotInList(values ...interface{}) *Object { _ = "STUB: not implemented"; return nil }

// ContainsKey succeeds if object contains given key.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123})
//	object.ContainsKey("foo")
func (o *Object) ContainsKey(key string) *Object { _ = "STUB: not implemented"; return nil }

// NotContainsKey succeeds if object doesn't contain given key.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123})
//	object.NotContainsKey("bar")
func (o *Object) NotContainsKey(key string) *Object { _ = "STUB: not implemented"; return nil }

// ContainsValue succeeds if object contains given value with any key.
// Before comparison, both object and value are converted to canonical form.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123})
//	object.ContainsValue(123)
func (o *Object) ContainsValue(value interface{}) *Object { _ = "STUB: not implemented"; return nil }

// NotContainsValue succeeds if object does not contain given value with any key.
// Before comparison, both object and value are converted to canonical form.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123})
//	object.NotContainsValue(456)
func (o *Object) NotContainsValue(value interface{}) *Object { _ = "STUB: not implemented"; return nil }

// ContainsSubset succeeds if given value is a subset of object.
// Before comparison, both object and value are converted to canonical form.
//
// value should be map[string]interface{} or struct.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{
//		"foo": 123,
//		"bar": []interface{}{"x", "y"},
//		"bar": map[string]interface{}{
//			"a": true,
//			"b": false,
//		},
//	})
//
//	object.ContainsSubset(map[string]interface{}{  // success
//		"foo": 123,
//		"bar": map[string]interface{}{
//			"a": true,
//		},
//	})
//
//	object.ContainsSubset(map[string]interface{}{  // failure
//		"foo": 123,
//		"qux": 456,
//	})
//
//	object.ContainsSubset(map[string]interface{}{  // failure, slices should match exactly
//		"bar": []interface{}{"x"},
//	})
func (o *Object) ContainsSubset(value interface{}) *Object { _ = "STUB: not implemented"; return nil }

// NotContainsSubset succeeds if given value is not a subset of object.
// Before comparison, both object and value are converted to canonical form.
//
// value should be map[string]interface{} or struct.
//
// Example:
//
//	object := NewObject(t, map[string]interface{}{"foo": 123, "bar": 456})
//	object.NotContainsSubset(map[string]interface{}{"foo": 123, "bar": "no-no-no"})
func (o *Object) NotContainsSubset(value interface{}) *Object {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: use ContainsSubset instead.
func (o *Object) ContainsMap(value interface{}) *Object { _ = "STUB: not implemented"; return nil }

// Deprecated: use NotContainsSubset instead.
func (o *Object) NotContainsMap(value interface{}) *Object { _ = "STUB: not implemented"; return nil }

type kv struct {
	key string
	val interface{}
}

func (o *Object) sortedKV() []kv { _ = "STUB: not implemented"; return nil }

func containsKey(
	opChain *chain, obj map[string]interface{}, key string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func containsValue(
	opChain *chain, obj map[string]interface{}, val interface{},
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func containsSubset(
	opChain *chain, obj map[string]interface{}, val interface{},
) bool {
	_ = "STUB: not implemented"
	return false
}

func isSubset(outer, inner map[string]interface{}) bool { _ = "STUB: not implemented"; return false }
