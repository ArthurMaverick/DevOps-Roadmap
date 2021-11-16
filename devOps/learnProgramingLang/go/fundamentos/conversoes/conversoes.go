package main

import (
	"fmt"
	converter "strconv"
)

func main() {
	x := 42.0
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(rune(97))) // 97 é o valor da tabela ASCII

	// int para string
	fmt.Println("Teste " + converter.Itoa(int(42)))

	// string para int
	num, _ := converter.Atoi("423")
	fmt.Println(num - 35)

	// 1 is true
	b, _ := converter.ParseBool("1")
	if b {
		fmt.Println("Verdadeiro")
	}

}
