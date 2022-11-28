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

func (v *Int64) String() string {
	return strconv.FormatInt(int64(*v), 10)
}

func (v *Int64) AsObject() lang.Object {
	return v
}

func (v *Int64) Equals(o lang.Object) bool {
	return lang.EqualsInt64(v, o)
}

func (v *Int64) HashCode() int32 {
	return int32(*v) ^ int32(*v>>32)
}

func (v *Int64) ToInt64() lang.Int64 {
	return lang.Int64(*v)
}

func (v *Int64) ToUInt64() lang.UInt64 {
	return lang.UInt64(*v)
}

func (v *Int64) ToFloat64() lang.Float64 {
	return lang.Float64(*v)
}
