package util

import (
	"sync"
	"unsafe"
)

type (
	Stringer struct {
		stringer func() string
	}
)

func OnceDone(once *sync.Once) bool {
	i := *(*uint32)(unsafe.Pointer(once))
	return i != 0
}

func NewStringer(stringer func() string) Stringer {
	return Stringer{
		stringer: stringer,
	}
}

func (s Stringer) String() string {
	if s.stringer == nil {
		return "nil"
	}
	return s.stringer()
}
