package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type Queue[T lang.Object] interface {
	// Iterator MutableIterable
	Iterator() Iterator[T]
	// MutableIterator MutableIterable
	MutableIterator() *linkedListIterator[T]

	Push(element T) exceptions.Exception
	Offer() (T, exceptions.Exception)
}
