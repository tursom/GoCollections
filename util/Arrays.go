/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package util

import (
	"github.com/tursom/GoCollections/collections"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

func AsList[T lang.Object](arr []T) collections.List[T] {
	return &arrayList[T]{array: arr}
}

func CheckedGet[T any](array []T, index int) (T, exceptions.Exception) {
	return exceptions.CatchIndexOutOfBound(func() T {
		return array[index]
	}, exceptions.Cfg().AddSkipStack(3))
}
