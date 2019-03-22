package linkedlist

import (
	"errors"
)

// 单链表 singly linked list

// LinkedNode ...
type LinkedNode struct {
	Data interface{}
	Next *LinkedNode
	Pre  *LinkedNode
}

// LinkedList ...
type LinkedList struct {
	len  int
	head *LinkedNode
}

// NewLinkedList ...
func NewLinkedList() *LinkedList {
	return &LinkedList{len: 0, head: nil}
}

// IsEmpty ...
func (list *LinkedList) IsEmpty() bool {
	return list.len == 0
}

// Clear ...
func (list *LinkedList) Clear() {
	list.len = 0
	list.head = nil
}

// Insert ...
func (list *LinkedList) Insert(pos int, e interface{}) error {
	err := list.checkPos(pos)
	if err != nil {
		return err
	}

	head := list.head
	index := 1
	for head != nil && index < pos-1 {
		head = head.Next
		index++
	}

	s := &LinkedNode{Data: e}
	if head == nil && list.len == 0 {
		list.head = s
	} else {
		s.Next = head.Next
		head.Next = s
	}
	list.len++
	return nil
}

// Append ...
func (list *LinkedList) Append(e interface{}) error {
	head := list.head
	for head != nil && head.Next != nil {
		head = head.Next
	}
	s := &LinkedNode{Data: e, Next: nil}
	if head == nil && list.len == 0 {
		list.head = s
	} else {
		head.Next = s
	}
	list.len++
	return nil
}

// Data ...
func (list *LinkedList) Data() (eles []interface{}) {
	if list.len == 0 {
		return eles
	}
	head := list.head
	eles = append(eles, head.Data)
	for head.Next != nil {
		head = head.Next
		eles = append(eles, head.Data)
	}
	return eles
}

// Del ...
func (list *LinkedList) Del(pos int) error {
	err := list.checkPosGetDel(pos)
	if err != nil {
		return err
	}

	head := list.head
	index := 1
	for head != nil && index < pos-1 {
		head = head.Next
		index++
	}

	if pos == 1 {
		list.head = head.Next
	} else {
		head.Next = head.Next.Next
	}
	list.len--
	return nil
}

// Get ...
func (list *LinkedList) Get(pos int) (interface{}, error) {
	err := list.checkPosGetDel(pos)
	if err != nil {
		return 0, err
	}

	head := list.head
	index := 1
	for head != nil && index < pos {
		head = head.Next
		index++
	}

	return head.Data, nil
}

// Len ...
func (list *LinkedList) Len() int {
	return list.len
}

func (list *LinkedList) checkPos(pos int) error {
	if pos < 1 || pos > list.len+1 {
		return errors.New("wrong pos")
	}
	return nil
}

func (list *LinkedList) checkPosGetDel(pos int) error {
	if pos < 1 || pos > list.len {
		return errors.New("wrong pos")
	}
	return nil
}
