/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package exceptions

import (
	"fmt"
	"reflect"
)

type PackageException struct {
	RuntimeException
	err any
}

func NewPackageException(err any, config *ExceptionConfig) *PackageException {
	message := ""
	switch e := err.(type) {
	case error:
		message = e.Error()
	default:
		message = fmt.Sprint(e)
	}
	t := reflect.TypeOf(err)
	message = fmt.Sprintf("%s (%s)", message, t.Name())
	return &PackageException{
		RuntimeException: NewRuntimeException(message, config.AddSkipStack(1).
			SetExceptionName("github.com.tursom.GoCollections.exceptions.PackageException")),
		err: err,
	}
}

func (p *PackageException) Err() any {
	return p.err
}

func UnpackException(err any) any {
	for err != nil {
		switch e := err.(type) {
		case *PackageException:
			return e.Err()
		case Exception:
			err = e.Cause()
			if err == nil {
				return e
			}
		default:
			return err
		}
	}
	return nil
}
