/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package atomic

import (
	"strconv"
	"sync/atomic"
	"unsafe"

	"github.com/tursom/GoCollections/lang"
)

type Float64 float64

func (v *Float64) P() *float64 {
	return (*float64)(v)
}

func (v *Float64) pint64() *int64 {
	return (*int64)(unsafe.Pointer(v))
}

func (v *Float64) Load() (val float64) {
	return intToFloat64(atomic.LoadInt64(v.pint64()))
}

func (v *Float64) Store(val float64) {
	atomic.StoreInt64(v.pint64(), floatToInt64(val))
}

func (v *Float64) Swap(new float64) (old float64) {
	i := atomic.SwapInt64(v.pint64(), floatToInt64(new))
	return intToFloat64(i)
}

func (v *Float64) CompareAndSwap(old, new float64) (swapped bool) {
	return atomic.CompareAndSwapInt64(v.pint64(), floatToInt64(old), floatToInt64(new))
}

func (v *Float64) Add(f float64) (new float64) {
	old := float64(*v)
	new = old + f
	for !v.CompareAndSwap(old, new) {
		old = float64(*v)
		new = old + f
	}

	return new
}

func (v *Float64) String() string {
	return strconv.FormatFloat(float64(*v), 'e', -1, 64)
}

func (v *Float64) AsObject() lang.Object {
	return v
}

func (v *Float64) Equals(o lang.Object) bool {
	return lang.EqualsFloat64(v, o)
}

func (v *Float64) HashCode() int32 {
	return lang.Hash64(v)
}

func (v *Float64) ToInt64() lang.Int64 {
	return lang.Int64(*v)
}

func (v *Float64) ToUInt64() lang.UInt64 {
	return lang.UInt64(*v)
}

func (v *Float64) ToFloat64() lang.Float64 {
	return lang.Float64(*v)
}

func floatToInt64(f float64) int64 {
	return *(*int64)(unsafe.Pointer(&f))
}

func intToFloat64(f int64) float64 {
	return *(*float64)(unsafe.Pointer(&f))
}
