package godeque

import "testing"

func newCircularList() dequelike {
	return new(CircularList)
}

func TestCircularList(t *testing.T) {
	testSuite(t, newCircularList)
}

func BenchmarkCircularList(b *testing.B) {
	benchmarkSuite(b, newCircularList)
}
