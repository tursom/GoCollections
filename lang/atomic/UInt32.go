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

type UInt32 uint32

func (v *UInt32) P() *uint32 {
	return (*uint32)(v)
}

func (v *UInt32) Load() (val uint32) {
	return atomic.LoadUint32(v.P())
}

func (v *UInt32) Store(val uint32) {
	atomic.StoreUint32(v.P(), val)
}

func (v *UInt32) Swap(new uint32) (old uint32) {
	return atomic.SwapUint32(v.P(), new)
}

func (v *UInt32) CompareAndSwap(old, new uint32) (swapped bool) {
	return atomic.CompareAndSwapUint32(v.P(), old, new)
}

func (v *UInt32) Add(i uint32) (new uint32) {
	return atomic.AddUint32(v.P(), i)
}

func (v *UInt32) BitLength() int {
	return 32
}

func (v *UInt32) SetBit(bit int, up bool) bool {
	return SetBit(CompareAndSwapUInt32, v.P(), bit, up)
}

func (v *UInt32) CompareAndSwapBit(bit int, old, new bool) bool {
	return CompareAndSwapBit(CompareAndSwapUInt32, v.P(), bit, old, new)
}

func (v *UInt32) String() string {
	return strconv.FormatUint(uint64(*v), 10)
}

func (v *UInt32) AsObject() lang.Object {
	return v
}

func (v *UInt32) Equals(o lang.Object) bool {
	return lang.EqualsUInt32(v, o)
}

func (v *UInt32) HashCode() int32 {
	return int32(*v)
}

func (v *UInt32) ToInt64() lang.Int64 {
	return lang.Int64(*v)
}

func (v *UInt32) ToUInt64() lang.UInt64 {
	return lang.UInt64(*v)
}

func (v *UInt32) ToFloat64() lang.Float64 {
	return lang.Float64(*v)
}
