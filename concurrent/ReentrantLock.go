package concurrent

import "sync"

type ReentrantLock struct {
	sync.Mutex
	cnt int32
}

func (l *ReentrantLock) Lock() {
	defer func() {
		r := recover()
		if r != nil {
			l.cnt++
		}
	}()

	l.Mutex.Lock()
}

func (l *ReentrantLock) Unlock() {
	l.cnt--
	if l.cnt == 0 {
		l.Mutex.Unlock()
	}
}
