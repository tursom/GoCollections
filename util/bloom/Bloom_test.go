/*
 * Copyright (c) 2023 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package bloom

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"testing"
)

func TestBloom_Contains(t *testing.T) {
	bloom := NewBloom(100_0000, 0.1)
	bloom.Add([]byte("hello, world!"))
	for i := 0; i < 10000; i++ {
		bloom.Add([]byte(fmt.Sprintf("%d", i)))
	}

	fmt.Println(len(bloom.m))
	fmt.Println(len(gz(bloom.m.Bytes())))

	if !bloom.Contains([]byte("hello, world!")) {
		t.Failed()
	}

	if bloom.Contains([]byte("hello, not world!")) {
		t.Failed()
	}
}

func gz(b []byte) []byte {
	buffer := bytes.NewBuffer(nil)

	writer := gzip.NewWriter(buffer)
	writer.Write(b)
	writer.Flush()

	return buffer.Bytes()
}

func TestCalcBitLength(t *testing.T) {
	//fmt.Printf("%d\n", CalcBitLength(100_0000, 0.1)/8)
	for i := 1; i < 63; i++ {
		var n uint = 1 << i
		numBytes := CalcBitLength(n, 0.1) / 8
		fmt.Printf("%d: %d, %s / %s = %f\n",
			i,
			NumHashFunctions(n, numBytes*8),
			storageFormat(numBytes),
			storageFormat(n),
			float64(numBytes)/float64(n))
	}
}

func storageFormat(size uint) string {
	var base uint = 1
	if size < 1024*base {
		return fmt.Sprintf("%db", size/base)
	}

	base *= 1024
	if size < 1024*base {
		return fmt.Sprintf("%fkb", float64(size)/float64(base))
	}

	base *= 1024
	if size < 1024*base {
		return fmt.Sprintf("%fmb", float64(size)/float64(base))
	}

	base *= 1024
	if size < 1024*base {
		return fmt.Sprintf("%fgb", float64(size)/float64(base))
	}

	base *= 1024
	return fmt.Sprintf("%ftb", float64(size)/float64(base))
}
