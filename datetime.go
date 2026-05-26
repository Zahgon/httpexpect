package httpexpect

import (
	"time"
)

// DateTime provides methods to inspect attached time.Time value.
type DateTime struct {
	noCopy noCopy
	chain  *chain
	value  time.Time
}

// NewDateTime returns a new DateTime instance.
//
// If reporter is nil, the function panics.
//
// Example:
//
//	dt := NewDateTime(t, time.Now())
//	dt.IsLe(time.Now())
//
//	time.Sleep(time.Second)
//	dt.IsLt(time.Now())
func NewDateTime(reporter Reporter, value time.Time) *DateTime {
	_ = "STUB: not implemented"
	return nil
}

// NewDateTimeC returns a new DateTime instance with config.
//
// Requirements for config are same as for WithConfig function.
//
// See NewDateTime for usage example.
func NewDateTimeC(config Config, value time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

func newDateTime(parent *chain, val time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

// Raw returns underlying time.Time value attached to DateTime.
// This is the value originally passed to NewDateTime.
//
// Example:
//
//	dt := NewDateTime(t, timestamp)
//	assert.Equal(t, timestamp, dt.Raw())
func (dt *DateTime) Raw() time.Time {
	_ = "STUB: not implemented"

	// Alias is similar to Value.Alias.
	return *new(time.Time)
}

func (dt *DateTime) Alias(name string) *DateTime { _ = "STUB: not implemented"; return nil }

// Zone returns a new String instance with datetime zone.
//
// Example:
//
//	tm, _ := time.Parse(time.UnixDate, "Fri Dec 30 15:04:05 IST 2022")
//	dt := NewDateTime(t, tm)
//	dt.Zone().IsEqual("IST")
func (dt *DateTime) Zone() *String { _ = "STUB: not implemented"; return nil }

// Year returns the year in which datetime occurs,
// in the range [0, 9999]
//
// Example:
//
//	tm, _ := time.Parse(time.UnixDate, "Fri Dec 30 15:04:05 IST 2022")
//	dt := NewDateTime(t, tm)
//	dt.Year().IsEqual(2022)
func (dt *DateTime) Year() *Number { _ = "STUB: not implemented"; return nil }

// Month returns the month of the year specified by datetime,
// in the range [1,12].
//
// Example:
//
//	tm, _ := time.Parse(time.UnixDate, "Fri Dec 30 15:04:05 IST 2022")
//	dt := NewDateTime(t, tm)
//	dt.Month().IsEqual(12)
func (dt *DateTime) Month() *Number { _ = "STUB: not implemented"; return nil }

// Day returns the day of the month specified datetime,
// in the range [1,31].
//
// Example:
//
//	tm, _ := time.Parse(time.UnixDate, "Fri Dec 30 15:04:05 IST 2022")
//	dt := NewDateTime(t, tm)
//	dt.Day().IsEqual(30)
func (dt *DateTime) Day() *Number { _ = "STUB: not implemented"; return nil }

// Weekday returns the day of the week specified by datetime,
// in the range [0, 6], 0 corresponds to Sunday
//
// Example:
//
//	tm, _ := time.Parse(time.UnixDate, "Fri Dec 30 15:04:05 IST 2022")
//	dt := NewDateTime(t, tm)
//	dt.WeekDay().IsEqual(time.Friday)
func (dt *DateTime) WeekDay() *Number { _ = "STUB: not implemented"; return nil }

// YearDay returns the day of the year specified by datetime,
// in the range [1,365] for non-leap years,
// and [1,366] in leap years.
//
// Example:
//
//	tm, _ := time.Parse(time.UnixDate, "Fri Dec 30 15:04:05 IST 2022")
//	dt := NewDateTime(t, tm)
//	dt.YearDay().IsEqual(364)
func (dt *DateTime) YearDay() *Number { _ = "STUB: not implemented"; return nil }

// Hour returns the hour within the day specified by datetime,
// in the range [0, 23].
//
// Example:
//
//	tm, _ := time.Parse(time.UnixDate, "Fri Dec 30 15:04:05 IST 2022")
//	dt := NewDateTime(t, tm)
//	dt.Hour().IsEqual(15)
func (dt *DateTime) Hour() *Number { _ = "STUB: not implemented"; return nil }

// Minute returns the minute offset within the hour specified by datetime,
// in the range [0, 59].
//
// Example:
//
//	tm, _ := time.Parse(time.UnixDate, "Fri Dec 30 15:04:05 IST 2022")
//	dt := NewDateTime(t, tm)
//	dt.Minute().IsEqual(4)
func (dt *DateTime) Minute() *Number { _ = "STUB: not implemented"; return nil }

// Second returns the second offset within the minute specified by datetime,
// in the range [0, 59].
//
// Example:
//
//	tm, _ := time.Parse(time.UnixDate, "Fri Dec 30 15:04:05 IST 2022")
//	dt := NewDateTime(t, tm)
//	dt.Second().IsEqual(5)
func (dt *DateTime) Second() *Number { _ = "STUB: not implemented"; return nil }

// Nanosecond returns the nanosecond offset within the second specified by datetime,
// in the range [0, 999999999].
//
// Example:
//
//	tm, _ := time.Parse(time.UnixDate, "Fri Dec 30 15:04:05 IST 2022")
//	dt := NewDateTime(t, tm)
//	dt.Nanosecond().IsEqual(0)
func (dt *DateTime) Nanosecond() *Number { _ = "STUB: not implemented"; return nil }

// Deprecated: use Zone instead.
func (dt *DateTime) GetZone() *String {
	_ = "STUB: not implemented"

	// Deprecated: use Year instead.
	return nil
}

func (dt *DateTime) GetYear() *Number {
	_ = "STUB: not implemented"

	// Deprecated: use Month instead.
	return nil
}

func (dt *DateTime) GetMonth() *Number {
	_ = "STUB: not implemented"

	// Deprecated: use Day instead.
	return nil
}

func (dt *DateTime) GetDay() *Number {
	_ = "STUB: not implemented"

	// Deprecated: use WeekDay instead.
	return nil
}

func (dt *DateTime) GetWeekDay() *Number { _ = "STUB: not implemented"; return nil }

// Deprecated: use YearDay instead.
func (dt *DateTime) GetYearDay() *Number { _ = "STUB: not implemented"; return nil }

// Deprecated: use Hour instead.
func (dt *DateTime) GetHour() *Number {
	_ = "STUB: not implemented"

	// Deprecated: use Minute instead.
	return nil
}

func (dt *DateTime) GetMinute() *Number {
	_ = "STUB: not implemented"

	// Deprecated: use Second instead.
	return nil
}

func (dt *DateTime) GetSecond() *Number {
	_ = "STUB: not implemented"

	// Deprecated: use Nanosecond instead.
	return nil
}

func (dt *DateTime) GetNanosecond() *Number { _ = "STUB: not implemented"; return nil }

// IsEqual succeeds if DateTime is equal to given value.
//
// Example:
//
//	dt := NewDateTime(t, time.Unix(0, 1))
//	dt.IsEqual(time.Unix(0, 1))
func (dt *DateTime) IsEqual(value time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

// NotEqual succeeds if DateTime is not equal to given value.
//
// Example:
//
//	dt := NewDateTime(t, time.Unix(0, 1))
//	dt.NotEqual(time.Unix(0, 2))
func (dt *DateTime) NotEqual(value time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsEqual instead.
func (dt *DateTime) Equal(value time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

// InRange succeeds if DateTime is within given range [min; max].
//
// Example:
//
//	dt := NewDateTime(t, time.Unix(0, 2))
//	dt.InRange(time.Unix(0, 1), time.Unix(0, 3))
//	dt.InRange(time.Unix(0, 2), time.Unix(0, 2))
func (dt *DateTime) InRange(min, max time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

// NotInRange succeeds if DateTime is not within given range [min; max].
//
// Example:
//
//	dt := NewDateTime(t, time.Unix(0, 10))
//	dt.NotInRange(time.Unix(0, 1), time.Unix(0, 9))
//	dt.NotInRange(time.Unix(0, 11), time.Unix(0, 20))
func (dt *DateTime) NotInRange(min, max time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

// InList succeeds if DateTime is equal to one of the values from given
// list of time.Time.
//
// Example:
//
//	dt := NewDateTime(t, time.Unix(0, 2))
//	dt.InRange(time.Unix(0, 1), time.Unix(0, 2))
func (dt *DateTime) InList(values ...time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

// NotInList succeeds if DateTime is not equal to any of the values from
// given list of time.Time.
//
// Example:
//
//	dt := NewDateTime(t, time.Unix(0, 2))
//	dt.InRange(time.Unix(0, 1), time.Unix(0, 3))
func (dt *DateTime) NotInList(values ...time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

// IsGt succeeds if DateTime is greater than given value.
//
// Example:
//
//	dt := NewDateTime(t, time.Unix(0, 2))
//	dt.IsGt(time.Unix(0, 1))
func (dt *DateTime) IsGt(value time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

// IsGe succeeds if DateTime is greater than or equal to given value.
//
// Example:
//
//	dt := NewDateTime(t, time.Unix(0, 2))
//	dt.IsGe(time.Unix(0, 1))
func (dt *DateTime) IsGe(value time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

// IsLt succeeds if DateTime is lesser than given value.
//
// Example:
//
//	dt := NewDateTime(t, time.Unix(0, 1))
//	dt.IsLt(time.Unix(0, 2))
func (dt *DateTime) IsLt(value time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

// IsLe succeeds if DateTime is lesser than or equal to given value.
//
// Example:
//
//	dt := NewDateTime(t, time.Unix(0, 1))
//	dt.IsLe(time.Unix(0, 2))
func (dt *DateTime) IsLe(value time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsGt instead.
func (dt *DateTime) Gt(value time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsGe instead.
func (dt *DateTime) Ge(value time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsLt instead.
func (dt *DateTime) Lt(value time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

// Deprecated: use IsLe instead.
func (dt *DateTime) Le(value time.Time) *DateTime { _ = "STUB: not implemented"; return nil }

// AsUTC returns a new DateTime instance in UTC timeZone.
//
// Example:
//
//	tm, _ := time.Parse(time.UnixDate, "Fri Dec 30 15:04:05 IST 2022")
//	dt := NewDateTime(t, tm)
//	dt.AsUTC().Zone().IsEqual("UTC")
func (dt *DateTime) AsUTC() *DateTime { _ = "STUB: not implemented"; return nil }

// AsLocal returns a new DateTime instance in Local timeZone.
//
// Example:
//
//	tm, _ := time.Parse(time.UnixDate, "Fri Dec 30 15:04:05 IST 2022")
//	dt := NewDateTime(t, tm)
//	dt.AsLocal().Zone().IsEqual("IST")
func (dt *DateTime) AsLocal() *DateTime { _ = "STUB: not implemented"; return nil }
