package exceptions

import (
	"fmt"
	"io"
	"runtime"
	"strings"
)

type StackTrace struct {
	pc   uintptr
	file string
	line int
}

func (s StackTrace) Pc() uintptr {
	return s.pc
}

func (s StackTrace) File() string {
	return s.file
}

func (s StackTrace) Line() int {
	return s.line
}

func (s StackTrace) Print(writer io.Writer) {
	_, err := fmt.Fprint(writer, "at", s.file, s.line)
	if err != nil {
		return
	}
}

func (s StackTrace) PrintLn(writer io.Writer) {
	_, err := fmt.Fprintf(writer, "\tat %s(%d)\n", s.file, s.line)
	if err != nil {
		return
	}
}

func (s StackTrace) WriteTo(builder *strings.Builder) {
	builder.WriteString(fmt.Sprintf("\tat %s(%d)\n", s.file, s.line))
}

func GetStackTrace() []StackTrace {
	return GetStackTraceSkipDeep(0)
}

func GetStackTraceSkipDeep(deep int) []StackTrace {
	stackTrace := make([]StackTrace, 0, 16)
	for i := deep + 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		stackTrace = append(stackTrace, StackTrace{pc, file, line})
	}
	return stackTrace
}
