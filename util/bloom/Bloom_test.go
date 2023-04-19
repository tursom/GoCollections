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
	"math"
	"testing"
	"time"
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

func TestBloom_miss(t *testing.T) {
	//HashFunc = func(data []byte, seed uint32) uint32 {
	//	return murmur3.Sum32WithSeed(data, seed)
	//	//	h1, _ := murmur3.Sum128WithSeed(data, seed)
	//	//	return uint32(h1)
	//}

	var base uint = 1000_0000
	bloom := NewBloom(base, 0.03)

	t1 := time.Now()

	for i := 0; i < int(base/1000); i++ {
		bloom.Add([]byte(fmt.Sprintf("%d", i)))
	}

	counter := make([]uint, 256)
	for _, value := range bloom.m {
		counter[value]++
	}

	miss := 0
	for i := base; i < base*2; i++ {
		if bloom.Contains([]byte(fmt.Sprintf("%d", i))) {
			miss += 1
		}
	}

	t2 := time.Now()

	fmt.Println(miss, float64(miss)/float64(base))
	fmt.Println(counter)

	var H float64
	for _, c := range counter {
		if c == 0 {
			continue
		}
		p := float64(c) / float64(len(bloom.m))
		H += -p * math.Log2(p)
	}
	fmt.Println(H / 8)
	fmt.Println(t2.Sub(t1))
}

func gz(b []byte) []byte {
	buffer := bytes.NewBuffer(nil)

	writer := gzip.NewWriter(buffer)
	writer.Write(b)
	writer.Flush()

	return buffer.Bytes()
}

func TestCalcBitLength(t *testing.T) {
	//fmt.Printf("%d\n", CalcBitLength(1024*1024*1024, 0.03)/8)
	for i := 0; i < 63; i++ {
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
