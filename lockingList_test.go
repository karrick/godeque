package godeque_test

import (
	"testing"

	"github.com/karrick/godeque"
)

func newLockingList() godeque.Deque {
	return godeque.NewLockingList()
}

func TestLockingDeque(t *testing.T) {
	testDeque(t, newLockingList)
}

func BenchmarkLockingDeque(b *testing.B) {
	benchmarkDeque(b, newLockingList)
}
