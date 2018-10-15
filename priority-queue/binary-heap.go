package pqueue

import "fmt"

// BinaryHeap represents priority queue implementation
// using binary heap (this example is for integers only)
type BinaryHeap struct {
	size int   // The number of elements currently inside the heap
	cap  int   // The internal capacity of the heap
	heap []int // A slice to track the elements inside the heap
}

// New creates a priority queue with a slice of elements
func (bh *BinaryHeap) New(elems []int) {
	bh.size = len(elems)
	bh.cap = len(elems)
	bh.heap = elems
}

// IsEmpty returns true/false depending on if the priority queue is empty
func (bh *BinaryHeap) IsEmpty() bool {
	return bh.size == 0
}

// Clear clears everything inside the heap, O(n)
func (bh *BinaryHeap) Clear() {
	bh.heap = []int{}
	bh.size = 0
}

// Size return heap's size
func (bh *BinaryHeap) Size() int {
	return bh.size
}

// Peek returns the value of the element with the lowest
// priority in this priority queue. If the priority
// queue is empty 0 is returned.
func (bh *BinaryHeap) Peek() (int, error) {
	if bh.IsEmpty() {
		err := fmt.Errorf("Heap is empty")
		return 0, err
	}

	return bh.heap[0], nil
}

// Poll removes the root of the heap, O(log(n))
func (bh *BinaryHeap) Poll() {
	_, _ = bh.RemoveAt(0)
}

// Add adds an element to the priority queue, the
// element must not be 0, O(log(n))
func (bh *BinaryHeap) Add(el int) error {
	if el == 0 {
		err := fmt.Errorf("Invalid element")
		return err
	}

	if bh.size < bh.cap {
		bh.heap[bh.size] = el
	} else {
		bh.heap = append(bh.heap, el)
		bh.cap++
	}

	bh.swim(bh.size)
	bh.size++
	return nil
}

// Tests if the value of node i <= node j
// This method assumes i & j are valid indices, O(1)
func (bh *BinaryHeap) less(i, j int) bool {
	node1 := bh.heap[i]
	node2 := bh.heap[j]

	return node1 <= node2
}

// Perform bottom up node swim, O(log(n))
func (bh *BinaryHeap) swim(k int) {
	// Grab the index of the next parent node WRT to k
	parent := (k - 1) / 2

	// Keep swimming while we have not reached the
	// root and while we're less than our parent.
	for k > 0 && bh.less(k, parent) {
		// Exchange k with the parent
		bh.swap(parent, k)
		k = parent

		// Grab the index of the next parent node WRT to k
		parent = (k - 1) / 2
	}
}

// Top down node sink, O(log(n))
func (bh *BinaryHeap) sink(k int) {
	for {
		left := 2*k + 1  // Left  node
		right := 2*k + 2 // Right node
		smallest := left // Assume left is the smallest node of the two children

		// Find which is smaller left or right
		// If right is smaller set smallest to be right
		if right < bh.size && bh.less(right, left) {
			smallest = right
		}

		// Stop if we're outside the bounds of the tree
		// or stop early if we cannot sink k anymore
		if left >= bh.size || bh.less(k, smallest) {
			break
		}

		// Move down the tree following the smallest node
		bh.swap(smallest, k)
		k = smallest
	}
}

// Swap two nodes. Assumes i & j are valid, O(1)
func (bh *BinaryHeap) swap(i, j int) {
	iElem := bh.heap[i]
	jElem := bh.heap[j]

	bh.heap[i] = jElem
	bh.heap[j] = iElem
}

// RemoveAt removes a node at particular index, O(log(n))
func (bh *BinaryHeap) RemoveAt(i int) (int, error) {
	if bh.IsEmpty() {
		err := fmt.Errorf("Heap is empty")
		return 0, err
	}

	bh.size--
	data := bh.heap[i]
	bh.swap(i, bh.size)

	// Remove element
	bh.heap = bh.heap[:bh.size]

	// Check if the last element was removed
	if i == bh.size {
		return data, nil
	}
	elem := bh.heap[i]

	// Try sinking element
	bh.sink(i)

	// If sinking did not work try swimming
	if bh.heap[i] == elem {
		bh.swim(i)

	}
	return data, nil
}

// IsMinHeap validates the heap
func (bh *BinaryHeap) IsMinHeap(k int) bool {
	// If we are outside the bounds of the heap return true
	if k >= bh.size {
		return true
	}

	left := 2*k + 1
	right := 2*k + 2

	// Make sure that the current node k is less than
	// both of its children left, and right if they exist
	// return false otherwise to indicate an invalid heap
	if left < bh.size && !bh.less(k, left) {
		return false
	}

	if right < bh.size && !bh.less(k, left) {
		return false
	}

	// Recurse on both children to make sure they're also valid heaps
	return bh.IsMinHeap(left) && bh.IsMinHeap(right)
}
