package concurrent

import (
	"sync"
	"unsafe"
)

type (
	Once sync.Once
)

func (o *Once) Do(f func()) {
	(*sync.Once)(o).Do(f)
}

func (o *Once) IsDone() bool {
	i := *(*uint32)(unsafe.Pointer(o))
	return i != 0
}
