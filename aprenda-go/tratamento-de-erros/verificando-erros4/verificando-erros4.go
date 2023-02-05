package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//Ex4:
	f, err := os.Open("names.txt") //abrindo o arquivo que criei
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close() //fecha o arquivo
	defer fmt.Println("rodou a função que estava em defer")

	bs, err := ioutil.ReadAll(f) //retorna um slice de bytes
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs)) //transforma em string
}
