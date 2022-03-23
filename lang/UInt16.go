package lang

import "strconv"

type UInt16 uint16

func (i UInt16) AsUInt16() uint16 {
	return uint16(i)
}

func (i UInt16) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

func (i UInt16) AsObject() Object {
	return i
}

func (i UInt16) Equals(e Object) bool {
	i2, ok := e.(UInt16)
	if !ok {
		return false
	}
	return i == i2
}

func (i UInt16) ToString() String {
	return NewString(i.String())
}

func (i UInt16) HashCode() int32 {
	return int32(i)
}

func (i UInt16) Compare(t UInt16) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
