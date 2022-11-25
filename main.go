/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package main

import (
	_ "github.com/tursom/GoCollections/collections"
	_ "github.com/tursom/GoCollections/concurrent"
	_ "github.com/tursom/GoCollections/concurrent/collections"
	"github.com/tursom/GoCollections/exceptions"
	_ "github.com/tursom/GoCollections/lang"
	_ "github.com/tursom/GoCollections/lang/atomic"
	_ "github.com/tursom/GoCollections/util"
)

func main() {
	exceptions.NewRuntimeException("test2", exceptions.DefaultExceptionConfig().SetCause(1)).PrintStackTrace()
}
