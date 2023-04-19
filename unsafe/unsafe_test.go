package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestAsBytes1(t *testing.T) {
	i := 1
	var a any = i
	bs := AsBytes2(i)
	fmt.Println(i, bs, AsBytes1(&a), *(*int)(ForceCast[iface](unsafe.Pointer(&a)).data))
}
