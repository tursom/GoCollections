/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package exceptions

type IllegalParameterException struct {
	RuntimeException
}

func NewIllegalParameterException(message string, config *ExceptionConfig) *IllegalParameterException {
	return &IllegalParameterException{
		NewRuntimeException(message, config.AddSkipStack(1).
			SetExceptionName("github.com.tursom.GoCollections.exceptions.IllegalParameterException")),
	}
}
