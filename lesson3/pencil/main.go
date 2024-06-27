package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

type Pencil struct {
	Color string
	Product
}

func (p Pencil) Print() {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Price: %v\n", p.Price)
	fmt.Printf("Color: %q\n", p.Color)
}

func main() {
	pencil := Pencil{
		Color: "grey",
		Product: Product{
			Name:  "Pencil",
			Price: 2.99,
		},
	}

	pencil.Print()
}
