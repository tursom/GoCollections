/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package util

import (
	"fmt"
	"testing"

	"github.com/tursom/GoCollections/lang"
)

func Test_arrayList_Get(t *testing.T) {
	l := arrayList[lang.Int]{array: []lang.Int{1, 2}}
	for i := 0; i < l.Size()+1; i++ {
		r, err := l.Get(i)
		if err != nil {
			err.PrintStackTrace()
		} else {
			fmt.Println(r)
		}
	}
}
