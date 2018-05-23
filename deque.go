package godeque

// dnode is a node to be used in a doubly linked list
type dnode struct {
	prev  *dnode
	next  *dnode
	datum interface{}
}

// Deque is a data structure that allows insertion and removal of datum values
// from both ends of its ordered list. Deques have the same capabilities of both
// a queue and a stack.
type Deque interface {
	Len() int
	Pop() (interface{}, bool)
	Push(interface{})
	Shift() (interface{}, bool)
	Unshift(interface{})
}

// Queue is a first-in-first-out (FIFO) data structure that allows insertion and
// removal of datum values.
type Queue interface {
	Enqueue(interface{})
	Dequeue() (interface{}, bool)
}

// Stack is a last-in-first-out (LIFO) data structure that allows insertion and
// removal of datum values.
type Stack interface {
	Pop() (interface{}, bool)
	Push(interface{})
}
