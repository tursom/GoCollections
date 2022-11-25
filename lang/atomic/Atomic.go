/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package atomic

import (
	"sync/atomic"
	"unsafe"

	"github.com/tursom/GoCollections/lang"
)

type (
	Atomizer[T any] interface {
		Swap(new T) (old T)
		CompareAndSwap(old, new T) (swapped bool)
		Load() (val T)
		Store(val T)
	}

	BitSet interface {
		lang.BitSet
		CompareAndSwapBit(bit int, old, new bool) (swapped bool)
	}

	Atomic[T any] interface {
		Swap() func(addr *T, new T) (old T)
		CompareAndSwap() func(addr *T, old, new T) (swapped bool)
		Load() func(addr *T) (val T)
		Store() func(addr *T, val T)
	}

	atomicImpl[T any] struct {
		swap           func(addr *T, new T) (old T)
		compareAndSwap func(addr *T, old, new T) (swapped bool)
		load           func(addr *T) (val T)
		store          func(addr *T, val T)
	}

	atomicTyped[T any] struct {
	}
)

//goland:noinspection GoUnusedGlobalVariable
var (
	Int32F Atomic[int32] = &atomicImpl[int32]{
		SwapInt32,
		CompareAndSwapInt32,
		LoadInt32,
		StoreInt32,
	}
	Int64F Atomic[int64] = &atomicImpl[int64]{
		SwapInt64,
		CompareAndSwapInt64,
		LoadInt64,
		StoreInt64,
	}
	UInt32F Atomic[uint32] = &atomicImpl[uint32]{
		SwapUInt32,
		CompareAndSwapUInt32,
		LoadUint32,
		StoreUInt32,
	}
	UInt64F Atomic[uint64] = &atomicImpl[uint64]{
		SwapUInt64,
		CompareAndSwapUInt64,
		LoadUint64,
		StoreUInt64,
	}
	PointerF Atomic[unsafe.Pointer] = &atomicImpl[unsafe.Pointer]{
		UnsafeSwapPointer,
		UnsafeCompareAndSwapPointer,
		UnsafeLoadPointer,
		UnsafeStorePointer,
	}
)

func GetAtomic[T any]() Atomic[*T] {
	return atomicTyped[T]{}
}

func (a atomicTyped[T]) Swap() func(addr **T, new *T) (old *T) {
	return SwapPointer[T]
}

func (a atomicTyped[T]) CompareAndSwap() func(addr **T, old *T, new *T) (swapped bool) {
	return CompareAndSwapPointer[T]
}

func (a atomicTyped[T]) Load() func(addr **T) (val *T) {
	return LoadPointer[T]
}

func (a atomicTyped[T]) Store() func(addr **T, val *T) {
	return StorePointer[T]
}

func (a *atomicImpl[T]) Swap() func(addr *T, new T) (old T) {
	return a.swap
}

func (a *atomicImpl[T]) CompareAndSwap() func(addr *T, old T, new T) (swapped bool) {
	return a.compareAndSwap
}

func (a *atomicImpl[T]) Load() func(addr *T) (val T) {
	return a.load
}

func (a *atomicImpl[T]) Store() func(addr *T, val T) {
	return a.store
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

func SetBit[T int32 | int64 | uint32 | uint64](cas func(p *T, old, new T) bool, p *T, bit int, up bool) bool {
	location := T(1) << bit
	var old T
	for {
		old = *p
		var newValue T
		if up {
			newValue = old | location
		} else {
			newValue = old & ^location
		}
		if cas(p, old, newValue) {
			break
		}
	}
	return old&location != 0
}

func SetBitInt32(p *int32, bit int, up bool) bool {
	return SetBit(CompareAndSwapInt32, p, bit, up)
}

func SetBitInt64(p *int64, bit int, up bool) bool {
	return SetBit(CompareAndSwapInt64, p, bit, up)
}

func SetBitUInt32(p *uint32, bit int, up bool) bool {
	return SetBit(CompareAndSwapUInt32, p, bit, up)
}

func SetBitUInt64(p *uint32, bit int, up bool) bool {
	return SetBit(CompareAndSwapUInt32, p, bit, up)
}

func CompareAndSwapBit[T int32 | int64 | uint32 | uint64](cas func(p *T, old, new T) bool, p *T, bit int, old, new bool) bool {
	location := T(1) << bit
	oldValue := *p
	if old {
		oldValue = oldValue | location
	} else {
		oldValue = oldValue & ^location
	}
	var newValue T
	if new {
		newValue = oldValue | location
	} else {
		newValue = oldValue & ^location
	}
	return cas(p, oldValue, newValue)
}

func CompareAndSwapBitInt32(p *int32, bit int, old, new bool) bool {
	return CompareAndSwapBit(CompareAndSwapInt32, p, bit, old, new)
}

func CompareAndSwapBitInt64(p *int64, bit int, old, new bool) bool {
	return CompareAndSwapBit(CompareAndSwapInt64, p, bit, old, new)
}

func CompareAndSwapBitUInt32(p *uint32, bit int, old, new bool) bool {
	return CompareAndSwapBit(CompareAndSwapUInt32, p, bit, old, new)
}

func CompareAndSwapBitUInt64(p *uint64, bit int, old, new bool) bool {
	return CompareAndSwapBit(CompareAndSwapUInt64, p, bit, old, new)
}
