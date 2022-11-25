/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package unsafe

import (
	"unsafe"

	"github.com/tursom/GoCollections/lang"
)

func Sizeof[T any]() uintptr {
	return unsafe.Sizeof(lang.Nil[T]())
}
