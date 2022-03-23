package lang

import "strconv"

type Int64 int64

func (i Int64) AsInt64() int64 {
	return int64(i)
}

func (i Int64) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int64) AsObject() Object {
	return i
}

func (i Int64) Equals(e Object) bool {
	i2, ok := e.(Int64)
	if !ok {
		return false
	}
	return i == i2
}

func (i Int64) ToString() String {
	return NewString(i.String())
}

func (i Int64) HashCode() int32 {
	return Hash64(&i)
}

func (i Int64) Compare(t Int64) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
