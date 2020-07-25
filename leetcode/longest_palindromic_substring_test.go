package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	s := "pwwkew"
	if LongestPalindrome(s) != "ww" {
		t.Error(LongestPalindrome(s))
	}

	s = "w"
	if LongestPalindrome(s) != "w" {
		t.Error(LongestPalindrome(s))
	}

	s = "babad"
	if LongestPalindrome(s) != "bab" {
		t.Error(LongestPalindrome(s))
	}

	s = "cbbd"
	if LongestPalindrome(s) != "bb" {
		t.Error(LongestPalindrome(s))
	}

	s = "ccc"
	if LongestPalindrome(s) != "ccc" {
		t.Error(LongestPalindrome(s))
	}

	s = "aaaa"
	if LongestPalindrome(s) != "aaaa" {
		t.Error(LongestPalindrome(s))
	}
}
