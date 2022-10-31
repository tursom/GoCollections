package lang

import "time"

type (
	SendChannel[T any] interface {
		Close()
		SCh() chan<- T
		Send(obj T)
		TrySend(obj T) bool
		SendTimeout(obj T, timeout time.Duration) bool
	}

	ReceiveChannel[T any] interface {
		Close()
		RCh() <-chan T
		Receive() T
		TryReceive() (T, bool)
		ReceiveTimeout(timeout time.Duration) (T, bool)
	}

	Channel[T any] interface {
		SendChannel[T]
		ReceiveChannel[T]
		Ch() chan T
	}

	RawChannel[T any] chan T

	withReceiveChannel[T any] struct {
		ReceiveChannel[T]
		closer func()
	}
)

func NewChannel[T any](cap int) RawChannel[T] {
	return make(chan T, cap)
}

func WithReceiveChannel[T any](channel ReceiveChannel[T], closer func()) ReceiveChannel[T] {
	return withReceiveChannel[T]{channel, closer}
}

func (ch RawChannel[T]) Ch() chan T {
	return ch
}

func (ch RawChannel[T]) Close() {
	close(ch)
}

func (ch RawChannel[T]) SCh() chan<- T {
	return ch
}

func (ch RawChannel[T]) Send(obj T) {
	ch <- obj
}

func (ch RawChannel[T]) TrySend(obj T) bool {
	select {
	case ch <- obj:
		return true
	default:
		return false
	}
}

func (ch RawChannel[T]) SendTimeout(obj T, timeout time.Duration) bool {
	select {
	case ch <- obj:
		return true
	case <-time.After(timeout):
		return false
	}
}

func (ch RawChannel[T]) RCh() <-chan T {
	return ch
}

func (ch RawChannel[T]) Receive() T {
	return <-ch
}

func (ch RawChannel[T]) TryReceive() (T, bool) {
	select {
	case obj := <-ch:
		return obj, true
	default:
		return Nil[T](), false
	}
}

func (ch RawChannel[T]) ReceiveTimeout(timeout time.Duration) (T, bool) {
	select {
	case obj := <-ch:
		return obj, true
	case <-time.After(timeout):
		return Nil[T](), false
	}
}

func (i withReceiveChannel[T]) Close() {
	if i.closer != nil {
		i.closer()
	} else {
		i.ReceiveChannel.Close()
	}
}
