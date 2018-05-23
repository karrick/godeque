package godeque_test

import (
	"math/rand"
	"testing"

	"github.com/karrick/godeque"
)

func testDeque(t *testing.T, init func() godeque.Deque) {
	t.Run("push pop", func(t *testing.T) {
		// behaves like a stack

		d := init()

		d.Push(11)
		d.Push(22)
		d.Push(33)
		d.Push(44)

		t.Run("first pop", func(t *testing.T) {
			if got, want := d.Len(), 4; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}

			v, ok := d.Pop()
			if got, want := v, 44; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("second pop", func(t *testing.T) {
			if got, want := d.Len(), 3; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			v, ok := d.Pop()
			if got, want := v, 33; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("third pop", func(t *testing.T) {
			if got, want := d.Len(), 2; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			v, ok := d.Pop()
			if got, want := v, 22; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("fourth pop", func(t *testing.T) {
			if got, want := d.Len(), 1; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			v, ok := d.Pop()
			if got, want := v, 11; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("fifth pop", func(t *testing.T) {
			if got, want := d.Len(), 0; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			_, ok := d.Pop()
			if got, want := ok, false; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
	})

	t.Run("push shift", func(t *testing.T) {
		// behaves like a fifo

		d := init()

		d.Push(11)
		d.Push(22)
		d.Push(33)
		d.Push(44)

		t.Run("first shift", func(t *testing.T) {
			v, ok := d.Shift()
			if got, want := v, 11; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("second shift", func(t *testing.T) {
			v, ok := d.Shift()
			if got, want := v, 22; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("third shift", func(t *testing.T) {
			v, ok := d.Shift()
			if got, want := v, 33; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("fourth shift", func(t *testing.T) {
			v, ok := d.Shift()
			if got, want := v, 44; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("fifth shift", func(t *testing.T) {
			_, ok := d.Shift()
			if got, want := ok, false; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
	})

	t.Run("unshift pop", func(t *testing.T) {
		// behaves like a fifo

		d := init()

		d.Unshift(11)
		d.Unshift(22)
		d.Unshift(33)
		d.Unshift(44)

		t.Run("first pop", func(t *testing.T) {
			v, ok := d.Pop()
			if got, want := v, 11; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("second pop", func(t *testing.T) {
			v, ok := d.Pop()
			if got, want := v, 22; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("third pop", func(t *testing.T) {
			v, ok := d.Pop()
			if got, want := v, 33; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("fourth pop", func(t *testing.T) {
			v, ok := d.Pop()
			if got, want := v, 44; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("fifth pop", func(t *testing.T) {
			_, ok := d.Pop()
			if got, want := ok, false; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
	})

	t.Run("unshift shift", func(t *testing.T) {
		// behaves like a stack

		d := init()

		d.Unshift(11)
		d.Unshift(22)
		d.Unshift(33)
		d.Unshift(44)

		t.Run("first shift", func(t *testing.T) {
			v, ok := d.Shift()
			if got, want := v, 44; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("second shift", func(t *testing.T) {
			v, ok := d.Shift()
			if got, want := v, 33; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("third shift", func(t *testing.T) {
			v, ok := d.Shift()
			if got, want := v, 22; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("fourth shift", func(t *testing.T) {
			v, ok := d.Shift()
			if got, want := v, 11; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("fifth shift", func(t *testing.T) {
			_, ok := d.Shift()
			if got, want := ok, false; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
	})

	t.Run("unshift shift", func(t *testing.T) {
		d := init()

		d.Unshift(11)
		d.Unshift(22)
		d.Unshift(33)
		d.Unshift(44)

		t.Run("first shift", func(t *testing.T) {
			v, ok := d.Shift()
			if got, want := v, 44; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("second shift", func(t *testing.T) {
			v, ok := d.Shift()
			if got, want := v, 33; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("third shift", func(t *testing.T) {
			v, ok := d.Shift()
			if got, want := v, 22; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("fourth shift", func(t *testing.T) {
			v, ok := d.Shift()
			if got, want := v, 11; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := ok, true; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("fifth shift", func(t *testing.T) {
			_, ok := d.Shift()
			if got, want := ok, false; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
	})
}

func benchmarkDeque(b *testing.B, init func() godeque.Deque) {
	values := rand.Perm(b.N)

	d := init()

	b.ReportAllocs()
	b.ResetTimer()

	for i, v := range values {
		switch i % 10 {
		case 0:
			switch i % 20 {
			case 0:
				_, _ = d.Pop()
			case 1:
				_, _ = d.Shift()
			}
		default:
			switch i % 2 {
			case 0:
				d.Push(v)
			case 1:
				d.Unshift(v)
			}
		}
	}
}
