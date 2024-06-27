package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countOfWords("Some sentence for test"))
	fmt.Println(countOfWords("Some line of text, with comma and point. This is second part, of line!!!"))
}

func countOfWords(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)

	for _, word := range words {
		counts[strings.Trim(word, ".,!?*`'\"")]++
	}
	return counts
}
