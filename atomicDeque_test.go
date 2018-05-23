package godeque

import (
	"sync"
	"sync/atomic"
	"testing"
)

func atomicDequeStress(tb testing.TB, toppers, bottomers int, itemCount int) {
	tb.Skip("TODO")
	d := newAtomicDeque()
	defer func() { _ = d.Close() }()

	const value = 0x1a

	var inserted, removed, spin int64
	var wg sync.WaitGroup

	for i := 0; i < toppers; i++ {
		wg.Add(2)
		go func(d *atomicDeque, wg *sync.WaitGroup) {
			for {
				d.insertTop(value)
				if i := atomic.AddInt64(&inserted, 1); i == int64(itemCount) {
					break
				}
			}
			wg.Done()
		}(d, &wg)
		go func(d *atomicDeque, wg *sync.WaitGroup) {
			for {
				if v, ok := d.removeTop(); ok {
					if i := atomic.AddInt64(&removed, 1); i == int64(itemCount) {
						break
					}
					if v != value {
						tb.Fatalf("Actual: %v; Expected: %v", v, value)
					}
				}
				s := atomic.AddInt64(&spin, 1)
				if s == int64(itemCount*2) {
					break
				}
				if s%100 == 0 {
					tb.Logf("spin: %d", s)
				}
			}
			wg.Done()
		}(d, &wg)
	}

	for i := 0; i < bottomers; i++ {
		wg.Add(2)
		go func(d *atomicDeque, wg *sync.WaitGroup) {
			for {
				d.insertBottom(value)
				if i := atomic.AddInt64(&inserted, 1); i == int64(itemCount) {
					break
				}
			}
			wg.Done()
		}(d, &wg)
		go func(d *atomicDeque, wg *sync.WaitGroup) {
			for {
				if v, ok := d.removeBottom(); ok {
					if i := atomic.AddInt64(&removed, 1); i == int64(itemCount) {
						break
					}
					if v != value {
						tb.Fatalf("Actual: %v; Expected: %v", v, value)
					}
				}
				s := atomic.AddInt64(&spin, 1)
				if s == int64(itemCount*2) {
					break
				}
				if s%100 == 0 {
					tb.Logf("spin: %d", s)
				}
			}
			wg.Done()
		}(d, &wg)
	}

	wg.Wait()

	if inserted != removed {
		tb.Errorf("inserted: %v; removed: %v", inserted, removed)
	}
	tb.Logf("spin: %d", spin)
}

func TestAtomicDequeBottomStress(t *testing.T) {
	// dequeBottomStress(t, 1000, 1000)
	atomicDequeStress(t, 0, 4, 10)
}

// func TestAtomicDequeTopStress(t *testing.T) {
// 	dequeTopStress(t, 1000, 1000)
// }

// func TestAtomicDequeStress(t *testing.T) {
// 	dequeStress(t, 1000, 1000)
// }

// func BenchmarkAtomicDequeBottomDeep(b *testing.B) {
// 	dequeBottomStress(b, 100, b.N)
// }

// func BenchmarkAtomicDequeBottomWide(b *testing.B) {
// 	dequeBottomStress(b, b.N, 100)
// }
