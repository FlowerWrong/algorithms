package tree

import "fmt"

// 二叉搜索树 page 176
// 若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值； 若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值； 它的左、右子树也分别为二叉排序树。
// 搜索/插入/删除的复杂度等于树高，O(log(n))

// BinarySearchTree https://blog.csdn.net/skh2015java/article/details/79242356 https://wuyin.io/2018/02/05/golang-data-structure-binary-search-tree/
type BinarySearchTree struct {
	value       int
	left, right *BinarySearchTree
}

// NewBinarySearchTree 初始化并添加根节点的值
func NewBinarySearchTree(rootValue int) *BinarySearchTree {
	return &BinarySearchTree{value: rootValue}
}

// Insert 添加元素
func (t *BinarySearchTree) Insert(value int) *BinarySearchTree {
	if t == nil {
		t = NewBinarySearchTree(value)
		return t
	}
	if value < t.value {
		t.left = t.left.Insert(value)
	} else {
		t.right = t.right.Insert(value)
	}
	return t
}

// Contains 是否包含某一个元素
func (t *BinarySearchTree) Contains(value int) bool {
	if t == nil {
		return false
	}
	v := t.compareTo(value)

	if v < 0 {
		return t.left.Contains(value)
	} else if v > 0 {
		return t.right.Contains(value)
	} else {
		return true
	}
}

func (t *BinarySearchTree) compareTo(value int) int {
	return value - t.value
}

// Remove 移除元素 TODO
func (t *BinarySearchTree) Remove(value int) *BinarySearchTree {
	if t == nil {
		return t
	}
	compareResult := t.compareTo(value)
	if compareResult < 0 {
		t.left = t.left.Remove(value)
	} else if compareResult > 0 {
		t.right = t.right.Remove(value)
	} else if t.left != nil && t.right != nil { // 待删除的节点就为t
		t.value = t.right.FindMin()
		t.right = t.right.Remove(t.value)
	} else if t.left != nil {
		t = t.left
	} else {
		t = t.right
	}
	return t
}

// FindMax 查找最大值
func (t *BinarySearchTree) FindMax() int {
	if t == nil {
		fmt.Println("tree is empty")
		return -1
	}
	if t.right == nil {
		return t.value
	}
	return t.right.FindMax()
	// 也可以用下面的方法
	// return t.FindMaxNode().value
}

// FindMaxNode 查找最大的节点
func (t *BinarySearchTree) FindMaxNode() *BinarySearchTree {
	if t != nil {
		for t.right != nil {
			t = t.right
		}
	}
	return t
}

// FindMin 查找最小值
func (t *BinarySearchTree) FindMin() int {
	if t == nil {
		fmt.Println("tree is empty")
		return -1
	}
	if t.left == nil {
		return t.value
	}
	return t.left.FindMin()
	// 也可以直接用下面的方法
	// return t.FindMinNode().value
}

// FindMinNode 查找最小的节点
func (t *BinarySearchTree) FindMinNode() *BinarySearchTree {
	if t != nil {
		for t.left != nil {
			t = t.left
		}
	}
	return t
}

// Value ...
func (t *BinarySearchTree) Value() int {
	return t.value
}

// GetAll 获取树种所有的元素值（并按从小到大排序）
func (t *BinarySearchTree) GetAll() []int {
	values := []int{}
	return appendValues(values, t)
}

func appendValues(values []int, t *BinarySearchTree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
