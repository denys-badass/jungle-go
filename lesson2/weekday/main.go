package main

import "fmt"

var (
	userDay string
)

func main() {
	fmt.Print("What day is it today: ")
	fmt.Scan(&userDay)

	switch userDay {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It`s workday")
	case "Saturday", "Sunday":
		fmt.Println("It`s weekend")
	}
}
