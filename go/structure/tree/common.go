package tree

// Node ...
type Node struct {
	Data       interface{}
	LeftChild  *Node
	RightChild *Node
}
