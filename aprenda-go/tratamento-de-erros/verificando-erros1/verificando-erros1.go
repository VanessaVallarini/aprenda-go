package main

import (
	"fmt"
)

func main() {
	//Ex1: esse é um caso onde podemos ignorar o erro
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
