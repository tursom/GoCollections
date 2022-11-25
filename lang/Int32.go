/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "strconv"

type Int32 int32

type Rune = Int32

type AsInt32 interface {
	Object
	AsInt32() Int32
}

func CastInt32(v any) (int32, bool) {
	switch i := v.(type) {
	case int32:
		return i, true
	case AsInt32:
		return i.AsInt32().V(), true
	default:
		return 0, false
	}
}

func EqualsInt32(i1 AsInt32, i2 any) bool {
	i2, ok := CastInt32(i2)
	return ok && i2 == i1.AsInt32().V()
}

func (i Int32) V() int32 {
	return int32(i)
}

func (i *Int32) P() *int32 {
	return (*int32)(i)
}

func (i Int32) AsInt32() Int32 {
	return i
}

func (i Int32) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int32) AsObject() Object {
	return i
}

func (i Int32) Equals(e Object) bool {
	return EqualsInt32(i, e)
}

func (i Int32) ToString() String {
	return NewString(i.String())
}

func (i Int32) HashCode() int32 {
	return int32(i)
}

func (i Int32) Compare(t Int32) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
