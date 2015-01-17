package main

import (
	"fmt"
	"github.com/markpitchless/stringutil"
)

func main() {
	for n := 0; n < 10; n = n + 1 {
		fmt.Printf(stringutil.Reverse("Hello, World\n"))
		fmt.Printf(stringutil.Reverse("!oG ,olleH\n"))
	}
	fmt.Printf("\n")
}
