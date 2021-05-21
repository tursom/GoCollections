package collections

import "github.com/tursom/GoCollections/exceptions"

type Queue interface {
	// Iterator MutableIterable
	Iterator() Iterator
	// MutableIterator MutableIterable
	MutableIterator() MutableIterator

	Push(element interface{}) exceptions.Exception
	Offer() (interface{}, exceptions.Exception)
}
