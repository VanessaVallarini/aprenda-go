package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//criando dois canais e enviando tudo para a função converge
	canal := converge(trabalho("maçã"), trabalho("pêra"))

	//lendo o canal converge
	for x := 0; x < 16; x++ {
		fmt.Println(<-canal)
	}

}

//cria um canal
//cria uma go func que manda dados pro canal
//retorna o canal
func trabalho(s string) chan string {
	canal := make(chan string)
	go func(s string, c chan string) {
		for i := 1; ; i++ { //não precisa de um parâmetro de parada pq está sendo controlado pelo for da main
			c <- fmt.Sprintf("Função %v diz: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3))) //simular um cenário real - dorme entre nada e 1 segundo
		}
	}(s, canal)
	return canal
}

//recebe 2 canais
//cria um canal novo
//cria duas go func que le os canais e passa tudo para o novo canal
//retorna o canal novo
func converge(x, y chan string) chan string {
	novo := make(chan string)
	go func() {
		for {
			novo <- <-x
		}
	}()
	go func() {
		for {
			novo <- <-y
		}
	}()
	return novo
}
