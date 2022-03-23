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
		AsObject() Object
		Equals(o Object) bool
		ToString() String
		HashCode() int32
	}

	Any = Object

	BaseObject struct {
	}
)

func Equals(e Object, t Object) bool {
	if e == nil {
		return t == nil
	}
	return e.Equals(t)
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

func (b *BaseObject) ToString() String {
	return NewString(fmt.Sprint(unsafe.Pointer(b)))
}

func (b *BaseObject) HashCode() int32 {
	return Hash64(b)
}
