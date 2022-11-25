/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package concurrent

import (
	"sync"
	"unsafe"
)

type (
	Once sync.Once
)

func (o *Once) Do(f func()) {
	(*sync.Once)(o).Do(f)
}

func (o *Once) IsDone() bool {
	i := *(*uint32)(unsafe.Pointer(o))
	return i != 0
}
