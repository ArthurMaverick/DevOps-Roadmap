package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numeros inteiros
	fmt.Println(1, 2, 3000)
	fmt.Println("literal inteiro é", reflect.TypeOf(22222))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := 1234
	maxNumber := math.MaxInt64
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))
	fmt.Println("O tipo de i1 é", reflect.TypeOf(maxNumber), maxNumber)

	// representa um mapeamento da tabela Unicode (int32)
	var unicode rune = 'a'
	fmt.Println("O tipo de unicode (a) é", reflect.TypeOf(unicode), unicode)

	// default é float64
	var floatNumber float32 = 12.3
	fmt.Println("O tipo de floatNumber é", reflect.TypeOf(floatNumber), floatNumber)

	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo), bo)
	fmt.Println("O tipo de bo é", reflect.TypeOf(!bo), !bo)

	// string
	s1 := "Olá meu nome é Caio"
	s2 := `
	Olá 
	meu 
	amigo`
	fmt.Println("O tipo de s1 é", reflect.TypeOf(s1), len(s1))
	fmt.Println("O tipo de s2 é", reflect.TypeOf(s2), len(s2))

	//char
	char := 'b'
	fmt.Println("O tipo de char é", reflect.TypeOf(char), char) //unicode table
}
