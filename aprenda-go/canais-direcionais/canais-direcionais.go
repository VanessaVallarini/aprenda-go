package main

import "fmt"

//canais direcionais
//1 só recebe informação
//1 só envia informação
func main() {
	//ex1
	canal := make(chan int) //canal bidirecional

	//as funções abaixo tem que rodar de maneira concorrente, portanto, uma das duas tem que ser uma go func
	go send(canal)

	receive(canal)

	//ex2
	c := make(chan int)    //bidirecional
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// general to specific converts
	//específico para específico não rola
	//específico para geral tb não rola
	//só atribuir tb não rola, tem que converter (c = cr)
	fmt.Println("-----")
	fmt.Printf("c\t%T\n", (<-chan int)(c)) //convertendo o canal bidirecional para um canal receive
	fmt.Printf("c\t%T\n", (chan<- int)(c)) //convertendo o canal bidirecional para um canal send
}

func send(s chan<- int) {
	s <- 42
}

func receive(r <-chan int) {
	fmt.Println("O valor recebido do canal foi:", <-r)
}
