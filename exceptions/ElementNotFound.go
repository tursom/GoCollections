package exceptions

import (
	"io"
	"os"
)

type ElementNotFoundException struct {
	message    string
	stackTrace []StackTrace
}

func NewElementNotFoundException(message string, getStackTrace bool) *ElementNotFoundException {
	var stackTrace []StackTrace = nil
	if getStackTrace {
		stackTrace = GetStackTrace()
	}
	return &ElementNotFoundException{
		message:    message,
		stackTrace: stackTrace,
	}
}

func (e ElementNotFoundException) Error() string {
	if len(e.message) == 0 {
		return "element not found"
	} else {
		return e.message
	}
}

func (e ElementNotFoundException) StackTrace() []StackTrace {
	return e.stackTrace
}

func (e ElementNotFoundException) PrintStackTrace() {
	e.PrintStackTraceTo(os.Stderr)
}

func (e ElementNotFoundException) PrintStackTraceTo(writer io.Writer) {
	PrintStackTrace(writer, e, "exception caused ElementNotFoundException:")
}
