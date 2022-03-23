package lang

import (
	"testing"
)

func TestUInt_HashCode(t *testing.T) {
	type Tests struct {
		name string
		i    UInt
		want int32
	}
	var tests []Tests
	for i := 0; i < 10; i++ {
		// -i - 1 is ^i
		// hash code of ^i equals i
		tests = append(tests, Tests{"test -1", UInt(-i - 1), int32(i)})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.HashCode(); got != tt.want {
				t.Errorf("HashCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
