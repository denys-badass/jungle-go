package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := fillSlice(25)

	fmt.Println(nums)
	fmt.Println(countOfNumber(nums, -4))
	fmt.Println(countOfNumber(nums, 0))
	fmt.Println(countOfNumber(nums, 9))
}

func fillSlice(size int) []int {
	result := make([]int, 0, size)
	for i := 0; i < size; i++ {
		result = append(result, rand.Intn(20)-10)
	}
	return result
}

func countOfNumber(nums []int, find int) int {
	numsMap := make(map[int]int)

	for _, num := range nums {
		numsMap[num]++
	}
	return numsMap[find]
}
