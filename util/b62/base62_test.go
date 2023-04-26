package b62

import (
	"math"
	"testing"
)

func TestBase62(t *testing.T) {
	encode := Encode(math.MaxUint64)
	if Decode(encode) != math.MaxUint64 {
		t.Fatal(Decode(encode))
	}
}
