package main

import (
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/FlowerWrong/algorithms/leetcode"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.Println("Start LeetCode")

	log.Println(leetcode.LengthOfLongestSubstring("abbca"))
}
