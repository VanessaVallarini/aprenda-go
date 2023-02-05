package main

import (
	"fmt"
	"os"
)

// Println formats using the default formats for its operands and writes to standard output.

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened", err)
	}
}
