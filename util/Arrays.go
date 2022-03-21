package util

import (
	"github.com/tursom/GoCollections/collections"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

func AsList[T lang.Object](arr []T) collections.List[T] {
	return &arrayList[T]{arr}
}

func CheckedGet[T any](array []T, index int) (T, exceptions.Exception) {
	return exceptions.CatchIndexOutOfBound(func() T {
		return array[index]
	}, exceptions.DefaultExceptionConfig().AddSkipStack(3))
}
