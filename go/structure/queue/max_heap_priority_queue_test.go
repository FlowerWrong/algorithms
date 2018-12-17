package queue

import (
	"testing"

	"github.com/FlowerWrong/algorithm/go/structure/queue"
	"github.com/stretchr/testify/assert"
)

// go test *_test.go -test.bench=".*"

func TestMaxHeapPriorityQueue(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	arr := []interface{}{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	q := queue.NewMaxHeapPriorityQueue(arr)
	assert.Equal(t, 16, q.Maximum(), "they should be equal")

	m, err := q.ExtractMax()
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, 16, m, "they should be equal")
	assert.Equal(t, len(arr)-1, q.H.HeapSize, "they should be equal")

	err = q.IncreaseKey(1, 18)
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, 18, q.Maximum(), "they should be equal")

	err = q.IncreaseKey(4, 11)
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, 11, q.Arr[1].(int), "they should be equal")

	err = q.Insert(12)
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, 12, q.Arr[1].(int), "they should be equal")
}
