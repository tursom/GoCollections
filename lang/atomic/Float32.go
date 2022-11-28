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

type Float32 float32

func (v *Float32) P() *float32 {
	return (*float32)(v)
}

func (v *Float32) pint32() *int32 {
	return (*int32)(unsafe.Pointer(v))
}

func (v *Float32) Load() (val float32) {
	return intToFloat32(atomic.LoadInt32(v.pint32()))
}

func (v *Float32) Store(val float32) {
	atomic.StoreInt32(v.pint32(), floatToInt32(val))
}

func (v *Float32) Swap(new float32) (old float32) {
	i := atomic.SwapInt32(v.pint32(), floatToInt32(new))
	return intToFloat32(i)
}

func (v *Float32) CompareAndSwap(old, new float32) (swapped bool) {
	return atomic.CompareAndSwapInt32(v.pint32(), floatToInt32(old), floatToInt32(new))
}

func (v *Float32) Add(f float32) (new float32) {
	old := float32(*v)
	new = old + f
	for !v.CompareAndSwap(old, new) {
		old = float32(*v)
		new = old + f
	}

	return new
}

func (v *Float32) String() string {
	return strconv.FormatFloat(float64(*v), 'e', -1, 32)
}

func (v *Float32) AsObject() lang.Object {
	return v
}

func (v *Float32) Equals(o lang.Object) bool {
	return lang.EqualsFloat32(v, o)
}

func (v *Float32) HashCode() int32 {
	return lang.Hash32(v)
}

func (v *Float32) ToInt64() lang.Int64 {
	return lang.Int64(*v)
}

func (v *Float32) ToUInt64() lang.UInt64 {
	return lang.UInt64(*v)
}

func (v *Float32) ToFloat64() lang.Float64 {
	return lang.Float64(*v)
}

func floatToInt32(f float32) int32 {
	return *(*int32)(unsafe.Pointer(&f))
}

func intToFloat32(f int32) float32 {
	return *(*float32)(unsafe.Pointer(&f))
}
