package structure

import (
	"errors"

	"github.com/FlowerWrong/algorithm/go/structure/linearlist/linkedlist"
)

// LIFO: Last In First Out

// FIFO: First-In-First-Out

// Stack ...
type Stack struct {
	top *linkedlist.LinkedNode
	len int
}

// NewStack ...
func NewStack() *Stack {
	return &Stack{len: 0, top: nil}
}

// IsEmpty ...
func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

// Push ...
func (s *Stack) Push(e interface{}) error {
	node := &linkedlist.LinkedNode{Data: e, Next: nil, Pre: nil}
	if s.top == nil && s.len == 0 {
		s.top = node
	} else {
		node.Pre = s.top
		s.top.Next = node
		s.top = node
	}
	s.len++
	return nil
}

// Pop ...
func (s *Stack) Pop() (interface{}, error) {
	if s.top == nil && s.len == 0 {
		return nil, errors.New("empty stack")
	}
	node := s.top
	if s.len == 1 {
		s.top = nil
	} else {
		s.top = node.Pre
		s.top.Next = nil
	}
	s.len--
	return node.Data, nil
}

// Len ...
func (s *Stack) Len() int {
	return s.len
}
