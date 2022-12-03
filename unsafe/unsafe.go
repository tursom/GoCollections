package unsafe

import (
	"reflect"
	"unsafe"

	"github.com/tursom/GoCollections/lang"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func ForceCast[T any](v unsafe.Pointer) *T {
	if v == nil {
		return nil
	} else {
		return (*T)(v)
	}
}

func Sizeof[T any]() uintptr {
	return unsafe.Sizeof(lang.Nil[T]())
}

// AsBytes cast any slice as []byte with same pinter, real len and real cap
func AsBytes[T any](arr []T) []byte {
	sarr := *ForceCast[slice](unsafe.Pointer(&arr))
	typeAlign := reflect.TypeOf(lang.Nil[T]()).Align()
	asBytes := unsafe.Pointer(&slice{
		array: sarr.array,
		len:   sarr.len * typeAlign,
		cap:   sarr.cap * typeAlign,
	})

	return *ForceCast[[]byte](asBytes)
}

// AsString cast bytes as string
func AsString(bytes []byte) string {
	return *ForceCast[string](unsafe.Pointer(&bytes))
}
