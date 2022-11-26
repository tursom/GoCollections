/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package exceptions

import "github.com/tursom/GoCollections/lang"

type ExceptionConfig = lang.ExceptionConfig

func Cfg() *ExceptionConfig {
	return DefaultExceptionConfig()
}

func DefaultExceptionConfig() *ExceptionConfig {
	return lang.DefaultExceptionConfig()
}
