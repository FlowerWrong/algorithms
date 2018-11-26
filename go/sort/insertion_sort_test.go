package sort

import (
	"."
	"github.com/stretchr/testify/assert"
	"testing"
)

// go test *_test.go -test.bench=".*"

func TestInsertionSort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	input := []int{1, 7, 4, 9, 6, 6}
	output := []int{1, 4, 6, 6, 7, 9}
	sort.InsertionSort(input)
	assert.Equal(t, input, output, "they should be equal")
}

func BenchmarkInsertionSort(b *testing.B) {
	input := []int{1, 7, 4, 9, 6, 6}
	for i := 0; i < b.N; i++ {
		sort.InsertionSort(input)
	}
}
