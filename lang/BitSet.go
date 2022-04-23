package lang

type (
	BitSet interface {
		BitLength() int
		SetBit(bit int, up bool) (old bool)
	}
)
