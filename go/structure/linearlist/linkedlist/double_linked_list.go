package linkedlist

import (
	"errors"
)

// 循环双链表

// Elem ...
// type Elem int

// DoubleLinkedNode ...
type DoubleLinkedNode struct {
	data Elem
	Next *DoubleLinkedNode
	Pre  *DoubleLinkedNode
}

// NewDoubleLinkedNode ...
func NewDoubleLinkedNode() *DoubleLinkedNode {
	list := &DoubleLinkedNode{}
	list.Next = list
	list.Pre = list
	return list
}

// IsEmpty ...
func (list *DoubleLinkedNode) IsEmpty() bool {
	return list.Next == list && list.Pre == list
}

// Clear ...
func (list *DoubleLinkedNode) Clear() {
	list.Next = list
	list.Pre = list
}

// Len ...
func (list *DoubleLinkedNode) Len() int {
	i := 0
	p := list.Next
	for p != list {
		p = p.Next
		i++
	}
	return i
}

// Data ...
func (list *DoubleLinkedNode) Data() (eles []Elem) {
	if list.Len() == 0 {
		return eles
	}
	p := list.Next
	for p != list {
		eles = append(eles, p.data)
		p = p.Next
	}
	return eles
}

// Get ...
func (list *DoubleLinkedNode) Get(pos int) (Elem, error) {
	err := list.checkPosGetDel(pos)
	if err != nil {
		return 0, err
	}

	i := 1
	p := list.Next
	for p != list && i < pos {
		p = p.Next
		i++
	}

	return p.data, nil
}

// Insert ...
func (list *DoubleLinkedNode) Insert(pos int, e Elem) error {
	err := list.checkPos(pos)
	if err != nil {
		return err
	}

	i := 2
	p := list.Next
	for p != list && i < pos {
		p = p.Next
		i++
	}

	node := &DoubleLinkedNode{data: e}
	node.Next = p.Next
	node.Pre = p

	p.Next = node

	return nil
}

// Del ...
func (list *DoubleLinkedNode) Del(pos int) error {
	err := list.checkPosGetDel(pos)
	if err != nil {
		return err
	}

	i := 1
	p := list.Next
	for p != list && i < pos {
		p = p.Next
		i++
	}

	pre := p.Pre
	next := p.Next
	pre.Next = next
	next.Pre = pre

	return nil
}

// Append ...
func (list *DoubleLinkedNode) Append(e Elem) error {
	p := list.Pre

	node := &DoubleLinkedNode{data: e}
	node.Next = list
	node.Pre = p

	p.Next = node
	list.Pre = node
	return nil
}

func (list *DoubleLinkedNode) checkPos(pos int) error {
	if pos < 1 || pos > list.Len()+1 {
		return errors.New("wrong pos")
	}
	return nil
}

func (list *DoubleLinkedNode) checkPosGetDel(pos int) error {
	if pos < 1 || pos > list.Len() {
		return errors.New("wrong pos")
	}
	return nil
}
