package godeque_test

import (
	"testing"

	"github.com/karrick/godeque"
)

func newList() godeque.Deque {
	return new(godeque.List)
}

func TestList(t *testing.T) {
	testDeque(t, newList)
}

func BenchmarkList(b *testing.B) {
	benchmarkDeque(b, newList)
}
