package main

import (
	"context"
	"fmt"
	"time"
)

//igual ao exemplo 4, só muda que o 50*time.Millisecond passo como parâmetro direto
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

}
