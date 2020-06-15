package leetcode

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	res := AddTwoNumbers(l1, l2)
	if res.Val != 7 {
		t.Error(res)
	}
}
