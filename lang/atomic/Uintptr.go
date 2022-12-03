/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package atomic

import (
	"sync/atomic"

	"github.com/tursom/GoCollections/unsafe"
)

type (
	Uintptr             uintptr
	TypedUintptr[T any] Uintptr
)

func (p *Uintptr) Raw() uintptr {
	return uintptr(*p)
}

func (p *Uintptr) RawP() *uintptr {
	return (*uintptr)(p)
}

func (p *Uintptr) AsPointer() Pointer {
	return Pointer(p)
}

func (p *Uintptr) AsPPointer() PPointer {
	return PPointer(p.AsPointer())
}
func (p *Uintptr) Load() (val Uintptr) {
	return Uintptr(atomic.LoadUintptr((*uintptr)(p)))
}

func (p *Uintptr) Store(val Uintptr) {
	atomic.StoreUintptr((*uintptr)(p), uintptr(val))
}

func (p *Uintptr) Swap(new Uintptr) (old Uintptr) {
	return Uintptr(atomic.SwapUintptr((*uintptr)(p), uintptr(new)))
}

func (p *Uintptr) CompareAndSwap(old, new Uintptr) (swapped bool) {
	return atomic.CompareAndSwapUintptr((*uintptr)(p), uintptr(old), uintptr(new))
}

func (p *TypedUintptr[T]) Raw() uintptr {
	return uintptr(*p)
}

func (p *TypedUintptr[T]) RawP() *uintptr {
	return (*uintptr)(p)
}

func (p *TypedUintptr[T]) AsPointer() Pointer {
	return Pointer(p)
}

func (p *TypedUintptr[T]) AsPPointer() PPointer {
	return PPointer(p.AsPointer())
}

func (tp *TypedUintptr[T]) AsReference() *Reference[T] {
	return unsafe.ForceCast[Reference[T]](Pointer(tp))
}

func (tp *TypedUintptr[T]) Uintptr() Uintptr {
	return Uintptr(*tp)
}

func (tp *TypedUintptr[T]) PUintptr() *Uintptr {
	return (*Uintptr)(tp)
}

func (tp *TypedUintptr[T]) Load() TypedUintptr[T] {
	return TypedUintptr[T](atomic.LoadUintptr((*uintptr)(tp)))
}

func (tp *TypedUintptr[T]) Store(val TypedUintptr[T]) {
	atomic.StoreUintptr((*uintptr)(tp), uintptr(val))
}

func (tp *TypedUintptr[T]) Swap(new TypedUintptr[T]) (old TypedUintptr[T]) {
	return TypedUintptr[T](atomic.SwapUintptr((*uintptr)(tp), uintptr(new)))
}

func (tp *TypedUintptr[T]) CompareAndSwap(old, new TypedUintptr[T]) (swapped bool) {
	return atomic.CompareAndSwapUintptr((*uintptr)(tp), uintptr(old), uintptr(new))
}
