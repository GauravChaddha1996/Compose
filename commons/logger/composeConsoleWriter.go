package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	colorBlack = iota + 30
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite

	colorBold     = 1
	colorDarkGray = 90
)

var (
	consoleBufPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, 100))
		},
	}
)

const (
	loggerTimeFormat = "02/01/2006 15:04:05.000 -07:00"
)

// Formatter transforms the input into a formatted string.
type Formatter func(interface{}) string

// ComposeConsoleWriter parses the JSON input and writes it in an
// (optionally) colorized, human-friendly format to Out.
type ComposeConsoleWriter struct {
	// Out is the output destination.
	Out io.Writer

	// NoColor disables the colorized output.
	NoColor bool

	// TimeFormat specifies the format for timestamp in output.
	TimeFormat string

	// PartsOrder defines the order of parts in output.
	PartsOrder []string

	// Shall write fields in different lines
	FieldsInDifferentLines bool

	FormatTimestamp     Formatter
	FormatLevel         Formatter
	FormatCaller        Formatter
	FormatMessage       Formatter
	FormatFieldName     Formatter
	FormatFieldValue    Formatter
	FormatErrFieldName  Formatter
	FormatErrFieldValue Formatter
}

// NewComposeConsoleWriter creates and initializes a new ComposeConsoleWriter.
func NewComposeConsoleWriter(options ...func(w *ComposeConsoleWriter)) ComposeConsoleWriter {
	w := ComposeConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: loggerTimeFormat,
		PartsOrder: consoleDefaultPartsOrder(),
	}

	for _, opt := range options {
		opt(&w)
	}

	return w
}

// Write transforms the JSON input with formatters and appends to w.Out.
func (w ComposeConsoleWriter) Write(p []byte) (n int, err error) {
	if w.PartsOrder == nil {
		w.PartsOrder = consoleDefaultPartsOrder()
	}

	var buf = consoleBufPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		consoleBufPool.Put(buf)
	}()

	var evt map[string]interface{}
	d := json.NewDecoder(bytes.NewReader(p))
	d.UseNumber()
	err = d.Decode(&evt)
	if err != nil {
		return n, fmt.Errorf("cannot decode event: %s", err)
	}

	for _, p := range w.PartsOrder {
		w.writePart(buf, evt, p)
	}


	w.writeFields(evt, buf)

	err = buf.WriteByte('\n')
	if err != nil {
		return n, err
	}
	_, err = buf.WriteTo(w.Out)
	return len(p), err
}

