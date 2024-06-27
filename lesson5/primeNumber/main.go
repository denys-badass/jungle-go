package main

import "fmt"

func main() {
	fmt.Println(findPrime(1, 25))
	fmt.Println(findPrime(43, 101))
	fmt.Println(findPrime(0, 10))
}

func findPrime(start, end int) []int {
	primes := make([]int, 0)

	for i := start; i <= end; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrime(a int) bool {
	if a <= 1 {
		return false
	}
	if a == 2 {
		return true
	}
	for i := 2; i < a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
