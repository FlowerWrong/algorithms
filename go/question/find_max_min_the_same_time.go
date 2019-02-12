package question

// FindMaxMin 同时找到最大最小值 O(1.5 * n)
func FindMaxMin(input []int) (max int, min int) {
	size := len(input)

	begin := 1
	if size%2 == 0 {
		begin = 2
		if input[0] < input[1] {
			max, min = input[1], input[0]
		} else {
			max, min = input[0], input[1]
		}
	} else {
		max, min = input[0], input[0]
	}

	var tmpMax, tmpMin int
	for i := begin; i < size; i += 2 {
		if input[i] < input[i+1] {
			tmpMax, tmpMin = input[i+1], input[i]
		} else {
			tmpMax, tmpMin = input[i], input[i+1]
		}

		if tmpMax > max {
			max = tmpMax
		}

		if tmpMin < min {
			min = tmpMin
		}
	}

	return max, min
}
