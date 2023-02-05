package main

import "fmt"

//ponteitos é armazenar algo na memória e utilizar ela
func main() {
	x := 0

	y := &x

	fmt.Print(x, y)

	*y++

	fmt.Println(*y) //dessa forma é bem comum utilizarmos
	fmt.Printf("%T, %T\n", x, y)
	fmt.Println(x, y)

	//quando usar?
	//se tenho uma grande quantidade de dados e não quero ficar copiando esses dados a todo momento
	//daí eu posso armazenar ele em memória e buscar ele lá quando precisar

	valorOriginal1 := 11
	valorOriginal2 := 11

	retornoFuncao1 := estarecebeovalor(valorOriginal1)
	retornoFuncao2 := estarecebeumponteiro(&valorOriginal2)

	fmt.Printf("\nSem ponteiro: Valor original: %v. Retorno da função: %v", valorOriginal1, retornoFuncao1)
	fmt.Printf("\nCom ponteiro: Valor original: %v. Retorno da função: %v", valorOriginal2, retornoFuncao2)

}

func estarecebeovalor(x int) int {
	x++
	return x
}

func estarecebeumponteiro(x *int) int {
	*x++
	return *x
}
