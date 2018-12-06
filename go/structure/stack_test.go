package structure

import (
	"testing"

	"github.com/FlowerWrong/algorithm/go/structure"
	"github.com/stretchr/testify/assert"
)

// go test *_test.go -test.bench=".*"

func TestStack(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	s := structure.NewStack()
	assert.Equal(t, true, s.IsEmpty(), "they should be equal")

	_, err := s.Pop()
	assert.NotEqual(t, nil, err, "they should be equal")

	s.Push(1)
	_, err = s.Pop()
	assert.Equal(t, nil, err, "they should be equal")
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	assert.Equal(t, 5, s.Len(), "they should be equal")

	e, err := s.Pop()
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, 5, e.(int), "they should be equal")
	assert.Equal(t, 4, s.Len(), "they should be equal")
}
