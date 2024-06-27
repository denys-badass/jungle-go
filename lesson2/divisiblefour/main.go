package main

import "fmt"

var (
	minValue = -20
	maxValue = 45
)

func main() {
	for i := minValue; i <= maxValue; i++ {
		if i == 0 || i%4 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
