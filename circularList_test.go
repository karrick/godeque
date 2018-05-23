package godeque_test

import (
	"testing"

	"github.com/karrick/godeque"
)

func newCircularList() godeque.Deque {
	return new(godeque.CircularList)
}

func TestCircularList(t *testing.T) {
	testDeque(t, newCircularList)
}

func BenchmarkCircularList(b *testing.B) {
	benchmarkDeque(b, newCircularList)
}
