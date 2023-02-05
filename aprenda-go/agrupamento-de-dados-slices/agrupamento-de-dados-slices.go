package main

import "fmt"

var x [5]int

func main() {
	//array - quase não se usa. Usamos slice
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5
	fmt.Println(x)
	fmt.Println(len(x))

	fmt.Println()

	//slice - tipo de dados composto
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	//mudando o tamanho de um slice, algo que não dá pra fazer com array
	slice2 := append(slice1, 6)
	fmt.Println(slice2)
	//alterando o valor de um índice
	slice2[2] = 31
	fmt.Println(slice2)

	fmt.Println()

	//função range
	frutas := []string{"banana", "maçã", "jaca"}
	for i, valor := range frutas {
		fmt.Printf("No índica %v temos o valor: %v \n", i, valor)
	}

	//função range jogando fora o índica
	for _, valor := range frutas {
		fmt.Println("Valor:", valor)
	}

	fmt.Println()

	//fatiando slices
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatro queijos", "marg"}
	//pegando a partir do índice 2
	fatia := sabores[2:]
	//dá pra
	//x[:], x[a:], x[:b], x[a:b]
	//a é incluso
	//b não é incluso
	fmt.Println(fatia)
	//fatiando slices -> não recomendado criar um novo slice, pois, a mudança em uma nova
	//variável pode avacalhar o slice antigo
	//recomenda-se percorrer o slice que deseja copiar com range e assim criar outro
	//ou usar a mesma variável
	//ex:
	pslice := []int{1, 2, 3, 4, 5}
	fmt.Println("primeiro slice: ", pslice)
	sslice := append(pslice[:2], pslice[4:]...)
	fmt.Println("segundo slice: ", sslice)
	fmt.Println("primeiro slice: ", pslice)
	fmt.Println()

	//deletando um item de uma fatia
	nomes := []string{"van", "fabio", "val", "maria", "viani"}
	nomes2 := append(nomes[:2], nomes[4:]...)
	fmt.Println(nomes2)

	fmt.Println()

	//append: anexando itens a um slice
	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{9, 10, 11, 12}
	fmt.Println(umaslice)
	umaslice = append(umaslice, 5, 6, 7, 8)
	fmt.Println(umaslice)
	//os pointinhos representam os itens do slice
	//pois, o slice é de int, portanto, não consigo inserir um slice dentro do slice
	//apenas os itens int do slice dentro de outro slice int
	umaslice = append(umaslice, outraslice...)
	fmt.Println(umaslice)

	fmt.Println()

	//make
	//é uma boa prática iniciar o tamanho do slice quando sabemos a quantiadde de
	//elementos que o mesmo irá ter, pois, melhora a performance já que o programa
	//não vai precisar criar um slice novo e copiar o sclice antigo e adicionar no
	//slice novo toda vez que eu adicionar um elemento novo no slice
	slice := make([]int, 5, 10)
	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, 6)
	fmt.Println(slice, len(slice), cap(slice))
}
