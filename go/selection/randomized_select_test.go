package selection

import (
	"log"
	"testing"

	"github.com/FlowerWrong/algorithm/go/selection"
	"github.com/stretchr/testify/assert"
)

// go test *_test.go -test.bench=".*"

func TestRandomizedSelect(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	input := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}

	assert.Equal(t, 4, selection.RandomizedSelect(input, 1, 1, 1), "they should be equal")
	assert.Equal(t, 16, selection.RandomizedSelect(input, 1, 4, 10), "they should be equal")

	output := 7
	outputActual := selection.RandomizedSelect(input, 1, 10, 5)
	log.Println(outputActual)
	assert.Equal(t, output, outputActual, "they should be equal")
}

func BenchmarkRandomizedSelect(b *testing.B) {
	input := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	for i := 0; i < b.N; i++ {
		selection.RandomizedSelect(input, 1, 10, 5)
	}
}
