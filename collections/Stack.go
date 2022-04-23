package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type (
	Stack[T any] interface {
		MutableIterable[T]

		Push(element T) exceptions.Exception
		PushAndGetNode(element T) (StackNode[T], exceptions.Exception)
		Pop() (T, exceptions.Exception)
	}

	StackNode[T lang.Object] interface {
		Set(value T) exceptions.Exception
		Get() (T, exceptions.Exception)
		Remove() exceptions.Exception
		RemoveAndGet() (T, exceptions.Exception)
	}
)
