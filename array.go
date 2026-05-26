package httpexpect

// Array provides methods to inspect attached []interface{} object
// (Go representation of JSON array).
type Array struct {
	noCopy noCopy
	chain  *chain
	value  []interface{}
}

// NewArray returns a new Array instance.
//
// If reporter is nil, the function panics.
// If value is nil, failure is reported.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123})
func NewArray(reporter Reporter, value []interface{}) *Array { _ = "STUB: not implemented"; return nil }

// NewArrayC returns a new Array instance with config.
//
// Requirements for config are same as for WithConfig function.
// If value is nil, failure is reported.
//
// Example:
//
//	array := NewArrayC(config, []interface{}{"foo", 123})
func NewArrayC(config Config, value []interface{}) *Array { _ = "STUB: not implemented"; return nil }

func newArray(parent *chain, val []interface{}) *Array { _ = "STUB: not implemented"; return nil }

// Raw returns underlying value attached to Array.
// This is the value originally passed to NewArray, converted to canonical form.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123})
//	assert.Equal(t, []interface{}{"foo", 123.0}, array.Raw())
func (a *Array) Raw() []interface{} {
	_ = "STUB: not implemented"

	// Decode unmarshals the underlying value attached to the Array to a target variable.
	// target should be one of these:
	//
	//   - pointer to an empty interface
	//   - pointer to a slice of any type
	//
	// Example:
	//
	//	type S struct{
	//		Foo int `json:foo`
	//	}
	//	value := []interface{}{
	//		map[string]interface{}{
	//			"foo": 123,
	//		},
	//		map[string]interface{}{
	//			"foo": 456,
	//		},
	//	}
	//	array := NewArray(t, value)
	//
	//	var target []S
	//	arr.Decode(&target)
	//
	//	assert.Equal(t, []S{{123}, {456}}, target)
	return nil
}

func (a *Array) Decode(target interface{}) *Array { _ = "STUB: not implemented"; return nil }

// Alias is similar to Value.Alias.
func (a *Array) Alias(name string) *Array { _ = "STUB: not implemented"; return nil }

// Path is similar to Value.Path.
func (a *Array) Path(path string) *Value { _ = "STUB: not implemented"; return nil }

// Schema is similar to Value.Schema.
func (a *Array) Schema(schema interface{}) *Array { _ = "STUB: not implemented"; return nil }

// Length returns a new Number instance with array length.
//
// Example:
//
//	array := NewArray(t, []interface{}{1, 2, 3})
//	array.Length().IsEqual(3)
func (a *Array) Length() *Number { _ = "STUB: not implemented"; return nil }

// Value returns a new Value instance with array element for given index.
//
// If index is out of array bounds, Value reports failure and returns empty
// (but non-nil) instance.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123})
//	array.Value(0).String().IsEqual("foo")
//	array.Value(1).Number().IsEqual(123)
func (a *Array) Value(index int) *Value { _ = "STUB: not implemented"; return nil }

// Deprecated: use Value instead.
func (a *Array) Element(index int) *Value { _ = "STUB: not implemented"; return nil }

// HasValue succeeds if array's value at the given index is equal to given value.
//
// Before comparison, both values are converted to canonical form. value should be
// map[string]interface{} or struct.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", "123"})
//	array.HasValue(1, 123)
func (a *Array) HasValue(index int, value interface{}) *Array {
	_ = "STUB: not implemented"
	return nil
}

