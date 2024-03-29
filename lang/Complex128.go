/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import (
	"fmt"
)

type Complex128 complex128

func (i Complex128) AsComplex128() complex128 {
	return complex128(i)
}

func (i Complex128) String() string {
	return fmt.Sprint(complex64(i))
}

func (i Complex128) AsObject() Object {
	return i
}

func (i Complex128) Equals(e Object) bool {
	i2, ok := e.(Complex128)
	if !ok {
		return false
	}
	return i == i2
}

func (i Complex128) ToString() String {
	return NewString(i.String())
}

func (i Complex128) HashCode() int32 {
	return HashFloat64(real(i)) ^ HashFloat64(imag(i))
}
