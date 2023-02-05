package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("context:\t", ctx)           //vendo o que tem no context
	fmt.Println("context err:\t", ctx.Err()) //vendo se tem algum erro
	fmt.Printf("context type:\t%T\n", ctx)   //vendo o tipo
	fmt.Println("----------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context:\t", ctx)           //vendo o que tem no context
	fmt.Println("context err:\t", ctx.Err()) //vendo se tem algum erro
	fmt.Printf("context type:\t%T\n", ctx)   //vendo o tipo
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("----------")

	cancel()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err()) //agora vai ter 1 erro pq eu dei um cancel
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("----------")
}
