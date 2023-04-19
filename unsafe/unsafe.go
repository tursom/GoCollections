package unsafe

import (
	"reflect"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

type iface struct {
	tab  *struct{}
	data unsafe.Pointer
}

func Nil[T any]() T {
	var n T
	return n
}

func ForceCast[T any](v unsafe.Pointer) *T {
	if v == nil {
		return nil
	} else {
		return (*T)(v)
	}
}

func Sizeof[T any]() uintptr {
	return unsafe.Sizeof(Nil[T]())
}

// AsBytes cast any slice as []byte with same pinter, real len and real cap
func AsBytes[T any](arr []T) []byte {
	sarr := *ForceCast[slice](unsafe.Pointer(&arr))
	typeAlign := reflect.TypeOf(Nil[T]()).Align()
	asBytes := unsafe.Pointer(&slice{
		array: sarr.array,
		len:   sarr.len * typeAlign,
		cap:   sarr.cap * typeAlign,
	})

	return *ForceCast[[]byte](asBytes)
}

func AsBytes1[T any](value *T) []byte {
	typeAlign := reflect.TypeOf(value).Align()
	asBytes := unsafe.Pointer(&slice{
		array: unsafe.Pointer(value),
		len:   typeAlign,
		cap:   typeAlign,
	})

	return *ForceCast[[]byte](asBytes)
}

func AsBytes2(value any) []byte {
	typeAlign := reflect.TypeOf(value).Align()
	asBytes := unsafe.Pointer(&slice{
		array: InterfacePointer(value),
		len:   typeAlign,
		cap:   typeAlign,
	})

	return *ForceCast[[]byte](asBytes)
}

func InterfacePointer(i any) unsafe.Pointer {
	return ForceCast[iface](unsafe.Pointer(&i)).data
}

func InterfaceForceCast[T any](i any) *T {
	return ForceCast[T](ForceCast[iface](unsafe.Pointer(&i)).data)
}

// AsString cast bytes as string
func AsString(bytes []byte) string {
	return *ForceCast[string](unsafe.Pointer(&bytes))
}

func IndexOf[T any](s []T, v *T) int {
	begin := *(*uintptr)(unsafe.Pointer(&s))
	addr := uintptr(unsafe.Pointer(v))
	return int((addr - begin) / reflect.TypeOf(*v).Size())
}
