package httpexpect

import (
	"math/big"
)

// Number provides methods to inspect attached float64 value
// (Go representation of JSON number).
type Number struct {
	noCopy noCopy
	chain  *chain
	value  float64
}

// NewNumber returns a new Number instance.
//
// If reporter is nil, the function panics.
//
// Example:
//
//	number := NewNumber(t, 123.4)
func NewNumber(reporter Reporter, value float64) *Number { _ = "STUB: not implemented"; return nil }

// NewNumberC returns a new Number instance with config.
//
// Requirements for config are same as for WithConfig function.
//
// Example:
//
//	number := NewNumberC(config, 123.4)
func NewNumberC(config Config, value float64) *Number { _ = "STUB: not implemented"; return nil }

func newNumber(parent *chain, val float64) *Number { _ = "STUB: not implemented"; return nil }

// Raw returns underlying value attached to Number.
// This is the value originally passed to NewNumber.
//
// Example:
//
//	number := NewNumber(t, 123.4)
//	assert.Equal(t, 123.4, number.Raw())
func (n *Number) Raw() float64 {
	_ = "STUB: not implemented"

	// Decode unmarshals the underlying value attached to the Number to a target variable.
	// target should be one of these:
	//
	//   - pointer to an empty interface
	//   - pointer to any integer or floating type
	//
	// Example:
	//
	//	value := NewNumber(t, 123)
	//
	//	var target interface{}
	//	valude.decode(&target)
	//
	//	assert.Equal(t, 123, target)
	return 0
}

func (n *Number) Decode(target interface{}) *Number { _ = "STUB: not implemented"; return nil }

// Alias is similar to Value.Alias.
func (n *Number) Alias(name string) *Number { _ = "STUB: not implemented"; return nil }

// Path is similar to Value.Path.
func (n *Number) Path(path string) *Value { _ = "STUB: not implemented"; return nil }

// Schema is similar to Value.Schema.
func (n *Number) Schema(schema interface{}) *Number { _ = "STUB: not implemented"; return nil }

// IsEqual succeeds if number is equal to given value.
//
// value should have numeric type convertible to float64. Before comparison,
// it is converted to float64.
//
// Example:
//
//	number := NewNumber(t, 123)
//	number.IsEqual(float64(123))
//	number.IsEqual(int32(123))
func (n *Number) IsEqual(value interface{}) *Number { _ = "STUB: not implemented"; return nil }

