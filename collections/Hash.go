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

func HashIterable[E lang.Object](iterable Iterable[E]) int32 {
	hashCode := int32(0)
	_ = Loop(iterable, func(element E) exceptions.Exception {
		hashCode = hashCode*31 ^ element.HashCode()
		return nil
	})
	return hashCode
}
