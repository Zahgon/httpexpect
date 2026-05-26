package httpexpect

import (
	"encoding/json"
	"regexp"
	"strings"
	"sync"
	"text/template"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
	"github.com/mitchellh/go-wordwrap"
)

// Formatter is used to format assertion messages into strings.
type Formatter interface {
	FormatSuccess(*AssertionContext) string
	FormatFailure(*AssertionContext, *AssertionFailure) string
}

// DefaultFormatter is the default Formatter implementation.
//
// DefaultFormatter gathers values from AssertionContext and AssertionFailure,
// converts them to strings, and creates FormatData struct. Then it passes
// FormatData to the template engine (text/template) to format message.
//
// You can control what is included and what is excluded from messages via
// several public fields.
//
// If desired, you can provide custom templates and function map. This may
// be easier than creating your own formatter from scratch.
type DefaultFormatter struct {
	// Exclude test name and request name from failure report.
	DisableNames bool

	// Exclude assertion path from failure report.
	DisablePaths bool

	// Exclude aliased assertion path from failure report.
	DisableAliases bool

	// Exclude diff from failure report.
	DisableDiffs bool

	// Exclude HTTP request from failure report.
	DisableRequests bool

	// Exclude HTTP response from failure report.
	DisableResponses bool

	// Thousand separator.
	// Default is DigitSeparatorUnderscore.
	DigitSeparator DigitSeparator

	// Float printing format.
	// Default is FloatFormatAuto.
	FloatFormat FloatFormat

	// Defines whether to print stacktrace on failure and in what format.
	// Default is StacktraceModeDisabled.
	StacktraceMode StacktraceMode

	// Colorization mode.
	// Default is ColorModeAuto.
	ColorMode ColorMode

	// Wrap text to keep lines below given width.
	// Use zero for default width, and negative value to disable wrapping.
	LineWidth int

	// If not empty, used to format success messages.
	// If empty, default template is used.
	SuccessTemplate string

	// If not empty, used to format failure messages.
	// If empty, default template is used.
	FailureTemplate string

	// When SuccessTemplate or FailureTemplate is set, this field
	// defines the function map passed to template engine.
	// May be nil.
	TemplateFuncs template.FuncMap
}

// FormatSuccess implements Formatter.FormatSuccess.
func (f *DefaultFormatter) FormatSuccess(ctx *AssertionContext) string {
	_ = "STUB: not implemented"
	return ""
}

// FormatFailure implements Formatter.FormatFailure.
func (f *DefaultFormatter) FormatFailure(
	ctx *AssertionContext, failure *AssertionFailure,
) string {
	_ = "STUB: not implemented"
	return ""
}

// DigitSeparator defines the separator used to format integers and floats.
type DigitSeparator int

const (
	// Separate using underscore
	DigitSeparatorUnderscore DigitSeparator = iota

	// Separate using comma
	DigitSeparatorComma

	// Separate using apostrophe
	DigitSeparatorApostrophe

	// Do not separate
	DigitSeparatorNone
)

// FloatFormat defines the format in which all floats are printed.
type FloatFormat int

const (
	// Print floats in scientific notation for large exponents,
	// otherwise print in decimal notation.
	// Precision is the smallest needed to identify the value uniquely.
	// Similar to %g format.
	FloatFormatAuto FloatFormat = iota

	// Always print floats in decimal notation.
	// Precision is the smallest needed to identify the value uniquely.
	// Similar to %f format.
	FloatFormatDecimal

	// Always print floats in scientific notation.
	// Precision is the smallest needed to identify the value uniquely.
	// Similar to %e format.
	FloatFormatScientific
)

// StacktraceMode defines the format of stacktrace.
type StacktraceMode int

const (
	// Don't print stacktrace.
	StacktraceModeDisabled StacktraceMode = iota

	// Standard, verbose format.
	StacktraceModeStandard

	// Compact format.
	StacktraceModeCompact
)

// ColorMode defines how the text color is enabled.
type ColorMode int

const (
	// Automatically enable colors if ALL of the following is true:
	//  - stdout is a tty / console
	//  - AssertionHandler is known to output to testing.T
	//  - testing.Verbose() is true
	//
	// Colors are forcibly enabled if FORCE_COLOR environment variable
	// is set to a positive integer.
	//
	// Colors are forcibly disabled if TERM is "dumb" or NO_COLOR
	// environment variable is set to non-empty string.
	ColorModeAuto ColorMode = iota

	// Unconditionally enable colors.
	ColorModeAlways

	// Unconditionally disable colors.
	ColorModeNever
)