// NotEqual succeeds if number is not equal to given value.
//
// value should have numeric type convertible to float64. Before comparison,
// it is converted to float64.
//
// Example:
//
//	number := NewNumber(t, 123)
//	number.NotEqual(float64(321))
//	number.NotEqual(int32(321))
func (n *Number) NotEqual(value interface{}) *Number { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsEqual instead.
func (n *Number) Equal(value interface{}) *Number { _ = "STUB: not implemented"; return nil }

// InDelta succeeds if two numerals are within delta of each other.
//
// Example:
//
//	number := NewNumber(t, 123.0)
//	number.InDelta(123.2, 0.3)
func (n *Number) InDelta(value, delta float64) *Number { _ = "STUB: not implemented"; return nil }

// NotInDelta succeeds if two numerals are not within delta of each other.
//
// Example:
//
//	number := NewNumber(t, 123.0)
//	number.NotInDelta(123.2, 0.1)
func (n *Number) NotInDelta(value, delta float64) *Number { _ = "STUB: not implemented"; return nil }

// Deprecated: use InDelta instead.
func (n *Number) EqualDelta(value, delta float64) *Number { _ = "STUB: not implemented"; return nil }

// Deprecated: use NotInDelta instead.
func (n *Number) NotEqualDelta(value, delta float64) *Number { _ = "STUB: not implemented"; return nil }

// InDeltaRelative succeeds if two numbers are within relative delta of each other.
//
// The relative delta is expressed as a decimal. For example, to determine if a number
// and a value are within 1% of each other, use 0.01.
//
// A number and a value are within relative delta if
// Abs(number-value) / Abs(number) < relative delta.
//
// Please note that number, value, and delta can't be NaN, number and value can't
// be opposite Inf and delta cannot be Inf.
//
// Example:
//
//	number := NewNumber(t, 123.0)
//	number.InDeltaRelative(126.5, 0.03)
func (n *Number) InDeltaRelative(value, delta float64) *Number {
	_ = "STUB: not implemented"
	return nil
}

// Fail if any of the numbers is NaN with specific error message

// Pass if number and value are +-Inf and equal,
// regardless if delta is 0 or positive number

// Fail if number and value are +=Inf and unequal with specific error message

// Normal comparison after filtering out all corner cases

// NotInDeltaRelative succeeds if two numbers aren't within relative delta of each other.
//
// The relative delta is expressed as a decimal. For example, to determine if a number
// and a value are within 1% of each other, use 0.01.
//
// A number and a value are within relative delta if
// Abs(number-value) / Abs(number) < relative delta.
//
// Please note that number, value, and delta can't be NaN, number and value can't
// be opposite Inf and delta cannot be Inf.
//
// Example:
//
//	number := NewNumber(t, 123.0)
//	number.NotInDeltaRelative(126.5, 0.01)
func (n *Number) NotInDeltaRelative(value, delta float64) *Number {
	_ = "STUB: not implemented"
	return nil
}

// Fail if any of the numbers is NaN with specific error message

// Fail if number and value are +-Inf and equal,
// regardless if delta is 0 or positive number

// Pass if number and value are +=Inf and unequal

// Normal comparison after filtering out all corner cases

// InRange succeeds if number is within given range [min; max].
//
// min and max should have numeric type convertible to float64. Before comparison,
// they are converted to float64.
//
// Example:
//
//	number := NewNumber(t, 123)
//	number.InRange(float32(100), int32(200))  // success
//	number.InRange(100, 200)                  // success
//	number.InRange(123, 123)                  // success
func (n *Number) InRange(min, max interface{}) *Number { _ = "STUB: not implemented"; return nil }

// NotInRange succeeds if number is not within given range [min; max].
//
// min and max should have numeric type convertible to float64. Before comparison,
// they are converted to float64.
//
// Example:
//
//	number := NewNumber(t, 100)
//	number.NotInRange(0, 99)
//	number.NotInRange(101, 200)
func (n *Number) NotInRange(min, max interface{}) *Number { _ = "STUB: not implemented"; return nil }

// InList succeeds if the number is equal to one of the values from given list
// of numbers. Before comparison, each value is converted to canonical form.
//
// Each value should be numeric type convertible to float64. If at least one
// value has wrong type, failure is reported.
//
// Example:
//
//	number := NewNumber(t, 123)
//	number.InList(float64(123), int32(123))
func (n *Number) InList(values ...interface{}) *Number { _ = "STUB: not implemented"; return nil }

// continue loop to check that all values are correct

// NotInList succeeds if the number is not equal to any of the values from given
// list of numbers. Before comparison, each value is converted to canonical form.
//
// Each value should be numeric type convertible to float64. If at least one
// value has wrong type, failure is reported.
//
// Example:
//
//	number := NewNumber(t, 123)
//	number.NotInList(float64(456), int32(456))
func (n *Number) NotInList(values ...interface{}) *Number { _ = "STUB: not implemented"; return nil }

// IsGt succeeds if number is greater than given value.
//
// value should have numeric type convertible to float64. Before comparison,
// it is converted to float64.
//
// Example:
//
//	number := NewNumber(t, 123)
//	number.IsGt(float64(122))
//	number.IsGt(int32(122))
func (n *Number) IsGt(value interface{}) *Number { _ = "STUB: not implemented"; return nil }

// IsGe succeeds if number is greater than or equal to given value.
//
// value should have numeric type convertible to float64. Before comparison,
// it is converted to float64.
//
// Example:
//
//	number := NewNumber(t, 123)
//	number.IsGe(float64(122))
//	number.IsGe(int32(122))
func (n *Number) IsGe(value interface{}) *Number { _ = "STUB: not implemented"; return nil }

// IsLt succeeds if number is lesser than given value.
//
// value should have numeric type convertible to float64. Before comparison,
// it is converted to float64.
//
// Example:
//
//	number := NewNumber(t, 123)
//	number.IsLt(float64(124))
//	number.IsLt(int32(124))
func (n *Number) IsLt(value interface{}) *Number { _ = "STUB: not implemented"; return nil }

// IsLe succeeds if number is lesser than or equal to given value.
//
// value should have numeric type convertible to float64. Before comparison,
// it is converted to float64.
//
// Example:
//
//	number := NewNumber(t, 123)
//	number.IsLe(float64(124))
//	number.IsLe(int32(124))
func (n *Number) IsLe(value interface{}) *Number { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsGt instead.
func (n *Number) Gt(value interface{}) *Number { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsGe instead.
func (n *Number) Ge(value interface{}) *Number { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsLt instead.
func (n *Number) Lt(value interface{}) *Number { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsLe instead.
func (n *Number) Le(value interface{}) *Number { _ = "STUB: not implemented"; return nil }

// IsInt succeeds if number is a signed integer of the specified bit width
// as an optional argument.
//
// Bits argument defines maximum allowed bitness for the given number.
// If bits is omitted, boundary check is omitted too.
//
// Example:
//
//	number := NewNumber(t, 1000000)
//	number.IsInt()   // success
//	number.IsInt(32) // success
//	number.IsInt(16) // failure
//
//	number := NewNumber(t, -1000000)
//	number.IsInt()   // success
//	number.IsInt(32) // success
//	number.IsInt(16) // failure
//
//	number := NewNumber(t, 0.5)
//	number.IsInt()   // failure
func (n *Number) IsInt(bits ...int) *Number { _ = "STUB: not implemented"; return nil }

// NotInt succeeds if number is not a signed integer of the specified bit
// width as an optional argument.
//
// Bits argument defines maximum allowed bitness for the given number.
// If bits is omitted, boundary check is omitted too.
//
// Example:
//
//	number := NewNumber(t, 1000000)
//	number.NotInt()   // failure
//	number.NotInt(32) // failure
//	number.NotInt(16) // success
//
//	number := NewNumber(t, -1000000)
//	number.NotInt()   // failure
//	number.NotInt(32) // failure
//	number.NotInt(16) // success
//
//	number := NewNumber(t, 0.5)
//	number.NotInt()   // success
func (n *Number) NotInt(bits ...int) *Number { _ = "STUB: not implemented"; return nil }

// IsUint succeeds if number is an unsigned integer of the specified bit
// width as an optional argument.
//
// Bits argument defines maximum allowed bitness for the given number.
// If bits is omitted, boundary check is omitted too.
//
// Example:
//
//	number := NewNumber(t, 1000000)
//	number.IsUint()   // success
//	number.IsUint(32) // success
//	number.IsUint(16) // failure
//
//	number := NewNumber(t, -1000000)
//	number.IsUint()   // failure
//	number.IsUint(32) // failure
//	number.IsUint(16) // failure
//
//	number := NewNumber(t, 0.5)
//	number.IsUint()   // failure
func (n *Number) IsUint(bits ...int) *Number { _ = "STUB: not implemented"; return nil }

// NotUint succeeds if number is not an unsigned integer of the specified bit
// width as an optional argument.
//
// Bits argument defines maximum allowed bitness for the given number.
// If bits is omitted, boundary check is omitted too.
//
// Example:
//
//	number := NewNumber(t, 1000000)
//	number.NotUint()   // failure
//	number.NotUint(32) // failure
//	number.NotUint(16) // success
//
//	number := NewNumber(t, -1000000)
//	number.NotUint()   // success
//	number.NotUint(32) // success
//	number.NotUint(16) // success
//
//	number := NewNumber(t, 0.5)
//	number.NotUint()   // success
func (n *Number) NotUint(bits ...int) *Number { _ = "STUB: not implemented"; return nil }

// IsFinite succeeds if number is neither ±Inf nor NaN.
//
// Example:
//
//	number := NewNumber(t, 1234.5)
//	number.IsFinite() // success
//
//	number := NewNumber(t, math.NaN())
//	number.IsFinite() // failure
//
//	number := NewNumber(t, math.Inf(+1))
//	number.IsFinite() // failure
func (n *Number) IsFinite() *Number { _ = "STUB: not implemented"; return nil }

// NotFinite succeeds if number is either ±Inf or NaN.
//
// Example:
//
//	number := NewNumber(t, 1234.5)
//	number.NotFinite() // failure
//
//	number := NewNumber(t, math.NaN())
//	number.NotFinite() // success
//
//	number := NewNumber(t, math.Inf(+1))
//	number.NotFinite() // success
func (n *Number) NotFinite() *Number { _ = "STUB: not implemented"; return nil }

type intBoundary struct {
	val  *big.Int
	sign int
	bits int
}

func (b intBoundary) String() string { _ = "STUB: not implemented"; return "" }

type relativeDelta float64

func (rd relativeDelta) String() string { _ = "STUB: not implemented"; return "" }

func deltaRelativeErrorCheck(inDeltaRelative bool, number, value, delta float64) bool {
	_ = "STUB: not implemented"
	return false
}
