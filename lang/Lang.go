/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import (
	"reflect"
)

func TypeName[T any]() string {
	t := reflect.TypeOf(Nil[T]())
	return t.Name()
}

func TypeNameOf(v any) string {
	t := reflect.TypeOf(v)
	return t.Name()
}

func Nil[T any]() T {
	var n T
	return n
}

func Len[T any](array []T) int {
	return len(array)
}

func Append[T any](slice []T, elems ...T) []T {
	return append(slice, elems...)
}

func TryCast[T any](v any) (T, bool) {
	if v == nil {
		return Nil[T](), true
	} else {
		t, ok := v.(T)
		return t, ok
	}
}

func Cast[T any](v any) T {
	if v == nil {
		return Nil[T]()
	} else {
		t, ok := v.(T)
		if !ok {
			panic(NewTypeCastException2[T](v, nil))
		}
		return t
	}
}
