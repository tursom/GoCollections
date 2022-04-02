package util

import (
	"fmt"
	"github.com/tursom/GoCollections/lang"
	"testing"
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
