package main

import "fmt"

type Student interface {
	Human
	Multiply(a, b float64)
}

type Human interface {
	Greet()
}

type Person struct {
}

func (p Person) Greet() {
	fmt.Println("Hello")
}

func (p Person) Multiply(a, b float64) {
	sum := a * b
	fmt.Println(sum)
}

func main() {
	var person Student = Person{}
	person.Greet()
	person.Multiply(2.5, 4)
}
