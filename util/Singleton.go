package util

import "sync"

// Singleton DCL singleton implement
type Singleton[T any] struct {
	value T
	init  func() T
	lock  sync.Mutex
}

func NewSingleton[T any](init func() T) *Singleton[T] {
	if init == nil {
		panic("nil singleton initializer")
	}

	return &Singleton[T]{
		init: init,
	}
}

func (s *Singleton[T]) Get() T {
	if s.init != nil {
		s.lock.Lock()
		s.lock.Unlock()
		if s.init != nil {
			s.value = s.init()
			s.init = nil
		}
	}

	return s.value
}
