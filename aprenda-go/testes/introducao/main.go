package main

import "fmt"

func main() {
	x := Soma(1, 2, 3)
	y := Multiplica(10, 10)
	fmt.Println(x, y)
	fmt.Println(Add(1, 2))
}

func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total = total + v
		// total += v
	}
	return total
}

func Multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}

func Add(x, y int) int {
	s := []int{x, y}
	return s[0] + s[1]
}
