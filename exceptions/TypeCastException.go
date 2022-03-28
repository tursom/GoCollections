package exceptions

import (
	"fmt"
	"github.com/tursom/GoCollections/lang"
	"reflect"
)

type TypeCastException struct {
	RuntimeException
}

func NewTypeCastException(message any, config *ExceptionConfig) *TypeCastException {
	return &TypeCastException{
		NewRuntimeException(
			message,
			"exception caused ElementNotFoundException:",
			config.AddSkipStack(1),
		),
	}
}

func NewTypeCastExceptionByType[T any](obj any, config *ExceptionConfig) *TypeCastException {
	return NewTypeCastException(
		fmt.Sprintf("object %s cannot cast to %s", obj, reflect.TypeOf(lang.Nil[T]()).Name()),
		config,
	)
}
