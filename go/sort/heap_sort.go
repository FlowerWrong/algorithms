package sort

import (
	"github.com/FlowerWrong/algorithm/go/structure/tree/heap"
)

// HeapSort O(n * lgn) in place
func HeapSort(input []interface{}) {
	h := heap.NewMaxHeap(input)
	length := len(input)
	for i := length; i >= 2; i-- {
		input[0], input[i-1] = input[i-1], input[0]
		h.HeapSize--
		h.MaxHeapify(input, 1)
	}
}
