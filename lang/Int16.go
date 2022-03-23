package lang

import "strconv"

type Int16 int16

func (i Int16) AsInt16() int16 {
	return int16(i)
}

func (i Int16) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int16) AsObject() Object {
	return i
}

func (i Int16) Equals(e Object) bool {
	i2, ok := e.(Int16)
	if !ok {
		return false
	}
	return i == i2
}

func (i Int16) ToString() String {
	return NewString(i.String())
}

func (i Int16) HashCode() int32 {
	return int32(i)
}

func (i Int16) Compare(t Int16) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
