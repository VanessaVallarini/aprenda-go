package main

import "fmt"

func main() {
	//for: inicialização, condição e pós
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	fmt.Println()

	//for: um dentro do outro
	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Hora:", horas)
		for minutos := 0; minutos < 60; minutos++ {
			fmt.Print(" ", minutos)
		}
		fmt.Println("\n")
	}

	//while: não existe em go, mas dá pra fazer da seguinte forma
	while := 0
	for while < 10 {
		fmt.Println("é menor que dez")
		while++
	}

	//continue
	for y := 0; y < 10; y++ {
		if y%2 == 0 {
			continue //continua dentro do for e não executa o println
		}
		fmt.Println(y)
	}

	//breack
	for y := 0; y < 10; y++ {
		fmt.Println(y)
		if y == 3 {
			break //para o for
		}
	}

}
