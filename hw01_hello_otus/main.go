package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	// Place your code here.
	strOriginal := "Hello, OTUS!"
	strReversed := stringutil.Reverse(strOriginal)

	fmt.Println(strReversed)
}