// FormatData defines data passed to template engine when DefaultFormatter
// formats assertion. You can use these fields in your custom templates.
type FormatData struct {
	TestName    string
	RequestName string

	AssertPath     []string
	AssertType     string
	AssertSeverity string

	Errors []string

	HaveActual bool
	Actual     string

	HaveExpected bool
	IsNegation   bool
	IsComparison bool
	ExpectedKind string
	Expected     []string

	HaveReference bool
	Reference     string

	HaveDelta bool
	Delta     string

	HaveDiff bool
	Diff     string

	HaveRequest bool
	Request     string

	HaveResponse bool
	Response     string

	HaveStacktrace bool
	Stacktrace     []string

	EnableColors bool
	LineWidth    int
}

const (
	kindRange      = "range"
	kindSchema     = "schema"
	kindPath       = "path"
	kindRegexp     = "regexp"
	kindFormat     = "format"
	kindFormatList = "formats"
	kindKey        = "key"
	kindElement    = "element"
	kindSubset     = "subset"
	kindValue      = "value"
	kindValueList  = "values"
)

func (f *DefaultFormatter) applyTemplate(
	templateName string,
	templateString string,
	templateFuncs template.FuncMap,
	ctx *AssertionContext,
	failure *AssertionFailure,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *DefaultFormatter) buildFormatData(
	ctx *AssertionContext, failure *AssertionFailure,
) *FormatData {
	_ = "STUB: not implemented"
	return nil
}

func (f *DefaultFormatter) fillGeneral(
	data *FormatData, ctx *AssertionContext,
) {
	_ = "STUB: not implemented"
	return
}

func (f *DefaultFormatter) fillErrors(
	data *FormatData, ctx *AssertionContext, failure *AssertionFailure,
) {
	_ = "STUB: not implemented"
	return
}

func (f *DefaultFormatter) fillActual(
	data *FormatData, ctx *AssertionContext, failure *AssertionFailure,
) {
	_ = "STUB: not implemented"
	return
	//nolint
}

func (f *DefaultFormatter) fillExpected(
	data *FormatData, ctx *AssertionContext, failure *AssertionFailure,
) {
	_ = "STUB: not implemented"
	return
}

func (f *DefaultFormatter) fillIsNegation(
	data *FormatData, ctx *AssertionContext, failure *AssertionFailure,
) {
	_ = "STUB: not implemented"
	return
}

func (f *DefaultFormatter) fillIsComparison(
	data *FormatData, ctx *AssertionContext, failure *AssertionFailure,
) {
	_ = "STUB: not implemented"
	return
	//nolint
}

func (f *DefaultFormatter) fillReference(
	data *FormatData, ctx *AssertionContext, failure *AssertionFailure,
) {
	_ = "STUB: not implemented"
	return
}

func (f *DefaultFormatter) fillDelta(
	data *FormatData, ctx *AssertionContext, failure *AssertionFailure,
) {
	_ = "STUB: not implemented"
	return
}

func (f *DefaultFormatter) fillRequest(
	data *FormatData, ctx *AssertionContext, failure *AssertionFailure,
) {
	_ = "STUB: not implemented"
	return
}

func (f *DefaultFormatter) fillResponse(
	data *FormatData, ctx *AssertionContext, failure *AssertionFailure,
) {
	_ = "STUB: not implemented"
	return
}

func (f *DefaultFormatter) fillStacktrace(
	data *FormatData, ctx *AssertionContext, failure *AssertionFailure,
) {
	_ = "STUB: not implemented"
	return
}

func (f *DefaultFormatter) formatValue(value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *DefaultFormatter) formatFloatValue(value float64, bits int) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *DefaultFormatter) formatTypedValue(value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *DefaultFormatter) formatMatchValue(value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *DefaultFormatter) formatRangeValue(value interface{}) []string {
	_ = "STUB: not implemented"
	return nil
}

func (f *DefaultFormatter) formatListValue(value interface{}) []string {
	_ = "STUB: not implemented"
	return nil
}

func (f *DefaultFormatter) formatDiff(expected, actual interface{}) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (f *DefaultFormatter) reformatNumber(numStr string) string {
	_ = "STUB: not implemented"
	return ""
}

var (
	decomposeRegexp = regexp.MustCompile(`^([+-])?(\d+)([.](\d+))?([eE]([+-]?\d+))?$`)
)

