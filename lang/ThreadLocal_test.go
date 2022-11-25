/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import (
	"fmt"
	"testing"
)

func TestThreadLocalImpl(t1 *testing.T) {
	local := NewThreadLocal[int]()
	fmt.Println(local.Get())
	local.Put(1)
	fmt.Println(local.Get())
	local.Remove()
	fmt.Println(local.Get())
}
