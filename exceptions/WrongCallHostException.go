/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package exceptions

type WrongCallHostException struct {
	RuntimeException
}

func NewWrongCallHostException(message string) WrongCallHostException {
	return WrongCallHostException{
		NewRuntimeException(message, DefaultExceptionConfig().AddSkipStack(1).
			SetExceptionName("github.com.tursom.GoCollections.exceptions.WrongCallHostException")),
	}
}
