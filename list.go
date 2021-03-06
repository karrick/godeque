package godeque

// List allows adding and removing datum values both to the head and the tail of
// the list.
type List struct {
	head, tail *dnode
	length     int
}

// Len returns the number of items in the list.
func (d *List) Len() int {
	return d.length
}

// Pop extracts and returns the datum from the tail of the list, causing the
// tail to move to its previous element.
func (d *List) Pop() (interface{}, bool) {
	if d.tail == nil {
		return nil, false
	}
	datum := d.tail.datum
	d.tail = d.tail.prev
	if d.tail == nil {
		d.head = nil
	} else {
		d.tail.next = nil
	}
	d.length--
	return datum, true
}

// Push appends datum to the tail of the list causing it to become the new list
// tail.
func (d *List) Push(datum interface{}) {
	n := &dnode{datum: datum, prev: d.tail}
	if d.tail == nil {
		d.head = n
	} else {
		d.tail.next = n
	}
	d.length++
	d.tail = n
}

// Shift extracts and returns the datum from the head of the list, advancing the
// head to the next item in the list.
func (d *List) Shift() (interface{}, bool) {
	if d.head == nil {
		return nil, false
	}
	datum := d.head.datum
	d.head = d.head.next
	if d.head == nil {
		d.tail = nil
	} else {
		d.head.prev = nil
	}
	d.length--
	return datum, true
}

// Unshift prepends datum to the head of the list causing it to become the new
// list head.
func (d *List) Unshift(datum interface{}) {
	n := &dnode{datum: datum, next: d.head}
	if d.head == nil {
		d.tail = n // head only nil when tail nil
	} else {
		d.head.prev = n
	}
	d.head = n
	d.length++
}

// Enqueue appends datum to the tail of the list causing it to become the new
// list tail.
func (d *List) Enqueue(datum interface{}) {
	d.Push(datum)
}

// Dequeue extracts and returns the datum from the head of the list, advancing
// the head to the next item in the list.
func (d *List) Dequeue() (interface{}, bool) {
	return d.Shift()
}
