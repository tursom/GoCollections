package atomic

import (
	"sync/atomic"
)

type UInt32 struct {
	i uint32
}

func (v *UInt32) Load() (val uint32) {
	return atomic.LoadUint32(&v.i)
}

func (v *UInt32) Store(val uint32) {
	atomic.StoreUint32(&v.i, val)
}

func (v *UInt32) Swap(new uint32) (old uint32) {
	return atomic.SwapUint32(&v.i, new)
}

func (v *UInt32) CompareAndSwap(old, new uint32) (swapped bool) {
	return atomic.CompareAndSwapUint32(&v.i, old, new)
}

func (v *UInt32) Add(i uint32) (new uint32) {
	return atomic.AddUint32(&v.i, i)
}
