package lang

import (
	"fmt"
	"testing"
)

func TestString_HashCode(t *testing.T) {
	for i := 1000; i < 1100; i++ {
		fmt.Println(Int(i).ToString().HashCode())
	}
}
