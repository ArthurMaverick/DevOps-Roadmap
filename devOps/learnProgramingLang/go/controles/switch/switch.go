package main

import "fmt"

func switch1(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		fallthrough
	case 9:
		return "Aprovado com Distinção"
	case 8, 7:
		return "Aprovado"
	case 6, 5:
		return "Recuperação"
	case 4, 3, 2, 1:
		return "Reprovado"
	default:
		return "nota invalida"
	}
}

func main() {
	fmt.Println(switch1(9.8))
	fmt.Println(switch1(6.9))
	fmt.Println(switch1(2.8))
	fmt.Println(switch1(10))
	fmt.Println(switch1(0))
}
