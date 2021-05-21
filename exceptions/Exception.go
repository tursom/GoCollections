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

func Try(f func() (ret interface{}, err error), catch func(interface{}) (ret interface{}, err error)) (ret interface{}, err error) {
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
