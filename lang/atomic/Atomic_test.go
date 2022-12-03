package atomic

import (
	"fmt"
	"testing"
)

func TestGetAtomizer(t *testing.T) {
	atomizer := GetAtomizer[int32]()
	var ref *int32
	var obj int32 = 1
	atomizer.CompareAndSwap()(&ref, nil, &obj)

	fmt.Println(*ref)
}
