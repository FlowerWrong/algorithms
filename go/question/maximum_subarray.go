package question

import "math"

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
func MaximumSubarrayLgN(input []int, low, high int) (int, int, int) {
	if low == high {
		return low, high, input[low]
	}

	mid1 := math.Floor(float64((low + high) / 2))
	mid := int(mid1)
	leftLow, leftHigh, leftSum := MaximumSubarrayLgN(input, low, mid)
	rightLow, rightHigh, rightSum := MaximumSubarrayLgN(input, mid+1, high)
	crossLow, crossHigh, crossSum := FindMaxCrossingSubarray(input, low, mid, high)

	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	} else {
		return crossLow, crossHigh, crossSum
	}
}

// FindMaxCrossingSubarray ...
func FindMaxCrossingSubarray(input []int, low, mid, high int) (int, int, int) {
	sum := 0
	leftSum, rightSum := 0, 0
	maxLeft, maxRight := 0, 0
	for i := mid; i >= low; i-- {
		sum += input[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	sum = 0
	for j := mid + 1; j < high; j++ {
		sum += input[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}
	return maxLeft, maxRight, leftSum + rightSum
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
