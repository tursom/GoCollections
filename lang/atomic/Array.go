/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package atomic

import "github.com/tursom/GoCollections/lang"

type (
	Array[T any] lang.Array[*T]
	Int32Array   lang.Array[Int32]
	Int64Array   lang.Array[Int64]
	UInt32Array  lang.Array[UInt32]
	UInt64Array  lang.Array[UInt64]
)

func NewArray[T any](size int) Array[T] {
	return make([]*T, size)
}

func CapArray[T any](array []*T) Array[T] {
	return array
}

func NewInt32Array(size int) Int32Array {
	return make(Int32Array, size)
}

func NewInt64Array(size int) Int64Array {
	return make(Int64Array, size)
}

func NewUInt32Array(size int) UInt32Array {
	return make(UInt32Array, size)
}

func NewUInt64Array(size int) UInt64Array {
	return make(UInt64Array, size)
}

func (a Array[T]) Len() int {
	return len(a)
}

func (a Array[T]) Array() []*T {
	return a
}

func (a Array[T]) Get(index int) *T {
	return GetAtomizer[T]().Load()(&a[index])
}

func (a Int32Array) Get(index int) int32 {
	return a[index].Load()
}

func (a Int64Array) Get(index int) int64 {
	return a[index].Load()
}

func (a UInt32Array) Get(index int) uint32 {
	return a[index].Load()
}

func (a UInt64Array) Get(index int) uint64 {
	return a[index].Load()
}

func (a Array[T]) Set(index int, p *T) {
	GetAtomizer[T]().Store()(&a[index], p)
}

func (a Int32Array) Set(index int, v int32) {
	a[index].Store(v)
}

func (a Int64Array) Set(index int, v int64) {
	a[index].Store(v)
}

func (a UInt32Array) Set(index int, v uint32) {
	a[index].Store(v)
}

func (a UInt64Array) Set(index int, v uint64) {
	a[index].Store(v)
}

func (a Array[T]) Swap(index int, p *T) (old *T) {
	return GetAtomizer[T]().Swap()(&a[index], p)
}

func (a Int32Array) Swap(index int, new int32) (old int32) {
	return a[index].Swap(new)
}

func (a Int64Array) Swap(index int, new int64) (old int64) {
	return a[index].Swap(new)
}

func (a UInt32Array) Swap(index int, new uint32) (old uint32) {
	return a[index].Swap(new)
}

func (a UInt64Array) Swap(index int, new uint64) (old uint64) {
	return a[index].Swap(new)
}

func (a Array[T]) CompareAndSwap(index int, old, new *T) (swapped bool) {
	return GetAtomizer[T]().CompareAndSwap()(&a[index], old, new)
}

func (a Int32Array) CompareAndSwap(index int, old, new int32) (swapped bool) {
	return a[index].CompareAndSwap(old, new)
}

func (a Int64Array) CompareAndSwap(index int, old, new int64) (swapped bool) {
	return a[index].CompareAndSwap(old, new)
}

func (a UInt32Array) CompareAndSwap(index int, old, new uint32) (swapped bool) {
	return a[index].CompareAndSwap(old, new)
}

func (a UInt64Array) CompareAndSwap(index int, old, new uint64) (swapped bool) {
	return a[index].CompareAndSwap(old, new)
}

func (a Int32Array) Add(index int, value int32) {
	a[index].Add(value)
}

func (a Int64Array) Add(index int, value int64) {
	a[index].Add(value)
}

func (a UInt32Array) Add(index int, value uint32) {
	a[index].Add(value)
}

func (a UInt64Array) Add(index int, value uint64) {
	a[index].Add(value)
}

func (a Int32Array) BitLength() int {
	return len(a) * 32
}

func (a Int64Array) BitLength() int {
	return len(a) * 64
}

func (a UInt32Array) BitLength() int {
	return len(a) * 32
}

func (a UInt64Array) BitLength() int {
	return len(a) * 64
}

func (a Int32Array) SetBit(bit int, up bool) (old bool) {
	return a[bit/32].SetBit(bit%32, up)
}

func (a Int64Array) SetBit(bit int, up bool) (old bool) {
	return a[bit/64].SetBit(bit%64, up)
}

func (a UInt32Array) SetBit(bit int, up bool) (old bool) {
	return a[bit/32].SetBit(bit%32, up)
}

func (a UInt64Array) SetBit(bit int, up bool) (old bool) {
	return a[bit/64].SetBit(bit%64, up)
}

func (a Int32Array) CompareAndSwapBit(bit int, old, new bool) (swapped bool) {
	return a[bit/32].CompareAndSwapBit(bit%32, old, new)
}

func (a Int64Array) CompareAndSwapBit(bit int, old, new bool) (swapped bool) {
	return a[bit/64].CompareAndSwapBit(bit%64, old, new)
}

func (a UInt32Array) CompareAndSwapBit(bit int, old, new bool) (swapped bool) {
	return a[bit/32].CompareAndSwapBit(bit%32, old, new)
}

func (a UInt64Array) CompareAndSwapBit(bit int, old, new bool) (swapped bool) {
	return a[bit/64].CompareAndSwapBit(bit%64, old, new)
}
