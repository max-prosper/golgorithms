package list

import "fmt"

// Node represents single node struct
type Node struct {
	data interface{}
	prev *Node
	next *Node
}

// DLinked represents the Doubly Linked List struct
type DLinked struct {
	size int
	head *Node
	tail *Node
}

// Clear the list, O(n)
func (l *DLinked) Clear() {
	trav := l.head
	for trav != nil {
		next := trav.next
		trav.prev = nil
		trav.next = nil
		trav.data = nil
		trav = next
	}

	l.head = nil
	l.tail = nil
	trav = nil
	l.size = 0
}

// ToString prints DLinked to estout
func (n *Node) ToString() string {
	if n != nil {
		return fmt.Sprintf("%v", n.data)
	}
	return ""
}

// Size returns the size of this linked list
func (l *DLinked) Size() (size int) {
	size = l.size
	return
}

// IsEmpty returns true if linked list is empty
func (l *DLinked) IsEmpty() bool {
	return l.size == 0
}

// Add adds an element to the tail of the linked list, O(1)
func (l *DLinked) Add(el interface{}) {
	if l.IsEmpty() {
		l.head = &Node{el, nil, nil}
		l.tail = l.head
	} else {
		l.tail.next = &Node{el, nil, nil}
		l.tail = l.tail.next
	}
	l.size++
}

// AddFirst adds an element to the beginning of this linked list, O(1)\
func (l *DLinked) AddFirst(el interface{}) {
	if l.IsEmpty() {
		l.head = &Node{el, nil, nil}
		l.tail = l.head
	} else {
		l.head.prev = &Node{el, nil, l.head}
		l.head = l.head.prev
	}

	l.size++
}

// PeakFirst chacks the value of the first node if it exists, O(1)
func (l *DLinked) PeakFirst() (interface{}, error) {
	if l != nil {
		data := l.head.data
		return data, nil
	}
	err := fmt.Errorf("List is empty")
	return nil, err
}

// PeakLast checks the value of the last node if it exists, O(1)
func (l *DLinked) PeakLast() (interface{}, error) {
	if l != nil {
		data := l.tail.data
		return data, nil
	}
	err := fmt.Errorf("List is empty")
	return nil, err
}

// RemoveFirst removes the first value at the head of the linked list, O(1)
func (l *DLinked) RemoveFirst() (interface{}, error) {
	// Can't remove data from an empty list
	if l.IsEmpty() {
		err := fmt.Errorf("List is empty")
		return nil, err
	}

	// Extract the data at the head and move
	// the head pointer forwards one node
	data := l.head.data
	l.head = l.head.next
	l.size--

	// If the list is empty after removeing the
	// first node set the tail to null and do
	// a memory clean up on the previous node
	if l.IsEmpty() {
		l.tail = nil
	} else {
		l.head.prev = nil
	}

	// Return the data that was at the first
	// node we just removed
	return data, nil
}

// RemoveLast removes the last value at the tail of the linked list, O(1)
func (l *DLinked) RemoveLast() (interface{}, error) {
	// Can't remove data from an empty list
	if l.IsEmpty() {
		err := fmt.Errorf("List is empty")
		return nil, err
	}

	// Extract the data at the tail and move
	// the tail pointer backwards one node
	data := l.tail.data
	l.tail = l.tail.prev
	l.size--

	// If the list is now empty set the head to null
	// Do a memory clean of the node that was just removed
	if l.IsEmpty() {
		l.head = nil
	} else {
		l.tail.next = nil
	}

	// Return the data that was in the last node we just removed
	return data, nil
}

// Remove removes an arbitrary node from the linked list, O(1)
func (l *DLinked) Remove(n *Node) (interface{}, error) {
	// If the node to remove is somewhere either at the
	// head or the tail handle those independently
	if n.prev == nil {
		data, err := l.RemoveFirst()
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	if n.next == nil {
		data, err := l.RemoveLast()
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	// Temporarily store the data we want to return
	data := n.data

	// Memory cleanup
	n.data = nil
	n.prev = nil
	n.next = nil
	n = nil
	l.size--

	// Return the data in the node we just removed
	return data, nil
}

// RemoveAt removes a node at a particular index, O(n)
func (l *DLinked) RemoveAt(i int) (interface{}, error) {
	// Make sure the index provided is valid -_-
	if i < 0 || i >= l.size {
		err := fmt.Errorf("Index out of range")
		return nil, err
	}

	var trav *Node

	// Search from the front of the list
	if i < l.size/2 {
		trav = l.head
		for index := 0; index != i; index++ {
			trav = trav.next
		}

		// Search from the back of the list
	} else {
		trav = l.tail
		for index := l.size - 1; i != index; i-- {
			trav = trav.prev
		}
	}

	data, err := l.Remove(trav)
	if err != nil {
		return nil, err
	}

	return data, nil
}
