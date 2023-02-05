package main

//- O que é yield? runtime.Gosched(): troca de processamento entre rotinas
//- Race condition:
//Função 1       var     Função 2
//Lendo: 0   →   0
//Yield          0   →   Lendo: 0
//var++: 1               Yield
//Grava: 1   →   1       var++: 1
//1   ←   Grava: 1
//Lendo: 1   ←   1
//Yield          1   →   Lendo: 1
//var++: 2               Yield
//Grava: 2   →   2       var++: 2
//2   ←   Grava: 2
//no exemplo acima o resultado deveria ser 4 pq incrementei a variável 4x
//porém, durante as trocas de processamento o programa não conseguia pegar o valor atual da variável compartilhada
//- E é por isso que vamos ver mutex, atomic e, por fim, channels.
//mutex: é uma trava na variável. Se eu leio a variável enquanto eu não parar de usá-la a mesma não aceitará alterar o seu valor
//atomic: mais baixo nível, não vamos usar
//canais: é o que vamos usar
func main() {
}
