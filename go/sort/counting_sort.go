package sort

// CountingSort 计数排序 O(n)
func CountingSort(input []int, k int) []int {
	c := make([]int, k+1)
	length := len(input)
	output := make([]int, length, length)

	for j := 1; j <= length; j++ {
		key := input[j-1]
		c[key]++
	}
	// c[i]表示input中值为i的元素个数

	for i := 1; i <= k; i++ {
		c[i] += c[i-1]
	}
	// c[i]表示有多少个元素小于等于i

	for _, v := range input {
		output[c[v]-1] = v
		c[v] = c[v] - 1
	}

	return output
}
