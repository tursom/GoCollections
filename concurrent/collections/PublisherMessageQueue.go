/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package collections

import (
	"log"
	"sync"

	"github.com/tursom/GoCollections/util/time"

	"github.com/tursom/GoCollections/concurrent"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
	"github.com/tursom/GoCollections/lang/atomic"
)

type (
	// PublisherMessageQueue
	// Enable an application to announce events to multiple interested consumers asynchronously,
	// without coupling the senders to the receivers
	PublisherMessageQueue[T any] struct {
		end  *publisherMessageQueueNode[T]
		lock sync.Mutex
		cond concurrent.Cond
	}

	publisherMessageQueueNode[T any] struct {
		index int
		value T
		next  *publisherMessageQueueNode[T]
	}
)

func (q *PublisherMessageQueue[T]) getEnd() *publisherMessageQueueNode[T] {
	if q.end == nil {
		q.lock.Lock()
		defer q.lock.Unlock()
		if q.end == nil {
			q.end = &publisherMessageQueueNode[T]{}
		}
	}
	return q.end
}

func (q *PublisherMessageQueue[T]) getCond() concurrent.Cond {
	if q.cond == nil {
		q.lock.Lock()
		defer q.lock.Unlock()
		q.cond = concurrent.NewCond(&q.lock)
	}
	return q.cond
}

func (q *PublisherMessageQueue[T]) Subscribe() lang.ReceiveChannel[T] {
	end := q.getEnd()
	ch := lang.NewChannel[T](0)
	canceled := false
	go func() {
		defer ch.Close()

		cond := q.getCond()
		node := &end.next
		for !canceled {
			// node may be nil when MQ created
			for *node != nil {
				if canceled {
					return
				}
				for !ch.SendTimeout((*node).value, time.Second) && !canceled {
					// check MessageQueueCapacity
					if MessageQueueCapacity != -1 {
						continue
					}
					diff := q.end.index - (*node).index
					if diff >= MessageQueueWarnLimit {
						log.Printf("MD is on warn stack")
					}
					if diff > MessageQueueCapacity {
						panic(exceptions.NewIndexOutOfBound("object buffer of this MQ is full", nil))
					}
				}
				node = &(*node).next
			}
			cond.Wait()
		}
	}()
	return lang.WithReceiveChannel[T](ch, func() {
		canceled = true
	})
}

func (q *PublisherMessageQueue[T]) Send(msg T) {
	index := 0
	if q.end != nil {
		index = q.end.index + 1
	}

	node := &publisherMessageQueueNode[T]{
		index: index,
		value: msg,
	}

	p := &q.getEnd().next
	for !atomic.CompareAndSwapPointer(p, nil, node) {
		for *p != nil {
			p = &q.end.next
		}
		node.index = q.end.index + 1
	}
	q.end = node
	q.getCond().Broadcast()
}
