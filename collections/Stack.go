package collections

import "github.com/tursom/GoCollections/exceptions"

type Stack[T any] interface {
	MutableIterable[T]

	Push(element T) exceptions.Exception
	Pop() (T, exceptions.Exception)
}
