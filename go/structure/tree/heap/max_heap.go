// Package heap https://www.geeksforgeeks.org/binary-heap/
package heap

// MaxHeap ...
type MaxHeap struct {
	size int
}

// LeftChild ...
func (h *MaxHeap) LeftChild(i int) int {
	return 2*i + 1
}

// RightChild ...
func (h *MaxHeap) RightChild(i int) int {
	return 2*i + 2
}

// BuildMaxHeap ...
func (h *MaxHeap) BuildMaxHeap() {

}
