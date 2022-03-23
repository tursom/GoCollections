package lang

import "strconv"

type Int32 int32

type Rune = Int32

func (i Int32) AsInt32() int32 {
	return int32(i)
}

func (i Int32) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int32) AsObject() Object {
	return i
}

func (i Int32) Equals(e Object) bool {
	i2, ok := e.(Int32)
	if !ok {
		return false
	}
	return i == i2
}

func (i Int32) ToString() String {
	return NewString(i.String())
}

func (i Int32) HashCode() int32 {
	return int32(i)
}

func (i Int32) Compare(t Int32) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
