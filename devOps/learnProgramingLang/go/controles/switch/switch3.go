package main

import "fmt"

func switch3(nota float64) string {
	var n = int(nota)

	switch true {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 5 && n < 9:
		return "b"
	case n >= 0 && n < 5:
		return "c"
	default:
		return "error"
	}
}

func main() {
	fmt.Println(switch3(9.3))
}
