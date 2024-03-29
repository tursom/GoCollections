/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package concurrent

import (
	"fmt"
	"sync"

	"github.com/tursom/GoCollections/exceptions"
)

type ReentrantLock struct {
	lock      sync.Mutex
	cond      sync.Cond
	recursion int32
	host      int64
}

func NewReentrantLock() *ReentrantLock {
	res := &ReentrantLock{
		recursion: 0,
		host:      0,
	}
	res.cond = *sync.NewCond(&res.lock)
	return res
}

func (rt *ReentrantLock) TryLock() bool {
	id := GetGoroutineID()
	rt.lock.Lock()
	defer rt.lock.Unlock()

	if rt.host == id {
		rt.recursion++
		return true
	}

	if rt.recursion == 0 {
		rt.host = id
		rt.recursion = 1
		return true
	}

	return false
}

func (rt *ReentrantLock) Lock() {
	id := GetGoroutineID()
	rt.lock.Lock()
	defer rt.lock.Unlock()

	if rt.host == id {
		rt.recursion++
		return
	}

	for rt.recursion != 0 {
		rt.cond.Wait()
	}
	rt.host = id
	rt.recursion = 1
}

func (rt *ReentrantLock) Unlock() {
	rt.lock.Lock()
	defer rt.lock.Unlock()

	if rt.recursion == 0 || rt.host != GetGoroutineID() {
		panic(exceptions.NewWrongCallHostException(fmt.Sprintf("the wrong call host: (%d); current_id: %d; recursion: %d", rt.host, GetGoroutineID(), rt.recursion)))
	}

	rt.recursion--
	if rt.recursion == 0 {
		rt.cond.Signal()
	}
}
