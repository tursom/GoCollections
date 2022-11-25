/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

type Comparable[T any] interface {
	Compare(t T) int
}

func Compare[T any, E Comparable[T]](e E, t T) int {
	//if e == nil {
	//	panic(exceptions.NewNPE("LessThan compare from an null object", true))
	//}

	return e.Compare(t)
}

func LessThan[T any, E Comparable[T]](e E, t T) bool {
	return e.Compare(t) < 0
}

func MoreThan[T any, E Comparable[T]](e E, t T) bool {
	return e.Compare(t) > 0
}

func LessThanEqual[T any, E Comparable[T]](e E, t T) bool {
	return e.Compare(t) <= 0
}

func MoreThanEqual[T any, E Comparable[T]](e E, t T) bool {
	return e.Compare(t) >= 0
}
