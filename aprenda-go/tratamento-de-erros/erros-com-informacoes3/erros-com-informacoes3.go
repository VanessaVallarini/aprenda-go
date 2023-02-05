package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//criando um log do tipo erro e um erro pq Errorf retorna um erro
		return 0, fmt.Errorf("norgate math again: square root of negative number: %v", f)
	}
	return 42, nil
}
