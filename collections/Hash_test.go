package collections

import (
	"fmt"
	"github.com/tursom/GoCollections/lang"
	"testing"
)

func TestHashIterable(t *testing.T) {
	list := NewArrayList[lang.String]()
	for i := 0; i < 201; i++ {
		list.Add(lang.Int(i).ToString())
	}
	fmt.Println(HashIterable[lang.String](list))
}
