package leetcode

import "testing"

func TestFindPeakElement(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	nums := []int{1, 2, 3, 1}
	if FindPeakElement(nums) != 2 {
		t.Error(FindPeakElement(nums))
	}

	nums = []int{1, 2, 1, 3, 5, 6, 4}
	if FindPeakElement(nums) != 1 {
		t.Error(FindPeakElement(nums))
	}
}
