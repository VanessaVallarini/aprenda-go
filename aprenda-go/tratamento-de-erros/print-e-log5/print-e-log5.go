package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Panicln(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run") //isso não será executado por conta do Panicln
}

/*
Panicln is equivalent to Println() followed by a call to panic().
*/

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).

//saída:
//2023/02/05 19:00:37 open no-file.txt: no such file or directory
//When os.Exit() is called, deferred functions don't run
//panic: open no-file.txt: no such file or directory

//goroutine 1 [running]:
//log.Panicln({0xc000092f58?, 0xb?, 0x0?})
//	/usr/local/go/src/log/log.go:399 +0x65
//main.main()
//	/Users/vanessa.vallarini/go/src/aprenda-go/tratamento-de-erros/print-e-log5/print-e-log5.go:13 +0x7d
//exit status 2
