package exceptions

import (
	"testing"
)

func TestRuntimeException_PrintStackTrace(t *testing.T) {
	exception := NewRuntimeException("test1", DefaultExceptionConfig().SetCause(1))
	exception = NewRuntimeException("test2", DefaultExceptionConfig().SetCause(exception))
	exception = NewRuntimeException("", DefaultExceptionConfig().SetCause(exception))
	exception.PrintStackTrace()
}
