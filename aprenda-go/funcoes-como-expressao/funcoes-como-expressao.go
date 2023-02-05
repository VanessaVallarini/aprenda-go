package main

import "fmt"

func main() {

	x := 387

	//atribuímos uma função a um valor de uma variável
	y := func(x int) {
		fmt.Print(x, " vezes 873648 é: ")
		fmt.Print(x * 873648)
	}

	y(x)

}
