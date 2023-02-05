package main

import "fmt"

func main() {
	x := retornaUmaFuncao()
	y := x(3)
	fmt.Println(y)
	//sem precisar usar a variável x
	fmt.Println(retornaUmaFuncao()(5))

}

func retornaUmaFuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
