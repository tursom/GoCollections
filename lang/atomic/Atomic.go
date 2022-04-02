package atomic

import (
	"sync/atomic"
	"unsafe"
)

type (
	Atomic[T any] struct {
		Swap           func(addr *T, new T) (old T)
		CompareAndSwap func(addr *T, old, new T) (swapped bool)
		Load           func(addr *T) (val T)
		Store          func(addr *T, val T)
	}
)

//goland:noinspection GoUnusedGlobalVariable
var (
	Int32F = Atomic[int32]{
		SwapInt32,
		CompareAndSwapInt32,
		LoadInt32,
		StoreInt32,
	}
	Int64F = Atomic[int64]{
		SwapInt64,
		CompareAndSwapInt64,
		LoadInt64,
		StoreInt64,
	}
	UInt32F = Atomic[uint32]{
		SwapUInt32,
		CompareAndSwapUInt32,
		LoadUint32,
		StoreUInt32,
	}
	UInt64F = Atomic[uint64]{
		SwapUInt64,
		CompareAndSwapUInt64,
		LoadUint64,
		StoreUInt64,
	}
	PointerF = Atomic[unsafe.Pointer]{
		UnsafeSwapPointer,
		UnsafeCompareAndSwapPointer,
		UnsafeLoadPointer,
		UnsafeStorePointer,
	}
)

func GetAtomic[T any]() *Atomic[*T] {
	return &Atomic[*T]{
		SwapPointer[T],
		CompareAndSwapPointer[T],
		LoadPointer[T],
		StorePointer[T],
	}
}

func SwapInt32(addr *int32, new int32) (old int32) {
	return atomic.SwapInt32(addr, new)
}

func SwapInt64(addr *int64, new int64) (old int64) {
	return atomic.SwapInt64(addr, new)
}

func SwapUInt32(addr *uint32, new uint32) (old uint32) {
	return atomic.SwapUint32(addr, new)
}

func SwapUInt64(addr *uint64, new uint64) (old uint64) {
	return atomic.SwapUint64(addr, new)
}

func SwapUintptr(addr *uintptr, new uintptr) (old uintptr) {
	return atomic.SwapUintptr(addr, new)
}

func SwapUnsafePointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer) {
	return atomic.SwapPointer(addr, new)
}

func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(addr, old, new)
}

func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool) {
	return atomic.CompareAndSwapInt64(addr, old, new)
}

func CompareAndSwapUInt32(addr *uint32, old, new uint32) (swapped bool) {
	return atomic.CompareAndSwapUint32(addr, old, new)
}

func CompareAndSwapUInt64(addr *uint64, old, new uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64(addr, old, new)
}

func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool) {
	return atomic.CompareAndSwapUintptr(addr, old, new)
}

func CompareAndSwapUnsafePointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool) {
	return atomic.CompareAndSwapPointer(addr, old, new)
}

func AddInt32(addr *int32, delta int32) (new int32) {
	return atomic.AddInt32(addr, delta)
}

func AddUInt32(addr *uint32, delta uint32) (new uint32) {
	return atomic.AddUint32(addr, delta)
}

func AddInt64(addr *int64, delta int64) (new int64) {
	return atomic.AddInt64(addr, delta)
}

func AddUInt64(addr *uint64, delta uint64) (new uint64) {
	return atomic.AddUint64(addr, delta)
}

func AddUintptr(addr *uintptr, delta uintptr) (new uintptr) {
	return atomic.AddUintptr(addr, delta)
}

func LoadInt32(addr *int32) (val int32) {
	return atomic.LoadInt32(addr)
}

func LoadInt64(addr *int64) (val int64) {
	return atomic.LoadInt64(addr)
}

func LoadUint32(addr *uint32) (val uint32) {
	return atomic.LoadUint32(addr)
}

func LoadUint64(addr *uint64) (val uint64) {
	return atomic.LoadUint64(addr)
}

func LoadUintptr(addr *uintptr) (val uintptr) {
	return atomic.LoadUintptr(addr)
}

func LoadUnsafePointer(addr *unsafe.Pointer) (val unsafe.Pointer) {
	return atomic.LoadPointer(addr)
}

func StoreInt32(addr *int32, val int32) {
	atomic.StoreInt32(addr, val)
}

func StoreInt64(addr *int64, val int64) {
	atomic.StoreInt64(addr, val)
}

func StoreUInt32(addr *uint32, val uint32) {
	atomic.StoreUint32(addr, val)
}

func StoreUInt64(addr *uint64, val uint64) {
	atomic.StoreUint64(addr, val)
}

func StoreUintptr(addr *uintptr, val uintptr) {
	atomic.StoreUintptr(addr, val)
}

func StoreUnsafePointer(addr *unsafe.Pointer, val unsafe.Pointer) {
	atomic.StorePointer(addr, val)
}
