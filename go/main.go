package main

import (
	"./sort"
	"log"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())

	t1 := time.Now()

	// input := []int{1, 7, 4, 9, 6, 6}
	input := g(5)
	log.Println(input)
	sort.InsertionSort(input)
	log.Println(input)

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
