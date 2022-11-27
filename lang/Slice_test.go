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

func TestSlice_Append(t *testing.T) {
	s := NewSlice[Int](0, 0)
	for i := 0; i < 16; i++ {
		s.Append(Int(i))
		fmt.Println(s)
	}
}
