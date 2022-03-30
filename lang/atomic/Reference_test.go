package atomic

import (
	"fmt"
	"github.com/tursom/GoCollections/lang"
	"testing"
)

func TestAtomic_Store(t *testing.T) {
	a := &Reference[lang.Int]{}
	var i lang.Int = 1
	a.Store(&i)
	i = 2
	fmt.Println(a.Load())
}
