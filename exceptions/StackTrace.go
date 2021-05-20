package exceptions

import (
	"fmt"
	"io"
	"runtime"
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

func GetStackTrace() []StackTrace {
	stackTraceMax := 16
	stackTraceUsed := 0
	stackTrace := make([]StackTrace, stackTraceMax)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if stackTraceUsed == stackTraceMax {
			stackTraceMax *= 2
			stackTraceOld := stackTrace
			stackTrace = make([]StackTrace, stackTraceMax)
			copy(stackTrace, stackTraceOld)
		}
		stackTrace[stackTraceUsed] = StackTrace{pc, file, line}
		stackTraceUsed++
	}
	return stackTrace[0:stackTraceUsed]
}
