package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Printf("Aprovado com nota %.2f \n", nota)
	} else if nota >= 5 && nota < 7 {
		fmt.Printf("recuperaçao com nota %.2f \n", nota)
	} else {
		fmt.Printf("reṕrovado com nota %.2f \n", nota)
	}
}

func main() {
	imprimirResultado(7.0)
	imprimirResultado(5.9)
	imprimirResultado(2.0)
	imprimirResultado(-2.0)
}
