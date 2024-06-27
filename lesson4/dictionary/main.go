package main

import (
	"fmt"
	"strings"
)

var dictionary = map[string]string{
	"hello":   "привіт",
	"world":   "світ",
	"shop":    "крамниця",
	"brother": "брат",
	"phone":   "телефон",
	"pizza":   "піца",
	"ukraine": "Україна",
	"tea":     "чай",
	"we":      "ми",
	"go":      "йти",
	"live":    "жити",
	"this":    "це",
}

func main() {
	sentence := "Hello world this is Ukraine"

	for _, word := range strings.Fields(sentence) {
		translate, err := Translate(strings.ToLower(word))
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%q : %s\n", word, translate)
	}
}

func Translate(word string) (translate string, err error) {
	if translate, ok := dictionary[word]; ok {
		return translate, nil
	}
	return "", fmt.Errorf("%q cannot found", word)
}
