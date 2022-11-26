/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package exceptions

import (
	"fmt"
	"reflect"

	"github.com/tursom/GoCollections/lang"
)

type TypeCastException = lang.TypeCastException

func NewTypeCastException(message string, config *ExceptionConfig) *TypeCastException {
	return &TypeCastException{
		*NewRuntimeException(message, config.AddSkipStack(1).
			SetExceptionName("github.com.tursom.GoCollections.exceptions.TypeCastException")),
	}
}

func NewTypeCastExceptionByType[T any](obj any, config *ExceptionConfig) *TypeCastException {
	return NewTypeCastException(
		fmt.Sprintf("object %s cannot cast to %s", obj, reflect.TypeOf(lang.Nil[T]()).Name()),
		config,
	)
}
