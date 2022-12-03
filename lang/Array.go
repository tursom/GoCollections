package lang

type (
	Array[T any] []T

	Int8Array   Array[int32]
	Int16Array  Array[int32]
	Int32Array  Array[int32]
	Int64Array  Array[int64]
	UInt8Array  Array[uint32]
	UInt16Array Array[uint32]
	UInt32Array Array[uint32]
	UInt64Array Array[uint64]
)

func (a Array[T]) Array() []T {
	return a
}

func (a Int8Array) SetBit(bit int, up bool) (old bool) {
	arrIndex := bit / 8
	i := &a[arrIndex]

	return SwapBit(i, bit%8, up)
}

func (a Int16Array) SetBit(bit int, up bool) (old bool) {
	arrIndex := bit / 16
	i := &a[arrIndex]

	return SwapBit(i, bit%16, up)
}

func (a Int32Array) SetBit(bit int, up bool) (old bool) {
	arrIndex := bit / 32
	i := &a[arrIndex]

	return SwapBit(i, bit%32, up)
}

func (a Int64Array) SetBit(bit int, up bool) (old bool) {
	arrIndex := bit / 64
	i := &a[arrIndex]

	return SwapBit(i, bit%64, up)
}

func (a UInt8Array) SetBit(bit int, up bool) (old bool) {
	arrIndex := bit / 8
	i := &a[arrIndex]

	return SwapBit(i, bit%8, up)
}

func (a UInt16Array) SetBit(bit int, up bool) (old bool) {
	arrIndex := bit / 16
	i := &a[arrIndex]

	return SwapBit(i, bit%16, up)
}

func (a UInt32Array) SetBit(bit int, up bool) (old bool) {
	arrIndex := bit / 32
	i := &a[arrIndex]

	return SwapBit(i, bit%32, up)
}

func (a UInt64Array) SetBit(bit int, up bool) (old bool) {
	arrIndex := bit / 64
	i := &a[arrIndex]

	return SwapBit(i, bit%64, up)
}
