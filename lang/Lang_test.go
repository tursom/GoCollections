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

func TestCast(t *testing.T) {
	fmt.Println(Cast[int](1))
	fmt.Println(Cast[int](nil))
}
