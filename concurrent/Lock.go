package concurrent

type (
	Lock interface {
		Lock()
		Unlock()
	}
	RWLock interface {
		Lock
		RLock()
		RUnlock()
	}
)

func WithLock[R any](lock Lock, f func() R) R {
	lock.Lock()
	defer lock.Unlock()
	return f()
}

func WithRLock[R any](lock RWLock, f func() R) R {
	lock.RLock()
	defer lock.RUnlock()
	return f()
}
