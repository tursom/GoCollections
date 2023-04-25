package mr

import "github.com/tursom/GoCollections/lang/atomic"

type (
	MapReduce[V, R any] interface {
		Map(value V) R
		Reduce(results <-chan R) R
	}
)

func LocalMap[V, R any](values <-chan V, m func(value V) R) <-chan R {
	rc := make(chan R)

	go func() {
		for value := range values {
			rc <- m(value)
		}

		close(rc)
	}()

	return rc
}

func LocalReduce[R any](values <-chan R, r func(results <-chan R) R) R {
	return r(values)
}

func Local[V, R any](values <-chan V, mr MapReduce[V, R]) R {
	return LocalReduce(LocalMap(values, mr.Map), mr.Reduce)
}

func MultiMap[V, R any](values <-chan V, m func(value V) R) <-chan R {
	rc := make(chan R)

	c := atomic.Int32(1)
	for value0 := range values {
		c.Add(1)
		value := value0
		go func() {
			rc <- m(value)

			if c.Add(-1) == 0 {
				close(rc)
			}
		}()
	}
	if c.Add(-1) == 0 {
		close(rc)
	}

	return rc
}

func MultiReduce[R any](values <-chan R, r func(results <-chan R) R) R {
	rc := make(chan R)
	go func() {
		rc <- r(values)
	}()

	return <-rc
}

func Multi[V, R any](values <-chan V, mr MapReduce[V, R]) R {
	return MultiReduce(MultiMap(values, mr.Map), mr.Reduce)
}
