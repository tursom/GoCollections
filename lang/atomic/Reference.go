package atomic

import (
	"sync/atomic"
	"unsafe"
)

type Reference[T any] struct {
	reference *T
}

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

func LoadPointer[T any](addr **T) (val *T) {
	return (*T)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(addr))))
}

func StorePointer[T any](addr **T, val *T) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(addr)), unsafe.Pointer(val))
}

func SwapPointer[T any](addr **T, new *T) (old *T) {
	return (*T)(atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(addr)), unsafe.Pointer(new)))
}

func CompareAndSwapPointer[T any](addr **T, old, new *T) (swapped bool) {
	return atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(addr)), unsafe.Pointer(old), unsafe.Pointer(new))
}
