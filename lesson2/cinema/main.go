package main

import "fmt"

var (
	userAge int
	minAge  = 12
)

func main() {
	fmt.Print("Please specify your age: ")
	fmt.Scan(&userAge)

	if userAge < minAge {
		fmt.Println("You cannot visit our cinema without your parents")
	} else {
		fmt.Println("Welcome to our cinema!!!")
	}
}
