package main

import "fmt"

func main() {
	fmt.Println(multiply(2, 2))
	fmt.Println(multiply(2.5, 3))
	fmt.Println(multiply(0, -1))
}

func multiply(a, b float64) float64 {
	return a * b
}
