/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package exceptions

import (
	"testing"
)

func TestRuntimeException_PrintStackTrace(t *testing.T) {
	exception := NewRuntimeException("test1", DefaultExceptionConfig().SetCause(1))
	exception = NewRuntimeException("test2", DefaultExceptionConfig().SetCause(exception))
	exception = NewRuntimeException("", DefaultExceptionConfig().SetCause(exception))
	exception.PrintStackTrace()
}
