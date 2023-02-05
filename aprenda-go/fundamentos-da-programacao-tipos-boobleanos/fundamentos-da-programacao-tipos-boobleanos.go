package main

import "fmt"

//resultado das operações são true ou false: ==, <=, >=, !=, >, <
//doc: https://go.dev/ref/spec#Boolean_types
func main() {
	x := 10
	fmt.Println(x < 100)
	fmt.Println(x == 100)
	fmt.Println(x > 100)
}
