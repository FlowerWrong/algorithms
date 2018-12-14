package tree

import (
	"log"
	"testing"

	"github.com/FlowerWrong/algorithm/go/structure/tree"
	"github.com/stretchr/testify/assert"
)

// go test *_test.go -test.bench=".*"

func TestCompleteBinaryTree(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	tt := tree.NewCompleteBinaryTree()
	tt.Push(1)
	tt.Push(2)
	tt.Push(3)
	assert.Equal(t, 1, tt.Root.Data, "they should be equal")
	assert.Equal(t, 2, tt.Root.LeftChild.Data, "they should be equal")
	assert.Equal(t, 3, tt.Root.RightChild.Data, "they should be equal")
	tt.Push(4)
	tt.Push(5)
	tt.MiddleOrder()
	log.Println("==================")
	tt.PreOrder()
	log.Println("==================")
	tt.PostOrder()
	log.Println("==================")
	err := tt.LevelOrder()
	if err != nil {
		log.Println(err)
	}
	log.Println("==================")
	assert.Equal(t, 5, tt.Len(), "they should be equal")

	f, err := tt.IsCompleteBinaryTree()
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, true, f, "they should be equal")

	tt2 := tree.NewCompleteBinaryTreeFromArr([]interface{}{1, 2, 3, 4, 5, 6})
	assert.Equal(t, 6, tt2.Len(), "they should be equal")
	f, err = tt2.IsCompleteBinaryTree()
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, true, f, "they should be equal")

	log.Println(tt2.Arr())

	n1, err := tt2.Get(4)
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, 4, n1.Data, "they should be equal")

	n2, err := tt2.Get(1)
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, 1, n2.Data, "they should be equal")

	n3, err := tt2.Get(3)
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, 3, n3.Data, "they should be equal")

	n5, err := tt2.Get(5)
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, 5, n5.Data, "they should be equal")

	err = tt2.Set(5, 100)
	assert.Equal(t, nil, err, "they should be equal")
	n5, err = tt2.Get(5)
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, 100, n5.Data, "they should be equal")
}
