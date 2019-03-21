package selection

import "github.com/FlowerWrong/algorithm/go/sort"

// RandomizedSelect o(n) 返回input[p..r]中第i小的元素 page 135 FIXME
func RandomizedSelect(input []int, p, r, i int) (res int) {
	if p == r {
		res = input[p-1]
		return
	}

	q := sort.RandomizedPartition(input, p, r)
	k := q - p + 1 // 元素个数
	if i == k {
		res = input[q-1]
	} else if i < k {
		res = RandomizedSelect(input, p, q-1, i)
	} else {
		res = RandomizedSelect(input, q+1, r, i-k)
	}
	return
}
