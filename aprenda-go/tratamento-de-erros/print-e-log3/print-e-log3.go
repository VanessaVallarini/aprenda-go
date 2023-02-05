package main

import (
	"fmt"
	"log"
	"os"
)

// Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
}
