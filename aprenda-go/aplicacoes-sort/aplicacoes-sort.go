package main

import (
	"fmt"
	"sort"
)

func main() {

	ss := []string{"abóbora", "maçã", "laranja", "beringela", "berinjela"}

	fmt.Println(ss) //antes

	sort.Strings(ss) //ordenando na ordem crescente

	fmt.Println(ss) //depois

	si := []int{123, 987, 324, 876, 234, 987, 234, 76}

	fmt.Println(si) //antes

	sort.Ints(si) //ordenando na ordem crescente

	fmt.Println(si) //depois

	//customizando a ordenação

	carros := []carro{{"Chevete", 50, 8}, {"Porsche", 300, 5}, {"Fusca", 20, 30}}

	fmt.Println("Inicial:\n", carros)

	sort.Sort(ordenarPorPotencia(carros))

	fmt.Println("Potência:\n", carros)

	sort.Sort(ordenarPorConsumo(carros))

	fmt.Println("Consumo:\n", carros)

	sort.Sort(ordenarPorLucroProDonoDoPosto(carros))

	fmt.Println("Lucro:\n", carros)

}

type carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordenarPorPotencia []carro

func (x ordenarPorPotencia) Len() int           { return len(x) }                        //retorna o tamanho do slice
func (x ordenarPorPotencia) Less(i, j int) bool { return x[i].potencia < x[j].potencia } //diz se o elemento vai antes do próximo elemento
func (a ordenarPorPotencia) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }              //troca a ordem dos elementos

type ordenarPorConsumo []carro

func (x ordenarPorConsumo) Len() int           { return len(x) }
func (x ordenarPorConsumo) Less(i, j int) bool { return x[i].consumo > x[j].consumo }
func (a ordenarPorConsumo) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ordenarPorLucroProDonoDoPosto []carro

func (x ordenarPorLucroProDonoDoPosto) Len() int           { return len(x) }
func (x ordenarPorLucroProDonoDoPosto) Less(i, j int) bool { return x[i].consumo < x[j].consumo }
func (a ordenarPorLucroProDonoDoPosto) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
