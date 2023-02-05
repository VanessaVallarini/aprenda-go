package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

//como executamos concorrência? Com goroutines
func main() {
	//nesse exemplo a func1 não será executada pq a função main encerrará antes da execução da função 1
	//go func1()

	fmt.Println(runtime.NumCPU())       //mostrar qts processadores eu tenho
	fmt.Println(runtime.NumGoroutine()) //número de goroutines antes de executar (nesse momento tenho só a função main)

	//a função 2 só é executada pq estou usando sync
	wg.Add(2) //quantidade de goroutines que serão executadas
	go func2()
	go func3()

	fmt.Println(runtime.NumGoroutine()) //número de goroutines depois de executar
	wg.Wait()                           //espera a execução

}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1:", i)
		time.Sleep(20) //só pra ver a execução de forma paralela
	}
}

func func2() {
	for i := 0; i < 100; i++ {
		fmt.Println("func2:", i)
		time.Sleep(20) //só pra ver a execução de forma paralela
	}
	wg.Done() //pronto, finalizou a execução
}

func func3() {
	for i := 0; i < 20; i++ {
		fmt.Println("func3:", i)
	}
	wg.Done() //pronto, finalizou a execução
}
