package mr

import (
	"github.com/tursom/GoCollections/collections"
	"github.com/tursom/GoCollections/lang"
)

type MapEntry[K comparable, V any] struct {
	key   K
	value V
}

func MapChannel[K comparable, V any](m map[K]V) lang.ReceiveChannel[*MapEntry[K, V]] {
	ch := make(lang.RawChannel[*MapEntry[K, V]])
	go func() {
		for k, v := range m {
			ch <- &MapEntry[K, V]{k, v}
		}

		close(ch)
	}()

	return ch
}

func ArrayChannel[E any](arr lang.Array[E]) lang.ReceiveChannel[E] {
	ch := make(lang.RawChannel[E])
	go func() {
		for _, e := range arr {
			ch <- e
		}

		close(ch)
	}()

	return ch
}

func SliceChannel[E any](arr []E) lang.ReceiveChannel[E] {
	ch := make(lang.RawChannel[E])
	go func() {
		for _, e := range arr {
			ch <- e
		}

		close(ch)
	}()

	return ch
}

func IteratorChannel[E any](iter collections.Iterator[E]) lang.ReceiveChannel[E] {
	ch := make(lang.RawChannel[E])
	go func() {
		for iter.HasNext() {
			next, e := iter.Next()
			if e != nil {
				panic(e)
			}

			ch <- next
		}

		close(ch)
	}()

	return ch
}
