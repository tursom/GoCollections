package atomic

import (
	"sync/atomic"
	"unsafe"
)

//goland:noinspection GoUnusedGlobalVariable
var (
	UnsafeLoadPointer           = atomic.LoadPointer
	UnsafeStorePointer          = atomic.StorePointer
	UnsafeSwapPointer           = atomic.SwapPointer
	UnsafeCompareAndSwapPointer = atomic.CompareAndSwapPointer
)

type (
	Pointer          = unsafe.Pointer
	PPointer         = *unsafe.Pointer
	Reference[T any] struct {
		reference *T
	}
)

func NewReference[T any](reference *T) *Reference[T] {
	return &Reference[T]{reference}
}

func (v *Reference[T]) Load() (val *T) {
	return LoadPointer(&v.reference)
}

func (v *Reference[T]) Store(val *T) {
	StorePointer(&v.reference, val)
}

func (v *Reference[T]) Swap(new *T) (old *T) {
	return SwapPointer(&v.reference, new)
}

func (v *Reference[T]) CompareAndSwap(old, new *T) (swapped bool) {
	return CompareAndSwapPointer(&v.reference, old, new)
}

func AsPPointer[T any](p **T) *unsafe.Pointer {
	return (*unsafe.Pointer)(unsafe.Pointer(p))
}

func LoadPointer[T any](addr **T) (val *T) {
	return (*T)(atomic.LoadPointer(AsPPointer(addr)))
}

func StorePointer[T any](addr **T, val *T) {
	atomic.StorePointer(AsPPointer(addr), unsafe.Pointer(val))
}

func SwapPointer[T any](addr **T, new *T) (old *T) {
	return (*T)(atomic.SwapPointer(AsPPointer(addr), unsafe.Pointer(new)))
}

func CompareAndSwapPointer[T any](addr **T, old, new *T) (swapped bool) {
	return atomic.CompareAndSwapPointer(AsPPointer(addr), Pointer(old), Pointer(new))
}
