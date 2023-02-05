package main

import "fmt"

func main() {
	c := make(chan int)

	go meuloop(10, c)
	prints(c)

}

func meuloop(t int, s chan<- int) { //recebe um canal que recebe informação
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s) //fechando o canal -> informo ao programa que desse canal não sai mais nada, portanto, vamos fechá-lo
	//e sem isso iria dar um exception pq minha função prints (range) iria ficar esperando sair mais alguma coisa do canal
	//exception : fatal error: all goroutines are asleep - deadlock!
}

func prints(r <-chan int) {
	//range é um pouco diferente para canais
	for v := range r {
		fmt.Println("Recebido do canal:", v)
	}
}
