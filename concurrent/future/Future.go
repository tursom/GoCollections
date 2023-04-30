package future

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
	"github.com/tursom/GoCollections/util/time"
)

type (
	Future[T any] interface {
		Get() (T, exceptions.Exception)
		GetT(timeout time.Duration) (T, exceptions.Exception)
		IsDone() bool
		IsCancelled() bool
		Cancel() bool
	}

	Task[T any] interface {
		Future[T]
		Done(value T)
		Fail(e exceptions.Exception)
	}

	futureState uint8

	impl[T any] struct {
		state     futureState
		value     T
		e         exceptions.Exception
		done      chan futureState
		canceller func() bool
	}
)

//goland:noinspection GoSnakeCaseUsage
const (
	futureState_start    futureState = iota
	futureState_done     futureState = iota
	futureState_canceled futureState = iota
)

func Get[T any](future Future[T]) T {
	get, e := future.Get()
	if e != nil {
		panic(e)
	}

	return get
}

func GetT[T any](future Future[T], timeout time.Duration) T {
	get, e := future.GetT(timeout)
	if e != nil {
		panic(e)
	}

	return get
}

func New[T any](canceller func() bool) Task[T] {
	return &impl[T]{
		done:      make(chan futureState, 1),
		canceller: canceller,
	}
}

func (i *impl[T]) Done(value T) {
	if i.IsDone() {
		return
	}

	i.value = value
	i.state = futureState_done
	i.done <- futureState_done
}

func (i *impl[T]) Fail(e exceptions.Exception) {
	if i.IsDone() {
		return
	}

	i.e = e
	i.state = futureState_done
	i.done <- futureState_done
}

func (i *impl[T]) Get() (T, exceptions.Exception) {
	if i.IsDone() {
		return i.value, i.e
	}

	<-i.done

	return i.value, i.e
}

func (i *impl[T]) GetT(timeout time.Duration) (T, exceptions.Exception) {
	if i.IsDone() {
		return i.value, i.e
	}

	select {
	case i.state = <-i.done:
	case <-time.After(timeout):
		return lang.Nil[T](), exceptions.NewTimeoutException(
			"Future.GetT timeout", nil)
	}

	return i.value, i.e
}

func (i *impl[T]) IsDone() bool {
	return i.state != futureState_start
}

func (i *impl[T]) IsCancelled() bool {
	return i.state == futureState_canceled
}

func (i *impl[T]) Cancel() bool {
	if i.IsDone() {
		return false
	}

	if !i.canceller() {
		return false
	}

	i.state = futureState_canceled
	i.done <- futureState_canceled
	return true
}
