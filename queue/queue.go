package queue

import (
	"fmt"

	"github.com/max-prosper/golgorithms/list"
)

// Queue represents a queue structure base in linked list
type Queue struct {
	list list.DLinked
}

// Size return the size of the queue
func (q *Queue) Size() int {
	return q.list.Size()
}

// IsEmpty returns whether or not the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.list.Size() == 0
}

// Peek the element at the front of the queue
func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		err := fmt.Errorf("Queue Empty")
		return nil, err
	}

	data, err := q.list.PeakFirst()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Enqueue the element
func (q *Queue) Enqueue(el interface{}) {
	q.list.Add(el)
}

// Dequeue the element
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		err := fmt.Errorf("Queue Empty")
		return nil, err
	}

	data, err := q.list.RemoveFirst()
	if err != nil {
		return nil, err
	}

	return data, nil
}
