package main

import (
	"fmt"
	m "math" // like a import {path as m} from "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo inferido pelo compilador
	area := PI * m.Pow(raio, 2)

	fmt.Printf("A área da circunferência é %f \n", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmts := fmt.Sprintf("%d %d %d %d", a, b, c, d)
	fmt.Println(fmts)

	var isOK, IsNotOk bool = true, false // declarando variaveis na mesma linha
	fmt.Println(isOK, IsNotOk)

	const one, two int64 = 1, 2
	fmt.Println(one, two)

	simple, three := "simples", 3
	fmt.Println(simple, three)

}
