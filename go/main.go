package main

import (
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/FlowerWrong/algorithm/go/sort"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())

	t1 := time.Now()

	// input := []int{1, 7, 4, 9, 6, 6}
	// input := g(5)
	// log.Println(input)
	// sort.InsertionSort(input)
	// log.Println(input)

	input := []int{1, 3, 5, 7, 8, 4, 6, 9, 11, 30, 40}
	log.Println(input)
	sort.MergeSort(input, 0, (len(input) - 1))
	log.Println(input)

	// input := []int{13, -3, -15, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	// log.Println(question.MaximumSubarrayNN(input))

	t2 := time.Now()
	log.Println(t2.Sub(t1))
}

func g(count int) []int {
	arr := make([]int, count)
	for index := 0; index < count; index++ {
		arr[index] = rand.Intn(100)
	}
	return arr
}
