package util

import "github.com/tursom/GoCollections/concurrent/collections"

type (
	pipeline[T, R any] struct {
		value  T
		sender collections.SequenceSender[R]
	}
)

func NewPipeline[T, R any](producer <-chan T, concurrency int, consumer func(T) R) <-chan R {
	var sequence collections.Sequence[R]
	c := make(chan pipeline[T, R])
	go func() {
		defer close(c)
		for value := range producer {
			c <- pipeline[T, R]{
				value:  value,
				sender: sequence.Alloc(),
			}
		}
	}()

	for i := 0; i < concurrency; i++ {
		go func() {
			defer sequence.Close()
			for value := range c {
				r := consumer(value.value)
				value.sender.Send(r)
			}
		}()
	}
	return sequence.Channel()
}
