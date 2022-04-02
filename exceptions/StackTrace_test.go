package exceptions

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestGetStackTrace(t *testing.T) {
	fmt.Println(unsafe.Sizeof(StackTrace{}))
	fmt.Println(unsafe.Sizeof(make([]StackTrace, 0, 16)))
	fmt.Println(unsafe.Sizeof(make([]StackTrace, 0, 16)) + unsafe.Sizeof(StackTrace{})*16)
}
