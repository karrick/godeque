package godeque

type CircularList struct {
	head   *dnode
	tail   *dnode
	length int
}

// Len returns the number of items in the list.
func (d *CircularList) Len() int {
	return d.length
}

// Pop extracts and returns the value from the tail of the list, causing the
// tail to move to its previous element.
func (cl *CircularList) Pop() (interface{}, bool) {
	if cl.head == nil {
		return nil, false
	}

	v := cl.tail.value

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
	return v, true
}

// Push appends v to the tail of the list causing it to become the new list
// tail.
func (cl *CircularList) Push(v interface{}) {
	n := &dnode{value: v}

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

// Shift extracts and returns the value from the head of the list, advancing the
// head to the next item in the list.
func (cl *CircularList) Shift() (interface{}, bool) {
	if cl.head == nil {
		return nil, false
	}

	v := cl.head.value

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
	return v, true
}

// Unshift prepends v to the head of the list causing it to become the new list
// head.
func (cl *CircularList) Unshift(v interface{}) {
	n := &dnode{value: v}

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

// Peek returns the value at the head of the list without modifying the list.
func (cl *CircularList) Peek() (interface{}, bool) {
	if cl.head == nil {
		return nil, false
	}
	return cl.head.value, true
}

// Insert appends v to the tail of the list causing it to become the new list
// tail.
func (cl *CircularList) Insert(value interface{}) {
	cl.Push(value)
}

// Remove extracts and returns the value from the head of the list, advancing
// the head to the next item in the list.
func (cl *CircularList) Remove() (interface{}, bool) {
	return cl.Shift()
}