// NotHasValue succeeds if array's value at the given index is not equal to given value.
//
// Before comparison, both values are converted to canonical form. value should be
// map[string]interface{} or struct.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", "123"})
//	array.NotHasValue(1, 234)
func (a *Array) NotHasValue(index int, value interface{}) *Array {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: use Value or HasValue instead.
func (a *Array) First() *Value { _ = "STUB: not implemented"; return nil }

// Deprecated: use Value or HasValue instead.
func (a *Array) Last() *Value { _ = "STUB: not implemented"; return nil }

// Iter returns a new slice of Values attached to array elements.
//
// Example:
//
//	strings := []interface{}{"foo", "bar"}
//	array := NewArray(t, strings)
//
//	for index, value := range array.Iter() {
//		value.String().IsEqual(strings[index])
//	}
func (a *Array) Iter() []Value { _ = "STUB: not implemented"; return nil }

// Every runs the passed function on all the elements in the array.
//
// If assertion inside function fails, the original Array is marked failed.
//
// Every will execute the function for all values in the array irrespective
// of assertion failures for some values in the array.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", "bar"})
//
//	array.Every(func(index int, value *httpexpect.Value) {
//		value.String().NotEmpty()
//	})
func (a *Array) Every(fn func(index int, value *Value)) *Array {
	_ = "STUB: not implemented"
	return nil
}

// Filter accepts a function that returns a boolean. The function is ran
// over the array elements. If the function returns true, the element passes
// the filter and is added to the new array of filtered elements. If false,
// the element is skipped (or in other words filtered out). After iterating
// through all the elements of the original array, the new filtered array
// is returned.
//
// If there are any failed assertions in the filtering function, the
// element is omitted without causing test failure.
//
// Example:
//
//	array := NewArray(t, []interface{}{1, 2, "foo", "bar"})
//	filteredArray := array.Filter(func(index int, value *httpexpect.Value) bool {
//		value.String().NotEmpty()		//fails on 1 and 2
//		return value.Raw() != "bar"		//fails on "bar"
//	})
//	filteredArray.IsEqual([]interface{}{"foo"})	//succeeds
func (a *Array) Filter(fn func(index int, value *Value) bool) *Array {
	_ = "STUB: not implemented"
	return nil
}

// Transform runs the passed function on all the elements in the array
// and returns a new array without effeecting original array.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", "bar"})
//	transformedArray := array.Transform(
//		func(index int, value interface{}) interface{} {
//			return strings.ToUpper(value.(string))
//		})
//	transformedArray.IsEqual([]interface{}{"FOO", "BAR"})
func (a *Array) Transform(fn func(index int, value interface{}) interface{}) *Array {
	_ = "STUB: not implemented"
	return nil
}

// Find accepts a function that returns a boolean, runs it over the array
// elements, and returns the first element on which it returned true.
//
// If there are any failed assertions in the predicate function, the
// element is skipped without causing test failure.
//
// If no elements were found, a failure is reported.
//
// Example:
//
//	array := NewArray(t, []interface{}{1, "foo", 101, "bar", 201})
//	foundValue := array.Find(func(index int, value *httpexpect.Value) bool {
//		num := value.Number()    // skip if element is not a number
//		return num.Raw() > 100   // check element value
//	})
//	foundValue.IsEqual(101) // succeeds
func (a *Array) Find(fn func(index int, value *Value) bool) *Value {
	_ = "STUB: not implemented"
	return nil
}

// FindAll accepts a function that returns a boolean, runs it over the array
// elements, and returns all the elements on which it returned true.
//
// If there are any failed assertions in the predicate function, the
// element is skipped without causing test failure.
//
// If no elements were found, empty slice is returned without reporting error.
//
// Example:
//
//	array := NewArray(t, []interface{}{1, "foo", 101, "bar", 201})
//	foundValues := array.FindAll(func(index int, value *httpexpect.Value) bool {
//		num := value.Number()   // skip if element is not a number
//		return num.Raw() > 100  // check element value
//	})
//
//	assert.Equal(t, len(foundValues), 2)
//	foundValues[0].IsEqual(101)
//	foundValues[1].IsEqual(201)
func (a *Array) FindAll(fn func(index int, value *Value) bool) []*Value {
	_ = "STUB: not implemented"
	return nil
}

// NotFind accepts a function that returns a boolean, runs it over the array
// elelements, and checks that it does not return true for any of the elements.
//
// If there are any failed assertions in the predicate function, the
// element is skipped without causing test failure.
//
// If the predicate function did not fail and returned true for at least
// one element, a failure is reported.
//
// Example:
//
//	array := NewArray(t, []interface{}{1, "foo", 2, "bar"})
//	array.NotFind(func(index int, value *httpexpect.Value) bool {
//		num := value.Number()    // skip if element is not a number
//		return num.Raw() > 100   // check element value
//	}) // succeeds
func (a *Array) NotFind(fn func(index int, value *Value) bool) *Array {
	_ = "STUB: not implemented"
	return nil
}

// IsEmpty succeeds if array is empty.
//
// Example:
//
//	array := NewArray(t, []interface{}{})
//	array.IsEmpty()
func (a *Array) IsEmpty() *Array { _ = "STUB: not implemented"; return nil }

// NotEmpty succeeds if array is non-empty.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123})
//	array.NotEmpty()
func (a *Array) NotEmpty() *Array { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsEmpty instead.
func (a *Array) Empty() *Array {
	_ = "STUB: not implemented"

	// IsEqual succeeds if array is equal to given value.
	// Before comparison, both array and value are converted to canonical form.
	//
	// value should be a slice of any type.
	//
	// Example:
	//
	//	array := NewArray(t, []interface{}{"foo", 123})
	//	array.IsEqual([]interface{}{"foo", 123})
	//
	//	array := NewArray(t, []interface{}{"foo", "bar"})
	//	array.IsEqual([]string{}{"foo", "bar"})
	//
	//	array := NewArray(t, []interface{}{123, 456})
	//	array.IsEqual([]int{}{123, 456})
	return nil
}

func (a *Array) IsEqual(value interface{}) *Array { _ = "STUB: not implemented"; return nil }

// NotEqual succeeds if array is not equal to given value.
// Before comparison, both array and value are converted to canonical form.
//
// value should be a slice of any type.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123})
//	array.NotEqual([]interface{}{123, "foo"})
func (a *Array) NotEqual(value interface{}) *Array { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsEqual instead.
func (a *Array) Equal(value interface{}) *Array { _ = "STUB: not implemented"; return nil }

// IsEqualUnordered succeeds if array is equal to another array, ignoring element
// order. Before comparison, both arrays are converted to canonical form.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123})
//	array.IsEqualUnordered([]interface{}{123, "foo"})
func (a *Array) IsEqualUnordered(value interface{}) *Array { _ = "STUB: not implemented"; return nil }

// NotEqualUnordered succeeds if array is not equal to another array, ignoring
// element order. Before comparison, both arrays are converted to canonical form.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123})
//	array.NotEqualUnordered([]interface{}{123, "foo", "bar"})
func (a *Array) NotEqualUnordered(value interface{}) *Array { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsEqualUnordered instead.
func (a *Array) EqualUnordered(value interface{}) *Array { _ = "STUB: not implemented"; return nil }

// InList succeeds if the whole array is equal to one of the values from given
// list of arrays. Before comparison, both array and each value are converted
// to canonical form.
//
// Each value should be a slice of any type. If at least one value has wrong
// type, failure is reported.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123})
//	array.InList([]interface{}{"foo", 123}, []interface{}{"bar", "456"})
func (a *Array) InList(values ...interface{}) *Array { _ = "STUB: not implemented"; return nil }

// continue loop to check that all values are correct

// NotInList succeeds if the whole array is not equal to any of the values from
// given list of arrays. Before comparison, both array and each value are
// converted to canonical form.
//
// Each value should be a slice of any type. If at least one value has wrong
// type, failure is reported.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123})
//	array.NotInList([]interface{}{"bar", 456}, []interface{}{"baz", "foo"})
func (a *Array) NotInList(values ...interface{}) *Array { _ = "STUB: not implemented"; return nil }

// ConsistsOf succeeds if array contains all given elements, in given order, and only
// them. Before comparison, array and all elements are converted to canonical form.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123})
//	array.ConsistsOf("foo", 123)
//
// These calls are equivalent:
//
//	array.ConsistsOf("a", "b")
//	array.IsEqual([]interface{}{"a", "b"})
func (a *Array) ConsistsOf(values ...interface{}) *Array { _ = "STUB: not implemented"; return nil }

// NotConsistsOf is opposite to ConsistsOf.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123})
//	array.NotConsistsOf("foo")
//	array.NotConsistsOf("foo", 123, 456)
//	array.NotConsistsOf(123, "foo")
//
// These calls are equivalent:
//
//	array.NotConsistsOf("a", "b")
//	array.NotEqual([]interface{}{"a", "b"})
func (a *Array) NotConsistsOf(values ...interface{}) *Array { _ = "STUB: not implemented"; return nil }

// Deprecated: use ConsistsOf instead.
func (a *Array) Elements(values ...interface{}) *Array { _ = "STUB: not implemented"; return nil }

// Deprecated: use NotConsistsOf instead.
func (a *Array) NotElements(values ...interface{}) *Array { _ = "STUB: not implemented"; return nil }

// Deprecated: use ContainsAll or ContainsAny instead.
func (a *Array) Contains(values ...interface{}) *Array { _ = "STUB: not implemented"; return nil }

// Deprecated: use NotContainsAll or NotContainsAny instead.
func (a *Array) NotContains(values ...interface{}) *Array { _ = "STUB: not implemented"; return nil }

// ContainsAll succeeds if array contains all given elements (in any order).
// Before comparison, array and all elements are converted to canonical form.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123})
//	array.ContainsAll(123, "foo")
func (a *Array) ContainsAll(values ...interface{}) *Array { _ = "STUB: not implemented"; return nil }

// NotContainsAll succeeds if array does not contain at least one of the elements.
// Before comparison, array and all elements are converted to canonical form.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123})
//	array.NotContainsAll("bar")         // success
//	array.NotContainsAll(123, "foo")    // failure
func (a *Array) NotContainsAll(values ...interface{}) *Array { _ = "STUB: not implemented"; return nil }

// ContainsAny succeeds if array contains at least one element from the given elements.
// Before comparison, array and all elements are converted to canonical form.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123, 123})
//	array.ContainsAny(123, "foo", "FOO") // success
//	array.ContainsAny("FOO") // failure
func (a *Array) ContainsAny(values ...interface{}) *Array { _ = "STUB: not implemented"; return nil }

// NotContainsAny succeeds if none of the given elements are in the array.
// Before comparison, array and all elements are converted to canonical form.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123})
//	array.NotContainsAny("bar", 124) // success
//	array.NotContainsAny(123) // failure
func (a *Array) NotContainsAny(values ...interface{}) *Array { _ = "STUB: not implemented"; return nil }

// ContainsOnly succeeds if array contains all given elements, in any order, and only
// them, ignoring duplicates. Before comparison, array and all elements are converted
// to canonical form.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123, 123})
//	array.ContainsOnly(123, "foo")
//
// These calls are equivalent:
//
//	array.ContainsOnly("a", "b")
//	array.ContainsOnly("b", "a")
func (a *Array) ContainsOnly(values ...interface{}) *Array { _ = "STUB: not implemented"; return nil }

// NotContainsOnly is opposite to ContainsOnly.
//
// Example:
//
//	array := NewArray(t, []interface{}{"foo", 123})
//	array.NotContainsOnly(123)
//	array.NotContainsOnly(123, "foo", "bar")
//
// These calls are equivalent:
//
//	array.NotContainsOnly("a", "b")
//	array.NotContainsOnly("b", "a")
func (a *Array) NotContainsOnly(values ...interface{}) *Array {
	_ = "STUB: not implemented"
	return nil
}

// IsOrdered succeeds if every element is not less than the previous element
// as defined on the given `less` comparator function.
// For default, it will use built-in comparator function for each data type.
// Built-in comparator requires all elements in the array to have same data type.
// Array with 0 or 1 element will always succeed
//
// Example:
//
//	array := NewArray(t, []interface{}{100, 101, 102})
//	array.IsOrdered() // succeeds
//	array.IsOrdered(func(x, y *httpexpect.Value) bool {
//		return x.Number().Raw() < y.Number().Raw()
//	}) // succeeds
func (a *Array) IsOrdered(less ...func(x, y *Value) bool) *Array {
	_ = "STUB: not implemented"
	return nil
}

// NotOrdered succeeds if at least one element is less than the previous element
// as defined on the given `less` comparator function.
// For default, it will use built-in comparator function for each data type.
// Built-in comparator requires all elements in the array to have same data type.
// Array with 0 or 1 element will always succeed
//
// Example:
//
//	array := NewArray(t, []interface{}{102, 101, 100})
//	array.NotOrdered() // succeeds
//	array.NotOrdered(func(x, y *httpexpect.Value) bool {
//		return x.Number().Raw() < y.Number().Raw()
//	}) // succeeds
func (a *Array) NotOrdered(less ...func(x, y *Value) bool) *Array {
	_ = "STUB: not implemented"
	return nil
}

func countElement(array []interface{}, element interface{}) int {
	_ = "STUB: not implemented"
	return 0
}

func builtinComparator(opChain *chain, array []interface{}) func(x, y *Value) bool {
	_ = "STUB: not implemented"
	return nil
}

// ok, do nothing

// `nil` is never less than `nil`

type unquotedType string

func (t unquotedType) String() string { _ = "STUB: not implemented"; return "" }
