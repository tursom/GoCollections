package bloom

import (
	"math"

	"github.com/spaolacci/murmur3"

	"github.com/tursom/GoCollections/lang"
)

var (
	HashFunc = murmur3.Sum32WithSeed
)

type (
	Bloom struct {
		m lang.UInt8Array
		k uint
		c uint
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
	b.c++
	for i := 0; i < int(b.k); i++ {
		hashCode := uint(HashFunc(data, uint32(i)))
		b.m.SetBit(hashCode%b.m.BitLength(), true)
	}
}
