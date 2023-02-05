package main

import "fmt"

//select de canal é igual o switch, ou seja, enquanto ele não receber um valor igual ao da expressão ele bloqueia
//se houver mais de um caso igual, o slect irá escolher um aleatoriamente

//duas go func enviando x numeros cada para um canal
func main() {
	a := make(chan int)
	b := make(chan int)
	x := 500

	go func(x int) {
		for i := 0; i < x; i++ {
			a <- i
		}
	}(x / 2)

	go func(x int) {
		for i := 0; i < x; i++ {
			b <- i
		}
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-a: //se eu receber o valor do canal a faz isso
			fmt.Println("Canal A:", v)
		case v := <-b: //se eu receber o valor do canal b faz isso
			fmt.Println("Canal B:", v)
		}
	}

}
