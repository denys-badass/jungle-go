package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := fillSlice(15)

	fmt.Println(nums)
	bubbleSort(nums)
	fmt.Println(nums)

}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func fillSlice(size int) []int {
	result := make([]int, 0, size)
	for i := 0; i < size; i++ {
		result = append(result, rand.Intn(20)-10)
	}

	return result
}
