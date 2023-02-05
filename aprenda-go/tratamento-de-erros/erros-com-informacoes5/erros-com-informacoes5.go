package main

import (
	"fmt"
	"log"
)

//representa um erro, sem isso não conseigo criar o erro
type norgateMathError struct {
	lat  string
	long string
	err  error
}

//método que retorna uma mensagem de erro personalizada de acordo com minha struct
func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//criando um log do tipo erro e um erro pq Errorf retorna um erro
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		//retorno uma struct que contém o nosso erro
		return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
	}
	return 42, nil
}

// see use of structs with error type in standard library:
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go
