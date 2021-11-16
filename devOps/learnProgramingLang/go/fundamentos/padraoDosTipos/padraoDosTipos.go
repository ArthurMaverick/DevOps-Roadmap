package main

import (
	"fmt"
)

func main() {
	var a int
	var b float64
	var c string
	var d bool
	var f []string
	var g *int // ponteiro para int

	fmt.Printf("%v %v %q %v %q %v", a, b, c, d, f, g)
}
