/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package util

import (
	"sync"

	"github.com/tursom/GoCollections/lang"
)

var (
	goroutineLocalMap = sync.Map{}
)

type (
	// GoroutineLocal goroutine local variable
	GoroutineLocal[T any] struct {
		lang.BaseObject
	}
)

func WithGoroutineLocal(f func()) {
	initGoroutineLocal()
	defer freeGoroutineLocal()

	f()
}

func (g *GoroutineLocal[T]) Get() T {
	//TODO implement me
	panic("implement me")
}

func (g *GoroutineLocal[T]) Put(value T) {
	//TODO implement me
	panic("implement me")
}

func (g *GoroutineLocal[T]) Remove() {
	//TODO implement me
	panic("implement me")
}

func initGoroutineLocal() {
	//TODO implement me
	goroutineLocalMap.Store(lang.GoId(), nil)
	panic("implement me")
}

func freeGoroutineLocal() {
	goroutineLocalMap.Delete(lang.GoId())
}
