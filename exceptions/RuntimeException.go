/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package exceptions

import (
	"github.com/tursom/GoCollections/lang"
)

type RuntimeException = lang.RuntimeException

func NewRuntimeException(message string, config *ExceptionConfig) *RuntimeException {
	return lang.NewRuntimeException(message, config.AddSkipStack(1))
}
