package exceptions

import (
	"github.com/tursom/GoCollections/lang"
	"reflect"
)

type NPE struct {
	RuntimeException
}

func NewNPE(message any, config *ExceptionConfig) *NPE {
	return &NPE{
		NewRuntimeException(
			message,
			"exception caused NullPointerException:",
			config.AddSkipStack(1),
		),
	}
}

func CheckNil[T any](p *T) {
	t := reflect.TypeOf(lang.Nil[T]())
	if p == nil {
		panic(NewNPE(t.Name()+" is null", DefaultExceptionConfig().AddSkipStack(1)))
	}
}
