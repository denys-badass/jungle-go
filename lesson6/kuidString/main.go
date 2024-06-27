package main

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

func main() {
	uid := ksuid.New()
	fmt.Println(uid)
}
