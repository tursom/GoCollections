package atomic

import (
	"sync/atomic"
)

type Int64 struct {
	i int64
}

func (v *Int64) Load() (val int64) {
	return atomic.LoadInt64(&v.i)
}

func (v *Int64) Store(val int64) {
	atomic.StoreInt64(&v.i, val)
}

func (v *Int64) Swap(new int64) (old int64) {
	return atomic.SwapInt64(&v.i, new)
}

func (v *Int64) CompareAndSwap(old, new int64) (swapped bool) {
	return atomic.CompareAndSwapInt64(&v.i, old, new)
}

func (v *Int64) Add(i int64) (new int64) {
	return atomic.AddInt64(&v.i, i)
}
