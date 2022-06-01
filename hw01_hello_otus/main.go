package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	const entry = "Hello, OTUS!"

	fmt.Println(stringutil.Reverse(entry))
}
