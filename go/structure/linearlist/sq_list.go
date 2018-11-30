package linearlist

import "errors"

// 顺序表

// Elem ...
type Elem int

// SqList ...
type SqList struct {
	maxSize int
	len     int
	data    []Elem
}

// NewSqList ...
func NewSqList(maxSize int) *SqList {
	return &SqList{maxSize: maxSize, len: 0, data: make([]Elem, maxSize)}
}

// IsEmpty ...
func (list *SqList) IsEmpty() bool {
	return list.len == 0
}

// IsFull ...
func (list *SqList) IsFull() bool {
	return list.len == list.maxSize
}

// Insert ...
func (list *SqList) Insert(pos int, e Elem) error {
	err := list.checkPos(pos)
	if list.checkPos(pos) != nil {
		return err
	}

	for index := list.len; index > pos-1; index-- {
		list.data[index] = list.data[index-1]
	}
	list.data[pos-1] = e
	list.len++
	return nil
}

// Append ...
func (list *SqList) Append(e Elem) error {
	if list.IsFull() {
		return errors.New("list is full")
	}
	list.data[list.len] = e
	list.len++
	return nil
}

// Clear ...
func (list *SqList) Clear() {
	list.len = 0
}

// Del ...
func (list *SqList) Del(pos int) error {
	err := list.checkPos(pos)
	if list.checkPos(pos) != nil {
		return err
	}

	for index := pos - 1; index < list.len-1; index++ {
		list.data[index] = list.data[index+1]
	}
	list.data[list.len-1] = 0
	list.len--
	return nil
}

// Data ...
func (list *SqList) Data() []Elem {
	return list.data[:list.len]
}

// Get ...
func (list *SqList) Get(pos int) (Elem, error) {
	err := list.checkPos(pos)
	if list.checkPos(pos) != nil {
		return 0, err
	}

	return list.data[pos-1], nil
}

// Len ...
func (list *SqList) Len() int {
	return list.len
}

func (list *SqList) checkPos(pos int) error {
	if pos < 1 || pos > list.len {
		return errors.New("wrong pos")
	}
	return nil
}
