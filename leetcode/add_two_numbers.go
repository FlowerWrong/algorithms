package leetcode

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers https://leetcode-cn.com/problems/add-two-numbers/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	t1 := l1
	t2 := l2
	tmp := res
	advanceOne := false
	for {
		if t1 != nil {
			tmp.Val += t1.Val
		}
		if t2 != nil {
			tmp.Val += t2.Val
		}

		if advanceOne {
			tmp.Val++
		}
		if tmp.Val >= 10 {
			advanceOne = true
			tmp.Val = tmp.Val % 10
		} else {
			advanceOne = false
		}

		if t1 != nil {
			t1 = t1.Next
		}
		if t2 != nil {
			t2 = t2.Next
		}

		if t1 != nil || t2 != nil {
			tmp.Next = new(ListNode)
			tmp = tmp.Next
		} else {
			if advanceOne {
				tmp.Next = new(ListNode)
				tmp = tmp.Next
				tmp.Val = 1
			}
			break
		}
	}
	return res
}
