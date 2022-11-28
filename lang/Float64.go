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

func CastFloat64(v any) (float64, bool) {
	switch i := v.(type) {
	case int:
		return float64(i), true
	case int8:
		return float64(i), true
	case int16:
		return float64(i), true
	case int32:
		return float64(i), true
	case int64:
		return float64(i), true
	case uint:
		return float64(i), true
	case uint8:
		return float64(i), true
	case uint16:
		return float64(i), true
	case uint32:
		return float64(i), true
	case uint64:
		return float64(i), true
	case float32:
		return float64(i), true
	case float64:
		return i, true
	case Number:
		return i.ToFloat64().V(), true
	default:
		return 0, false
	}
}

func EqualsFloat64(i1 Number, i2 any) bool {
	i2, ok := CastFloat64(i2)
	return ok && i2 == i1.ToFloat64().V()
}

func (i Float64) V() float64 {
	return float64(i)
}

func (i Float64) String() string {
	return fmt.Sprint(float64(i))
}

func (i Float64) AsObject() Object {
	return i
}

func (i Float64) Equals(e Object) bool {
	i2, ok := CastFloat64(e)
	if !ok {
		return false
	}
	return i.V() == i2
}

func (i Float64) ToString() String {
	return NewString(i.String())
}

func (i Float64) HashCode() int32 {
	return HashFloat64(float64(i))
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

func HashFloat64(f float64) int32 {
	return Hash64(&f)
}
