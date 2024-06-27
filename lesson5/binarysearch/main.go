package main

import "fmt"

func main() {
	nums := []int{-9, -8, -6, -3, 0, 1, 5, 7, 12, 14, 19, 24, 39, 41, 48, 55, 60, 79, 99, 100, 101, 144}

	fmt.Println(nums)

	fmt.Println(binarySearch(nums, 101, 0, len(nums)))
	fmt.Println(binarySearch(nums, 144, 0, len(nums)))
	fmt.Println(binarySearch(nums, 0, 0, len(nums)))
	fmt.Println(binarySearch(nums, -9, 0, len(nums)))
	fmt.Println(binarySearch(nums, 24, 0, len(nums)))
	fmt.Println(binarySearch(nums, 987, 0, len(nums)))

}

func binarySearch(nums []int, value, low, high int) int {

	if low >= high {
		return -1
	}
	mid := (low + high) / 2

	switch {
	case value > nums[mid]:
		return binarySearch(nums, value, mid+1, high)
	case value < nums[mid]:
		return binarySearch(nums, value, low, mid)
	default:
		return mid
	}
}
