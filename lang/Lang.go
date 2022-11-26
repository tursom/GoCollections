/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unsafe"

	"github.com/tursom/GoCollections/exceptions"
)

type TypeCastException struct {
	BaseObject
	message    string
	stackTrace []StackTrace
	cause      any
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
	defer func() {
		r := recover()
		if r != nil {
			panic(&TypeCastException{
				message:    "cast failed",
				stackTrace: GetStackTraceSkipDeep(1),
				cause:      r,
			})
		}
	}()

	if v == nil {
		return Nil[T]()
	} else {
		return v.(T)
	}
}

func ForceCast[T any](v unsafe.Pointer) *T {
	defer func() {
		r := recover()
		if r != nil {
			panic(&TypeCastException{
				message:    "cast failed",
				stackTrace: GetStackTraceSkipDeep(1),
				cause:      r,
			})
		}
	}()

	if v == nil {
		return nil
	} else {
		return (*T)(v)
	}
}

func (t *TypeCastException) Cause() exceptions.Exception {
	return nil
}

func (t *TypeCastException) Error() string {
	return fmt.Sprintf("TypeCastException: %s\ncause by: %s", t.message, t.cause)
}

func (t *TypeCastException) Name() string {
	return "github.com.tursom.GoCollections.lang.TypeCastException"
}

func (t *TypeCastException) Message() string {
	return t.message
}

func (t *TypeCastException) StackTrace() []StackTrace {
	return t.stackTrace
}

func (t *TypeCastException) PrintStackTrace() {
	t.PrintStackTraceTo(os.Stderr)
}

func (t *TypeCastException) PrintStackTraceTo(writer io.Writer) {
	builder := strings.Builder{}
	t.BuildPrintStackTrace(&builder)
	bytes := []byte(builder.String())
	writeBytes := 0
	_, _ = writer.Write(bytes[writeBytes:])
}

func (t *TypeCastException) BuildPrintStackTrace(builder *strings.Builder) {
	builder.WriteString(t.Error())
	builder.WriteString("\n")
	if t.StackTrace() == nil {
		return
	}
	for _, stackTrace := range t.StackTrace() {
		stackTrace.WriteTo(builder)
	}
}
