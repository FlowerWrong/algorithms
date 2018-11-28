package sort

import (
	"testing"

	"github.com/FlowerWrong/algorithm/go/sort"
	"github.com/stretchr/testify/assert"
)

// go test *_test.go -test.bench=".*"

func TestMergeSort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	input := []int{9, 11, 3, 7, 8, 4, 6, 1, 38, 5, 40}
	output := []int{1, 3, 4, 5, 6, 7, 8, 9, 11, 38, 40}
	r := len(input) - 1
	sort.MergeSort(input, 0, r)
	assert.Equal(t, input, output, "they should be equal")
}

func BenchmarkMergeSort(b *testing.B) {
	input := []int{9, 11, 3, 7, 8, 4, 6, 1, 38, 5, 40}
	r := len(input) - 1
	for i := 0; i < b.N; i++ {
		sort.MergeSort(input, 0, r)
	}
}
