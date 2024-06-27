package main

import "fmt"

func main() {
	startSlice := []string{"Some", "thing", "into", "slice"}
	fmt.Printf("%#v\n", addToSlice(startSlice, 10))
	fmt.Printf("%#v\n", addToSlice(startSlice, 8))
	fmt.Printf("%#v\n", addToSlice(startSlice, 2))
	fmt.Printf("%#v\n", addToSlice(startSlice, 15))
}

func addToSlice(s []string, maxLen int) []string {
	if maxLen < len(s) {
		return s[:maxLen]
	}

	result := make([]string, maxLen)
	copy(result, s)
	return result
}
