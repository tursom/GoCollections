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

func TestGoMap_get(t *testing.T) {
	m := NewGoMap[lang.Int, int]()
	for i := 0; i < 10; i++ {
		m.Put(lang.Int(i), i+1)
		fmt.Println(m.ToString().String())
		fmt.Println(m)
	}
}
