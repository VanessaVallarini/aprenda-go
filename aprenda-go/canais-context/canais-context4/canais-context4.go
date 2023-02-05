package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//d := time.Now().Add(50 * time.Millisecond) //tempo de execução + 50 milisegundos
	d := time.Now().Add(50 * time.Second) //tempo de execução + 50 milisegundos

	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel() //depois de rodar o select isso aqui vai rodar
	//50 * time.Millisecond
	//mas como nunca vai rodar o select pq o tempo é maior que o d (50 * time.Millisecond vs 1 * time.Second)
	//retorno:context deadline exceeded

	//50 * time.Second
	//agora o select vai executar pq 50 * time.Second é maior que 1 * time.Second
	//retorno:overslept

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
