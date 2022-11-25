/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package collections

import (
	"github.com/tursom/GoCollections/exceptions"
)

type (
	Stack[T any] interface {
		MutableIterable[T]

		Push(element T) exceptions.Exception
		PushAndGetNode(element T) (StackNode[T], exceptions.Exception)
		Pop() (T, exceptions.Exception)
	}

	StackNode[T any] interface {
		Set(value T) exceptions.Exception
		Get() (T, exceptions.Exception)
		Remove() exceptions.Exception
		RemoveAndGet() (T, exceptions.Exception)
	}
)
