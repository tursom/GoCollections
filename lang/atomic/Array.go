/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package atomic

type (
	Array[T any] struct {
		atomic Atomic[T]
		array  []T
	}
	Int32Array struct {
		Array[int32]
	}
	Int64Array struct {
		Array[int64]
	}
	UInt32Array struct {
		Array[uint32]
	}
	UInt64Array struct {
		Array[uint64]
	}
)

func NewArray[T any](size int) *Array[*T] {
	return &Array[*T]{
		atomic: GetAtomic[T](),
		array:  make([]*T, size),
	}
}

func CapArray[T any](array []*T) *Array[*T] {
	return &Array[*T]{
		atomic: GetAtomic[T](),
		array:  array,
	}
}

func NewInt32Array(size int) *Int32Array {
	return &Int32Array{
		Array[int32]{
			atomic: Int32F,
			array:  make([]int32, size),
		},
	}
}

func NewInt64Array(size int) *Int64Array {
	return &Int64Array{
		Array[int64]{
			atomic: Int64F,
			array:  make([]int64, size),
		},
	}
}

func NewUInt32Array(size int) *UInt32Array {
	return &UInt32Array{
		Array[uint32]{
			atomic: UInt32F,
			array:  make([]uint32, size),
		},
	}
}

func NewUInt64Array(size int) *UInt64Array {
	return &UInt64Array{
		Array[uint64]{
			atomic: UInt64F,
			array:  make([]uint64, size),
		},
	}
}

func (a *Array[T]) Len() int {
	return len(a.array)
}

func (a *Array[T]) Array() []T {
	return a.array
}

func (a *Array[T]) Get(index int) T {
	return a.atomic.Load()(&a.array[index])
}

func (a *Array[T]) Set(index int, p T) {
	a.atomic.Store()(&a.array[index], p)
}

func (a *Array[T]) Swap(index int, p T) (old T) {
	return a.atomic.Swap()(&a.array[index], p)
}

func (a *Array[T]) CompareAndSwap(index int, old, new T) (swapped bool) {
	return a.atomic.CompareAndSwap()(&a.array[index], old, new)
}

func (a *Int32Array) Add(index int, value int32) {
	AddInt32(&a.array[index], value)
}

func (a *Int64Array) Add(index int, value int64) {
	AddInt64(&a.array[index], value)
}

func (a *UInt32Array) Add(index int, value uint32) {
	AddUInt32(&a.array[index], value)
}

func (a *UInt64Array) Add(index int, value uint64) {
	AddUInt64(&a.array[index], value)
}
