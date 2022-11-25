/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package collections

import (
	"fmt"
	"testing"

	"github.com/tursom/GoCollections/lang"
)

func TestHashIterable(t *testing.T) {
	list := NewArrayList[lang.String]()
	for i := 0; i < 201; i++ {
		list.Add(lang.Int(i).ToString())
	}
	fmt.Println(HashIterable[lang.String](list))
}
