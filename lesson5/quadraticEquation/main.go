package main

import (
	"fmt"
	"math"
)

func main() {
	x1, x2, err := quadraticEquation(2, 5, -3) // 2x^2 + 5x - 3 = 0, x1 = -0.25, x2 = -5
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x1, x2)

	x1, x2, err = quadraticEquation(9, -30, 25) // 9x^2 - 30x + 25 = 0, x = 1 * 2/3, One root
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x1, x2)

	x1, x2, err = quadraticEquation(3, -1, 8) // 3x^2 - x + 8 = 0, Doesn't have roots
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x1, x2)
}

func quadraticEquation(a, b, c float64) (x1, x2 float64, err error) {
	if a == 0 {
		x1 = -b - c
		return x1, x1, nil
	}

	d := discriminant(a, b, c)

	switch {
	case d < 0:
		return x1, x2, fmt.Errorf("roots don`t exists")
	case d == 0:
		x1 = -b / (2 * a)
		return x1, x1, nil
	}

	x1 = (-b + math.Sqrt(d)) / (2 * a)
	x2 = (-b - math.Sqrt(d)) / (2 * a)
	return x1, x2, nil
}

func discriminant(a, b, c float64) float64 {
	return b*b - 4*a*c
}
