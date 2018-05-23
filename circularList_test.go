package godeque

import "testing"

func newCircularList() deque {
	return new(circularList)
}

func TestCircularList(t *testing.T) {
	testSuite(t, newCircularList)
}

func BenchmarkCircularList(b *testing.B) {
	benchmarkSuite(b, newCircularList)
}
