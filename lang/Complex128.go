package lang

import (
	"fmt"
)

type Complex128 complex128

func (i Complex128) AsComplex128() complex128 {
	return complex128(i)
}

func (i Complex128) String() string {
	return fmt.Sprint(complex64(i))
}

func (i Complex128) AsObject() Object {
	return i
}

func (i Complex128) Equals(e Object) bool {
	i2, ok := e.(Complex128)
	if !ok {
		return false
	}
	return i == i2
}

func (i Complex128) ToString() String {
	return NewString(i.String())
}

func (i Complex128) HashCode() int32 {
	return Float64(real(complex64(i))).HashCode() ^ Float64(imag(complex64(i))).HashCode()
}
