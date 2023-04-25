package bloom

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
	"unsafe"

	"github.com/spaolacci/murmur3"

	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

var (
	HashFunc = murmur3.Sum32WithSeed
)

type (
	Bloom struct {
		lang.BaseObject
		m        lang.UInt8Array
		k        uint
		c        uint
		hashCode int32
	}
)

func max(i1, i2 uint) uint {
	if i1 < i2 {
		return i2
	} else {
		return i1
	}
}

func numHashFunctions(n, m float64) uint {
	return max(1, uint(math.Floor(0.5+m/n*math.Ln2)))
}

func NumHashFunctions(n, m uint) uint {
	return numHashFunctions(float64(n), float64(m))
}

func calcBitLength(n float64, p float64) uint {
	if p == 0 {
		p = math.SmallestNonzeroFloat64
	}
	return uint(-n * math.Log(p) / (math.Ln2 * math.Ln2))
}

func CalcBitLength(n uint, p float64) uint {
	return calcBitLength(float64(n), p)
}

func NewBloom(n uint, p float64) *Bloom {
	m := CalcBitLength(n, p) - 1
	return &Bloom{
		m: make(lang.UInt8Array, m/8+1),
		k: NumHashFunctions(n, m),
	}
}

func (b *Bloom) C() uint {
	return b.c
}

func (b *Bloom) K() uint {
	return b.k
}

func (b *Bloom) M() uint {
	return uint(len(b.m)) * 8
}

func (b *Bloom) Contains(data []byte) bool {
	for i := 0; i < int(b.k); i++ {
		hashCode := uint(HashFunc(data, uint32(i)))
		if !b.m.GetBit(hashCode % b.m.BitLength()) {
			return false
		}
	}

	return true
}

func (b *Bloom) Add(data []byte) {
	b.hashCode = 0
	b.c++
	for i := 0; i < int(b.k); i++ {
		hashCode := uint(HashFunc(data, uint32(i)))
		b.m.SetBit(hashCode%b.m.BitLength(), true)
	}
}

func (b *Bloom) Marshal(writer io.Writer) {
	if err := binary.Write(writer, binary.BigEndian, uint32(b.k)); err != nil {
		panic(exceptions.Package(err))
	}
	if err := binary.Write(writer, binary.BigEndian, uint32(b.c)); err != nil {
		panic(exceptions.Package(err))
	}

	if _, err := writer.Write(b.m.Bytes()); err != nil {
		panic(exceptions.Package(err))
	}
}

func Unmarshal(data []byte) *Bloom {
	k := binary.BigEndian.Uint32(data)
	c := binary.BigEndian.Uint32(data[4:])

	m := data[8:]

	return &Bloom{
		m: *(*lang.UInt8Array)(unsafe.Pointer(&m)),
		k: uint(k),
		c: uint(c),
	}
}

func (b *Bloom) Equals(t lang.Object) bool {
	tb, ok := t.(*Bloom)
	if !ok {
		return false
	}

	return tb.k == b.k && bytes.Compare(b.m.Bytes(), tb.m.Bytes()) == 0
}

func (b *Bloom) HashCode() int32 {
	if b.hashCode == 0 {
		b.hashCode = int32(murmur3.Sum32(b.m.Bytes()))
	}
	return b.hashCode
}

func (b *Bloom) Merge(t *Bloom) bool {
	if b.k != t.k {
		return false
	} else if len(b.m) != len(t.m) {
		return false
	}

	b.hashCode = 0
	for i := range b.m {
		b.m[i] |= t.m[i]
	}

	return true
}

func (b *Bloom) MergeBM(bitMap []byte) bool {
	if len(b.m) != len(bitMap) {
		return false
	}

	b.hashCode = 0
	for i := range b.m {
		b.m[i] |= lang.UInt8(bitMap[i])
	}

	return true
}
