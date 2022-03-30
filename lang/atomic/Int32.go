package atomic

import (
	"sync/atomic"
)

type Int32 struct {
	i int32
}

func (v *Int32) Load() (val int32) {
	return atomic.LoadInt32(&v.i)
}

func (v *Int32) Store(val int32) {
	atomic.StoreInt32(&v.i, val)
}

func (v *Int32) Swap(new int32) (old int32) {
	return atomic.SwapInt32(&v.i, new)
}

func (v *Int32) CompareAndSwap(old, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(&v.i, old, new)
}

func (v *Int32) Add(i int32) (new int32) {
	return atomic.AddInt32(&v.i, i)
}
