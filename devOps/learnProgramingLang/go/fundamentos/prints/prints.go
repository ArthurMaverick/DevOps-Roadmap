package main

import "fmt"

func main() {
	fmt.Println("Olá")
	fmt.Println("estoy en bajo")

	a := fmt.Sprintf("10") //retorna sempre uma string
	fmt.Println(a)

	pi := 3.1415
	fmt.Printf("o valor é %.2f", pi)
	// %d => int
	// %f => float
	// %s => string
	// %c => char
	// %t => bool
	// %b => binario
	// %o => octal
	// %x => hexadecimal
	// %p => pointer
	// %q => string com aspas duplas
	// %v => valor
	// %T => tipo
	// %e => exponencial

	xs := []float64{98.23, 93.234, 77.421, 82.53, 83.456}
	fmt.Printf("float64: %.2f", xs)
}
