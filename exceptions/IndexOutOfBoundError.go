package exceptions

import (
	"fmt"
	"io"
	"os"
)

type IndexOutOfBound struct {
	message    string
	stackTrace []StackTrace
}

func NewIndexOutOfBound(message string, getStackTrace bool) IndexOutOfBound {
	var stackTrace []StackTrace = nil
	if getStackTrace {
		stackTrace = GetStackTrace()
	}
	return IndexOutOfBound{
		message:    message,
		stackTrace: stackTrace,
	}
}

func (i IndexOutOfBound) StackTrace() []StackTrace {
	return i.stackTrace
}

func (i IndexOutOfBound) PrintStackTrace() {
	i.PrintStackTraceTo(os.Stderr)
}

func (i IndexOutOfBound) PrintStackTraceTo(writer io.Writer) {
	_, err := fmt.Fprintln(writer, "exception caused IndexOutOfBound:", i.Error())
	if err != nil {
		return
	}
	if i.stackTrace == nil {
		return
	}
	for _, stackTrace := range i.stackTrace {
		stackTrace.PrintLn(writer)
	}
}

func (i IndexOutOfBound) Error() string {
	if len(i.message) == 0 {
		return "index out of bound"
	} else {
		return i.message
	}
}
