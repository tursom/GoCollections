package collections

import (
	"sync"

	"github.com/tursom/GoCollections/util/time"
)

type Park struct {
	lock sync.Mutex
	ch   chan struct{}
}

func (p *Park) getCh() chan struct{} {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.ch == nil {
		p.ch = make(chan struct{})
	}

	return p.ch
}

func (p *Park) Park() {
	<-p.getCh()
}

func (p *Park) ParkT(timeout time.Duration) {
	select {
	case <-p.getCh():
	case <-time.After(timeout):
	}
}

func (p *Park) Unpark() {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.ch == nil {
		return
	}

	close(p.ch)
	p.ch = nil
}
