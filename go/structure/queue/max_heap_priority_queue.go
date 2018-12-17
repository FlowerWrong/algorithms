package queue

import (
	"errors"

	"github.com/FlowerWrong/algorithm/go/structure/tree/heap"
)

// 最大优先队列

// MaxHeapPriorityQueue ...
type MaxHeapPriorityQueue struct {
	H   *heap.MaxHeap // FIXME
	Arr []interface{}
}

// NewMaxHeapPriorityQueue ...
func NewMaxHeapPriorityQueue(arr []interface{}) *MaxHeapPriorityQueue {
	h := heap.NewMaxHeap(arr)
	return &MaxHeapPriorityQueue{H: h, Arr: h.Tree().Arr()}
}

// Maximum O(1)
func (q *MaxHeapPriorityQueue) Maximum() interface{} {
	return q.Arr[0]
}

// ExtractMax O(lg n)
func (q *MaxHeapPriorityQueue) ExtractMax() (interface{}, error) {
	arr := q.Arr
	if q.H.HeapSize < 1 {
		return nil, errors.New("heap underflow")
	}
	max := arr[0]
	arr[0] = arr[q.H.HeapSize-1]
	q.H.HeapSize--
	q.H.MaxHeapify(arr, 1)
	q.Arr = arr[:q.H.HeapSize]
	return max, nil
}

// IncreaseKey i is array index O(lg n)
func (q *MaxHeapPriorityQueue) IncreaseKey(i int, key interface{}) error {
	arr := q.Arr
	if key.(int) < arr[i-1].(int) {
		return errors.New("new key is smaller than current key")
	}
	arr[i-1] = key

	for i > 1 && arr[q.H.ParentIndex(i)-1].(int) < arr[i-1].(int) {
		arr[i-1], arr[q.H.ParentIndex(i)-1] = arr[q.H.ParentIndex(i)-1], arr[i-1]
		i = q.H.ParentIndex(i)
	}
	return nil
}

// Insert O(lg n)
func (q *MaxHeapPriorityQueue) Insert(key interface{}) error {
	q.H.HeapSize++
	q.Arr = append(q.Arr, key)
	return q.IncreaseKey(q.H.HeapSize, key)
}
