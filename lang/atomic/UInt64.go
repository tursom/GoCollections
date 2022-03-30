package atomic

import (
	"sync/atomic"
)

type UInt64 struct {
	i uint64
}

func (v *UInt64) Load() (val uint64) {
	return atomic.LoadUint64(&v.i)
}

func (v *UInt64) Store(val uint64) {
	atomic.StoreUint64(&v.i, val)
}

func (v *UInt64) Swap(new uint64) (old uint64) {
	return atomic.SwapUint64(&v.i, new)
}

func (v *UInt64) CompareAndSwap(old, new uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64(&v.i, old, new)
}

func (v *UInt64) Add(i uint64) (new uint64) {
	return atomic.AddUint64(&v.i, i)
}
