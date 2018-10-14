package stack

import (
	"fmt"

	"github.com/max-prosper/golgorithms/list"
)

// Stack represents a stack structure base in linked list
type Stack struct {
	list list.DLinked
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return s.list.Size()
}

// Check if the stack is empty
func (s *Stack) isEmpty() bool {
	return s.Size() == 0
}

// Push an element on the stack
func (s *Stack) Push(el interface{}) {
	s.list.Add(el)
}

// Pop an element off the stack
func (s *Stack) Pop() (interface{}, error) {
	if s.isEmpty() {
		err := fmt.Errorf("List is empty")
		return nil, err
	}

	data, err := s.list.RemoveLast()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Peek the top of the stack without removing an element
func (s *Stack) Peek() (interface{}, error) {
	if s.isEmpty() {
		err := fmt.Errorf("List is empty")
		return nil, err
	}

	data, err := s.list.PeakLast()
	if err != nil {
		return nil, err
	}

	return data, nil
}
