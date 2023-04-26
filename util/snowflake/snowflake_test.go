package snowflake

import (
	"sync"
	"testing"

	"github.com/tursom/GoCollections/util/time"
)

func Test_snowflake_Close(t *testing.T) {
	s := New(0)
	s.Close()
	s.Close()
}

func Test_snowflake_New(t *testing.T) {
	s := New(1)

	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()

	var wg sync.WaitGroup

	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			old := s.New()
			for i := 0; i < 10_000_000; i++ {
				n := s.New()
				if n <= old {
					panic(n)
				}
				old = n
			}

			wg.Done()
		}()
	}

	wg.Wait()
}
