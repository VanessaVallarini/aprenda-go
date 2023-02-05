package main

import (
	"encoding/json"
	"os"
)

type pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

func main() {
	jamesbond := pessoa{
		Nome:          "James",
		Sobrenome:     "Bond",
		Idade:         40,
		Profissao:     "Agente Secreto",
		Contabancaria: 1000000.50,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(jamesbond)
	//o modelo acima elimina a etapa de salvar em uma variável e depois mandar ela pra outro lugar
	//d, err := json.Marshal(darthvader)
	//fmt.Println(string(j))

}
