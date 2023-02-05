package main

import "fmt"

//Closure: capturamos um contexto e o utilizamos quando quisermos
func main() {

	a := i()

	fmt.Println(a()) //1
	fmt.Println(a()) //2
	fmt.Println(a()) //3

	b := i()
	//b reinicia a contagem pq x da função
	//está fora do escopo da função interna
	fmt.Println(b()) //1
	fmt.Println(b()) //2
	fmt.Println(b()) //2
}

func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
