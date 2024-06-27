package main

import (
	"fmt"
	"math"
	"math/rand"
	// "slices"
)

func main() {
	nums := fillSlice(10)

	fmt.Println("nums:", nums)
	fmt.Println("min:", min(nums))
	fmt.Println("max:", max(nums))

	// OR
	// fmt.Println("min:", slices.Min(nums))
	// fmt.Println("max:", slices.Max(nums))
}

func fillSlice(size int) []int {
	result := make([]int, 0, size)
	for i := 0; i < size; i++ {
		result = append(result, rand.Intn(20)-10)
	}
	return result
}

func max(nums []int) int {
	max := math.MinInt

	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func min(nums []int) int {
	min := math.MaxInt

	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}
