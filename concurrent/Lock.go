/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

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
