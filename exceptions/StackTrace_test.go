/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package exceptions

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestGetStackTrace(t *testing.T) {
	fmt.Println(unsafe.Sizeof(StackTrace{}))
	fmt.Println(unsafe.Sizeof(make([]StackTrace, 0, 16)))
	fmt.Println(unsafe.Sizeof(make([]StackTrace, 0, 16)) + unsafe.Sizeof(StackTrace{})*16)
}
