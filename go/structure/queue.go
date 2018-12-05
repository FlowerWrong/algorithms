package structure

import (
	"errors"

	"github.com/FlowerWrong/algorithm/go/structure/linearlist/linkedlist"
)

// FIFO: First-In-First-Out

// Queue ...
type Queue struct {
	front *linkedlist.LinkedNode
	rear  *linkedlist.LinkedNode
	len   int
}

// NewQueue ...
func NewQueue() *Queue {
	return &Queue{len: 0}
}

// IsEmpty ...
func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

// Enqueue ...
func (q *Queue) Enqueue(e interface{}) error {
	node := &linkedlist.LinkedNode{Data: e, Next: nil, Pre: nil}
	if q.front == nil && q.rear == nil && q.len == 0 {
		q.front = node
		q.rear = node
	} else {
		node.Next = q.rear
		q.rear.Pre = node
		q.rear = node
	}
	q.len++
	return nil
}

// Dequeue ...
func (q *Queue) Dequeue() (interface{}, error) {
	if q.front == nil && q.rear == nil && q.len == 0 {
		return nil, errors.New("empty queue")
	}
	node := q.front
	if node.Pre == nil && q.rear == q.front && q.len == 1 {
		q.front = nil
		q.rear = nil
	} else {
		q.front = node.Pre
		node.Pre.Next = nil
	}
	q.len--
	return node.Data, nil
}

// Len ...
func (q *Queue) Len() int {
	return q.len
}
