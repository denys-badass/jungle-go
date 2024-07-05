package main

import "fmt"

func main() {
	printLog("Hello")
}

func printLog(log interface{}) {
	fmt.Printf("Log: %v\n", log)
}
