package exceptions

import (
	"fmt"
	"io"
	"strings"
)

type Exception interface {
	Cause() Exception
	Error() string
	StackTrace() []StackTrace
	PrintStackTrace()
	PrintStackTraceTo(writer io.Writer)
	BuildPrintStackTrace(builder *strings.Builder)
}

func PrintStackTrace(builder *strings.Builder, e Exception, exceptionMsg string) {
	builder.WriteString(fmt.Sprintln(exceptionMsg, e.Error()))
	if e.StackTrace() == nil {
		return
	}
	for _, stackTrace := range e.StackTrace() {
		stackTrace.WriteTo(builder)
	}
}

func Try(
	f func() (ret interface{}, err error),
	catch func(panic interface{}) (ret interface{}, err error),
) (ret interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			ret, err = catch(r)
		}
	}()
	ret, err = f()
	if err != nil {
		ret, err = catch(err)
	}
	return
}
