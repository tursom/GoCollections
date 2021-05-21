package exceptions

import (
	"io"
	"os"
)

type RuntimeException struct {
	message          string
	exceptionMessage string
	stackTrace       []StackTrace
}

func NewRuntimeException(message, exceptionMessage string, getStackTrace bool) RuntimeException {
	var stackTrace []StackTrace = nil
	if getStackTrace {
		stackTrace = GetStackTrace()
	}

	if len(exceptionMessage) == 0 {
		exceptionMessage = "exception caused:"
	}

	return RuntimeException{
		message:          message,
		exceptionMessage: exceptionMessage,
		stackTrace:       stackTrace,
	}
}

func (o RuntimeException) Error() string {
	if len(o.message) == 0 {
		return "index out of bound"
	} else {
		return o.message
	}
}

func (o RuntimeException) StackTrace() []StackTrace {
	return o.stackTrace
}

func (o RuntimeException) PrintStackTrace() {
	o.PrintStackTraceTo(os.Stderr)
}

func (o RuntimeException) PrintStackTraceTo(writer io.Writer) {
	PrintStackTrace(writer, o, o.exceptionMessage)
}
