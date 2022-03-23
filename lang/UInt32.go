package lang

import "strconv"

type UInt32 uint32

func (i UInt32) AsUInt32() uint32 {
	return uint32(i)
}

func (i UInt32) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

func (i UInt32) AsObject() Object {
	return i
}

func (i UInt32) Equals(e Object) bool {
	i2, ok := e.(UInt32)
	if !ok {
		return false
	}
	return i == i2
}

func (i UInt32) ToString() String {
	return NewString(i.String())
}

func (i UInt32) HashCode() int32 {
	return int32(i)
}

func (i UInt32) Compare(t UInt32) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
