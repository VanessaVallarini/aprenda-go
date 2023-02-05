package main

import "fmt"

func main() {
	fmt.Println("Soma: ", soma(2, 2))
	basica()
	params := []int{1, 2, 3, 4}
	result, cont := soma2(params...)
	fmt.Printf("Soma2: result[%v], cont[%v] ", result, cont)
	//por ser uma função variática, posso passar 0 argumentos que não vai dar erro
	result, cont = soma2()
	fmt.Printf("\nSoma2: result[%v], cont[%v] ", result, cont)
}

func soma(x, y int) int {
	return x + y
}

func basica() {
	fmt.Println("Oi, bom dia")
}

func soma2(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)
}
