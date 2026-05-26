package httpexpect

import (
	"regexp"
	"runtime"
)

// Stacktrace entry.
type StacktraceEntry struct {
	Pc uintptr // Program counter

	File string // File path
	Line int    // Line number

	Func *runtime.Func // Function information

	FuncName    string  // Function name (without package and parenthesis)
	FuncPackage string  // Function package
	FuncOffset  uintptr // Program counter offset relative to function start

	// True if this is program entry point
	// (like main.main or testing.tRunner)
	IsEntrypoint bool
}

var stacktraceFuncRe = regexp.MustCompile(`^(.+/[^.]+)\.(.+)$`)

func stacktrace() []StacktraceEntry { _ = "STUB: not implemented"; return nil }
