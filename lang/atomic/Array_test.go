package atomic

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/tursom/GoCollections/encoding/hex"
	"github.com/tursom/GoCollections/lang"
)

func TestNewInt32Array_SetBit(t *testing.T) {
	fmt.Println(reflect.TypeOf(lang.Nil[atomizerImpl[int]]()).Align())

	array := NewInt32Array(16)
	for i := 0; i < 16; i++ {
		array.SetBit(10*i, true)
	}
	fmt.Println(array)
	fmt.Println(hex.EncodeToString(array))
}
