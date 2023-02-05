package main

import "fmt"

//doc: https://go.dev/ref/spec#Numeric_types
//#dica: não declare o tipo, deixe o computador escolher pq se vc acabar
//passando um valor maior que o permitido para o tamaho da vaiável declarada
//o sistema aplicará 0 ou 1 como valor
func main() {
	unit8Type := 255
	fmt.Printf("unit8Type: %v, %T\n", unit8Type, unit8Type)

	unit16Type := 65535
	fmt.Printf("unit16Type: %v, %T\n", unit16Type, unit16Type)

	unit32Type := 4294967295
	fmt.Printf("unit32Type: %v, %T\n", unit32Type, unit32Type)

	unit64Type := 1844674407370955161
	fmt.Printf("unit64Type: %v, %T\n", unit64Type, unit64Type)

	//a diferença para unit é que aceita negativos
	int8Type := -128
	fmt.Printf("int8Type: %v, %T\n", int8Type, int8Type)

	int16Type := -32768
	fmt.Printf("int16Type: %v, %T\n", int16Type, int16Type)

	int32Type := -2147483648
	fmt.Printf("int32Type: %v, %T\n", int32Type, int32Type)

	int64Type := -9223372036854775808
	fmt.Printf("int64Type: %v, %T\n", int64Type, int64Type)

	//float
	floatExample := 128.0
	fmt.Printf("floatExample: %v, %T\n", floatExample, floatExample)

	//rune é um int da tabelas ascii
	runeType := "e"
	fmt.Printf("runeType: %v, %T\n", runeType, runeType)

	runeToByte := []byte(runeType)
	fmt.Printf("runeToByte: %v, %T\n", runeToByte, runeToByte)

}
