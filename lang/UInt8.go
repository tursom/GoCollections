package lang

import "strconv"

type UInt8 uint8

type AsUInt8 interface {
	Object
	AsUInt8() UInt8
}

func CastUInt8(v any) (uint8, bool) {
	switch i := v.(type) {
	case uint8:
		return i, true
	case AsUInt8:
		return i.AsUInt8().V(), true
	default:
		return 0, false
	}
}

func EqualsUInt8(i1 AsUInt8, i2 any) bool {
	i2, ok := CastUInt8(i2)
	return ok && i2 == i1.AsUInt8().V()
}

func (i UInt8) V() uint8 {
	return uint8(i)
}

func (i *UInt8) P() *uint8 {
	return (*uint8)(i)
}

func (i UInt8) AsUInt8() UInt8 {
	return i
}

func (i UInt8) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

func (i UInt8) AsObject() Object {
	return i
}

func (i UInt8) Equals(e Object) bool {
	return EqualsUInt8(i, e)
}

func (i UInt8) ToString() String {
	return NewString(i.String())
}

func (i UInt8) HashCode() int32 {
	return int32(i)
}

func (i UInt8) Compare(t UInt8) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
