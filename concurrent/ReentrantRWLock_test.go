/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package concurrent

import (
	"fmt"
	"testing"

	"github.com/tursom/GoCollections/util/time"
)

func TestReentrantRWLock_RLock(t *testing.T) {
	lock := NewReentrantRWLock()
	lock.Lock()
	defer lock.Unlock()

	go func() {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("get lock")
	}()
	time.Sleep(time.Second)
	lock.Lock()
	defer lock.Unlock()
	lock.RLock()
	defer lock.RUnlock()
	lock.RLock()
	defer lock.RUnlock()
	fmt.Println("release lock")
}
