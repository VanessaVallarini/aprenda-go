package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r) //vai ser executado qd dar o panic e volta pra main
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

//é uma função recursiva, cada vez que é chamada é acrescentado 1 no valor do seu parâmetro que inicia em 0
func g(i int) {
	if i > 3 { //chegou no g4 é dado panic
		fmt.Println("Panicking!") //o panic executa os defers de forma reversa até encontrar o recover
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
