package sort

import (
	"math"
)

// MergeSort c * n * lgn
func MergeSort(input []int, p, r int) {
	// r := len(input) - 1
	// p := 0
	if p < r {
		q1 := math.Floor(float64((p + r) / 2))
		q := int(q1)

		MergeSort(input, p, q)
		MergeSort(input, q+1, r)
		merge(input, p, int(q), r)
	}
}

// a[p:q] a[q+1:r] p <= q < r, p,q,r is array index
// 1,3,5,7,8
// 4,6,9,11,30,40
// return [1 3 4 5 6 7 8 9 11 30 40]
func merge(a []int, p, q, r int) {
	left := make([]int, (q + 1))
	right := make([]int, (r - q))
	copy(left, a[:(q+1)])
	copy(right, a[(q+1):(r+1)])

	i := 0
	j := 0
	for k := 0; k < len(a); k++ {
		if len(left) == i {
			copy(a[k:], right[j:])
			break
		}

		if len(right) == j {
			copy(a[k:], left[i:])
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
