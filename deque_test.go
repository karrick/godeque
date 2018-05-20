package godeque

import "testing"

func newDeque() dequelike {
	return new(Deque)
}

func TestDeque(t *testing.T) {
	testSuite(t, newDeque)
}

func BenchmarkDeque(b *testing.B) {
	benchmarkSuite(b, newDeque)
}
