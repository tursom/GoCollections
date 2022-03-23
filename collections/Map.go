package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type Map[K lang.Object, V any] interface {
	Put(k K, v V) (bool, exceptions.Exception)
	Get(k K) (V, exceptions.Exception)
}
