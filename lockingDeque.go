package godeque

import "sync"

// LockingDeque is a go routine safe deque that serializes operations through a
// mutex.
type LockingDeque struct {
	l sync.RWMutex
	d *Deque
}

// NewLockingDeque returns a go routine safe deque that serializes operations
// through a mutex.
func NewLockingDeque() *LockingDeque {
	return &LockingDeque{d: new(Deque)}
}

// Len returns the number of items in the deque.
func (ld *LockingDeque) Len() int {
	ld.l.RLock()
	l := ld.d.Len()
	ld.l.RUnlock()
	return l
}

// Pop extracts and returns the datum from the tail of the deque, causing the
// tail to move to its previous element.
func (ld *LockingDeque) Pop() (interface{}, bool) {
	ld.l.Lock()
	v, o := ld.d.Pop()
	ld.l.Unlock()
	return v, o
}

// Push appends datum to the tail of the deque causing it to become the new list
// tail.
func (ld *LockingDeque) Push(datum interface{}) bool {
	ld.l.Lock()
	ld.d.Push(datum)
	ld.l.Unlock()
}

// Shift extracts and returns the datum from the head of the deque, advancing
// the head to the next item in the deque.
func (ld *LockingDeque) Shift() (interface{}, bool) {
	ld.l.Lock()
	v, o := ld.d.Shift()
	ld.l.Unlock()
	return v, o
}

// Unshift prepends datum to the head of the deque causing it to become the new
// list head.
func (ld *LockingDeque) Unshift(datum interface{}) bool {
	ld.l.Lock()
	ld.d.Unshift(datum)
	ld.l.Unlock()
}

// Enqueue appends datum to the tail of the deque causing it to become the new
// list tail.
func (ld *LockingDeque) Enqueue(datum interface{}) {
	ld.Push(datum)
}

// Dequeue extracts and returns the datum from the head of the deque, advancing
// the head to the next item in the deque.
func (ld *LockingDeque) Dequeue() (interface{}, bool) {
	return ld.Shift()
}
