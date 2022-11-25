/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import (
	"github.com/timandy/routine"
)

//goland:noinspection GoUnusedGlobalVariable
var (
	ThreadLocalGo           = routine.Go
	ThreadLocalGoWait       = routine.GoWait
	ThreadLocalGoWaitResult = routine.GoWaitResult
)

type (
	ThreadLocal[T any] interface {
		Object
		Get() T
		Put(value T)
		Remove()
	}

	threadLocalImpl[T any] struct {
		BaseObject
		threadLocal routine.ThreadLocal
	}
)

func GoId() int64 {
	return routine.Goid()
}

//goland:noinspection GoUnusedExportedFunction
func NewThreadLocal[T any]() ThreadLocal[T] {
	return &threadLocalImpl[T]{
		threadLocal: routine.NewInheritableThreadLocal(),
	}
}

//goland:noinspection GoUnusedExportedFunction
func NewThreadLocalWithInitial[T any](supplier func() T) ThreadLocal[T] {
	return &threadLocalImpl[T]{
		threadLocal: routine.NewThreadLocalWithInitial(func() any {
			return supplier()
		}),
	}
}

//goland:noinspection GoUnusedExportedFunction
func NewInheritableThreadLocal[T any]() ThreadLocal[T] {
	return &threadLocalImpl[T]{
		threadLocal: routine.NewInheritableThreadLocal(),
	}
}

//goland:noinspection GoUnusedExportedFunction
func NewInheritableThreadLocalWithInitial[T any](supplier func() T) ThreadLocal[T] {
	return &threadLocalImpl[T]{
		threadLocal: routine.NewInheritableThreadLocalWithInitial(func() any {
			return supplier()
		}),
	}
}

func (t *threadLocalImpl[T]) Get() T {
	return Cast[T](t.threadLocal.Get())
}

func (t *threadLocalImpl[T]) Put(value T) {
	t.threadLocal.Set(value)
}

func (t *threadLocalImpl[T]) Remove() {
	t.threadLocal.Remove()
}
