package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const strOriginal = "Hello, OTUS!"

func main() {
	strReversed := stringutil.Reverse(strOriginal)
	fmt.Println(strReversed)
}
