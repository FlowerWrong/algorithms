package sort

import (
	"testing"

	"github.com/FlowerWrong/algorithm/go/sort"
	"github.com/stretchr/testify/assert"
)

// go test *_test.go -test.bench=".*"

func TestCountingSort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	input := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	output := []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}
	outputActual := sort.CountingSort(input, 20)
	assert.Equal(t, output, outputActual, "they should be equal")
}

func BenchmarkCountingSort(b *testing.B) {
	input := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	for i := 0; i < b.N; i++ {
		sort.CountingSort(input, 20)
	}
}
