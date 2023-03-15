package lang

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

func (a Int8Array) SetBit(bit int, up bool) (old bool) {
	arrIndex := bit / 8
	i := &a[arrIndex]

	return SwapBit[int8]((*int8)(i), bit%8, up)
}

func (a Int16Array) SetBit(bit int, up bool) (old bool) {
	arrIndex := bit / 16
	i := &a[arrIndex]

	return SwapBit[int16]((*int16)(i), bit%16, up)
}

func (a Int32Array) SetBit(bit int, up bool) (old bool) {
	arrIndex := bit / 32
	i := &a[arrIndex]

	return SwapBit[int32]((*int32)(i), bit%32, up)
}

func (a Int64Array) SetBit(bit int, up bool) (old bool) {
	arrIndex := bit / 64
	i := &a[arrIndex]

	return SwapBit[int64]((*int64)(i), bit%64, up)
}

func (a UInt8Array) SetBit(bit int, up bool) (old bool) {
	arrIndex := bit / 8
	i := &a[arrIndex]

	return SwapBit[uint8]((*uint8)(i), bit%8, up)
}

func (a UInt16Array) SetBit(bit int, up bool) (old bool) {
	arrIndex := bit / 16
	i := &a[arrIndex]

	return SwapBit[uint16]((*uint16)(i), bit%16, up)
}

func (a UInt32Array) SetBit(bit int, up bool) (old bool) {
	arrIndex := bit / 32
	i := &a[arrIndex]

	return SwapBit[uint32]((*uint32)(i), bit%32, up)
}

func (a UInt64Array) SetBit(bit int, up bool) (old bool) {
	arrIndex := bit / 64
	i := &a[arrIndex]

	return SwapBit[uint64]((*uint64)(i), bit%64, up)
}
