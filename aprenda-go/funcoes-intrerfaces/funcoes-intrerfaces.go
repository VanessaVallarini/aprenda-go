package main

import "fmt"

//em go, diferente de outras linguagens, não precisamos dizer
//classe cachorro implementa interface bichoDeEstimacao
//se a classe cachorro implementar os métodos da interface bichoDeEstimacao
//já fica implícito que a classe implementa a interface
type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type Dentista struct {
	pessoa           Pessoa
	dentesarrancados int
	salário          float64
}

type Arquiteto struct {
	pessoa           Pessoa
	tipodeconstrução string
	tamanhodaloucura string
}

func (x Dentista) oibomdia() {
	fmt.Println("Meu nome é", x.pessoa.nome, "e eu já arranquei", x.dentesarrancados, "dentes, e ouve só: Bom dia!")
}

func (x Arquiteto) oibomdia() {
	fmt.Println("Meu nome é", x.pessoa.nome, "e ouve só: Bom dia!")
}

type Gente interface {
	oibomdia()
}

func serhumano(g Gente) {
	g.oibomdia()
	switch g.(type) {
	case Dentista:
		fmt.Println("Eu ganho:", g.(Dentista).salário)

	case Arquiteto:
		fmt.Println("Eu construo:", g.(Arquiteto).tipodeconstrução)
	}
}

func main() {
	mrdente := Dentista{
		pessoa: Pessoa{
			nome:      "Mister Dente",
			sobrenome: "da Silva",
			idade:     50,
		},
		dentesarrancados: 8973,
		salário:          98736.06,
	}

	mrprédio := Arquiteto{
		pessoa: Pessoa{
			nome:      "Mister Prédio",
			sobrenome: "Sobrenome",
			idade:     51,
		},
		tipodeconstrução: "Hospícios",
		tamanhodaloucura: "Demais",
	}

	serhumano(mrdente)
	fmt.Println("")
	serhumano(mrprédio)

}
