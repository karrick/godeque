package godeque

// dnode is a node used in a doubly linked list
type dnode struct {
	next  *dnode
	prev  *dnode
	value interface{}
}

type Deque struct {
	head, tail *dnode
	length     int
}

// Len returns the number of items in the list.
func (d *Deque) Len() int {
	return d.length
}

// Pop extracts and returns the value from the tail of the list, causing the
// tail to move to its previous element.
func (d *Deque) Pop() (interface{}, bool) {
	if d.tail == nil {
		return nil, false
	}
	v := d.tail.value
	d.tail = d.tail.prev
	if d.tail == nil {
		d.head = nil
	} else {
		d.tail.next = nil
	}
	d.length--
	return v, true
}

// Push appends v to the tail of the list causing it to become the new list
// tail.
func (d *Deque) Push(value interface{}) {
	n := &dnode{value: value, prev: d.tail}
	if d.tail == nil {
		d.head = n
	} else {
		d.tail.next = n
	}
	d.length++
	d.tail = n
}

// Shift extracts and returns the value from the head of the list, advancing the
// head to the next item in the list.
func (d *Deque) Shift() (interface{}, bool) {
	if d.head == nil {
		return nil, false
	}
	v := d.head.value
	d.head = d.head.next
	if d.head == nil {
		d.tail = nil
	} else {
		d.head.prev = nil
	}
	d.length--
	return v, true
}

// Unshift prepends v to the head of the list causing it to become the new list
// head.
func (d *Deque) Unshift(value interface{}) {
	n := &dnode{value: value, next: d.head}
	if d.head == nil {
		d.tail = n // head only nil when tail nil
	} else {
		d.head.prev = n
	}
	d.head = n
	d.length++
}

// Insert appends v to the tail of the list causing it to become the new list
// tail.
func (d *Deque) Insert(value interface{}) {
	d.Push(value)
}

// Remove extracts and returns the value from the head of the list, advancing
// the head to the next item in the list.
func (d *Deque) Remove() (interface{}, bool) {
	return d.Shift()
}
