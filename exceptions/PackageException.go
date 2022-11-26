/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package exceptions

import (
	"github.com/tursom/GoCollections/lang"
)

type PackageException = lang.PackageException

func NewPackageException(err any, config *ExceptionConfig) *PackageException {
	return lang.NewPackageException(err, config.AddSkipStack(1))
}
func UnpackException(err any) any {
	return lang.UnpackException(err)
}
