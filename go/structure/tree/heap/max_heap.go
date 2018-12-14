// Package heap https://www.geeksforgeeks.org/binary-heap/
// https://blog.csdn.net/guoweimelon/article/details/50904346
package heap

import (
	"github.com/FlowerWrong/algorithm/go/structure/tree"
)

// MaxHeap ...
type MaxHeap struct {
	tree     *tree.CompleteBinaryTree
	len      int
	HeapSize int
}

// NewMaxHeap ...
func NewMaxHeap(arr []interface{}) *MaxHeap {
	h := &MaxHeap{tree: nil, len: 0, HeapSize: 0}
	h.BuildMaxHeap(arr)
	t := tree.NewCompleteBinaryTreeFromArr(arr)
	h.tree = t
	h.len = t.Len()
	h.HeapSize = t.Len()
	return h
}

// LeftChildIndex i is from 1..len
func (h *MaxHeap) LeftChildIndex(i int) int {
	return 2 * i
}

// RightChildIndex i is from 1..len
func (h *MaxHeap) RightChildIndex(i int) int {
	return 2*i + 1
}

// ParentIndex i is from 1..len
func (h *MaxHeap) ParentIndex(i int) int {
	return i / 2
}

// Tree ...
func (h *MaxHeap) Tree() *tree.CompleteBinaryTree {
	return h.tree
}

// MaxHeapify arr is tree array, i is index, from 0
func (h *MaxHeap) MaxHeapify(arr []interface{}, i int) {
	l := h.LeftChildIndex(i)
	r := h.RightChildIndex(i)
	length := h.HeapSize
	largest := i
	if l <= length && arr[l-1].(int) > arr[i-1].(int) {
		largest = l
	}

	if r <= length && arr[r-1].(int) > arr[largest-1].(int) {
		largest = r
	}

	if largest != i {
		arr[i-1], arr[largest-1] = arr[largest-1], arr[i-1]
		h.MaxHeapify(arr, largest)
	}
}

// BuildMaxHeap ...
func (h *MaxHeap) BuildMaxHeap(arr []interface{}) error {
	h.HeapSize = len(arr)
	for i := len(arr) / 2; i >= 1; i-- {
		h.MaxHeapify(arr, i)
	}
	return nil
}

// Len ...
func (h *MaxHeap) Len() int {
	return h.len
}
