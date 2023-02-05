package main

import "fmt"

//iota são constantes não tipadas, mesmo exemplo usado nas constantes
//posso usar tanto como int quanto float
const (
	x = iota
)

var c int
var d float64

func main() {
	fmt.Printf("x: %v, %T\n", x, x)

	c = x //vai dar ceto pq usei x como int
	fmt.Printf("c: %v, %T\n", c, c)

	d = x //vai dar ceto pq qd não declarada funciona como int ou float
	fmt.Printf("d: %v, %T\n", d, d)

}
