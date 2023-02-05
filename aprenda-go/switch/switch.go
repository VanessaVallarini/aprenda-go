package main

import "fmt"

func main() {

	x := 5

	//executa a primeira opção e sai do switch
	switch {
	case x < 5:
		fmt.Println("x é menor que 5")
	case x == 5:
		fmt.Println("x é igual a 5")
	case x > 5:
		fmt.Println("x é maior que 5")
	}

	fmt.Println("")

	//pode ser com base no valor de uma variável
	quemTaNoEscritorioHoje := "zezinho"
	switch quemTaNoEscritorioHoje {
	case "zezinho":
		fmt.Println("Hoje quem tá no escritório é o Zezinho")
	case "marquinhos":
		fmt.Println("Hoje quem tá no escritório é o Marquinhos")
	case "joana":
		fmt.Println("Hoje quem tá no escritório é a Joana")
	case "maria":
		fmt.Println("Hoje quem tá no escritório é a Maria")
	}

	fmt.Println("")

	//se uma condição for verdadeira executa ela e a seguinte dela
	switch quemTaNoEscritorioHoje {
	case "zezinho":
		fmt.Println("Hoje quem tá no escritório é o Zezinho")
		fallthrough
	case "marquinhos":
		fmt.Println("Sem pre que o Zezinho vem, o Marquinhos vem também")
	case "joana":
		fmt.Println("Hoje quem tá no escritório é a Joana")
	case "maria":
		fmt.Println("Hoje quem tá no escritório é a Maria")
	}

	fmt.Println("")

	//default
	quemTaNoEscritorioHoje = ""
	switch quemTaNoEscritorioHoje {
	case "zezinho":
		fmt.Println("Hoje quem tá no escritório é o Zezinho")
		fallthrough
	case "marquinhos":
		fmt.Println("Sem pre que o Zezinho vem, o Marquinhos vem também")
	case "joana":
		fmt.Println("Hoje quem tá no escritório é a Joana")
	case "maria":
		fmt.Println("Hoje quem tá no escritório é a Maria")
	default:
		fmt.Println("Tá vazio, ou a balada tava ótima, ou é feriado")
	}

	fmt.Println("")

	//cases compostos
	quemTaNoEscritorioHoje = "marquinhos"
	switch quemTaNoEscritorioHoje {
	case "zezinho", "marquinhos":
		fmt.Println("Sem pre que o Zezinho vem, o Marquinhos vem também")
	case "joana":
		fmt.Println("Hoje quem tá no escritório é a Joana")
	case "maria":
		fmt.Println("Hoje quem tá no escritório é a Maria")
	default:
		fmt.Println("Tá vazio, ou a balada tava ótima, ou é feriado")
	}

}
