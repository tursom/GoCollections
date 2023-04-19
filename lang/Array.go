package lang

import "unsafe"

type (
	Array[T any] []T

	Int8Array       Array[Int8]
	Int16Array      Array[Int16]
	Int32Array      Array[Int32]
	Int64Array      Array[Int64]
	UInt8Array      Array[UInt8]
	UInt16Array     Array[UInt16]
	UInt32Array     Array[UInt32]
	UInt64Array     Array[UInt64]
	Float32Array    Array[Float32]
	Float64Array    Array[Float64]
	Complex64Array  Array[Complex64]
	Complex128Array Array[Complex128]
)

func (a Array[T]) Array() []T {
	return a
}

func FromRawInt8Array(arr []int8) Int8Array {
	return *(*Int8Array)(unsafe.Pointer(&arr))
}

func FromRawInt16Array(arr []int16) Int16Array {
	return *(*Int16Array)(unsafe.Pointer(&arr))
}

func FromRawInt32Array(arr []int32) Int32Array {
	return *(*Int32Array)(unsafe.Pointer(&arr))
}

func FromRawInt64Array(arr []int64) Int64Array {
	return *(*Int64Array)(unsafe.Pointer(&arr))
}

func FromRawUInt8Array(arr []uint8) UInt8Array {
	return *(*UInt8Array)(unsafe.Pointer(&arr))
}

func FromRawUInt16Array(arr []uint16) UInt16Array {
	return *(*UInt16Array)(unsafe.Pointer(&arr))
}

func FromRawUInt32Array(arr []uint32) UInt32Array {
	return *(*UInt32Array)(unsafe.Pointer(&arr))
}

func FromRawUInt64Array(arr []uint64) UInt64Array {
	return *(*UInt64Array)(unsafe.Pointer(&arr))
}

func FromRawFloat32Array(arr []float32) Float32Array {
	return *(*Float32Array)(unsafe.Pointer(&arr))
}

func FromRawFloat64Array(arr []float64) Float64Array {
	return *(*Float64Array)(unsafe.Pointer(&arr))
}

func FromRawComplex64Array(arr []complex64) Complex64Array {
	return *(*Complex64Array)(unsafe.Pointer(&arr))
}

func FromRawComplex128Array(arr []complex128) Complex128Array {
	return *(*Complex128Array)(unsafe.Pointer(&arr))
}

func (a UInt8Array) Bytes() []byte {
	return a.Raw()
}

func (a Int8Array) Raw() []int8 {
	return *(*[]int8)(unsafe.Pointer(&a))
}

func (a Int16Array) Raw() []int16 {
	return *(*[]int16)(unsafe.Pointer(&a))
}

func (a Int32Array) Raw() []int32 {
	return *(*[]int32)(unsafe.Pointer(&a))
}

func (a Int64Array) Raw() []int64 {
	return *(*[]int64)(unsafe.Pointer(&a))
}

func (a UInt8Array) Raw() []uint8 {
	return *(*[]uint8)(unsafe.Pointer(&a))
}

func (a UInt16Array) Raw() []uint16 {
	return *(*[]uint16)(unsafe.Pointer(&a))
}

func (a UInt32Array) Raw() []uint32 {
	return *(*[]uint32)(unsafe.Pointer(&a))
}

func (a UInt64Array) Raw() []uint64 {
	return *(*[]uint64)(unsafe.Pointer(&a))
}

func (a Float32Array) Raw() []float32 {
	return *(*[]float32)(unsafe.Pointer(&a))
}

func (a Float64Array) Raw() []float64 {
	return *(*[]float64)(unsafe.Pointer(&a))
}

func (a Complex64Array) Raw() []complex64 {
	return *(*[]complex64)(unsafe.Pointer(&a))
}

func (a Complex128Array) Raw() []complex128 {
	return *(*[]complex128)(unsafe.Pointer(&a))
}

func (a Int8Array) BitLength() uint {
	return uint(len(a) * 8)
}

func (a Int16Array) BitLength() uint {
	return uint(len(a) * 16)
}

func (a Int32Array) BitLength() uint {
	return uint(len(a) * 32)
}

func (a Int64Array) BitLength() uint {
	return uint(len(a) * 64)
}

func (a UInt8Array) BitLength() uint {
	return uint(len(a) * 8)
}

func (a UInt16Array) BitLength() uint {
	return uint(len(a) * 16)
}

func (a UInt32Array) BitLength() uint {
	return uint(len(a) * 32)
}

func (a UInt64Array) BitLength() uint {
	return uint(len(a) * 64)
}

func (a Int8Array) GetBit(index uint) (ok bool) {
	return GetBit(a[index/8], index%8)
}

func (a Int16Array) GetBit(index uint) (ok bool) {
	return GetBit(a[index/16], index%16)
}

func (a Int32Array) GetBit(index uint) (ok bool) {
	return GetBit(a[index/32], index%32)
}

func (a Int64Array) GetBit(index uint) (ok bool) {
	return GetBit(a[index/64], index%64)
}

func (a UInt8Array) GetBit(index uint) (ok bool) {
	return GetBit(a[index/8], index%8)
}

func (a UInt16Array) GetBit(index uint) (ok bool) {
	return GetBit(a[index/16], index%16)
}

func (a UInt32Array) GetBit(index uint) (ok bool) {
	return GetBit(a[index/32], index%32)
}

func (a UInt64Array) GetBit(index uint) (ok bool) {
	return GetBit(a[index/64], index%64)
}

func (a Int8Array) SetBit(index uint, up bool) {
	SetBit(&a[index/8], index%8, up)
}

func (a Int16Array) SetBit(index uint, up bool) {
	SetBit(&a[index/16], index%16, up)
}

func (a Int32Array) SetBit(index uint, up bool) {
	SetBit(&a[index/32], index%32, up)
}

func (a Int64Array) SetBit(index uint, up bool) {
	SetBit(&a[index/64], index%64, up)
}

func (a UInt8Array) SetBit(index uint, up bool) {
	SetBit(&a[index/8], index%8, up)
}

func (a UInt16Array) SetBit(index uint, up bool) {
	SetBit(&a[index/16], index%16, up)
}

func (a UInt32Array) SetBit(index uint, up bool) {
	SetBit(&a[index/32], index%32, up)
}

func (a UInt64Array) SetBit(index uint, up bool) {
	SetBit(&a[index/64], index%64, up)
}

func (a Int8Array) SwapBit(index uint, up bool) (old bool) {
	return SwapBit(&a[index/8], index%8, up)
}

func (a Int16Array) SwapBit(index uint, up bool) (old bool) {
	return SwapBit(&a[index/16], index%16, up)
}

func (a Int32Array) SwapBit(index uint, up bool) (old bool) {
	return SwapBit(&a[index/32], index%32, up)
}

func (a Int64Array) SwapBit(index uint, up bool) (old bool) {
	return SwapBit(&a[index/64], index%64, up)
}

func (a UInt8Array) SwapBit(index uint, up bool) (old bool) {
	return SwapBit(&a[index/8], index%8, up)
}

func (a UInt16Array) SwapBit(index uint, up bool) (old bool) {
	return SwapBit(&a[index/16], index%16, up)
}

func (a UInt32Array) SwapBit(index uint, up bool) (old bool) {
	return SwapBit(&a[index/32], index%32, up)
}

func (a UInt64Array) SwapBit(index uint, up bool) (old bool) {
	return SwapBit(&a[index/64], index%64, up)
}
