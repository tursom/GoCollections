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

func TestString_HashCode(t *testing.T) {
	for i := 1000; i < 1100; i++ {
		fmt.Println(Int(i).ToString().HashCode())
	}
}
