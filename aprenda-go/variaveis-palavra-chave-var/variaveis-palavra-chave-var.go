package main

import "fmt"

//só usamos o var para ter o acesso a nível de pacote
var p = "variável disponível a nível de pacote"

//quando declarada sem valor, só pode ser atribuído um valor dentro de um escopo
var pSemDeclararValor int

func main() {
	q := "variável disponível a nível de escopo {}"
	fmt.Printf("p: %v, %T\n", p, p)
	fmt.Printf("q: %v, %T\n", q, q)
	pSemDeclararValor = 10
	fmt.Printf("pSemDeclararValor: %v, %T\n", pSemDeclararValor, pSemDeclararValor)
}
