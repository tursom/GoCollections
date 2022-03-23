package lang

import "strconv"

type UInt8 uint8

type Byte = UInt8

func (i UInt8) AsUInt8() uint8 {
	return uint8(i)
}

func (i UInt8) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

func (i UInt8) AsObject() Object {
	return i
}

func (i UInt8) Equals(e Object) bool {
	i2, ok := e.(UInt8)
	if !ok {
		return false
	}
	return i == i2
}

func (i UInt8) ToString() String {
	return NewString(i.String())
}

func (i UInt8) HashCode() int32 {
	return int32(i)
}

func (i UInt8) Compare(t UInt8) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
