package main

import "fmt"

//em go os tipos são estáticos (pra sempre)
//primitivos: int, string e bool
//compostos: slice, array, struct e map
//strings interpretadas: x := "Olá \"tudo bem?'"""

func main() {
	//strings interpretadas
	x := "Olá \"tudo bem?\" \n contigo?"
	fmt.Println(x)
	//strings row a forma como escrevi
	y := `"Olá \"tudo bem?\" \n contigo?"`
	fmt.Println(y)
}
