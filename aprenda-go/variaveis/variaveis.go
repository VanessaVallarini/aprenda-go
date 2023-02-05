//package main -> package principal do programa
package main

import "fmt"

//fun main -> onde inicia e finaliza o programa
func main() {
	//qualquer função para usar sua delaração é package + . + a função que quero usar
	//link doc func fmt.Println: https://pkg.go.dev/github.com/reactivego/rx/test/Println
	fmt.Println("Hello word")
	//o que ela retorna
	inteiro, err := fmt.Println("Hello word")
	fmt.Println(inteiro, err)
	//ignorando um retorno
	_, err2 := fmt.Println("Hello word")
	fmt.Println(err2)
	//tipos primitivos
	x := 10
	y := "strings"
	z := true
	fmt.Println(x, y, z)
}