func (f *DefaultFormatter) decomposeNumber(numStr string) (
	signPart, intPart, fracPart, expPart string,
) {
	_ = "STUB: not implemented"
	return "", "", "", ""
}

func (f *DefaultFormatter) applySeparator(numStr string, dir int) string {
	_ = "STUB: not implemented"
	return ""
}

func extractString(value interface{}) *string { _ = "STUB: not implemented"; return nil }

func extractFloat32(value interface{}) *float64 { _ = "STUB: not implemented"; return nil }

func extractFloat64(value interface{}) *float64 { _ = "STUB: not implemented"; return nil }

func exctractRange(value interface{}) *AssertionRange { _ = "STUB: not implemented"; return nil }

// invalid, but we handle it

func extractList(value interface{}) *AssertionList { _ = "STUB: not implemented"; return nil }

// invalid, but we handle it

var (
	colorsSupportedOnce sync.Once
	colorsSupportedMode int
)

const (
	colorsUnsupported = iota
	colorsSupported
	colorsForced
)

func colorMode() int { _ = "STUB: not implemented"; return 0 }

const (
	defaultIndent    = "  "
	defaultLineWidth = 60
)

var defaultColors = map[string]color.Attribute{
	// regular
	"Black":   color.FgBlack,
	"Red":     color.FgRed,
	"Green":   color.FgGreen,
	"Yellow":  color.FgYellow,
	"Magenta": color.FgMagenta,
	"Cyan":    color.FgCyan,
	"White":   color.FgWhite,
	// bright
	"HiBlack":   color.FgHiBlack,
	"HiRed":     color.FgHiRed,
	"HiGreen":   color.FgHiGreen,
	"HiYellow":  color.FgHiYellow,
	"HiMagenta": color.FgHiMagenta,
	"HiCyan":    color.FgHiCyan,
	"HiWhite":   color.FgHiWhite,
}

var defaultTemplateFuncs = template.FuncMap{
	"trim": func(input string) string {
		return strings.TrimSpace(input)
	},
	"indent": func(input string) string {
		var sb strings.Builder

		for _, s := range strings.Split(input, "\n") {
			if sb.Len() != 0 {
				sb.WriteString("\n")
			}
			sb.WriteString(defaultIndent)
			sb.WriteString(s)
		}

		return sb.String()
	},
	"wrap": func(width int, input string) string {
		input = strings.TrimSpace(input)

		width -= len(defaultIndent)
		if width <= 0 {
			return input
		}

		return wordwrap.WrapString(input, uint(width))
	},
	"join": func(width int, tokenList []string) string {
		width -= len(defaultIndent)
		if width <= 0 {
			return strings.Join(tokenList, ".")
		}

		var sb strings.Builder

		lineLen := 0
		lineNum := 0

		write := func(s string) {
			sb.WriteString(s)
			lineLen += len(s)
		}

		for n, token := range tokenList {
			if lineLen+len(token)+1 > width {
				write("\n")
				lineLen = 0
				if lineNum < 2 {
					lineNum++
				}
			}
			if lineLen == 0 {
				for l := 0; l < lineNum; l++ {
					write(defaultIndent)
				}
			}
			write(token)
			if n != len(tokenList)-1 {
				write(".")
			}
		}

		return sb.String()
	},
	"color": func(enable bool, colorName, input string) string {
		if !enable {
			return input
		}
		colorAttr := color.Reset
		if ca, ok := defaultColors[colorName]; ok {
			colorAttr = ca
		}
		return color.New(colorAttr).Sprint(input)
	},
	"colorhttp": func(enable bool, isResponse bool, input string) string {
		if !enable {
			return input
		}

		methodColor := color.New(defaultColors["HiMagenta"])
		statusColor := color.New(defaultColors["HiMagenta"])
		headerColor := color.New(defaultColors["Cyan"])

		var sb strings.Builder

		isFirstLine := true
		for _, line := range strings.Split(input, "\n") {
			if sb.Len() != 0 {
				sb.WriteString("\n")
			}

			line = strings.TrimSuffix(line, "\n")
			line = strings.TrimSuffix(line, "\r")

			var words []string
			if isFirstLine {
				words = strings.SplitN(line, " ", -1)
			} else {
				words = strings.SplitN(line, " ", 2)
			}

			wordLen := len(words)
			for index, word := range words {
				var applyColor *color.Color

				if isFirstLine {
					if isResponse {
						if index != 0 && index != wordLen-1 {
							applyColor = statusColor
						}
					} else {
						if index == 0 {
							applyColor = methodColor
						}
					}
				} else {
					if index == 0 {
						applyColor = headerColor
					}
				}

				if word != "" && applyColor != nil {
					sb.WriteString(applyColor.Sprint(word))
				} else {
					sb.WriteString(word)
				}

				sb.WriteString(" ")
			}

			isFirstLine = false
		}

		return sb.String()
	},
	"colorjson": func(enable bool, colorName, input string) string {
		if !enable {
			return input
		}

		fallbackColor := color.Reset
		if attr, ok := defaultColors[colorName]; ok {
			fallbackColor = attr
		}

		var parsedInput interface{}
		err := json.Unmarshal([]byte(input), &parsedInput)
		if err != nil {
			return color.New(fallbackColor).Sprint(input)
		}

		formatter := colorjson.NewFormatter()
		formatter.KeyColor = color.New(color.Reset)
		formatter.StringColor = color.New(defaultColors["HiMagenta"])
		formatter.NumberColor = color.New(defaultColors["Cyan"])
		formatter.BoolColor = color.New(defaultColors["Cyan"])
		formatter.NullColor = color.New(defaultColors["Cyan"])
		formatter.Indent = 2

		b, err := formatter.Marshal(parsedInput)
		if err != nil {
			return color.New(fallbackColor).Sprint(input)
		}

		return string(b)
	},
	"colordiff": func(enable bool, input string) string {
		if !enable {
			return input
		}

		prefixColor := [...]struct {
			prefix string
			color  color.Attribute
		}{
			{"---", color.FgWhite},
			{"+++", color.FgWhite},
			{"-", color.FgRed},
			{"+", color.FgGreen},
		}

		lineColor := func(s string) color.Attribute {
			for _, pc := range prefixColor {
				if strings.HasPrefix(s, pc.prefix) {
					return pc.color
				}
			}

			return color.Reset
		}

		var sb strings.Builder
		for _, line := range strings.Split(input, "\n") {
			if sb.Len() != 0 {
				sb.WriteString("\n")
			}

			sb.WriteString(color.New(lineColor(line)).Sprint(line))
		}

		return sb.String()
	},
}

