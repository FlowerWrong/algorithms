package structure

import (
	"testing"

	"github.com/FlowerWrong/algorithm/go/structure"
	"github.com/stretchr/testify/assert"
)

// go test *_test.go -test.bench=".*"

func TestSqList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	q := structure.NewQueue()
	assert.Equal(t, true, q.IsEmpty(), "they should be equal")

	_, err := q.Dequeue()
	assert.NotEqual(t, nil, err, "they should be equal")

	q.Enqueue(1)
	_, err = q.Dequeue()
	assert.Equal(t, nil, err, "they should be equal")
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	assert.Equal(t, 5, q.Len(), "they should be equal")

	e, err := q.Dequeue()
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, 1, e.(int), "they should be equal")
	assert.Equal(t, 4, q.Len(), "they should be equal")
}
