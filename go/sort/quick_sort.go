package sort

// QuickSort O(n * lgn) -> O(n * n) p <= q < r, p,q,r from 1 to array length; it is in place.
func QuickSort(input []int, p, r int) {
	if p < r {
		q := partition(input, p, r)
		QuickSort(input, p, q-1)
		QuickSort(input, q+1, r)
	}
}

func partition(input []int, p, r int) int {
	// 随机化版本 tmp := input[random(p, r) - 1]
	tmp := input[r-1]
	i := p - 1

	for j := p; j <= r-1; j++ {
		if input[j-1] <= tmp {
			i++
			input[i-1], input[j-1] = input[j-1], input[i-1]
		}
	}

	input[i], input[r-1] = input[r-1], input[i]
	return i + 1
}
