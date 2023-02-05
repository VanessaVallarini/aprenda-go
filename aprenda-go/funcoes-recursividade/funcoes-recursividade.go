package main

import "fmt"

func main() {
	fmt.Println(fatorial(6))
	fmt.Println(loops(6))

}

func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}

//mesma coisa que a função recursiva
//usa menos memória
//menos chance de travar tudo e ir pro brejo ushush
//mais seguro
func loops(x int) int {
	total := 1
	for x > 1 {
		total *= x // total = total * x
		x--
	}
	return total
}
