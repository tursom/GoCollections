package lang

import "strconv"

type UInt64 uint64

type AsUInt64 interface {
	Object
	AsUInt64() UInt64
}

func CastUInt64(v any) (uint64, bool) {
	switch i := v.(type) {
	case uint64:
		return i, true
	case AsUInt64:
		return i.AsUInt64().V(), true
	default:
		return 0, false
	}
}

func EqualsUInt64(i1 AsUInt64, i2 any) bool {
	i2, ok := CastUInt64(i2)
	return ok && i2 == i1.AsUInt64().V()
}

func (i UInt64) V() uint64 {
	return uint64(i)
}

func (i *UInt64) P() *uint64 {
	return (*uint64)(i)
}

func (i UInt64) AsUInt64() UInt64 {
	return i
}

func (i UInt64) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

func (i UInt64) AsObject() Object {
	return i
}

func (i UInt64) Equals(e Object) bool {
	return EqualsUInt64(i, e)
}

func (i UInt64) ToString() String {
	return NewString(i.String())
}

func (i UInt64) HashCode() int32 {
	return Hash64(&i)
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
