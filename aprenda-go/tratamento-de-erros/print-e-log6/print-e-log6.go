package main

import (
	"fmt"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run") //isso não será executado por conta do Panicln
}

// http://godoc.org/builtin#panic

//saída:
//When os.Exit() is called, deferred functions don't run
//panic: open no-file.txt: no such file or directory

//goroutine 1 [running]:
//main.main()
//	/Users/vanessa.vallarini/go/src/aprenda-go/tratamento-de-erros/print-e-log6/print-e-log6.go:12 +0x7a
//exit status 2//
