package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMaxDelimeter(12, 4))
	fmt.Println(findMaxDelimeter(15, 25))
	fmt.Println(findMaxDelimeter(1207, 1349))
	fmt.Println(findMaxDelimeter(9960, 1320))

}

func findMaxDelimeter(a, b int) int {
	min := int(math.Min(float64(a), float64(b)))
	max := int(math.Max(float64(a), float64(b)))
	if max%min == 0 {
		return min
	}
	return findMaxDelimeter(min, max%min)
}
