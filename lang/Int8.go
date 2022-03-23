package lang

import "strconv"

type Int8 int8

func (i Int8) AsInt8() int8 {
	return int8(i)
}

func (i Int8) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int8) AsObject() Object {
	return i
}

func (i Int8) Equals(e Object) bool {
	i2, ok := e.(Int8)
	if !ok {
		return false
	}
	return i == i2
}

func (i Int8) ToString() String {
	return NewString(i.String())
}

func (i Int8) HashCode() int32 {
	return int32(i)
}

func (i Int8) Compare(t Int8) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
