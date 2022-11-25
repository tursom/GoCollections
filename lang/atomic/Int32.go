/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package atomic

import (
	"fmt"
	"sync/atomic"

	"github.com/tursom/GoCollections/lang"
)

type Int32 int32

func (v *Int32) AsInt32() lang.Int32 {
	return lang.Int32(*v)
}

func (v *Int32) String() string {
	return fmt.Sprint(*v)
}

func (v *Int32) AsObject() lang.Object {
	return v
}

func (v *Int32) Equals(o lang.Object) bool {
	return lang.EqualsInt32(v, o)
}

func (v *Int32) HashCode() int32 {
	return v.AsInt32().V()
}

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
