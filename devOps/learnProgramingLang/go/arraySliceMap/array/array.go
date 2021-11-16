package main

import "fmt"

func main() {
	// estrutura homogênea (mesmo tipo) e estatica (fixo)
	var two [2]int
	fmt.Println(two)

	two[0] = 1
	two[1] = 2
	fmt.Println("initial declaration", two)

	two[0], two[1] = 3, 4
	fmt.Println("first change ", two)

	//media

	total := 0
	for i := 0; i < len(two); i++ {
		total += two[i]
	}

	fmt.Println("média: ", total/len(two))
}
