/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package atomic

import (
	"sync/atomic"
)

type Int64 int64

func (v *Int64) P() *int64 {
	return (*int64)(v)
}

func (v *Int64) Load() (val int64) {
	return atomic.LoadInt64(v.P())
}

func (v *Int64) Store(val int64) {
	atomic.StoreInt64(v.P(), val)
}

func (v *Int64) Swap(new int64) (old int64) {
	return atomic.SwapInt64(v.P(), new)
}

func (v *Int64) CompareAndSwap(old, new int64) (swapped bool) {
	return atomic.CompareAndSwapInt64(v.P(), old, new)
}

func (v *Int64) Add(i int64) (new int64) {
	return atomic.AddInt64(v.P(), i)
}

func (v *Int64) BitLength() int {
	return 64
}

func (v *Int64) SetBit(bit int, up bool) bool {
	return SetBit(CompareAndSwapInt64, v.P(), bit, up)
}

func (v *Int64) CompareAndSwapBit(bit int, old, new bool) bool {
	return CompareAndSwapBit(CompareAndSwapInt64, v.P(), bit, old, new)
}
