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

func (a UInt8Array) Bytes() []byte {
	return *(*[]byte)(unsafe.Pointer(&a))
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

func (a Int8Array) SetBit(index uint, up bool) (old bool) {
	return SwapBit(&a[index/8], index%8, up)
}

func (a Int16Array) SetBit(index uint, up bool) (old bool) {
	return SwapBit(&a[index/16], index%16, up)
}

func (a Int32Array) SetBit(index uint, up bool) (old bool) {
	return SwapBit(&a[index/32], index%32, up)
}

func (a Int64Array) SetBit(index uint, up bool) (old bool) {
	return SwapBit(&a[index/64], index%64, up)
}

func (a UInt8Array) SetBit(index uint, up bool) (old bool) {
	return SwapBit(&a[index/8], index%8, up)
}

func (a UInt16Array) SetBit(index uint, up bool) (old bool) {
	return SwapBit(&a[index/16], index%16, up)
}

func (a UInt32Array) SetBit(index uint, up bool) (old bool) {
	return SwapBit(&a[index/32], index%32, up)
}

func (a UInt64Array) SetBit(index uint, up bool) (old bool) {
	return SwapBit(&a[index/64], index%64, up)
}
