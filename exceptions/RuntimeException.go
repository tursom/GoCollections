package exceptions

import (
	"fmt"
	"github.com/tursom/GoCollections/lang"
	"io"
	"os"
	"strings"
)

type RuntimeException struct {
	lang.BaseObject
	message       string
	exceptionName string
	stackTrace    []StackTrace
	cause         Exception
}

func NewRuntimeException(message string, config *ExceptionConfig) RuntimeException {
	if config == nil {
		config = DefaultExceptionConfig()
	}

	var stackTrace []StackTrace = nil
	if config.GetStackTrace {
		stackTrace = GetStackTraceSkipDeep(config.SkipStack + 1)
	}

	var causeException Exception = nil
	if config.Cause != nil {
		switch e := config.Cause.(type) {
		case Exception:
			causeException = e
		default:
			causeException = NewPackageException(config.Cause, DefaultExceptionConfig().
				SetGetStackTrace(false))
		}
	}

	exceptionName := "github.com.tursom.GoCollections.exceptions.RuntimeException"
	if len(config.ExceptionName) != 0 {
		exceptionName = config.ExceptionName
	}

	return RuntimeException{
		BaseObject:    lang.NewBaseObject(),
		message:       message,
		stackTrace:    stackTrace,
		cause:         causeException,
		exceptionName: exceptionName,
	}
}

func (o RuntimeException) Cause() Exception {
	return o.cause
}

func (o RuntimeException) Error() string {
	message := o.message
	if len(message) == 0 {
		if o.cause != nil {
			message = fmt.Sprintf("%s: %s", o.Name(), o.cause.Error())
		} else {
			message = o.Name()
		}
	} else {
		message = fmt.Sprintf("%s: %s", o.Name(), message)
	}
	return message
}

func (o RuntimeException) Message() string {
	return o.message
}

func (o RuntimeException) Name() string {
	return o.exceptionName
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
			Print(err)
			return
		}
		writeBytes += write
	}
}

func (o RuntimeException) BuildPrintStackTrace(builder *strings.Builder) {
	BuildStackTrace(builder, o)
	if o.cause != nil {
		builder.WriteString("caused by: ")
		o.cause.BuildPrintStackTrace(builder)
	}
}
