package leetcode

// LengthOfLongestSubstring1 https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// 暴力解法
func LengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	length := 1
	for i := 0; i < len(s); i++ {
		LongestSubstr := make(map[byte]bool)
		LongestSubstr[s[i]] = true
		for j := i + 1; j < len(s); j++ {
			if _, ok := LongestSubstr[s[j]]; !ok {
				LongestSubstr[s[j]] = true
			} else {
				break
			}
		}
		if len(LongestSubstr) > length {
			length = len(LongestSubstr)
		}
	}
	return length
}

// LengthOfLongestSubstring 滑动窗口
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/hua-jie-suan-fa-3-wu-zhong-fu-zi-fu-de-zui-chang-z/
func LengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	ans := 0

	start, end := 0, 0
	for end < n {
		if i, ok := m[s[end]]; ok && i >= start {
			start = m[s[end]] + 1
		} else {
			ans = max(ans, end-start+1)
		}
		m[s[end]] = end
		end++
	}

	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
