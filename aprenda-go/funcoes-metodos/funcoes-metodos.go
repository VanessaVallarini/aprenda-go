package main

import "fmt"

func main() {
	mauricio := Pessoa{"Maurício", 30}
	mauricio.oiBomDia()

}

//func (receiver) identifier(parameters) (returns) { code }
//receiver faz com que a função só seja executada se existir o receiver
type Pessoa struct {
	nome  string
	idade int
}

func (p Pessoa) oiBomDia() {
	fmt.Println(p.nome, "diz bom dia!")
}
