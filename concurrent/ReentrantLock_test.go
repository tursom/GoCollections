/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package concurrent

import (
	"testing"
)

func TestMutex(t *testing.T) {
	mutex := &ReentrantLock{}
	mutex.Lock()
	mutex.Lock()
}
