package lang

import (
	"fmt"
	"testing"
)

func TestHashAddr(t *testing.T) {
	str := "test"
	fmt.Println(HashAddr(&str))
}

func TestHashString(t *testing.T) {
	fmt.Println(HashString("testwerwefdcsd"))
}
