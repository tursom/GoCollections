package lang

import "unsafe"

func Nil[T any]() T {
	var n T
	return n
}

func Len[T any](array []T) int {
	return len(array)
}

func Append[T any](slice []T, elems ...T) []T {
	return append(slice, elems...)
}

func TryCast[T any](v any) (T, bool) {
	if v == nil {
		return Nil[T](), true
	} else {
		t, ok := v.(T)
		return t, ok
	}
}

func Cast[T any](v any) T {
	if v == nil {
		return Nil[T]()
	} else {
		return v.(T)
	}
}

func ForceCast[T any](v unsafe.Pointer) T {
	if v == nil {
		return Nil[T]()
	} else {
		return Cast[T](v)
	}
}
