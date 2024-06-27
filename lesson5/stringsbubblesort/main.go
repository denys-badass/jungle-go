package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields("Nori grape silver beet broccoli kombu beet greens fava bean potato quandong celery")

	fmt.Println(words)
	bubbleSortDesc(words)
	fmt.Println(words)
}

func bubbleSortDesc(words []string) {
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words)-1; j++ {
			if len(words[j]) < len(words[j+1]) {
				words[j], words[j+1] = words[j+1], words[j]
			}
		}
	}
}
