package main

import "fmt"

var p = "variável disponível a nível de pacote"

func main() {

	//:= operador curto de declaração atribui um tipo de acordo com o valor
	x := 10
	y := "Hy"

	//imprimindo o valor da variável + seu tipo
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	//= operador de atribuição
	x = 20
	fmt.Printf("x: %v, %T\n", x, x)

	//declarando uma e atribuindo em outra variável já existente
	x, z := 30, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)

	q := "variável disponível a nível de escopo {}"
	fmt.Printf("p: %v, %T\n", p, p)
	fmt.Printf("q: %v, %T\n", q, q)

	//nomes reservados que não podemos usar para nome de variávels
	//package
	//func
	//doc: https://go.dev/ref/spec#Keywords
}
