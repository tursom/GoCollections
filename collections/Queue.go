/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type (
	Queue[T lang.Object] interface {
		MutableIterable[T]

		Offer(element T) exceptions.Exception
		OfferAndGetNode(element T) (QueueNode[T], exceptions.Exception)
		Poll() (T, exceptions.Exception)
	}

	QueueNode[T lang.Object] StackNode[T]
)
