/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package atomic

import (
	"sync/atomic"
)

type UInt64 uint64

func (v *UInt64) P() *uint64 {
	return (*uint64)(v)
}

func (v *UInt64) Load() (val uint64) {
	return atomic.LoadUint64(v.P())
}

func (v *UInt64) Store(val uint64) {
	atomic.StoreUint64(v.P(), val)
}

func (v *UInt64) Swap(new uint64) (old uint64) {
	return atomic.SwapUint64(v.P(), new)
}

func (v *UInt64) CompareAndSwap(old, new uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64(v.P(), old, new)
}

func (v *UInt64) Add(i uint64) (new uint64) {
	return atomic.AddUint64(v.P(), i)
}

func (v *UInt64) BitLength() int {
	return 64
}

func (v *UInt64) SetBit(bit int, up bool) bool {
	return SetBit(CompareAndSwapUInt64, v.P(), bit, up)
}

func (v *UInt64) CompareAndSwapBit(bit int, old, new bool) bool {
	return CompareAndSwapBit(CompareAndSwapUInt64, v.P(), bit, old, new)
}
