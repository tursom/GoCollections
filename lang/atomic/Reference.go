package atomic

import (
	"sync/atomic"
	"unsafe"

	"github.com/tursom/GoCollections/lang"
)

//goland:noinspection GoUnusedGlobalVariable
var (
	UnsafeLoadPointer           = atomic.LoadPointer
	UnsafeStorePointer          = atomic.StorePointer
	UnsafeSwapPointer           = atomic.SwapPointer
	UnsafeCompareAndSwapPointer = atomic.CompareAndSwapPointer
)

type (
	Pointer  = unsafe.Pointer
	PPointer = *unsafe.Pointer

	// Reference atomic type T reference
	Reference[T any] struct {
		lang.BaseObject
		p *T
	}
)

// NewReference new *Reference[T] init by given reference
func NewReference[T any](reference *T) *Reference[T] {
	return lang.ForceCast[Reference[T]](Pointer(&reference))
}

// ReferenceOf cast **T to *Reference[T]
func ReferenceOf[T any](reference **T) *Reference[T] {
	return lang.ForceCast[Reference[T]](Pointer(reference))
}

func ReferenceUintptr[T any](reference *uintptr) *Reference[T] {
	return lang.ForceCast[Reference[T]](Pointer(reference))
}

func (r *Reference[T]) AsPointer() Pointer {
	return Pointer(r)
}

func (r *Reference[T]) AsPPointer() PPointer {
	return PPointer(r.AsPointer())
}

func (r *Reference[T]) AsUintptr() *TypedUintptr[T] {
	return lang.ForceCast[TypedUintptr[T]](r.AsPointer())
}

func (r *Reference[T]) pointer() **T {
	return &r.p
}

func (r *Reference[T]) Load() (val *T) {
	return LoadPointer(r.pointer())
}

func (r *Reference[T]) Store(val *T) {
	StorePointer(r.pointer(), val)
}

func (r *Reference[T]) Swap(new *T) (old *T) {
	return SwapPointer(r.pointer(), new)
}

func (r *Reference[T]) CompareAndSwap(old, new *T) (swapped bool) {
	return CompareAndSwapPointer(r.pointer(), old, new)
}

func AsPPointer[T any](p **T) PPointer {
	return PPointer(Pointer(p))
}

func LoadPointer[T any](addr **T) (val *T) {
	return (*T)(atomic.LoadPointer(AsPPointer(addr)))
}

func StorePointer[T any](addr **T, val *T) {
	atomic.StorePointer(AsPPointer(addr), Pointer(val))
}

func SwapPointer[T any](addr **T, new *T) (old *T) {
	return (*T)(atomic.SwapPointer(AsPPointer(addr), Pointer(new)))
}

func CompareAndSwapPointer[T any](addr **T, old, new *T) (swapped bool) {
	return atomic.CompareAndSwapPointer(AsPPointer(addr), Pointer(old), Pointer(new))
}
