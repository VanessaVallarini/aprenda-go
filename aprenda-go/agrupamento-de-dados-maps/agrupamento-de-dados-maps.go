package main

import "fmt"

func main() {
	//são pares de chave valor
	//não tem ordem
	//performance extremamente boa
	amigos := map[string]int{
		"alfredo": 5551234,
		"joana":   9996674,
	}
	amigos["gopher"] = 4444444
	sera := amigos["fantasma"]
	fmt.Println(sera) //imprime 0
	if sera, ok := amigos["fantasma"]; ok {
		fmt.Println(sera)
	} else {
		fmt.Println("Esse amigo não existe")
	}
	fmt.Println(amigos["gopher"])
	fmt.Println(amigos)

	fmt.Println()

	//map com range
	//já que não é ordenado recomenda-se ler o map com range
	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		18:  "idade de ir pra festa",
	}
	fmt.Println(qualquercoisa)
	for key, value := range qualquercoisa {
		fmt.Println(key, value)
	}
	//remover um elemento
	delete(qualquercoisa, 123)
	fmt.Println(qualquercoisa)
}
