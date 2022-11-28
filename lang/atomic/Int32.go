/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package atomic

import (
	"strconv"
	"sync/atomic"

	"github.com/tursom/GoCollections/lang"
)

type Int32 int32

func (v *Int32) P() *int32 {
	return (*int32)(v)
}

func (v *Int32) Load() (val int32) {
	return atomic.LoadInt32(v.P())
}

func (v *Int32) Store(val int32) {
	atomic.StoreInt32(v.P(), val)
}

func (v *Int32) Swap(new int32) (old int32) {
	return atomic.SwapInt32(v.P(), new)
}

func (v *Int32) CompareAndSwap(old, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(v.P(), old, new)
}

func (v *Int32) Add(i int32) (new int32) {
	return atomic.AddInt32(v.P(), i)
}

func (v *Int32) BitLength() int {
	return 32
}

func (v *Int32) SetBit(bit int, up bool) bool {
	return SetBit(CompareAndSwapInt32, v.P(), bit, up)
}

func (v *Int32) CompareAndSwapBit(bit int, old, new bool) bool {
	return CompareAndSwapBit(CompareAndSwapInt32, v.P(), bit, old, new)
}

func (v *Int32) String() string {
	return strconv.FormatInt(int64(*v), 10)
}

func (v *Int32) AsObject() lang.Object {
	return v
}

func (v *Int32) Equals(o lang.Object) bool {
	return lang.EqualsInt32(v, o)
}

func (v *Int32) HashCode() int32 {
	return lang.Hash32(v)
}

func (v *Int32) ToInt64() lang.Int64 {
	return lang.Int64(*v)
}

func (v *Int32) ToUInt64() lang.UInt64 {
	return lang.UInt64(*v)
}

func (v *Int32) ToFloat64() lang.Float64 {
	return lang.Float64(*v)
}
