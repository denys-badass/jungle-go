package main

import "fmt"

func main() {
	fmt.Println(multiplyOrAdd(-5))
	fmt.Println(multiplyOrAdd(0))
	fmt.Println(multiplyOrAdd(101))

}

func multiplyOrAdd(a float64) (result float64) {
	defer fmt.Println("“функція повернула результат”")
	if a > 0 {
		result = a * 2
	} else {
		result = a + 8
	}
	return
}
