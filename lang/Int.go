package lang

import "strconv"

type Int int

func (i Int) AsInt() int {
	return int(i)
}

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func (i Int) Equals(e Object) bool {
	i2, ok := e.(Int)
	if !ok {
		return false
	}
	return i == i2
}

func (i Int) AsObject() Object {
	return i
}

func (i Int) ToString() String {
	return NewString(i.String())
}

func (i Int) HashCode() int32 {
	return Int64(i).HashCode()
}

func (i Int) Compare(t Int) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
