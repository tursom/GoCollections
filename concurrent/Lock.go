package concurrent

type (
	Lock interface {
		Lock()
		Unlock()
	}
	RWLock interface {
		Lock()
		Unlock()
		RLock()
		RUnlock()
	}
)
