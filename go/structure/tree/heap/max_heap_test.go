package heap

import (
	"log"
	"testing"

	"github.com/FlowerWrong/algorithm/go/structure/tree/heap"
	"github.com/stretchr/testify/assert"
)

// go test *_test.go -test.bench=".*"

func TestCompleteBinaryTree(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	arr := []interface{}{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	h := heap.NewMaxHeap(arr)
	assert.Equal(t, 2, h.ParentIndex(5), "they should be equal")
	assert.Equal(t, 3, h.ParentIndex(6), "they should be equal")
	log.Println(h.Tree().Arr())
}