// writeFields appends formatted key-value pairs to buf.
func (w ComposeConsoleWriter) writeFields(evt map[string]interface{}, buf *bytes.Buffer) {
	var fields = make([]string, 0, len(evt))
	for field := range evt {
		switch field {
		case zerolog.LevelFieldName, zerolog.TimestampFieldName, zerolog.MessageFieldName, zerolog.CallerFieldName:
			continue
		}
		fields = append(fields, field)
	}
	sort.Strings(fields)

	if len(fields) > 0 {
		buf.WriteByte(' ')
	}

	// Move the "error" field to the front
	ei := sort.Search(len(fields), func(i int) bool { return fields[i] >= zerolog.ErrorFieldName })
	if ei < len(fields) && fields[ei] == zerolog.ErrorFieldName {
		fields[ei] = ""
		fields = append([]string{zerolog.ErrorFieldName}, fields...)
		var xfields = make([]string, 0, len(fields))
		for _, field := range fields {
			if field == "" { // Skip empty fields
				continue
			}
			xfields = append(xfields, field)
		}
		fields = xfields
	}

	if len(fields) > 0 {
		if w.FieldsInDifferentLines {
			buf.WriteString(" --> \n")
		} else {
			buf.WriteString(" --> [")
		}
	}
	for i, field := range fields {
		var fn Formatter
		var fnQuotes Formatter
		var fv Formatter
		var fvQuotes Formatter

		if field == zerolog.ErrorFieldName {
			if w.FormatErrFieldName == nil {
				fn = consoleDefaultFormatErrFieldName(w.NoColor)
			} else {
				fn = w.FormatErrFieldName
			}

			if w.FormatErrFieldValue == nil {
				fv = consoleDefaultFormatErrFieldValue(w.NoColor)
			} else {
				fv = w.FormatErrFieldValue
			}
		} else {
			if w.FormatFieldName == nil {
				fn = consoleDefaultFormatFieldName(w.NoColor)
				fnQuotes = consoleDefaultFormatFieldNameQuotes(w.NoColor)
			} else {
				fn = w.FormatFieldName
				fnQuotes = w.FormatFieldName
			}

			if w.FormatFieldValue == nil {
				fv = consoleDefaultFormatFieldValue
				fvQuotes = consoleDefaultFormatFieldValueQuotes
			} else {
				fv = w.FormatFieldValue
				fvQuotes = w.FormatFieldValue
			}
		}

		if w.FieldsInDifferentLines {
			buf.WriteString(fnQuotes(field))
		} else {
			buf.WriteString(fn(field))
		}

		switch fValue := evt[field].(type) {
		case string:
			if needsQuote(fValue) || w.FieldsInDifferentLines {
				buf.WriteString(fvQuotes(strconv.Quote(fValue)))
			} else {
				buf.WriteString(fv(fValue))
			}
		case json.Number:
			buf.WriteString(fv(fValue))
		default:
			b, err := json.Marshal(fValue)
			if err != nil {
				fmt.Fprintf(buf, colorize("[error: %v]", colorRed, w.NoColor), err)
			} else {
				fmt.Fprint(buf, fv(b))
			}
		}

		if i < len(fields)-1 { // Skip space for last field
			buf.WriteString(",")
			if w.FieldsInDifferentLines {
				buf.WriteString("\n")
			}
		}
	}
	if len(fields) > 0 && !w.FieldsInDifferentLines {
		buf.WriteString("]")
	}
}

// writePart appends a formatted part to buf.
func (w ComposeConsoleWriter) writePart(buf *bytes.Buffer, evt map[string]interface{}, p string) {
	var f Formatter

	switch p {
	case zerolog.LevelFieldName:
		if w.FormatLevel == nil {
			f = consoleDefaultFormatLevel(w.NoColor)
		} else {
			f = w.FormatLevel
		}
	case zerolog.TimestampFieldName:
		if w.FormatTimestamp == nil {
			f = consoleDefaultFormatTimestamp(w.TimeFormat, w.NoColor)
		} else {
			f = w.FormatTimestamp
		}
	case zerolog.MessageFieldName:
		if w.FormatMessage == nil {
			f = consoleDefaultFormatMessage
		} else {
			f = w.FormatMessage
		}
	case zerolog.CallerFieldName:
		if w.FormatCaller == nil {
			f = consoleDefaultFormatCaller(w.NoColor)
		} else {
			f = w.FormatCaller
		}
	default:
		if w.FormatFieldValue == nil {
			f = consoleDefaultFormatFieldValue
		} else {
			f = w.FormatFieldValue
		}
	}

	var s = f(evt[p])

	if len(s) > 0 {
		buf.WriteString(s)
		if p != w.PartsOrder[len(w.PartsOrder)-1] { // Skip space for last part
			buf.WriteByte(' ')
		}
	}
}

// needsQuote returns true when the string s should be quoted in output.
func needsQuote(s string) bool {
	for i := range s {
		if s[i] < 0x20 || s[i] > 0x7e || s[i] == ' ' || s[i] == '\\' || s[i] == '"' {
			return true
		}
	}
	return false
}

// colorize returns the string s wrapped in ANSI code c, unless disabled is true.
func colorize(s interface{}, c int, disabled bool) string {
	if disabled {
		return fmt.Sprintf("%s", s)
	}
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", c, s)
}

// ----- DEFAULT FORMATTERS ---------------------------------------------------

func consoleDefaultPartsOrder() []string {
	return []string{
		zerolog.TimestampFieldName,
		zerolog.LevelFieldName,
		zerolog.CallerFieldName,
		zerolog.MessageFieldName,
	}
}

