package main

import "fmt"

// func 1 recebe x valores de canal, depois manda qualquer coisa para chan quit
// func 2 for infinito, select: case envia para canal, case recebe de quit

func main() {

	canal := make(chan int)
	quit := make(chan int)

	go recebeQuit(canal, quit)
	enviaPraCanal(canal, quit)
}

//recebe x valores de canal, depois manda qualquer coisa para chan quit
func recebeQuit(canal chan int, quit chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("Recebido:", <-canal)
	}
	quit <- 0 //sair
}

//for infinito, select: case envia para canal, case recebe de quit
func enviaPraCanal(canal chan int, quit chan int) {
	qualquercoisa := 1
	for {
		select {
		case canal <- qualquercoisa: //manda coisa pro canal
			qualquercoisa++
		case <-quit: //rece coisa do canal
			return
		}
	}
}
