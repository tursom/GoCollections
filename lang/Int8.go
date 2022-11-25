/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "strconv"

type Int8 int8

type AsInt8 interface {
	Object
	AsInt8() Int8
}

func CastInt8(v any) (int8, bool) {
	switch i := v.(type) {
	case int8:
		return i, true
	case AsInt8:
		return i.AsInt8().V(), true
	default:
		return 0, false
	}
}

func EqualsInt8(i1 AsInt8, i2 any) bool {
	i2, ok := CastInt8(i2)
	return ok && i2 == i1.AsInt8().V()
}

func (i Int8) V() int8 {
	return int8(i)
}

func (i *Int8) P() *int8 {
	return (*int8)(i)
}

func (i Int8) AsInt8() Int8 {
	return i
}

func (i Int8) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int8) AsObject() Object {
	return i
}

func (i Int8) Equals(e Object) bool {
	return EqualsInt8(i, e)
}

func (i Int8) ToString() String {
	return NewString(i.String())
}

func (i Int8) HashCode() int32 {
	return int32(i)
}

func (i Int8) Compare(t Int8) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
