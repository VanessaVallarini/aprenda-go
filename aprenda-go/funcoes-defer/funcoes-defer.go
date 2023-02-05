package main

import "fmt"

//defer: literalmente deixar para última hora
//O primeiro inserido no defer é o último a sair
//exemplos de uso:
// - fechamento de arquivo
func main() {
	defer fmt.Println("com defer vem depois")
	fmt.Println("sem defer vem antes")

	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
}
