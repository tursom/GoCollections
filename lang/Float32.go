/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import (
	"fmt"
)

type Float32 float32

func CastFloat32(v any) (float32, bool) {
	switch i := v.(type) {
	case int:
		return float32(i), true
	case int8:
		return float32(i), true
	case int16:
		return float32(i), true
	case int32:
		return float32(i), true
	case int64:
		return float32(i), true
	case uint:
		return float32(i), true
	case uint8:
		return float32(i), true
	case uint16:
		return float32(i), true
	case uint32:
		return float32(i), true
	case uint64:
		return float32(i), true
	case float32:
		return i, true
	case float64:
		return float32(i), true
	case Number:
		return float32(i.ToFloat64().V()), true
	default:
		return 0, false
	}
}

func EqualsFloat32(i1 Number, i2 any) bool {
	i2, ok := CastFloat32(i2)
	return ok && i2 == i1.ToFloat64().V()
}

func (i Float32) V() float32 {
	return float32(i)
}

func (i Float32) String() string {
	return fmt.Sprint(float32(i))
}

func (i Float32) AsObject() Object {
	return i
}

func (i Float32) Equals(e Object) bool {
	return EqualsFloat32(i, e)
}

func (i Float32) ToString() String {
	return NewString(i.String())
}

func (i Float32) HashCode() int32 {
	return HashFloat32(float32(i))
}

func (i Float32) Compare(t Float32) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}

func (i Float32) ToInt64() Int64 {
	return Int64(i)
}

func (i Float32) ToUInt64() UInt64 {
	return UInt64(i)
}

func (i Float32) ToFloat64() Float64 {
	return Float64(i)
}

func HashFloat32(f float32) int32 {
	return Hash32(&f)
}
