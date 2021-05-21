package exceptions

import (
	"io"
	"os"
	"strings"
)

type RuntimeException struct {
	message          string
	exceptionMessage string
	stackTrace       []StackTrace
	cause            Exception
}

func NewRuntimeException(message, exceptionMessage string, getStackTrace bool, cause Exception) RuntimeException {
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
		cause:            cause,
	}
}

func (o RuntimeException) Cause() Exception {
	return o.cause
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
	builder := strings.Builder{}
	o.BuildPrintStackTrace(&builder)
	bytes := []byte(builder.String())
	writeBytes := 0
	for writeBytes < len(bytes) {
		write, err := writer.Write(bytes[writeBytes:])
		if err != nil {
			return
		}
		writeBytes += write
	}
}

func (o RuntimeException) BuildPrintStackTrace(builder *strings.Builder) {
	PrintStackTrace(builder, o, o.exceptionMessage)
	if o.cause != nil {
		builder.WriteString("caused by: ")
		o.cause.BuildPrintStackTrace(builder)
	}
}
