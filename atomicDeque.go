package godeque

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type atomicDeque struct {
	head, tail *dnode
}

func newAtomicDeque() *atomicDeque {
	return new(atomicDeque)
}

func (d *atomicDeque) enumerate() int {
	fmt.Printf("--- enumerate ---\n")
	var itemCount int

	for np := d.tail; np != nil; np = np.prev {
		itemCount++
		fmt.Printf("node: %v; %v\n", unsafe.Pointer(np), np)
	}
	return itemCount
}

// Close should be named Close
func (d *atomicDeque) Close() error {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&d.head)), nil)
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&d.tail)), nil)
	return nil
}

// insertTop enqueues a datum to top of deque.  It is a wait-free method.
//
// NOTE: When pushing an item to a deque, the following 4 step procedure is used.
//
// STEP 1: Create a new node.
//
// STEP 2: oldTop = SwapPointer(d.top, newNode)
//
// STEP 3: When deque was previously empty, also point bottom to new
//         node. Otherwise point previous top to new node and new node's down to
//         previous top, appending new node on top of the list.
//
// STEP 4: When dequeue was not previously empty, update newNode's down pointer
//         to point to oldTop.
//
// BEFORE            |  STEP 1            |  STEP 2                 |  STEP 3a                |  STEP 3b
// ------------------+--------------------+-------------------------+-------------------------+-------------------------
//                   |                    |                         |
//                   |    newNode -> [3]  |  newNode, d.top -> [3]  |  newNode, d.top -> [3]  |  newNode, d.top -> [3]
//                   |                    |                         |                     ^   |                     ^
//                   |                    |                         |                     |   |                     |
//                   |                    |                         |                     |   |                     v
//     d.top -> [2]  |      d.top -> [2]  |          oldTop -> [2]  |          oldTop -> [2]  |          oldTop -> [2]
//               ^   |                ^   |                     ^   |                     ^   |                     ^
//               |   |                |   |                     |   |                     |   |                     |
//               v   |                v   |                     v   |                     v   |                     v
//              [1]  |               [1]  |                    [1]  |                    [1]  |                    [1]
//               ^   |                ^   |                     ^   |                     ^   |                     ^
//               |   |                |   |                     |   |                     |   |                     |
//               v   |                v   |                     v   |                     v   |                     v
//  d.bottom -> [0]  |   d.bottom -> [0]  |        d.bottom -> [0]  |        d.bottom -> [0]  |        d.bottom -> [0]
//
func (d *atomicDeque) insertTop(datum interface{}) {
	// STEP 1
	newNode := &dnode{datum: datum}

	// STEP 2 (oldTop, top = top, newNode)
	oldTop := (*dnode)(atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&d.head)), unsafe.Pointer(newNode)))

	// STEP 3 (whichever of these is appropriate)
	if oldTop != nil {
		// STEP 3a: deque was non-empty, therefore point oldTop's up to newNode
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&oldTop.prev)), unsafe.Pointer(newNode))
		// STEP 3b: and update newNode's down to oldTop
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&newNode.next)), unsafe.Pointer(oldTop))
	} else {
		// deque was empty (where bottom is already nil), therefore update bottom to point to newNode
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&d.tail)), unsafe.Pointer(newNode))
	}
}

// insertBottom enqueues a datum to bottom of deque.  It is wait-free.
func (d *atomicDeque) insertBottom(datum interface{}) {
	// STEP 1
	newNode := &dnode{datum: datum}

	// STEP 2 (oldBottom, d.bottom = d.bottom, newNode)
	oldBottom := (*dnode)(atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&d.tail)), unsafe.Pointer(newNode)))

	// STEP 3 (whichever of these is appropriate)
	if oldBottom != nil {
		// STEP 3a: deque was non-empty, therefore point oldBottom's down to newNode
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&oldBottom.next)), unsafe.Pointer(newNode))
		// STEP 3b: and update newNode's up to oldBottom
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&newNode.prev)), unsafe.Pointer(oldBottom))
	} else {
		// deque was empty (where dequeue is already nil), therefore update top to point to newNode
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&d.head)), unsafe.Pointer(newNode))
	}
}

