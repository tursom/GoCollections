package collections

import "github.com/tursom/GoCollections/exceptions"

type Stack[T any] interface {
	// Iterator MutableIterable
	Iterator() Iterator[T]
	// MutableIterator MutableIterable
	MutableIterator() MutableIterator[T]

	Push(element T) exceptions.Exception
	Pop() (T, exceptions.Exception)
}
