package httpexpect

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertReporter implements Reporter interface using `testify/assert'
// package. Failures are non-fatal with this reporter.
type AssertReporter struct {
	backend *assert.Assertions
}

// NewAssertReporter returns a new AssertReporter object.
func NewAssertReporter(t assert.TestingT) *AssertReporter { _ = "STUB: not implemented"; return nil }

// Errorf implements Reporter.Errorf.
func (r *AssertReporter) Errorf(message string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// RequireReporter implements Reporter interface using `testify/require'
// package. Failures are fatal with this reporter.
type RequireReporter struct {
	backend *require.Assertions
}

// NewRequireReporter returns a new RequireReporter object.
func NewRequireReporter(t require.TestingT) *RequireReporter { _ = "STUB: not implemented"; return nil }

// Errorf implements Reporter.Errorf.
func (r *RequireReporter) Errorf(message string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// FatalReporter is a struct that implements the Reporter interface
// and calls t.Fatalf() when a test fails.
type FatalReporter struct {
	backend testing.TB
}

// NewFatalReporter returns a new FatalReporter object.
func NewFatalReporter(t testing.TB) *FatalReporter { _ = "STUB: not implemented"; return nil }

// Errorf implements Reporter.Errorf.
func (r *FatalReporter) Errorf(message string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// PanicReporter is a struct that implements the Reporter interface
// and panics when a test fails.
// Useful for multithreaded tests when you want to report fatal
// failures from goroutines other than the main goroutine, because
// the main goroutine is forbidden to call t.Fatal.
type PanicReporter struct{}

// NewPanicReporter returns a new PanicReporter object.
func NewPanicReporter() *PanicReporter { _ = "STUB: not implemented"; return nil }

// Errorf implements Reporter.Errorf
func (r *PanicReporter) Errorf(message string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}
