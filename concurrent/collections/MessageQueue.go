/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package collections

import (
	"sync"

	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

var (
	// MessageQueueCapacity message capacity of MQ
	// -1 to unlimited
	// this variable can let you discover problems before OOM crash
	MessageQueueCapacity  = 128
	MessageQueueWarnLimit = MessageQueueCapacity / 2
)

type (
	MessageQueue[T any] interface {
		// Subscribe subscribe this message queue
		Subscribe() lang.ReceiveChannel[T]
		Send(msg T)
	}
	MessageQueueImpl[T lang.Object] struct {
		chLock sync.Mutex
		ch     lang.Channel[T]
	}
)

func (m *MessageQueueImpl[T]) checkCh() {
	if m.ch != nil {
		return
	}

	m.chLock.Lock()
	defer m.chLock.Unlock()

	if m.ch != nil {
		return
	}

	m.ch = lang.NewChannel[T](MessageQueueCapacity)
}

func (m *MessageQueueImpl[T]) Subscribe() lang.ReceiveChannel[T] {
	m.checkCh()
	// package ch, remove closer to empty body
	// closer is nil will just close this channel
	return lang.WithReceiveChannel[T](m.ch, func() {})
}

func (m *MessageQueueImpl[T]) Send(msg T) {
	m.checkCh()
	if !m.ch.TrySend(msg) {
		panic(exceptions.NewIndexOutOfBound("object buffer of this MQ is full", nil))
	}
}
