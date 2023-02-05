package main

import "fmt"

func main() {
	//link doc func fmt: https://pkg.go.dev/fmt

	x := 10
	a := "Oi"
	b := "tudo bem?"

	fmt.Print("Hello word")
	//Retorna número de bytes, erro e coloca a informação na tela
	fmt.Println("\nHello word")
	//Retorna uma string e não coloca a informação na tela
	ret := fmt.Sprintln(a, b)
	fmt.Println(ret)
	fmt.Printf("x: %v, %T\n", x, x)
}
