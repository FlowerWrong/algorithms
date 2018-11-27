package question

// MaximumSubarrayN O(n)
// eg: input [13, -3, -15, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7] and output [18, 20, -7, 12]
func MaximumSubarrayN(input []int) []int {
	start, end := 0, 0
	flag := true
	maxSum, tmpSum := 0, 0
	for i := 0; i < len(input); i++ {
		tmpSum += input[i]
		if flag {
			start = i
			flag = false
		}

		if tmpSum > maxSum {
			maxSum = tmpSum
			end = i
		}

		if tmpSum < 0 {
			tmpSum = 0
			start, end = 0, 0
			flag = true
		}
	}
	return input[start:(end + 1)]
}

// MaximumSubarrayLgN O(n * lgn)
func MaximumSubarrayLgN(input []int) []int {
	return input[:]
}

// MaximumSubarrayNN O(n * n)
func MaximumSubarrayNN(input []int) []int {
	start, end := 0, 0
	tmpSum := 0
	for i := 0; i < len(input); i++ {
		s, e, m := calMaxSum(input, i)
		if m > tmpSum {
			tmpSum = m
			start, end = s, e
		}
	}
	return input[start:(end + 1)]
}

func calMaxSum(input []int, j int) (start, end int, maxSum int) {
	flag := true
	tmpSum := 0
	for i := j; i < len(input); i++ {
		tmpSum += input[i]
		if flag {
			start = i
			flag = false
		}
		if tmpSum > maxSum {
			maxSum = tmpSum
			end = i
		}
	}
	return
}
