package lang

import "strconv"

type UInt uint

func (i UInt) AsUInt() uint {
	return uint(i)
}

func (i UInt) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

func (i UInt) AsObject() Object {
	return i
}

func (i UInt) Equals(e Object) bool {
	i2, ok := e.(UInt)
	if !ok {
		return false
	}
	return i == i2
}

func (i UInt) ToString() String {
	return NewString(i.String())
}

func (i UInt) HashCode() int32 {
	return Int64(i).HashCode()
}

func (i UInt) Compare(t UInt) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
