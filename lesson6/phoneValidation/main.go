package main

import (
	"fmt"
	"log"

	"github.com/ttacon/libphonenumber"
)

var (
	userNumber string
	region     string
)

func main() {
	fmt.Print("Enter you phone number: ")
	if _, err := fmt.Scan(&userNumber); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter you region: ")
	if _, err := fmt.Scan(&region); err != nil {
		log.Fatal(err)
	}

	phoneNumber, err := libphonenumber.Parse(userNumber, region)
	if err != nil {
		log.Fatal(err)
	}

	if libphonenumber.IsValidNumber(phoneNumber) {
		fmt.Println("You number is valid")
		return
	}

	fmt.Println("Your number isn't valid")
}
