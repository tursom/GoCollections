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

func TestHashAddr(t *testing.T) {
	str := "test"
	fmt.Println(HashAddr(&str))
}

func TestHashString(t *testing.T) {
	fmt.Println(HashString("testwerwefdcsd"))
}
