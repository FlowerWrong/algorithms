package tree

import (
	"log"

	"github.com/FlowerWrong/algorithm/go/structure"
)

// 完全二叉树

// CompleteBinaryTree ...
type CompleteBinaryTree struct {
	Root *Node
	len  int
}

// NewCompleteBinaryTree ...
func NewCompleteBinaryTree() *CompleteBinaryTree {
	return &CompleteBinaryTree{Root: nil, len: 0}
}

// IsCompleteBinaryTree ...
func (t *CompleteBinaryTree) IsCompleteBinaryTree() (bool, error) {
	q := structure.NewQueue()
	node := t.Root
	if node == nil {
		return true, nil
	}
	err := q.Enqueue(node)
	if err != nil {
		return false, err
	}

	for !q.IsEmpty() {
		dNode, err := q.Dequeue()
		if err != nil {
			return false, err
		}

		if dNode == nil || dNode.(*Node) == nil {
			return false, nil
		}

		l := dNode.(*Node).LeftChild
		r := dNode.(*Node).RightChild
		if l != nil {
			err = q.Enqueue(l)
			if err != nil {
				return false, err
			}
		}
		if r != nil {
			err = q.Enqueue(r)
			if err != nil {
				return false, err
			}
		}
	}
	return true, nil
}

// Push ...
func (t *CompleteBinaryTree) Push(e interface{}) error {
	newNode := &Node{Data: e, LeftChild: nil, RightChild: nil}
	p := t.Root
	if p == nil {
		t.Root = newNode
	} else {
		pL := p
		pR := p
		iL := 0
		iR := 0

		for pL.LeftChild != nil {
			iL++
			pL = pL.LeftChild
		}

		for pR.RightChild != nil {
			iR++
			pR = pR.RightChild
		}

		if iL != iR {
			pR.RightChild = newNode
		} else {
			pL.LeftChild = newNode
		}
	}

	t.len++
	return nil
}

// Len ...
func (t *CompleteBinaryTree) Len() int {
	return t.len
}

// PreOrder ...
func (t *CompleteBinaryTree) PreOrder() {
	t.Root.preOrder()
}

func (node *Node) preOrder() {
	if node == nil {
		return
	}
	log.Println(node.Data)
	node.LeftChild.preOrder()
	node.RightChild.preOrder()
}

// MiddleOrder ...
func (t *CompleteBinaryTree) MiddleOrder() {
	t.Root.middleOrder()
}

func (node *Node) middleOrder() {
	if node == nil {
		return
	}
	node.LeftChild.middleOrder()
	log.Println(node.Data)
	node.RightChild.middleOrder()
}

// PostOrder ...
func (t *CompleteBinaryTree) PostOrder() {
	t.Root.postOrder()
}

func (node *Node) postOrder() {
	if node == nil {
		return
	}
	node.LeftChild.postOrder()
	node.RightChild.postOrder()
	log.Println(node.Data)
}

// LevelOrder ...
func (t *CompleteBinaryTree) LevelOrder() error {
	q := structure.NewQueue()
	node := t.Root
	if node == nil {
		return nil
	}
	err := q.Enqueue(node)
	if err != nil {
		log.Println(err)
		return err
	}

	for !q.IsEmpty() {
		dNode, err := q.Dequeue()
		if err != nil {
			return err
		}
		l := dNode.(*Node).LeftChild
		r := dNode.(*Node).RightChild
		if l != nil {
			err = q.Enqueue(l)
			if err != nil {
				return err
			}
		}
		if r != nil {
			err = q.Enqueue(r)
			if err != nil {
				return err
			}
		}

		log.Println(dNode.(*Node).Data)
	}
	return nil
}
