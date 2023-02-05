package main

import "fmt"

//Se não declararmos o tipo da cosntante o programa não irá definí-lo enquanto a
//mesma não ser usada, diferente da variável
const (
	x = 10
)

var c int
var d float64

func main() {
	fmt.Printf("x: %v, %T\n", x, x)

	c = x //vai dar ceto pq usei x como int
	fmt.Printf("c: %v, %T\n", c, c)

	d = x //vai dar ceto pq qd não declarada funciona como int ou float
	fmt.Printf("d: %v, %T\n", d, d)

	//x = d //não dá pq x é int

}
