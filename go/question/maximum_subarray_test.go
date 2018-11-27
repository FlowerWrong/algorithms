package question

import (
	"testing"

	"github.com/FlowerWrong/algorithm/go/question"
	"github.com/stretchr/testify/assert"
)

// go test *_test.go -test.bench=".*"

func TestMaximumSubarrayN(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	input := []int{13, -3, -15, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	expected := []int{18, 20, -7, 12}
	output := question.MaximumSubarrayN(input)
	assert.Equal(t, expected, output, "they should be equal")
}

func BenchmarkMaximumSubarrayN(b *testing.B) {
	input := []int{13, -3, -15, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	for i := 0; i < b.N; i++ {
		question.MaximumSubarrayN(input)
	}
}

func TestMaximumSubarrayNN(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	input := []int{13, -3, -15, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	expected := []int{18, 20, -7, 12}
	output := question.MaximumSubarrayNN(input)
	assert.Equal(t, expected, output, "they should be equal")
}

func BenchmarkMaximumSubarrayNN(b *testing.B) {
	input := []int{13, -3, -15, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	for i := 0; i < b.N; i++ {
		question.MaximumSubarrayNN(input)
	}
}
