package main

import "fmt"

func main() {
	//multi-dimencional é um slice de slices
	ss := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(ss[0][0])
	fmt.Println(ss[0][1])
	fmt.Println(ss[0][2])
	fmt.Println()
	fmt.Println(ss[1][0])
	fmt.Println(ss[1][1])
	fmt.Println(ss[1][2])
	fmt.Println()
	fmt.Println(ss[2][0])
	fmt.Println(ss[2][1])
	fmt.Println(ss[2][2])
}
