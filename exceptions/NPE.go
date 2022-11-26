/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package exceptions

import (
	"reflect"

	"github.com/tursom/GoCollections/lang"
)

type NPE struct {
	RuntimeException
}

func NewNPE(message string, config *ExceptionConfig) *NPE {
	return &NPE{
		*NewRuntimeException(message, config.AddSkipStack(1).
			SetExceptionName("github.com.tursom.GoCollections.exceptions.NPE")),
	}
}

func CheckNil[T any](p *T) {
	t := reflect.TypeOf(lang.Nil[T]())
	if p == nil {
		panic(NewNPE(t.Name()+" is null", DefaultExceptionConfig().AddSkipStack(1)))
	}
}
