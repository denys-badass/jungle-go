package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := fillSlice(15)

	fmt.Println(nums)
	selectionSort(nums)
	fmt.Println(nums)

}

func selectionSort(nums []int) {

	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

func fillSlice(size int) []int {
	result := make([]int, 0, size)
	for i := 0; i < size; i++ {
		result = append(result, rand.Intn(20)-10)
	}
	return result
}
