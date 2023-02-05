package main

import "fmt"

func main() {

	x := 15

	if x > 100 {
		fmt.Println("x é maior que 100")
	} else if x < 10 {
		fmt.Println("x é menor que 10")
	} else {
		fmt.Println("x não é menor que 10 e nem maior que 100")
	}

}
