// Package heap https://www.geeksforgeeks.org/binary-heap/
package heap

// MinHeap ...
type MinHeap struct {
	size int
}

// LeftChild ...
func (h *MinHeap) LeftChild(i int) int {
	return 2*i + 1
}

// RightChild ...
func (h *MinHeap) RightChild(i int) int {
	return 2*i + 2
}

// BuildMaxHeap ...
func (h *MinHeap) BuildMaxHeap() {

}
