package b62

import (
	"bytes"

	"github.com/tursom/GoCollections/util"
)

var (
	digits = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
)

func Encode(n uint64) string {
	buffer := bytes.NewBuffer(nil)

	for n != 0 {
		remainder := n % 62
		buffer.WriteByte(digits[remainder])
		n /= 62
	}

	return string(util.Reverse(buffer.Bytes()))
}

func Decode(str string) uint64 {
	var sum uint64
	var base uint64 = 1

	for i := range str {
		indexByte := bytes.IndexByte(digits, str[len(str)-i-1])
		if indexByte < 0 || indexByte >= 62 {
			panic(indexByte)
		}
		sum += uint64(indexByte) * base
		base *= 62
	}

	return sum
}