// removeBottom dequeues a datam from from bottom of deque.  It is wait-free and returns immediately when no items are in deque.
//
// NOTE: When pushing an item to a deque, the following 4 step procedure is used.
//
// STEP 1: Create a new node.
//
// STEP 2: oldEnqueue = SwapPointer(d.enqueue, newNode)
//
// STEP 3: When deque was previously empty, also point bottom to new node.  Otherwise point previous
//         top to new node and new node's down to previous top, appending new node on top of the
//         list.
//
// STEP 4: When dequeue was not previously empty, update newNode's down pointer to pont to oldTop.
//
// |  BEFORE           |  STEP 1              |  STEP 2              |  STEP 3                |  STEP 4                |  STEP 5
// +-------------------+----------------------+----------------------+------------------------+------------------------+----------------------
// |  [3] <- d.top     |  [3] <- d.top        |  [3] <- d.top        |  [3] <- d.top          |  [3] <- d.top	       |  [3] <- d.top
// |   ^               |   ^                  |   ^                  |   ^                    |   ^		       |   ^
// |   |               |   |                  |   |                  |   |                    |   |		       |   |
// |   v               |   v                  |   v                  |   v                    |   v		       |   v
// |  [2]              |  [2]                 |  [2]                 |  [2]                   |  [2]		       |  [2]
// |   ^               |   ^                  |   ^                  |   ^                    |   ^		       |   ^
// |   |               |   |                  |   |                  |   |                    |   |		       |   |
// |   v               |   v                  |   v                  |   v                    |   v		       |   v
// |  [1]              |  [1]                 |  [1] <- up0, up1     |  [1] <- d.bottom, up0  |  [1] <- d.bottom, up0  |  [1] <- d.bottom, up0
// |   ^               |   ^                  |   ^                  |   ^                    |   ^		       |   ^
// |   |               |   |                  |   |                  |   |                    |   |		       |   |
// |   v               |   v                  |   v                  |   v                    |   |		       |   |
// |  [0] <- d.bottom  |  [0] <- d.bottom, b  |  [0] <- d.bottom, b  |  [0] <- b              |  [0] <- b              |  [0] <- b
//
// One characteristic of this algorithm is that in high contention environments removing an element
// from one of the ends will result in lots of `return nil, false` even though there are elements
// that could be removed from the deque.  Although not ideal behavior, and unacceptable for use by
// code that requires the result to be false only when no items are on the deque, this particular
// data structure is meant to be consumed by other data structure which implements this checking.
//
func (d *atomicDeque) removeBottom() (interface{}, bool) {
	var bottom *dnode
	for {
		// STEP 1
		bottom = (*dnode)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&d.tail))))
		if bottom == nil {
			// The list is empty
			return nil, false
		}

		// STEP 2
		newBottom := (*dnode)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&bottom.prev))))

		// STEP 3
		if !atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&d.tail)), unsafe.Pointer(bottom), unsafe.Pointer(newBottom)) {
			// A different thread already moved bottom up
			return nil, false
		}
		break
	}

	// // STEP 4: If new bottom holds onto pointer to b, then b will not be released
	// atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&d.bottom.down)), nil)

	// STEP 5: When top matches the old bottom, the list is now empty
	atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&d.head)), unsafe.Pointer(bottom), nil)

	return bottom.datum, true
}

// removeTop dequeues a datum from top of deque.  It is wait-free and returns immediately when no items are in deque.
func (d *atomicDeque) removeTop() (interface{}, bool) {
	var top *dnode
	for {
		// STEP 1
		top = (*dnode)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&d.head))))
		if top == nil {
			// The list is empty
			return nil, false
		}

		// STEP 2
		newTop := (*dnode)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&top.prev))))

		// STEP 3
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&d.head)), unsafe.Pointer(top), unsafe.Pointer(newTop)) {
			break
		}
		// A different thread already moved top up
	}

	// // STEP 4: If new top holds onto pointer to b, then b will not be released
	// atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&d.top.down)), nil)

	// STEP 5: When top matches the old top, the list is now empty
	atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&d.head)), unsafe.Pointer(top), nil)

	return top.datum, true
}
