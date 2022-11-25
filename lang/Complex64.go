/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import (
	"fmt"
)

type Complex64 complex64

func (i Complex64) AsComplex64() complex64 {
	return complex64(i)
}

func (i Complex64) String() string {
	return fmt.Sprint(complex64(i))
}

func (i Complex64) AsObject() Object {
	return i
}

func (i Complex64) Equals(e Object) bool {
	i2, ok := e.(Complex64)
	if !ok {
		return false
	}
	return i == i2
}

func (i Complex64) ToString() String {
	return NewString(i.String())
}

func (i Complex64) HashCode() int32 {
	return Float32(real(complex64(i))).HashCode() ^ Float32(imag(complex64(i))).HashCode()
}
