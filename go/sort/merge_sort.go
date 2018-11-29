package sort

import (
	"math"
)

// MergeSort O(n * lgn) not in place
func MergeSort(input []int, p, r int) {
	if p < r {
		q1 := math.Floor(float64((p + r) / 2))
		q := int(q1)

		MergeSort(input, p, q)
		MergeSort(input, q+1, r)
		merge(input, p, q, r)
	}
}

// a[p:q] a[q+1:r] p <= q < r, p,q,r is array index
// 1,3,5,7,8
// 4,6,9,11,30,40
// return [1 3 4 5 6 7 8 9 11 30 40]
func merge(a []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	left := make([]int, n1)
	right := make([]int, n2)
	copy(left, a[p:(q+1)])
	copy(right, a[(q+1):(r+1)])

	i := 0
	j := 0
	for k := p; k <= r; k++ {
		if len(left) == i {
			copy(a[k:(r+1)], right[j:])
			break
		}

		if len(right) == j {
			copy(a[k:(r+1)], left[i:])
			break
		}

		if left[i] <= right[j] {
			a[k] = left[i]
			i++
		} else {
			a[k] = right[j]
			j++
		}
	}
}
