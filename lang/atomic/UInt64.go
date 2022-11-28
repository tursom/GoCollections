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

func (v *UInt64) String() string {
	return strconv.FormatUint(uint64(*v), 10)
}

func (v *UInt64) AsObject() lang.Object {
	return v
}

func (v *UInt64) Equals(o lang.Object) bool {
	return lang.EqualsUInt64(v, o)
}

func (v *UInt64) HashCode() int32 {
	return int32(*v) ^ int32(*v>>32)
}

func (v *UInt64) ToInt64() lang.Int64 {
	return lang.Int64(*v)
}

func (v *UInt64) ToUInt64() lang.UInt64 {
	return lang.UInt64(*v)
}

func (v *UInt64) ToFloat64() lang.Float64 {
	return lang.Float64(*v)
}