func consoleDefaultFormatTimestamp(timeFormat string, noColor bool) Formatter {
	if timeFormat == "" {
		timeFormat = loggerTimeFormat
	}
	return func(i interface{}) string {
		t := "<nil>"
		switch tt := i.(type) {
		case string:
			ts, err := time.Parse(zerolog.TimeFieldFormat, tt)
			if err != nil {
				t = tt
			} else {
				t = ts.Format(timeFormat)
			}
		case json.Number:
			i, err := tt.Int64()
			if err != nil {
				t = tt.String()
			} else {
				var sec, nsec int64 = i, 0
				switch zerolog.TimeFieldFormat {
				case zerolog.TimeFormatUnixMs:
					nsec = int64(time.Duration(i) * time.Millisecond)
					sec = 0
				case zerolog.TimeFormatUnixMicro:
					nsec = int64(time.Duration(i) * time.Microsecond)
					sec = 0
				}
				ts := time.Unix(sec, nsec).UTC()
				t = ts.Format(timeFormat)
			}
		}
		return colorize(t, colorBlack, noColor)
	}
}

func consoleDefaultFormatLevel(noColor bool) Formatter {
	return func(i interface{}) string {
		var l string
		if ll, ok := i.(string); ok {
			switch ll {
			case "trace":
				l = colorize("TRACE", colorMagenta, noColor)
			case "debug":
				l = colorize("DEBUG", colorYellow, noColor)
			case "info":
				l = colorize("INFO", colorGreen, noColor)
			case "warn":
				l = colorize("WARN", colorRed, noColor)
			case "error":
				l = colorize(colorize("ERROR", colorRed, noColor), colorBold, noColor)
			case "fatal":
				l = colorize(colorize("FATAL", colorRed, noColor), colorBold, noColor)
			case "panic":
				l = colorize(colorize("PANIC", colorRed, noColor), colorBold, noColor)
			default:
				l = colorize("???", colorBold, noColor)
			}
		} else {
			if i == nil {
				l = colorize("???", colorBold, noColor)
			} else {
				l = strings.ToUpper(fmt.Sprintf("%s", i))[0:3]
			}
		}
		return l + " --> "
	}
}

func consoleDefaultFormatCaller(noColor bool) Formatter {
	return func(i interface{}) string {
		var c string
		if cc, ok := i.(string); ok {
			c = cc
		}
		if len(c) > 0 {
			cwd, err := os.Getwd()
			if err == nil {
				c = strings.TrimPrefix(c, cwd)
				c = strings.TrimPrefix(c, "/")
			}
			c = colorize(c, colorBold, noColor) + colorize(" >", colorCyan, noColor)
		}
		return c
	}
}

func consoleDefaultFormatMessage(i interface{}) string {
	if i == nil {
		return ""
	}
	return fmt.Sprintf("\"%s\"", i)
}

func consoleDefaultFormatFieldName(noColor bool) Formatter {
	return func(i interface{}) string {
		return colorize(fmt.Sprintf("'%s'=", i), colorCyan, noColor)
	}
}

func consoleDefaultFormatFieldNameQuotes(noColor bool) Formatter {
	return func(i interface{}) string {
		return colorize(fmt.Sprintf("%s=", i), colorCyan, noColor)
	}
}

func consoleDefaultFormatFieldValue(i interface{}) string {
	return fmt.Sprintf("'%s'", i)
}

func consoleDefaultFormatFieldValueQuotes(i interface{}) string {
	return fmt.Sprintf("%s", i)
}

func consoleDefaultFormatErrFieldName(noColor bool) Formatter {
	return func(i interface{}) string {
		return colorize(fmt.Sprintf("%s=", i), colorRed, noColor)
	}
}

func consoleDefaultFormatErrFieldValue(noColor bool) Formatter {
	return func(i interface{}) string {
		return colorize(fmt.Sprintf("%s", i), colorRed, noColor)
	}
}
