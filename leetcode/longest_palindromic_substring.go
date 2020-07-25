package leetcode

// LongestPalindrome https://leetcode-cn.com/problems/longest-palindromic-substring/
func LongestPalindrome(s string) string {
	l := len(s)

	if s == "" {
		return ""
	}
	if l == 1 {
		return s
	}
	if l == 2 && s[0] != s[1] {
		return string(s[0:1])
	}

	ans := ""

	var prev byte
	var pprev byte
	var current byte
	for i := 0; i < len(s); i++ {
		current = s[i]
		if i > 0 {
			if pprev == current {
				m := i - 3
				n := i + 1
				for {
					if m < 0 {
						ans = string([]byte{pprev, prev, current})
						break
					}
					if n > l-1 || s[m] != s[n] {
						res := string(s[m+1 : n])
						if len(res) > len(ans) {
							ans = res
						}
						break
					}

					m--
					n++
				}
			} else if prev == current {
				m := i - 2
				n := i + 1
				for {
					if m < 0 {
						ans = string([]byte{prev, current})
						break
					}
					if n > l-1 || s[m] != s[n] {
						res := string(s[m+1 : n])
						if len(res) > len(ans) {
							ans = res
						}
						break
					}

					m--
					n++
				}
			}
		}

		prev = s[i]
		if i > 0 {
			pprev = s[i-1]
		}
	}
	return ans
}
