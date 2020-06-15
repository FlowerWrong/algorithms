package leetcode

// TwoSum https://leetcode-cn.com/problems/two-sum/ O(n) nums重复数据 FIXME
// 优化: 一次for循环搞定
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	res := make([]int, 2)
	for index, v := range nums {
		m[v] = index
	}
	for index, v := range nums {
		diff := target - v
		if _, ok := m[diff]; ok && index != m[diff] {
			res[0] = index
			res[1] = m[diff]
			return res
		}
	}
	return []int{}
}