var defaultSuccessTemplate = `[OK] {{ join .LineWidth .AssertPath }}`

var defaultFailureTemplate = `
{{- range $n, $err := .Errors }}
{{ if eq $n 0 -}}
{{ $err | wrap $.LineWidth | color $.EnableColors "Red" }}
{{- else -}}
{{ $err | wrap $.LineWidth | indent | color $.EnableColors "Red" }}
{{- end -}}
{{- end -}}
{{- if .TestName }}

test name: {{ .TestName | color $.EnableColors "Cyan" }}
{{- end -}}
{{- if .RequestName }}

request name: {{ .RequestName | color $.EnableColors "Cyan" }}
{{- end -}}
{{- if .HaveRequest }}

request: {{ .Request | colorhttp $.EnableColors false | indent | trim }}
{{- end -}}
{{- if .HaveResponse }}

response: {{ .Response | colorhttp $.EnableColors true | indent | trim }}
{{- end -}}
{{- if .HaveStacktrace }}

trace:
{{- range $n, $call := .Stacktrace }}
{{ $call | indent }}
{{- end -}}
{{- end -}}
{{- if .AssertPath }}

assertion:
{{ join .LineWidth .AssertPath | indent | color .EnableColors "Yellow" }}
{{- end -}}
{{- if .HaveExpected }}

{{ if .IsNegation }}denied
{{- else if .IsComparison }}compared
{{- else }}expected
{{- end }} {{ .ExpectedKind }}:
{{- range $n, $exp := .Expected }}
{{ $exp | colorjson $.EnableColors "HiMagenta" | indent }}
{{- end -}}
{{- end -}}
{{- if .HaveActual }}

actual value:
{{ .Actual | colorjson .EnableColors "HiMagenta" | indent }}
{{- end -}}
{{- if .HaveReference }}

reference value:
{{ .Reference | colorjson .EnableColors "HiMagenta" | indent }}
{{- end -}}
{{- if .HaveDelta }}

allowed delta:
{{ .Delta | indent | color .EnableColors "Cyan" }}
{{- end -}}
{{- if .HaveDiff }}

diff:
{{ .Diff | colordiff .EnableColors | indent }}
{{- end -}}
`
