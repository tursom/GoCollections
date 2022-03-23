package lang

import (
	"fmt"
)

type Float32 float32

func (i Float32) AsFloat32() float32 {
	return float32(i)
}

func (i Float32) String() string {
	return fmt.Sprint(float32(i))
}

func (i Float32) AsObject() Object {
	return i
}

func (i Float32) Equals(e Object) bool {
	i2, ok := e.(Float32)
	if !ok {
		return false
	}
	return i == i2
}

func (i Float32) ToString() String {
	return NewString(i.String())
}

func (i Float32) HashCode() int32 {
	return Hash32(&i)
}

func (i Float32) Compare(t Float32) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
