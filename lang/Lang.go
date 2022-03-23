package lang

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
