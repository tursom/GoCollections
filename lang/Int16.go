/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "strconv"

type Int16 int16

type AsInt16 interface {
	Object
	AsInt16() Int16
}

func CastInt16(v any) (int16, bool) {
	switch i := v.(type) {
	case int16:
		return i, true
	case AsInt16:
		return i.AsInt16().V(), true
	default:
		return 0, false
	}
}

func EqualsInt16(i1 AsInt16, i2 any) bool {
	i2, ok := CastInt16(i2)
	return ok && i2 == i1.AsInt16().V()
}

func (i Int16) V() int16 {
	return int16(i)
}

func (i *Int16) P() *int16 {
	return (*int16)(i)
}

func (i Int16) AsInt16() Int16 {
	return i
}

func (i Int16) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int16) AsObject() Object {
	return i
}

func (i Int16) Equals(e Object) bool {
	return EqualsInt16(i, e)
}

func (i Int16) ToString() String {
	return NewString(i.String())
}

func (i Int16) HashCode() int32 {
	return int32(i)
}

func (i Int16) Compare(t Int16) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
