package lang

import (
	"fmt"
	"unsafe"
)

type (
	AsObject interface {
		AsObject() Object
	}

	Equable interface {
		Equals(o Object) bool
	}

	Object interface {
		fmt.Stringer
		AsObject
		Equable
		HashCode() int32
	}

	Any = Object

	BaseObject struct {
	}
)

func ToString(obj Object) String {
	return NewString(obj.String())
}

func Equals(e Object, t Object) bool {
	if e == nil {
		return t == nil
	}
	return e.Equals(t)
}

func HashCode(obj Object) int32 {
	if obj == nil {
		return 0
	} else {
		return obj.HashCode()
	}
}

func NewBaseObject() BaseObject {
	return BaseObject{}
}

func (b *BaseObject) AsObject() Object {
	return b
}

func (b *BaseObject) Equals(o Object) bool {
	return b == o
}

func (b *BaseObject) GoString() string {
	return b.String()
}

func (b *BaseObject) String() string {
	return fmt.Sprintf("BaseObject@%p", unsafe.Pointer(b))
}

func (b *BaseObject) ToString() String {
	return NewString(b.String())
}

func (b *BaseObject) HashCode() int32 {
	return Hash64(b)
}
