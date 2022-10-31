package unsafe

import (
	"unsafe"

	"github.com/tursom/GoCollections/lang"
)

func Sizeof[T any]() uintptr {
	return unsafe.Sizeof(lang.Nil[T]())
}
