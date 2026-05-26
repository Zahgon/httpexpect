package httpexpect

// Boolean provides methods to inspect attached bool value
// (Go representation of JSON boolean).
type Boolean struct {
	noCopy noCopy
	chain  *chain
	value  bool
}

// NewBoolean returns a new Boolean instance.
//
// If reporter is nil, the function panics.
//
// Example:
//
//	boolean := NewBoolean(t, true)
//	boolean.IsTrue()
func NewBoolean(reporter Reporter, value bool) *Boolean { _ = "STUB: not implemented"; return nil }

// NewBooleanC returns a new Boolean instance with config.
//
// Requirements for config are same as for WithConfig function.
//
// Example:
//
//	boolean := NewBooleanC(config, true)
//	boolean.IsTrue()
func NewBooleanC(config Config, value bool) *Boolean { _ = "STUB: not implemented"; return nil }

func newBoolean(parent *chain, val bool) *Boolean { _ = "STUB: not implemented"; return nil }

// Raw returns underlying value attached to Boolean.
// This is the value originally passed to NewBoolean.
//
// Example:
//
//	boolean := NewBoolean(t, true)
//	assert.Equal(t, true, boolean.Raw())
func (b *Boolean) Raw() bool {
	_ = "STUB: not implemented"

	// Decode unmarshals the underlying value attached to the Boolean to a target variable.
	// target should be one of these:
	//
	//   - pointer to an empty interface
	//   - pointer to a boolean
	//
	// Example:
	//
	//	value := NewBoolean(t, true)
	//
	//	var target bool
	//	value.Decode(&target)
	//
	//	assert.Equal(t, true, target)
	return false
}

func (b *Boolean) Decode(target interface{}) *Boolean { _ = "STUB: not implemented"; return nil }

// Alias is similar to Value.Alias.
func (b *Boolean) Alias(name string) *Boolean { _ = "STUB: not implemented"; return nil }

// Path is similar to Value.Path.
func (b *Boolean) Path(path string) *Value { _ = "STUB: not implemented"; return nil }

// Schema is similar to Value.Schema.
func (b *Boolean) Schema(schema interface{}) *Boolean { _ = "STUB: not implemented"; return nil }

// IsTrue succeeds if boolean is true.
//
// Example:
//
//	boolean := NewBoolean(t, true)
//	boolean.IsTrue()
func (b *Boolean) IsTrue() *Boolean { _ = "STUB: not implemented"; return nil }

// IsFalse succeeds if boolean is false.
//
// Example:
//
//	boolean := NewBoolean(t, false)
//	boolean.IsFalse()
func (b *Boolean) IsFalse() *Boolean { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsTrue instead.
func (b *Boolean) True() *Boolean {
	_ = "STUB: not implemented"

	// Deprecated: use IsFalse instead.
	return nil
}

func (b *Boolean) False() *Boolean {
	_ = "STUB: not implemented"

	// IsEqual succeeds if boolean is equal to given value.
	//
	// Example:
	//
	//	boolean := NewBoolean(t, true)
	//	boolean.IsEqual(true)
	return nil
}

func (b *Boolean) IsEqual(value bool) *Boolean { _ = "STUB: not implemented"; return nil }

// NotEqual succeeds if boolean is not equal to given value.
//
// Example:
//
//	boolean := NewBoolean(t, true)
//	boolean.NotEqual(false)
func (b *Boolean) NotEqual(value bool) *Boolean { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsEqual instead.
func (b *Boolean) Equal(value bool) *Boolean { _ = "STUB: not implemented"; return nil }

// InList succeeds if boolean is equal to one of the values from given
// list of booleans.
//
// Example:
//
//	boolean := NewBoolean(t, true)
//	boolean.InList(true, false)
func (b *Boolean) InList(values ...bool) *Boolean { _ = "STUB: not implemented"; return nil }

// NotInList succeeds if boolean is not equal to any of the values from
// given list of booleans.
//
// Example:
//
//	boolean := NewBoolean(t, true)
//	boolean.NotInList(true, false) // failure
func (b *Boolean) NotInList(values ...bool) *Boolean { _ = "STUB: not implemented"; return nil }
