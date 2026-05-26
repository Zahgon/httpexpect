package httpexpect

import (
	"time"
)

// Duration provides methods to inspect attached time.Duration value.
type Duration struct {
	noCopy noCopy
	chain  *chain
	value  *time.Duration
}

// NewDuration returns a new Duration instance.
//
// If reporter is nil, the function panics.
//
// Example:
//
//	d := NewDuration(t, time.Second)
//	d.IsLe(time.Minute)
func NewDuration(reporter Reporter, value time.Duration) *Duration {
	_ = "STUB: not implemented"
	return nil
}

// NewDurationC returns a new Duration instance with config.
//
// Requirements for config are same as for WithConfig function.
//
// Example:
//
//	d := NewDurationC(config, time.Second)
//	d.IsLe(time.Minute)
func NewDurationC(config Config, value time.Duration) *Duration {
	_ = "STUB: not implemented"
	return nil
}

func newDuration(parent *chain, val *time.Duration) *Duration {
	_ = "STUB: not implemented"
	return nil
}

// Raw returns underlying time.Duration value attached to Duration.
// This is the value originally passed to NewDuration.
//
// Example:
//
//	d := NewDuration(t, duration)
//	assert.Equal(t, timestamp, d.Raw())
func (d *Duration) Raw() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// Alias is similar to Value.Alias.
func (d *Duration) Alias(name string) *Duration { _ = "STUB: not implemented"; return nil }

// Deprecated: support for unset durations will be removed. The only method that
// can create unset duration is Cookie.MaxAge. Instead of Cookie.MaxAge().IsSet(),
// please use Cookie.ContainsMaxAge().
func (d *Duration) IsSet() *Duration { _ = "STUB: not implemented"; return nil }

// Deprecated: support for unset durations will be removed. The only method that
// can create unset duration is Cookie.MaxAge. Instead of Cookie.MaxAge().NotSet(),
// please use Cookie.NotContainsMaxAge().
func (d *Duration) NotSet() *Duration { _ = "STUB: not implemented"; return nil }

// IsEqual succeeds if Duration is equal to given value.
//
// Example:
//
//	d := NewDuration(t, time.Second)
//	d.IsEqual(time.Second)
func (d *Duration) IsEqual(value time.Duration) *Duration { _ = "STUB: not implemented"; return nil }

// NotEqual succeeds if Duration is not equal to given value.
//
// Example:
//
//	d := NewDuration(t, time.Second)
//	d.NotEqual(time.Minute)
func (d *Duration) NotEqual(value time.Duration) *Duration { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsEqual instead.
func (d *Duration) Equal(value time.Duration) *Duration { _ = "STUB: not implemented"; return nil }

// IsGt succeeds if Duration is greater than given value.
//
// Example:
//
//	d := NewDuration(t, time.Minute)
//	d.IsGt(time.Second)
func (d *Duration) IsGt(value time.Duration) *Duration { _ = "STUB: not implemented"; return nil }

// IsGe succeeds if Duration is greater than or equal to given value.
//
// Example:
//
//	d := NewDuration(t, time.Minute)
//	d.IsGe(time.Second)
func (d *Duration) IsGe(value time.Duration) *Duration { _ = "STUB: not implemented"; return nil }

// IsLt succeeds if Duration is lesser than given value.
//
// Example:
//
//	d := NewDuration(t, time.Second)
//	d.IsLt(time.Minute)
func (d *Duration) IsLt(value time.Duration) *Duration { _ = "STUB: not implemented"; return nil }

// IsLe succeeds if Duration is lesser than or equal to given value.
//
// Example:
//
//	d := NewDuration(t, time.Second)
//	d.IsLe(time.Minute)
func (d *Duration) IsLe(value time.Duration) *Duration { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsGt instead.
func (d *Duration) Gt(value time.Duration) *Duration { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsGe instead.
func (d *Duration) Ge(value time.Duration) *Duration { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsLt instead.
func (d *Duration) Lt(value time.Duration) *Duration { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsLe instead.
func (d *Duration) Le(value time.Duration) *Duration { _ = "STUB: not implemented"; return nil }

// InRange succeeds if Duration is within given range [min; max].
//
// Example:
//
//	d := NewDuration(t, time.Minute)
//	d.InRange(time.Second, time.Hour)
//	d.InRange(time.Minute, time.Minute)
func (d *Duration) InRange(min, max time.Duration) *Duration { _ = "STUB: not implemented"; return nil }

// NotInRange succeeds if Duration is not within given range [min; max].
//
// Example:
//
//	d := NewDuration(t, time.Minute*10)
//	d.NotInRange(time.Minute, time.Minute-time.Nanosecond)
//	d.NotInRange(time.Minute+time.Nanosecond, time.Minute*10)
func (d *Duration) NotInRange(min, max time.Duration) *Duration {
	_ = "STUB: not implemented"
	return nil
}

// InList succeeds if Duration is equal to one of the values from given
// list of time.Duration.
//
// Example:
//
//	d := NewDuration(t, time.Minute)
//	d.InList(time.Minute, time.Hour)
func (d *Duration) InList(values ...time.Duration) *Duration { _ = "STUB: not implemented"; return nil }

// NotInList succeeds if Duration is not equal to any of the values from
// given list of time.Duration.
//
// Example:
//
//	d := NewDuration(t, time.Minute)
//	d.NotInList(time.Second, time.Hour)
func (d *Duration) NotInList(values ...time.Duration) *Duration {
	_ = "STUB: not implemented"
	return nil
}
