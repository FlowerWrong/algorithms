package leetcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	if LengthOfLongestSubstring("pwwkew") != 3 {
		t.Error(LengthOfLongestSubstring("pwwkew"))
	}
}
