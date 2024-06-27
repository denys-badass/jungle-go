package main

import "fmt"

func main() {
	number := 10

	fmt.Println(number)

	change := func() {
		number++
	}

	change()
	fmt.Println(number)

	change()
	change()
	fmt.Println(number)
}
