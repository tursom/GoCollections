package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type (
	Queue[T lang.Object] interface {
		MutableIterable[T]

		Offer(element T) exceptions.Exception
		OfferAndGetNode(element T) (QueueNode[T], exceptions.Exception)
		Poll() (T, exceptions.Exception)
	}

	QueueNode[T lang.Object] = StackNode[T]
)
