# godeque

Data structure that allows insertions and removals from either side of a list

## Overview [![GoDoc](https://godoc.org/github.com/karrick/godeque?status.svg)](https://godoc.org/github.com/karrick/godeque)

Deque data structures have Push and Pop to add and remove items from
the right side of the list, and Unshift and Shift to add and remove
items from the left side of the list.

A deque may be used as a Last-In-First-Out (LIFO) stack simply by
calling Push to add items and Pop to remove items in reverse order.

A deque may be used as a queue by using Push to insert items and Shift
to remove them, however the data structures in this library also
provide Enqueue and Dequeue methods for doing the same thing.

This library provides three implementations of a deque, namely List,
LockingList, and CircularList. LockingList provides go-routine safety
by wrapping a List with a sync.RWMutex. Due to its relative
simplicity, List is slightly more performant than CircularList in
benchmarks.

*NOTE* If all you need is a Queue, then I highly recommend using
[https://github.com/karrick/goqueue](https://github.com/karrick/goqueue),
which is a lock-free queue, which is much more performant than
LockingList provided by this library.

```Go
package main

import (
    "fmt"
    
    "github.com/karrick/godeque"
)

func main() {
    var l godeque.List
    
    for i := 0; i < 4; i++ {
        l.Push(i)
    }
    
    fmt.Printf("list has %d items\n", l.Len())

    for {
        value, ok := l.Shift()
        if !ok {
            break
        }
        fmt.Println(value)
    }
}
```

## Install

    go get github.com/karrick/godeque

## License

MIT.
