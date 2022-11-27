/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "fmt"

type TypeCastException struct {
	RuntimeException
}

func NewTypeCastException(message string, err any, config *ExceptionConfig) *TypeCastException {
	if message == "" {
		message = "type cast failed"
	}

	return &TypeCastException{
		RuntimeException: *NewRuntimeException(message, config.AddSkipStack(1).
			SetCause(err).
			SetExceptionName("github.com.tursom.GoCollections.exceptions.TypeCastException")),
	}
}

func NewTypeCastException2[T any](v any, config *ExceptionConfig) *TypeCastException {
	return NewTypeCastException(fmt.Sprintf("type %s cannot cast to %s", TypeNameOf(v), TypeName[T]()), nil, config)
}
