/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import (
	"fmt"
)

type Float64 float64

func (i Float64) AsFloat64() float64 {
	return float64(i)
}

func (i Float64) String() string {
	return fmt.Sprint(float64(i))
}

func (i Float64) AsObject() Object {
	return i
}

func (i Float64) Equals(e Object) bool {
	i2, ok := e.(Float64)
	if !ok {
		return false
	}
	return i == i2
}

func (i Float64) ToString() String {
	return NewString(i.String())
}

func (i Float64) HashCode() int32 {
	return Hash64(&i)
}

func (i Float64) Compare(t Float64) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
