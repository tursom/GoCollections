/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package concurrent

import (
	"sync"

	"github.com/petermattis/goid"
)

func GetGoroutineID() int64 {
	return goid.Get()
}

func WaitCond(cond *sync.Cond) {
	cond.L.Lock()
	defer cond.L.Unlock()
	cond.Wait()
}
