package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	res := TwoSum([]int{2, 7, 11, 15}, 9)
	if res[0] != 0 || res[1] != 1 {
		t.Error(res)
	}

	res = TwoSum([]int{3, 2, 4}, 6)
	if res[0] != 1 || res[1] != 2 {
		t.Error(res)
	}
}
