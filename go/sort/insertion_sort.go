package sort

// InsertionSort O(n * n) in place
func InsertionSort(input []int) {
	for i := 1; i < len(input); i++ {
		tmp := input[i]

		for j := i - 1; j >= 0; j-- {
			if tmp < input[j] {
				input[j+1] = input[j]
				input[j] = tmp
			} else {
				break
			}
		}
	}
}
