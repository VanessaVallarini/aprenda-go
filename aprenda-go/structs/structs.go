package main

import "fmt"

type Cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

type Pessoa struct {
	nome  string
	idade int
}

type Profissional struct {
	pessoa  Pessoa
	titulo  string
	salario int
}

func main() {
	cliente1 := Cliente{
		nome:      "João",
		sobrenome: "Silva",
		fumante:   true,
	}
	fmt.Printf("Cliente: %v %v. Fumante: %v", cliente1.nome, cliente1.sobrenome, cliente1.fumante)

	//podemos declarar de forma implícita, mas não é recomendado
	//pois, quem for ler o código vai ficar mais difícil de entender
	cliente2 := Cliente{"Maurício", "Silva", false}
	fmt.Printf("\nCliente: %v %v. Fumante: %v", cliente2.nome, cliente2.sobrenome, cliente2.fumante)

	//structs embutidos (1 dentro do outro)
	pessoa1 := Profissional{
		pessoa: Pessoa{
			nome:  "Alfredo",
			idade: 30,
		},
		titulo:  "Pizzaiola",
		salario: 10000,
	}
	fmt.Printf("\nProfissional: \n Nome: %v \n Idade: %v \n Cargo: %v \n Salário: %v",
		pessoa1.pessoa.nome, pessoa1.pessoa.idade, pessoa1.titulo, pessoa1.salario)

	//structs anônimos
	//não declaramos o tipo
	//é algo descartável
	x := struct {
		nome  string
		idade int
	}{
		nome:  "Mario",
		idade: 50,
	}
	fmt.Println(x)

}
