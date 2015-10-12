package main

import (
	"fmt"
	"os"
)

// main function
func main() {
	// Read first user input string
	input := os.Args[1:2]
	fmt.Println(input)
}
