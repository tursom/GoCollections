package hex

import (
	"encoding/hex"

	"github.com/tursom/GoCollections/unsafe"
)

func EncodeToString[T any](arr []T) string {
	return hex.EncodeToString(unsafe.AsBytes(arr))
}

func DecodeString(s string) ([]byte, error) {
	return hex.DecodeString(s)
}
