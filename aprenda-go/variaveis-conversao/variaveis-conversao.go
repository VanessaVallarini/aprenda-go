package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	x := 10
	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("x: %v, %T\n", x, x)

	y := int(b)
	fmt.Printf("y: %v, %T\n", y, y)
}
