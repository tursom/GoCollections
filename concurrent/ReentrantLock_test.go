package concurrent

import (
	"testing"
)

func TestMutex(t *testing.T) {
	mutex := &ReentrantLock{}
	mutex.Lock()
	mutex.Lock()
}
