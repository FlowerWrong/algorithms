package tree

import (
	"log"
	"testing"

	"github.com/FlowerWrong/algorithm/go/structure/tree"
	"github.com/stretchr/testify/assert"
)

// go test *_test.go -test.bench=".*"

func TestBinarySearchTree(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	binaryTree := tree.NewBinarySearchTree(50)
	binaryTree.Insert(20)
	binaryTree.Insert(10)
	binaryTree.Insert(100)
	binaryTree.Insert(60)
	binaryTree.Insert(70)
	binaryTree.Insert(5)
	binaryTree.Insert(35)
	binaryTree.Insert(40)
	log.Println(binaryTree.GetAll())

	assert.Equal(t, true, binaryTree.Contains(20), "they should be equal")
	assert.Equal(t, true, binaryTree.Contains(60), "they should be equal")
	assert.Equal(t, false, binaryTree.Contains(32), "they should be equal")

	log.Println(binaryTree.FindMin())
	log.Println(binaryTree.FindMinNode().Value())
	assert.Equal(t, 5, binaryTree.FindMin(), "they should be equal")
	assert.Equal(t, 5, binaryTree.FindMinNode().Value(), "they should be equal")

	assert.Equal(t, 100, binaryTree.FindMax(), "they should be equal")
	assert.Equal(t, 100, binaryTree.FindMaxNode().Value(), "they should be equal")
	log.Println(binaryTree.FindMax())
	log.Println(binaryTree.FindMaxNode().Value())

	binaryTree.Remove(20)
	log.Println(binaryTree.GetAll())
}
