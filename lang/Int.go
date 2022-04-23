package lang

import "strconv"

type Int int

type AsInt interface {
	Object
	AsInt() Int
}

func CastInt(v any) (int, bool) {
	switch i := v.(type) {
	case int:
		return i, true
	case AsInt:
		return i.AsInt().V(), true
	default:
		return 0, false
	}
}

func EqualsInt(i1 AsInt, i2 any) bool {
	i2, ok := CastInt(i2)
	return ok && i2 == i1.AsInt().V()
}

func (i Int) V() int {
	return int(i)
}

func (i *Int) P() *int {
	return (*int)(i)
}

func (i Int) AsInt() Int {
	return i
}

func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int) AsObject() Object {
	return i
}

func (i Int) Equals(e Object) bool {
	return EqualsInt(i, e)
}

func (i Int) ToString() String {
	return NewString(i.String())
}

func (i Int) HashCode() int32 {
	return int32(i)
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
