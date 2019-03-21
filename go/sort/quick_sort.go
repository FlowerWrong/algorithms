package sort

import (
	"math/rand"
	"time"
)

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

// RandomizedPartition 随机化版本
func RandomizedPartition(input []int, p, r int) int {
	// 产生 p 到 r 之间的随机数
	tmp := input[RandomInt(p, r)-1] // pivot
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

// RandomInt ...
func RandomInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
