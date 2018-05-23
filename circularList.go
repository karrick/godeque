package godeque

// CircularList is a deque implemented with a circular list.
//
// While correct, this implementation is not as fast as a List, but provided for
// reference.
type CircularList struct {
	head   *dnode
	tail   *dnode
	length int
}

// Len returns the number of items in the list.
func (d *CircularList) Len() int {
	return d.length
}

// Pop extracts and returns the datum from the tail of the list, causing the
// tail to move to its previous element.
func (cl *CircularList) Pop() (interface{}, bool) {
	if cl.head == nil {
		return nil, false
	}

	datum := cl.tail.datum

	if cl.head == cl.head.next {
		// list had a single value and ought now have none
		cl.head = nil
		cl.tail = nil
	} else {
		// list had at least two values
		cl.head.prev = cl.tail.prev
		cl.tail = cl.tail.prev
		cl.tail.next = cl.head
	}

	cl.length--
	return datum, true
}

// Push appends datum to the tail of the list causing it to become the new list
// tail.
func (cl *CircularList) Push(datum interface{}) {
	n := &dnode{datum: datum}

	if cl.head == nil {
		// list had no items
		n.next = n
		n.prev = n
		cl.head = n
		cl.tail = n
	} else {
		// list already has one or more items
		n.next = cl.head
		n.prev = cl.tail
		cl.tail.next = n
		cl.head.prev = n
		cl.tail = n
	}
	cl.length++
}

// Shift extracts and returns the datum from the head of the list, advancing the
// head to the next item in the list.
func (cl *CircularList) Shift() (interface{}, bool) {
	if cl.head == nil {
		return nil, false
	}

	datum := cl.head.datum

	if cl.head == cl.head.next {
		// list had a single value and ought now have none
		cl.head = nil
		cl.tail = nil
	} else {
		// list had at least two values
		cl.tail.next = cl.head.next
		cl.head.next.prev = cl.tail
		cl.head = cl.head.next
	}

	cl.length--
	return datum, true
}

// Unshift prepends datum to the head of the list causing it to become the new
// list head.
func (cl *CircularList) Unshift(datum interface{}) {
	n := &dnode{datum: datum}

	if cl.head == nil {
		// list had no items
		n.next = n
		n.prev = n
		cl.head = n
		cl.tail = n
	} else {
		// list already had one or more items
		n.next = cl.head
		n.prev = cl.tail
		cl.tail.next = n
		cl.head.prev = n
		cl.head = n
	}
	cl.length++
}

// Forward advances the head and tail one element further into the list.
func (cl *CircularList) Forward() {
	if cl.head == nil {
		return
	}
	cl.tail = cl.head
	cl.head = cl.tail.next
}

// Reverse retards the head and tail one element back from the list.
func (cl *CircularList) Reverse() {
	if cl.head == nil {
		return
	}
	cl.head = cl.tail
	cl.tail = cl.head.prev
}

// Peek returns the datum at the head of the list without modifying the list.
func (cl *CircularList) Peek() (interface{}, bool) {
	if cl.head == nil {
		return nil, false
	}
	return cl.head.datum, true
}

// Enqueue appends datum to the tail of the list causing it to become the new
// list tail.
func (cl *CircularList) Enqueue(datum interface{}) {
	cl.Push(datum)
}

// Dequeue extracts and returns the datum from the head of the list, advancing
// the head to the next item in the list.
func (cl *CircularList) Dequeue() (interface{}, bool) {
	return cl.Shift()
}
