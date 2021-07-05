package exceptions

import (
	"fmt"
	"io"
	"os"
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

func PrintStackTraceByArray(writer io.Writer, trace []StackTrace) {
	if trace == nil {
		return
	}
	builder := &strings.Builder{}
	for _, stackTrace := range trace {
		stackTrace.WriteTo(builder)
	}
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

func BuildStackTraceByArray(builder *strings.Builder, trace []StackTrace) {
	if trace == nil {
		return
	}
	for _, stackTrace := range trace {
		stackTrace.WriteTo(builder)
	}
}

func BuildStackTrace(builder *strings.Builder, e Exception, exceptionMsg string) {
	builder.WriteString(fmt.Sprintln(exceptionMsg, e.Error()))
	if e.StackTrace() == nil {
		return
	}
	for _, stackTrace := range e.StackTrace() {
		stackTrace.WriteTo(builder)
	}
}

func Try(
	f func() (ret interface{}, err Exception),
	catch func(panic interface{}) (ret interface{}, err Exception),
) (ret interface{}, err Exception) {
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

func Print(err error) {
	if err == nil {
		return
	}
	switch err.(type) {
	case Exception:
		err.(Exception).PrintStackTrace()
	default:
		_, _ = fmt.Fprintln(os.Stderr, err)
		PrintStackTraceByArray(os.Stderr, GetStackTrace())
	}
}

func Package(err error) Exception {
	if err == nil {
		return nil
	}
	switch err.(type) {
	case Exception:
		return err.(Exception)
	}
	return NewRuntimeException(err, "", true, err)
}

func PackageAny(err interface{}) Exception {
	if err == nil {
		return nil
	}
	switch err.(type) {
	case error:
		return Package(err.(error))
	default:
		return NewRuntimeException(err, "", true, nil)
	}
}

func PackagePanic(panic interface{}) Exception {
	switch panic.(type) {
	case error:
		return NewRuntimeException(
			panic,
			"an panic caused on handle WebSocket message:",
			true, panic,
		)
	default:
		return NewRuntimeException(
			panic,
			"an panic caused on handle WebSocket message:",
			true, nil,
		)
	}
}
