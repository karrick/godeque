package godeque

import "sync"

// LockingList is a go routine safe list that serializes operations through a
// mutex.
type LockingList struct {
	l sync.RWMutex
	d *List
}

// NewLockingList returns a go routine safe list that serializes operations
// through a mutex.
func NewLockingList() *LockingList {
	return &LockingList{d: new(List)}
}

// Len returns the number of items in the list.
func (ld *LockingList) Len() int {
	ld.l.RLock()
	l := ld.d.Len()
	ld.l.RUnlock()
	return l
}

// Pop extracts and returns the datum from the tail of the list, causing the
// tail to move to its previous element.
func (ld *LockingList) Pop() (interface{}, bool) {
	ld.l.Lock()
	v, o := ld.d.Pop()
	ld.l.Unlock()
	return v, o
}

// Push appends datum to the tail of the list causing it to become the new list
// tail.
func (ld *LockingList) Push(datum interface{}) {
	ld.l.Lock()
	ld.d.Push(datum)
	ld.l.Unlock()
}

// Shift extracts and returns the datum from the head of the list, advancing
// the head to the next item in the list.
func (ld *LockingList) Shift() (interface{}, bool) {
	ld.l.Lock()
	v, o := ld.d.Shift()
	ld.l.Unlock()
	return v, o
}

// Unshift prepends datum to the head of the list causing it to become the new
// list head.
func (ld *LockingList) Unshift(datum interface{}) {
	ld.l.Lock()
	ld.d.Unshift(datum)
	ld.l.Unlock()
}

// Enqueue appends datum to the tail of the list causing it to become the new
// list tail.
func (ld *LockingList) Enqueue(datum interface{}) {
	ld.Push(datum)
}

// Dequeue extracts and returns the datum from the head of the list, advancing
// the head to the next item in the list.
func (ld *LockingList) Dequeue() (interface{}, bool) {
	return ld.Shift()
}
