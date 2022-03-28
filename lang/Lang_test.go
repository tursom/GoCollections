package lang

import (
	"fmt"
	"testing"
)

func TestCast(t *testing.T) {
	fmt.Println(Cast[int](1))
	fmt.Println(Cast[int](nil))
}
