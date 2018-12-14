package sort

import (
	"testing"

	"github.com/FlowerWrong/algorithm/go/sort"
	"github.com/stretchr/testify/assert"
)

// go test *_test.go -test.bench=".*"

func TestHeapSort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	input := []interface{}{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	output := []interface{}{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}
	sort.HeapSort(input)
	assert.Equal(t, output, input, "they should be equal")
}

func BenchmarkHeapSort(b *testing.B) {
	input := []interface{}{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	for i := 0; i < b.N; i++ {
		sort.HeapSort(input)
	}
}
