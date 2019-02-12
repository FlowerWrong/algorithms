package question

import (
	"testing"

	"github.com/FlowerWrong/algorithm/go/question"
	"github.com/stretchr/testify/assert"
)

// go test *_test.go -test.bench=".*"

func TestFindMaxMin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	input := []int{13, -3, -15, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	expectedMax := 20
	expectedMin := -23
	max, min := question.FindMaxMin(input)
	assert.Equal(t, expectedMax, max, "they should be equal")
	assert.Equal(t, expectedMin, min, "they should be equal")
}

func BenchmarkFindMaxMin(b *testing.B) {
	input := []int{13, -3, -15, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	for i := 0; i < b.N; i++ {
		question.FindMaxMin(input)
	}
}
