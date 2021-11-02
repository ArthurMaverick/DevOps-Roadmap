package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5, 6, 7} // dynamic array

	for i, numero := range numeros { // for with destructure
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, numero := range numeros { // for with destructure but ignore one value
		fmt.Printf("%d\n", numero)
	}
}
