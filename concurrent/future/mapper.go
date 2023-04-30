package future

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
	"github.com/tursom/GoCollections/util/time"
)

type (
	mapperFuture[V1, V2 any] struct {
		future Future[V1]
		mapper func(v1 V1) (V2, exceptions.Exception)
		v      V2
		e      exceptions.Exception
		done   bool
	}
)

func Map[V1, V2 any](
	future Future[V1],
	mapper func(v1 V1) (V2, exceptions.Exception),
) Future[V2] {
	return &mapperFuture[V1, V2]{
		future: future,
		mapper: mapper,
	}
}

func (m *mapperFuture[V1, V2]) Get() (V2, exceptions.Exception) {
	if m.done {
		return m.v, m.e
	}

	defer func() {
		m.done = true
	}()

	v1, e := m.future.Get()
	if e != nil {
		m.e = e
		return lang.Nil[V2](), e
	}

	v2, e := m.mapper(v1)
	if e != nil {
		m.e = e
		return lang.Nil[V2](), e
	}

	m.v = v2
	return v2, nil
}

func (m *mapperFuture[V1, V2]) GetT(timeout time.Duration) (V2, exceptions.Exception) {
	if m.done {
		return m.v, m.e
	}

	v1, e := m.future.GetT(timeout)
	if e != nil {
		return m.processException(e)
	}

	v2, e := m.mapper(v1)
	if e != nil {
		return m.processException(e)
	}

	m.v = v2
	return v2, nil
}

func (m *mapperFuture[V1, V2]) processException(e exceptions.Exception) (V2, exceptions.Exception) {
	if _, ok := e.(*exceptions.TimeoutException); !ok {
		m.done = true
		m.e = e
	}
	return lang.Nil[V2](), e
}

func (m *mapperFuture[V1, V2]) IsDone() bool {
	return m.future.IsDone()
}

func (m *mapperFuture[V1, V2]) IsCancelled() bool {
	return m.future.IsCancelled()
}

func (m *mapperFuture[V1, V2]) Cancel() bool {
	return m.future.Cancel()
}
