package collections

import (
	"fmt"
	"github.com/tursom/GoCollections/lang"
	"testing"
)

func TestGoMap_get(t *testing.T) {
	m := NewGoMap[lang.Int, int]()
	for i := 0; i < 10; i++ {
		m.Put(lang.Int(i), i+1)
		fmt.Println(m.ToString().String())
		fmt.Println(m)
	}
}
