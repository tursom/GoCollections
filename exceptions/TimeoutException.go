/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package exceptions

type TimeoutException struct {
	RuntimeException
}

func NewTimeoutException(message string, config *ExceptionConfig) *TimeoutException {
	return &TimeoutException{
		*NewRuntimeException(message, config.AddSkipStack(1).
			SetExceptionName("github.com.tursom.GoCollections.exceptions.TimeoutException")),
	}
}
