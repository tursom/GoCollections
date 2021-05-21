package exceptions

import (
	"fmt"
	"io"
)

type Exception interface {
	Error() string
	StackTrace() []StackTrace
	PrintStackTrace()
	PrintStackTraceTo(writer io.Writer)
}

func PrintStackTrace(writer io.Writer, e Exception, exceptionMsg string) {
	_, err := fmt.Fprintln(writer, exceptionMsg, e.Error())
	if err != nil {
		return
	}
	if e.StackTrace() == nil {
		return
	}
	for _, stackTrace := range e.StackTrace() {
		stackTrace.PrintLn(writer)
	}
}

func Try(f func() (interface{}, error), catch func(interface{}) interface{}) (ret interface{}) {
	defer func() {
		if r := recover(); r != nil {
			ret = catch(r)
		}
	}()
	ret, err := f()
	if err != nil {
		ret = catch(err)
	}
	return
}
