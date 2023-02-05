package main

import "fmt"

func main() {
	x := 387

	//declarar e chamar a função na mesma linha
	//muito usado para go rotines e funções descartáveis
	func(x int) {
		fmt.Print(x, " vezes 873648 é: ")
		fmt.Print(x * 873648)
	}(x)

}
