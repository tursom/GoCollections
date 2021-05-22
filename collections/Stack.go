package collections

import "github.com/tursom/GoCollections/exceptions"

type Stack interface {
	// Iterator MutableIterable
	Iterator() Iterator
	// MutableIterator MutableIterable
	MutableIterator() MutableIterator

	Push(element interface{}) exceptions.Exception
	Pop() (interface{}, exceptions.Exception)
}
