package godeque

import "testing"

func newDeque() deque {
	return new(Deque)
}

func TestDeque(t *testing.T) {
	testSuite(t, newDeque)
}

func BenchmarkDeque(b *testing.B) {
	benchmarkSuite(b, newDeque)
}
