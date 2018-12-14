package tree

import (
	"errors"
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

// NewCompleteBinaryTreeFromArr ...
func NewCompleteBinaryTreeFromArr(arr []interface{}) *CompleteBinaryTree {
	t := &CompleteBinaryTree{Root: nil, len: 0}
	for _, v := range arr {
		t.Push(v)
	}
	return t
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

// Get ...
func (t *CompleteBinaryTree) Get(i int) (*Node, error) {
	q := structure.NewQueue()
	node := t.Root
	if node == nil {
		return nil, errors.New("empty tree")
	}
	err := q.Enqueue(node)
	if err != nil {
		return nil, err
	}

	j := 1
	for !q.IsEmpty() && j < i {
		dNode, err := q.Dequeue()
		if err != nil {
			return nil, err
		}
		l := dNode.(*Node).LeftChild
		r := dNode.(*Node).RightChild
		if l != nil {
			err = q.Enqueue(l)
			if err != nil {
				return nil, err
			}
		}
		if r != nil {
			err = q.Enqueue(r)
			if err != nil {
				return nil, err
			}
		}
		j++
	}
	if q.IsEmpty() {
		return nil, errors.New("not found")
	}
	resNode, err := q.Dequeue()
	if err != nil {
		return nil, err
	}
	return resNode.(*Node), nil
}

// Set ...
func (t *CompleteBinaryTree) Set(i int, e interface{}) error {
	q := structure.NewQueue()
	node := t.Root
	if node == nil {
		return errors.New("empty tree")
	}
	err := q.Enqueue(node)
	if err != nil {
		return err
	}

	j := 1
	for !q.IsEmpty() && j < i {
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
		j++
	}
	if q.IsEmpty() {
		return errors.New("not found")
	}
	resNode, err := q.Dequeue()
	if err != nil {
		return err
	}
	resNode.(*Node).Data = e
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

// Arr ...
func (t *CompleteBinaryTree) Arr() (arr []interface{}) {
	q := structure.NewQueue()
	node := t.Root
	if node == nil {
		return arr
	}
	err := q.Enqueue(node)
	if err != nil {
		return arr
	}

	for !q.IsEmpty() {
		dNode, err := q.Dequeue()
		if err != nil {
			return arr
		}
		l := dNode.(*Node).LeftChild
		r := dNode.(*Node).RightChild
		if l != nil {
			err = q.Enqueue(l)
			if err != nil {
				return arr
			}
		}
		if r != nil {
			err = q.Enqueue(r)
			if err != nil {
				return arr
			}
		}

		arr = append(arr, dNode.(*Node).Data)
	}
	return arr
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
