package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//Ex3: cria um arquivo e se der erro exibe o erro na tela e encerra
	f, err := os.Create("names.txt")
	//se usar a função Create, vai criar o arquivo ao invés de tentar abrir, daí não da erro
	//se usar a função Open dá erro pq não encontra o arquivo
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	defer fmt.Println("rodou a função que estava em defer")

	r := strings.NewReader("James Bond") //colocou o texto dentro do arquivo

	io.Copy(f, r) //pega o que eu posso ler e coloca o que posso escrever
}
