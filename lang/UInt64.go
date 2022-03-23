package lang

import "strconv"

type UInt64 uint64

func (i UInt64) AsUInt64() uint64 {
	return uint64(i)
}

func (i UInt64) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

func (i UInt64) AsObject() Object {
	return i
}

func (i UInt64) Equals(e Object) bool {
	i2, ok := e.(UInt64)
	if !ok {
		return false
	}
	return i == i2
}

func (i UInt64) ToString() String {
	return NewString(i.String())
}

func (i UInt64) HashCode() int32 {
	return Int64(i).HashCode()
}

func (i UInt64) Compare(t UInt64) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
