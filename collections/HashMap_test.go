package collections

import (
	"fmt"
	"github.com/tursom/GoCollections/lang"
	"testing"
)

func TestHashMap_put(t *testing.T) {
	m := NewHashMap[lang.Int, int]()
	for i := 0; i < 100; i++ {
		fmt.Println(m)
		m.Put(lang.Int(i), i+1)
	}
	fmt.Println(m)
	m.resize()
	fmt.Println(m)
}
