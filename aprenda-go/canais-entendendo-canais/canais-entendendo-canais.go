package main

import "fmt"

//canais é uma maneira de transferência de dados entre go rotines
//lembrando que a main tb é uma go rotine

func main() {
	canal := make(chan int)

	//exemplo1
	//canal <- 42
	//fmt.Println(<-canal) //essa linha não executa pq eu fico esperando alguém pegar o 4 que inseri no canal

	//exemplo2
	//pra usar o canal eu preciso de 1 go routine pra inseri e outra pra remover
	go func() {
		canal <- 42
	}()

	//exemplo 3
	//com o buffer funciona o exemplo1 pq eu inseri no canal e to falando que não é necessário que a mensagem seja comsumida
	//instantaneamente
	//via de regra não usamos buffer
	canal2 := make(chan int, 1)
	canal2 <- 43

	fmt.Println(<-canal)
	fmt.Println(<-canal2)
}
